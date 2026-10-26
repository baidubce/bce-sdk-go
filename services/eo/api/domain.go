package api

import (
	"encoding/json"
	"errors"
	"strconv"

	"github.com/baidubce/bce-sdk-go/bce"
)

// OriginPoolConfig defined the origin pool definition of a site.
type OriginPoolConfig struct {
	Name         string           `json:"name"`
	OriginConfig []OriginPoolItem `json:"originConfig"`
}

// OriginPoolItem defined a single origin inside an origin pool.
type OriginPoolItem struct {
	Addr            string                `json:"addr"`
	Type            string                `json:"type"`
	Host            *string               `json:"host,omitempty"`
	Weight          int                   `json:"weight"`
	ThirdBucketAuth *OriginPoolBucketAuth `json:"thirdBucketAuth,omitempty"`
}

// OriginPoolBucketAuth defined the object storage authentication of an origin pool item.
type OriginPoolBucketAuth struct {
	AuthType string  `json:"authType"`
	Enabled  bool    `json:"enabled"`
	Ak       *string `json:"ak,omitempty"`
	Sk       *string `json:"sk,omitempty"`
	Bucket   *string `json:"bucket,omitempty"`
	Region   *string `json:"region,omitempty"`
	Service  *string `json:"service,omitempty"`
}

// CreateOriginPoolRequest defined the request of creating an origin pool.
type CreateOriginPoolRequest struct {
	Site       string            `json:"site"`
	OriginPool *OriginPoolConfig `json:"originPool"`
}

// OriginPoolResult defined the shared response of create/update/delete origin pool.
type OriginPoolResult struct {
	Success bool `json:"success"`
}

// CreateOriginPool - create an origin pool for a site
//
// PARAMS:
//   - cli: the client agent which can perform sending request
//   - site: the site name
//   - pool: the origin pool definition
//
// RETURNS:
//   - *OriginPoolResult: whether the creation succeeded
//   - error: nil if success otherwise the specific error
func CreateOriginPool(cli bce.Client, site string, pool *OriginPoolConfig) (*OriginPoolResult, error) {
	if site == "" {
		return nil, errors.New("site is required")
	}
	if pool == nil {
		return nil, errors.New("originPool is required")
	}
	if pool.Name == "" {
		return nil, errors.New("originPool.name is required")
	}
	if len(pool.OriginConfig) == 0 {
		return nil, errors.New("originPool.originConfig is required")
	}

	req := &CreateOriginPoolRequest{Site: site, OriginPool: pool}
	params := map[string]string{"action": "create"}
	respObj := &OriginPoolResult{}
	err := httpRequest(cli, "POST", "/v2/geo/origin-pool", params, req, respObj)
	if err != nil {
		return nil, err
	}

	return respObj, nil
}

// UpdateOriginPoolRequest defined the request of updating an origin pool.
type UpdateOriginPoolRequest struct {
	Site         string            `json:"site"`
	OriginPoolId string            `json:"originPoolId"`
	OriginPool   *OriginPoolConfig `json:"originPool"`
}

// UpdateOriginPool - update an origin pool of a site
//
// PARAMS:
//   - cli: the client agent which can perform sending request
//   - site: the site name
//   - originPoolId: the origin pool id to be updated
//   - pool: the origin pool definition
//
// RETURNS:
//   - *OriginPoolResult: whether the update succeeded
//   - error: nil if success otherwise the specific error
func UpdateOriginPool(cli bce.Client, site string, originPoolId string, pool *OriginPoolConfig) (*OriginPoolResult, error) {
	if site == "" {
		return nil, errors.New("site is required")
	}
	if originPoolId == "" {
		return nil, errors.New("originPoolId is required")
	}
	if pool == nil {
		return nil, errors.New("originPool is required")
	}
	if pool.Name == "" {
		return nil, errors.New("originPool.name is required")
	}
	if len(pool.OriginConfig) == 0 {
		return nil, errors.New("originPool.originConfig is required")
	}

	req := &UpdateOriginPoolRequest{Site: site, OriginPoolId: originPoolId, OriginPool: pool}
	params := map[string]string{"action": "update"}
	respObj := &OriginPoolResult{}
	err := httpRequest(cli, "POST", "/v2/geo/origin-pool", params, req, respObj)
	if err != nil {
		return nil, err
	}

	return respObj, nil
}

