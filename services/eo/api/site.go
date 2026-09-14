package api

import (
	"bytes"
	"encoding/json"
	"errors"

	"github.com/baidubce/bce-sdk-go/bce"
)

// CacheTtl defined cache rule for a site
type CacheTtl struct {
	Value          string `json:"value"`
	Weight         int    `json:"weight"`
	OverrideOrigin bool   `json:"override_origin"`
	Ttl            int    `json:"ttl"`
	Type           string `json:"type"`
}

// CacheCodeTtl defined status-code cache rule for a site.
type CacheCodeTtl struct {
	Value          string `json:"value"`
	Weight         int    `json:"weight"`
	OverrideOrigin bool   `json:"overrideOrigin"`
	Ttl            int    `json:"ttl"`
	Type           string `json:"type"`
}

// CacheKey defined the query string strategy for cache key.
type CacheKey struct {
	Query       *bool     `json:"query,omitempty"`
	IncludeArgs *[]string `json:"include_args,omitempty"`
	ExcludeArgs *[]string `json:"exclude_args,omitempty"`
	IgnoreCase  *bool     `json:"ignore_case,omitempty"`
}

// HSTS defined the HSTS strategy for a site.
type HSTS struct {
	MaxAge            *int  `json:"maxAge,omitempty"`
	IncludeSubDomains *bool `json:"includeSubDomains,omitempty"`
	Preload           *bool `json:"preload,omitempty"`
}

// HTTP3 defined the HTTP3 strategy for a site.
type HTTP3 struct {
	Enable *bool `json:"enable,omitempty"`
}

// PageRule defined a rule-engine entry of a site.
type PageRule struct {
	Name   string      `json:"name"`
	Status string      `json:"status"`
	Rules  [][]Rule    `json:"rules"`
	Config *RuleConfig `json:"config,omitempty"`
}

// Rule defined a single match condition inside PageRule.Rules.
// Outer slice means OR; inner slice means AND.
// Values accepts []string, a single string (e.g. remoteIsp / regex operator),
// or []RemoteGeo (matchFrom = remoteGeos); it can be omitted for the `exists` operator.
type Rule struct {
	MatchFrom  string      `json:"matchFrom"`
	Operator   string      `json:"operator"`
	MatchKey   string      `json:"matchKey,omitempty"`
	Values     interface{} `json:"values,omitempty"`
	IgnoreCase *bool       `json:"ignoreCase,omitempty"`
}

// RemoteGeo defined a geo entry used by Rule.Values when matchFrom is "remoteGeos".
// Use Pro for mainland China provinces and Cty for other countries.
type RemoteGeo struct {
	Pro string `json:"pro,omitempty"`
	Cty string `json:"cty,omitempty"`
}

// RuleCacheKey defined the query string strategy for cache key in rule scope.
// It supports two extra fields (`headers` / `cookie`) compared to the site-scope CacheKey.
type RuleCacheKey struct {
	Query       *bool     `json:"query,omitempty"`
	IncludeArgs *[]string `json:"include_args,omitempty"`
	ExcludeArgs *[]string `json:"exclude_args,omitempty"`
	IgnoreCase  *bool     `json:"ignore_case,omitempty"`
	Headers     *[]string `json:"headers,omitempty"`
	Cookie      *[]string `json:"cookie,omitempty"`
}

// RefreshRevalidate defined the refresh-revalidate strategy in rule scope.
type RefreshRevalidate struct {
	Enabled *bool `json:"enabled,omitempty"`
}

// WebSocket defined the WebSocket strategy for a site or in rule scope.
type WebSocket struct {
	Enabled *bool `json:"enabled,omitempty"`
	Timeout *int  `json:"timeout,omitempty"`
}

// OriginTimeout defined the origin timeout strategy in rule scope.
type OriginTimeout struct {
	LoadTimeout    *int `json:"loadTimeout,omitempty"`
	ConnectTimeout *int `json:"connectTimeout,omitempty"`
}

