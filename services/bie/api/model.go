/*
 * Copyright 2017 Baidu, Inc.
 *
 * Licensed under the Apache License, Version 2.0 (the "License"); you may not use this file
 * except in compliance with the License. You may obtain a copy of the License at
 *
 * http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software distributed under the
 * License is distributed on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND,
 * either express or implied. See the License for the specific language governing permissions
 * and limitations under the License.
 */

// model.go - definitions of the request arguments and results data structure model

package api

type Group struct {
	GroupUuid   string `json:"groupUuid"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Platform    string `json:"platform"`
	CreateTime  string `json:"createTime"`
	UpdateTime  string `json:"updateTime"`
}

type ListGroupReq struct {
	PageNo           int     `json:"pageNo"`
	PageSize         int     `json:"pageSize"`
	Name             string `json:"name"`
}

type ListGroupResult struct {
	TotalCountByUser int     `json:"totalCountByUser"`
	TotalCount       int     `json:"totalCount"`
	PageNo           int     `json:"pageNo"`
	PageSize         int     `json:"pageSize"`
	Result           []Group `json:"result"`
}

type CreateGroupReq struct {
	Name        string `json:"name"`
	Description string `json:"description"`
	AuthType    string `json:"authType"`
	Platform    string `json:"platform"`
}

type CoreAuth struct {
	AuthType    string `json:"authType"`
	Password    string `json:"password"`
	PrivateKey  string `json:"privateKey"`
	Certificate string `json:"certificate"`
	UserName    string `json:"username"`
	Tcp         string `json:"tcp"`
	Hostname    string `json:"hostname"`
	Ssl         string `json:"ssl"`
	Wss         string `json:"wss"`
}

type CoreInfo struct {
	DeviceUuid  string   `json:"deviceUuid"`
	Name        string   `json:"name"`
	Description string   `json:"description"`
	AuthType    string   `json:"authType"`
	Platform    string   `json:"platform"`
	BinVersion  string   `json:"binVersion"`
	CoreAuth    CoreAuth `json:"coreAuth"`
	DownloadUrl string   `json:"downloadUrl"`
	CreateTime  string   `json:"createTime"`
	UpdateTime  string   `json:"updateTime"`
}

type CoreResult struct {
	DeviceUuid  string `json:"deviceUuid"`
	Name        string `json:"name"`
	ConfVersion string `json:"confVersion"`
	Report      string `json:"report"`
	AuthType    string `json:"authType"`
	Description string `json:"description"`
	Platform    string `json:"platform"`
	CreateTime  string `json:"createTime"`
	UpdateTime  string `json:"updateTime"`
}

type ListCoreResult struct {
	Result []CoreResult `json:"result"`
}

type CreateGroupResult struct {
	GroupInfo Group    `json:"groupInfo"`
	CoreInfo  CoreInfo `json:"coreInfo"`
}

type EditGroupReq struct {
	Name        string `json:"name"`
	Description string `json:"description"`
}

type CoreStatus struct {
	IsCoreOnline bool `json:"isCoreOnline"`
}

type Config struct {
	Uuid           string `json:"uuid"`
	Status         string `json:"status"`
	Version        string `json:"version"`
	Description    string `json:"description"`
	LastDeployTime string `json:"lastDeployTime"`
	CreateTime     string `json:"createTime"`
}

type ListConfigResult struct {
	TotalCount int      `json:"totalCount"`
	Result     []Config `json:"result"`
	Language   string   `json:"language"`
	PageNo     int      `json:"pageNo"`
	PageSize   int      `json:"pageSize"`
	Region     bool     `json:"region"`
}

type ListConfigReq struct {
	Status   string `json:"status"`
	PageNo   int    `json:"pageNo"`
	PageSize int    `json:"pageSize"`
}

type CfgService struct {
	Uuid           string   `json:"uuid"`
	Name           string   `json:"name"`
	ModuleUuid     string   `json:"moduleUuid"`
	ModuleImage    string   `json:"moduleImage"`
	ModuleName     string   `json:"moduleName"`
	ModuleTags     string   `json:"moduleTags"`
	ModuleCategory string   `json:"moduleCategory"`
	Mounts         []string `json:"mounts"`
	Description    string   `json:"description"`
	CreateTime     string   `json:"createTime"`
}

type CfgVolume struct {
	Name       string `json:"name"`
	Version    string `json:"version"`
	HostPath   string `json:"hostPath"`
	Deletable  bool   `json:"deletable"`
	CreateTime string `json:"createTime"`
}

type CfgResult struct {
	Uuid           string       `json:"uuid"`
	Status         string       `json:"status"`
	Version        string       `json:"version"`
	Description    string       `json:"description"`
	LastDeployTime string         `json:"lastDeployTime"`
	CreateTime     string       `json:"createTime"`
	ConfigServices []CfgService `json:"configServices"`
	ConfigVolumes  []CfgVolume  `json:"configVolumes"`
}

type CoreidVersion struct {
	Coreid  string `json:"coreid"`
	Version string `json:"version"`
}

type CfgPubBody struct {
	Description string `json:"description"`
}

type CfgDownloadReq struct {
	Coreid  string `json:"coreid"`
	Version string `json:"version"`
	WithBin bool   `json:"withBin"`
}

type CfgDownloadResult struct {
	Url string `json:"url"`
	Md5 string `json:"md5"`
}

type CfgRestart struct {
	Retry map[string]interface{} `json:"restart"`
}

type CreateServiceReq struct {
	Name           string                   `json:"name"`
	ModuleUuid     string                   `json:"moduleUuid"`
	ModuleCategory string                   `json:"moduleCategory"`
	Description    string                   `json:"description"`
	Replica        int                      `json:"replica"`
	Mounts         []map[string]interface{} `json:"mounts"`
	Ports          []string                 `json:"ports"`
	Args           []string                 `json:"args"`
	Env            map[string]string        `json:"env"`
	Devs           []string                 `json:"devs"`
	Restart        map[string]interface{}   `json:"restart"`
	Resources      map[string]interface{}   `json:"resources"`
}

type ServiceResult struct {
	Uuid           string                   `json:"uuid"`
	Name           string                   `json:"name"`
	ModuleImage    string                   `json:"moduleImage"`
	ModuleName     string                   `json:"moduleName"`
	ModuleTags     string                   `json:"moduleTags"`
	ModuleCategory string                   `json:"moduleCategory"`
	Description    string                   `json:"description"`
	Mounts         []map[string]interface{} `json:"mounts"`
	Ports          []string                 `json:"ports"`
	Args           []string                 `json:"args"`
	Env            map[string]string        `json:"env"`
	Devs           []string                 `json:"devs"`
	Restart        map[string]interface{}   `json:"restart"`
	Resources      map[string]interface{}   `json:"resources"`
	CreateTime     string                   `json:"createTime"`
}

type EditServiceReq struct {
	Description string                   `json:"description"`
	Replica     int                      `json:"replica"`
	Mounts      []map[string]interface{} `json:"mounts"`
	Ports       []string                 `json:"ports"`
	Args        []string                 `json:"args"`
	Env         map[string]string        `json:"env"`
	Devs        []string                 `json:"devs"`
	Restart     map[string]interface{}   `json:"restart"`
	Resources   map[string]interface{}   `json:"resources"`
}

type IdVerName struct {
	Coreid  string `json:"coreid"`
	Version string `json:"version"`
	Name    string `json:"name"`
}

type NameVersion struct {
	Name    string `json:"name"`
	Version string `json:"version"`
}

type VolumeOpReq struct {
	Attach []NameVersion `json:"attach"`
	Detach []NameVersion `json:"detach"`
}

// Volume
type VolTemplate struct {
	Name     string   `json:"name"`
	Code     string   `json:"code"`
	Category string   `json:"category"`
	Tags     []string `json:"tags"`
}

type ListVolTemplate struct {
	Result []VolTemplate `json:"result"`
}

type CreateVolReq struct {
	Name         string   `json:"name"`
	TemplateCode string   `json:"templateCode"`
	Description  string   `json:"description"`
	Tags         []string `json:"tags"`
	HostPath	string `json:"hostPath"`
}

type VolFile struct {
	FileName  string `json:"fileName"`
	Type      string `json:"type"`
	Viewable  bool   `json:"viewable"`
	Editable  bool   `json:"editable"`
	Deletable bool   `json:"deletable"`
}

type VolumeResult struct {
	Uuid        string        `json:"uuid"`
	Name        string        `json:"name"`
	Tags        []string      `json:"tags"`
	HostPath    string        `json:"hostPath"`
	Files       []VolFile     `json:"files"`
	Description string        `json:"description"`
	Version     string        `json:"version"`
	Template    VolTemplate `json:"template"`
	NewFile     bool          `json:"newFile"`
	CleanFile   bool          `json:"cleanFile"`
	CreateTime  string        `json:"createTime"`
}

type ListVolumeReq struct {
	PageNo   int    `json:"pageNo"`
	PageSize int    `json:"pageSize"`
	Tag      string    `json:"tag"`
	Name     string `json:"name"`
}

type ListVolumeResult struct {
	TotalCount int            `json:"totalCount"`
	PageNo     int            `json:"pageNo"`
	PageSize   int            `json:"pageSize"`
	Result     []VolumeResult `json:"result"`
}

type EditVolumeReq struct {
	Description string   `json:"description"`
	Tags        []string `json:"tags"`
	HostPath    string   `json:"hostPath"`
}

type ListVolumeVerResult struct {
	Result []VolumeResult `json:"result"`
}

type VolDownloadResult struct {
	Url string `json:"url"`
}

type CreateVolFileReq struct {
	Content  string `json:"content"`
	FileName string `json:"fileName"`
}

type GetVolFileReq struct {
	Name  string `json:"name"`
	Version  string `json:"version"`
	FileName string `json:"fileName"`
}

type Name2 struct {
	Name     string `json:"name"`
	FileName string `json:"fileName"`
}

type EditVolFileReq struct {
	Content string `json:"content"`
}

type ListVolCoreResult struct {
	PageNo   int `json:"pageNo"`
	PageSize int `json:"pageSize"`
	Result   []struct {
		ConfigVersion string `json:"configVersion"`
		DeviceName    string `json:"deviceName"`
		DeviceUUID    string `json:"deviceUuid"`
		GroupUUID     string `json:"groupUuid"`
		VolumeName    string `json:"volumeName"`
		VolumeVersion string `json:"volumeVersion"`
	} `json:"result"`
	TotalCount int `json:"totalCount"`
}

type ListVolCoreReq struct {
	Name     string `json:"name"`
	PageNo   int    `json:"pageNo"`
	PageSize int    `json:"pageSize"`
}

type EditCoreVolVerReq struct {
	Jobs []EditCoreVolVerEntry `json:"jobs"`
}

type EditCoreVolVerEntry struct {
	DeviceUUID string `json:"deviceUuid"`
	NewVersion string `json:"newVersion"`
	OldVersion string `json:"oldVersion"`
}

type ImportCfcReq struct {
	Name    string `json:"name"`
	Version string `json:"version"`
}

type ImportBosReq struct {
	BosBucket    string `json:"bosBucket"`
	BosKey string `json:"bosKey"`
}

// image
type ListImageReq struct {
	PageNo     int `json:"pageNo"`
	PageSize   int `json:"pageSize"`
	Tag        string `json:"tag"`
}

type Image struct {
	Category    string `json:"category"`
	CreateTime  string `json:"createTime"`
	Description string `json:"description"`
	Image       string `json:"image"`
	Name        string `json:"name"`
	Tags        string `json:"tags"`
	UUID        string `json:"uuid"`
	Warehouse   string `json:"warehouse"`
	Replica	struct {
		Min	string `json:"min"`
	} `json:"replica"`
}

type ListImageResult struct {
	Result     []Image `json:"result"`
	PageNo     int `json:"pageNo"`
	PageSize   int `json:"pageSize"`
	TotalCount int `json:"totalCount"`
}

type CreateImageReq struct {
	Description string `json:"description"`
	Image       string `json:"image"`
	Name        string `json:"name"`
}

type EditImageReq struct {
	Description string `json:"description"`
}