// QueryOriginPoolRequest defined the request of querying origin pools.
type QueryOriginPoolRequest struct {
	Site         string `json:"site"`
	Name         string `json:"name,omitempty"`
	OriginPoolId string `json:"originPoolId,omitempty"`
	PageNo       int    `json:"pageNo,omitempty"`
	PageSize     int    `json:"pageSize,omitempty"`
}

// QueryOriginPoolResult defined the response of querying origin pools.
type QueryOriginPoolResult struct {
	OriginPools []OriginPoolDetail `json:"originPools"`
	PageNo      int                `json:"pageNo"`
	PageSize    int                `json:"pageSize"`
	Total       int                `json:"total"`
}

// OriginPoolDetail defined a single origin pool returned by the query.
type OriginPoolDetail struct {
	Id           int                    `json:"id"`
	Uid          string                 `json:"uid"`
	Site         string                 `json:"site"`
	Name         string                 `json:"name"`
	OriginPoolId string                 `json:"originPoolId"`
	Status       string                 `json:"status"`
	Count        int                    `json:"count"`
	OriginConfig []OriginPoolDetailItem `json:"originConfig"`
	Instance     []OriginPoolInstance   `json:"instance"`
	Slb          []interface{}          `json:"slb"`
}

// OriginPoolDetailItem defined a single origin inside a queried origin pool.
type OriginPoolDetailItem struct {
	Addr            string                `json:"addr"`
	Type            string                `json:"type"`
	Host            *string               `json:"host,omitempty"`
	Weight          int                   `json:"weight"`
	Backup          bool                  `json:"backup"`
	ThirdBucketAuth *OriginPoolBucketAuth `json:"thirdBucketAuth,omitempty"`
}

// OriginPoolInstance defined a reference instance of an origin pool.
type OriginPoolInstance struct {
	Domain          string `json:"domain,omitempty"`
	Name            string `json:"name,omitempty"`
	LoadBalanceName string `json:"loadBalanceName,omitempty"`
	Type            string `json:"type"`
	Site            string `json:"site"`
}

// QueryOriginPool - query the origin pools of a site
//
// PARAMS:
//   - cli: the client agent which can perform sending request
//   - query: the query conditions, `Site` is required
//
// RETURNS:
//   - *QueryOriginPoolResult: the matched origin pools and pagination info
//   - error: nil if success otherwise the specific error
func QueryOriginPool(cli bce.Client, query *QueryOriginPoolRequest) (*QueryOriginPoolResult, error) {
	if query == nil {
		return nil, errors.New("query is required")
	}
	if query.Site == "" {
		return nil, errors.New("site is required")
	}

	params := map[string]string{"action": "query"}
	respObj := &QueryOriginPoolResult{}
	err := httpRequest(cli, "POST", "/v2/geo/origin-pool", params, query, respObj)
	if err != nil {
		return nil, err
	}

	return respObj, nil
}

// DeleteOriginPoolRequest defined the request of deleting an origin pool.
type DeleteOriginPoolRequest struct {
	Site         string `json:"site"`
	OriginPoolId string `json:"originPoolId"`
}

// DeleteOriginPool - delete an origin pool of a site
//
// PARAMS:
//   - cli: the client agent which can perform sending request
//   - site: the site name
//   - originPoolId: the origin pool id to be deleted
//
// RETURNS:
//   - *OriginPoolResult: whether the deletion succeeded
//   - error: nil if success otherwise the specific error
func DeleteOriginPool(cli bce.Client, site string, originPoolId string) (*OriginPoolResult, error) {
	if site == "" {
		return nil, errors.New("site is required")
	}
	if originPoolId == "" {
		return nil, errors.New("originPoolId is required")
	}

	req := &DeleteOriginPoolRequest{Site: site, OriginPoolId: originPoolId}
	params := map[string]string{"action": "delete"}
	respObj := &OriginPoolResult{}
	err := httpRequest(cli, "POST", "/v2/geo/origin-pool", params, req, respObj)
	if err != nil {
		return nil, err
	}

	return respObj, nil
}

