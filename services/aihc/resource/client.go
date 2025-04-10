package resource

import (
	"fmt"
	"strconv"

	"github.com/baidubce/bce-sdk-go/bce"
	"github.com/baidubce/bce-sdk-go/http"
	v1 "github.com/baidubce/bce-sdk-go/services/aihc/api/v1"
	"github.com/baidubce/bce-sdk-go/services/aihc/client"
)

const (
	URI_PREFIX                = bce.URI_PREFIX + "api" + bce.URI_PREFIX + "v1"
	REQUEST_RESOURCE_POOL_URL = "/resourcepools"
	REQUEST_NODE_URL          = "/nodes"
	REQUEST_QUEUE_URL         = "/queue"
	REQUEST_JOB_URL           = "/aijobs"

	BHCMP_CLUSTER      = "aihc-bhcmp"
	SERVERLESS_CLUSTER = "aihc-serverless"
	BosPrefixPath      = "aihc/tempfile/"
)

type Client struct {
	client.Client
}

// NewClient make the aihc inference service client with default configuration.
func NewClient(ak, sk, endPoint string) (*Client, error) {
	aihcClient, err := client.NewClient(ak, sk, endPoint)
	if err != nil {
		return nil, err
	}
	newClient := Client{*aihcClient}
	return &newClient, nil
}

func (c *Client) GetBceClient() *bce.BceClient {
	return c.DefaultClient
}

func (c *Client) SetBceClient(client *bce.BceClient) {
	c.DefaultClient = client
}

// NewClientWithSTS make the aihc inference service client with STS configuration.
func NewClientWithSTS(accessKey, secretKey, sessionToken, endPoint string) (*Client, error) {
	aihcClient, err := client.NewClientWithSTS(accessKey, secretKey, sessionToken, endPoint)
	if err != nil {
		return nil, err
	}

	newClient := Client{
		*aihcClient,
	}
	return &newClient, nil
}

func (c *Client) GetResourcePool(resourcePoolID string) (result *v1.GetResourcePoolResponse, err error) {
	if resourcePoolID == "" {
		return nil, fmt.Errorf("resourcePoolID is empty")
	}
	result = &v1.GetResourcePoolResponse{}
	err = bce.NewRequestBuilder(c.GetBceClient()).
		WithMethod(http.GET).
		WithURL(getResourcePoolUriWithID(resourcePoolID)).
		WithResult(result).
		Do()

	return result, err
}

func (c *Client) ListResourcePool(args *v1.ListResourcePoolRequest) (result *v1.ListResourcePoolResponse, err error) {
	if args == nil {
		return nil, fmt.Errorf("args is nil")
	}
	if args.PageNo <= 0 {
		args.PageNo = 1
	}
	if args.PageSize <= 0 {
		args.PageSize = 10
	}

	result = &v1.ListResourcePoolResponse{}
	err = bce.NewRequestBuilder(c.DefaultClient).
		WithMethod(http.GET).
		WithURL(listResourcePoolUri()).
		WithQueryParamFilter("keywordType", string(args.KeywordType)).
		WithQueryParamFilter("keyword", args.Keyword).
		WithQueryParamFilter("orderBy", string(args.OrderBy)).
		WithQueryParamFilter("order", string(args.Order)).
		WithQueryParamFilter("pageNo", strconv.Itoa(args.PageNo)).
		WithQueryParamFilter("pageSize", strconv.Itoa(args.PageSize)).
		WithResult(result).
		Do()
	return result, err
}

