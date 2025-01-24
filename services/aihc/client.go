package aihc

import (
	resourcepoolv1 "github.com/baidubce/bce-sdk-go/services/aihc/api/v1"
	"github.com/baidubce/bce-sdk-go/services/aihc/client"
	"github.com/baidubce/bce-sdk-go/services/aihc/inference"
	inferenceapi "github.com/baidubce/bce-sdk-go/services/aihc/inference/api"
	"github.com/baidubce/bce-sdk-go/services/aihc/resource"
)

type Interface interface {
	resource.Interface
	inference.Interface
}

type Client struct {
	resourceClient  *resource.Client
	inferenceClient *inference.Client
	GlobalClient    *client.Client
}

func NewClient(ak, sk, endpoint string) (Interface, error) {
	clientset := &Client{}
	resourceClient, err := resource.NewClient(ak, sk, endpoint)
	if err != nil {
		return nil, err
	}
	clientset.resourceClient = resourceClient

	inferenceEndpoint := endpoint
	if client.IsRegionedEndpoint(inferenceEndpoint) {
		inferenceEndpoint = client.DEFAULT_GLOBAL_ENDPOINT
	}
	inferenceClient, err := inference.NewClient(ak, sk, inferenceEndpoint)
	if err != nil {
		return nil, err
	}
	clientset.inferenceClient = inferenceClient
	return clientset, nil
}

// NewClientWithSTS make the aihc inference service client with STS configuration.
func NewClientWithSTS(ak, sk, sessionToken, endpoint string) (Interface, error) {
	clientset := &Client{}
	resourceClient, err := resource.NewClientWithSTS(ak, sk, sessionToken, endpoint)
	if err != nil {
		return nil, err
	}
	clientset.resourceClient = resourceClient

	inferenceClient, err := inference.NewClientWithSTS(ak, sk, sessionToken, endpoint)
	if err != nil {
		return nil, err
	}
	clientset.inferenceClient = inferenceClient
	return clientset, nil
}

// resourcepool&queue
func (clientset *Client) ListResourcePool(args *resourcepoolv1.ListResourcePoolRequest) (*resourcepoolv1.ListResourcePoolResponse, error) {
	return clientset.resourceClient.ListResourcePool(args)
}
func (clientset *Client) GetResourcePool(resourcePoolID string) (*resourcepoolv1.GetResourcePoolResponse, error) {
	return clientset.resourceClient.GetResourcePool(resourcePoolID)
}
func (clientset *Client) ListNodeByResourcePoolID(resourcePoolID string,
	args *resourcepoolv1.ListResourcePoolNodeRequest) (*resourcepoolv1.ListNodeByResourcePoolResponse, error) {
	return clientset.resourceClient.ListNodeByResourcePoolID(resourcePoolID, args)
}
func (clientset *Client) ListQueue(resourcePoolID string, args *resourcepoolv1.ListQueueRequest) (*resourcepoolv1.ListQueuesResponse, error) {
	return clientset.resourceClient.ListQueue(resourcePoolID, args)
}
func (clientset *Client) GetQueue(resourcePoolID, queueName string) (*resourcepoolv1.GetQueuesResponse, error) {
	return clientset.resourceClient.GetQueue(resourcePoolID, queueName)
}
func (clientset *Client) ListJobs(args *resourcepoolv1.OpenAPIJobListRequest) (*resourcepoolv1.OpenAPIJobListResponse, error) {
	return clientset.resourceClient.ListJobs(args)
}