// DetectConfig defined the health check configuration of a load balance.
type DetectConfig struct {
	HealthCheckType   string `json:"healthCheckType,omitempty"`
	ProbeInterval     int    `json:"probeInterval"`
	ProbeMinSuccesses int    `json:"probeMinSuccesses"`
	ProbeMaxFails     int    `json:"probeMaxFails"`
	ProbeTimeout      int    `json:"probeTimeout"`
	ProbeSucceedCode  string `json:"probeSucceedCode,omitempty"`
	ProbeUrl          string `json:"probeUrl,omitempty"`
	ProbeMethod       string `json:"probeMethod,omitempty"`
	CheckData         string `json:"checkData,omitempty"`
	CheckResult       string `json:"checkResult,omitempty"`
}

// LoadBalanceRequest defined the request of creating or updating a load balance.
type LoadBalanceRequest struct {
	SiteId               int           `json:"siteId"`
	Name                 string        `json:"name"`
	DispatchMode         string        `json:"dispatchMode"`
	SessionMode          string        `json:"sessionMode"`
	DetectConfig         *DetectConfig `json:"detectConfig"`
	OriginPoolIds        []string      `json:"originPoolIds"`
	Weights              []int         `json:"weights,omitempty"`
	FallbackOriginPoolId string        `json:"fallbackOriginPoolId,omitempty"`
}

// LoadBalanceResult defined the shared response of load balance create/update/delete.
type LoadBalanceResult struct {
	Success bool `json:"success"`
}

// CreateLoadBalance - create a load balance for a site
//
// PARAMS:
//   - cli: the client agent which can perform sending request
//   - req: the load balance definition
//
// RETURNS:
//   - *LoadBalanceResult: whether the creation succeeded
//   - error: nil if success otherwise the specific error
func CreateLoadBalance(cli bce.Client, req *LoadBalanceRequest) (*LoadBalanceResult, error) {
	if req == nil {
		return nil, errors.New("request is required")
	}
	if req.Name == "" {
		return nil, errors.New("name is required")
	}
	if req.DispatchMode == "" {
		return nil, errors.New("dispatchMode is required")
	}
	if req.SessionMode == "" {
		return nil, errors.New("sessionMode is required")
	}
	if req.DetectConfig == nil {
		return nil, errors.New("detectConfig is required")
	}
	if len(req.OriginPoolIds) == 0 {
		return nil, errors.New("originPoolIds is required")
	}

	respObj := &LoadBalanceResult{}
	err := httpRequest(cli, "POST", "/v2/geo/load-balance", nil, req, respObj)
	if err != nil {
		return nil, err
	}

	return respObj, nil
}

// UpdateLoadBalance - update a load balance of a site
//
// PARAMS:
//   - cli: the client agent which can perform sending request
//   - id: the load balance id to be updated
//   - req: the load balance definition
//
// RETURNS:
//   - *LoadBalanceResult: whether the update succeeded
//   - error: nil if success otherwise the specific error
func UpdateLoadBalance(cli bce.Client, id string, req *LoadBalanceRequest) (*LoadBalanceResult, error) {
	if id == "" {
		return nil, errors.New("id is required")
	}
	if req == nil {
		return nil, errors.New("request is required")
	}
	if req.Name == "" {
		return nil, errors.New("name is required")
	}
	if req.DispatchMode == "" {
		return nil, errors.New("dispatchMode is required")
	}
	if req.SessionMode == "" {
		return nil, errors.New("sessionMode is required")
	}
	if req.DetectConfig == nil {
		return nil, errors.New("detectConfig is required")
	}
	if len(req.OriginPoolIds) == 0 {
		return nil, errors.New("originPoolIds is required")
	}

	respObj := &LoadBalanceResult{}
	err := httpRequest(cli, "PUT", "/v2/geo/load-balance/"+id, nil, req, respObj)
	if err != nil {
		return nil, err
	}

	return respObj, nil
}