func (c *Client) ListNodeByResourcePoolID(resourcePoolID string, args *v1.ListResourcePoolNodeRequest) (result *v1.ListNodeByResourcePoolResponse, err error) {
	if resourcePoolID == "" {
		return nil, fmt.Errorf("resourcePoolID is empty")
	}
	if args != nil && (args.PageNo <= 0 || args.PageSize <= 0) {
		return nil, fmt.Errorf("invlaid pageNo or pageSize")
	}

	result = &v1.ListNodeByResourcePoolResponse{}
	err = bce.NewRequestBuilder(c.GetBceClient()).
		WithMethod(http.GET).
		WithURL(listResourcePoolNodesUri(resourcePoolID)).
		WithQueryParamFilter("orderBy", string(args.OrderBy)).
		WithQueryParamFilter("order", string(args.Order)).
		WithQueryParamFilter("pageNo", strconv.Itoa(args.PageNo)).
		WithQueryParamFilter("pageSize", strconv.Itoa(args.PageSize)).
		WithResult(result).
		Do()
	return result, err
}

func (c *Client) GetQueue(resourcePoolID, queueName string) (result *v1.GetQueuesResponse, err error) {
	if resourcePoolID == "" {
		return nil, fmt.Errorf("resourcePoolID is empty")
	}
	if queueName == "" {
		return nil, fmt.Errorf("queueName is empty")
	}
	result = &v1.GetQueuesResponse{}
	err = bce.NewRequestBuilder(c.GetBceClient()).
		WithMethod(http.GET).
		WithURL(getQueueUri(resourcePoolID, queueName)).
		WithResult(result).
		Do()

	return result, err
}

func (c *Client) ListQueue(resourcePoolID string, args *v1.ListQueueRequest) (result *v1.ListQueuesResponse, err error) {
	if resourcePoolID == "" {
		return nil, fmt.Errorf("resourcePoolID is empty")
	}
	if args.PageNo <= 0 {
		args.PageNo = 1
	}
	if args.PageSize <= 0 {
		args.PageSize = 10
	}
	result = &v1.ListQueuesResponse{}
	err = bce.NewRequestBuilder(c.GetBceClient()).
		WithMethod(http.GET).
		WithURL(listQueueUri(resourcePoolID)).
		WithQueryParamFilter("keywordType", string(args.KeywordType)).
		WithQueryParamFilter("keyword", args.Keyword).
		WithQueryParamFilter("orderBy", string(args.OrderBy)).
		WithQueryParamFilter("order", string(args.Order)).
		WithQueryParamFilter("pageNo", strconv.Itoa(args.PageNo)).
		WithQueryParamFilter("pageSize", strconv.Itoa(args.PageSize)).
		WithResult(result).
		Do()
	return result, err
}

func (c *Client) GetJob(options *v1.GetAIJobOptions) (*v1.OpenAPIGetJobResponse, error) {
	if options == nil {
		return nil, fmt.Errorf("GetAIJobOptions is nil")
	}
	jobID, resourcePoolId, queueID := options.JobID, options.ResourcePoolID, options.QueueID

	if resourcePoolId == "" {
		return nil, fmt.Errorf("resourcePoolId is empty")
	}
	if jobID == "" {
		return nil, fmt.Errorf("jobID is empty")
	}
	if (resourcePoolId == BHCMP_CLUSTER || resourcePoolId == SERVERLESS_CLUSTER) && queueID == "" {
		return nil, fmt.Errorf("bhcmp cluster or serverless cluster must be set queueID")
	}
	result := &v1.OpenAPIGetJobResponse{}
	err := bce.NewRequestBuilder(c.GetBceClient()).
		WithMethod(http.GET).
		WithURL(getAIJobUri(jobID)).
		WithQueryParamFilter("resourcePoolId", resourcePoolId).
		WithQueryParamFilter("queueID", queueID).
		WithResult(result).
		Do()
	return result, err
}

