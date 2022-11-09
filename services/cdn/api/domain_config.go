package api

import (
	"encoding/json"
	"errors"
	"fmt"
	"strings"

	"github.com/baidubce/bce-sdk-go/bce"
)

// DomainConfig defined a struct for a specified domain's configuration
type DomainConfig struct {
	Domain         string         `json:"domain"`
	Cname          string         `json:"cname"`
	Status         string         `json:"status"`
	CreateTime     string         `json:"createTime"`
	LastModifyTime string         `json:"lastModifyTime"`
	IsBan          string         `json:"isBan"`
	Origin         []OriginPeer   `json:"origin"`
	DefaultHost    string         `json:"defaultHost,omitempty"`
	CacheTTL       []CacheTTL     `json:"cacheTTL"`
	LimitRate      int            `json:"limitRate"`
	RequestAuth    *RequestAuth   `json:"requestAuth,omitempty"`
	Https          *HTTPSConfig   `json:"https,omitempty"`
	FollowProtocol bool           `json:"followProtocol"`
	SeoSwitch      *SeoSwitch     `json:"seoSwitch"`
	Form           string         `json:"form"`
	RangeSwitch    string         `json:"rangeSwitch"`
	OfflineMode    bool           `json:"offlineMode"`
	ClientIp       *ClientIp      `json:"clientIp"`
	OCSP           bool           `json:"ocsp"`
	HttpHeader     []HttpHeader   `json:"httpHeader"`
	MediaDragConf  *MediaDragConf `json:"mediaDragConf"`
	FileTrim       bool           `json:"fileTrim"`
	QUIC           bool           `json:"quic"`
	RefererACL     *RefererACL    `json:"refererACL"`
	IpACL          *IpACL         `json:"ipACL"`
	UaAcl          *UaACL         `json:"uaAcl"`
	AccessLimit    *AccessLimit   `json:"accessLimit"`
	TrafficLimit   *TrafficLimit  `json:"trafficLimit"`
	ErrorPage      []ErrorPage    `json:"errorPage"`
	CacheShare     *CacheShared   `json:"cacheShare"`
}

// CacheTTL defined a struct for cached rules setting
type CacheTTL struct {
	Type   string `json:"type"`
	Value  string `json:"value"`
	Weight int    `json:"weight,omitempty"`
	TTL    int    `json:"ttl"`
}

// CacheShared defined a struct for sharing cache with the other domain
type CacheShared struct {
	Enabled    bool   `json:"enabled"`
	SharedWith string `json:"domain,omitempty"`
}

// RequestAuth defined a struct for the authorization setting
type RequestAuth struct {
	Type            string   `json:"type"`
	Key1            string   `json:"key1"`
	Key2            string   `json:"key2,omitempty"`
	Timeout         int      `json:"timeout,omitempty"`
	WhiteList       []string `json:"whiteList,omitempty"`
	SignArg         string   `json:"signArg,omitempty"`
	TimeArg         string   `json:"timeArg,omitempty"`
	TimestampMetric int      `json:"timestampMetric,omitempty"`
}

// HTTPSConfig defined a struct for configuration about HTTPS
type HTTPSConfig struct {
	Enabled           bool     `json:"enabled"`
	CertId            string   `json:"certId,omitempty"`
	HttpRedirect      bool     `json:"httpRedirect"`
	HttpRedirectCode  int      `json:"httpRedirectCode,omitempty"`
	HttpsRedirect     bool     `json:"httpsRedirect"`
	HttpsRedirectCode int      `json:"httpsRedirectCode"`
	Http2Enabled      bool     `json:"http2Enabled"`
	SslVersion        string   `json:"sslVersion,omitempty"`
	VerifyClient      bool     `json:"verifyClient"`
	SslProtocols      []string `json:"sslProtocols,omitempty"`

	// Deprecated: You can no longer use this field,
	// The better choice is use SetOriginProtocol/GetOriginProtocol.
	HttpOrigin bool `json:"-"`
}

// SeoSwitch defined a struct for SEO setting
type SeoSwitch struct {
	DirectlyOrigin string `json:"diretlyOrigin"`
	PushRecord     string `json:"pushRecord"`
}

// TrafficLimit defined a struct for configure the traffic limitation
type TrafficLimit struct {
	Enabled          bool   `json:"enable"`
	LimitRate        int    `json:"limitRate,omitempty"`
	LimitStartHour   int    `json:"limitStartHour"`
	LimitEndHour     int    `json:"limitEndHour"`
	LimitRateAfter   int    `json:"limitRateAfter,omitempty"`
	TrafficLimitArg  string `json:"trafficLimitArg,omitempty"`
	TrafficLimitUnit string `json:"trafficLimitUnit,omitempty"`
}

// HttpHeader defined a struct for a operation about HTTP header
type HttpHeader struct {
	Type     string `json:"type"`
	Header   string `json:"header"`
	Value    string `json:"value"`
	Action   string `json:"action,omitempty"`
	Describe string `json:"describe,omitempty"`
}

// RefererACL defined a struct for referer ACL setting
type RefererACL struct {
	BlackList  []string `json:"blackList"`
	WhiteList  []string `json:"whiteList"`
	AllowEmpty bool     `json:"allowEmpty"`
}

// IpACL defined a struct for black IP and white IP
type IpACL struct {
	BlackList []string `json:"blackList"`
	WhiteList []string `json:"whiteList"`
}

// UaACL defined a struct for black UA and white UA
type UaACL struct {
	BlackList []string `json:"blackList"`
	WhiteList []string `json:"whiteList"`
}

// ErrorPage defined a struct for redirecting to the custom page when error occur
type ErrorPage struct {
	Code         int    `json:"code"`
	RedirectCode int    `json:"redirectCode,omitempty"`
	Url          string `json:"url"`
}

// MediaCfg defined a struct for a media dragging config
type MediaCfg struct {
	FileSuffix   []string `json:"fileSuffix"`
	StartArgName string   `json:"startArgName,omitempty"`
	EndArgName   string   `json:"endArgName,omitempty"`
	DragMode     string   `json:"dragMode"`
}

// MediaDragConf defined a struct for dragging configs about 'Mp4' and 'Flv'
type MediaDragConf struct {
	Mp4 *MediaCfg `json:"mp4,omitempty"`
	Flv *MediaCfg `json:"flv,omitempty"`
}

// ClientIp defined a struct for how to get client IP
type ClientIp struct {
	Enabled bool   `json:"enabled"`
	Name    string `json:"name,omitempty"`
}

// RetryOrigin defined a struct for how to retry origin servers
type RetryOrigin struct {
	Codes []int `json:"codes"`
}

// AccessLimit defined a struct for access restriction in one client
type AccessLimit struct {
	Enabled bool `json:"enabled"`
	Limit   int  `json:"limit,omitempty"`
}

// CacheUrlArgs defined a struct for cache keys
type CacheUrlArgs struct {
	CacheFullUrl bool     `json:"cacheFullUrl"`
	CacheUrlArgs []string `json:"cacheUrlArgs,omitempty"`
}

// CorsCfg defined a struct for cors setting
type CorsCfg struct {
	IsAllow bool
	Origins []string
}

