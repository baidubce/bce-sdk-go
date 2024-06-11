package api

// create pipline args
type CreatePiplineArgs struct {
	PipelineName string               `json:"pipelineName"`
	Description  string               `json:"description,omitempty"`
	SourceBucket string               `json:"sourceBucket"`
	TargetBucket string               `json:"targetBucket"`
	Config       *CreatePiplineConfig `json:"config"`
}

type CreatePiplineConfig struct {
	Capacity     int    `json:"capacity"`
	Notification string `json:"notification,omitempty"`
	PipelineType string `json:"pipelineType,omitempty"`
}

type ListPipelinesResponse struct {
	Pipelines []PipelineStatus `json:"pipelines"`
}

type JobStatus struct {
	Total   int `json:"total,omitempty"`
	Running int `json:"running,omitempty"`
	Pending int `json:"pending,omitempty"`
	Failed  int `json:"failed,omitempty"`
}

type PipelineStatus struct {
	PipelineName string              `json:"pipelineName"`
	Description  string              `json:"description,omitempty"`
	SourceBucket string              `json:"sourceBucket"`
	TargetBucket string              `json:"targetBucket"`
	Config       CreatePiplineConfig `json:"config"`
	State        string              `json:"state,omitempty"`
	Createtime   string              `json:"createtime,omitempty"`
	JobStatus    JobStatus           `json:"jobStatus,omitempty"`
}

type UpdatePipelineArgs struct {
	PipelineName         string                `json:"pipelineName,omitempty"`
	Description          string                `json:"description,omitempty"`
	SourceBucket         string                `json:"sourceBucket,omitempty"`
	TargetBucket         string                `json:"targetBucket,omitempty"`
	UpdatePipelineConfig *UpdatePipelineConfig `json:"config,omitempty"`
}

type UpdatePipelineConfig struct {
	Capacity     int    `json:"capacity,omitempty"`
	Notification string `json:"notification,omitempty"`
}

type CreateJobArgs struct {
	PipelineName string  `json:"pipelineName,omitempty"`
	Source       *Source `json:"source"`
	Target       *Target `json:"target"`
}

type Source struct {
	SourceKey string        `json:"sourceKey,omitempty"`
	Clips     *[]SourceClip `json:"clips,omitempty"`
}

type SourceClip struct {
	Bucket                 string `json:"bucket,omitempty"`
	SourceKey              string `json:"sourceKey,omitempty"`
	StartTimeInSecond      int    `json:"startTimeInSecond,omitempty"`
	DurationInSecond       int    `json:"durationInSecond,omitempty"`
	StartTimeInMillisecond int    `json:"startTimeInMillisecond,omitempty"`
	DurationInMillisecond  int    `json:"durationInMillisecond,omitempty"`
	EnableLogo             bool   `json:"enableLogo,omitempty"`
	AsMasterClip           bool   `json:"asMasterClip,omitempty"`
	EnableDelogo           bool   `json:"enableDelogo,omitempty"`
	EnableCrop             bool   `json:"enableCrop,omitempty"`
}

type Target struct {
	TargetKey            string    `json:"targetKey,omitempty"`
	PresetName           string    `json:"presetName,omitempty"`
	AutoDelogo           bool      `json:"autoDelogo,omitempty"`
	DelogoMode           string    `json:"delogoMode,omitempty"`
	DelogoArea           *Area     `json:"delogoArea,omitempty"`
	DelogoAreas          *[]Area   `json:"delogoAreas,omitempty"`
	AutoCrop             bool      `json:"autoCrop,omitempty"`
	Crop                 *Area     `json:"crop,omitempty"`
	WatermarkIds         []string  `json:"watermarkIds,omitempty"`
	Inserts              *[]Insert `json:"inserts,omitempty"`
	DigitalWmSecretKeyId string    `json:"digitalWmSecretKeyId,omitempty"`
	DigitalWmTextContent string    `json:"digitalWmTextContent,omitempty"`
	JobCfg               *JobCfg   `json:"jobCfg,omitempty"`
}

type JobCfg struct {
	Notification string `json:"notification,omitempty"`
}

type Area struct {
	X      int `json:"x,omitempty"`
	Y      int `json:"y,omitempty"`
	Width  int `json:"width,omitempty"`
	Height int `json:"height,omitempty"`
}