func (c *Client) DeleteJob(options *v1.DeleteAIJobOptions) (*v1.OpenAPIJobDeleteResponse, error) {
	if options == nil {
		return nil, fmt.Errorf("DeleteAIJobOptions is nil")
	}
	jobID, resourcePoolId, queueID := options.JobID, options.ResourcePoolID, options.QueueID

	if resourcePoolId == "" {
		return nil, fmt.Errorf("resourcePoolId is empty")
	}
	if jobID == "" {
		return nil, fmt.Errorf("jobID is empty")
	}
	if (resourcePoolId == BHCMP_CLUSTER || resourcePoolId == SERVERLESS_CLUSTER) && queueID == "" {
		return nil, fmt.Errorf("bhcmp cluster or serverless cluster must be set queueID")
	}
	result := &v1.OpenAPIJobDeleteResponse{}
	err := bce.NewRequestBuilder(c.GetBceClient()).
		WithMethod(http.DELETE).
		WithURL(getAIJobUri(jobID)).
		WithQueryParamFilter("resourcePoolId", resourcePoolId).
		WithQueryParamFilter("queueID", queueID).
		WithResult(result).
		Do()
	return result, err
}

func (c *Client) CreateJob(job *v1.OpenAPIJobCreateRequest, options *v1.CreateAIJobOptions) (*v1.OpenAPIJobCreateResponse, error) {
	if options == nil {
		return nil, fmt.Errorf("CreateAIJobOptions is nil")
	}
	resourcePoolId := options.ResourcePoolID

	if job == nil {
		return nil, fmt.Errorf("job is empty")
	}
	if resourcePoolId == "" {
		return nil, fmt.Errorf("resourcePoolId is empty")
	}
	if job.Name == "" {
		return nil, fmt.Errorf("job name is empty")
	}
	queueID := ""
	if resourcePoolId == BHCMP_CLUSTER || resourcePoolId == SERVERLESS_CLUSTER {
		queueID = job.Queue
	}
	result := &v1.OpenAPIJobCreateResponse{}
	err := bce.NewRequestBuilder(c.GetBceClient()).
		WithMethod(http.POST).
		WithURL(listJobUri()).
		WithBody(job).
		WithQueryParamFilter("resourcePoolId", resourcePoolId).
		WithQueryParamFilter("queueID", queueID).
		WithResult(result).
		Do()
	return result, err
}

func (c *Client) UpdateJob(job *v1.OpenAPIJobUpdateRequest, options *v1.UpdateAIJobOptions) (*v1.OpenAPIJobUpdateResponse, error) {
	if options == nil {
		return nil, fmt.Errorf("UpdateJob is nil")
	}
	jobID, resourcePoolId, queueID := options.JobID, options.ResourcePoolID, options.QueueID

	if job == nil {
		return nil, fmt.Errorf("job is empty")
	}
	if resourcePoolId == "" {
		return nil, fmt.Errorf("resourcePoolId is empty")
	}
	if job.Priority == "" {
		return nil, fmt.Errorf("job priority is empty")
	}
	if (resourcePoolId == BHCMP_CLUSTER || resourcePoolId == SERVERLESS_CLUSTER) && queueID == "" {
		return nil, fmt.Errorf("bhcmp cluster or serverless cluster must be set queueID")
	}
	result := &v1.OpenAPIJobUpdateResponse{}
	err := bce.NewRequestBuilder(c.GetBceClient()).
		WithMethod(http.PUT).
		WithURL(getAIJobUri(jobID)).
		WithQueryParamFilter("resourcePoolId", resourcePoolId).
		WithQueryParamFilter("queueID", queueID).
		WithBody(job).
		WithResult(result).
		Do()
	return result, err
}

func (c *Client) StopJob(options *v1.StopAIJobOptions) (*v1.OpenAPIJobStopResponse, error) {
	if options == nil {
		return nil, fmt.Errorf("StopAIJobOptions is nil")
	}
	jobID, resourcePoolId, queueID := options.JobID, options.ResourcePoolID, options.QueueID

	if jobID == "" {
		return nil, fmt.Errorf("jobID is empty")
	}
	if resourcePoolId == "" {
		return nil, fmt.Errorf("resourcePoolId is empty")
	}
	if (resourcePoolId == BHCMP_CLUSTER || resourcePoolId == SERVERLESS_CLUSTER) && queueID == "" {
		return nil, fmt.Errorf("bhcmp cluster or serverless cluster must be set queueID")
	}
	result := &v1.OpenAPIJobStopResponse{}
	err := bce.NewRequestBuilder(c.GetBceClient()).
		WithMethod(http.POST).
		WithURL(getStopAIJobUri(jobID)).
		WithQueryParamFilter("resourcePoolId", resourcePoolId).
		WithQueryParamFilter("queueID", queueID).
		WithResult(result).
		Do()
	return result, err
}

