package eo

import (
	"github.com/baidubce/bce-sdk-go/auth"
	"github.com/baidubce/bce-sdk-go/bce"
	"github.com/baidubce/bce-sdk-go/services/eo/api"
)

const (
	DEFAULT_SERVICE_DOMAIN = "geo.baidubce.com"
)

// Client of EO service is a kind of BceClient, so derived from BceClient
type Client struct {
	*bce.BceClient
}

// NewClient make the EO service client with default configuration
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
//   - method: the HTTP requested method, e.g. "GET", "POST", "PUT" ...
//   - urlPath: a path component, consisting of a sequence of path segments separated by a slash ( / ).
//   - params: the query params, which will be append to the query path, and separate by "&"
//     e.g. http://www.baidu.com?query_param1=value1&query_param2=value2
//   - reqHeaders: the request http headers
//   - bodyObj: the HTTP requested body content transferred to a goland object
//   - respObj: the HTTP response content transferred to a goland object
//
// RETURNS:
//   - error: nil if success otherwise the specific error
func (cli *Client) SendCustomRequest(method string, urlPath string, params, reqHeaders map[string]string, bodyObj interface{}, respObj interface{}) error {
	return api.SendCustomRequest(cli, method, urlPath, params, reqHeaders, bodyObj, respObj)
}

// Purge - tells the EO system to purge the specified files
// For more details, please refer https://cloud.baidu.com/doc/GEO/s/4mhsrv9ry
//
// PARAMS:
//   - site: the site that the purge tasks belong to
//   - tasks: the tasks about purging the files from the EO nodes
//
// RETURNS:
//   - PurgedId: an ID representing a purged task, using it to search the task progress
//   - error: nil if success otherwise the specific error
func (cli *Client) Purge(site string, tasks []api.PurgeTask) (api.PurgedId, error) {
	return api.Purge(cli, site, tasks)
}

// GetPurgedStatus - get the purged progress
// For more details, please refer https://cloud.baidu.com/doc/GEO/s/mmhssw91q
//
// PARAMS:
//   - queryData: querying conditions, it contains the site, the time interval, the task ID, the type and the marker
//
// RETURNS:
//   - *PurgeRecords: the details about the purged
//   - error: nil if success otherwise the specific error
func (cli *Client) GetPurgedStatus(queryData *api.PurgeStatusQueryData) (*api.PurgeRecords, error) {
	return api.GetPurgedStatus(cli, queryData)
}

// Prefetch - tells the EO system to prefetch the specified files
// For more details, please refer https://cloud.baidu.com/doc/GEO/s/5mhsuituv
//
// PARAMS:
//   - site: the site that the prefetch tasks belong to
//   - tasks: the tasks about prefetch the files from the EO nodes
//
// RETURNS:
//   - PrefetchId: an ID representing a prefetch task, using it to search the task progress
//   - error: nil if success otherwise the specific error
func (cli *Client) Prefetch(site string, tasks []api.PrefetchTask) (api.PrefetchId, error) {
	return api.Prefetch(cli, site, tasks)
}

// GetPrefetchStatus - get the prefetch progress
// For more details, please refer https://cloud.baidu.com/doc/GEO/s/Bmhsv5i9u
//
// PARAMS:
//   - queryData: querying conditions, it contains the site, the time interval, the task ID and the marker
//
// RETURNS:
//   - *PrefetchRecords: the details about the prefetch
//   - error: nil if success otherwise the specific error
func (cli *Client) GetPrefetchStatus(queryData *api.PrefetchStatusQueryData) (*api.PrefetchRecords, error) {
	return api.GetPrefetchStatus(cli, queryData)
}

// GetOfflineLog - get the offline log download urls of one or multi domains under a site
//
// PARAMS:
//   - queryData: querying conditions, it contains the site, the time interval, the domain list and pagination
//
// RETURNS:
//   - *LogQueryResult: the offline log entries and total count
//   - error: nil if success otherwise the specific error
func (cli *Client) GetOfflineLog(queryData *api.LogQueryData) (*api.LogQueryResult, error) {
	return api.GetOfflineLog(cli, queryData)
}