type Insert struct {
	Bucket   string    `json:"bucket,omitempty"`
	Key      string    `json:"key,omitempty"`
	Type     string    `json:"type,omitempty"`
	Text     string    `json:"text"`
	Font     *Font     `json:"font,omitempty"`
	Timeline *Timeline `json:"timeline"`
	Layout   *Layout   `json:"layout,omitempty"`
}

type Font struct {
	Family      string `json:"family,omitempty"`
	SizeInPoint int    `json:"sizeInPoint,omitempty"`
}

type Timeline struct {
	StartTimeInMillisecond int `json:"startTimeInMillisecond,omitempty"`
	DurationInMillisecond  int `json:"durationInMillisecond,omitempty"`
}

type Layout struct {
	VerticalAlignment       string `json:"verticalAlignment,omitempty"`
	HorizontalAlignment     string `json:"horizontalAlignment,omitempty"`
	VerticalOffsetInPixel   int    `json:"verticalOffsetInPixel,omitempty"`
	HorizontalOffsetInPixel int    `json:"horizontalOffsetInPixel,omitempty"`
}

type CreateJobResponse struct {
	JobId string `json:"jobId"`
}

type ListTranscodingJobsResponse struct {
	Jobs []Job `json:"jobs"`
}

type Job struct {
	JobID        string `json:"jobId"`
	PipelineName string `json:"pipelineName"`
	Source       Source `json:"source"`
	Target       Target `json:"target"`
	JobStatus    string `json:"jobStatus"`
	StartTime    string `json:"startTime"`
	EndTime      string `json:"endTime"`
	Error        Error  `json:"error"`
}

type GetTranscodingJobResponse struct {
	JobID        string `json:"jobId"`
	PipelineName string `json:"pipelineName"`
	Source       Source `json:"source"`
	Target       Target `json:"target"`
	JobStatus    string `json:"jobStatus"`
	StartTime    string `json:"startTime"`
	EndTime      string `json:"endTime"`
	Error        Error  `json:"error"`
}

type Error struct {
	Code    string `json:"code"`
	Message string `json:"message"`
}

type ListPresetsResponse struct {
	Presets []Preset `json:"presets"`
}

type Preset struct {
	PresetName         string      `json:"presetName,omitempty"`
	Description        string      `json:"description,omitempty"`
	Container          string      `json:"container,omitempty"`
	Transmux           bool        `json:"transmux,omitempty"`
	Clip               *Clip       `json:"clip,omitempty"`
	Audio              *Audio      `json:"audio,omitempty"`
	Video              *Video      `json:"video,omitempty"`
	Encryption         *Encryption `json:"encryption,omitempty"`
	WatermarkID        string      `json:"watermarkId,omitempty"`
	Watermarks         *Watermarks `json:"watermarks,omitempty"`
	TransCfg           *TransCfg   `json:"transCfg,omitempty"`
	ExtraCfg           *ExtraCfg   `json:"extraCfg,omitempty"`
	State              string      `json:"state,omitempty"`
	CreatedTime        string      `json:"createdTime,omitempty"`
	DigitalWmId        string      `json:"digitalWmId,omitempty"`
	DigitalWmSecretKey string      `json:"digitalWmSecretKey,omitempty"`
}

type Clip struct {
	StartTimeInSecond int `json:"startTimeInSecond,omitempty"`
	DurationInSecond  int `json:"durationInSecond,omitempty"`
}

type Audio struct {
	BitRateInBps   int           `json:"bitRateInBps,omitempty"`
	SampleRateInHz int           `json:"sampleRateInHz,omitempty"`
	Channels       int           `json:"channels,omitempty"`
	PcmFormat      string        `json:"pcmFormat,omitempty"`
	VolumeAdjust   *VolumeAdjust `json:"volumeAdjust,omitempty"`
	Codec          string        `json:"codec,omitemptyc"`
}

type VolumeAdjust struct {
	Mute bool `json:"mute,omitempty"`
	Norm bool `json:"norm,omitempty"`
	Gain int  `json:"gain,omitempty"`
}

