package cdn

import (
	"github.com/baidubce/bce-sdk-go/auth"
	"github.com/baidubce/bce-sdk-go/bce"
	"github.com/baidubce/bce-sdk-go/services/cdn/api"
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

// ListDomains - list all domains that in CDN service
// For details, please refer https://cloud.baidu.com/doc/CDN/s/sjwvyewt1
//
// PARAMS:
//     - marker: a marker is a start point of searching
// RETURNS:
//     - []string: domains belongs to the user
//     - string: a marker for next searching, empty if is in the end
//     - error: nil if success otherwise the specific error
func (cli *Client) ListDomains(marker string) ([]string, string, error) {
	return api.ListDomains(cli, marker)
}

// GetDomainStatus - get domains' details
// For details, please refer https://cloud.baidu.com/doc/CDN/s/8jwvyewf1
//
// PARAMS:
//     - status: the specified running status, the available values are "RUNNING", "STOPPED", OPERATING or "ALL"
//     - rule: the regex matching rule
// RETURNS:
//     - []DomainStatus: domain details list
//     - error: nil if success otherwise the specific error
func (cli *Client) GetDomainStatus(status string, rule string) ([]api.DomainStatus, error) {
	return api.GetDomainStatus(cli, status, rule)
}

// IsValidDomain - check the specified domain whether it can be added to CDN service or not.
// For details, please refer https://cloud.baidu.com/doc/CDN/s/qjwvyexh6
//
// PARAMS:
//     - domain: the specified domain
// RETURNS:
//     - *DomainValidInfo: available information about the specified domain
//     - error: nil if success otherwise the specific error
func (cli *Client) IsValidDomain(domain string) (*api.DomainValidInfo, error) {
	return api.IsValidDomain(cli, domain)
}

// CreateDomain - add a specified domain into CDN service
// For details, please refer https://cloud.baidu.com/doc/CDN/s/gjwvyex4o
//
// PARAMS:
//     - domain: the specified domain
//     - originInit: initialized data for a CDN domain
// RETURNS:
//     - *DomainCreatedInfo: the details about created a CDN domain
//     - error: nil if success otherwise the specific error
func (cli *Client) CreateDomain(domain string, originInit *api.OriginInit) (*api.DomainCreatedInfo, error) {
	return api.CreateDomain(cli, domain, originInit)
}

// EnableDomain - enable a specified domain
// For details, please refer https://cloud.baidu.com/doc/CDN/s/Jjwvyexv8
//
// PARAMS:
//     - domain: the specified domain
// RETURNS:
//     - error: nil if success otherwise the specific error
func (cli *Client) EnableDomain(domain string) error {
	return api.EnableDomain(cli, domain)
}

// DisableDomain - disable a specified domain
// For details, please refer https://cloud.baidu.com/doc/CDN/s/9jwvyew3e
//
// PARAMS:
//     - domain: the specified domain
// RETURNS:
//     - error: nil if success otherwise the specific error
func (cli *Client) DisableDomain(domain string) error {
	return api.DisableDomain(cli, domain)
}

// DeleteDomain - delete a specified domain from BCE CDN system.
// For details, please refer https://cloud.baidu.com/doc/CDN/s/Njwvyey7f
//
// PARAMS:
//     - domain: the specified domain
// RETURNS:
//     - error: nil if success otherwise the specific error
func (cli *Client) DeleteDomain(domain string) error {
	return api.DeleteDomain(cli, domain)
}

// GetIpInfo - retrieves information about the specified IP
// For details, please refer https://cloud.baidu.com/doc/CDN/s/8jwvyeunq
//
// PARAMS:
//     - ip: the specified ip addr
//     - action: the action for operating the ip addr
// RETURNS:
//     - *IpInfo: the information about the specified ip addr
//     - error: nil if success otherwise the specific error
func (cli *Client) GetIpInfo(ip string, action string) (*api.IpInfo, error) {
	return api.GetIpInfo(cli, ip, action)
}

// GetDomainConfig - get the configuration for the specified domain.
// For details, please refer https://cloud.baidu.com/doc/CDN/s/2jwvyf39o
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
// For details, please refer https://cloud.baidu.com/doc/CDN/s/xjxzi7729
//
// PARAMS:
//     - domain: the specified domain
//     - origins: the origin servers
//     - defaultHost: the default host
// RETURNS:
//     - error: nil if success otherwise the specific error
func (cli *Client) SetDomainOrigin(domain string, origins []api.OriginPeer, defaultHost string) error {
	return api.SetDomainOrigin(cli, domain, origins, defaultHost)
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
func (cli *Client) SetOriginProtocol(domain string, originProtocol string) error {
	return api.SetOriginProtocol(cli, domain, originProtocol)
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
func (cli *Client) GetOriginProtocol(domain string) (string, error) {
	return api.GetOriginProtocol(cli, domain)
}

// SetDomainSeo - set SEO setting
// For details, please refer https://cloud.baidu.com/doc/CDN/s/Jjxziuq4y
//
// PARAMS:
//     - domain: the specified domain
//     - seoSwitch: the setting about SEO
// RETURNS:
//     - error: nil if success otherwise the specific error
func (cli *Client) SetDomainSeo(domain string, seoSwitch *api.SeoSwitch) error {
	return api.SetDomainSeo(cli, domain, seoSwitch)
}

// GetDomainSeo - retrieve the setting about SEO
// There are two types of data that the server responds to
// 1. `{"seoSwitch":[]}` indicates no setting about SEO
// 2. `{"seoSwitch":{"diretlyOrigin":"ON","pushRecord":"OFF"}}` indicates it had normal setting about SEO.
// So the code need to handle the complex affairs
//
// For details, please refer https://cloud.baidu.com/doc/CDN/s/Djxzjfz8f
//
// PARAMS:
//     - domain: the specified domain
// RETURNS:
//     - *SeoSwitch: the setting about SEO
//     - error: nil if success otherwise the specific error
func (cli *Client) GetDomainSeo(domain string) (*api.SeoSwitch, error) {
	return api.GetDomainSeo(cli, domain)
}

// GetCacheTTL - get the current cached setting
// For details, please refer https://cloud.baidu.com/doc/CDN/s/ljxzhl9bu
//
// PARAMS:
//     - domain: the specified domain
// RETURNS:
//     - []CacheTTL: the cache setting list
//     - error: nil if success otherwise the specific error
func (cli *Client) GetCacheTTL(domain string) ([]api.CacheTTL, error) {
	return api.GetCacheTTL(cli, domain)
}

// SetCacheTTL - add some rules for cached setting
// For details, please refer https://cloud.baidu.com/doc/CDN/s/wjxzhgxnx
//
// PARAMS:
//     - domain: the specified domain
//     - cacheTTLs: the cache setting list
// RETURNS:
//     - error: nil if success otherwise the specific error
func (cli *Client) SetCacheTTL(domain string, cacheTTLs []api.CacheTTL) error {
	return api.SetCacheTTL(cli, domain, cacheTTLs)
}

// SetRefererACL - set a rule for filter some HTTP request, blackList and whiteList only one can be set
// For details, please refer https://cloud.baidu.com/doc/CDN/s/yjxzhvf21
//
// PARAMS:
//     - domain: the specified domain
//     - blackList: the forbidden host
//     - whiteList: the available host
// RETURNS:
//     - error: nil if success otherwise the specific error
func (cli *Client) SetRefererACL(domain string, blackList []string, whiteList []string, isAllowEmpty bool) error {
	return api.SetRefererACL(cli, domain, blackList, whiteList, isAllowEmpty)
}

// GetRefererACL - get referer ACL setting
// For details, please refer https://cloud.baidu.com/doc/CDN/s/Ujzkotvtb
//
// PARAMS:
//     - domain: the specified domain
// RETURNS:
//     - *api.RefererACL: referer ACL setting
//     - error: nil if success otherwise the specific error
func (cli *Client) GetRefererACL(domain string) (*api.RefererACL, error) {
	return api.GetRefererACL(cli, domain)
}

// SetRefererACL - set a rule for filter some HTTP request, blackList and whiteList only one can be set
// For details, please refer https://cloud.baidu.com/doc/CDN/s/8jxzhwc4d
//
// PARAMS:
//     - domain: the specified domain
//     - blackList: the forbidden ip
//     - whiteList: the available ip
// RETURNS:
//     - error: nil if success otherwise the specific error
func (cli *Client) SetIpACL(domain string, blackList []string, whiteList []string) error {
	return api.SetIpACL(cli, domain, blackList, whiteList)
}

// GetIpACL - get black IP or white IP
// For details, please refer https://cloud.baidu.com/doc/CDN/s/jjzkp5ku7
//
// PARAMS:
//     - domain: the specified domain
// RETURNS:
//     - *api.IpACL: ip setting
//     - error: nil if success otherwise the specific error
func (cli *Client) GetIpACL(domain string) (*api.IpACL, error) {
	return api.GetIpACL(cli, domain)
}

// SetLimitRate - set limited speed
// For details, please refer https://cloud.baidu.com/doc/CDN/s/Kjy6v02wt
//
// PARAMS:
//     - domain: the specified domain
//     - limitRate: the limited rate, "1024" means the transmittal speed is less than 1024 Byte/s
// RETURNS:
//     - error: nil if success otherwise the specific error
func (cli *Client) SetLimitRate(domain string, limitRate int) error {
	return api.SetLimitRate(cli, domain, limitRate)
}

// SetDomainHttps - set a rule for speed HTTPS' request
// For details, please refer https://cloud.baidu.com/doc/CDN/s/rjy6v3tnr
//
// PARAMS:
//     - domain: the specified domain
//     - httpsConfig: the rules about the HTTP configure
// RETURNS:
//     - error: nil if success otherwise the specific error
func (cli *Client) SetDomainHttps(domain string, httpsConfig *api.HTTPSConfig) error {
	return api.SetDomainHttps(cli, domain, httpsConfig)
}

// GetDomainHttps - get the setting about HTTPS
//
// PARAMS:
//     - domain: the specified domain
// RETURNS:
//     - *HTTPSConfig: the rules about the HTTP configure
//     - error: nil if success otherwise the specific error
func (cli *Client) GetDomainHttps(domain string) (*api.HTTPSConfig, error) {
	return api.GetDomainHttps(cli, domain)
}

// SetDomainRequestAuth - set the authorized rules for requesting
// For details, please refer https://cloud.baidu.com/doc/CDN/s/njxzi59g9
//
// PARAMS:
//     - domain: the specified domain
//     - requestAuth: the rules about the auth
// RETURNS:
//     - error: nil if success otherwise the specific error
func (cli *Client) SetDomainRequestAuth(domain string, requestAuth *api.RequestAuth) error {
	return api.SetDomainRequestAuth(cli, domain, requestAuth)
}

// Deprecated: We suggest use the SetOriginProtocol as an alternative
// SetFollowProtocol - set whether using the same protocol or not when back to the sourced server
// For details, please refer https://cloud.baidu.com/doc/CDN/s/9jxzi89k2
//
// PARAMS:
//     - domain: the specified domain
//     - isFollowProtocol: true in using the same protocol or not when back to the sourced server, false for other
// RETURNS:
//     - error: nil if success otherwise the specific error
func (cli *Client) SetFollowProtocol(domain string, isFollowProtocol bool) error {
	return api.SetFollowProtocol(cli, domain, isFollowProtocol)
}

// SetHttpHeader -set some HTTP headers which can be added or deleted when response form CDN edge node
// For details, please refer https://cloud.baidu.com/doc/CDN/s/Jjxzil1sd
//
// PARAMS:
//     - domain: the specified domain
//     - httpHeaders: the HTTP headers' setting
// RETURNS:
//     - error: nil if success otherwise the specific error
func (cli *Client) SetHttpHeader(domain string, httpHeaders []api.HttpHeader) error {
	return api.SetHttpHeader(cli, domain, httpHeaders)
}

// GetHttpHeader - get the HTTP headers' setting
// For details, please refer https://cloud.baidu.com/doc/CDN/s/6jxzip3wn
//
// PARAMS:
//     - domain: the specified domain
// RETURNS:
//     - error: nil if success otherwise the specific error
//     - []HttpHeader: the HTTP headers in setting
func (cli *Client) GetHttpHeader(domain string) ([]api.HttpHeader, error) {
	return api.GetHttpHeader(cli, domain)
}

// SetErrorPage - set the page that redirected to when error occurred
// For details, please refer https://cloud.baidu.com/doc/CDN/s/Ejy6vc4yb
//
// PARAMS:
//     - domain: the specified domain
//     - errorPages: the custom pages' setting
// RETURNS:
//     - error: nil if success otherwise the specific error
func (cli *Client) SetErrorPage(domain string, errorPages []api.ErrorPage) error {
	return api.SetErrorPage(cli, domain, errorPages)
}

// GetErrorPage - get the custom pages' setting
// For details, please refer https://cloud.baidu.com/doc/CDN/s/qjy6vfk2u
//
// PARAMS:
//     - domain: the specified domain
// RETURNS:
//     - []ErrorPage: the pages' setting
//     - error: nil if success otherwise the specific error
func (cli *Client) GetErrorPage(domain string) ([]api.ErrorPage, error) {
	return api.GetErrorPage(cli, domain)
}

// SetMediaDrag - set the media setting about mp4 and flv
// For details, please refer https://cloud.baidu.com/doc/CDN/s/4jy6v6xk3
//
// PARAMS:
//     - domain: the specified domain
//     - mediaDragConf: media setting about mp4 and flv
// RETURNS:
//     - error: nil if success otherwise the specific error
func (cli *Client) SetMediaDrag(domain string, mediaDragConf *api.MediaDragConf) error {
	return api.SetMediaDrag(cli, domain, mediaDragConf)
}

// GetMediaDrag - get the media setting
// For details, please refer https://cloud.baidu.com/doc/CDN/s/Ojy6v9q8f
//
// PARAMS:
//     - domain: the specified domain
// RETURNS:
//     - *MediaDragConf: the media setting about mp4 and flv
//     - error: nil if success otherwise the specific error
func (cli *Client) GetMediaDrag(domain string) (*api.MediaDragConf, error) {
	return api.GetMediaDrag(cli, domain)
}

// SetFileTrim - trim the text file or not
// For details, please refer https://cloud.baidu.com/doc/CDN/s/Xjy6vimct
//
// PARAMS:
//     - domain: the specified domain
//     - fileTrim: true means trimming the text file, false means do nothing
// RETURNS:
//     - error: nil if success otherwise the specific error
func (cli *Client) SetFileTrim(domain string, fileTrim bool) error {
	return api.SetFileTrim(cli, domain, fileTrim)
}

// GetFileTrim - get the trim setting about text file
// For details, please refer https://cloud.baidu.com/doc/CDN/s/Ujy6vjxnl
//
// PARAMS:
//     - domain: the specified domain
// RETURNS:
//     - bool: true means trimming the text file, false means do nothing
//     - error: nil if success otherwise the specific error
func (cli *Client) GetFileTrim(domain string) (bool, error) {
	return api.GetFileTrim(cli, domain)
}

// SetMobileAccess - distinguish the client or not
// For details, please refer https://cloud.baidu.com/doc/CDN/s/Mjy6vmv6g
//
// PARAMS:
//     - domain: the specified domain
//     - distinguishClient: true means distinguishing the client, false means not
// RETURNS:
//     - error: nil if success otherwise the specific error
func (cli *Client) SetMobileAccess(domain string, distinguishClient bool) error {
	return api.SetMobileAccess(cli, domain, distinguishClient)
}

// GetMobileAccess - get the setting about distinguishing the client or not
// For details, please refer https://cloud.baidu.com/doc/CDN/s/Mjy6vo9z2
//
// PARAMS:
//     - domain: the specified domain
// RETURNS:
//     - bool: true means distinguishing the client, false means not
//     - error: nil if success otherwise the specific error
func (cli *Client) GetMobileAccess(domain string) (bool, error) {
	return api.GetMobileAccess(cli, domain)
}

// SetClientIp - set the specified HTTP header for the origin server
// For details, please refer https://cloud.baidu.com/doc/CDN/s/Kjy6umyrm
//
// PARAMS:
//     - domain: the specified domain
//     - clientIp: header setting
// RETURNS:
//     - error: nil if success otherwise the specific error
func (cli *Client) SetClientIp(domain string, clientIp *api.ClientIp) error {
	return api.SetClientIp(cli, domain, clientIp)
}

// GetClientIp - get the setting about getting client IP
// For details, please refer https://cloud.baidu.com/doc/CDN/s/8jy6urcq5
//
// PARAMS:
//     - domain: the specified domain
// RETURNS:
//     - *ClientIp: the HTTP header setting for origin server to get client IP
//     - error: nil if success otherwise the specific error
func (cli *Client) GetClientIp(domain string) (*api.ClientIp, error) {
	return api.GetClientIp(cli, domain)
}

// SetAccessLimit - set the qps for on one client
// For details, please refer https://cloud.baidu.com/doc/CDN/s/Kjy6v02wt
//
// PARAMS:
//     - domain: the specified domain
//     - accessLimit: the access setting
// RETURNS:
//     - error: nil if success otherwise the specific error
func (cli *Client) SetAccessLimit(domain string, accessLimit *api.AccessLimit) error {
	return api.SetAccessLimit(cli, domain, accessLimit)
}

// GetAccessLimit - get the qps setting
// For details, please refer https://cloud.baidu.com/doc/CDN/s/rjy6v3tnr
//
// PARAMS:
//     - domain: the specified domain
// RETURNS:
//     - *AccessLimit: the access setting
//     - error: nil if success otherwise the specific error
func (cli *Client) GetAccessLimit(domain string) (*api.AccessLimit, error) {
	return api.GetAccessLimit(cli, domain)
}

// SetCacheUrlArgs - tell the CDN system cache the url's params or not
// For details, please refer https://cloud.baidu.com/doc/CDN/s/vjxzho0kx
//
// PARAMS:
//     - domain: the specified domain
//     - cacheFullUrl: whether cache the full url or not, full url means include params, also some extra params can be avoided
// RETURNS:
//     - error: nil if success otherwise the specific error
func (cli *Client) SetCacheUrlArgs(domain string, cacheFullUrl *api.CacheUrlArgs) error {
	return api.SetCacheUrlArgs(cli, domain, cacheFullUrl)
}

// GetCacheUrlArgs - get the cached rules
// For details, please refer https://cloud.baidu.com/doc/CDN/s/sjxzhsb6h
//
// PARAMS:
//     - domain: the specified domain
// RETURNS:
//     - *CacheUrlArgs: the details about cached rules
//     - error: nil if success otherwise the specific error
func (cli *Client) GetCacheUrlArgs(domain string) (*api.CacheUrlArgs, error) {
	return api.GetCacheUrlArgs(cli, domain)
}

// SetCors - set about Cross-origin resource sharing
// For details, please refer https://cloud.baidu.com/doc/CDN/s/Rjxzi1cfs
// PARAMS:
//     - domain: the specified domain
//     - isAllow: true means allow Cors, false means not allow
//     - originList: the origin setting, it's invalid when isAllow is false
// RETURNS:
//     - error: nil if success otherwise the specific error
func (cli *Client) SetCors(domain string, isAllow bool, originList []string) error {
	return api.SetCors(cli, domain, isAllow, originList)
}

// GetCors - get the Cors setting
// For details, please refer https://cloud.baidu.com/doc/CDN/s/tjxzi2d7t
//
// PARAMS:
//     - domain: the specified domain
// RETURNS:
//     - *CorsCfg: the Cors setting
//     - error: nil if success otherwise the specific error
func (cli *Client) GetCors(domain string) (*api.CorsCfg, error) {
	return api.GetCors(cli, domain)
}

// SetRangeSwitch - set the range setting
// For details, please refer https://cloud.baidu.com/doc/CDN/s/Fjxziabst
//
// PARAMS:
//     - domain: the specified domain
//     - enabled: true means enable range cached, false means disable range cached
// RETURNS:
//     - error: nil if success otherwise the specific error
func (cli *Client) SetRangeSwitch(domain string, enabled bool) error {
	return api.SetRangeSwitch(cli, domain, enabled)
}

// GetRangeSwitch - get the range setting
// For details, please refer https://cloud.baidu.com/doc/CDN/s/5jxzid6o9
//
// PARAMS:
//     - domain: the specified domain
// RETURNS:
//     - bool: true means enable range cached, false means disable range cached
//     - error: nil if success otherwise the specific error
func (cli *Client) GetRangeSwitch(domain string) (bool, error) {
	return api.GetRangeSwitch(cli, domain)
}

// SetContentEncoding - set Content-Encoding
// For details, please refer https://cloud.baidu.com/doc/CDN/s/0jyqyahsb
//
// PARAMS:
//     - domain: the specified domain
//     - enabled: true means using the specified encoding algorithm indicated by "encodingType" in transferring,
//         false means disable encoding
// RETURNS:
//     - error: nil if success otherwise the specific error
func (cli *Client) SetContentEncoding(domain string, enabled bool, encodingType string) error {
	return api.SetContentEncoding(cli, domain, enabled, encodingType)
}

// GetContentEncoding - get the setting about Content-Encoding
// For details, please refer https://cloud.baidu.com/doc/CDN/s/bjyqycw8g
//
// PARAMS:
//     - domain: the specified domain
// RETURNS:
//     - string: the encoding algorithm for transferring, empty means disable encoding in transferring
//     - error: nil if success otherwise the specific error
func (cli *Client) GetContentEncoding(domain string) (string, error) {
	return api.GetContentEncoding(cli, domain)
}

// Purge - tells the CDN system to purge the specified files
// For more details, please refer https://cloud.baidu.com/doc/CDN/s/ijwvyeyyj
//
// PARAMS:
//     - tasks: the tasks about purging the files from the CDN nodes
// RETURNS:
//     - PurgedId: an ID representing a purged task, using it to search the task progress
//     - error: nil if success otherwise the specific error.
func (cli *Client) Purge(tasks []api.PurgeTask) (api.PurgedId, error) {
	return api.Purge(cli, tasks)
}

// GetPurgedStatus - get the purged progress
// For details, please refer https://cloud.baidu.com/doc/CDN/s/ujwvyezqm
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
// For details, please refer https://cloud.baidu.com/doc/CDN/s/Rjwvyf0ff
//
// PARAMS:
//     - tasks: the tasks about prefetch the files from the CDN nodes
//     - error: nil if success otherwise the specific error
func (cli *Client) Prefetch(tasks []api.PrefetchTask) (api.PrefetchId, error) {
	return api.Prefetch(cli, tasks)
}

// GetPrefetchStatus - get the prefetch progress
// For details, please refer https://cloud.baidu.com/doc/CDN/s/4jwvyf01w
//
// PARAMS:
//     - queryData: querying conditions, it contains the time interval, the task ID and the specified url.
// RETURNS:
//     - *PrefetchStatus: the details about the prefetch
//     - error: nil if success otherwise the specific error
func (cli *Client) GetPrefetchStatus(queryData *api.CStatusQueryData) (*api.PrefetchStatus, error) {
	return api.GetPrefetchStatus(cli, queryData)
}

// GetQuota - get the quota about purge and prefetch
// For details, please refer https://cloud.baidu.com/doc/CDN/s/zjwvyeze3
//
// RETURNS:
//     - QuotaDetail: the quota details about a specified user
//     - error: nil if success otherwise the specific error
func (cli *Client) GetQuota() (*api.QuotaDetail, error) {
	return api.GetQuota(cli)
}

// GetCacheOpRecords get the history operating records
// For details, please refer https://cloud.baidu.com/doc/CDN/s/5jypnzjqt
//
// PARAMS:
//     - queryData: querying conditions, it contains the time interval, the task type and the specified url
// RETURNS:
//     - *RecordDetails: the details about the records
//     - error: nil if success otherwise the specific error
func (cli *Client) GetCacheOpRecords(queryData *api.CRecordQueryData) (*api.RecordDetails, error) {
	return api.GetCacheOpRecords(cli, queryData)
}

// EnableDsa - enable DSA
// For details, please refer https://cloud.baidu.com/doc/CDN/s/7jwvyf1h5
//
// RETURNS:
//     - error: nil if success otherwise the specific error
func (cli *Client) EnableDsa() error {
	return api.EnableDsa(cli)
}

// DisableDsa - disable DSA
// For details, please refer https://cloud.baidu.com/doc/CDN/s/7jwvyf1h5
//
// RETURNS:
//     - error: nil if success otherwise the specific error
func (cli *Client) DisableDsa() error {
	return api.DisableDsa(cli)
}

// ListDsaDomains - retrieve all the domains in DSA service
// For details, please refer https://cloud.baidu.com/doc/CDN/s/5jwvyf1sq
//
// RETURNS:
//     - []DSADomain: the details about DSA domains
//     - error: nil if success otherwise the specific error
func (cli *Client) ListDsaDomains() ([]api.DSADomain, error) {
	return api.ListDsaDomains(cli)
}

// SetDsaConfig - set the DSA configuration
// For details, please refer https://cloud.baidu.com/doc/CDN/s/0jwvyf26d
//
// PARAMS:
//     - domain: the specified domain
//     - dsaConfig: the specified configuration for the specified domain
// RETURNS:
//     - error: nil if success otherwise the specific error
func (cli *Client) SetDsaConfig(domain string, dsaConfig *api.DSAConfig) error {
	return api.SetDsaConfig(cli, domain, dsaConfig)
}

// GetDomainLog -get one domain's log urls
// For details, please refer https://cloud.baidu.com/doc/CDN/s/cjwvyf0r9
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

// GetMultiDomainLog - get many domains' log urls
// For details, please refer https://cloud.baidu.com/doc/CDN/s/cjwvyf0r9
//
// PARAMS:
//     - queryData: the querying conditions
//     - error: nil if success otherwise the specific error
func (cli *Client) GetMultiDomainLog(queryData *api.LogQueryData) ([]api.LogEntry, error) {
	return api.GetMultiDomainLog(cli, queryData)
}

// GetAvgSpeed - get the average speed
// For details, please refer https://cloud.baidu.com/doc/CDN/s/5jwvyf8zn#%E6%9F%A5%E8%AF%A2%E5%B9%B3%E5%9D%87%E9%80%9F%E7%8E%87
//
// PARAMS:
//     - queryCondition: the querying conditions
// RETURNS:
//     - []AvgSpeedDetail: the detail list about the average speed
//     - error: nil if success otherwise the specific error
func (cli *Client) GetAvgSpeed(queryCondition *api.QueryCondition) ([]api.AvgSpeedDetail, error) {
	return api.GetAvgSpeed(cli, queryCondition)
}

// GetAvgSpeed - get the average speed filter by location
// For details, please refer https://cloud.baidu.com/doc/CDN/s/5jwvyf8zn#%E5%AE%A2%E6%88%B7%E7%AB%AF%E8%AE%BF%E9%97%AE%E5%88%86%E5%B8%83%E6%9F%A5%E8%AF%A2%E5%B9%B3%E5%9D%87%E9%80%9F%E7%8E%87
//
// PARAMS:
//     - queryCondition: the querying conditions
//     - prov: the specified area, like "beijing"
//     - isp: the specified ISP, like "ct"
// RETURNS:
//     - []AvgSpeedRegionDetail: the detail list about the average speed
//     - error: nil if success otherwise the specific error
func (cli *Client) GetAvgSpeedByRegion(queryCondition *api.QueryCondition, prov string, isp string) ([]api.AvgSpeedRegionDetail, error) {
	return api.GetAvgSpeedByRegion(cli, queryCondition, prov, isp)
}

// GetPv - get the PV data
// For details, please refer https://cloud.baidu.com/doc/CDN/s/5jwvyf8zn#pvqps%E6%9F%A5%E8%AF%A2
//
// PARAMS:
//     - queryCondition: the querying conditions
//     - level: the node level, the available values are "edge", "internal" and "all"
// RETURNS:
//     - []PvDetail: the detail list about page view
//     - error: nil if success otherwise the specific error
func (cli *Client) GetPv(queryCondition *api.QueryCondition, level string) ([]api.PvDetail, error) {
	return api.GetPv(cli, queryCondition, level)
}

// GetSrcPv - get the PV data in back to the sourced server
// For details, please refer https://cloud.baidu.com/doc/CDN/s/5jwvyf8zn#%E5%9B%9E%E6%BA%90pvqps%E6%9F%A5%E8%AF%A2
//
// PARAMS:
//     - queryCondition: the querying conditions
// RETURNS:
//     - []PvDetail: the detail list about page view
//     - error: nil if success otherwise the specific error
func (cli *Client) GetSrcPv(queryCondition *api.QueryCondition) ([]api.PvDetail, error) {
	return api.GetSrcPv(cli, queryCondition)
}

// GetAvgPvByRegion - get the PV data filter by location
// For details, please refer https://cloud.baidu.com/doc/CDN/s/5jwvyf8zn#%E6%9F%A5%E8%AF%A2pvqps%E5%88%86%E5%AE%A2%E6%88%B7%E7%AB%AF%E8%AE%BF%E9%97%AE%E5%88%86%E5%B8%83
//
// PARAMS:
//     - queryCondition: the querying conditions
//     - prov: the specified area, like "beijing"
//     - isp: the specified ISP, like "ct"
// RETURNS:
//     - []PvRegionDetail: the detail list about page view
//     - error: nil if success otherwise the specific error
func (cli *Client) GetPvByRegion(queryCondition *api.QueryCondition, prov string, isp string) ([]api.PvRegionDetail, error) {
	return api.GetPvByRegion(cli, queryCondition, prov, isp)
}

// GetUv - get the UV data
// For details, please refer https://cloud.baidu.com/doc/CDN/s/5jwvyf8zn#uv%E6%9F%A5%E8%AF%A2
//
// PARAMS:
//     - queryCondition: the querying conditions
// RETURNS:
//     - []UvDetail: the detail list about unique visitor
//     - error: nil if success otherwise the specific error
func (cli *Client) GetUv(queryCondition *api.QueryCondition) ([]api.UvDetail, error) {
	return api.GetUv(cli, queryCondition)
}

// GetFlow - get the flow data
// For details, please refer https://cloud.baidu.com/doc/CDN/s/5jwvyf8zn#%E6%9F%A5%E8%AF%A2%E6%B5%81%E9%87%8F%E3%80%81%E5%B8%A6%E5%AE%BD
//
// PARAMS:
//     - queryCondition: the querying conditions
// RETURNS:
//     - []FlowDetail: the detail list about flow
//     - error: nil if success otherwise the specific error
func (cli *Client) GetFlow(queryCondition *api.QueryCondition, level string) ([]api.FlowDetail, error) {
	return api.GetFlow(cli, queryCondition, level)
}

// GetFlowByProtocol - get the flow data filter by protocol
// For details, please refer https://cloud.baidu.com/doc/CDN/s/5jwvyf8zn#%E6%9F%A5%E8%AF%A2%E6%B5%81%E9%87%8F%E3%80%81%E5%B8%A6%E5%AE%BD%E5%88%86%E5%8D%8F%E8%AE%AE
//
// PARAMS:
//     - queryCondition: the querying conditions
//     - protocol: the specified HTTP protocol, like "http" or "https", "all" means both "http" and "https"
// RETURNS:
//     - []FlowDetail: the detail list about flow
//     - error: nil if success otherwise the specific error
func (cli *Client) GetFlowByProtocol(queryCondition *api.QueryCondition, protocol string) ([]api.FlowDetail, error) {
	return api.GetFlowByProtocol(cli, queryCondition, protocol)
}

// GetFlowByRegion - get the flow data filter by location.
// For details, please refer https://cloud.baidu.com/doc/CDN/s/5jwvyf8zn#%E6%9F%A5%E8%AF%A2%E6%B5%81%E9%87%8F%E3%80%81%E5%B8%A6%E5%AE%BD%EF%BC%88%E5%88%86%E5%AE%A2%E6%88%B7%E7%AB%AF%E8%AE%BF%E9%97%AE%E5%88%86%E5%B8%83%EF%BC%89
//
// PARAMS:
//     - queryCondition: the querying conditions
//     - prov: the specified area, like "beijing"
//     - isp: the specified ISP, like "ct"
// RETURNS:
//     - []FlowRegionDetail: the detail list about flow
//     - error: nil if success otherwise the specific error
func (cli *Client) GetFlowByRegion(queryCondition *api.QueryCondition, prov string, isp string) ([]api.FlowRegionDetail, error) {
	return api.GetFlowByRegion(cli, queryCondition, prov, isp)
}

// GetSrcFlow - get the flow data in backed to sourced server
// For details, please refer https://cloud.baidu.com/doc/CDN/s/5jwvyf8zn#%E6%9F%A5%E8%AF%A2%E5%9B%9E%E6%BA%90%E6%B5%81%E9%87%8F%E3%80%81%E5%9B%9E%E6%BA%90%E5%B8%A6%E5%AE%BD
//
// PARAMS:
//     - queryCondition: the querying conditions
// RETURNS:
//     - []FlowDetail: the detail list about flow
//     - error: nil if success otherwise the specific error
func (cli *Client) GetSrcFlow(queryCondition *api.QueryCondition) ([]api.FlowDetail, error) {
	return api.GetSrcFlow(cli, queryCondition)
}

// GetRealHit - get the detail about byte hit rate
// For details, please refer https://cloud.baidu.com/doc/CDN/s/5jwvyf8zn#%E5%AD%97%E8%8A%82%E5%91%BD%E4%B8%AD%E7%8E%87%E6%9F%A5%E8%AF%A2
//
// PARAMS:
//     - queryCondition: the querying conditions
// RETURNS:
//     - []HitDetail: the detail list about byte rate
//     - error: nil if success otherwise the specific error
func (cli *Client) GetRealHit(queryCondition *api.QueryCondition) ([]api.HitDetail, error) {
	return api.GetRealHit(cli, queryCondition)
}

// GetPvHit - get the detail about PV hit rate.
// For details, please refer https://cloud.baidu.com/doc/CDN/s/5jwvyf8zn#%E8%AF%B7%E6%B1%82%E5%91%BD%E4%B8%AD%E7%8E%87%E6%9F%A5%E8%AF%A2
//
// PARAMS:
//     - queryCondition: the querying conditions
// RETURNS:
//     - []HitDetail: the detail list about pv rate
//     - error: nil if success otherwise the specific error
func (cli *Client) GetPvHit(queryCondition *api.QueryCondition) ([]api.HitDetail, error) {
	return api.GetPvHit(cli, queryCondition)
}

// GetHttpCode - get the http code's statistics
// For details, please refer https://cloud.baidu.com/doc/CDN/s/5jwvyf8zn#%E7%8A%B6%E6%80%81%E7%A0%81%E7%BB%9F%E8%AE%A1%E6%9F%A5%E8%AF%A2
//
// PARAMS:
//     - queryCondition: the querying conditions
// RETURNS:
//     - []HttpCodeDetail: the detail list about http code
//     - error: nil if success otherwise the specific error
func (cli *Client) GetHttpCode(queryCondition *api.QueryCondition) ([]api.HttpCodeDetail, error) {
	return api.GetHttpCode(cli, queryCondition)
}

// GetSrcHttpCode - get the http code's statistics in backed to sourced server
// For details, please refer https://cloud.baidu.com/doc/CDN/s/5jwvyf8zn#%E5%9B%9E%E6%BA%90%E7%8A%B6%E6%80%81%E7%A0%81%E6%9F%A5%E8%AF%A2
//
// PARAMS:
//     - queryCondition: the querying conditions
// RETURNS:
//     - []HttpCodeDetail: the detail list about http code
//     - error: nil if success otherwise the specific error
func (cli *Client) GetSrcHttpCode(queryCondition *api.QueryCondition) ([]api.HttpCodeDetail, error) {
	return api.GetSrcHttpCode(cli, queryCondition)
}

// GetHttpCodeByRegion - get the http code's statistics filter by location
// For details, please refer https://cloud.baidu.com/doc/CDN/s/5jwvyf8zn#%E7%8A%B6%E6%80%81%E7%A0%81%E7%BB%9F%E8%AE%A1%E6%9F%A5%E8%AF%A2%EF%BC%88%E5%88%86%E5%AE%A2%E6%88%B7%E7%AB%AF%E8%AE%BF%E9%97%AE%E5%88%86%E5%B8%83%EF%BC%89
//
// PARAMS:
//     - queryCondition: the querying conditions
//     - prov: the specified area, like "beijing"
//     - isp: the specified ISP, like "ct"
// RETURNS:
//     - []HttpCodeRegionDetail: the detail list about http code
//     - error: nil if success otherwise the specific error
func (cli *Client) GetHttpCodeByRegion(queryCondition *api.QueryCondition, prov string, isp string) ([]api.HttpCodeRegionDetail, error) {
	return api.GetHttpCodeByRegion(cli, queryCondition, prov, isp)
}

// GetTopNUrls - get the top N urls that requested
// For details, please refer https://cloud.baidu.com/doc/CDN/s/5jwvyf8zn#topn-urls
//
// PARAMS:
//     - queryCondition: the querying conditions
//     - httpCode: the specified HTTP code, like "200"
// RETURNS:
//     - []TopNDetail: the top N urls' detail
//     - error: nil if success otherwise the specific error
func (cli *Client) GetTopNUrls(queryCondition *api.QueryCondition, httpCode string) ([]api.TopNDetail, error) {
	return api.GetTopNUrls(cli, queryCondition, httpCode)
}

// GetTopNReferers - get the top N urls that brought by requested
// For details, please refer https://cloud.baidu.com/doc/CDN/s/5jwvyf8zn#topn-referers
//
// PARAMS:
//     - queryCondition: the querying conditions
//     - httpCode: the specified HTTP code, like "200"
// RETURNS:
//     - []TopNDetail: the top N referer urls' detail
//     - error: nil if success otherwise the specific error
func (cli *Client) GetTopNReferers(queryCondition *api.QueryCondition, httpCode string) ([]api.TopNDetail, error) {
	return api.GetTopNReferers(cli, queryCondition, httpCode)
}

// GetTopNDomains - get the top N domains that equested
// For details, please refer https://cloud.baidu.com/doc/CDN/s/5jwvyf8zn#topn-domains
//
// PARAMS:
//     - queryCondition: the querying conditions
//     - httpCode: the specified HTTP code, like "200"
// RETURNS:
//     - []TopNDetail: the top N domains' detail
//     - error: nil if success otherwise the specific error
func (cli *Client) GetTopNDomains(queryCondition *api.QueryCondition, httpCode string) ([]api.TopNDetail, error) {
	return api.GetTopNDomains(cli, queryCondition, httpCode)
}

// GetError - get the error code's data
// For details, please refer https://cloud.baidu.com/doc/CDN/s/5jwvyf8zn#cdn%E9%94%99%E8%AF%AF%E7%A0%81%E5%88%86%E7%B1%BB%E7%BB%9F%E8%AE%A1%E6%9F%A5%E8%AF%A2
//
// PARAMS:
//     - queryCondition: the querying conditions
// RETURNS:
//     - []ErrorDetail: the top N error details
//     - error: nil if success otherwise the specific error
func (cli *Client) GetError(queryCondition *api.QueryCondition) ([]api.ErrorDetail, error) {
	return api.GetError(cli, queryCondition)
}

// GetPeak95Bandwidth - get peak 95 bandwidth for the specified tags or domains.
// For details, pleader refer https://cloud.baidu.com/doc/CDN/s/5jwvyf8zn#%E6%9F%A5%E8%AF%A2%E6%9C%8895%E5%B3%B0%E5%80%BC%E5%B8%A6%E5%AE%BD
//
// PARAMS:
//     - startTime: start time which in `YYYY-mm-ddTHH:ii:ssZ` style
//     - endTime: end time which in `YYYY-mm-ddTHH:ii:ssZ` style
//     - domains: a list of domains, only one of `tags` and `domains` can contains item
//     - tags: a list of tag names, only one of `tags` and `domains` can contains item
// RETURNS:
//     - string: the peak95 time which in `YYYY-mm-ddTHH:ii:ssZ` style
//     - int64: peak95 bandwidth
//     - error: nil if success otherwise the specific error
func (cli *Client) GetPeak95Bandwidth(startTime, endTime string, domains, tags []string) (string, int64, error) {
	return api.GetPeak95Bandwidth(cli, startTime, endTime, domains, tags)
}
