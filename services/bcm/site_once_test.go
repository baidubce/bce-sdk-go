package bcm

import (
	"fmt"
	"testing"
)

func TestClient_AgainExec(t *testing.T) {
	params := SiteOnceTaskRequest{
		UserID: "453bf***************c984129090dc",
		SiteID: "jspjUbhwHVotroGFKeRChlriwxlftkxH",
	}
	resp, err := bcmClient.AgainExec(&params)
	if err != nil {
		fmt.Print(err)
	}
	fmt.Printf("%+v\n", resp)
}

func TestClient_CreateSiteOnce(t *testing.T) {
	params := SiteOnceRequest{
		UserID:       "453bf***************c984129090dc",
		Address:      "www.baidu.com",
		AdvancedFlag: false,
		IpType:       "ipv4",
		Idc:          "beijing-CMNET,beijing-UNICOM,beijing-CHINANET,guangdong-CMNET,fujian-CMNET,henan-CMNET,hebei-CHINANET",
		Timeout:      60,
		ProtocolType: "HTTP",
		TaskType:     "NET_QUAILTY",
		OnceConfig: SiteOnceConfig{
			Method:      "get",
			PostContent: "",
		},
	}

	resp, err := bcmClient.CreateSiteOnce(&params)
	if err != nil {
		fmt.Print(err)
	}

	fmt.Printf("%+v\n", resp)
}

func TestClient_DeleteSiteOnceTask(t *testing.T) {

	params := SiteOnceTaskRequest{
		SiteID: "YRxqHGeoXADvgZmlrPlGYKSIRCFUHpBE",
		UserID: "453bf***************c984129090dc",
	}

	resp, err := bcmClient.DeleteSiteOnceTask(&params)
	if err != nil {
		fmt.Print(err)
	}
	fmt.Printf("%+v\n", resp)
}

func TestClient_DetailTask(t *testing.T) {

	params := SiteOnceTaskRequest{
		PageSize: 10,
		PageNo:   1,
		UserID:   "453bf***************c984129090dc",
		SiteID:   "jspjUbhwHVotroGFKeRChlriwxlftkxH",
	}

	resp, err := bcmClient.DetailTask(&params)
	if err != nil {
		fmt.Print(err)
	}
	fmt.Printf("%+v\n", resp)
}

func TestClient_GetSiteAgent(t *testing.T) {

	resp, err := bcmClient.GetSiteAgent("453bf***************c984129090dc", "")
	if err != nil {
		fmt.Print(err)
	}
	fmt.Printf("%+v\n", resp)
}

func TestClient_ListHistoryTasks(t *testing.T) {

	params := SiteOnceTaskRequest{
		PageSize: 10,
		PageNo:   1,
		UserID:   "453bf***************c984129090dc",
		URL:      "baidu",
	}

	resp, err := bcmClient.ListHistoryTasks(&params)
	if err != nil {
		fmt.Print(err)
	}
	fmt.Printf("%+v\n", resp)
}

func TestClient_ListSiteOnceTasks(t *testing.T) {

	params := SiteOnceTaskRequest{
		PageSize: 10,
		PageNo:   1,
		UserID:   "453bf***************c984129090dc",
		URL:      "baidu",
	}
	tasks, err := bcmClient.ListSiteOnceTasks(&params)
	if err != nil {
		fmt.Print(err)
	}
	fmt.Printf("%+v\n", tasks)
}

func TestClient_LoadData(t *testing.T) {

	params := SiteOnceTaskRequest{
		PageSize: 10,
		PageNo:   1,
		UserID:   "453bf***************c984129090dc",
		SiteID:   "jspjUbhwHVotroGFKeRChlriwxlftkxH",
	}

	resp, err := bcmClient.LoadData(&params)
	if err != nil {
		fmt.Print(err)
	}
	fmt.Printf("%+v\n", resp)
}
