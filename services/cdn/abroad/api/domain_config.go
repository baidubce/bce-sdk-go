package api

import (
	"errors"
	"fmt"

	"github.com/baidubce/bce-sdk-go/bce"
	"github.com/baidubce/bce-sdk-go/model"
)

const (
	HTTPOrigin  = "http"
	HTTPSOrigin = "https"
)

// DomainConfig defined a struct for total configurations.
type DomainConfig struct {
	Domain         string           `json:"domain"`
	Status         string           `json:"status"`
	Cname          string           `json:"cname"`
	Origin         []OriginPeer     `json:"originConfig"`
	CacheTTL       []CacheTTL       `json:"cacheTtl"`
	CacheFullUrl   bool             `json:"cacheFullUrl"`
	OriginHost     *string          `json:"originHost"`
	RefererACL     *RefererACL      `json:"refererACL"`
	IpACL          *IpACL           `json:"ipACL"`
	HTTPSConfig    *HTTPSConfig     `json:"https"`
	OriginProtocol string           `json:"originProtocol"`
	Tags           []model.TagModel `json:"tags"`
}

// OriginPeer defined a struct for Origin server.
type OriginPeer struct {
	Type   string `json:"type"`
	Addr   string `json:"addr"`
	Backup bool   `json:"backup"`
}

// CacheTTL defined a struct for cached rules.
type CacheTTL struct {
	Type           string `json:"type"`
	Value          string `json:"value"`
	Weight         int    `json:"weight,omitempty"`
	TTL            int    `json:"ttl"`
	OverrideOrigin bool   `json:"override_origin"`
}

// RefererACL defined a struct for Referer whitelist or blacklist to enable hotlink protection.
type RefererACL struct {
	BlackList  []string `json:"blackList"`
	WhiteList  []string `json:"whiteList"`
	AllowEmpty bool     `json:"allowEmpty"`
}

// IpACL defined a struct for IP address blacklist or whitelist.
type IpACL struct {
	BlackList []string `json:"blackList"`
	WhiteList []string `json:"whiteList"`
}

type HTTPSConfig struct {
	Enabled      bool   `json:"enabled"`
	CertId       string `json:"certId,omitempty"`
	HttpRedirect bool   `json:"httpRedirect,omitempty"`
	Http2Enabled bool   `json:"http2Enabled,omitempty"`
}

// GetDomainConfig - get the configuration for the specified domain
// For details, please refer https://cloud.baidu.com/doc/CDN-ABROAD/s/9kbsye6k8
//
// PARAMS:
//   - cli: the client agent which can perform sending request
//   - domain: the specified domain
//
// RETURNS:
//   - *DomainConfig: the configuration about the specified domain
//   - error: nil if success otherwise the specific error
func GetDomainConfig(cli bce.Client, domain string) (*DomainConfig, error) {
	urlPath := fmt.Sprintf("/v2/abroad/domain/%s/config", domain)
	var config DomainConfig
	err := httpRequest(cli, "GET", urlPath, nil, nil, &config)
	if err != nil {
		return nil, err
	}

	return &config, nil
}

// SetDomainOrigin - set the origin setting for the new
// For details, please refer https://cloud.baidu.com/doc/CDN-ABROAD/s/Gkbstcgaa
//
// PARAMS:
//   - cli: the client agent which can perform sending request
//   - domain: the specified domain
//   - origins: the origin servers
//
// RETURNS:
//   - error: nil if success otherwise the specific error
func SetDomainOrigin(cli bce.Client, domain string, origins []OriginPeer) error {
	urlPath := fmt.Sprintf("/v2/abroad/domain/%s/config", domain)
	params := map[string]string{
		"origin": "",
	}

	err := httpRequest(cli, "PUT", urlPath, params, &struct {
		Origin []OriginPeer `json:"originConfig"`
	}{
		Origin: origins,
	}, nil)

	return err
}

