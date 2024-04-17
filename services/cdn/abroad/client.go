package abroad

import (
	"github.com/baidubce/bce-sdk-go/auth"
	"github.com/baidubce/bce-sdk-go/bce"
	"github.com/baidubce/bce-sdk-go/services/cdn/abroad/api"
)

const (
	DEFAULT_SERVICE_DOMAIN = "cdn.baidubce.com"
)

// Client of CDN service is a kind of BceClient, so derived from BceClient
type Client struct {
	*bce.BceClient
}

// NewClient make the CDN service client with default configuration
// Use `cli.Config.xxx` to access the config or change it to non-default value
func NewClient(ak, sk, endpoint string) (*Client, error) {
	var credentials *auth.BceCredentials
	var err error
	if len(ak) == 0 && len(sk) == 0 { // to support public-read-write request
		credentials, err = nil, nil
	} else {
		credentials, err = auth.NewBceCredentials(ak, sk)
		if err != nil {
			return nil, err
		}
	}
	if len(endpoint) == 0 {
		endpoint = DEFAULT_SERVICE_DOMAIN
	}
	defaultSignOptions := &auth.SignOptions{
		HeadersToSign: auth.DEFAULT_HEADERS_TO_SIGN,
		ExpireSeconds: auth.DEFAULT_EXPIRE_SECONDS}
	defaultConf := &bce.BceClientConfiguration{
		Endpoint:                  endpoint,
		Region:                    bce.DEFAULT_REGION,
		UserAgent:                 bce.DEFAULT_USER_AGENT,
		Credentials:               credentials,
		SignOption:                defaultSignOptions,
		Retry:                     bce.DEFAULT_RETRY_POLICY,
		ConnectionTimeoutInMillis: bce.DEFAULT_CONNECTION_TIMEOUT_IN_MILLIS}
	v1Signer := &auth.BceV1Signer{}

	client := &Client{bce.NewBceClient(defaultConf, v1Signer)}
	return client, nil
}

// SendCustomRequest - send a HTTP request, and response data or error, it use the default times for retrying
//
// PARAMS:
//     - method: the HTTP requested method, e.g. "GET", "POST", "PUT" ...
//     - urlPath: a path component, consisting of a sequence of path segments separated by a slash ( / ).
//     - params: the query params, which will be append to the query path, and separate by "&"
//         e.g. http://www.baidu.com?query_param1=value1&query_param2=value2
//     - reqHeaders: the request http headers
//     - bodyObj: the HTTP requested body content transferred to a goland object
//     - respObj: the HTTP response content transferred to a goland object
// RETURNS:
//     - error: nil if success otherwise the specific error
func (cli *Client) SendCustomRequest(method string, urlPath string, params, reqHeaders map[string]string, bodyObj interface{}, respObj interface{}) error {
	return api.SendCustomRequest(cli, method, urlPath, params, reqHeaders, bodyObj, respObj)
}

// GetDomainConfig - get the configuration for the specified domain.
// For details, please refer https://cloud.baidu.com/doc/CDN-ABROAD/s/9kbsye6k8
//
// PARAMS:
//     - domain: the specified domain
// RETURNS:
//     - *DomainConfig: the configuration about the specified domain
//     - error: nil if success otherwise the specific error
func (cli *Client) GetDomainConfig(domain string) (*api.DomainConfig, error) {
	return api.GetDomainConfig(cli, domain)
}

// SetDomainOrigin - set the origin setting for the new
// For details, please refer https://cloud.baidu.com/doc/CDN-ABROAD/s/Gkbstcgaa
//
// PARAMS:
//     - domain: the specified domain
//     - origins: the origin servers
// RETURNS:
//     - error: nil if success otherwise the specific error
func (cli *Client) SetDomainOrigin(domain string, origins []api.OriginPeer) error {
	return api.SetDomainOrigin(cli, domain, origins)
}

// DomainOriginOption defined a method for setting optional origin configurations.
type DomainOriginOption func(interface{})

// SetDomainOriginWithOptions - set the origin setting for the new
// For details, please refer https://cloud.baidu.com/doc/CDN-ABROAD/s/Gkbstcgaa
//
// PARAMS:
//     - domain: the specified domain
//     - origins: the origin servers
//     - opts: optional configurations for origin, unused now!!!
// RETURNS:
//     - error: nil if success otherwise the specific error
func (cli *Client) SetDomainOriginWithOptions(domain string, origins []api.OriginPeer, opts ...DomainOriginOption) error {
	return api.SetDomainOrigin(cli, domain, origins)
}