// DeleteLoadBalanceRequest defined the request of deleting a load balance.
type DeleteLoadBalanceRequest struct {
	SiteId int `json:"siteId"`
}

// DeleteLoadBalance - delete a load balance of a site
//
// PARAMS:
//   - cli: the client agent which can perform sending request
//   - id: the load balance id to be deleted
//   - siteId: the site id that the load balance belongs to
//
// RETURNS:
//   - *LoadBalanceResult: whether the deletion succeeded
//   - error: nil if success otherwise the specific error
func DeleteLoadBalance(cli bce.Client, id string, siteId int) (*LoadBalanceResult, error) {
	if id == "" {
		return nil, errors.New("id is required")
	}

	req := &DeleteLoadBalanceRequest{SiteId: siteId}
	respObj := &LoadBalanceResult{}
	err := httpRequest(cli, "DELETE", "/v2/geo/load-balance/"+id, nil, req, respObj)
	if err != nil {
		return nil, err
	}

	return respObj, nil
}

// QueryLoadBalanceRequest defined the query conditions of querying load balances.
type QueryLoadBalanceRequest struct {
	SiteId   int    `json:"siteId"`
	Name     string `json:"name,omitempty"`
	Id       string `json:"id,omitempty"`
	FetchAll bool   `json:"fetchAll,omitempty"`
	PageNo   int    `json:"pageNo,omitempty"`
	PageSize int    `json:"pageSize,omitempty"`
}

// QueryLoadBalanceResult defined the response of querying load balances.
type QueryLoadBalanceResult struct {
	Total             int                 `json:"total"`
	LoadBalanceDetail []LoadBalanceDetail `json:"loadBalanceDetail"`
	PageNo            int                 `json:"pageNo,omitempty"`
	PageSize          int                 `json:"pageSize,omitempty"`
}

// LoadBalanceDetail defined a single load balance returned by the query.
type LoadBalanceDetail struct {
	LoadBalanceId        string                `json:"loadBalanceId"`
	Name                 string                `json:"name"`
	Status               string                `json:"status"`
	DispatchMode         string                `json:"dispatchMode"`
	SessionMode          string                `json:"sessionMode"`
	DetectConfig         *DetectConfig         `json:"detectConfig,omitempty"`
	OriginPoolIds        []string              `json:"originPoolIds"`
	WeightList           []int                 `json:"weightList"`
	FallbackOriginPoolId *string               `json:"fallbackOriginPoolId,omitempty"`
	OriginPoolCount      int                   `json:"originPoolCount"`
	Instance             []LoadBalanceInstance `json:"instance"`
}

// LoadBalanceInstance defined a reference instance of a load balance.
type LoadBalanceInstance struct {
	Domain string `json:"domain,omitempty"`
	Type   string `json:"type,omitempty"`
	Site   string `json:"site,omitempty"`
}

// QueryLoadBalance - query the load balances of a site
//
// PARAMS:
//   - cli: the client agent which can perform sending request
//   - query: the query conditions, `SiteId` is required
//
// RETURNS:
//   - *QueryLoadBalanceResult: the matched load balances and pagination info
//   - error: nil if success otherwise the specific error
func QueryLoadBalance(cli bce.Client, query *QueryLoadBalanceRequest) (*QueryLoadBalanceResult, error) {
	if query == nil {
		return nil, errors.New("query is required")
	}
	if query.SiteId == 0 {
		return nil, errors.New("siteId is required")
	}

	params := map[string]string{"siteId": strconv.Itoa(query.SiteId)}
	if query.Name != "" {
		params["name"] = query.Name
	}
	if query.Id != "" {
		params["id"] = query.Id
	}
	if query.FetchAll {
		params["fetchAll"] = "true"
	}
	if query.PageNo > 0 {
		params["pageNo"] = strconv.Itoa(query.PageNo)
	}
	if query.PageSize > 0 {
		params["pageSize"] = strconv.Itoa(query.PageSize)
	}

	respObj := &QueryLoadBalanceResult{}
	err := httpRequest(cli, "GET", "/v2/geo/load-balance", params, nil, respObj)
	if err != nil {
		return nil, err
	}

	return respObj, nil
}

