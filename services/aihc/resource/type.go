package resource

import v1 "github.com/baidubce/bce-sdk-go/services/aihc/api/v1"

type Interface interface {
	GetResourcePool(resourcePoolID string) (*v1.GetResourcePoolResponse, error)
	ListResourcePool(args *v1.ListResourcePoolRequest) (*v1.ListResourcePoolResponse, error)
	ListNodeByResourcePoolID(resourcePoolID string, args *v1.ListResourcePoolNodeRequest) (*v1.ListNodeByResourcePoolResponse, error)
	GetQueue(resourcePoolID, queueName string) (*v1.GetQueuesResponse, error)
	ListQueue(resourcePoolID string, args *v1.ListQueueRequest) (*v1.ListQueuesResponse, error)

	GetJob(options *v1.GetAIJobOptions) (*v1.OpenAPIGetJobResponse, error)
	DeleteJob(options *v1.DeleteAIJobOptions) (*v1.OpenAPIJobDeleteResponse, error)
	ListJobs(args *v1.OpenAPIJobListRequest) (*v1.OpenAPIJobListResponse, error)
	CreateJob(job *v1.OpenAPIJobCreateRequest, options *v1.CreateAIJobOptions) (*v1.OpenAPIJobCreateResponse, error)
	UpdateJob(job *v1.OpenAPIJobUpdateRequest, options *v1.UpdateAIJobOptions) (*v1.OpenAPIJobUpdateResponse, error)
	StopJob(options *v1.StopAIJobOptions) (*v1.OpenAPIJobStopResponse, error)
	GetTaskEvent(args *v1.GetJobEventsRequest) (*v1.GetJobEventsResponse, error)
	GetPodLogs(args *v1.GetPodLogsRequest) (*v1.GetPodLogResponse, error)
	GetJobNodesList(options *v1.GetJobNodesListOptions) (*v1.JobNodesListResponse, error)
	GetPodEvents(args *v1.GetPodEventsRequest) (*v1.GetPodEventsResponse, error)
	GetTaskMetrics(args *v1.GetTaskMetricsRequest) (*v1.GetTaskMetricsResponse, error)
	GetWebSSHUrl(arg *v1.GetWebShellURLRequest) (*v1.GetWebShellURLResponse, error)
	FileUpload(args *v1.FileUploadRequest) (*v1.FileUploaderResponse, error)
}