// SetCacheTTL - add rules to cache asserts.
// For details, please refer https://cloud.baidu.com/doc/CDN-ABROAD/s/Zkbstm0vg
//
// PARAMS:
//     - domain: the specified domain
//     - cacheTTLs: the cache setting list
// RETURNS:
//     - error: nil if success otherwise the specific error
func (cli *Client) SetCacheTTL(domain string, cacheTTLs []api.CacheTTL) error {
	return api.SetCacheTTL(cli, domain, cacheTTLs)
}

// SetCacheFullUrl - set the rule to calculate the cache key, set `cacheFullUrl` to true
// means the whole query(the string right after the question mark in URL) will be added to the cache key.
// For details, please refer https://cloud.baidu.com/doc/CDN-ABROAD/s/nkbsxko6t
//
// PARAMS:
//     - domain: the specified domain
//     - cacheFullUrl: the query part in URL being added to the cache key or not
// RETURNS:
//     - error: nil if success otherwise the specific error
func (cli *Client) SetCacheFullUrl(domain string, cacheFullUrl bool) error {
	return api.SetCacheFullUrl(cli, domain, cacheFullUrl)
}

// SetHostToOrigin - Specify a default value for the HOST header for virtual sites in your origin server.
// For details, please refer https://cloud.baidu.com/doc/CDN-ABROAD/s/Pkbsxw8xw
//
// PARAMS:
//     - domain: the specified domain
//     - originHost: specified HOST header for origin server
// RETURNS:
//     - error: nil if success otherwise the specific error
func (cli *Client) SetHostToOrigin(domain string, originHost string) error {
	return api.SetHostToOrigin(cli, domain, originHost)
}

// SetRefererACL - Set a Referer whitelist or blacklist to enable hotlink protection.
// For details, please refer https://cloud.baidu.com/doc/CDN-ABROAD/s/ekbsxow69
//
// PARAMS:
//     - domain: the specified domain
//     - refererACL: referer of whitelist or blacklist
// RETURNS:
//     - error: nil if success otherwise the specific error
func (cli *Client) SetRefererACL(domain string, refererACL *api.RefererACL) error {
	return api.SetRefererACL(cli, domain, refererACL)
}

// SetIpACL - Set an IP whitelist or blacklist to block or allow requests from specific IP addresses.
// For details, please refer https://cloud.baidu.com/doc/CDN-ABROAD/s/2kbsxt693
//
// PARAMS:
//     - domain: the specified domain
//     - ipACL: IP whitelist or blacklist
// RETURNS:
//     - error: nil if success otherwise the specific error
func (cli *Client) SetIpACL(domain string, ipACL *api.IpACL) error {
	return api.SetIpACL(cli, domain, ipACL)
}

// Purge - tells the CDN system to purge the specified files
// For more details, please refer https://cloud.baidu.com/doc/CDN-ABROAD/s/Zkbsy0k8j
//
// PARAMS:
//     - tasks: the tasks about purging the files from the CDN nodes
// RETURNS:
//     - PurgedId: an ID representing a purged task, using it to search the task progress
//     - error: nil if success otherwise the specific error
func (cli *Client) Purge(tasks []api.PurgeTask) (api.PurgedId, error) {
	return api.Purge(cli, tasks)
}

// GetPurgedStatus - get the purged progress
// For details, please refer https://cloud.baidu.com/doc/CDN-ABROAD/s/ikbsy9cvb
//
// PARAMS:
//     - queryData: querying conditions, it contains the time interval, the task ID and the specified url
// RETURNS:
//     - *PurgedStatus: the details about the purged
//     - error: nil if success otherwise the specific error
func (cli *Client) GetPurgedStatus(queryData *api.CStatusQueryData) (*api.PurgedStatus, error) {
	return api.GetPurgedStatus(cli, queryData)
}

// Prefetch - tells the CDN system to prefetch the specified files
// For details, please refer https://cloud.baidu.com/doc/CDN-ABROAD/s/Dlj5ch09q
//
// PARAMS:
//     - tasks: the tasks about prefetch the files from the CDN nodes
//     - error: nil if success otherwise the specific error
func (cli *Client) Prefetch(tasks []api.PrefetchTask) (api.PrefetchId, error) {
	return api.Prefetch(cli, tasks)
}

