package media

import (
	"github.com/baidubce/bce-sdk-go/auth"
	"github.com/baidubce/bce-sdk-go/bce"
	"github.com/baidubce/bce-sdk-go/services/media/api"
)

const DEFAULT_SERVICE_DOMAIN = "media.bj.baidubce.com"

// mcp(media) client extends BceClient
type Client struct {
	*bce.BceClient
}

// NewClient make MCP(media) service client with defualt configuration
// endPoint value can chose bj and sz or gz defualt bj
func NewClient(ak, sk, endPoint string) (*Client, error) {
	credentials, err := auth.NewBceCredentials(ak, sk)
	if err != nil {
		return nil, err
	}
	if endPoint == "" {
		endPoint = DEFAULT_SERVICE_DOMAIN
	}
	defaultSignOptions := &auth.SignOptions{
		HeadersToSign: auth.DEFAULT_HEADERS_TO_SIGN,
		ExpireSeconds: auth.DEFAULT_EXPIRE_SECONDS}
	defaultConf := &bce.BceClientConfiguration{
		Endpoint:                  endPoint,
		Region:                    bce.DEFAULT_REGION,
		UserAgent:                 bce.DEFAULT_USER_AGENT,
		Credentials:               credentials,
		SignOption:                defaultSignOptions,
		Retry:                     bce.DEFAULT_RETRY_POLICY,
		ConnectionTimeoutInMillis: bce.DEFAULT_CONNECTION_TIMEOUT_IN_MILLIS}
	v1Signer := &auth.BceV1Signer{}

	client := &Client{bce.NewBceClient(defaultConf, v1Signer)}
	return client, nil
}

// create a simple pipeline with pipelieName,soureBucket,targetBucket,capacity
func (cli *Client) CreatePipeline(pipelineName, sourceBucket, targetBucket string, capacity int) error {

	return api.CreatePipeline(cli, pipelineName, sourceBucket, targetBucket, capacity)
}

// list all pipelines for user
func (cli *Client) ListPipelines() (*api.ListPipelinesResponse, error) {
	return api.ListPipelines(cli)
}

// query pipeline by piplineName
func (cli *Client) GetPipeline(pipelineName string) (*api.PipelineStatus, error) {
	return api.GetPipeline(cli, pipelineName)
}

func (cli *Client) GetPipelineUpdate(pipelineName string) (*api.UpdatePipelineArgs, error) {
	return api.GetPipelineUpdate(cli, pipelineName)
}

// delete pipeline by pipelineName
func (cli *Client) DeletePipeline(pipelineName string) error {
	return api.DeletePipeline(cli, pipelineName)
}

// update pipeline with UpdatePipelineArgs
func (cli *Client) UpdatePipeline(pipelineName string, updatePipelineArgs *api.UpdatePipelineArgs) error {
	return api.UpdatePipeline(cli, pipelineName, updatePipelineArgs)
}

// create transcoding job with pipelineName, sourceKey, targetKey, presetName
func (cli *Client) CreateJob(pipelineName, sourceKey, targetKey, presetName string) (*api.CreateJobResponse, error) {
	return api.CreateJob(cli, pipelineName, sourceKey, targetKey, presetName)
}

// create trandcoding job with customize params
func (cli *Client) CreateJobCustomize(args *api.CreateJobArgs) (*api.CreateJobResponse, error) {
	return api.CreateJobCustomize(cli, args)
}

// list all jobs with piplineName
func (cli *Client) ListTranscodingJobs(pipelineName string) (*api.ListTranscodingJobsResponse, error) {
	return api.ListTranscodingJobs(cli, pipelineName)
}

// get transcoding job by jobId
func (cli *Client) GetTranscodingJob(jobId string) (*api.GetTranscodingJobResponse, error) {
	return api.GetTranscodingJob(cli, jobId)
}

// list all presets
func (cli *Client) ListPresets() (*api.ListPresetsResponse, error) {
	return api.ListPresets(cli)
}

// get preset by presetName
func (cli *Client) GetPreset(presetName string) (*api.Preset, error) {
	return api.GetPreset(cli, presetName)
}

// create preset at the same time perform container format conversion
func (cli *Client) CreatePreset(presetName, description, container string) error {
	return api.CreatePreset(cli, presetName, description, container)
}

// create preset with user-defined configuration
func (cli *Client) CreatePrestCustomize(preset *api.Preset) error {
	return api.CreatePrestCustomize(cli, preset)
}

