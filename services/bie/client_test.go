package bie

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"reflect"
	"runtime"
	"testing"

	"github.com/baidubce/bce-sdk-go/services/bie/api"
	"github.com/baidubce/bce-sdk-go/util/log"
)

var CLIENT *Client

type Conf struct {
	AK       string
	SK       string
	Endpoint string
}

func init() {
	fmt.Printf("init \n")
	_, f, _, _ := runtime.Caller(0)
	for i := 0; i < 6; i++ {
		f = filepath.Dir(f)
	}
	conf := filepath.Join(f, "config.json")
	fp, err := os.Open(conf)
	if err != nil {
		log.Fatal("config json file of ak/sk not given:", conf)
		os.Exit(1)
	}
	decoder := json.NewDecoder(fp)
	confObj := &Conf{}
	decoder.Decode(confObj)

	CLIENT, _ = NewClient(confObj.AK, confObj.SK, confObj.Endpoint)

	log.SetLogHandler(log.STDERR)
	log.SetLogLevel(log.DEBUG)
}

// ExpectEqual is the helper function for test each case
func ExpectEqual(alert func(format string, args ...interface{}),
	expected interface{}, actual interface{}) bool {
	expectedValue, actualValue := reflect.ValueOf(expected), reflect.ValueOf(actual)
	equal := false
	switch {
	case expected == nil && actual == nil:
		return true
	case expected != nil && actual == nil:
		equal = expectedValue.IsNil()
	case expected == nil && actual != nil:
		equal = actualValue.IsNil()
	default:
		if actualType := reflect.TypeOf(actual); actualType != nil {
			if expectedValue.IsValid() && expectedValue.Type().ConvertibleTo(actualType) {
				equal = reflect.DeepEqual(expectedValue.Convert(actualType).Interface(), actual)
			}
		}
	}
	if !equal {
		_, file, line, _ := runtime.Caller(1)
		alert("%s:%d: missmatch, expect %v but %v", file, line, expected, actual)
		return false
	}
	return true
}

//////////////////////////////////////////////
// group API
//////////////////////////////////////////////

func TestListGroup(t *testing.T) {
	listReq := &api.ListGroupReq{PageNo: 1, PageSize: 1000}
	res, err := CLIENT.ListGroup(listReq)
	ExpectEqual(t.Errorf, nil, err)
	t.Logf("%+v", res)
}

func TestGroupCRUD(t *testing.T) {
	// 1, list group first
	listReq := &api.ListGroupReq{PageNo: 1, PageSize: 1000}
	groups, err := CLIENT.ListGroup(listReq)
	ExpectEqual(t.Errorf, nil, err)

	name := "TestGroupCRUD"
	deleted := 0

	// delete core with the same name
	for _, g := range groups.Result {
		if g.Name == name {
			CLIENT.DeleteGroup(g.GroupUuid)
			deleted = 1
			break
		}
	}

	// 2, create a new group
	createReq := &api.CreateGroupReq{
		Name:        name,
		Description: "descr",
		AuthType:    "CERT",
		Platform:    "Linux-amd64",
	}

	createdGroup, err := CLIENT.CreateGroup(createReq)
	ExpectEqual(t.Errorf, nil, err)

	groups2, err := CLIENT.ListGroup(listReq)
	ExpectEqual(t.Errorf, nil, err)
	ExpectEqual(t.Errorf, groups2.TotalCount, groups.TotalCount+1)

	// 3, edit group
	newName := name + "new"
	newDesc := "descrnew"
	editReq := &api.EditGroupReq{
		Name:        newName,
		Description: newDesc,
	}

	newGroup, err := CLIENT.EditGroup(createdGroup.GroupInfo.GroupUuid, editReq)
	ExpectEqual(t.Errorf, nil, err)
	ExpectEqual(t.Errorf, newGroup.Name, newName)
	ExpectEqual(t.Errorf, newGroup.Description, newDesc)

	// 4, delete group
	err = CLIENT.DeleteGroup(newGroup.GroupUuid)
	ExpectEqual(t.Errorf, nil, err)

	// 5, list again
	groups2, _ = CLIENT.ListGroup(listReq)
	ExpectEqual(t.Errorf, groups2.TotalCount, groups.TotalCount-deleted)

}