func (c *Client) GetJobNodesList(options *v1.GetJobNodesListOptions) (*v1.JobNodesListResponse, error) {
	if options == nil {
		return nil, fmt.Errorf("StopAIJobOptions is nil")
	}
	jobId, resourcePoolId, namespace := options.JobID, options.ResourcePoolID, options.Namespace

	if jobId == "" {
		return nil, fmt.Errorf("jobID is empty")
	}
	if resourcePoolId == "" {
		return nil, fmt.Errorf("resourcePoolId is empty")
	}
	result := &v1.JobNodesListResponse{}
	err := bce.NewRequestBuilder(c.GetBceClient()).
		WithMethod(http.GET).
		WithURL(getJobNodeListUri(jobId)).
		WithQueryParamFilter("resourcePoolId", resourcePoolId).
		WithQueryParamFilter("nameSpace", namespace).
		WithResult(result).
		Do()
	return result, err
}

func (c *Client) GetTaskEvent(args *v1.GetJobEventsRequest) (*v1.GetJobEventsResponse, error) {
	if args.JobID == "" {
		return nil, fmt.Errorf("jobID is empty")
	}
	if args.ResourcePoolID == "" {
		return nil, fmt.Errorf("resourcePoolId is empty")
	}
	if args.JobFramework == "" {
		return nil, fmt.Errorf("jobFramework is empty")
	}
	result := &v1.GetJobEventsResponse{}
	err := bce.NewRequestBuilder(c.GetBceClient()).
		WithMethod(http.GET).
		WithURL(getAIJobEventsUri(args.JobID)).
		WithQueryParamFilter("resourcePoolId", args.ResourcePoolID).
		WithQueryParamFilter("startTime", args.StartTime).
		WithQueryParamFilter("endTime", args.EndTime).
		WithQueryParamFilter("jobFramework", args.JobFramework).
		WithQueryParamFilter("nameSpace", args.Namespace).
		WithResult(result).
		Do()
	return result, err
}

func (c *Client) GetPodEvents(args *v1.GetPodEventsRequest) (*v1.GetPodEventsResponse, error) {
	if args.JobID == "" {
		return nil, fmt.Errorf("jobID is empty")
	}
	if args.ResourcePoolID == "" {
		return nil, fmt.Errorf("resourcePoolId is empty")
	}
	if args.JobFramework == "" {
		return nil, fmt.Errorf("jobFramework is empty")
	}
	if args.PodName == "" {
		return nil, fmt.Errorf("podName is empty")
	}
	result := &v1.GetPodEventsResponse{}
	err := bce.NewRequestBuilder(c.GetBceClient()).
		WithMethod(http.GET).
		WithURL(getPodEventsUri(args.JobID, args.PodName)).
		WithQueryParamFilter("resourcePoolId", args.ResourcePoolID).
		WithQueryParamFilter("nameSpace", args.Namespace).
		WithQueryParamFilter("startTime", args.StartTime).
		WithQueryParamFilter("endTime", args.EndTime).
		WithQueryParamFilter("nameSpace", args.Namespace).
		WithQueryParamFilter("jobFramework", args.JobFramework).
		WithResult(result).
		Do()
	return result, err
}