// aijob
func (clientset *Client) GetJob(jobID, resourcePoolId string) (*resourcepoolv1.OpenAPIGetJobResponse, error) {
	return clientset.resourceClient.GetJob(jobID, resourcePoolId)
}
func (clientset *Client) DeleteJob(jobID, resourcePoolId string) (*resourcepoolv1.OpenAPIJobDeleteResponse, error) {
	return clientset.resourceClient.DeleteJob(jobID, resourcePoolId)
}
func (clientset *Client) CreateJob(args *resourcepoolv1.OpenAPIJobCreateRequest, resourcePoolId string) (*resourcepoolv1.OpenAPIJobCreateResponse, error) {
	return clientset.resourceClient.CreateJob(args, resourcePoolId)
}
func (clientset *Client) UpdateJob(args *resourcepoolv1.OpenAPIJobUpdateRequest, jobID, resourcePoolId string) (*resourcepoolv1.OpenAPIJobUpdateResponse, error) {
	return clientset.resourceClient.UpdateJob(args, jobID, resourcePoolId)
}
func (clientset *Client) StopJob(jobID, resourcePoolId string) (*resourcepoolv1.OpenAPIJobStopResponse, error) {
	return clientset.resourceClient.StopJob(jobID, resourcePoolId)
}
func (clientset *Client) GetTaskEvent(args *resourcepoolv1.GetJobEventsRequest) (*resourcepoolv1.GetJobEventsResponse, error) {
	return clientset.resourceClient.GetTaskEvent(args)
}
func (clientset *Client) GetPodEvents(args *resourcepoolv1.GetPodEventsRequest) (*resourcepoolv1.GetPodEventsResponse, error) {
	return clientset.resourceClient.GetPodEvents(args)
}
func (clientset *Client) GetPodLogs(args *resourcepoolv1.GetPodLogsRequest) (*resourcepoolv1.GetPodLogResponse, error) {
	return clientset.resourceClient.GetPodLogs(args)
}
func (clientset *Client) GetJobNodesList(jobID, resourcePoolId, namespace string) (*resourcepoolv1.JobNodesListResponse, error) {
	return clientset.resourceClient.GetJobNodesList(jobID, resourcePoolId, namespace)
}
func (clientset *Client) GetTaskMetrics(args *resourcepoolv1.GetTaskMetricsRequest) (*resourcepoolv1.GetTaskMetricsResponse, error) {
	return clientset.resourceClient.GetTaskMetrics(args)
}
func (clientset *Client) GetWebSSHUrl(args *resourcepoolv1.GetWebShellURLRequest) (*resourcepoolv1.GetWebShellURLResponse, error) {
	return clientset.resourceClient.GetWebSSHUrl(args)
}

// inference
func (clientset *Client) CreateApp(args *inferenceapi.CreateAppArgs, region string, extraInfo map[string]string) (*inferenceapi.CreateAppResult, error) {
	return clientset.inferenceClient.CreateApp(args, region, extraInfo)
}
func (clientset *Client) ListApp(args *inferenceapi.ListAppArgs, region string, extraInfo map[string]string) (*inferenceapi.ListAppResult, error) {
	return clientset.inferenceClient.ListApp(args, region, extraInfo)
}

func (clientset *Client) ListAppStats(args *inferenceapi.ListAppStatsArgs, region string) (*inferenceapi.ListAppStatsResult, error) {
	return clientset.inferenceClient.ListAppStats(args, region)
}

func (clientset *Client) AppDetails(args *inferenceapi.AppDetailsArgs,
	region string) (*inferenceapi.AppDetailsResult, error) {
	return clientset.inferenceClient.AppDetails(args, region)
}

func (clientset *Client) UpdateApp(args *inferenceapi.UpdateAppArgs,
	region string) (*inferenceapi.UpdateAppResult, error) {
	return clientset.inferenceClient.UpdateApp(args, region)
}

func (clientset *Client) ScaleApp(args *inferenceapi.ScaleAppArgs,
	region string) (*inferenceapi.ScaleAppResult, error) {
	return clientset.inferenceClient.ScaleApp(args, region)
}

func (clientset *Client) PubAccess(args *inferenceapi.PubAccessArgs,
	region string) (*inferenceapi.PubAccessResult, error) {
	return clientset.inferenceClient.PubAccess(args, region)
}

func (clientset *Client) ListChange(args *inferenceapi.ListChangeArgs,
	region string) (*inferenceapi.ListChangeResult, error) {
	return clientset.inferenceClient.ListChange(args, region)
}

func (clientset *Client) ChangeDetail(args *inferenceapi.ChangeDetailArgs, region string) (*inferenceapi.ChangeDetailResult, error) {
	return clientset.inferenceClient.ChangeDetail(args, region)
}

func (clientset *Client) DeleteApp(args *inferenceapi.DeleteAppArgs,
	region string) (*inferenceapi.DeleteAppResult, error) {
	return clientset.inferenceClient.DeleteApp(args, region)
}

func (clientset *Client) ListPod(args *inferenceapi.ListPodArgs,
	region string) (*inferenceapi.ListPodResult, error) {
	return clientset.inferenceClient.ListPod(args, region)
}

func (clientset *Client) BlockPod(args *inferenceapi.BlockPodArgs,
	region string) (*inferenceapi.BlockPodResult, error) {
	return clientset.inferenceClient.BlockPod(args, region)
}

func (clientset *Client) DeletePod(args *inferenceapi.DeletePodArgs,
	region string) (*inferenceapi.DeletePodResult, error) {
	return clientset.inferenceClient.DeletePod(args, region)
}

func (clientset *Client) ListBriefResPool(args *inferenceapi.ListBriefResPoolArgs, region string) (*inferenceapi.ListBriefResPoolResult, error) {
	return clientset.inferenceClient.ListBriefResPool(args, region)
}

func (clientset *Client) ResPoolDetail(args *inferenceapi.ResPoolDetailArgs, region string) (*inferenceapi.ResPoolDetailResult, error) {
	return clientset.inferenceClient.ResPoolDetail(args, region)
}