type Video struct {
	Codec                string        `json:"codec,omitempty"`
	CodecOptions         *CodecOptions `json:"codecOptions,omitempty"`
	RateControl          string        `json:"rateControl,omitempty"`
	CodecEnhance         bool          `json:"codecEnhance,omitempty"`
	BitRateInBps         int           `json:"bitRateInBps,omitempty"`
	MaxFrameRate         float64       `json:"maxFrameRate,omitempty"`
	MaxWidthInPixel      int           `json:"maxWidthInPixel,omitempty"`
	MaxHeigtInPixel      int           `json:"maxHeightInPixel,omitempty"`
	SizingPolicy         string        `json:"sizingPolicy,omitempty"`
	PlaybackSpeed        float64       `json:"playbackSpeed,omitempty"`
	Crf                  int           `json:"crf,omitempty"`
	AutoAdjustResolution bool          `json:"autoAdjustResolution,omitempty"`
}

type CodecOptions struct {
	Profile string `json:"profile,omitempty"`
}

type Encryption struct {
	Strategy     string `json:"strategy,omitempty"`
	AesKey       string `json:"aesKey,omitempty"`
	KeyServerURL string `json:"keyServerUrl,omitempty"`
}

type Watermarks struct {
	Image []string `json:"image,omitempty"`
}

type TransCfg struct {
	TransMode string `json:"transMode,omitempty"`
}

type ExtraCfg struct {
	WatermarkDisableWhitelist []string `json:"watermarkDisableWhitelist,omitempty"`
	SegmentDurationInSecond   float64  `json:"segmentDurationInSecond,omitempty"`
	GopLength                 int      `json:"gopLength,omitempty"`
	SkipBlackFrame            bool     `json:"skipBlackFrame,omitempty"`
}

type GetPresetResponse struct {
	PresetName  string      `json:"presetName"`
	Description string      `json:"description"`
	Container   string      `json:"container"`
	Transmux    bool        `json:"transmux"`
	Clip        Clip        `json:"clip"`
	Audio       Audio       `json:"audio"`
	Video       *Video      `json:"video"`
	Encryption  *Encryption `json:"encryption"`
	WatermarkID string      `json:"watermarkId"`
	Watermarks  *Watermarks `json:"watermarks"`
	TransCfg    *TransCfg   `json:"transCfg"`
	ExtraCfg    *ExtraCfg   `json:"extraCfg"`
	State       string      `json:"state"`
	CreatedTime string      `json:"createdTime"`
}

type GetMediaInfoOfFileResponse struct {
	Bucket                string     `json:"bucket"`
	Key                   string     `json:"key"`
	FileSizeInByte        int        `json:"fileSizeInByte"`
	Container             string     `json:"container"`
	DurationInSecond      int        `json:"durationInSecond"`
	DurationInMillisecond int        `json:"durationInMillisecond"`
	Etag                  string     `json:"etag"`
	Type                  string     `json:"type"`
	VideoInfo             *VideoInfo `json:"video"`
	AudioInfo             *AudioInfo `json:"audio"`
}

type VideoInfo struct {
	Codec         string  `json:"codec"`
	HeightInPixel int     `json:"heightInPixel"`
	WidthInPixel  int     `json:"widthInPixel"`
	BitRateInBps  int     `json:"bitRateInBps"`
	FrameRate     float64 `json:"frameRate"`
	Rotate        int     `json:"rotate"`
	Dar           string  `json:"dar"`
}

type AudioInfo struct {
	Codec          string `json:"codec"`
	Channels       int    `json:"channels"`
	SampleRateInHz int    `json:"sampleRateInHz"`
	BitRateInBps   int    `json:"bitRateInBps"`
}

type CreateThumbnailJobArgs struct {
	PipelineName     string            `json:"pipelineName,omitempty"`
	ThumbnailSource  *ThumbnailSource  `json:"source"`
	PresetName       string            `json:"presetName,omitempty"`
	ThumbnailTarget  *ThumbnailTarget  `json:"target,omitempty"`
	ThumbnailCapture *ThumbnailCapture `json:"capture,omitempty"`
	Area             *Area             `json:"delogoArea,omitempty"`
	Crop             *Area             `json:"crop,omitempty"`
}

type ThumbnailSource struct {
	Key string `json:"key,omitempty"`
}

