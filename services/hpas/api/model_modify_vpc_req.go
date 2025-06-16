package api

type ModifyInstanceVpcRequest struct {
	HpasId            string   `json:"hpasId"`
	SubnetId          string   `json:"subnetId"`
	PrivateIp         string   `json:"privateIp"`
	SecurityGroupType string   `json:"securityGroupType"`
	SecurityGroupIds  []string `json:"securityGroupIds"`
}