// update preset
func (cli *Client) UpdatePreset(preset *api.Preset) error {
	return api.UpdatePreset(cli, preset)
}

// get media information with bucket and key
func (cli *Client) GetMediaInfoOfFile(bucket, key string) (*api.GetMediaInfoOfFileResponse, error) {
	return api.GetMediaInfoOfFile(cli, bucket, key)
}

// this option implements create thumbnail job function overloading
type Option func(thumbnailOptional *api.ThumbnailOptional)

func PresetNameOp(presetName string) Option {
	return func(thumbnailOptional *api.ThumbnailOptional) {
		thumbnailOptional.PresetName = presetName
	}
}

func TargetOp(target *api.ThumbnailTarget) Option {
	return func(thumbnailOptional *api.ThumbnailOptional) {
		thumbnailOptional.Target = target
	}
}

func CaptureOp(capture *api.ThumbnailCapture) Option {
	return func(thumbnailOptional *api.ThumbnailOptional) {
		thumbnailOptional.Capture = capture
	}
}

func DelogoAreaOp(delogoArea *api.Area) Option {
	return func(thumbnailOptional *api.ThumbnailOptional) {
		thumbnailOptional.DelogoArea = delogoArea
	}
}

func CropOp(crop *api.Area) Option {
	return func(thumbnailOptional *api.ThumbnailOptional) {
		thumbnailOptional.Crop = crop
	}
}

func SourceOp(source *api.ThumbnailSource) Option {
	return func(thumbnailOptional *api.ThumbnailOptional) {
		thumbnailOptional.ThumbnailSource = source
	}
}

// create thumbnail job.
// you can create a thumbnail job with pipelineName and sourceKey and ThumbnailCapture and ThumbnailTarget or other args
func (cli *Client) CreateThumbnailJob(pipelineName, sourceKey string, ops ...Option) (*api.CreateJobResponse, error) {
	var thumbnailOptional api.ThumbnailOptional
	for _, op := range ops {
		op(&thumbnailOptional)
	}

	createThumbnialArgs := &api.CreateThumbnailJobArgs{}
	createThumbnialArgs.PipelineName = pipelineName
	source := &api.ThumbnailSource{}
	source.Key = sourceKey
	createThumbnialArgs.ThumbnailSource = source
	createThumbnialArgs.PresetName = thumbnailOptional.PresetName
	target := thumbnailOptional.Target
	createThumbnialArgs.ThumbnailTarget = target
	createThumbnialArgs.ThumbnailCapture = thumbnailOptional.Capture
	createThumbnialArgs.Area = thumbnailOptional.DelogoArea
	createThumbnialArgs.Crop = thumbnailOptional.Crop
	return api.CreateThumbnailJob(cli, pipelineName, sourceKey, createThumbnialArgs)
}

// query thumbanil job by jobId
func (cli *Client) GetThumbanilJob(jobId string) (*api.GetThumbnailJobResponse, error) {
	return api.GetThumbanilJob(cli, jobId)
}

// get thumbanil job by pipelineName
func (cli *Client) ListThumbnailJobs(pipelineName string) (*api.ListThumbnailJobsResponse, error) {
	return api.ListThumbnailJobs(cli, pipelineName)
}

// create watermark job
func (cli *Client) CreateWaterMark(watermarks *api.CreateWaterMarkArgs) (*api.CreateWaterMarkResponse, error) {
	return api.CreateWaterMark(cli, watermarks)
}

// get watermark by watermarkId
func (cli *Client) GetWaterMark(watermarkId string) (*api.GetWaterMarkResponse, error) {
	return api.GetWaterMark(cli, watermarkId)
}

// list user`s watermark by watermarkId
func (cli *Client) ListWaterMark() (*api.ListWaterMarkResponse, error) {
	return api.ListWaterMark(cli)
}

// delete watermark by watermarkId
func (cli *Client) DeleteWaterMark(watermarkId string) error {
	return api.DeleteWaterMark(cli, watermarkId)
}

// create notification with name and endpoint
func (cli *Client) CreateNotification(name, endpoint string) error {
	return api.CreateNotification(cli, name, endpoint)
}

// get notification by notification`name
func (cli *Client) GetNotification(name string) (*api.CreateNotificationArgs, error) {
	return api.GetNotification(cli, name)
}

// list all of user`s notification
func (cli *Client) ListNotification() (*api.ListNotificationsResponse, error) {
	return api.ListNotification(cli)
}

// delete notification by name
func (cli *Client) DeleteNotification(name string) error {
	return api.DeleteNotification(cli, name)
}