type ThumbnailCapture struct {
	Mode                string              `json:"mode,omitempty"`
	StartTimeInSecond   float64             `json:"startTimeInSecond,omitempty"`
	EndTimeInSecond     float64             `json:"endTimeInSecond,omitempty"`
	IntervalInSecond    float64             `json:"intervalInSecond,omitempty"`
	MinIntervalInSecond float64             `json:"minIntervalInSecond,omitempty"`
	FrameNumber         int                 `json:"frameNumber,omitempty"`
	SkipBlackFrame      bool                `json:"skipBlackFrame,omitempty"`
	HighlightOutputCfg  *HighlightOutputCfg `json:"highlightOutputCfg,omitempty"`
	SpriteOutputCfg     *SpriteOutputCfg    `json:"spriteOutputCfg,omitempty"`
}

type ThumbnailTarget struct {
	KeyPrefix       string           `json:"keyPrefix,omitempty"`
	Format          string           `json:"format,omitempty"`
	FrameRate       float64          `json:"frameRate,omitempty"`
	GifQuality      string           `json:"gifQuality,omitempty"`
	SizingPolicy    string           `json:"sizingPolicy,omitempty"`
	WidthInPixel    int              `json:"widthInPixel,omitempty"`
	HeightInPixel   int              `json:"heightInPixel,omitempty"`
	SpriteOutputCfg *SpriteOutputCfg `json:"spriteOutputCfg,omitempty"`
}

type HighlightOutputCfg struct {
	DurationInSecond float64 `json:"durationInSecond,omitempty"`
	FrameRate        float64 `json:"frameRate,omitempty"`
	PlaybackSpeed    float64 `json:"playbackSpeed,omitempty"`
	ReverseConcat    bool    `json:"reverseConcat,omitempty"`
}

type SpriteOutputCfg struct {
	Rows         int    `json:"rows,omitempty"`
	Columns      int    `json:"columns,omitempty"`
	Margin       int    `json:"margin,omitempty"`
	Padding      int    `json:"padding,omitempty"`
	KeepCellPic  bool   `json:"keepCellPic,omitempty"`
	SpriteKeyTag string `json:"spriteKeyTag,omitempty"`
}

type ThumbnailOptional struct {
	PresetName      string
	Target          *ThumbnailTarget
	Capture         *ThumbnailCapture
	DelogoArea      *Area
	Crop            *Area
	ThumbnailSource *ThumbnailSource
}

type GetThumbnailJobResponse struct {
	JobID        string                 `json:"jobId,omitempty"`
	JobStatus    string                 `json:"jobStatus,omitempty"`
	PipelineName string                 `json:"pipelineName,omitempty"`
	Source       *ThumbnailSource       `json:"source,omitempty"`
	PresetName   string                 `json:"presetName,omitempty"`
	Target       *ThumbnailTargetStatus `json:"target,omitempty"`
	Capture      *ThumbnailCapture      `json:"capture,omitempty"`
	DelogoArea   *Area                  `json:"delogoArea,omitempty"`
	Error        *Error                 `json:"error,omitempty"`
}

type ThumbnailTargetStatus struct {
	KeyPrefix       string           `json:"keyPrefix,omitempty"`
	Format          string           `json:"format,omitempty"`
	FrameRate       float64          `json:"frameRate,omitempty"`
	GifQuality      string           `json:"gifQuality,omitempty"`
	SizingPolicy    string           `json:"sizingPolicy,omitempty"`
	WidthInPixel    int              `json:"widthInPixel,omitempty"`
	HeightInPixel   int              `json:"heightInPixel,omitempty"`
	SpriteOutputCfg *SpriteOutputCfg `json:"spriteOutputCfg,omitempty"`
	Keys            []string         `json:"keys,omitempty"`
}

type ListThumbnailJobsResponse struct {
	Thumbnails []ThumbnailJobStatus `json:"thumbnails"`
}

type ThumbnailJobStatus struct {
	JobID        string                 `json:"jobId"`
	JobStatus    string                 `json:"jobStatus"`
	PipelineName string                 `json:"pipelineName"`
	Source       *ThumbnailSource       `json:"source,omitempty"`
	Target       *ThumbnailTargetStatus `json:"target,omitempty"`
	Capture      *Area                  `json:"capture,omitempty"`
	DelogoArea   *Area                  `json:"delogoArea,omitempty"`
	Error        *Error                 `json:"error,omitempty"`
}