func createNewGroup(cli *Client, name string) (*api.CreateGroupResult, error) {
	listReq := &api.ListGroupReq{PageNo: 1, PageSize: 1000}
	// 1, list group first
	groups, err := cli.ListGroup(listReq)
	if err != nil {
		return nil, err
	}

	// delete core with the same name
	for _, g := range groups.Result {
		if g.Name == name {
			CLIENT.DeleteGroup(g.GroupUuid)
			break
		}
	}

	// 2, create a new group
	createReq := &api.CreateGroupReq{
		Name:        name,
		Description: "descr",
		AuthType:    "CERT",
		Platform:    "Linux-amd64",
	}

	return cli.CreateGroup(createReq)
}

// ////////////////////////////////////////////
// core API
// ////////////////////////////////////////////
func TestCoreQuery(t *testing.T) {
	name := "TestCoreQuery"
	createdGroup, err := createNewGroup(CLIENT, name)
	ExpectEqual(t.Errorf, nil, err)

	// 3, list the core
	cores, err := CLIENT.ListCore(createdGroup.GroupInfo.GroupUuid)
	ExpectEqual(t.Errorf, nil, err)
	ExpectEqual(t.Errorf, len(cores.Result), 1)

	// 4, get the core
	core, err := CLIENT.GetCore(createdGroup.GroupInfo.GroupUuid, cores.Result[0].DeviceUuid)
	ExpectEqual(t.Errorf, nil, err)

	// 5, renew core auth
	_, err = CLIENT.RenewCoreAuth(core.DeviceUuid)
	ExpectEqual(t.Errorf, nil, err)

	// 6, get online status
	status, err := CLIENT.GetCoreStatus(core.DeviceUuid)
	ExpectEqual(t.Errorf, nil, err)
	ExpectEqual(t.Errorf, status.IsCoreOnline, false)

	// 7, delete group
	CLIENT.DeleteGroup(createdGroup.GroupInfo.GroupUuid)
}

// ////////////////////////////////////////////
// Config API
// ////////////////////////////////////////////
func TestConfigCrudAndOp(t *testing.T) {
	// 1, create a group and core
	name := "TestConfigCrudAndOp"
	newGroup, err := createNewGroup(CLIENT, name)
	ExpectEqual(t.Errorf, nil, err)

	// 2, create a config
	text := []byte(`{
		"moduleUuid": "b1f91fbe6fe54d2eaf70ef0025f1c3c2",
		"moduleCategory": "SYSTEM",
		"name": "first016",
		"description": "test",
		"replica": 1,
		"mounts1": [
			{"name": "open-hub-001", "version": "LATEST", "readonly": false, "target": "/etc/avr"}
		],
		"ports": ["100:101", "8000:8080"],
		"args": ["abc", "-v", "1.1.1"],
		"env": {
			"name": "wangxiaochen",
			"pwd": "abcabc"
		},
		"restart": {
			"retry": {
				"max": 1
			},
			"policy": "always",
			"backoff": {
				"min": "10m",
				"max": "10m",
				"factor": 12
			}
		},
		"resources": {
			"cpu": {
				"cpus": 2,
				"setcpus": "2"
			},
			"pids": {
				"limit": 20
			},
			"memory": {
				"limit": "10m",
				"swap": "20m"
			}
		},
		"devs": ["/dev/01:/dev/02", "/dev/cat1:/dev/tom1"]
	}`)

	var req api.CreateServiceReq
	err = json.Unmarshal(text, &req)
	ExpectEqual(t.Errorf, nil, err)

	idver := &api.CoreidVersion{Coreid: newGroup.CoreInfo.DeviceUuid, Version: "LATEST"}
	newConf, err := CLIENT.CreateService(idver, &req)
	ExpectEqual(t.Errorf, nil, err)
	ExpectEqual(t.Errorf, newConf.Name, "first016")

	// 3, list config
	listReq := &api.ListConfigReq{}
	confList, err := CLIENT.ListConfig(idver.Coreid, listReq)
	ExpectEqual(t.Errorf, nil, err)
	ExpectEqual(t.Errorf, confList.TotalCount, 2)
	// 4, get config
	confGot, err := CLIENT.GetConfig(idver.Coreid, idver.Version)
	ExpectEqual(t.Errorf, nil, err)
	ExpectEqual(t.Errorf, confGot.ConfigServices[0].ModuleCategory, "SYSTEM")

	// 5, pub config
	pubRet, err := CLIENT.PubConfig(idver, &api.CfgPubBody{Description: "test"})
	ExpectEqual(t.Errorf, nil, err)
	ExpectEqual(t.Errorf, pubRet.Status, "DONE")

	// 6, download url
	_, err = CLIENT.DownloadConfig(&api.CfgDownloadReq{
		Coreid: idver.Coreid, Version: pubRet.Version, WithBin: true})
	ExpectEqual(t.Errorf, nil, err)

	// 7, deploy
	err = CLIENT.DeployConfig(idver.Coreid, pubRet.Version)
	ExpectEqual(t.Errorf, false, err == nil)
}

