package ddcrds

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/baidubce/bce-sdk-go/auth"
	"github.com/baidubce/bce-sdk-go/bce"
	"github.com/baidubce/bce-sdk-go/services/rds"
)

const (
	DEFAULT_ENDPOINT                = "ddc.su.baidubce.com"
	DDC_NOT_SUPPORTED               = "DDC does not support this feature."
	RDS_NOT_SUPPORTED               = "RDS does not support this feature."
	URI_PREFIX                      = bce.URI_PREFIX + "v1/ddc"
	URI_PREFIX_V2                   = bce.URI_PREFIX + "v2/ddc"
	REQUEST_DDC_INSTANCE_URL        = "/instance"
	REQUEST_DDC_POOL_URL            = "/pool"
	REQUEST_DDC_HOST_URL            = "/host"
	REQUEST_DDC_TASK_URL            = "/task"
	REQUEST_DDC_DEPLOY_URL          = "/deploy"
	REQUEST_DDC_DATABASE_URL        = "/database"
	REQUEST_DDC_TABLE_URL           = "/table"
	REQUEST_DDC_HARDLINK_URL        = "/link"
	REQUEST_DDC_ACCOUNT_URL         = "/account"
	REQUEST_DDC_ROGROUP_URL         = "/roGroup"
	REQUEST_DDC_RECYCLER_URL        = "/recycler"
	REQUEST_DDC_SECURITYGROUP_URL   = "/security"
	REQUEST_DDC_LOG_URL             = "/logs"
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

func getDdcUriWithInstanceIdV2(instanceId string) string {
	return URI_PREFIX_V2 + REQUEST_DDC_INSTANCE_URL + "/" + instanceId
}

// Database URL
func getDatabaseUriWithInstanceId(instanceId string) string {
	return URI_PREFIX + REQUEST_DDC_INSTANCE_URL + "/" + instanceId + REQUEST_DDC_DATABASE_URL
}

func getDatabaseUriWithDbName(instanceId string, dbName string) string {
	return URI_PREFIX + REQUEST_DDC_INSTANCE_URL + "/" + instanceId + REQUEST_DDC_DATABASE_URL + "/" + dbName
}

func getQueryDatabaseUriWithDbName(instanceId string, dbName string) string {
	return URI_PREFIX + REQUEST_DDC_INSTANCE_URL + "/" + instanceId + REQUEST_DDC_DATABASE_URL + "/" + dbName + "/amount"
}

func getDatabaseDiskUsageUriWithInstanceId(instanceId string) string {
	return URI_PREFIX + REQUEST_DDC_INSTANCE_URL + "/" + instanceId + REQUEST_DDC_DATABASE_URL + "/usage"
}

func getDatabaseRecoverTimeUriWithInstanceId(instanceId string) string {
	return URI_PREFIX + REQUEST_DDC_INSTANCE_URL + "/" + instanceId + REQUEST_DDC_DATABASE_URL + "/recoverableDateTimes"
}

func getRecoverInstanceDatabaseUriWithInstanceId(instanceId string) string {
	return URI_PREFIX + REQUEST_DDC_INSTANCE_URL + "/" + instanceId + "/recoveryToSourceInstanceByDatetime"
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

// MaintainTime URL
func getMaintainTimeUriWithInstanceId(instanceId string) string {
	return URI_PREFIX + REQUEST_DDC_INSTANCE_URL + "/" + instanceId + REQUEST_DDC_MAINTAINTIME_URL
}

// MaintainTime URL
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

// Recycler URL
func getRecyclerUrl() string {
	return URI_PREFIX + REQUEST_DDC_RECYCLER_URL
}

// Recycler Recover URL
func getRecyclerRecoverUrl() string {
	return URI_PREFIX + REQUEST_DDC_RECYCLER_URL + "/batchRecover"
}

// Recycler Recover URL
func getRecyclerDeleteUrl() string {
	return URI_PREFIX + REQUEST_DDC_RECYCLER_URL + "/batchDelete"
}

// List Security Group By Vpc URL
func getSecurityGroupWithVpcIdUrl(vpcId string) string {
	return URI_PREFIX + REQUEST_DDC_SECURITYGROUP_URL + "/" + vpcId + "/listByVpc"
}

// List Security Group By Instance URL
func getSecurityGroupWithInstanceIdUrl(instanceId string) string {
	return URI_PREFIX + REQUEST_DDC_SECURITYGROUP_URL + "/" + instanceId + "/list"
}

// Bind Security Group To Instance URL
func getBindSecurityGroupWithUrl() string {
	return URI_PREFIX + REQUEST_DDC_SECURITYGROUP_URL + "/bind"
}

// UnBind Security Group To Instance URL
func getUnBindSecurityGroupWithUrl() string {
	return URI_PREFIX + REQUEST_DDC_SECURITYGROUP_URL + "/unbind"
}

// Batch Replace Security Group URL
func getReplaceSecurityGroupWithUrl() string {
	return URI_PREFIX + REQUEST_DDC_SECURITYGROUP_URL + "/updateSecurityGroup"
}

func getAccessLogUrl() string {
	return URI_PREFIX + REQUEST_DDC_LOG_URL + "/accessLog"
}

func getLogsUrlWithInstanceId(instanceId string) string {
	return URI_PREFIX + REQUEST_DDC_INSTANCE_URL + "/" + instanceId + REQUEST_DDC_LOG_URL
}

func getLogsUrlWithLogId(instanceId, logId string) string {
	return URI_PREFIX + REQUEST_DDC_INSTANCE_URL + "/" + instanceId + REQUEST_DDC_LOG_URL + "/" + logId
}

func getErrorLogsUrlWithInstanceId(instanceId string) string {
	return URI_PREFIX + REQUEST_DDC_INSTANCE_URL + "/" + instanceId + REQUEST_DDC_LOG_URL + "/logErrorDetail"
}

func getSlowLogsUrlWithInstanceId(instanceId string) string {
	return URI_PREFIX + REQUEST_DDC_INSTANCE_URL + "/" + instanceId + REQUEST_DDC_LOG_URL + "/logSlowDetail"
}

func getCreateTableHardLinkUrl(instanceId, dbName string) string {
	return URI_PREFIX + REQUEST_DDC_INSTANCE_URL + "/" + instanceId +
		REQUEST_DDC_DATABASE_URL + "/" + dbName +
		REQUEST_DDC_TABLE_URL + REQUEST_DDC_HARDLINK_URL
}

func getTableHardLinkUrl(instanceId, dbName, tableName string) string {
	return URI_PREFIX + REQUEST_DDC_INSTANCE_URL + "/" + instanceId +
		REQUEST_DDC_DATABASE_URL + "/" + dbName +
		REQUEST_DDC_TABLE_URL + "/" + tableName + REQUEST_DDC_HARDLINK_URL
}

func getChangeSemiSyncStatusUrlWithId(instanceId string) string {
	return URI_PREFIX + REQUEST_DDC_INSTANCE_URL + "/" + instanceId + "/semisync"
}

func getKillSessionUri(instanceId string) string {
	return URI_PREFIX + REQUEST_DDC_INSTANCE_URL + "/" + instanceId + "/session/kill"
}

func getKillSessionTaskUri(instanceId string, taskId int) string {
	return URI_PREFIX + REQUEST_DDC_INSTANCE_URL + "/" + instanceId + "/session/task" + "/" + strconv.Itoa(taskId)
}

func getMaintainTaskUri() string {
	return URI_PREFIX + REQUEST_DDC_TASK_URL
}

func getMaintainTaskDetailUri() string {
	return URI_PREFIX + REQUEST_DDC_TASK_URL + "/detail"
}

func getMaintainTaskUriWithTaskId(taskId string) string {
	return URI_PREFIX + REQUEST_DDC_TASK_URL + "/" + taskId
}

func getInstanceBackupStatusUrl(instanceId string) string {
	return URI_PREFIX + REQUEST_DDC_INSTANCE_URL + "/" + instanceId + "/backupStatus"
}

func getInstanceSyncDelayUrl(instanceId string) string {
	return URI_PREFIX + REQUEST_DDC_INSTANCE_URL + "/" + instanceId + "/sync_delay"
}

func getInstanceSyncDelayReplicationUrl(instanceId string) string {
	return URI_PREFIX + REQUEST_DDC_INSTANCE_URL + "/" + instanceId + "/replication"
}