// GetDomainConfig - get the configuration for the specified domain
// For details, please refer https://cloud.baidu.com/doc/CDN/s/2jwvyf39o
//
// PARAMS:
//     - cli: the client agent which can perform sending request
//     - domain: the specified domain
// RETURNS:
//     - *DomainConfig: the configuration about the specified domain
//     - error: nil if success otherwise the specific error
func GetDomainConfig(cli bce.Client, domain string) (*DomainConfig, error) {
	urlPath := fmt.Sprintf("/v2/domain/%s/config", domain)

	config := &DomainConfig{}

	err := httpRequest(cli, "GET", urlPath, nil, nil, config)
	if err != nil {
		return nil, err
	}

	return config, nil
}

// SetDomainOrigin - set the origin setting for the new
// For details, please refer https://cloud.baidu.com/doc/CDN/s/xjxzi7729
//
// PARAMS:
//     - cli: the client agent which can perform sending request
//     - domain: the specified domain
//     - origins: the origin servers
//     - defaultHost: the default host
// RETURNS:
//     - error: nil if success otherwise the specific error
func SetDomainOrigin(cli bce.Client, domain string, origins []OriginPeer, defaultHost string) error {
	urlPath := fmt.Sprintf("/v2/domain/%s/config", domain)
	params := map[string]string{
		"origin": "",
	}

	err := httpRequest(cli, "PUT", urlPath, params, &struct {
		Origin      []OriginPeer `json:"origin"`
		DefaultHost string       `json:"defaultHost"`
	}{
		Origin:      origins,
		DefaultHost: defaultHost,
	}, nil)

	return err
}

// SetOriginProtocol - set the http protocol back to backend server.
// The valid "originProtocol" must be "http", "https" or "*",
// "http" means send the HTTP request to the backend server,
// "https" means send the HTTPS request to the backend server,
// "*" means send the request follow the client's requesting protocol.
// For details, please refer https://cloud.baidu.com/doc/CDN/s/7k9jdhhlm
//
// PARAMS:
//     - cli: the client agent which can perform sending request
//     - domain: the specified domain
//     - originProtocol: the protocol used for back to the backend server
// RETURNS:
//     - error: nil if success otherwise the specific error
func SetOriginProtocol(cli bce.Client, domain string, originProtocol string) error {
	validOriginProtocols := map[string]bool{
		"http":  true,
		"https": false,
		"*":     true,
	}
	if _, ok := validOriginProtocols[originProtocol]; !ok {
		return fmt.Errorf("invalid originProtocol \"%s\", "+
			"valid value must be \"http\", \"https\" or \"*\"", originProtocol)
	}

	urlPath := fmt.Sprintf("/v2/domain/%s/config", domain)
	params := map[string]string{
		"originProtocol": "",
	}

	err := httpRequest(cli, "PUT", urlPath, params, map[string]interface{}{
		"originProtocol": map[string]string{
			"value": originProtocol,
		},
	}, nil)

	return err
}

// GetOriginProtocol - get the protocol used for back to the backend server.
// For details, please refer https://cloud.baidu.com/doc/CDN/s/dk9jdoob4
//
// PARAMS:
//     - cli: the client agent which can perform sending request
//     - domain: the specified domain
// RETURNS:
//     - string: the protocol used for back to the backend server, it's value must be "http", "https" or "*"
//     - error: nil if success otherwise the specific error
func GetOriginProtocol(cli bce.Client, domain string) (string, error) {
	urlPath := fmt.Sprintf("/v2/domain/%s/config", domain)
	params := map[string]string{
		"originProtocol": "",
	}

	respObj := &struct {
		OriginProtocol struct {
			Value string `json:"value"`
		} `json:"originProtocol"`
	}{}
	err := httpRequest(cli, "GET", urlPath, params, nil, respObj)
	if err != nil {
		return "", err
	}
	if respObj.OriginProtocol.Value == "" {
		return "http", nil
	}
	return respObj.OriginProtocol.Value, nil
}

// SetDomainSeo - set SEO setting
// For details, please refer https://cloud.baidu.com/doc/CDN/s/Jjxziuq4y
//
// PARAMS:
//     - cli: the client agent which can perform sending request
//     - domain: the specified domain
//     - seoSwitch: the setting about SEO
// RETURNS:
//     - error: nil if success otherwise the specific error
func SetDomainSeo(cli bce.Client, domain string, seoSwitch *SeoSwitch) error {
	urlPath := fmt.Sprintf("/v2/domain/%s/config", domain)
	params := map[string]string{
		"seoSwitch": "",
	}

	respObj := &struct {
		Status string `json:"status"`
	}{}

	err := httpRequest(cli, "PUT", urlPath, params, &struct {
		SeoSwitch *SeoSwitch `json:"seoSwitch"`
	}{
		SeoSwitch: seoSwitch,
	}, respObj)
	if err != nil {
		return err
	}

	return nil
}

// GetDomainSeo - retrieve the setting about SEO
// There are two types of data that the server responds to
// 1. `{"seoSwitch":[]}` indicates no setting about SEO
// 2. `{"seoSwitch":{"diretlyOrigin":"ON","pushRecord":"OFF"}}` indicates it had normal setting about SEO
// So the code need to handle the complex affairs
//
// For details, please refer https://cloud.baidu.com/doc/CDN/s/Djxzjfz8f
//
// PARAMS:
//     - cli: the client agent which can perform sending request
//     - domain: the specified domain
// RETURNS:
//     - *SeoSwitch: the setting about SEO
//     - error: nil if success otherwise the specific error
func GetDomainSeo(cli bce.Client, domain string) (*SeoSwitch, error) {
	urlPath := fmt.Sprintf("/v2/domain/%s/config", domain)
	params := map[string]string{
		"seoSwitch": "",
	}

	var respMap map[string]interface{}

	err := httpRequest(cli, "GET", urlPath, params, nil, &respMap)
	if err != nil {
		return nil, err
	}

	if _, ok := respMap["seoSwitch"]; ok {
		if _, ok := respMap["seoSwitch"].([]interface{}); !ok {
			respData, _ := json.Marshal(respMap)
			respObj := &struct {
				SeoSwitch *SeoSwitch `json:"seoSwitch"`
			}{}
			err = json.Unmarshal(respData, respObj)
			if err != nil {
				return nil, err
			}
			return respObj.SeoSwitch, nil
		}
	}

	return nil, nil
}

// GetCacheTTL - get the current cached setting
// For details, please refer https://cloud.baidu.com/doc/CDN/s/ljxzhl9bu
//
// PARAMS:
//     - cli: the client agent which can perform sending request
//     - domain: the specified domain
// RETURNS:
//     - []CacheTTL: the cache setting list
//     - error: nil if success otherwise the specific error
func GetCacheTTL(cli bce.Client, domain string) ([]CacheTTL, error) {
	urlPath := fmt.Sprintf("/v2/domain/%s/config", domain)
	params := map[string]string{
		"cacheTTL": "",
	}

	respObj := &struct {
		CacheTTLs []CacheTTL `json:"cacheTTL"`
	}{}

	err := httpRequest(cli, "GET", urlPath, params, nil, respObj)
	if err != nil {
		return nil, err
	}

	return respObj.CacheTTLs, nil
}

// SetCacheTTL - add some rules for cached setting
// For details, please refer https://cloud.baidu.com/doc/CDN/s/wjxzhgxnx
//
// PARAMS:
//     - cli: the client agent which can perform sending request
//     - domain: the specified domain
//     - cacheTTLs: the cache setting list
// RETURNS:
//     - error: nil if success otherwise the specific error
func SetCacheTTL(cli bce.Client, domain string, cacheTTLs []CacheTTL) error {
	urlPath := fmt.Sprintf("/v2/domain/%s/config", domain)
	params := map[string]string{
		"cacheTTL": "",
	}

	err := httpRequest(cli, "PUT", urlPath, params, &struct {
		CacheTTLs []CacheTTL `json:"cacheTTL"`
	}{
		CacheTTLs: cacheTTLs,
	}, nil)
	if err != nil {
		return err
	}

	return nil
}

