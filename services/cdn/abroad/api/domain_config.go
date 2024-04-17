package api

import (
	"fmt"

	"github.com/baidubce/bce-sdk-go/bce"
)

// DomainConfig defined a struct for total configurations.
type DomainConfig struct {
	Domain       string       `json:"domain"`
	Origin       []OriginPeer `json:"originConfig"`
	CacheTTL     []CacheTTL   `json:"cacheTtl"`
	CacheFullUrl bool         `json:"cacheFullUrl"`
	OriginHost   *string      `json:"originHost"`
	RefererACL   *RefererACL  `json:"refererACL"`
	IpACL        *IpACL       `json:"ipACL"`
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

// GetDomainConfig - get the configuration for the specified domain
// For details, please refer https://cloud.baidu.com/doc/CDN-ABROAD/s/9kbsye6k8
//
// PARAMS:
//     - cli: the client agent which can perform sending request
//     - domain: the specified domain
// RETURNS:
//     - *DomainConfig: the configuration about the specified domain
//     - error: nil if success otherwise the specific error
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
//     - cli: the client agent which can perform sending request
//     - domain: the specified domain
//     - origins: the origin servers
// RETURNS:
//     - error: nil if success otherwise the specific error
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
//     - cli: the client agent which can perform sending request
//     - domain: the specified domain
//     - cacheTTLs: the cache setting list
// RETURNS:
//     - error: nil if success otherwise the specific error
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
//     - cli: the client agent which can perform sending request
//     - domain: the specified domain
//     - cacheFullUrl: the query part in URL being added to the cache key or not
// RETURNS:
//     - error: nil if success otherwise the specific error
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
//     - cli: the client agent which can perform sending request
//     - domain: the specified domain
//     - originHost: specified HOST header for origin server
// RETURNS:
//     - error: nil if success otherwise the specific error
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
//     - cli: the client agent which can perform sending request
//     - domain: the specified domain
//     - refererACL: referer of whitelist or blacklist
// RETURNS:
//     - error: nil if success otherwise the specific error
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
//     - cli: the client agent which can perform sending request
//     - domain: the specified domain
//     - ipACL: IP whitelist or blacklist
// RETURNS:
//     - error: nil if success otherwise the specific error
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
