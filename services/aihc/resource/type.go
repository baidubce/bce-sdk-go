package resource

import v1 "github.com/baidubce/bce-sdk-go/services/aihc/api/v1"

type Interface interface {
	GetResourcePool(resourcePoolID string) (*v1.GetResourcePoolResponse, error)
	ListResourcePool(args *v1.ListResourcePoolRequest) (*v1.ListResourcePoolResponse, error)
	ListNodeByResourcePoolID(resourcePoolID string, args *v1.ListResourcePoolNodeRequest) (*v1.ListNodeByResourcePoolResponse, error)
	GetQueue(resourcePoolID, queueName string) (*v1.GetQueuesResponse, error)
	ListQueue(resourcePoolID string, args *v1.ListQueueRequest) (*v1.ListQueuesResponse, error)
}