// OriginRedirectOptions defined the origin redirect strategy in rule scope.
type OriginRedirectOptions struct {
	EnableRedirectFollow   *string `json:"enableRedirectFollow,omitempty"`
	MaxRedirectFollowCount *int    `json:"maxRedirectFollowCount,omitempty"`
}

// OriginOptions defined origin range/part-size strategy in rule scope.
type OriginOptions struct {
	Range    *string `json:"range,omitempty"`
	PartSize *string `json:"partSize,omitempty"`
}

// HeaderMap defined a string key-value header object.
// The server represents "no header configured" as an empty JSON array `[]` instead of `{}`,
// so both forms are accepted when decoding, and an empty map is encoded back as `[]`.
type HeaderMap map[string]string

func (h HeaderMap) MarshalJSON() ([]byte, error) {
	if len(h) == 0 {
		return []byte("[]"), nil
	}
	return json.Marshal(map[string]string(h))
}

func (h *HeaderMap) UnmarshalJSON(data []byte) error {
	trimmed := bytes.TrimSpace(data)
	if len(trimmed) == 0 || bytes.Equal(trimmed, []byte("null")) || bytes.Equal(trimmed, []byte("[]")) {
		*h = HeaderMap{}
		return nil
	}
	m := map[string]string{}
	if err := json.Unmarshal(trimmed, &m); err != nil {
		return err
	}
	*h = m
	return nil
}

// OriginRequest defined the origin request header strategy for a site.
type OriginRequest struct {
	AddHeaders    *HeaderMap `json:"addHeaders,omitempty"`
	RemoveHeaders *HeaderMap `json:"removeHeaders,omitempty"`
}

// ClientResponse defined the edge node response header strategy for a site.
type ClientResponse struct {
	AddHeaders    *HeaderMap `json:"addHeaders,omitempty"`
	RemoveHeaders *HeaderMap `json:"removeHeaders,omitempty"`
}

// ClientCacheTtl defined the browser cache TTL strategy for a site.
type ClientCacheTtl struct {
	Mode *string `json:"mode,omitempty"`
	Ttl  *int    `json:"ttl,omitempty"`
}

// ThirdBucketAuth defined the object storage authentication of an origin.
// Fields ak/sk/bucket/region/service are not required when authType is "bos".
type ThirdBucketAuth struct {
	AuthType string  `json:"authType"`
	Enabled  *bool   `json:"enabled,omitempty"`
	Ak       *string `json:"ak,omitempty"`
	Sk       *string `json:"sk,omitempty"`
	Bucket   *string `json:"bucket,omitempty"`
	Region   *string `json:"region,omitempty"`
	Service  *string `json:"service,omitempty"`
}

// OriginItem defined a single origin entry in rule scope.
type OriginItem struct {
	Addr             string           `json:"addr"`
	Type             string           `json:"type"`
	UpstreamProtocol string           `json:"upstreamProtocol"`
	HttpPort         *int             `json:"httpPort,omitempty"`
	HttpsPort        *int             `json:"httpsPort,omitempty"`
	Host             *string          `json:"host,omitempty"`
	ThirdBucketAuth  *ThirdBucketAuth `json:"thirdBucketAuth,omitempty"`
}

// OriginPool defined the origin pool strategy in rule scope.
// It is mutually exclusive with RuleConfig.OriginConfig.
type OriginPool struct {
	OriginPoolId     string `json:"originPoolId"`
	UpstreamProtocol string `json:"upstreamProtocol"`
	HttpPort         *int   `json:"httpPort,omitempty"`
	HttpsPort        *int   `json:"httpsPort,omitempty"`
}

// RealIp defined the real-IP strategy for a site or in rule scope.
type RealIp struct {
	Enabled *bool   `json:"enabled,omitempty"`
	Name    *string `json:"name,omitempty"`
}

// ErrorPage defined a single custom error-page entry in rule scope.
type ErrorPage struct {
	Code *int    `json:"code,omitempty"`
	Url  *string `json:"url,omitempty"`
}

// TrafficLimit defined the traffic-limit strategy in rule scope.
type TrafficLimit struct {
	Enable         *bool   `json:"enable,omitempty"`
	LimitRate      *int    `json:"limitRate,omitempty"`
	LimitStartHour *int    `json:"limitStartHour,omitempty"`
	LimitEndHour   *int    `json:"limitEndHour,omitempty"`
	LimitRateUnit  *string `json:"limitRateUnit,omitempty"`
}

