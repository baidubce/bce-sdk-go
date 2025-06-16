package api

type ModifyInstancesSubnetRequest struct {
	HpasIds   []string `json:"hpasIds"`
	SubnetId  string   `json:"subnetId"`
	PrivateIp string   `json:"privateIp"`
}
