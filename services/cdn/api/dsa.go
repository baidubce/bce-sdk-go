package api

import (
	"fmt"

	"github.com/baidubce/bce-sdk-go/bce"
)

// DSARule configure the rule of speed up which asserts.
// Valid rule types are:
// - suffix: the file suffix in the URL, such as http://test.com/1.php has the suffix .php, or you can call it as extension.
// - path: the directory about URL, such as http://test.com/web/1.php, located at two directories, "/" and the "/web".
// - exactPath: the absolute URL path.
// - method: the HTTP request method, now we only support 5 methods, they are GET, POST, PUT, DELETE and OPTIONS.
//
// The value is related to the rule type, for example, we can configure "method" type rule with 3 methods,
// the value of it are "GET;POST;PUT" which come from 3 method string connect to each other by semicolon.
// Here shows how to construct a DSARule object of above example:
// 	var rule = &DSARule{
//		Type:  "method",
//		Value: "GET;POST;PUT",
//	}
type DSARule struct {
	Type  string `json:"type"`
	Value string `json:"value"`
}

// DSADomain defined a struct about the specified domain's DSA setting
type DSADomain struct {
	Domain     string    `json:"domain"`
	Rules      []DSARule `json:"rules"`
	ModifyTime string    `json:"modifyTime"`
	Comment    string    `json:"comment"`
}

// DSAConfig defined a struct for DSA configuration
type DSAConfig struct {
	Enabled bool      `json:"enabled"`
	Rules   []DSARule `json:"rules"`
	Comment string    `json:"comment,omitempty"`
}

func setDsa(cli bce.Client, action string) error {
	err := httpRequest(cli, "PUT", "/v2/dsa", nil, &struct {
		Action string `json:"action"`
	}{
		Action: action,
	}, nil)
	if err != nil {
		return err
	}

	return nil
}

// EnableDsa - enable DSA
// For details, please refer https://cloud.baidu.com/doc/CDN/s/7jwvyf1h5
//
// PARAMS:
//     - cli: the client agent which can perform sending request
// RETURNS:
//     - error: nil if success otherwise the specific error
func EnableDsa(cli bce.Client) error {
	return setDsa(cli, "enable")
}

// DisableDsa - disable DSA
// For details, please refer https://cloud.baidu.com/doc/CDN/s/7jwvyf1h5
//
// PARAMS:
//     - cli: the client agent which can perform sending request
// RETURNS:
//     - error: nil if success otherwise the specific error
func DisableDsa(cli bce.Client) error {
	return setDsa(cli, "disable")
}

// ListDsaDomains - retrieve all the domains in DSA service
// For details, please refer https://cloud.baidu.com/doc/CDN/s/5jwvyf1sq
//
// PARAMS:
//     - cli: the client agent which can perform sending request
// RETURNS:
//     - []DSADomain: the details about DSA domains
//     - error: nil if success otherwise the specific error
func ListDsaDomains(cli bce.Client) ([]DSADomain, error) {
	respObj := &struct {
		Domains []DSADomain `json:"domains"`
	}{}

	err := httpRequest(cli, "GET", "/v2/dsa/domain", nil, nil, respObj)
	if err != nil {
		return nil, err
	}

	return respObj.Domains, nil
}

// SetDsaConfig - set the DSA configuration
// For details, please refer https://cloud.baidu.com/doc/CDN/s/0jwvyf26d
//
// PARAMS:
//     - cli: the client agent which can perform sending request
//     - domain: the specified domain
//     - dsaConfig: the specified configuration for the specified domain
// RETURNS:
//     - error: nil if success otherwise the specific error
func SetDsaConfig(cli bce.Client, domain string, dsaConfig *DSAConfig) error {
	urlPath := fmt.Sprintf("/v2/domain/%s/config", domain)
	params := map[string]string{
		"dsa": "",
	}

	err := httpRequest(cli, "PUT", urlPath, params, &struct {
		Dsa *DSAConfig `json:"dsa"`
	}{
		Dsa: dsaConfig,
	}, nil)
	if err != nil {
		return err
	}

	return nil
}