// SetRefererACL - set a rule for filter some HTTP request, blackList and whiteList only one can be set
// For details, please refer https://cloud.baidu.com/doc/CDN/s/yjxzhvf21
//
// PARAMS:
//     - cli: the client agent which can perform sending request
//     - domain: the specified domain
//     - blackList: the forbidden host
//     - whiteList: the available host
// RETURNS:
//     - error: nil if success otherwise the specific error
func SetRefererACL(cli bce.Client, domain string, blackList []string, whiteList []string, isAllowEmpty bool) error {
	if len(blackList) != 0 && len(whiteList) != 0 {
		return errors.New("blackList and whiteList cannot exist at the same time")
	}

	refererACLObj := &RefererACL{
		AllowEmpty: isAllowEmpty,
	}

	if blackList != nil {
		refererACLObj.BlackList = blackList
	} else if whiteList != nil {
		refererACLObj.WhiteList = whiteList
	} else {
		return errors.New("blackList and whiteList cannot be nil at the same time")
	}

	urlPath := fmt.Sprintf("/v2/domain/%s/config", domain)
	params := map[string]string{
		"refererACL": "",
	}
	err := httpRequest(cli, "PUT", urlPath, params, &struct {
		RefererACL *RefererACL `json:"refererACL"`
	}{
		RefererACL: refererACLObj,
	}, nil)
	if err != nil {
		return err
	}

	return nil
}

// GetRefererACL - get referer ACL setting
// For details, please refer https://cloud.baidu.com/doc/CDN/s/Ujzkotvtb
//
// PARAMS:
//     - cli: the client agent which can perform sending request
//     - domain: the specified domain
// RETURNS:
//     - *api.RefererACL: referer ACL setting
//     - error: nil if success otherwise the specific error
func GetRefererACL(cli bce.Client, domain string) (*RefererACL, error) {
	urlPath := fmt.Sprintf("/v2/domain/%s/config", domain)
	params := map[string]string{
		"refererACL": "",
	}

	refererACLObj := &RefererACL{}

	err := httpRequest(cli, "GET", urlPath, params, nil, &struct {
		RefererACL *RefererACL `json:"refererACL"`
	}{
		RefererACL: refererACLObj,
	})
	if err != nil {
		return nil, err
	}

	return refererACLObj, nil
}

// SetRefererACL - set a rule for filter some HTTP request, blackList and whiteList only one can be set
// For details, please refer https://cloud.baidu.com/doc/CDN/s/8jxzhwc4d
//
// PARAMS:
//     - cli: the client agent which can perform sending request
//     - domain: the specified domain
//     - blackList: the forbidden ip
//     - whiteList: the available ip
// RETURNS:
//     - error: nil if success otherwise the specific error
func SetIpACL(cli bce.Client, domain string, blackList []string, whiteList []string) error {
	if len(blackList) != 0 && len(whiteList) != 0 {
		return errors.New("blackList and whiteList cannot exist at the same time")
	}

	ipACLObj := &IpACL{}

	if blackList != nil {
		ipACLObj.BlackList = blackList
	} else if whiteList != nil {
		ipACLObj.WhiteList = whiteList
	} else {
		return errors.New("blackList and whiteList cannot be nil at the same time")
	}

	urlPath := fmt.Sprintf("/v2/domain/%s/config", domain)
	params := map[string]string{
		"ipACL": "",
	}
	err := httpRequest(cli, "PUT", urlPath, params, &struct {
		IpACL *IpACL `json:"ipACL"`
	}{
		IpACL: ipACLObj,
	}, nil)
	if err != nil {
		return err
	}

	return nil
}

// GetIpACL - get black IP or white IP
// For details, please refer https://cloud.baidu.com/doc/CDN/s/jjzkp5ku7
//
// PARAMS:
//     - cli: the client agent which can perform sending request
//     - domain: the specified domain
// RETURNS:
//     - *api.IpACL: ip setting
//     - error: nil if success otherwise the specific error
func GetIpACL(cli bce.Client, domain string) (*IpACL, error) {
	urlPath := fmt.Sprintf("/v2/domain/%s/config", domain)
	params := map[string]string{
		"ipACL": "",
	}

	ipACLObj := &IpACL{}
	err := httpRequest(cli, "GET", urlPath, params, nil, &struct {
		IpACL *IpACL `json:"ipACL"`
	}{
		IpACL: ipACLObj,
	})

	if err != nil {
		return nil, err
	}

	return ipACLObj, nil
}

// SetUaACL - set a rule for filter the specified HTTP header named "User-Agent"
// For details, please refer https://cloud.baidu.com/doc/CDN/s/uk88i2a86
//
// PARAMS:
//     - cli: the client agent can execute sending request
//     - domain: the specified domain
//     - blackList: the forbidden UA
//     - whiteList: the available UA
// RETURNS:
//     - error: nil if success otherwise the specific error
func SetUaACL(cli bce.Client, domain string, blackList []string, whiteList []string) error {
	if len(blackList) != 0 && len(whiteList) != 0 {
		return errors.New("blackList and whiteList cannot exist at the same time")
	}

	uaAclObj := &UaACL{}

	if blackList != nil {
		uaAclObj.BlackList = blackList
	} else if whiteList != nil {
		uaAclObj.WhiteList = whiteList
	} else {
		return errors.New("blackList and whiteList cannot be nil at the same time")
	}

	urlPath := fmt.Sprintf("/v2/domain/%s/config", domain)
	params := map[string]string{
		"uaAcl": "",
	}
	err := httpRequest(cli, "PUT", urlPath, params, &struct {
		UaACL *UaACL `json:"uaAcl"`
	}{
		UaACL: uaAclObj,
	}, nil)
	if err != nil {
		return err
	}

	return nil
}

// GetUaACL - get black UA or white UA
// For details, please refer https://cloud.baidu.com/doc/CDN/s/ak88ix19h
//
// PARAMS:
//     - cli: the client agent can execute sending request
//     - domain: the specified domain
// RETURNS:
//     - *api.UaACL: filter config for UA
//     - error: nil if success otherwise the specific error
func GetUaACL(cli bce.Client, domain string) (*UaACL, error) {
	urlPath := fmt.Sprintf("/v2/domain/%s/config", domain)
	params := map[string]string{
		"uaAcl": "",
	}

	uaACLObj := &UaACL{}
	err := httpRequest(cli, "GET", urlPath, params, nil, &struct {
		UaACL *UaACL `json:"uaAcl"`
	}{
		UaACL: uaACLObj,
	})

	if err != nil {
		return nil, err
	}

	return uaACLObj, nil
}

// Deprecated
// SetLimitRate - set limited speed
//
// PARAMS:
//     - cli: the client agent which can perform sending request
//     - domain: the specified domain
//     - limitRate: the limited rate, "1024" means the transmittal speed is less than 1024 Byte/s
// RETURNS:
//     - error: nil if success otherwise the specific error
func SetLimitRate(cli bce.Client, domain string, limitRate int) error {
	urlPath := fmt.Sprintf("/v2/domain/%s/config", domain)
	params := map[string]string{
		"limitRate": "",
	}

	err := httpRequest(cli, "PUT", urlPath, params, &struct {
		LimitRate int `json:"limitRate"`
	}{
		LimitRate: limitRate,
	}, nil)
	if err != nil {
		return err
	}

	return nil
}