// AntiHotLink defined the anti-hot-link strategy in rule scope.
type AntiHotLink struct {
	AntiType        *string `json:"antiType,omitempty"`
	SecretKey       *string `json:"secretKey,omitempty"`
	NewsecretKey    *string `json:"newsecretKey,omitempty"`
	Timeout         *int    `json:"timeout,omitempty"`
	TimestampFormat *string `json:"timestampFormat,omitempty"`
	AuthArg         *string `json:"authArg,omitempty"`
}

// OriginArg defined the origin-args strategy in rule scope.
type OriginArg struct {
	Ignore *bool     `json:"ignore,omitempty"`
	Args   *[]string `json:"args,omitempty"`
}

// UrlRules defined a single URL-rewrite/redirect rule in rule scope.
type UrlRules struct {
	Scheme  *string `json:"scheme,omitempty"`
	Host    *string `json:"host,omitempty"`
	SrcPath *string `json:"srcPath,omitempty"`
	DstPath *string `json:"dstPath,omitempty"`
	Query   *string `json:"query,omitempty"`
	Status  *int    `json:"status,omitempty"`
}

// RuleConfig defined the per-rule configurations that can be applied when a PageRule matches.
type RuleConfig struct {
	CacheTtl              *[]CacheTtl            `json:"cacheTtl,omitempty"`
	CacheKey              *RuleCacheKey          `json:"cacheKey,omitempty"`
	OfflineMode           *string                `json:"offlineMode,omitempty"`
	RefreshRevalidate     *RefreshRevalidate     `json:"refreshRevalidate,omitempty"`
	HttpToHttpsEnabled    *string                `json:"httpToHttpsEnabled,omitempty"`
	HttpToHttpsCode       *string                `json:"httpToHttpsCode,omitempty"`
	Hsts                  *HSTS                  `json:"hsts,omitempty"`
	Isa                   *string                `json:"isa,omitempty"`
	Http2Disable          *string                `json:"http2Disable,omitempty"`
	Http3                 *HTTP3                 `json:"http3,omitempty"`
	WebSocket             *WebSocket             `json:"webSocket,omitempty"`
	Http2Origin           *string                `json:"http2Origin,omitempty"`
	ClientMaxBodySize     *string                `json:"clientMaxBodySize,omitempty"`
	Compress              *string                `json:"compress,omitempty"`
	CompressMethodArray   *[]string              `json:"compressMethodArray,omitempty"`
	OriginTimeout         *OriginTimeout         `json:"originTimeout,omitempty"`
	OriginRedirectOptions *OriginRedirectOptions `json:"originRedirectOptions,omitempty"`
	OriginOptions         *OriginOptions         `json:"originOptions,omitempty"`
	RealIp                *RealIp                `json:"realIp,omitempty"`
	ErrorPage             *[]ErrorPage           `json:"errorPage,omitempty"`
	TrafficLimit          *TrafficLimit          `json:"trafficLimit,omitempty"`
	AntiHotLink           *AntiHotLink           `json:"antiHotLink,omitempty"`
	OriginArg             *OriginArg             `json:"originArg,omitempty"`
	UrlRules              *[]UrlRules            `json:"urlRules,omitempty"`
	SslProtocols          *[]string              `json:"sslProtocols,omitempty"`
	SslMode               **string               `json:"sslMode,omitempty"`
	SslCipherList         **string               `json:"sslCipherList,omitempty"`
	Ocsp                  *string                `json:"ocsp,omitempty"`
	CacheCodeTtl          *[]CacheCodeTtl        `json:"cacheCodeTtl,omitempty"`
	ClientCacheTtl        *ClientCacheTtl        `json:"clientCacheTtl,omitempty"`
	OriginConfig          *[]OriginItem          `json:"originConfig,omitempty"`
	OriginPool            *OriginPool            `json:"originPool,omitempty"`
	OriginRequest         *OriginRequest         `json:"originRequest,omitempty"`
	ClientResponse        *ClientResponse        `json:"clientResponse,omitempty"`
}