// GetStatTime - type 1: showType=time , no group
// For details, please refer https://cloud.baidu.com/doc/GEO/s/4mp0unbx3
//
// PARAMS:
//   - queryData: the querying conditions, `Metric` is required
//
// RETURNS:
//   - api.StatResult: a map keyed by metric name, value is the data point list
//   - error: nil if success otherwise the specific error
func (cli *Client) GetStatTime(queryData *api.StatTimeQueryData) (api.StatResult, error) {
	return api.GetStatTime(cli, queryData)
}

// GetStatPeak - type 2: showType=peak, no group
// For details, please refer https://cloud.baidu.com/doc/GEO/s/4mp0unbx3
//
// PARAMS:
//   - queryData: the querying conditions, `Metric` is required
//
// RETURNS:
//   - api.StatResult: a map keyed by metric name, value is the data point list
//   - error: nil if success otherwise the specific error
func (cli *Client) GetStatPeak(queryData *api.StatPeakQueryData) (api.StatResult, error) {
	return api.GetStatPeak(cli, queryData)
}

// GetStatTimeByGroup - type 3: showType=time, group required
// For details, please refer https://cloud.baidu.com/doc/GEO/s/4mp0unbx3
//
// PARAMS:
//   - queryData: the querying conditions, `Metric` and `Group` are required
//
// RETURNS:
//   - api.StatResult: a map keyed by metric name, value is the data point list
//   - error: nil if success otherwise the specific error
func (cli *Client) GetStatTimeByGroup(queryData *api.StatTimeByGroupQueryData) (api.StatResult, error) {
	return api.GetStatTimeByGroup(cli, queryData)
}

// GetStatSumByGroup - type 4: showType=sum, group required
// For details, please refer https://cloud.baidu.com/doc/GEO/s/4mp0unbx3
//
// PARAMS:
//   - queryData: the querying conditions, `Metric` and `Group` are required
//
// RETURNS:
//   - api.StatResult: a map keyed by metric name, value is the data point list
//   - error: nil if success otherwise the specific error
func (cli *Client) GetStatSumByGroup(queryData *api.StatSumByGroupQueryData) (api.StatResult, error) {
	return api.GetStatSumByGroup(cli, queryData)
}

// GetStatTopByGroup - type 5: showType=top, group required
// Note: only `sum_flow`, `upstream_flow`, `download_flow`, `pv` metrics are supported by the server for top queries.
// For details, please refer https://cloud.baidu.com/doc/GEO/s/4mp0unbx3
//
// PARAMS:
//   - queryData: the querying conditions, `Metric` and `Group` are required
//
// RETURNS:
//   - api.StatResult: a map keyed by metric name, value is the data point list
//   - error: nil if success otherwise the specific error
func (cli *Client) GetStatTopByGroup(queryData *api.StatTopByGroupQueryData) (api.StatResult, error) {
	return api.GetStatTopByGroup(cli, queryData)
}

// SetSiteConfig - set the site configurations
//
// PARAMS:
//   - site: the site to be configured
//   - config: the configurations to be set;
//
// RETURNS:
//   - *api.SiteConfigUpdateResult: the update status returned by the server
//   - error: nil if success otherwise the specific error
func (cli *Client) SetSiteConfig(site string, config *api.SiteConfig) (*api.SiteConfigUpdateResult, error) {
	return api.SetSiteConfig(cli, site, config)
}

// GetSiteConfig - get the site configurations
//
// PARAMS:
//   - site: the site to be queried
//
// RETURNS:
//   - *api.SiteConfig: the configurations of the site
//   - error: nil if success otherwise the specific error
func (cli *Client) GetSiteConfig(site string) (*api.SiteConfig, error) {
	return api.GetSiteConfig(cli, site)
}