// SetTrafficLimit - set the traffic limitation for the specified domain
// For details, please refer https://cloud.baidu.com/doc/CDN/s/ujxzi418e
//
// PARAMS:
//     - cli: the client agent which can perform sending request
//     - domain: the specified domain
//     - trafficLimit: config of traffic limitation
// RETURNS:
//     - error: nil if success otherwise the specific error
func SetTrafficLimit(cli bce.Client, domain string, trafficLimit *TrafficLimit) error {
	urlPath := fmt.Sprintf("/v2/domain/%s/config", domain)
	params := map[string]string{
		"trafficLimit": "",
	}

	err := httpRequest(cli, "PUT", urlPath, params, &struct {
		TrafficLimit *TrafficLimit `json:"trafficLimit"`
	}{
		TrafficLimit: trafficLimit,
	}, nil)
	if err != nil {
		return err
	}

	return nil
}

// GetTrafficLimit - get the traffic limitation of the specified domain
// For details, please refer https://cloud.baidu.com/doc/CDN/s/7k4npdru0
//
// PARAMS:
//     - cli: the client agent which can perform sending request
//     - domain: the specified domain
// RETURNS:
//     - *TrafficLimit: config of traffic limitation
//     - error: nil if success otherwise the specific error
func GetTrafficLimit(cli bce.Client, domain string) (*TrafficLimit, error) {
	urlPath := fmt.Sprintf("/v2/domain/%s/config", domain)
	params := map[string]string{
		"trafficLimit": "",
	}

	respObj := &struct {
		TrafficLimit TrafficLimit `json:"trafficLimit"`
	}{}
	err := httpRequest(cli, "GET", urlPath, params, nil, respObj)
	if err != nil {
		return nil, err
	}

	return &respObj.TrafficLimit, nil
}

// SetDomainHttps - set a rule for speed HTTPS' request
// For details, please refer https://cloud.baidu.com/doc/CDN/s/rjy6v3tnr
//
// PARAMS:
//     - cli: the client agent which can perform sending request
//     - domain: the specified domain
//     - httpsConfig: the rules about the HTTP configure
// RETURNS:
//     - error: nil if success otherwise the specific error
func SetDomainHttps(cli bce.Client, domain string, httpsConfig *HTTPSConfig) error {
	if httpsConfig.HttpRedirect && httpsConfig.HttpsRedirect {
		return errors.New("httpRedirect and httpsRedirect can not be true at the same time")
	}
	if httpsConfig.HttpRedirect {
		if httpsConfig.HttpRedirectCode != 0 && httpsConfig.HttpRedirectCode != 301 && httpsConfig.HttpRedirectCode != 302 {
			return errors.New("invalid httpRedirectCode")
		}
	}

	if httpsConfig.HttpsRedirect {
		if httpsConfig.HttpsRedirectCode != 0 && httpsConfig.HttpsRedirectCode != 301 && httpsConfig.HttpsRedirectCode != 302 {
			return errors.New("invalid httpsRedirectCode")
		}
	}

	urlPath := fmt.Sprintf("/v2/domain/%s/config", domain)
	params := map[string]string{
		"https": "",
	}

	err := httpRequest(cli, "PUT", urlPath, params, &struct {
		HttpsConfig *HTTPSConfig `json:"https"`
	}{
		HttpsConfig: httpsConfig,
	}, nil)
	if err != nil {
		return err
	}

	return nil
}

// GetDomainHttps - get the setting about HTTPS
//
// PARAMS:
//     - cli: the client agent which can perform sending request
//     - domain: the specified domain
// RETURNS:
//     - *HTTPSConfig: the rules about the HTTP configure
//     - error: nil if success otherwise the specific error
func GetDomainHttps(cli bce.Client, domain string) (*HTTPSConfig, error) {
	urlPath := fmt.Sprintf("/v2/domain/%s/config", domain)
	params := map[string]string{
		"https": "",
	}

	respObj := &struct {
		HttpsConfig *HTTPSConfig `json:"https"`
	}{}

	err := httpRequest(cli, "GET", urlPath, params, nil, respObj)
	if err != nil {
		return nil, err
	}

	return respObj.HttpsConfig, nil
}

// SetOCSP - set "OCSP" for the specified domain,
// For details, please refer https://cloud.baidu.com/doc/CDN/s/Pkf2c0ugn
//
// PARAMS:
//     - cli: the client agent can execute sending request
//     - domain: the specified domain
//     - enabled: true for "OCSP" opening otherwise closed
// RETURNS:
//     - error: nil if success otherwise the specific error
func SetOCSP(cli bce.Client, domain string, enabled bool) error {
	urlPath := fmt.Sprintf("/v2/domain/%s/config", domain)
	params := map[string]string{
		"ocsp": "",
	}

	reqObj := map[string]interface{}{
		"ocsp": enabled,
	}
	err := httpRequest(cli, "PUT", urlPath, params, reqObj, nil)
	if err != nil {
		return err
	}

	return nil
}

// GetOCSP - get "OCSP" switch details for the specified domain
// For details, please refer https://cloud.baidu.com/doc/CDN/s/Xkhyjzcvd
//
// PARAMS:
//     - cli: the client agent can execute sending request
//     - domain: the specified domain
// RETURNS:
//     - bool: true for "OCSP" opening otherwise closed
//     - error: nil if success otherwise the specific error
func GetOCSP(cli bce.Client, domain string) (bool, error) {
	urlPath := fmt.Sprintf("/v2/domain/%s/config", domain)
	params := map[string]string{
		"ocsp": "",
	}

	respObj := &struct {
		OfflineMode bool `json:"ocsp"`
	}{}

	err := httpRequest(cli, "GET", urlPath, params, nil, respObj)
	if err != nil {
		return false, err
	}

	return respObj.OfflineMode, nil
}

// SetDomainRequestAuth - set the authorized rules for requesting
// For details, please refer https://cloud.baidu.com/doc/CDN/s/njxzi59g9
//
// PARAMS:
//     - cli: the client agent which can perform sending request
//     - domain: the specified domain
//     - requestAuth: the rules about the auth
// RETURNS:
//     - error: nil if success otherwise the specific error
func SetDomainRequestAuth(cli bce.Client, domain string, requestAuth *RequestAuth) error {
	urlPath := fmt.Sprintf("/v2/domain/%s/config", domain)
	params := map[string]string{
		"requestAuth": "",
	}

	var body interface{}
	if requestAuth == nil {
		body = &struct {
			RequestAuth []string `json:"requestAuth"`
		}{
			RequestAuth: []string{},
		}
	} else {
		body = &struct {
			RequestAuth *RequestAuth `json:"requestAuth"`
		}{
			RequestAuth: requestAuth,
		}
	}

	err := httpRequest(cli, "PUT", urlPath, params, body, nil)
	if err != nil {
		return err
	}

	return nil
}

// SetFollowProtocol - set whether using the same protocol or not when back to the sourced server
// For details, please refer https://cloud.baidu.com/doc/CDN/s/9jxzi89k2
//
// PARAMS:
//     - cli: the client agent which can perform sending request
//     - domain: the specified domain
//     - isFollowProtocol: true in using the same protocol or not when back to the sourced server, false for other
// RETURNS:
//     - error: nil if success otherwise the specific error
func SetFollowProtocol(cli bce.Client, domain string, isFollowProtocol bool) error {
	urlPath := fmt.Sprintf("/v2/domain/%s/config", domain)
	params := map[string]string{
		"followProtocol": "",
	}

	err := httpRequest(cli, "PUT", urlPath, params, &struct {
		FollowProtocol bool `json:"followProtocol"`
	}{
		FollowProtocol: isFollowProtocol,
	}, nil)
	if err != nil {
		return err
	}

	return nil
}