// DomainOriginConfig defined a single origin of a domain.
type DomainOriginConfig struct {
	Addr             string                `json:"addr"`
	Type             string                `json:"type"`
	Host             *string               `json:"host,omitempty"`
	UpstreamProtocol string                `json:"upstreamProtocol"`
	HttpPort         *int                  `json:"httpPort,omitempty"`
	HttpsPort        *int                  `json:"httpsPort,omitempty"`
	Backup           *bool                 `json:"backup,omitempty"`
	Weight           *int                  `json:"weight,omitempty"`
	ThirdBucketAuth  *OriginPoolBucketAuth `json:"thirdBucketAuth,omitempty"`
}

// Ipv6Dispatch defined the ipv6 dispatch configuration of a domain.
type Ipv6Dispatch struct {
	Enable string `json:"enable"`
}

// CreateDomainRequest defined the request of creating a domain.
type CreateDomainRequest struct {
	Domain           string               `json:"domain"`
	SiteId           int                  `json:"siteId"`
	OriginConfig     []DomainOriginConfig `json:"originConfig,omitempty"`
	OriginPoolId     string               `json:"originPoolId,omitempty"`
	UpstreamProtocol string               `json:"upstreamProtocol,omitempty"`
	Ipv6Dispatch     *Ipv6Dispatch        `json:"ipv6Dispatch,omitempty"`
	HttpPort         *int                 `json:"httpPort,omitempty"`
	HttpsPort        *int                 `json:"httpsPort,omitempty"`
	LoadBalanceId    string               `json:"loadBalanceId,omitempty"`
}

// CreateDomainResult defined the response of creating a domain.
type CreateDomainResult struct {
	Status string `json:"status"`
	Cname  string `json:"cname"`
}

// CreateDomain - create an accelerated domain for a site
//
// PARAMS:
//   - cli: the client agent which can perform sending request
//   - req: the domain definition
//
// RETURNS:
//   - *CreateDomainResult: the creation status and generated cname
//   - error: nil if success otherwise the specific error
func CreateDomain(cli bce.Client, req *CreateDomainRequest) (*CreateDomainResult, error) {
	if req == nil {
		return nil, errors.New("request is required")
	}
	if req.Domain == "" {
		return nil, errors.New("domain is required")
	}

	params := map[string]string{"action": "create"}
	respObj := &CreateDomainResult{}
	err := httpRequest(cli, "POST", "/v2/geo/domain", params, req, respObj)
	if err != nil {
		return nil, err
	}

	return respObj, nil
}

// QueryDomainConfigRequest defined the request of querying a domain config.
type QueryDomainConfigRequest struct {
	Domain string `json:"domain"`
}

// QueryDomainConfigResult defined the response of querying a domain config.
// Config keeps the raw JSON of the domain configuration so that every field the
// server returns is preserved without modeling each one.
type QueryDomainConfigResult struct {
	Domain string          `json:"domain"`
	Config json.RawMessage `json:"config"`
}

// QueryDomainConfig - query the configuration of a domain
//
// PARAMS:
//   - cli: the client agent which can perform sending request
//   - domain: the accelerated domain to be queried
//
// RETURNS:
//   - *QueryDomainConfigResult: the domain and its full configuration
//   - error: nil if success otherwise the specific error
func QueryDomainConfig(cli bce.Client, domain string) (*QueryDomainConfigResult, error) {
	if domain == "" {
		return nil, errors.New("domain is required")
	}

	req := &QueryDomainConfigRequest{Domain: domain}
	params := map[string]string{"action": "query"}
	respObj := &QueryDomainConfigResult{}
	err := httpRequest(cli, "POST", "/v2/geo/domain/config", params, req, respObj)
	if err != nil {
		return nil, err
	}

	return respObj, nil
}

