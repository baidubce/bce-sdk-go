package resource

import (
	"encoding/json"
	"fmt"

	"github.com/baidubce/bce-sdk-go/services/aihc"
	v1 "github.com/baidubce/bce-sdk-go/services/aihc/api/v1"
	"github.com/baidubce/bce-sdk-go/util/log"
)

const (
	ak_test       = ""
	sk_test       = ""
	endpoint_test = ""
)

var (
	RESOURCE_POOL_ID string
	AIJobID          string
	PodName          string
	ImageID          string
	AIJobName        string
	MetricType       string
	QueueID          string
	DataSourceType   string
	SourcePath       string
	MountPath        string
	DataSourceName   string
)

func GetJob() {
	ak, sk, endpoint := ak_test, sk_test, endpoint_test
	resourcePoolID, JobID := RESOURCE_POOL_ID, AIJobID

	client, _ := aihc.NewClient(ak, sk, endpoint)
	result, err := client.GetJob(&v1.GetAIJobOptions{
		JobID:          JobID,
		ResourcePoolID: resourcePoolID,
		QueueID:        QueueID,
	})

	if err != nil {
		panic(err)
	}

	jsonBytes, _ := json.Marshal(result)
	fmt.Println(string(jsonBytes))
}

func DeleleJob() {
	ak, sk, endpoint := ak_test, sk_test, endpoint_test
	resourcePoolID, JobID := RESOURCE_POOL_ID, AIJobID

	client, _ := aihc.NewClient(ak, sk, endpoint)
	result, err := client.DeleteJob(&v1.DeleteAIJobOptions{
		JobID:          JobID,
		ResourcePoolID: resourcePoolID,
		QueueID:        QueueID,
	})

	if err != nil {
		panic(err)
	}

	jsonBytes, _ := json.Marshal(result)
	fmt.Println(string(jsonBytes))
}

func CreateJob() {
	ak, sk, endpoint := ak_test, sk_test, endpoint_test
	resourcePoolID := RESOURCE_POOL_ID

	jobConfig := &v1.OpenAPIJobCreateRequest{
		Name: AIJobName,
		Datasources: []v1.OpenAPIDatasource{
			{
				Type:       DataSourceType,
				SourcePath: SourcePath,
				MountPath:  MountPath,
				Name:       DataSourceName,
				Options:    v1.AIJobDatasourceOptions{},
			},
		},
		JobSpec: v1.OpenAPIAIJobSpec{
			Command:  `echo "hello sdk"; sleep infinity`,
			Replicas: 1,
			Image:    ImageID,
			Resources: []v1.OpenAPIResource{
				{
					Name:     "cpu",
					Quantity: 1,
				},
			},
			EnableRDMA: false,
		},
		Queue:      QueueID,
		EnableBccl: false,
	}
	client, _ := aihc.NewClient(ak, sk, endpoint)
	result, err := client.CreateJob(jobConfig, &v1.CreateAIJobOptions{
		ResourcePoolID: resourcePoolID,
	})

	if err != nil {
		panic(err)
	}

	jsonBytes, _ := json.Marshal(result)
	fmt.Println(string(jsonBytes))
}

func UpdateJob() {
	ak, sk, endpoint := ak_test, sk_test, endpoint_test
	resourcePoolID := RESOURCE_POOL_ID
	jobID := AIJobID

	jobConfig := &v1.OpenAPIJobUpdateRequest{
		Priority: "high",
	}
	client, _ := aihc.NewClient(ak, sk, endpoint)
	result, err := client.UpdateJob(jobConfig, &v1.UpdateAIJobOptions{
		JobID:          jobID,
		ResourcePoolID: resourcePoolID,
		QueueID:        QueueID,
	})

	if err != nil {
		panic(err)
	}
	jsonBytes, _ := json.Marshal(result)
	fmt.Println(string(jsonBytes))
}

func StopJob() {
	ak, sk, endpoint := ak_test, sk_test, endpoint_test
	resourcePoolID := RESOURCE_POOL_ID
	jobID := AIJobID

	client, _ := aihc.NewClient(ak, sk, endpoint)
	result, err := client.StopJob(&v1.StopAIJobOptions{
		JobID:          jobID,
		ResourcePoolID: resourcePoolID,
		QueueID:        QueueID,
	})
	log.Infof("stop job result: %v", result)
	if err != nil {
		panic(err)
	}
	jsonBytes, _ := json.Marshal(result)
	fmt.Println(string(jsonBytes))
}

