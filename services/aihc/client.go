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