// ////////////////////////////////////////////
// Volume API
// ////////////////////////////////////////////
func TestVolumeCrudAndOp(t *testing.T) {
	// 0, list volume template
	templates, err := CLIENT.ListVolumeTpl()
	ExpectEqual(t.Errorf, nil, err)
	ExpectEqual(t.Errorf, true, len(templates.Result) > 0)

	// 1, create volume
	name := "TestVolumeCrudAndOp"
	hostPath := "/var/local"
	createReq := &api.CreateVolReq{
		Name:         name,
		Tags:         []string{"tag1", "tag2"},
		TemplateCode: templates.Result[0].Code,
		HostPath:     hostPath,
		Description:  "desc",
	}
	createRet, err := CLIENT.CreateVolume(createReq)
	ExpectEqual(t.Errorf, nil, err)
	ExpectEqual(t.Errorf, name, createRet.Name)
	ExpectEqual(t.Errorf, templates.Result[0].Code, createRet.Template.Code)

	// 2, get volume
	volumeGot, err := CLIENT.GetVolume(name)
	ExpectEqual(t.Errorf, nil, err)
	ExpectEqual(t.Errorf, name, volumeGot.Name)
	ExpectEqual(t.Errorf, templates.Result[0].Code, volumeGot.Template.Code)

	// 3, list volume
	list, err := CLIENT.ListVolume(&api.ListVolumeReq{})
	ExpectEqual(t.Errorf, nil, err)
	ExpectEqual(t.Errorf, true, list.TotalCount > 0)

	// 4, update
	newDesc := "new description"
	newTags := []string{"tag3", "tag4", "tag5", "tag6"}
	newHostpath := hostPath + "/new"
	updateReq := &api.EditVolumeReq{
		Tags: newTags, Description: newDesc, HostPath: newHostpath}
	updateRet, err := CLIENT.EditVolume(name, updateReq)
	ExpectEqual(t.Errorf, nil, err)
	ExpectEqual(t.Errorf, newDesc, updateRet.Description)
	ExpectEqual(t.Errorf, true, len(updateRet.Tags) >= 4)

	// 5, query version list
	verListRet, err := CLIENT.ListVolumeVer(&api.NameVersion{Name: name, Version: updateRet.Version})
	ExpectEqual(t.Errorf, nil, err)
	ExpectEqual(t.Errorf, true, len(verListRet.Result) > 0)

	// 6, publish new version
	pubRet, err := CLIENT.PubVolumeVer(name)
	ExpectEqual(t.Errorf, nil, err)
	ExpectEqual(t.Errorf, false, pubRet.Version == updateRet.Version)
	verListRet2, err := CLIENT.ListVolumeVer(&api.NameVersion{Name: name})
	ExpectEqual(t.Errorf, nil, err)
	ExpectEqual(t.Errorf, true, len(verListRet2.Result) > 1)

	// 7, download url
	downRet, err := CLIENT.DownloadVolVer(name, pubRet.Version)
	ExpectEqual(t.Errorf, nil, err)
	ExpectEqual(t.Errorf, true, len(downRet.Url) > 0)

	// 8, delete
	err = CLIENT.DeleteVolume(name)
	ExpectEqual(t.Errorf, nil, err)
}