// SetLimitConsumption - set the usage-capping configuration
//
// PARAMS:
//   - config: the limit-consumption configuration to be set
//
// RETURNS:
//   - *api.LimitConsumptionResult: whether the setting succeeded
//   - error: nil if success otherwise the specific error
func (cli *Client) SetLimitConsumption(config *api.LimitConsumption) (*api.LimitConsumptionResult, error) {
	return api.SetLimitConsumption(cli, config)
}

// GetLimitConsumption - query the usage-capping rules of a site
//
// PARAMS:
//   - query: the query conditions, `Site` is required
//
// RETURNS:
//   - *api.LimitConsumptionQueryResult: the matched rules and pagination info
//   - error: nil if success otherwise the specific error
func (cli *Client) GetLimitConsumption(query *api.LimitConsumptionQuery) (*api.LimitConsumptionQueryResult, error) {
	return api.GetLimitConsumption(cli, query)
}

// DeleteLimitConsumption - delete the usage-capping rule of a site or a domain
//
// PARAMS:
//   - target: the delete target, both `Site` and `Object` are required
//
// RETURNS:
//   - *api.LimitConsumptionResult: whether the deletion succeeded
//   - error: nil if success otherwise the specific error
func (cli *Client) DeleteLimitConsumption(target *api.LimitConsumptionDelete) (*api.LimitConsumptionResult, error) {
	return api.DeleteLimitConsumption(cli, target)
}

// CreateOriginPool - create an origin pool for a site
//
// PARAMS:
//   - site: the site name
//   - pool: the origin pool definition
//
// RETURNS:
//   - *api.CreateOriginPoolResult: whether the creation succeeded
//   - error: nil if success otherwise the specific error
func (cli *Client) CreateOriginPool(site string, pool *api.OriginPoolConfig) (*api.OriginPoolResult, error) {
	return api.CreateOriginPool(cli, site, pool)
}

// UpdateOriginPool - update an origin pool of a site
//
// PARAMS:
//   - site: the site name
//   - originPoolId: the origin pool id to be updated
//   - pool: the origin pool definition
//
// RETURNS:
//   - *api.OriginPoolResult: whether the update succeeded
//   - error: nil if success otherwise the specific error
func (cli *Client) UpdateOriginPool(site string, originPoolId string, pool *api.OriginPoolConfig) (*api.OriginPoolResult, error) {
	return api.UpdateOriginPool(cli, site, originPoolId, pool)
}

// DeleteOriginPool - delete an origin pool of a site
//
// PARAMS:
//   - site: the site name
//   - originPoolId: the origin pool id to be deleted
//
// RETURNS:
//   - *api.OriginPoolResult: whether the deletion succeeded
//   - error: nil if success otherwise the specific error
func (cli *Client) DeleteOriginPool(site string, originPoolId string) (*api.OriginPoolResult, error) {
	return api.DeleteOriginPool(cli, site, originPoolId)
}

// QueryOriginPool - query the origin pools of a site
//
// PARAMS:
//   - query: the query conditions, `Site` is required
//
// RETURNS:
//   - *api.QueryOriginPoolResult: the matched origin pools and pagination info
//   - error: nil if success otherwise the specific error
func (cli *Client) QueryOriginPool(query *api.QueryOriginPoolRequest) (*api.QueryOriginPoolResult, error) {
	return api.QueryOriginPool(cli, query)
}

// CreateLoadBalance - create a load balance for a site
//
// PARAMS:
//   - req: the load balance definition
//
// RETURNS:
//   - *api.LoadBalanceResult: whether the creation succeeded
//   - error: nil if success otherwise the specific error
func (cli *Client) CreateLoadBalance(req *api.LoadBalanceRequest) (*api.LoadBalanceResult, error) {
	return api.CreateLoadBalance(cli, req)
}

// UpdateLoadBalance - update a load balance of a site
//
// PARAMS:
//   - id: the load balance id to be updated
//   - req: the load balance definition
//
// RETURNS:
//   - *api.LoadBalanceResult: whether the update succeeded
//   - error: nil if success otherwise the specific error
func (cli *Client) UpdateLoadBalance(id string, req *api.LoadBalanceRequest) (*api.LoadBalanceResult, error) {
	return api.UpdateLoadBalance(cli, id, req)
}

