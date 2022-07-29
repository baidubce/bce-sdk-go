package bes

type ESClusterModule struct {
	InstanceNum  int    `json:"instanceNum"`
	SlotType     string `json:"slotType"`
	DiskSlotInfo struct {
		Size int    `json:"size"`
		Type string `json:"type"`
	} `json:"diskSlotInfo,omitempty"`
	Type string `json:"type"`
}
type ESClusterRequest struct {
	Name            string             `json:"name"`
	Password        string             `json:"password"`
	SecurityGroupID string             `json:"securityGroupId"`
	SubnetUUID      string             `json:"subnetUuid"`
	AvailableZone   string             `json:"availableZone"`
	VpcID           string             `json:"vpcId"`
	IsOldPackage    bool               `json:"isOldPackage"`
	Version         string             `json:"version"`
	Modules         []*ESClusterModule `json:"modules"`
	Billing         struct {
		PaymentType string `json:"paymentType"`
		Time        int    `json:"time"`
	} `json:"billing"`
}

type DetailESClusterResponse struct {
	Result struct {
		DesireStatus string `json:"desireStatus"`
		Subnet       string `json:"subnet"`
		EsURL        string `json:"esUrl"`
		KibanaURL    string `json:"kibanaUrl"`
		KibanaEip    string `json:"kibanaEip"`
		Instances    []struct {
			InstanceID    string `json:"instanceId"`
			ModuleType    string `json:"moduleType"`
			HostIP        string `json:"hostIp"`
			ModuleVersion string `json:"moduleVersion"`
			Status        string `json:"status"`
		} `json:"instances"`
		Vpc           string `json:"vpc"`
		ClusterID     string `json:"clusterId"`
		SecurityGroup string `json:"securityGroup"`
		Modules       []struct {
			SlotDescription   string `json:"slotDescription"`
			SlotType          string `json:"slotType"`
			ActualInstanceNum int    `json:"actualInstanceNum"`
			Type              string `json:"type"`
			Version           string `json:"version"`
		} `json:"modules"`
		Billing struct {
			PaymentType string `json:"paymentType"`
		} `json:"billing"`
		Network []struct {
			SubnetID      string `json:"subnetId"`
			Subnet        string `json:"subnet"`
			AvailableZone string `json:"availableZone"`
		} `json:"network"`
		AdminUsername string `json:"adminUsername"`
		AvailableZone string `json:"availableZone"`
		ExpireTime    string `json:"expireTime"`
		ActualStatus  string `json:"actualStatus"`
		ClusterName   string `json:"clusterName"`
		VpcID         string `json:"vpcId"`
		EsEip         string `json:"esEip"`
		Region        string `json:"region"`
	} `json:"result"`
	Success bool `json:"success"`
	Status  int  `json:"status"`
}

type GetESClusterRequest struct {
	ClusterId string `json:"clusterId"`
}

type ESClusterResponse struct {
	Status  int  `json:"status"`
	Success bool `json:"success"`
	Result  struct {
		OrderId   string `json:"orderId"`
		ClusterId string `json:"clusterId"`
	} `json:"result"`
}

type DeleteESClusterResponse struct {
	Result  string `json:"result"`
	Success bool   `json:"success"`
	Status  int    `json:"status"`
}