func TestVolumeFileCrud(t *testing.T) {
	// 0, list volume template
	templates, err := CLIENT.ListVolumeTpl()
	ExpectEqual(t.Errorf, nil, err)
	ExpectEqual(t.Errorf, true, len(templates.Result) > 0)
	t.Logf("%+v", templates.Result[0])

	var tmplate api.VolTemplate
	for _, t := range templates.Result {
		if t.Code == "CUSTOMIZE" {
			tmplate = t
		}
	}

	// 1, create volume

	// 1.1 delete old volume
	name := "TestVolumeFileCrud"
	listVol, err := CLIENT.ListVolume(&api.ListVolumeReq{})
	for _, v := range listVol.Result {
		if v.Name == name {
			CLIENT.DeleteVolume(name)
		}
	}

	hostPath := "/var/local"
	createReq := &api.CreateVolReq{
		Name:         name,
		Tags:         []string{"tag1", "tag2"},
		TemplateCode: tmplate.Code,
		HostPath:     hostPath,
		Description:  "desc",
	}
	createVolRet, err := CLIENT.CreateVolume(createReq)
	ExpectEqual(t.Errorf, nil, err)
	t.Logf("%+v", createVolRet)

	// 2, add file
	filename := "file1"
	content := "{\"a\":1}"
	createFileReq := &api.CreateVolFileReq{
		FileName: filename,
		Content:  content,
	}
	createFileRet, err := CLIENT.CreateVolFile(name, createFileReq)
	t.Logf("%+v", createFileRet)
	ExpectEqual(t.Errorf, nil, err)
	ExpectEqual(t.Errorf, filename, createFileRet.FileName)
	ExpectEqual(t.Errorf, content, createFileRet.Content)

	// 3, get file
	getFileReq := &api.GetVolFileReq{
		Name:     name,
		FileName: filename,
		Version:  createVolRet.Version,
	}
	getFileRet, err := CLIENT.GetVolumeFile(getFileReq)
	ExpectEqual(t.Errorf, nil, err)
	ExpectEqual(t.Errorf, filename, getFileRet.FileName)
	ExpectEqual(t.Errorf, content, getFileRet.Content)

	// 4, update file
	newContent := content + "new"
	editReq := &api.EditVolFileReq{Content: newContent}
	editRet, err := CLIENT.EditVolumeFile(&api.Name2{Name: name, FileName: filename}, editReq)
	ExpectEqual(t.Errorf, nil, err)
	ExpectEqual(t.Errorf, filename, editRet.FileName)
	ExpectEqual(t.Errorf, newContent, editRet.Content)

	// 5, get file
	getFileRet, err = CLIENT.GetVolumeFile(getFileReq)
	ExpectEqual(t.Errorf, nil, err)
	ExpectEqual(t.Errorf, filename, getFileRet.FileName)
	ExpectEqual(t.Errorf, newContent, getFileRet.Content)

	// 6, delete file
	err = CLIENT.DeleteVolFile(name, filename)
	ExpectEqual(t.Errorf, nil, err)

	// 7, delete volume
	CLIENT.DeleteVolume(name)
}

func TestVolumeMisc(t *testing.T) {
	// 0, list volume template
	templates, err := CLIENT.ListVolumeTpl()
	ExpectEqual(t.Errorf, nil, err)
	ExpectEqual(t.Errorf, true, len(templates.Result) > 0)

	var tmplate api.VolTemplate
	for _, tmp := range templates.Result {
		if tmp.Code == "CFC" {
			tmplate = tmp
			break
		}
	}

	// 1, create volume

	// 1.1 delete old volume
	name := "TestVolumeMisc"
	listVol, err := CLIENT.ListVolume(&api.ListVolumeReq{})
	for _, v := range listVol.Result {
		if v.Name == name {
			CLIENT.DeleteVolume(name)
		}
	}

	hostPath := "/var/local"
	createReq := &api.CreateVolReq{
		Name:         name,
		Tags:         []string{"tag1", "tag2"},
		TemplateCode: tmplate.Code,
		HostPath:     hostPath,
		Description:  "desc",
	}
	createVolRet, err := CLIENT.CreateVolume(createReq)
	ExpectEqual(t.Errorf, nil, err)
	t.Logf("%+v", createVolRet)

	// 2, import cfc
	importCfcReq := &api.ImportCfcReq{
		Name:    "non-exist-name",
		Version: "1",
	}

	err = CLIENT.ImportCfc(name, importCfcReq)
	ExpectEqual(t.Errorf, true, err != nil)
	fmt.Println(err)

	// 3, get core info
	listCoreReq := &api.ListVolCoreReq{
		Name: name,
	}
	listCoreRet, err := CLIENT.ListVolCore(listCoreReq)
	ExpectEqual(t.Errorf, nil, err)
	ExpectEqual(t.Errorf, 0, listCoreRet.TotalCount)

	// 4, publish a new version and update core col version
	// publish colume version
	pubRet, err := CLIENT.PubVolumeVer(name)
	ExpectEqual(t.Errorf, nil, err)

	entry := api.EditCoreVolVerEntry{
		DeviceUUID: "",
		OldVersion: "V2",
		NewVersion: pubRet.Version,
	}
	editCoreVerReq := &api.EditCoreVolVerReq{
		Jobs: []api.EditCoreVolVerEntry{entry},
	}

	err = CLIENT.EditCoreVolVer(name, editCoreVerReq)
	ExpectEqual(t.Errorf, true, err != nil)

	// 5, delete volume
	CLIENT.DeleteVolume(name)
}