// SiteConfig defined a unified container for site configurations.
// All fields use pointer types so that callers can distinguish "not set" (nil, will be omitted)
// from "explicitly set to empty" (non-nil empty value, will be sent as `[]` / `{}`).
// New configuration items should also be added as pointer fields with `omitempty`.
type SiteConfig struct {
	CacheTtl            *[]CacheTtl     `json:"cacheTtl,omitempty"`
	CacheKey            *CacheKey       `json:"cacheKey,omitempty"`
	OfflineMode         *string         `json:"offlineMode,omitempty"`
	HttpToHttpsEnabled  *string         `json:"httpToHttpsEnabled,omitempty"`
	HttpToHttpsCode     *string         `json:"httpToHttpsCode,omitempty"`
	Hsts                *HSTS           `json:"hsts,omitempty"`
	Http2Disable        *string         `json:"http2Disable,omitempty"`
	Http3               *HTTP3          `json:"http3,omitempty"`
	ClientMaxBodySize   *string         `json:"clientMaxBodySize,omitempty"`
	Compress            *string         `json:"compress,omitempty"`
	CompressMethodArray *[]string       `json:"compressMethodArray,omitempty"`
	Isa                 *string         `json:"isa,omitempty"`
	CacheCodeTtl        *[]CacheCodeTtl `json:"cacheCodeTtl,omitempty"`
	GrpcOrigin          *string         `json:"grpcOrigin,omitempty"`
	Http2Origin         *string         `json:"http2Origin,omitempty"`
	PageRules           *[]PageRule     `json:"pageRules,omitempty"`
	OriginRequest       *OriginRequest  `json:"originRequest,omitempty"`
	ClientResponse      *ClientResponse `json:"clientResponse,omitempty"`
	ClientCacheTtl      *ClientCacheTtl `json:"clientCacheTtl,omitempty"`
	WebSocket           *WebSocket      `json:"webSocket,omitempty"`
	SslProtocols        *[]string       `json:"sslProtocols,omitempty"`
	SslMode             **string        `json:"sslMode,omitempty"`
	SslCipherList       **string        `json:"sslCipherList,omitempty"`
	Ocsp                *string         `json:"ocsp,omitempty"`
	RealIp              *RealIp         `json:"realIp,omitempty"`
}

// SiteConfigUpdateResult defined the response of SetSiteConfig
type SiteConfigUpdateResult struct {
	Status string `json:"status"`
}

// SetSiteConfig - set the site-level configurations
// For details, please refer to https://cloud.baidu.com/doc/GEO/s/vmiia4s0j
//
// PARAMS:
//   - cli: the client agent which can perform sending request
//   - site: the site to be configured
//   - config: the configurations to be set;
//
// RETURNS:
//   - *SiteConfigUpdateResult: the update status returned by the server
//   - error: nil if success otherwise the specific error
func SetSiteConfig(cli bce.Client, site string, config *SiteConfig) (*SiteConfigUpdateResult, error) {
	if site == "" {
		return nil, errors.New("site is required")
	}
	if config == nil {
		return nil, errors.New("config is required")
	}

	respObj := &SiteConfigUpdateResult{}
	err := httpRequest(cli, "PUT", "/v2/geo/site/"+site+"/config", nil, config, respObj)
	if err != nil {
		return nil, err
	}

	return respObj, nil
}

// GetSiteConfig - get the site configurations
//
// PARAMS:
//   - cli: the client agent which can perform sending request
//   - site: the site to be queried
//
// RETURNS:
//   - *SiteConfig: the configurations of the site
//   - error: nil if success otherwise the specific error
func GetSiteConfig(cli bce.Client, site string) (*SiteConfig, error) {
	if site == "" {
		return nil, errors.New("site is required")
	}

	respObj := &SiteConfig{}
	err := httpRequest(cli, "GET", "/v2/geo/site/"+site+"/config", nil, nil, respObj)
	if err != nil {
		return nil, err
	}

	return respObj, nil
}
