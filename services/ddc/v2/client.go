package ddcrds

import (
	"fmt"
	"github.com/baidubce/bce-sdk-go/auth"
	"github.com/baidubce/bce-sdk-go/bce"
	"github.com/baidubce/bce-sdk-go/services/rds"
	"strings"
)

const (
	DEFAULT_ENDPOINT                = "ddc.su.baidubce.com"
	DDC_NOT_SUPPORTED               = "DDC does not support this feature."
	RDS_NOT_SUPPORTED               = "RDS does not support this feature."
	URI_PREFIX                      = bce.URI_PREFIX + "v1/ddc"
	REQUEST_DDC_INSTANCE_URL        = "/instance"
	REQUEST_DDC_POOL_URL            = "/pool"
	REQUEST_DDC_HOST_URL            = "/host"
	REQUEST_DDC_DEPLOY_URL          = "/deploy"
	REQUEST_DDC_DATABASE_URL        = "/database"
	REQUEST_DDC_ACCOUNT_URL         = "/account"
	REQUEST_DDC_ROGROUP_URL         = "/roGroup"
	REQUEST_DDC_UPDATE_ACTION       = "/update"
	REQUEST_DDC_MAINTAINTIME_URL    = "/maintenTimeInfo"
	REQUEST_UPDATE_MAINTAINTIME_URL = "/updateMaintenTime"
)

func DDCNotSupportError() error {
	return fmt.Errorf(DDC_NOT_SUPPORTED)
}
func RDSNotSupportError() error {
	return fmt.Errorf(RDS_NOT_SUPPORTED)
}

// Client of DDC service is a kind of BceClient, so derived from BceClient
type Client struct {
	rdsClient *rds.Client
	ddcClient *DDCClient
}

func (c *Client) Config(config *bce.BceClientConfiguration) {
	c.ddcClient.Config = config
	c.rdsClient.Config = config
}

func (c *Client) ConfigEndpoint(endPoint string) {
	// 替换Endpoint,优先创建ddc Client
	ddcEndpoint := strings.Replace(endPoint, "rds.", "ddc.", 1)
	c.ddcClient.Config.Endpoint = ddcEndpoint
	// 替换Endpoint
	rdsEndpoint := strings.Replace(endPoint, "ddc.", "rds.", 1)
	c.rdsClient.Config.Endpoint = rdsEndpoint
}

func (c *Client) ConfigRegion(region string) {
	c.ddcClient.Config.Region = region
	c.rdsClient.Config.Region = region
}

func (c *Client) ConfigRetry(policy bce.RetryPolicy) {
	c.ddcClient.Config.Retry = policy
	c.rdsClient.Config.Retry = policy
}

func (c *Client) ConfigSignOption(option *auth.SignOptions) {
	c.ddcClient.Config.SignOption = option
	c.rdsClient.Config.SignOption = option
}

func (c *Client) ConfigSignOptionHeadersToSign(header map[string]struct{}) {
	c.ddcClient.Config.SignOption.HeadersToSign = header
	c.rdsClient.Config.SignOption.HeadersToSign = header
}

func (c *Client) ConfigSignOptionExpireSeconds(seconds int) {
	c.ddcClient.Config.SignOption.ExpireSeconds = seconds
	c.rdsClient.Config.SignOption.ExpireSeconds = seconds
}

func (c *Client) ConfigCredentials(credentials *auth.BceCredentials) {
	c.ddcClient.Config.Credentials = credentials
	c.rdsClient.Config.Credentials = credentials
}

func (c *Client) ConfigProxyUrl(proxyUrl string) {
	c.ddcClient.Config.ProxyUrl = proxyUrl
	c.rdsClient.Config.ProxyUrl = proxyUrl
}

func (c *Client) ConfigConnectionTimeoutInMillis(millis int) {
	c.ddcClient.Config.ConnectionTimeoutInMillis = millis
	c.rdsClient.Config.ConnectionTimeoutInMillis = millis
}

// 内部创建rds和ddc两个client
func NewClient(ak, sk, endPoint string) (*Client, error) {
	if len(endPoint) == 0 {
		endPoint = DEFAULT_ENDPOINT
	}
	// 替换Endpoint,优先创建ddc Client
	ddcEndpoint := strings.Replace(endPoint, "rds.", "ddc.", 1)
	ddcClient, err := NewDDCClient(ak, sk, ddcEndpoint)
	if err != nil {
		return nil, err
	}
	// 替换Endpoint
	rdsEndpoint := strings.Replace(endPoint, "ddc.", "rds.", 1)
	rdsClient, err := rds.NewClient(ak, sk, rdsEndpoint)
	if err != nil {
		return nil, err
	}
	return &Client{rdsClient: rdsClient, ddcClient: ddcClient}, nil
}