// GetPrefetchStatus - get the prefetch progress
// For details, please refer https://cloud.baidu.com/doc/CDN-ABROAD/s/Mlj5e9y0i
//
// PARAMS:
//     - queryData: querying conditions, it contains the time interval, the task ID and the specified url
// RETURNS:
//     - *PrefetchStatus: the details about the prefetch
//     - error: nil if success otherwise the specific error
func (cli *Client) GetPrefetchStatus(queryData *api.CStatusQueryData) (*api.PrefetchStatus, error) {
	return api.GetPrefetchStatus(cli, queryData)
}

// GetQuota - get the quota about purge and prefetch
// For details, please refer https://cloud.baidu.com/doc/CDN-ABROAD/s/flnoakciq
//
// RETURNS:
//     - QuotaDetail: the quota details about a specified user
//     - error: nil if success otherwise the specific error
func (cli *Client) GetQuota() (*api.QuotaDetail, error) {
	return api.GetQuota(cli)
}

// GetDomainLog -get one domain's log urls
// For details, please refer https://cloud.baidu.com/doc/CDN-ABROAD/s/mjwvxj0ec
//
// PARAMS:
//     - domain: the specified domain
//     - timeInterval: the specified time interval
// RETURNS:
//     - []LogEntry: the log detail list
//     - error: nil if success otherwise the specific error
func (cli *Client) GetDomainLog(domain string, timeInterval api.TimeInterval) ([]api.LogEntry, error) {
	return api.GetDomainLog(cli, domain, timeInterval)
}

// GetFlow - get the statistics of traffic(flow).
// For details, please refer https://cloud.baidu.com/doc/CDN-ABROAD/s/Bkbszintg
//
// PARAMS:
//     - options: the querying conditions, valid options are:
//                1. QueryStatByTimeRange
//                2. QueryStatByPeriod
//                3. QueryStatByDomains
//                4. QueryStatByCountry
// RETURNS:
//     - []FlowDetail: the details about traffic
//     - error: nil if success otherwise the specific error
func (cli *Client) GetFlow(options ...api.QueryStatOption) ([]api.FlowDetail, error) {
	return api.GetFlow(cli, options...)
}

// GetPv - get the statistics of pv/qps.
// For details, please refer https://cloud.baidu.com/doc/CDN-ABROAD/s/dkbszg48s
//
// PARAMS:
//     - options: the querying conditions, valid options are:
//                1. QueryStatByTimeRange
//                2. QueryStatByPeriod
//                3. QueryStatByDomains
//                4. QueryStatByCountry
// RETURNS:
//     - []PvDetail: the details about pv/qps
//     - error: nil if success otherwise the specific error
func (cli *Client) GetPv(options ...api.QueryStatOption) ([]api.PvDetail, error) {
	return api.GetPv(cli, options...)
}

// GetSrcFlow - get the statistics of traffic(flow) to your origin server.
// For details, please refer https://cloud.baidu.com/doc/CDN-ABROAD/s/rkbsznt4v
//
// PARAMS:
//     - options: the querying conditions, valid options are:
//                1. QueryStatByTimeRange
//                2. QueryStatByPeriod
//                3. QueryStatByDomains
// RETURNS:
//     - []FlowDetail: the details about traffic to your origin server.
//     - error: nil if success otherwise the specific error
func (cli *Client) GetSrcFlow(options ...api.QueryStatOption) ([]api.FlowDetail, error) {
	return api.GetSrcFlow(cli, options...)
}

// GetHttpCode - get the statistics of accessing HTTP codes.
// For details, please refer https://cloud.baidu.com/doc/CDN-ABROAD/s/ekbszvxv5
//
// PARAMS:
//     - options: the querying conditions, valid options are:
//                1. QueryStatByTimeRange
//                2. QueryStatByPeriod
//                3. QueryStatByDomains
// RETURNS:
//     - []HttpCodeDetail: the details about accessing HTTP codes.
//     - error: nil if success otherwise the specific error
func (cli *Client) GetHttpCode(options ...api.QueryStatOption) ([]api.HttpCodeDetail, error) {
	return api.GetHttpCode(cli, options...)
}

// GetRealHit - get the statistics of hit rate of accessing by traffic.
// For details, please refer https://cloud.baidu.com/doc/CDN-ABROAD/s/ckbszuehh
//
// PARAMS:
//     - options: the querying conditions, valid options are:
//                1. QueryStatByTimeRange
//                2. QueryStatByPeriod
//                3. QueryStatByDomains
// RETURNS:
//     - []HitDetail: the details about traffic hit rate.
//     - error: nil if success otherwise the specific error
func (cli *Client) GetRealHit(options ...api.QueryStatOption) ([]api.HitDetail, error) {
	return api.GetRealHit(cli, options...)
}