// DomainConfigUpdate defined the config to be updated of a domain.
type DomainConfigUpdate struct {
	OriginConfig     []DomainOriginConfig `json:"originConfig,omitempty"`
	OriginPoolId     string               `json:"originPoolId,omitempty"`
	UpstreamProtocol string               `json:"upstreamProtocol,omitempty"`
	Ipv6Dispatch     *Ipv6Dispatch        `json:"ipv6Dispatch,omitempty"`
	HttpPort         *int                 `json:"httpPort,omitempty"`
	HttpsPort        *int                 `json:"httpsPort,omitempty"`
	LoadBalanceId    string               `json:"loadBalanceId,omitempty"`
	CertId           string               `json:"certId,omitempty"`
	HttpsEnabled     string               `json:"httpsEnabled,omitempty"`
}

// UpdateDomainConfigRequest defined the request of updating a domain config.
type UpdateDomainConfigRequest struct {
	Site    string              `json:"site"`
	Domain  string              `json:"domain,omitempty"`
	Domains []string            `json:"domains,omitempty"`
	Config  *DomainConfigUpdate `json:"config"`
}

// DomainOperationResult defined the shared response of the batch domain operations
// (update config / start / stop / delete).
type DomainOperationResult struct {
	Status  string `json:"status"`
	Total   int    `json:"total"`
	Success int    `json:"success"`
	ErrMsg  string `json:"errMsg,omitempty"`
}

// UpdateDomainConfig - update the configuration of a domain
//
// PARAMS:
//   - cli: the client agent which can perform sending request
//   - req: the update request, `Site` and `Config` are required
//
// RETURNS:
//   - *DomainOperationResult: the update result
//   - error: nil if success otherwise the specific error
func UpdateDomainConfig(cli bce.Client, req *UpdateDomainConfigRequest) (*DomainOperationResult, error) {
	if req == nil {
		return nil, errors.New("request is required")
	}
	if req.Site == "" {
		return nil, errors.New("site is required")
	}
	if req.Config == nil {
		return nil, errors.New("config is required")
	}

	params := map[string]string{"action": "update"}
	respObj := &DomainOperationResult{}
	err := httpRequest(cli, "POST", "/v2/geo/domain/config", params, req, respObj)
	if err != nil {
		return nil, err
	}

	return respObj, nil
}

// QueryDomainsRequest defined the request of querying the domain list.
type QueryDomainsRequest struct {
	SiteId   int    `json:"siteId,omitempty"`
	Keyword  string `json:"keyword,omitempty"`
	Status   string `json:"status,omitempty"`
	PageNo   int    `json:"pageNo,omitempty"`
	PageSize int    `json:"pageSize,omitempty"`
}

// DomainSummary defined a single domain entry in the query result.
type DomainSummary struct {
	Domain            string               `json:"domain"`
	SiteId            int                  `json:"siteId"`
	Cname             string               `json:"cname"`
	CnameStatus       *int                 `json:"cnameStatus,omitempty"`
	Status            string               `json:"status"`
	BanStatus         string               `json:"banStatus,omitempty"`
	HttpsEnabled      string               `json:"httpsEnabled"`
	OriginConfig      []DomainOriginConfig `json:"originConfig"`
	OriginInformation string               `json:"originInformation"`
	OriginType        string               `json:"originType"`
	LoadBalanceId     string               `json:"loadBalanceId,omitempty"`
	OriginPoolId      string               `json:"originPoolId,omitempty"`
	OriginPools       *DomainOriginPools   `json:"originPools,omitempty"`
	Ipv6Dispatch      *Ipv6Dispatch        `json:"ipv6Dispatch,omitempty"`
}

// DomainOriginPoolItem defined a single origin pool entry inside DomainOriginPools.Origins.
// OriginPoolWeight is only present when the load balance dispatch mode is weight_mode.
type DomainOriginPoolItem struct {
	OriginPoolWeight *int                 `json:"originPoolWeight,omitempty"`
	OriginConfig     []DomainOriginConfig `json:"originConfig"`
}

// DomainOriginPools defined the origin pools config when a domain uses load balance.
type DomainOriginPools struct {
	Origins              []DomainOriginPoolItem `json:"origins"`
	DispatchMode         string                 `json:"dispatchMode"`
	SessionMode          string                 `json:"sessionMode"`
	DetectConfig         *DetectConfig          `json:"detectConfig,omitempty"`
	FallbackOriginConfig []DomainOriginConfig   `json:"fallbackOriginConfig"`
}