type CreateWaterMarkArgs struct {
	Bucket                  string    `json:"bucket,omitempty"`
	Key                     string    `json:"key,omitempty"`
	VerticalAlignment       string    `json:"verticalAlignment,omitempty"`
	HorizontalAlignment     string    `json:"horizontalAlignment,omitempty"`
	VerticalOffsetInPixel   int       `json:"verticalOffsetInPixel,omitempty"`
	HorizontalOffsetInPixel int       `json:"horizontalOffsetInPixel,omitempty"`
	Timeline                *Timeline `json:"timeline,omitempty"`
	Repeated                int       `json:"repeated,omitempty"`
	AllowScaling            bool      `json:"allowScaling,omitempty"`
	Dx                      string    `json:"dx,omitempty"`
	Dy                      string    `json:"dy,omitempty"`
	Width                   string    `json:"width,omitempty"`
	Height                  string    `json:"height,omitempty"`
}

type CreateWaterMarkResponse struct {
	WatermarkId string `json:"watermarkId"`
}

type GetWaterMarkResponse struct {
	Bucket                  string    `json:"bucket"`
	Key                     string    `json:"key"`
	VerticalAlignment       string    `json:"verticalAlignment"`
	HorizontalAlignment     string    `json:"horizontalAlignment"`
	VerticalOffsetInPixel   int       `json:"verticalOffsetInPixel"`
	HorizontalOffsetInPixel int       `json:"horizontalOffsetInPixel"`
	Timeline                *Timeline `json:"timeline"`
	Repeated                int       `json:"repeated"`
	AllowScaling            bool      `json:"allowScaling"`
	Dx                      string    `json:"dx"`
	Dy                      string    `json:"dy"`
	Width                   string    `json:"width"`
	Height                  string    `json:"height"`
}

type ListWaterMarkResponse struct {
	Watermarks []Watermark `json:"watermarks"`
}

type Watermark struct {
	Bucket                  string    `json:"bucket"`
	Key                     string    `json:"key"`
	VerticalOffsetInPixel   int       `json:"verticalOffsetInPixel"`
	HorizontalOffsetInPixel int       `json:"horizontalOffsetInPixel"`
	WatermarkID             string    `json:"watermarkId"`
	CreateTime              string    `json:"createTime"`
	VerticalAlignment       string    `json:"verticalAlignment"`
	HorizontalAlignment     string    `json:"horizontalAlignment"`
	Dx                      string    `json:"dx"`
	Dy                      string    `json:"dy"`
	Width                   string    `json:"width"`
	Height                  string    `json:"height"`
	Timeline                *Timeline `json:"timeline,omitempty"`
	Repeated                int       `json:"repeated"`
	AllowScaling            bool      `json:"allowScaling"`
}

type CreateNotificationArgs struct {
	Name     string `json:"name,omitempty"`
	Endpoint string `json:"endpoint,omitempty"`
}

type ListNotificationsResponse struct {
	Notifications []CreateNotificationArgs `json:"notifications"`
}

type DigitalWmPreset struct {
	DigitalWmId   string `json:"digitalWmId,omitempty"`
	Description   string `json:"description,omitempty"`
	DigitalWmType string `json:"digitalWmType,omitempty"`
	Bucket        string `json:"bucket,omitempty"`
	Key           string `json:"key,omitempty"`
	TextContent   string `json:"textContent,omitempty"`
}

type CreateDigitalWmPresetResponse struct {
	WatermarkId string `json:"watermarkId,omitempty"`
}

type GetDigitalWmPresetResponse struct {
	DigitalWmId   string `json:"digitalWmId,omitempty"`
	Description   string `json:"description,omitempty"`
	CreateTime    string `json:"createTime,omitempty"`
	State         string `json:"state,omitempty"`
	DigitalWmType string `json:"digitalWmType,omitempty"`
	Bucket        string `json:"bucket,omitempty"`
	Key           string `json:"key,omitempty"`
	TextContent   string `json:"textContent,omitempty"`
}