// SetHttpHeader -set some HTTP headers which can be added or deleted when response form CDN edge node
// For details, please refer https://cloud.baidu.com/doc/CDN/s/Jjxzil1sd
//
// PARAMS:
//     - cli: the client agent which can perform sending request
//     - domain: the specified domain
//     - httpHeaders: the HTTP headers' setting
// RETURNS:
//     - error: nil if success otherwise the specific error
func SetHttpHeader(cli bce.Client, domain string, httpHeaders []HttpHeader) error {
	urlPath := fmt.Sprintf("/v2/domain/%s/config", domain)
	params := map[string]string{
		"httpHeader": "",
	}

	err := httpRequest(cli, "PUT", urlPath, params, &struct {
		HttpHeaders []HttpHeader `json:"httpHeader"`
	}{
		HttpHeaders: httpHeaders,
	}, nil)
	if err != nil {
		return err
	}

	return nil
}

// GetHttpHeader - get the HTTP headers' setting
// For details, please refer https://cloud.baidu.com/doc/CDN/s/6jxzip3wn
//
// PARAMS:
//     - cli: the client agent which can perform sending request
//     - domain: the specified domain
// RETURNS:
//     - error: nil if success otherwise the specific error
//     - []HttpHeader: the HTTP headers in setting
func GetHttpHeader(cli bce.Client, domain string) ([]HttpHeader, error) {
	urlPath := fmt.Sprintf("/v2/domain/%s/config", domain)
	params := map[string]string{
		"httpHeader": "",
	}

	respObj := &struct {
		HttpHeaders []HttpHeader `json:"httpHeader"`
	}{}

	err := httpRequest(cli, "GET", urlPath, params, nil, respObj)
	if err != nil {
		return nil, err
	}

	return respObj.HttpHeaders, nil
}

// SetErrorPage - set the page that redirected to when error occurred
// For details, please refer https://cloud.baidu.com/doc/CDN/s/Ejy6vc4yb
//
// PARAMS:
//     - cli: the client agent which can perform sending request
//     - domain: the specified domain
//     - errorPages: the custom pages' setting
// RETURNS:
//     - error: nil if success otherwise the specific error
func SetErrorPage(cli bce.Client, domain string, errorPages []ErrorPage) error {
	urlPath := fmt.Sprintf("/v2/domain/%s/config", domain)
	params := map[string]string{
		"errorPage": "",
	}

	err := httpRequest(cli, "PUT", urlPath, params, &struct {
		ErrorPage []ErrorPage `json:"errorPage"`
	}{
		ErrorPage: errorPages,
	}, nil)
	if err != nil {
		return err
	}

	return nil
}

// GetErrorPage - get the custom pages' setting
// For details, please refer https://cloud.baidu.com/doc/CDN/s/qjy6vfk2u
//
// PARAMS:
//     - cli: the client agent which can perform sending request
//     - domain: the specified domain
// RETURNS:
//     - []ErrorPage: the pages' setting
//     - error: nil if success otherwise the specific error
func GetErrorPage(cli bce.Client, domain string) ([]ErrorPage, error) {
	urlPath := fmt.Sprintf("/v2/domain/%s/config", domain)
	params := map[string]string{
		"errorPage": "",
	}

	respObj := &struct {
		ErrorPage []ErrorPage `json:"errorPage"`
	}{}

	err := httpRequest(cli, "GET", urlPath, params, nil, respObj)
	if err != nil {
		return nil, err
	}

	return respObj.ErrorPage, nil
}

// SetCacheShared - set sharing cache with the other domain.
// For example, 1.test.com shared cache with 2.test.com.
// First, we query http://2.test.com/index.html and got missed.
// Secondly, we query http://1.test.com/index.html and got hit
// because of the CacheShared setting before.
// For details, please refer https://cloud.baidu.com/doc/CDN/s/0kf272ds7
//
// PARAMS:
//     - cli: the client agent can execute sending request
//     - domain: the specified domain
//     - cacheSharedConfig: enabled sets true for shared with the specified domain, otherwise no shared.
// RETURNS:
//     - error: nil if success otherwise the specific error
func SetCacheShared(cli bce.Client, domain string, cacheSharedConfig *CacheShared) error {
	urlPath := fmt.Sprintf("/v2/domain/%s/config", domain)
	params := map[string]string{
		"cacheShare": "",
	}

	err := httpRequest(cli, "PUT", urlPath, params, map[string]interface{}{
		"cacheShare": cacheSharedConfig,
	}, nil)
	if err != nil {
		return err
	}

	return nil
}

// GetCacheShared - get shared cache setting
// For details, please refer https://cloud.baidu.com/doc/CDN/s/Mjy6vo9z2
//
// PARAMS:
//     - cli: the client agent which can perform sending request
//     - domain: the specified domain
// RETURNS:
//     - *CacheShared: shared cache setting
//     - error: nil if success otherwise the specific error
func GetCacheShared(cli bce.Client, domain string) (*CacheShared, error) {
	urlPath := fmt.Sprintf("/v2/domain/%s/config", domain)
	params := map[string]string{
		"cacheShare": "",
	}

	respObj := &struct {
		CacheShared CacheShared `json:"cacheShare"`
	}{}
	err := httpRequest(cli, "GET", urlPath, params, nil, respObj)
	if err != nil {
		return nil, err
	}

	return &respObj.CacheShared, nil
}

// SetMediaDrag - set the media setting about mp4 and flv
// For details, please refer https://cloud.baidu.com/doc/CDN/s/4jy6v6xk3
//
// PARAMS:
//     - cli: the client agent which can perform sending request
//     - domain: the specified domain
//     - mediaDragConf: media setting about mp4 and flv
// RETURNS:
//     - error: nil if success otherwise the specific error
func SetMediaDrag(cli bce.Client, domain string, mediaDragConf *MediaDragConf) error {
	urlPath := fmt.Sprintf("/v2/domain/%s/config", domain)
	params := map[string]string{
		"mediaDrag": "",
	}

	err := httpRequest(cli, "PUT", urlPath, params, &struct {
		MediaDragConf *MediaDragConf `json:"mediaDragConf"`
	}{
		MediaDragConf: mediaDragConf,
	}, nil)
	if err != nil {
		return err
	}

	return nil
}

// GetMediaDrag - get the media setting
// For details, please refer https://cloud.baidu.com/doc/CDN/s/Ojy6v9q8f
//
// PARAMS:
//     - cli: the client agent which can perform sending request
//     - domain: the specified domain
// RETURNS:
//     - *MediaDragConf: the media setting about mp4 and flv
//     - error: nil if success otherwise the specific error
func GetMediaDrag(cli bce.Client, domain string) (*MediaDragConf, error) {
	urlPath := fmt.Sprintf("/v2/domain/%s/config", domain)
	params := map[string]string{
		"mediaDrag": "",
	}

	respObj := &struct {
		MediaDragConf *MediaDragConf `json:"mediaDragConf"`
	}{}

	err := httpRequest(cli, "GET", urlPath, params, nil, respObj)
	if err != nil {
		return nil, err
	}

	return respObj.MediaDragConf, nil
}