// Client for DDC service
type DDCClient struct {
	*bce.BceClient
}

func NewDDCClient(ak, sk, endPoint string) (*DDCClient, error) {
	if len(endPoint) == 0 {
		endPoint = DEFAULT_ENDPOINT
	}
	client, err := bce.NewBceClientWithAkSk(ak, sk, endPoint)
	if err != nil {
		return nil, err
	}
	return &DDCClient{client}, nil
}

func getDdcUri() string {
	return URI_PREFIX
}

func getDdcInstanceUri() string {
	return URI_PREFIX + REQUEST_DDC_INSTANCE_URL
}

// Pool URL
func getPoolUri() string {
	return URI_PREFIX + REQUEST_DDC_POOL_URL
}

func getPoolUriWithId(poolId string) string {
	return URI_PREFIX + REQUEST_DDC_POOL_URL + "/" + poolId
}

// Host URL
func getHostUri() string {
	return URI_PREFIX + REQUEST_DDC_HOST_URL
}

func getHostUriWithPnetIp(poolId, hostPnetIP string) string {
	return URI_PREFIX + REQUEST_DDC_POOL_URL + "/" + poolId + REQUEST_DDC_HOST_URL + "/" + hostPnetIP
}

// DeploySet URL
func getDeploySetUri(poolId string) string {
	return URI_PREFIX + REQUEST_DDC_POOL_URL + "/" + poolId + REQUEST_DDC_DEPLOY_URL
}

func getDeploySetUriWithId(poolId, id string) string {
	return URI_PREFIX + REQUEST_DDC_POOL_URL + "/" + poolId + REQUEST_DDC_DEPLOY_URL + "/" + id
}

func getDdcUriWithInstanceId(instanceId string) string {
	return URI_PREFIX + REQUEST_DDC_INSTANCE_URL + "/" + instanceId
}

// Database URL
func getDatabaseUriWithInstanceId(instanceId string) string {
	return URI_PREFIX + REQUEST_DDC_INSTANCE_URL + "/" + instanceId + REQUEST_DDC_DATABASE_URL
}

func getDatabaseUriWithDbName(instanceId string, dbName string) string {
	return URI_PREFIX + REQUEST_DDC_INSTANCE_URL + "/" + instanceId + REQUEST_DDC_DATABASE_URL + "/" + dbName
}

// Account URL
func getAccountUriWithInstanceId(instanceId string) string {
	return URI_PREFIX + REQUEST_DDC_INSTANCE_URL + "/" + instanceId + REQUEST_DDC_ACCOUNT_URL
}

func getAccountUriWithAccountName(instanceId string, accountName string) string {
	return URI_PREFIX + REQUEST_DDC_INSTANCE_URL + "/" + instanceId + REQUEST_DDC_ACCOUNT_URL + "/" + accountName
}

// RoGroup URL
func getRoGroupUriWithInstanceId(instanceId string) string {
	return URI_PREFIX + REQUEST_DDC_INSTANCE_URL + "/" + instanceId + REQUEST_DDC_ROGROUP_URL
}

// MaintenTime URL
func getMaintainTimeUriWithInstanceId(instanceId string) string {
	return URI_PREFIX + REQUEST_DDC_INSTANCE_URL + "/" + instanceId + REQUEST_DDC_MAINTAINTIME_URL
}

// MaintenTime URL
func getUpdateMaintainTimeUriWithInstanceId(instanceId string) string {
	return URI_PREFIX + REQUEST_DDC_INSTANCE_URL + "/" + instanceId + REQUEST_UPDATE_MAINTAINTIME_URL
}

// RoGroup URL
func getUpdateRoGroupUriWithId(roGroupId string) string {
	return URI_PREFIX + REQUEST_DDC_INSTANCE_URL + REQUEST_DDC_ROGROUP_URL + "/" + roGroupId + REQUEST_DDC_UPDATE_ACTION
}

// RoGroupWeight URL
func getUpdateRoGroupWeightUriWithId(roGroupId string) string {
	return URI_PREFIX + REQUEST_DDC_INSTANCE_URL + REQUEST_DDC_ROGROUP_URL + "/" + roGroupId + "/updateWeight"
}

// ReBalance RoGroup URL
func getReBalanceRoGroupUriWithId(roGroupId string) string {
	return URI_PREFIX + REQUEST_DDC_INSTANCE_URL + REQUEST_DDC_ROGROUP_URL + "/" + roGroupId + "/balanceRoLoad"
}