type ListDigitalWmPresetResponse struct {
	DigitalWatermarks []GetDigitalWmPresetResponse `json:"digitalWatermarks,omitempty"`
}

type DwmSecretkeyPreset struct {
	DigitalWmSecretKeyId string `json:"digitalWmSecretKeyId,omitempty"`
	Description          string `json:"description,omitempty"`
	SecretKey            string `json:"secretKey,omitempty"`
}

type CreateDwmSecretkeyPresetResponse struct {
	DigitalWmSecretKeyId string `json:"digitalWmSecretKeyId,omitempty"`
}

type GetDwmSecretkeyPresetResponse struct {
	DigitalWmSecretKeyId string `json:"digitalWmSecretKeyId,omitempty"`
	Description          string `json:"description,omitempty"`
	CreateTime           string `json:"createTime,omitempty"`
	State                string `json:"state,omitempty"`
	SecretKey            string `json:"secretKey,omitempty"`
}

type ListDwmPresetSecretkeyResponse struct {
	DwmPresetSecretkeys []GetDwmSecretkeyPresetResponse `json:"secretKeys,omitempty"`
}

type Dwmdetect struct {
	PipelineName         string           `json:"pipelineName,omitempty"`
	Source               *DwmSource       `json:"source,omitempty"`
	Target               *DwmTarget       `json:"target,omitempty"`
	DigitalWmType        string           `json:"digitalWmType,omitempty"`
	DigitalWmSecretKeyId string           `json:"digitalWmSecretKeyId,omitempty"`
	DigitalWmId          string           `json:"digitalWmId,omitempty"`
	TextWmContent        string           `json:"textWmContent,omitempty"`
	RefResolutions       []RefResolutions `json:"refResolutions,omitempty"`
}

type DwmSource struct {
	Bucket string `json:"bucket,omitempty"`
	Key    string `json:"key,omitempty"`
	Url    string `json:"url,omitempty"`
}

type DwmTarget struct {
	Bucket    string   `json:"bucket,omitempty"`
	Key       string   `json:"key,omitempty"`
	Keys      []string `json:"keys,omitempty"`
	KeyPrefix string   `json:"keyPrefix,omitempty"`
}

type RefResolutions struct {
	OriginalVideoWidth  int `json:"originalVideoWidth,omitempty"`
	OriginalVideoHeight int `json:"originalVideoHeight,omitempty"`
}

type GetDwmdetectResponse struct {
	Dwmdetect
	JobId            string   `json:"jobId,omitempty"`
	JobStatus        string   `json:"jobStatus,omitempty"`
	CreateTime       string   `json:"createTime,omitempty"`
	StartTime        string   `json:"startTime,omitempty"`
	EndTime          string   `json:"endTime,omitempty"`
	DetectFramesNum  int      `json:"detectFramesNum,omitempty"`
	DetectedTexts    []string `json:"detectedTexts,omitempty"`
	DetectSuccessNum int      `json:"detectSuccessNum,omitempty"`
}

type Imagedwm struct {
	PipelineName     string     `json:"pipelineName,omitempty"`
	Source           *DwmSource `json:"source,omitempty"`
	Target           *DwmTarget `json:"target,omitempty"`
	TaskType         string     `json:"taskType,omitempty"`
	Strength         float64    `json:"strength,omitempty"`
	DigitalWm        *DigitalWm `json:"digitalWm,omitempty"`
	AlgorithmVersion int        `json:"algorithmVersion,omitempty"`
}

type DigitalWm struct {
	ImageBucket string `json:"imageBucket,omitempty"`
	ImageKey    string `json:"imageKey,omitempty"`
	ImageUrl    string `json:"imageUrl,omitempty"`
	TextContent string `json:"textContent,omitempty"`
}

type GetImagedwmResponse struct {
	JobId string `json:"jobId,omitempty"`
	Imagedwm
	JobStatus  string   `json:"jobStatus,omitempty"`
	CreateTime string   `json:"createTime,omitempty"`
	StartTime  string   `json:"startTime,omitempty"`
	EndTime    string   `json:"endTime,omitempty"`
	Error      JobError `json:"error,omitempty"`
}

type JobError struct {
	Code    string `json:"code,omitempty"`
	Message string `json:"message,omitempty"`
}
