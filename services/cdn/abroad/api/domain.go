package api

import (
	"errors"
	"fmt"

	"github.com/baidubce/bce-sdk-go/bce"
	"github.com/baidubce/bce-sdk-go/model"
)

// DomainStatus defined a struct for domain info,
// the status only include "RUNNING", "OPERATING" and "STOPPED",
// the other status is undefined
type DomainStatus struct {
	Domain string `json:"domain"`
	Status string `json:"status"`
}

// DomainCreatedInfo defined a struct for `CreateDomain` return
type DomainCreatedInfo struct {
	Status string `json:"status"`
	Cname  string `json:"cname"`
}

// DomainInfo defined a struct for domain information
type DomainInfo struct {
	Name string `json:"name"`
	Form string `json:"form"`
}

// ListDomains - list all domains that in ABROAD-CDN service
// For details, please refer https://cloud.baidu.com/doc/CDN-ABROAD/s/1kbsyj9m6
//
// PARAMS:
//   - cli: the client agent which can perform sending request
//   - marker: a marker is a start point of searching
//
// RETURNS:
//   - []string: domains belongs to the user
//   - string: a marker for next searching, empty if is in the end
//   - error: nil if success otherwise the specific error
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

	err := httpRequest(cli, "GET", "/v2/abroad/domains", params, nil, respObj)
	if err != nil {
		return nil, "", err
	}

	domains := make([]string, len(respObj.Domains))
	for i, domain := range respObj.Domains {
		domains[i] = domain.Name
	}

	return domains, respObj.NextMarker, nil
}

// ListDomainInfos - list all domains that in ABROAD-CDN service
// For details, please refer https://cloud.baidu.com/doc/CDN-ABROAD/s/1kbsyj9m6
//
// PARAMS:
//   - cli: the client agent which can perform sending request
//   - marker: a marker is a start point of searching
//
// RETURNS:
//   - []DomainInfo: domains belongs to the user
//   - string: a marker for next searching, empty if is in the end
//   - error: nil if success otherwise the specific error
func ListDomainInfos(cli bce.Client, marker string) ([]DomainInfo, string, error) {
	respObj := &struct {
		IsTruncated bool         `json:"isTruncated"`
		Domains     []DomainInfo `json:"domains"`
		NextMarker  string       `json:"nextMarker"`
	}{}

	params := map[string]string{}
	if marker != "" {
		params["marker"] = marker
	}

	err := httpRequest(cli, "GET", "/v2/abroad/domains", params, nil, respObj)
	if err != nil {
		return nil, "", err
	}

	return respObj.Domains, respObj.NextMarker, nil
}

// CreateDomain - add a specified domain into ABROAD-CDN service
// For details, please refer https://cloud.baidu.com/doc/CDN-ABROAD/s/ekbsyn5o5
//
// PARAMS:
//   - cli: the client agent which can perform sending request
//   - domain: the specified domain
//   - originInit: initialized data for a ABROAD-CDN domain
//   - tags: bind with the specified tags
//
// RETURNS:
//   - *DomainCreatedInfo: the details about created a ABROAD-CDN domain
//   - error: nil if success otherwise the specific error
func CreateDomain(cli bce.Client, domain string, originConfig []OriginPeer, tags []model.TagModel) (*DomainCreatedInfo, error) {
	urlPath := fmt.Sprintf("/v2/abroad/domain/%s", domain)
	respObj := &DomainCreatedInfo{}

	type Request struct {
		OriginConfig []OriginPeer     `json:"originConfig"`
		Tags         []model.TagModel `json:"tags,omitempty"`
	}

	requestObject := &Request{
		OriginConfig: originConfig,
		Tags:         tags,
	}
	err := httpRequest(cli, "POST", urlPath, nil, requestObject, respObj)
	if err != nil {
		return nil, err
	}

	return respObj, nil
}

// EnableDomain - enable a specified domain
// For details, please refer https://cloud.baidu.com/doc/CDN-ABROAD/s/Zkbsypv9b
//
// PARAMS:
//   - cli: the client agent which can perform sending request
//   - domain: the specified domain
//
// RETURNS:
//   - error: nil if success otherwise the specific error
func EnableDomain(cli bce.Client, domain string) error {
	if domain == "" {
		return errors.New("domain parameter could not be empty")
	}

	params := map[string]string{
		"enable": "",
	}
	urlPath := fmt.Sprintf("/v2/abroad/domain/%s", domain)

	err := httpRequest(cli, "PUT", urlPath, params, nil, nil)
	if err != nil {
		return err
	}

	return nil
}

// DisableDomain - disable a specified domain
// For details, please refer https://cloud.baidu.com/doc/CDN-ABROAD/s/gkbsyrdck
//
// PARAMS:
//   - cli: the client agent which can perform sending request
//   - domain: the specified domain
//
// RETURNS:
//   - error: nil if success otherwise the specific error
func DisableDomain(cli bce.Client, domain string) error {
	if domain == "" {
		return errors.New("domain parameter could not be empty")
	}

	params := map[string]string{
		"disable": "",
	}
	urlPath := fmt.Sprintf("/v2/abroad/domain/%s", domain)

	err := httpRequest(cli, "PUT", urlPath, params, nil, nil)
	if err != nil {
		return err
	}

	return nil
}

// DeleteDomain - delete a specified domain from BCE CDN system
// For details, please refer https://cloud.baidu.com/doc/CDN-ABROAD/s/4kbsytf7q
//
// PARAMS:
//   - cli: the client agent which can perform sending request
//   - domain: the specified domain
//
// RETURNS:
//   - error: nil if success otherwise the specific error
func DeleteDomain(cli bce.Client, domain string) error {
	if domain == "" {
		return errors.New("domain parameter could not be empty")
	}

	urlPath := fmt.Sprintf("/v2/abroad/domain/%s", domain)

	err := httpRequest(cli, "DELETE", urlPath, nil, nil, nil)
	if err != nil {
		return err
	}

	return nil
}