// SetCacheTTL - add rules to cache asserts.
// For details, please refer https://cloud.baidu.com/doc/CDN-ABROAD/s/Zkbstm0vg
//
// PARAMS:
//   - cli: the client agent which can perform sending request
//   - domain: the specified domain
//   - cacheTTLs: the cache setting list
//
// RETURNS:
//   - error: nil if success otherwise the specific error
func SetCacheTTL(cli bce.Client, domain string, cacheTTLs []CacheTTL) error {
	urlPath := fmt.Sprintf("/v2/abroad/domain/%s/config", domain)
	params := map[string]string{
		"cacheTtl": "",
	}

	err := httpRequest(cli, "PUT", urlPath, params, &struct {
		CacheTTLs []CacheTTL `json:"cacheTtl"`
	}{
		CacheTTLs: cacheTTLs,
	}, nil)
	if err != nil {
		return err
	}

	return nil
}

// SetCacheFullUrl - set the rule to calculate the cache key, set `cacheFullUrl` to true
// means the whole query(the string right after the question mark in URL) will be added to the cache key.
// For details, please refer https://cloud.baidu.com/doc/CDN-ABROAD/s/nkbsxko6t
//
// PARAMS:
//   - cli: the client agent which can perform sending request
//   - domain: the specified domain
//   - cacheFullUrl: the query part in URL being added to the cache key or not
//
// RETURNS:
//   - error: nil if success otherwise the specific error
func SetCacheFullUrl(cli bce.Client, domain string, cacheFullUrl bool) error {
	urlPath := fmt.Sprintf("/v2/abroad/domain/%s/config", domain)
	params := map[string]string{
		"cacheFullUrl": "",
	}

	err := httpRequest(cli, "PUT", urlPath, params, &struct {
		CacheFullUrl bool `json:"cacheFullUrl"`
	}{
		CacheFullUrl: cacheFullUrl,
	}, nil)
	if err != nil {
		return err
	}

	return nil
}

// SetHostToOrigin - Specify a default value for the HOST header for virtual sites in your origin server.
// For details, please refer https://cloud.baidu.com/doc/CDN-ABROAD/s/Pkbsxw8xw
//
// PARAMS:
//   - cli: the client agent which can perform sending request
//   - domain: the specified domain
//   - originHost: specified HOST header for origin server
//
// RETURNS:
//   - error: nil if success otherwise the specific error
func SetHostToOrigin(cli bce.Client, domain string, originHost string) error {
	urlPath := fmt.Sprintf("/v2/abroad/domain/%s/config", domain)
	params := map[string]string{
		"designateHostToOrigin": "",
	}

	err := httpRequest(cli, "PUT", urlPath, params, &struct {
		DesignateHostToOrigin string `json:"designateHostToOrigin"`
	}{
		DesignateHostToOrigin: originHost,
	}, nil)
	if err != nil {
		return err
	}

	return nil
}

// SetRefererACL - Set a Referer whitelist or blacklist to enable hotlink protection.
// For details, please refer https://cloud.baidu.com/doc/CDN-ABROAD/s/ekbsxow69
//
// PARAMS:
//   - cli: the client agent which can perform sending request
//   - domain: the specified domain
//   - refererACL: referer of whitelist or blacklist
//
// RETURNS:
//   - error: nil if success otherwise the specific error
func SetRefererACL(cli bce.Client, domain string, refererACL *RefererACL) error {
	urlPath := fmt.Sprintf("/v2/abroad/domain/%s/config", domain)
	params := map[string]string{
		"refererACL": "",
	}

	err := httpRequest(cli, "PUT", urlPath, params, &struct {
		RefererACL *RefererACL `json:"refererACL"`
	}{
		RefererACL: refererACL,
	}, nil)
	if err != nil {
		return err
	}

	return nil
}

