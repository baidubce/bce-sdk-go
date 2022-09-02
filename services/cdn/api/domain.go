package api

import (
	"errors"
	"fmt"

	"github.com/baidubce/bce-sdk-go/bce"
)

// DomainStatus defined a struct for domain info,
// the status only include "RUNNING", "OPERATING" and "STOPPED",
// the other status is undefined
type DomainStatus struct {
	Domain string `json:"domain"`
	Status string `json:"status"`
}

// DomainValidInfo defined a struct for `IsValidDomain` return
type DomainValidInfo struct {
	Domain  string
	IsValid bool   `json:"isValid"`
	Message string `json:"message"`
}

// OriginPeer defined a struct for an origin server setting
type OriginPeer struct {
	Peer      string `json:"peer"`
	Host      string `json:"host,omitempty"`
	Backup    bool   `json:"backup"`
	Follow302 bool   `json:"follow302"`
	Weight    int    `json:"weight,omitempty"`
	ISP       string `json:"isp,omitempty"`
}

// OriginInit defined a struct for creating a new CDN service in `OPENCDN`
type OriginInit struct {
	Origin      []OriginPeer `json:"origin"`
	DefaultHost string       `json:"defaultHost,omitempty"`
	Form        string       `json:"form,omitempty"`
}

// DomainCreatedInfo defined a struct for `CreateDomain` return
type DomainCreatedInfo struct {
	Domain string `json:"domain"`
	Status string `json:"status"`
	Cname  string `json:"cname"`
}

// ListDomains - list all domains that in CDN service
// For details, please refer https://cloud.baidu.com/doc/CDN/s/sjwvyewt1
//
// PARAMS:
//     - cli: the client agent which can perform sending request
//     - marker: a marker is a start point of searching
// RETURNS:
//     - []string: domains belongs to the user
//     - string: a marker for next searching, empty if is in the end
//     - error: nil if success otherwise the specific error
func ListDomains(cli bce.Client, marker string) ([]string, string, error) {
	type domainInfo struct {
		Name string `json:"name"`
	}

	respObj := &struct {
		IsTruncated bool         `json:"isTruncated"`
		Domains     []domainInfo `json:"domains"`
		NextMarker  string       `json:"nextMarker"`
	}{}

	params := map[string]string{}
	if marker != "" {
		params["marker"] = marker
	}

	err := httpRequest(cli, "GET", "/v2/domain", params, nil, respObj)
	if err != nil {
		return nil, "", err
	}

	domains := make([]string, len(respObj.Domains))
	for i, domain := range respObj.Domains {
		domains[i] = domain.Name
	}

	return domains, respObj.NextMarker, nil
}

// GetDomainStatus - get domains' details
// For details, please refer https://cloud.baidu.com/doc/CDN/s/8jwvyewf1
//
// PARAMS:
//     - cli: the client agent which can perform sending request
//     - status: the specified running status, the available values are "RUNNING", "STOPPED", OPERATING or "ALL"
//     - rule: the regex matching rule
// RETURNS:
//     - []DomainStatus: domain details list
//     - error: nil if success otherwise the specific error
func GetDomainStatus(cli bce.Client, status string, rule string) ([]DomainStatus, error) {
	if status == "" {
		return nil, errors.New("domain status parameter could not be empty")
	}

	params := map[string]string{
		"status": status,
	}

	if rule != "" {
		params["rule"] = rule
	}

	respObj := &struct {
		Domains []DomainStatus `json:"domains"`
	}{}

	err := httpRequest(cli, "GET", "/v2/user/domains", params, nil, respObj)
	if err != nil {
		return nil, err
	}

	return respObj.Domains, nil
}

// IsValidDomain - check the specified domain whether it can be added to CDN service or not
// For details, please refer https://cloud.baidu.com/doc/CDN/s/qjwvyexh6
//
// PARAMS:
//     - cli: the client agent which can perform sending request
//     - domain: the specified domain
// RETURNS:
//     - *DomainValidInfo: available information about the specified domain
//     - error: nil if success otherwise the specific error
func IsValidDomain(cli bce.Client, domain string) (*DomainValidInfo, error) {
	urlPath := fmt.Sprintf("/v2/domain/%s/valid", domain)
	respObj := &DomainValidInfo{}

	err := httpRequest(cli, "GET", urlPath, nil, nil, respObj)
	if err != nil {
		return nil, err
	}

	respObj.Domain = domain
	return respObj, nil
}

// CreateDomain - add a specified domain into CDN service
// For details, please refer https://cloud.baidu.com/doc/CDN/s/gjwvyex4o
//
// PARAMS:
//     - cli: the client agent which can perform sending request
//     - domain: the specified domain
//     - originInit: initialized data for a CDN domain
// RETURNS:
//     - *DomainCreatedInfo: the details about created a CDN domain
//     - error: nil if success otherwise the specific error
func CreateDomain(cli bce.Client, domain string, originInit *OriginInit) (*DomainCreatedInfo, error) {
	urlPath := fmt.Sprintf("/v2/domain/%s", domain)
	respObj := &DomainCreatedInfo{}

	err := httpRequest(cli, "PUT", urlPath, nil, originInit, respObj)
	if err != nil {
		return nil, err
	}

	return respObj, nil
}

// EnableDomain - enable a specified domain
// For details, please refer https://cloud.baidu.com/doc/CDN/s/Jjwvyexv8
//
// PARAMS:
//     - cli: the client agent which can perform sending request
//     - domain: the specified domain
// RETURNS:
//     - error: nil if success otherwise the specific error
func EnableDomain(cli bce.Client, domain string) error {
	if domain == "" {
		return errors.New("domain parameter could not be empty")
	}

	params := map[string]string{
		"enable": "",
	}
	urlPath := fmt.Sprintf("/v2/domain/%s", domain)

	err := httpRequest(cli, "POST", urlPath, params, nil, nil)
	if err != nil {
		return err
	}

	return nil
}

// DisableDomain - disable a specified domain
// For details, please refer https://cloud.baidu.com/doc/CDN/s/9jwvyew3e
//
// PARAMS:
//     - cli: the client agent which can perform sending request
//     - domain: the specified domain
// RETURNS:
//     - error: nil if success otherwise the specific error
func DisableDomain(cli bce.Client, domain string) error {
	if domain == "" {
		return errors.New("domain parameter could not be empty")
	}

	params := map[string]string{
		"disable": "",
	}
	urlPath := fmt.Sprintf("/v2/domain/%s", domain)

	err := httpRequest(cli, "POST", urlPath, params, nil, nil)
	if err != nil {
		return err
	}

	return nil
}

// DeleteDomain - delete a specified domain from BCE CDN system
// For details, please refer https://cloud.baidu.com/doc/CDN/s/Njwvyey7f
//
// PARAMS:
//     - cli: the client agent which can perform sending request
//     - domain: the specified domain
// RETURNS:
//     - error: nil if success otherwise the specific error
func DeleteDomain(cli bce.Client, domain string) error {
	if domain == "" {
		return errors.New("domain parameter could not be empty")
	}

	urlPath := fmt.Sprintf("/v2/domain/%s", domain)

	err := httpRequest(cli, "DELETE", urlPath, nil, nil, nil)
	if err != nil {
		return err
	}

	return nil
}