func (c *Client) GetPodLogs(args *v1.GetPodLogsRequest) (*v1.GetPodLogResponse, error) {
	if args.JobID == "" {
		return nil, fmt.Errorf("jobID is empty")
	}
	if args.ResourcePoolID == "" {
		return nil, fmt.Errorf("resourcePoolId is empty")
	}
	if args.PodName == "" {
		return nil, fmt.Errorf("podName is empty")
	}
	result := &v1.GetPodLogResponse{}
	err := bce.NewRequestBuilder(c.GetBceClient()).
		WithMethod(http.GET).
		WithURL(getAIJobPodLogsUri(args.JobID, args.PodName)).
		WithQueryParamFilter("resourcePoolId", args.ResourcePoolID).
		WithQueryParamFilter("startTime", args.StartTime).
		WithQueryParamFilter("endTime", args.EndTime).
		WithQueryParamFilter("maxLines", args.MaxLines).
		WithQueryParamFilter("namespace", args.Namespace).
		WithQueryParamFilter("chunck", args.Chunk).
		WithQueryParamFilter("container", args.Container).
		WithQueryParamFilter("marker", args.Marker).
		WithQueryParamFilter("filePath", args.FilePath).
		WithQueryParamFilter("logSource", args.LogSource).
		WithResult(result).
		Do()
	return result, err
}

func (c *Client) GetTaskMetrics(args *v1.GetTaskMetricsRequest) (*v1.GetTaskMetricsResponse, error) {
	if args.JobID == "" {
		return nil, fmt.Errorf("jobID is empty")
	}
	if args.ResourcePoolID == "" {
		return nil, fmt.Errorf("resourcePoolId is empty")
	}
	result := &v1.GetTaskMetricsResponse{}
	err := bce.NewRequestBuilder(c.GetBceClient()).
		WithMethod(http.GET).
		WithURL(getJobMetricsUri(args.JobID)).
		WithQueryParamFilter("resourcePoolId", args.ResourcePoolID).
		WithQueryParamFilter("startTime", args.StartTime).
		WithQueryParamFilter("endTime", args.EndTime).
		WithQueryParamFilter("timeStep", args.TimeStep).
		WithQueryParamFilter("nameSpace", args.Namespace).
		WithQueryParamFilter("rateInterval", args.RateInterval).
		WithQueryParamFilter("metricType", args.MetricType).
		WithResult(result).
		Do()
	return result, err
}

func (c *Client) GetWebSSHUrl(arg *v1.GetWebShellURLRequest) (*v1.GetWebShellURLResponse, error) {
	if arg.JobID == "" {
		return nil, fmt.Errorf("jobID is empty")
	}
	if arg.ResourcePoolID == "" {
		return nil, fmt.Errorf("resourcePoolId is empty")
	}
	if arg.PodName == "" {
		return nil, fmt.Errorf("podName is empty")
	}
	if (arg.ResourcePoolID == BHCMP_CLUSTER || arg.ResourcePoolID == SERVERLESS_CLUSTER) && arg.QueueID == "" {
		return nil, fmt.Errorf("bhcmp cluster or serverless cluster must be set queueID")
	}
	result := &v1.GetWebShellURLResponse{}
	err := bce.NewRequestBuilder(c.GetBceClient()).
		WithMethod(http.GET).
		WithURL(getWebSSHUri(arg.JobID, arg.PodName)).
		WithQueryParamFilter("resourcePoolId", arg.ResourcePoolID).
		WithQueryParamFilter("podName", arg.PodName).
		WithQueryParamFilter("nameSpace", arg.Namespace).
		WithQueryParamFilter("pingTimeoutSecond", arg.PingTimeoutSecond).
		WithQueryParamFilter("handshakeTimeoutSecond", arg.HandshakeTimeoutSecond).
		WithQueryParamFilter("queueID", arg.QueueID).
		WithResult(result).
		Do()
	return result, err
}