// SetIpACL - Set an IP whitelist or blacklist to block or allow requests from specific IP addresses.
// For details, please refer https://cloud.baidu.com/doc/CDN-ABROAD/s/2kbsxt693
//
// PARAMS:
//   - cli: the client agent which can perform sending request
//   - domain: the specified domain
//   - ipACL: IP whitelist or blacklist
//
// RETURNS:
//   - error: nil if success otherwise the specific error
func SetIpACL(cli bce.Client, domain string, ipACL *IpACL) error {
	urlPath := fmt.Sprintf("/v2/abroad/domain/%s/config", domain)
	params := map[string]string{
		"ipACL": "",
	}

	err := httpRequest(cli, "PUT", urlPath, params, &struct {
		IpACL *IpACL `json:"ipACL"`
	}{
		IpACL: ipACL,
	}, nil)
	if err != nil {
		return err
	}

	return nil
}

// HTTPSConfigOption defined a method for setting optional configurations for HTTPS config.
type HTTPSConfigOption func(interface{})

// HTTPSConfigCertID defined a method to pass certId witch represents a certificate uploaded in BCE SSL platform:
// https://console.bce.baidu.com/cas/#/cas/purchased/common/list.
// This Option works while the HTTPS in enabled state.
func HTTPSConfigCertID(certId string) HTTPSConfigOption {
	return func(o interface{}) {
		cfg, ok := o.(*httpsConfig)
		if ok {
			cfg.certId = certId
		}
	}
}

// HTTPSConfigRedirectWith301 defined a method to enable the CDN PoPs to redirect the requests in HTTP protocol
// to the HTTPS ones, with the status 301.
// e.g. Assume that you have a CDN domain cdn.baidu.com, configured HTTPSConfigRedirectWith301, while the request
// comes just like "http://cdn.baidu.com/1.txt", the CDN Edge server would respond 301 page with Location header
// https://cdn.baidu.com/1.txt which change the scheme from "http" to "https".
// This Option works while the HTTPS in enabled state.
func HTTPSConfigRedirectWith301() HTTPSConfigOption {
	return func(o interface{}) {
		cfg, ok := o.(*httpsConfig)
		if ok {
			cfg.httpRedirect301 = true
		}
	}
}

// HTTPSConfigEnableH2 defined a method to enable HTTP2 in CDN Edge server.
// This Option works while the HTTPS in enabled state.
func HTTPSConfigEnableH2() HTTPSConfigOption {
	return func(o interface{}) {
		cfg, ok := o.(*httpsConfig)
		if ok {
			cfg.enableH2 = true
		}
	}
}

type httpsConfig struct {
	certId          string
	httpRedirect301 bool
	enableH2        bool
}

// SetHTTPSConfigWithOptions - enable or disable HTTPS.
// For details, please refer https://cloud.baidu.com/doc/CDN-ABROAD/s/ckb0fx9ea
//
// PARAMS:
//   - cli: the client agent which can perform sending request
//   - domain: the specified domain
//   - enabled: true means enable HTTPS, otherwise disable.
//   - options: more operations to configure HTTPS, the valid options are:
//     1. HTTPSConfigCertID
//     2. HTTPSConfigRedirectWith301
//     3. HTTPSConfigEnableH2
//
// RETURNS:
//   - error: nil if success otherwise the specific error
func SetHTTPSConfigWithOptions(cli bce.Client, domain string, enabled bool, options ...HTTPSConfigOption) error {
	urlPath := fmt.Sprintf("/v2/abroad/domain/%s/config", domain)
	params := map[string]string{
		"https": "",
	}

	var config httpsConfig
	for _, opt := range options {
		opt(&config)
	}

	requestObj := map[string]interface{}{
		"enabled": enabled,
	}
	if enabled {
		if config.certId == "" {
			return errors.New("try enable HTTPS but the certId is empty")
		}
		requestObj["certId"] = config.certId
		requestObj["httpRedirect"] = config.httpRedirect301
		requestObj["http2Enabled"] = config.enableH2
	} else {
		if config.enableH2 {
			return errors.New("config conflict: try enable HTTP2 and disable HTTPS")
		}
		if config.httpRedirect301 {
			return errors.New("config conflict: try enable redirecting HTTPS requests to HTTP ones and disable HTTPS")
		}
	}

	err := httpRequest(cli, "PUT", urlPath, params, &struct {
		HTTPS interface{} `json:"https"`
	}{
		HTTPS: requestObj,
	}, nil)
	if err != nil {
		return err
	}

	return nil
}

