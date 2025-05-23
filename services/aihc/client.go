package aihc

import (
	resourcepoolv1 "github.com/baidubce/bce-sdk-go/services/aihc/api/v1"
	"github.com/baidubce/bce-sdk-go/services/aihc/client"
	inference "github.com/baidubce/bce-sdk-go/services/aihc/inference/v2"
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
func (clientset *Client) GetJob(options *resourcepoolv1.GetAIJobOptions) (*resourcepoolv1.OpenAPIGetJobResponse, error) {
	return clientset.resourceClient.GetJob(options)
}
func (clientset *Client) DeleteJob(options *resourcepoolv1.DeleteAIJobOptions) (*resourcepoolv1.OpenAPIJobDeleteResponse, error) {
	return clientset.resourceClient.DeleteJob(options)
}
func (clientset *Client) CreateJob(args *resourcepoolv1.OpenAPIJobCreateRequest, options *resourcepoolv1.CreateAIJobOptions) (*resourcepoolv1.OpenAPIJobCreateResponse, error) {
	return clientset.resourceClient.CreateJob(args, options)
}
func (clientset *Client) UpdateJob(args *resourcepoolv1.OpenAPIJobUpdateRequest, options *resourcepoolv1.UpdateAIJobOptions) (*resourcepoolv1.OpenAPIJobUpdateResponse, error) {
	return clientset.resourceClient.UpdateJob(args, options)
}
func (clientset *Client) StopJob(options *resourcepoolv1.StopAIJobOptions) (*resourcepoolv1.OpenAPIJobStopResponse, error) {
	return clientset.resourceClient.StopJob(options)
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
func (clientset *Client) GetJobNodesList(options *resourcepoolv1.GetJobNodesListOptions) (*resourcepoolv1.JobNodesListResponse, error) {
	return clientset.resourceClient.GetJobNodesList(options)
}
func (clientset *Client) GetTaskMetrics(args *resourcepoolv1.GetTaskMetricsRequest) (*resourcepoolv1.GetTaskMetricsResponse, error) {
	return clientset.resourceClient.GetTaskMetrics(args)
}
func (clientset *Client) GetWebSSHUrl(args *resourcepoolv1.GetWebShellURLRequest) (*resourcepoolv1.GetWebShellURLResponse, error) {
	return clientset.resourceClient.GetWebSSHUrl(args)
}
func (clientset *Client) FileUpload(args *resourcepoolv1.FileUploadRequest) (*resourcepoolv1.FileUploaderResponse, error) {
	return clientset.resourceClient.FileUpload(args)
}

// inference
func (clientset *Client) CreateService(args *inference.ServiceConf, clientToken string) (*inference.CreateServiceResult, error) {
	return clientset.inferenceClient.CreateService(args, clientToken)
}
func (clientset *Client) ListService(args *inference.ListServiceArgs) (*inference.ListServiceResult, error) {
	return clientset.inferenceClient.ListService(args)
}

func (clientset *Client) ListServiceStats(args *inference.ListServiceStatsArgs) (*inference.ListServiceStatsResult, error) {
	return clientset.inferenceClient.ListServiceStats(args)
}

func (clientset *Client) ServiceDetails(args *inference.ServiceDetailsArgs) (*inference.ServiceDetailsResult, error) {
	return clientset.inferenceClient.ServiceDetails(args)
}

func (clientset *Client) UpdateService(args *inference.UpdateServiceArgs) (*inference.UpdateServiceResult, error) {
	return clientset.inferenceClient.UpdateService(args)
}

func (clientset *Client) ScaleService(args *inference.ScaleServiceArgs) (*inference.ScaleServiceResult, error) {
	return clientset.inferenceClient.ScaleService(args)
}

func (clientset *Client) PubAccess(args *inference.PubAccessArgs) (*inference.PubAccessResult, error) {
	return clientset.inferenceClient.PubAccess(args)
}

func (clientset *Client) ListChange(args *inference.ListChangeArgs) (*inference.ListChangeResult, error) {
	return clientset.inferenceClient.ListChange(args)
}

func (clientset *Client) ChangeDetail(args *inference.ChangeDetailArgs) (*inference.ChangeDetailResult, error) {
	return clientset.inferenceClient.ChangeDetail(args)
}

func (clientset *Client) DeleteService(args *inference.DeleteServiceArgs) (*inference.DeleteServiceResult, error) {
	return clientset.inferenceClient.DeleteService(args)
}

func (clientset *Client) ListPod(args *inference.ListPodArgs) (*inference.ListPodResult, error) {
	return clientset.inferenceClient.ListPod(args)
}

func (clientset *Client) BlockPod(args *inference.BlockPodArgs) (*inference.BlockPodResult, error) {
	return clientset.inferenceClient.BlockPod(args)
}

func (clientset *Client) DeletePod(args *inference.DeletePodArgs) (*inference.DeletePodResult, error) {
	return clientset.inferenceClient.DeletePod(args)
}

func (clientset *Client) ListPodGroups(args *inference.ListPodGroupsArgs) (*inference.ListPodGroupsResult, error) {
	return clientset.inferenceClient.ListPodGroups(args)
}