func (c *Client) ListJobs(args *v1.OpenAPIJobListRequest) (*v1.OpenAPIJobListResponse, error) {
	if args.ResourcePoolID == "" {
		return nil, fmt.Errorf("resourcePoolId is empty")
	}

	if (args.ResourcePoolID == BHCMP_CLUSTER || args.ResourcePoolID == SERVERLESS_CLUSTER) && args.Queue == "" {
		return nil, fmt.Errorf("bhcmp cluster or serverless cluster must be set queueID")
	}

	result := &v1.OpenAPIJobListResponse{}
	err := bce.NewRequestBuilder(c.GetBceClient()).
		WithMethod(http.GET).
		WithURL(listJobUri()).
		WithQueryParamFilter("resourcePoolId", args.ResourcePoolID).
		WithQueryParamFilter("pageNo", strconv.Itoa(args.PageNo)).
		WithQueryParamFilter("pageSize", strconv.Itoa(args.PageSize)).
		WithQueryParamFilter("queue", args.Queue).
		WithQueryParamFilter("queueID", args.Queue).
		WithQueryParamFilter("order", args.Order).
		WithQueryParamFilter("orderBy", args.OrderBy).
		WithResult(result).
		Do()
	return result, err
}

func (c *Client) FileUpload(args *v1.FileUploadRequest) (*v1.FileUploaderResponse, error) {
	resp := &v1.FileUploaderResponse{}

	err := bce.NewRequestBuilder(c.GetBceClient()).
		WithMethod(http.GET).
		WithURL(getFileUploadUri()).
		WithQueryParamFilter("resourcePoolId", args.ResourcePoolID).
		WithResult(resp).Do()
	if err != nil {
		return resp, fmt.Errorf("upload file error: %w", err)
	}

	sdkResp := &v1.FileUploaderResponse{
		RequestId: resp.RequestId,
		Result: v1.FileUploader{
			FilePath: resp.Result.FilePath,
			Token:    resp.Result.Token,
			FileID:   resp.Result.FileID,
			AK:       resp.Result.AK,
			SK:       resp.Result.SK,
			Bucket:   resp.Result.Bucket,
			Endpoint: resp.Result.Endpoint,
		},
	}
	return sdkResp, nil
}

func listResourcePoolUri() string {
	return URI_PREFIX + REQUEST_RESOURCE_POOL_URL
}
func listJobUri() string {
	return URI_PREFIX + REQUEST_JOB_URL
}

func getResourcePoolUriWithID(resourcePoolID string) string {
	return listResourcePoolUri() + "/" + resourcePoolID
}

func listResourcePoolNodesUri(resourcePoolID string) string {
	return getResourcePoolUriWithID(resourcePoolID) + REQUEST_NODE_URL
}

func listQueueUri(resourcePoolID string) string {
	return getResourcePoolUriWithID(resourcePoolID) + REQUEST_QUEUE_URL
}

func getQueueUri(resourcePoolID, queueName string) string {
	return getResourcePoolUriWithID(resourcePoolID) + REQUEST_QUEUE_URL + "/" + queueName
}

func getAIJobUri(jobID string) string {
	return listJobUri() + "/" + jobID
}

func getStopAIJobUri(jobID string) string {
	return getAIJobUri(jobID) + "/stop"
}

func getAIJobEventsUri(jobID string) string {
	return getAIJobUri(jobID) + "/events"
}

func getPodEventsUri(jobID, podName string) string {
	return getAIJobUri(jobID) + "/pods/" + podName + "/events"
}

func getAIJobPodLogsUri(jobID, podName string) string {
	return getAIJobUri(jobID) + "/pods/" + podName + "/logs"
}

func getJobNodeListUri(jobID string) string {
	return getAIJobUri(jobID) + "/nodes"
}

func getJobMetricsUri(jobID string) string {
	return getAIJobUri(jobID) + "/metrics"
}

func getWebSSHUri(jobID, podName string) string {
	return getAIJobUri(jobID) + "/pods/" + podName + "/webterminal"
}

func getFileUploadUri() string {
	return listJobUri() + "/cluster/ai/uploadCode"
}

func getCreateNotifyRuleUri() string {
	return listJobUri() + "/notify/rule"
}