// DeleteLoadBalance - delete a load balance of a site
//
// PARAMS:
//   - id: the load balance id to be deleted
//   - siteId: the site id that the load balance belongs to
//
// RETURNS:
//   - *api.LoadBalanceResult: whether the deletion succeeded
//   - error: nil if success otherwise the specific error
func (cli *Client) DeleteLoadBalance(id string, siteId int) (*api.LoadBalanceResult, error) {
	return api.DeleteLoadBalance(cli, id, siteId)
}

// QueryLoadBalance - query the load balances of a site
//
// PARAMS:
//   - query: the query conditions, `SiteId` is required
//
// RETURNS:
//   - *api.QueryLoadBalanceResult: the matched load balances and pagination info
//   - error: nil if success otherwise the specific error
func (cli *Client) QueryLoadBalance(query *api.QueryLoadBalanceRequest) (*api.QueryLoadBalanceResult, error) {
	return api.QueryLoadBalance(cli, query)
}

// CreateDomain - create an accelerated domain for a site
//
// PARAMS:
//   - req: the domain definition
//
// RETURNS:
//   - *api.CreateDomainResult: the creation status and generated cname
//   - error: nil if success otherwise the specific error
func (cli *Client) CreateDomain(req *api.CreateDomainRequest) (*api.CreateDomainResult, error) {
	return api.CreateDomain(cli, req)
}

// QueryDomainConfig - query the configuration of a domain
//
// PARAMS:
//   - domain: the accelerated domain to be queried
//
// RETURNS:
//   - *api.QueryDomainConfigResult: the domain and its full configuration
//   - error: nil if success otherwise the specific error
func (cli *Client) QueryDomainConfig(domain string) (*api.QueryDomainConfigResult, error) {
	return api.QueryDomainConfig(cli, domain)
}

// UpdateDomainConfig - update the configuration of a domain
//
// PARAMS:
//   - req: the update request, `Site` and `Config` are required
//
// RETURNS:
//   - *api.DomainOperationResult: the update result
//   - error: nil if success otherwise the specific error
func (cli *Client) UpdateDomainConfig(req *api.UpdateDomainConfigRequest) (*api.DomainOperationResult, error) {
	return api.UpdateDomainConfig(cli, req)
}

// QueryDomains - query the accelerated domain list and their simple configs
//
// PARAMS:
//   - req: the query conditions, all fields are optional; nil means query all
//
// RETURNS:
//   - *api.QueryDomainsResult: the domain list and pagination info
//   - error: nil if success otherwise the specific error
func (cli *Client) QueryDomains(req *api.QueryDomainsRequest) (*api.QueryDomainsResult, error) {
	return api.QueryDomains(cli, req)
}

// StartDomain - enable one or more accelerated domains
//
// PARAMS:
//   - req: the domains to be enabled, set either `Domain` or `Domains`
//
// RETURNS:
//   - *api.DomainOperationResult: the operation result
//   - error: nil if success otherwise the specific error
func (cli *Client) StartDomain(req *api.DomainBatchRequest) (*api.DomainOperationResult, error) {
	return api.StartDomain(cli, req)
}

// StopDomain - disable one or more accelerated domains
//
// PARAMS:
//   - req: the domains to be disabled, set either `Domain` or `Domains`
//
// RETURNS:
//   - *api.DomainOperationResult: the operation result
//   - error: nil if success otherwise the specific error
func (cli *Client) StopDomain(req *api.DomainBatchRequest) (*api.DomainOperationResult, error) {
	return api.StopDomain(cli, req)
}

// DeleteDomain - delete one or more accelerated domains
//
// PARAMS:
//   - req: the domains to be deleted, set either `Domain` or `Domains`
//
// RETURNS:
//   - *api.DomainOperationResult: the operation result
//   - error: nil if success otherwise the specific error
func (cli *Client) DeleteDomain(req *api.DomainBatchRequest) (*api.DomainOperationResult, error) {
	return api.DeleteDomain(cli, req)
}
