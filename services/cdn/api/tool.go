package api

import "github.com/baidubce/bce-sdk-go/bce"

// IpInfo defined a struct for IP info
type IpInfo struct {
	IP      string `json:"ip"`
	IsCdnIp bool   `json:"cdnIP"`
	Isp     string `json:"isp"`
	Region  string `json:"region"`
}

// BackOriginNode defined a struct for CDN node which may request the origin server if cache missed.
type BackOriginNode struct {
	CNNAME   string `json:"cnname"`
	IP       string `json:"ip"`
	Level    string `json:"level"`
	City     string `json:"city"`
	Province string `json:"province"`
	ISP      string `json:"isp"`
}

// GetIpInfo - retrieves information about the specified IP
// For details, please refer https://cloud.baidu.com/doc/CDN/s/8jwvyeunq#%E5%8D%95%E4%B8%AAip%E6%9F%A5%E8%AF%A2%E6%8E%A5%E5%8F%A3
//
// PARAMS:
//     - cli: the client agent can execute sending request
//     - ip: the specified ip addr
//     - action: the action for operating the ip addr
// RETURNS:
//     - *IpInfo: the information about the specified ip addr
//     - error: nil if success otherwise the specific error
func GetIpInfo(cli bce.Client, ip string, action string) (*IpInfo, error) {
	params := map[string]string{
		"ip":     ip,
		"action": action,
	}

	respObj := &IpInfo{}
	err := httpRequest(cli, "GET", "/v2/utils", params, nil, respObj)
	if err != nil {
		return nil, err
	}
	respObj.IP = ip
	return respObj, nil
}

// GetIpListInfo - retrieves information about the specified IP list
// For details, please refer https://cloud.baidu.com/doc/CDN/s/8jwvyeunq#ip-list-%E6%9F%A5%E8%AF%A2%E6%8E%A5%E5%8F%A3
//
// PARAMS:
//     - cli: the client agent can execute sending request
//     - ips: IP list
//     - action: the action for operating the ip addr
// RETURNS:
//     - []IpInfo: IP list's information
//     - error: nil if success otherwise the specific error
func GetIpListInfo(cli bce.Client, ips []string, action string) ([]IpInfo, error) {
	reqObj := map[string]interface{}{
		"ips":    ips,
		"action": action,
	}

	var respObj []IpInfo
	err := httpRequest(cli, "POST", "/v2/utils/ips", nil, reqObj, &respObj)
	if err != nil {
		return nil, err
	}
	return respObj, nil
}

// GetBackOriginNodes - get CDN nodes that may request the origin server if cache missed
// For details, please refer https://cloud.baidu.com/doc/CDN/s/8jwvyeunq#%E7%99%BE%E5%BA%A6%E5%9B%9E%E6%BA%90ip%E5%9C%B0%E5%9D%80%E6%AE%B5%E6%9F%A5%E8%AF%A2%E6%8E%A5%E5%8F%A3
//
// PARAMS:
//     - cli: the client agent can execute sending request
// RETURNS:
//     - []BackOriginNode: list of CDN node
//     - error: nil if success otherwise the specific error
func GetBackOriginNodes(cli bce.Client) ([]BackOriginNode, error) {
	respObj := &struct {
		Status          int              `json:"status"`
		BackOriginNodes []BackOriginNode `json:"details"`
	}{}
	err := httpRequest(cli, "GET", "/v2/nodes/list", nil, nil, &respObj)
	if err != nil {
		return nil, err
	}
	return respObj.BackOriginNodes, nil
}