func TestVolumeImportBos(t *testing.T) {
	// 0, list volume template
	templates, err := CLIENT.ListVolumeTpl()
	ExpectEqual(t.Errorf, nil, err)
	ExpectEqual(t.Errorf, true, len(templates.Result) > 0)

	var tmplate api.VolTemplate
	for _, tmp := range templates.Result {
		if tmp.Code == "BOS" {
			tmplate = tmp
			break
		}
	}

	// 1, create volume

	// 1.1 delete old volume
	name := "TestVolumeImportBos"
	listVol, err := CLIENT.ListVolume(&api.ListVolumeReq{})
	for _, v := range listVol.Result {
		if v.Name == name {
			CLIENT.DeleteVolume(name)
		}
	}

	hostPath := "/var/local"
	createReq := &api.CreateVolReq{
		Name:         name,
		Tags:         []string{"tag1", "tag2"},
		TemplateCode: tmplate.Code,
		HostPath:     hostPath,
		Description:  "desc",
	}
	createVolRet, err := CLIENT.CreateVolume(createReq)
	ExpectEqual(t.Errorf, nil, err)
	t.Logf("%+v", createVolRet)

	// 2, import cfc
	importBosReq := &api.ImportBosReq{
		BosBucket: "non-exist-bucket",
		BosKey:    "non-exist-key",
	}

	err = CLIENT.ImportBos(name, importBosReq)
	ExpectEqual(t.Errorf, true, err != nil)
	fmt.Println(err)

	// 3, delete volume
	CLIENT.DeleteVolume(name)
}

func TestDockerImages(t *testing.T) {
	listReq := &api.ListImageReq{PageNo: 1, PageSize: 1000}
	list, err := CLIENT.ListImageSys(listReq)
	ExpectEqual(t.Errorf, nil, err)
	ExpectEqual(t.Errorf, true, list.TotalCount >= 0)
	_, err = CLIENT.GetImageSys("notexistimageid")
	ExpectEqual(t.Errorf, true, err != nil)
	resultSize := listReq.PageSize
	if resultSize > list.TotalCount {
		resultSize = list.TotalCount
	}
	ExpectEqual(t.Errorf, resultSize, len(list.Result))

	list, err = CLIENT.ListImageUser(listReq)
	ExpectEqual(t.Errorf, nil, err)
	ExpectEqual(t.Errorf, true, list.TotalCount >= 0)
	cnt1 := list.TotalCount
	_, err = CLIENT.GetImageUser("notexistimageid")
	ExpectEqual(t.Errorf, true, err != nil)

	name := "TestDockerImages"
	// 0, delete user iamge left by previous test
	for _, m := range list.Result {
		if name == m.Name {
			CLIENT.DeleteImageUser(m.UUID)
		}
	}

	//  1, create a new user image
	url := "hub.c.163.com/library/nginx:latest"
	desc := "desc"
	createReq := &api.CreateImageReq{Name: name, Description: desc, Image: url}
	img, err := CLIENT.CreateImageUser(createReq)
	ExpectEqual(t.Errorf, nil, err)
	ExpectEqual(t.Errorf, name, img.Name)
	ExpectEqual(t.Errorf, desc, img.Description)
	ExpectEqual(t.Errorf, url, img.Image)

	// 2, list user image again, and compare the totalCount
	list, err = CLIENT.ListImageUser(listReq)
	ExpectEqual(t.Errorf, nil, err)
	ExpectEqual(t.Errorf, cnt1+1, list.TotalCount)

	// 2, delete the create image
	err = CLIENT.DeleteImageUser(img.UUID)
	ExpectEqual(t.Errorf, nil, err)
}