// SetHTTPSConfig - enable or disable HTTPS.
// For details, please refer https://cloud.baidu.com/doc/CDN-ABROAD/s/ckb0fx9ea
//
// PARAMS:
//   - cli: the client agent which can perform sending request
//   - domain: the specified domain
//   - httpsConfig: HTTPS configurations.
//
// RETURNS:
//   - error: nil if success otherwise the specific error
func SetHTTPSConfig(cli bce.Client, domain string, httpsConfig *HTTPSConfig) error {
	if httpsConfig == nil {
		return errors.New("HTTPS config is empty")
	}

	var options []HTTPSConfigOption
	if httpsConfig.CertId != "" {
		options = append(options, HTTPSConfigCertID(httpsConfig.CertId))
	}
	if httpsConfig.HttpRedirect {
		options = append(options, HTTPSConfigRedirectWith301())
	}
	if httpsConfig.Http2Enabled {
		options = append(options, HTTPSConfigEnableH2())
	}
	return SetHTTPSConfigWithOptions(cli, domain, httpsConfig.Enabled, options...)
}

// SetOriginProtocol - set originProtocol.
//
// PARAMS:
//   - cli: the client agent which can perform sending request
//   - domain: the specified domain
//   - originProtocol: http or https.
//
// RETURNS:
//   - error: nil if success otherwise the specific error
func SetOriginProtocol(cli bce.Client, domain, originProtocol string) error {
	if originProtocol != "http" && originProtocol != "https" {
		return errors.New("originProtocol must be http or https")
	}

	urlPath := fmt.Sprintf("/v2/abroad/domain/%s/config", domain)
	params := map[string]string{
		"originProtocol": "",
	}

	err := httpRequest(cli, "PUT", urlPath, params, map[string]string{
		"originProtocol": originProtocol,
	}, nil)
	if err != nil {
		return err
	}

	return nil
}

// SetTags - bind ABROAD-CDN domain with the specified tags.
//
// PARAMS:
//   - cli: the client agent can execute sending request
//   - domain: the specified domain
//   - tags: identifying CDN domain as something
//
// RETURNS:
//   - error: nil if success otherwise the specific error
func SetTags(cli bce.Client, domain string, tags []model.TagModel) error {
	urlPath := fmt.Sprintf("/v2/abroad/domain/%s/config", domain)
	params := map[string]string{
		"tags": "",
	}
	err := httpRequest(cli, "PUT", urlPath, params, &struct {
		Tags []model.TagModel `json:"tags"`
	}{
		Tags: tags,
	}, nil)
	if err != nil {
		return err
	}

	return nil
}

// GetTags - get tags the ABROAD-CDN domain bind with.
//
// PARAMS:
//   - cli: the client agent which can perform sending request
//   - domain: the specified domain
//
// RETURNS:
//   - []Tag: tags the ABROAD-CDN domain bind with
//   - error: nil if success otherwise the specific error
func GetTags(cli bce.Client, domain string) ([]model.TagModel, error) {
	urlPath := fmt.Sprintf("/v2/abroad/domain/%s/config", domain)
	params := map[string]string{
		"tags": "",
	}

	respObj := struct {
		Tags []model.TagModel `json:"tags"`
	}{}

	err := httpRequest(cli, "GET", urlPath, params, nil, &respObj)

	if err != nil {
		return nil, err
	}

	return respObj.Tags, nil
}
