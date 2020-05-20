package api

import (
	"fmt"

	"github.com/baidubce/bce-sdk-go/bce"
)

// DSARule defined a struct for DSA urls
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