// QueryDomainsResult defined the response of querying the domain list.
type QueryDomainsResult struct {
	SiteDomainCount int             `json:"siteDomainCount,omitempty"`
	PlanDomainQuota int             `json:"planDomainQuota,omitempty"`
	PlanDomainCount int             `json:"planDomainCount,omitempty"`
	Domains         []DomainSummary `json:"domains"`
	Total           int             `json:"total"`
	PageNo          int             `json:"pageNo"`
	PageSize        int             `json:"pageSize"`
}

// QueryDomains - query the accelerated domain list and their simple configs
//
// PARAMS:
//   - cli: the client agent which can perform sending request
//   - req: the query conditions, all fields are optional; nil means query all
//
// RETURNS:
//   - *QueryDomainsResult: the domain list and pagination info
//   - error: nil if success otherwise the specific error
func QueryDomains(cli bce.Client, req *QueryDomainsRequest) (*QueryDomainsResult, error) {
	if req == nil {
		req = &QueryDomainsRequest{}
	}

	params := map[string]string{"action": "queryDomainsWithConfig"}
	respObj := &QueryDomainsResult{}
	err := httpRequest(cli, "POST", "/v2/geo/domain", params, req, respObj)
	if err != nil {
		return nil, err
	}

	return respObj, nil
}

// DomainBatchRequest defined the request of the batch domain operations (start / stop / delete).
// Domain takes precedence over Domains when both are set.
type DomainBatchRequest struct {
	Domain  string   `json:"domain,omitempty"`
	Domains []string `json:"domains,omitempty"`
}

// StartDomain - enable one or more accelerated domains
//
// PARAMS:
//   - cli: the client agent which can perform sending request
//   - req: the domains to be enabled, set either `Domain` or `Domains`
//
// RETURNS:
//   - *DomainOperationResult: the operation result
//   - error: nil if success otherwise the specific error
func StartDomain(cli bce.Client, req *DomainBatchRequest) (*DomainOperationResult, error) {
	if req == nil {
		return nil, errors.New("request is required")
	}
	if req.Domain == "" && len(req.Domains) == 0 {
		return nil, errors.New("domain or domains is required")
	}

	params := map[string]string{"action": "start"}
	respObj := &DomainOperationResult{}
	err := httpRequest(cli, "POST", "/v2/geo/domain", params, req, respObj)
	if err != nil {
		return nil, err
	}

	return respObj, nil
}

// StopDomain - disable one or more accelerated domains
//
// PARAMS:
//   - cli: the client agent which can perform sending request
//   - req: the domains to be disabled, set either `Domain` or `Domains`
//
// RETURNS:
//   - *DomainOperationResult: the operation result
//   - error: nil if success otherwise the specific error
func StopDomain(cli bce.Client, req *DomainBatchRequest) (*DomainOperationResult, error) {
	if req == nil {
		return nil, errors.New("request is required")
	}
	if req.Domain == "" && len(req.Domains) == 0 {
		return nil, errors.New("domain or domains is required")
	}

	params := map[string]string{"action": "stop"}
	respObj := &DomainOperationResult{}
	err := httpRequest(cli, "POST", "/v2/geo/domain", params, req, respObj)
	if err != nil {
		return nil, err
	}

	return respObj, nil
}

// DeleteDomain - delete one or more accelerated domains
//
// PARAMS:
//   - cli: the client agent which can perform sending request
//   - req: the domains to be deleted, set either `Domain` or `Domains`
//
// RETURNS:
//   - *DomainOperationResult: the operation result
//   - error: nil if success otherwise the specific error
func DeleteDomain(cli bce.Client, req *DomainBatchRequest) (*DomainOperationResult, error) {
	if req == nil {
		return nil, errors.New("request is required")
	}
	if req.Domain == "" && len(req.Domains) == 0 {
		return nil, errors.New("domain or domains is required")
	}

	params := map[string]string{"action": "delete"}
	respObj := &DomainOperationResult{}
	err := httpRequest(cli, "POST", "/v2/geo/domain", params, req, respObj)
	if err != nil {
		return nil, err
	}

	return respObj, nil
}