func GetTaskEvents() {
	ak, sk, endpoint := ak_test, sk_test, endpoint_test

	req := &v1.GetJobEventsRequest{
		Namespace:      "",
		JobFramework:   "PyTorchJob",
		StartTime:      "",
		EndTime:        "",
		JobID:          AIJobID,
		ResourcePoolID: RESOURCE_POOL_ID,
	}

	client, _ := aihc.NewClient(ak, sk, endpoint)
	result, err := client.GetTaskEvent(req)

	if err != nil {
		panic(err)
	}
	jsonBytes, _ := json.Marshal(result)
	fmt.Println(string(jsonBytes))
}

func GetTaskPodLogs() {
	ak, sk, endpoint := ak_test, sk_test, endpoint_test

	req := &v1.GetPodLogsRequest{
		JobID:          AIJobID,
		ResourcePoolID: RESOURCE_POOL_ID,
		PodName:        PodName,
		Namespace:      "default",
		StartTime:      "",
		EndTime:        "",
		MaxLines:       "",
		Container:      "",
		Chunk:          "",
	}

	client, _ := aihc.NewClient(ak, sk, endpoint)
	result, err := client.GetPodLogs(req)

	if err != nil {
		panic(err)
	}
	jsonBytes, _ := json.Marshal(result)
	fmt.Println(string(jsonBytes))
}

func GetJobNodesList() {
	ak, sk, endpoint := ak_test, sk_test, endpoint_test
	resourcePoolID := RESOURCE_POOL_ID
	jobID := AIJobID
	namespace := ""

	client, _ := aihc.NewClient(ak, sk, endpoint)
	result, err := client.GetJobNodesList(&v1.GetJobNodesListOptions{
		JobID:          jobID,
		ResourcePoolID: resourcePoolID,
		Namespace:      namespace,
	})

	if err != nil {
		panic(err)
	}
	jsonBytes, _ := json.Marshal(result)
	fmt.Println(string(jsonBytes))
}

func GetPodEvents() {
	ak, sk, endpoint := ak_test, sk_test, endpoint_test
	req := &v1.GetPodEventsRequest{
		JobID:          AIJobID,
		ResourcePoolID: RESOURCE_POOL_ID,
		Namespace:      "",
		JobFramework:   "PyTorchJob",
		StartTime:      "",
		EndTime:        "",
		PodName:        PodName,
	}

	client, _ := aihc.NewClient(ak, sk, endpoint)
	result, err := client.GetPodEvents(req)

	if err != nil {
		panic(err)
	}
	jsonBytes, _ := json.Marshal(result)
	fmt.Println(string(jsonBytes))
}

func GetTaskMetrics() {
	ak, sk, endpoint := ak_test, sk_test, endpoint_test
	req := &v1.GetTaskMetricsRequest{
		StartTime:      "",
		ResourcePoolID: RESOURCE_POOL_ID,
		EndTime:        "",
		TimeStep:       "",
		MetricType:     MetricType,
		JobID:          AIJobID,
		Namespace:      "",
		RateInterval:   "",
	}

	client, _ := aihc.NewClient(ak, sk, endpoint)
	result, err := client.GetTaskMetrics(req)

	if err != nil {
		panic(err)
	}
	jsonBytes, _ := json.Marshal(result)
	fmt.Println(string(jsonBytes))
}

func GetWebSSHUrl() {
	ak, sk, endpoint := ak_test, sk_test, endpoint_test

	req := &v1.GetWebShellURLRequest{
		JobID:                  AIJobID,
		ResourcePoolID:         RESOURCE_POOL_ID,
		PodName:                PodName,
		Namespace:              "",
		PingTimeoutSecond:      "",
		HandshakeTimeoutSecond: "",
		QueueID:                QueueID,
	}

	client, _ := aihc.NewClient(ak, sk, endpoint)
	result, err := client.GetWebSSHUrl(req)

	if err != nil {
		panic(err)
	}
	jsonBytes, _ := json.Marshal(result)
	fmt.Println(string(jsonBytes))
}

func ListJobs() {
	ak, sk, endpoint := ak_test, sk_test, endpoint_test
	client, _ := aihc.NewClient(ak, sk, endpoint)
	req := &v1.OpenAPIJobListRequest{
		ResourcePoolID: RESOURCE_POOL_ID,
		PageNo:         1,
		PageSize:       3,
		Queue:          QueueID,
	}
	result, err := client.ListJobs(req)

	if err != nil {
		panic(err)
	}
	jsonBytes, _ := json.Marshal(result)
	fmt.Println(string(jsonBytes))
}

func FileUpload() {
	ak, sk, endpoint := ak_test, sk_test, endpoint_test
	client, _ := aihc.NewClient(ak, sk, endpoint)
	req := &v1.FileUploadRequest{
		FilePaths:      []string{},
		ResourcePoolID: RESOURCE_POOL_ID,
	}

	result, err := client.FileUpload(req)
	if err != nil {
		panic(err)
	}
	jsonBytes, _ := json.Marshal(result)
	fmt.Println(string(jsonBytes))
}