// SetFileTrim - trim the text file or not
// For details, please refer https://cloud.baidu.com/doc/CDN/s/Xjy6vimct
//
// PARAMS:
//     - cli: the client agent which can perform sending request
//     - domain: the specified domain
//     - fileTrim: true means trimming the text file, false means do nothing
// RETURNS:
//     - error: nil if success otherwise the specific error
func SetFileTrim(cli bce.Client, domain string, fileTrim bool) error {
	urlPath := fmt.Sprintf("/v2/domain/%s/config", domain)
	params := map[string]string{
		"fileTrim": "",
	}

	err := httpRequest(cli, "PUT", urlPath, params, &struct {
		FileTrim bool `json:"fileTrim"`
	}{
		FileTrim: fileTrim,
	}, nil)
	if err != nil {
		return err
	}

	return nil
}

// GetFileTrim - get the trim setting about text file
// For details, please refer https://cloud.baidu.com/doc/CDN/s/Ujy6vjxnl
//
// PARAMS:
//     - cli: the client agent which can perform sending request
//     - domain: the specified domain
// RETURNS:
//     - bool: true means trimming the text file, false means do nothing
//     - error: nil if success otherwise the specific error
func GetFileTrim(cli bce.Client, domain string) (bool, error) {
	urlPath := fmt.Sprintf("/v2/domain/%s/config", domain)
	params := map[string]string{
		"fileTrim": "",
	}

	respObj := &struct {
		FileTrim bool `json:"fileTrim"`
	}{}

	err := httpRequest(cli, "GET", urlPath, params, nil, respObj)
	if err != nil {
		return false, err
	}

	return respObj.FileTrim, nil
}

// SetIPv6 - open/close IPv6
// For details, please refer https://cloud.baidu.com/doc/CDN/s/qkggncsxp
//
// PARAMS:
//     - cli: the client agent can execute sending request
//     - domain: the specified domain
//     - enabled: true for setting IPv6 switch on otherwise closed
// RETURNS:
//     - error: nil if success otherwise the specific error
func SetIPv6(cli bce.Client, domain string, enabled bool) error {
	urlPath := fmt.Sprintf("/v2/domain/%s/config", domain)
	params := map[string]string{
		"ipv6Dispatch": "",
	}

	reqObj := map[string]interface{}{
		"ipv6Dispatch": map[string]interface{}{
			"enable": enabled,
		},
	}
	err := httpRequest(cli, "PUT", urlPath, params, reqObj, nil)
	if err != nil {
		return err
	}

	return nil
}

// GetIPv6 - get IPv6 switch details for the specified domain
// For details, please refer https://cloud.baidu.com/doc/CDN/s/Ykggnobxd
//
// PARAMS:
//     - cli: the client agent can execute sending request
//     - domain: the specified domain
// RETURNS:
//     - bool: true for setting IPv6 switch on otherwise closed
//     - error: nil if success otherwise the specific error
func GetIPv6(cli bce.Client, domain string) (bool, error) {
	urlPath := fmt.Sprintf("/v2/domain/%s/config", domain)
	params := map[string]string{
		"ipv6Dispatch": "",
	}

	respObj := &struct {
		Ipv6Dispatch struct {
			Enabled bool `json:"enable"`
		} `json:"ipv6Dispatch"`
	}{}

	err := httpRequest(cli, "GET", urlPath, params, nil, respObj)
	if err != nil {
		return false, err
	}

	return respObj.Ipv6Dispatch.Enabled, nil
}

// SetQUIC - open or close QUIC. open QUIC require enabled HTTPS first
// For details, please refer https://cloud.baidu.com/doc/CDN/s/Qkggmoz7p
//
// PARAMS:
//     - cli: the client agent can execute sending request
//     - domain: the specified domain
//     - enabled: true for QUIC opening otherwise closed
// RETURNS:
//     - error: nil if success otherwise the specific error
func SetQUIC(cli bce.Client, domain string, enabled bool) error {
	urlPath := fmt.Sprintf("/v2/domain/%s/config", domain)
	params := map[string]string{
		"quic": "",
	}

	reqObj := map[string]interface{}{
		"quic": enabled,
	}
	err := httpRequest(cli, "PUT", urlPath, params, reqObj, nil)
	if err != nil {
		return err
	}

	return nil
}

// GetQUIC - get QUIC switch details for the specified domain
// For details, please refer https://cloud.baidu.com/doc/CDN/s/pkggn6l1f
//
// PARAMS:
//     - cli: the client agent can execute sending request
//     - domain: the specified domain
// RETURNS:
//     - bool: true for QUIC opening otherwise closed
//     - error: nil if success otherwise the specific error
func GetQUIC(cli bce.Client, domain string) (bool, error) {
	urlPath := fmt.Sprintf("/v2/domain/%s/config", domain)
	params := map[string]string{
		"quic": "",
	}

	respObj := &struct {
		QUIC bool `json:"quic"`
	}{}

	err := httpRequest(cli, "GET", urlPath, params, nil, respObj)
	if err != nil {
		return false, err
	}

	return respObj.QUIC, nil
}

// SetOfflineMode - set "offlineMode" for the specified domain,
// setting true means also response old cached object when got origin server error
// instead of response error to client directly.
// For details, please refer https://cloud.baidu.com/doc/CDN/s/xkhopuj48
//
// PARAMS:
//     - cli: the client agent can execute sending request
//     - domain: the specified domain
//     - enabled: true for offlineMode opening otherwise closed
// RETURNS:
//     - error: nil if success otherwise the specific error
func SetOfflineMode(cli bce.Client, domain string, enabled bool) error {
	urlPath := fmt.Sprintf("/v2/domain/%s/config", domain)
	params := map[string]string{
		"offlineMode": "",
	}

	reqObj := map[string]interface{}{
		"offlineMode": enabled,
	}
	err := httpRequest(cli, "PUT", urlPath, params, reqObj, nil)
	if err != nil {
		return err
	}

	return nil
}

// GetOfflineMode - get "offlineMode" switch details for the specified domain
// For details, please refer https://cloud.baidu.com/doc/CDN/s/tkhopvlkj
//
// PARAMS:
//     - cli: the client agent can execute sending request
//     - domain: the specified domain
// RETURNS:
//     - bool: true for offlineMode opening otherwise closed
//     - error: nil if success otherwise the specific error
func GetOfflineMode(cli bce.Client, domain string) (bool, error) {
	urlPath := fmt.Sprintf("/v2/domain/%s/config", domain)
	params := map[string]string{
		"offlineMode": "",
	}

	respObj := &struct {
		OfflineMode bool `json:"offlineMode"`
	}{}

	err := httpRequest(cli, "GET", urlPath, params, nil, respObj)
	if err != nil {
		return false, err
	}

	return respObj.OfflineMode, nil
}

// SetMobileAccess - distinguish the client or not
// For details, please refer https://cloud.baidu.com/doc/CDN/s/Mjy6vmv6g
//
// PARAMS:
//     - cli: the client agent which can perform sending request
//     - domain: the specified domain
//     - distinguishClient: true means distinguishing the client, false means not
// RETURNS:
//     - error: nil if success otherwise the specific error
func SetMobileAccess(cli bce.Client, domain string, distinguishClient bool) error {
	urlPath := fmt.Sprintf("/v2/domain/%s/config", domain)
	params := map[string]string{
		"mobileAccess": "",
	}

	type mobileAccess struct {
		DistinguishClient bool `json:"distinguishClient"`
	}

	err := httpRequest(cli, "PUT", urlPath, params, &struct {
		MobileAccess *mobileAccess `json:"mobileAccess"`
	}{
		MobileAccess: &mobileAccess{
			DistinguishClient: distinguishClient,
		},
	}, nil)

	return err
}

// GetMobileAccess - get the setting about distinguishing the client or not
// For details, please refer https://cloud.baidu.com/doc/CDN/s/Mjy6vo9z2
//
// PARAMS:
//     - cli: the client agent which can perform sending request
//     - domain: the specified domain
// RETURNS:
//     - bool: true means distinguishing the client, false means not
//     - error: nil if success otherwise the specific error
func GetMobileAccess(cli bce.Client, domain string) (bool, error) {
	urlPath := fmt.Sprintf("/v2/domain/%s/config", domain)
	params := map[string]string{
		"mobileAccess": "",
	}

	respObj := &struct {
		MobileAccess *struct {
			DistinguishClient bool `json:"distinguishClient"`
		} `json:"mobileAccess"`
	}{}

	err := httpRequest(cli, "GET", urlPath, params, nil, respObj)
	if err != nil {
		return false, err
	}

	return respObj.MobileAccess.DistinguishClient, nil
}

// SetClientIp - set the specified HTTP header for the origin server
// For details, please refer https://cloud.baidu.com/doc/CDN/s/Kjy6umyrm
//
// PARAMS:
//     - cli: the client agent which can perform sending request
//     - domain: the specified domain
//     - clientIp: header setting
// RETURNS:
//     - error: nil if success otherwise the specific error
func SetClientIp(cli bce.Client, domain string, clientIp *ClientIp) error {
	urlPath := fmt.Sprintf("/v2/domain/%s/config", domain)
	params := map[string]string{
		"clientIp": "",
	}

	err := httpRequest(cli, "PUT", urlPath, params, &struct {
		ClientIp *ClientIp `json:"clientIp"`
	}{
		ClientIp: clientIp,
	}, nil)

	return err
}

// GetClientIp - get the setting about getting client IP
// For details, please refer https://cloud.baidu.com/doc/CDN/s/8jy6urcq5
//
// PARAMS:
//     - cli: the client agent which can perform sending request
//     - domain: the specified domain
// RETURNS:
//     - *ClientIp: the HTTP header setting for origin server to get client IP
//     - error: nil if success otherwise the specific error
func GetClientIp(cli bce.Client, domain string) (*ClientIp, error) {
	urlPath := fmt.Sprintf("/v2/domain/%s/config", domain)
	params := map[string]string{
		"clientIp": "",
	}

	respObj := &struct {
		ClientIp *ClientIp `json:"clientIp"`
	}{}

	err := httpRequest(cli, "GET", urlPath, params, nil, respObj)
	if err != nil {
		return nil, err
	}

	return respObj.ClientIp, nil
}

// SetRetryOrigin - set the policy for retry origin servers if got failed
// For details, please refer https://cloud.baidu.com/doc/CDN/s/ukhopl3bq
//
// PARAMS:
//     - cli: the client agent can execute sending request
//     - domain: the specified domain
//     - retryOrigin: retry policy
// RETURNS:
//     - error: nil if success otherwise the specific error
func SetRetryOrigin(cli bce.Client, domain string, retryOrigin *RetryOrigin) error {
	urlPath := fmt.Sprintf("/v2/domain/%s/config", domain)
	params := map[string]string{
		"retryOrigin": "",
	}

	err := httpRequest(cli, "PUT", urlPath, params, &struct {
		RetryOrigin *RetryOrigin `json:"retryOrigin"`
	}{
		RetryOrigin: retryOrigin,
	}, nil)

	return err
}

// GetRetryOrigin - get the policy for retry origin servers
// For details, please refer https://cloud.baidu.com/doc/CDN/s/bkhoppbhd
//
// PARAMS:
//     - cli: the client agent which can perform sending request
//     - domain: the specified domain
// RETURNS:
//     - *RetryOrigin: policy of retry origin servers
//     - error: nil if success otherwise the specific error
func GetRetryOrigin(cli bce.Client, domain string) (*RetryOrigin, error) {
	urlPath := fmt.Sprintf("/v2/domain/%s/config", domain)
	params := map[string]string{
		"retryOrigin": "",
	}

	respObj := &struct {
		RetryOrigin *RetryOrigin `json:"retryOrigin"`
	}{}

	err := httpRequest(cli, "GET", urlPath, params, nil, respObj)
	if err != nil {
		return nil, err
	}

	return respObj.RetryOrigin, nil
}

// SetAccessLimit - set the qps for on one client
// For details, please refer https://cloud.baidu.com/doc/CDN/s/Kjy6v02wt
//
// PARAMS:
//     - cli: the client agent which can perform sending request
//     - domain: the specified domain
//     - accessLimit: the access setting
// RETURNS:
//     - error: nil if success otherwise the specific error
func SetAccessLimit(cli bce.Client, domain string, accessLimit *AccessLimit) error {
	urlPath := fmt.Sprintf("/v2/domain/%s/config", domain)
	params := map[string]string{
		"accessLimit": "",
	}

	err := httpRequest(cli, "PUT", urlPath, params, &struct {
		AccessLimit *AccessLimit `json:"accessLimit"`
	}{
		AccessLimit: accessLimit,
	}, nil)

	return err
}

// GetAccessLimit - get the qps setting
// For details, please refer https://cloud.baidu.com/doc/CDN/s/rjy6v3tnr
//
// PARAMS:
//     - cli: the client agent which can perform sending request
//     - domain: the specified domain
// RETURNS:
//     - *AccessLimit: the access setting
//     - error: nil if success otherwise the specific error
func GetAccessLimit(cli bce.Client, domain string) (*AccessLimit, error) {
	urlPath := fmt.Sprintf("/v2/domain/%s/config", domain)
	params := map[string]string{
		"accessLimit": "",
	}

	respObj := &struct {
		AccessLimit *AccessLimit `json:"accessLimit"`
	}{}

	err := httpRequest(cli, "GET", urlPath, params, nil, respObj)
	if err != nil {
		return nil, err
	}

	return respObj.AccessLimit, nil
}

// SetCacheUrlArgs - tell the CDN system cache the url's params or not
// For details, please refer https://cloud.baidu.com/doc/CDN/s/vjxzho0kx
//
// PARAMS:
//     - cli: the client agent which can perform sending request
//     - domain: the specified domain
//     - cacheFullUrl: whether cache the full url or not, full url means include params, also some extra params can be avoided
// RETURNS:
//     - error: nil if success otherwise the specific error
func SetCacheUrlArgs(cli bce.Client, domain string, cacheFullUrl *CacheUrlArgs) error {
	urlPath := fmt.Sprintf("/v2/domain/%s/config", domain)
	params := map[string]string{
		"cacheFullUrl": "",
	}

	err := httpRequest(cli, "PUT", urlPath, params, cacheFullUrl, nil)

	return err
}

// GetCacheUrlArgs - get the cached rules
// For details, please refer https://cloud.baidu.com/doc/CDN/s/sjxzhsb6h
//
// PARAMS:
//     - cli: the client agent which can perform sending request
//     - domain: the specified domain
// RETURNS:
//     - *CacheUrlArgs: the details about cached rules
//     - error: nil if success otherwise the specific error
func GetCacheUrlArgs(cli bce.Client, domain string) (*CacheUrlArgs, error) {
	urlPath := fmt.Sprintf("/v2/domain/%s/config", domain)
	params := map[string]string{
		"cacheFullUrl": "",
	}

	respObj := &CacheUrlArgs{}

	err := httpRequest(cli, "GET", urlPath, params, nil, respObj)
	if err != nil {
		return nil, err
	}

	return respObj, nil
}

// SetTlsVersions - set some TLS versions that you want
//
// PARAMS:
//     - cli: the client agent which can perform sending request
//     - domain: the specified domain
//     - tlsVersions: TLS version settings
// RETURNS:
//     - error: nil if success otherwise the specific error
func SetTlsVersions(cli bce.Client, domain string, tlsVersions []string) error {
	urlPath := fmt.Sprintf("/v2/domain/%s/config", domain)
	params := map[string]string{
		"tlsVersions": "",
	}

	reqData := &struct {
		Data []string `json:"data"`
	}{
		Data: tlsVersions,
	}

	err := httpRequest(cli, "PUT", urlPath, params, &struct {
		TlsVersions interface{} `json:"tlsVersions"`
	}{
		TlsVersions: reqData,
	}, nil)

	return err
}

// SetCors - set about Cross-origin resource sharing
// For details, please refer https://cloud.baidu.com/doc/CDN/s/Rjxzi1cfs
//
// PARAMS:
//     - cli: the client agent which can perform sending request
//     - domain: the specified domain
//     - isAllow: true means allow Cors, false means not allow
//     - originList: the origin setting, it's invalid when isAllow is false
// RETURNS:
//     - error: nil if success otherwise the specific error
func SetCors(cli bce.Client, domain string, isAllow bool, originList []string) error {
	urlPath := fmt.Sprintf("/v2/domain/%s/config", domain)
	params := map[string]string{
		"cors": "",
	}

	allow := "off"
	var origins []string
	origins = []string{}

	if isAllow {
		allow = "on"
		origins = originList
	}

	reqData := &struct {
		Allow      string   `json:"allow"`
		OriginList []string `json:"originList"`
	}{
		Allow:      allow,
		OriginList: origins,
	}

	err := httpRequest(cli, "PUT", urlPath, params, &struct {
		Cors interface{} `json:"cors"`
	}{
		Cors: reqData,
	}, nil)

	return err
}

// GetCors - get the Cors setting
// For details, please refer https://cloud.baidu.com/doc/CDN/s/tjxzi2d7t
//
// PARAMS:
//     - cli: the client agent which can perform sending request
//     - domain: the specified domain
// RETURNS:
//     - *CorsCfg: the Cors setting
//     - error: nil if success otherwise the specific error
func GetCors(cli bce.Client, domain string) (*CorsCfg, error) {
	urlPath := fmt.Sprintf("/v2/domain/%s/config", domain)
	params := map[string]string{
		"cors": "",
	}

	respObj := &struct {
		Allow      string   `json:"allow"`
		OriginList []string `json:"originList,omitempty"`
	}{}

	err := httpRequest(cli, "GET", urlPath, params, nil, &struct {
		Cors interface{} `json:"cors"`
	}{
		Cors: respObj,
	})

	if err != nil {
		return nil, err
	}

	corsCfg := &CorsCfg{
		IsAllow: false,
		Origins: respObj.OriginList,
	}

	if strings.ToUpper(respObj.Allow) == "ON" {
		corsCfg.IsAllow = true
	}

	return corsCfg, nil
}

// SetRangeSwitch - set the range setting
// For details, please refer https://cloud.baidu.com/doc/CDN/s/Fjxziabst
//
// PARAMS:
//     - cli: the client agent which can perform sending request
//     - domain: the specified domain
//     - enabled: true means enable range cached, false means disable range cached
// RETURNS:
//     - error: nil if success otherwise the specific error
func SetRangeSwitch(cli bce.Client, domain string, enabled bool) error {
	urlPath := fmt.Sprintf("/v2/domain/%s/config", domain)
	params := map[string]string{
		"rangeSwitch": "",
	}

	err := httpRequest(cli, "PUT", urlPath, params, &struct {
		RangeSwitch bool `json:"rangeSwitch"`
	}{
		RangeSwitch: enabled,
	}, nil)

	return err
}

// GetRangeSwitch - get the range setting
// For details, please refer https://cloud.baidu.com/doc/CDN/s/5jxzid6o9
//
// PARAMS:
//     - cli: the client agent which can perform sending request
//     - domain: the specified domain
// RETURNS:
//     - bool: true means enable range cached, false means disable range cached
//     - error: nil if success otherwise the specific error
func GetRangeSwitch(cli bce.Client, domain string) (bool, error) {
	urlPath := fmt.Sprintf("/v2/domain/%s/config", domain)
	params := map[string]string{
		"rangeSwitch": "",
	}

	respObj := &struct {
		RangeSwitch string `json:"rangeSwitch"`
	}{}

	err := httpRequest(cli, "GET", urlPath, params, nil, respObj)

	if err != nil {
		return false, err
	}

	enabled := false
	if strings.ToUpper(respObj.RangeSwitch) == "ON" {
		enabled = true
	}

	return enabled, nil
}

// SetContentEncoding - set Content-Encoding
// For details, please refer https://cloud.baidu.com/doc/CDN/s/0jyqyahsb
//
// PARAMS:
//     - cli: the client agent which can perform sending request
//     - domain: the specified domain
//     - enabled: true means using the specified encoding algorithm indicated by "encodingType" in transferring,
//         false means disable encoding
// RETURNS:
//     - error: nil if success otherwise the specific error
func SetContentEncoding(cli bce.Client, domain string, enabled bool, encodingType string) error {
	if enabled && encodingType != "gzip" && encodingType != "br" && encodingType != "all" {
		errMsg := fmt.Sprintf("invalid encoding type \"%s\" for setting Content-Encoding,"+
			" it must in \"gzip\", \"br\" and \"all\"", encodingType)
		return errors.New(errMsg)
	}

	if !enabled {
		encodingType = ""
	}

	urlPath := fmt.Sprintf("/v2/domain/%s/config", domain)
	params := map[string]string{
		"compress": "",
	}

	respObj := &struct {
		Allow bool   `json:"allow"`
		Type  string `json:"type,omitempty"`
	}{
		Allow: enabled,
		Type:  encodingType,
	}

	err := httpRequest(cli, "PUT", urlPath, params, &struct {
		ContentEncoding interface{} `json:"compress"`
	}{
		ContentEncoding: respObj,
	}, nil)

	return err
}

// GetContentEncoding - get the setting about Content-Encoding
// For details, please refer https://cloud.baidu.com/doc/CDN/s/bjyqycw8g
//
// PARAMS:
//     - cli: the client agent which can perform sending request
//     - domain: the specified domain
// RETURNS:
//     - string: the encoding algorithm for transferring, empty means disable encoding in transferring
//     - error: nil if success otherwise the specific error
func GetContentEncoding(cli bce.Client, domain string) (string, error) {
	urlPath := fmt.Sprintf("/v2/domain/%s/config", domain)
	params := map[string]string{
		"compress": "",
	}

	respObj := &struct {
		Allow string `json:"allow"`
		Type  string `json:"type,omitempty"`
	}{}

	err := httpRequest(cli, "GET", urlPath, params, nil, &struct {
		ContentEncoding interface{} `json:"compress"`
	}{
		ContentEncoding: respObj,
	})

	if err != nil {
		return "", err
	}

	contentEncoding := respObj.Type
	if respObj.Allow == "off" {
		contentEncoding = ""
	}

	return contentEncoding, nil
}
