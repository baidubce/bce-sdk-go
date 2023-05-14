package api

type MatlibUploadRequest struct {
	MediaType    string `json:"mediaType"`
	Title        string `json:"title"`
	Bucket       string `json:"bucket"`
	Key          string `json:"key"`
	Notification string `json:"notification"`
}

type MatlibUploadResponse struct {
	Id string `json:"materialId"`
}

type MaterialSearchRequest struct {
	InfoType     string `json:"infoType"`
	MediaType    string `json:"mediaType"`
	SourceType   string `json:"sourceType"`
	Status       string `json:"status"`
	TitleKeyword string `json:"titleKeyword"`
	PageNo       int    `json:"pageNo"`
	Size         int    `json:"size"`
	Begin        string `json:"begin"`
	End          string `json:"end"`
}

type MaterialSearchResponse struct {
	Items []MaterialGetResponse `json:"items,omitempty"`
}

type MaterialGetResponse struct {
	ID            string   `json:"id"`
	UserID        string   `json:"userId"`
	InfoType      string   `json:"infoType,omitempty"`
	MediaType     string   `json:"mediaType,omitempty"`
	SourceType    string   `json:"sourceType"`
	Status        string   `json:"status"`
	Title         string   `json:"title"`
	SourceURL     string   `json:"sourceUrl"`
	SourceURL360P string   `json:"sourceUrl360p"`
	ThumbnailList []string `json:"thumbnailList"`
	SubtitleUrls  []string `json:"subtitleUrls"`
	CreateTime    string   `json:"createTime"`
	UpdateTime    string   `json:"updateTime"`
	Duration      float64  `json:"duration"`
	Height        int      `json:"height"`
	Width         int      `json:"width"`
	Bucket        string   `json:"bucket"`
	Key           string   `json:"key"`
	Key360P       string   `json:"key360p"`
	Key720P       string   `json:"key720p"`
	AudioKey      string   `json:"audioKey"`
	ThumbnailKeys []string `json:"thumbnailKeys"`
	Subtitles     []string `json:"subtitles"`
}

type MaterialPresetUploadResponse struct {
	Id string `json:"id"`
}

type MaterialPresetSearchRequest struct {
	SourceType string `json:"sourceType"`
	Status     string `json:"status"`
	MediaType  string `json:"type"`
	PageNo     string `json:"pageNo"`
	PageSize   string `json:"pageSize"`
}

type MaterialPresetSearchResponse struct {
	Result     []PresetResponseWrapper `json:"result"`
	PageNo     int                     `json:"pageNo"`
	PageSize   int                     `json:"pageSize"`
	TotalCount int                     `json:"totalCount"`
}

type PresetResponseWrapper struct {
	MaterialType string                      `json:"type"`
	Addons       []MaterialPresetGetResponse `json:"addons"`
}

type MaterialPresetGetResponse struct {
	ID                 string                 `json:"id"`
	Status             string                 `json:"status"`
	UserID             string                 `json:"userId"`
	Title              string                 `json:"title"`
	Tag                string                 `json:"tag"`
	Type               string                 `json:"type"`
	SourceType         string                 `json:"sourceType"`
	PreviewMaterialIds map[string]interface{} `json:"previewMaterialIds"`
	PreviewBucket      map[string]interface{} `json:"previewBucket"`
	PreviewKeys        map[string]interface{} `json:"previewKeys"`
	PreviewUrls        map[string]interface{} `json:"previewUrls"`
	MaterialID         string                 `json:"materialId"`
	Bucket             string                 `json:"bucket"`
	Key                string                 `json:"key"`
	SourceURL          string                 `json:"sourceUrl"`
	Config             string                 `json:"config"`
	CreateTime         string                 `json:"createTime"`
	UpdateTime         string                 `json:"updateTime"`
}

type MatlibConfigGetResponse struct {
	Bucket string `json:"bucket"`
}

type MatlibConfigBaseRequest struct {
	Bucket string `json:"bucket"`
}

type MatlibConfigUpdateRequest struct {
	Bucket string `json:"bucket"`
}

type CreateDraftRequest struct {
	Titile   string `json:"title"`
	Ratio    string `json:"ratio"`
	Duration string `json:"duration"`
}

type MatlibTaskResponse struct {
	Id int `json:"id"`
}

type Otimeline struct {
	Itimeline Itimeline   `json:"timeline"`
	Meta      interface{} `json:"meta"`
}

type Itimeline struct {
	Video    []MediaPair   `json:"video"`
	Audio    []MediaPair   `json:"audio"`
	Sticker  []MediaPair   `json:"sticker"`
	Subtitle []interface{} `json:"subtitle"`
}

type MatlibTaskRequest struct {
	Title          string                `json:"title"`
	Ratio          string                `json:"ratio"`
	Duration       string                `json:"duration"`
	Timeline       Otimeline             `json:"timeline"`
	ResourceList   []MaterialGetResponse `json:"resourceList"`
	CoverBucekt    string                `json:"coverBucket"`
	CoverKey       string                `json:"coverKey"`
	LastUpdateTime string                `json:"lastUpdateTime"`
}

type GetDraftResponse struct {
	Otimeline      Otimeline             `json:"timeline"`
	ResourceList   []GetMaterialResponse `json:"resourceList,omitempty"`
	Title          string                `json:"title"`
	Ratio          string                `json:"ratio"`
	CoverBucket    string                `json:"coverBucket"`
	CoverKey       string                `json:"coverKey"`
	Endpoint       string                `json:"endpoint"`
	LastUpdateTime string                `json:"lastUpdateTime"`
}
type GetMaterialResponse struct {
	ID             string   `json:"id,omitempty"`
	UserID         string   `json:"userId,omitempty"`
	ActualUserID   string   `json:"actualUserId,omitempty"`
	SaasType       string   `json:"saasType,omitempty"`
	InfoType       string   `json:"infoType,omitempty"`
	MediaType      string   `json:"mediaType,omitempty"`
	SourceType     string   `json:"sourceType,omitempty"`
	Status         string   `json:"status,omitempty"`
	Title          string   `json:"title,omitempty"`
	SourceURL      string   `json:"sourceUrl,omitempty"`
	SourceURL360P  string   `json:"sourceUrl360p,omitempty"`
	AudioURL       string   `json:"audioUrl,omitempty"`
	ThumbnailList  []string `json:"thumbnailList,omitempty"`
	SubtitleUrls   []string `json:"subtitleUrls,omitempty"`
	CreateTime     string   `json:"createTime,omitempty"`
	UpdateTime     string   `json:"updateTime,omitempty"`
	Duration       float64  `json:"duration,omitempty"`
	Height         int      `json:"height,omitempty"`
	Width          int      `json:"width,omitempty"`
	FileSizeInByte int      `json:"fileSizeInByte,omitempty"`
	Bucket         string   `json:"bucket,omitempty"`
	Key            string   `json:"key,omitempty"`
	Key360P        string   `json:"key360p,omitempty"`
	Key720P        string   `json:"key720p,omitempty"`
	AudioKey       string   `json:"audioKey,omitempty"`
	ThumbnailKeys  []string `json:"thumbnailKeys,omitempty"`
	Subtitles      []string `json:"subtitles,omitempty"`
	Endpoint       string   `json:"endpoint,omitempty"`
}

type MediaPair struct {
	Key      string             `json:"key"`
	List     []TimelineMaterial `json:"list"`
	IsMaster bool               `json:"isMaster"`
}

type MediaInfo struct {
	FileType        string   `json:"fileType,omitempty"`
	SourceType      string   `json:"sourceType,omitempty"`
	SourceURL       string   `json:"sourceUrl,omitempty"`
	AudioURL        string   `json:"audioUrl,omitempty"`
	Bucket          string   `json:"bucket,omitempty"`
	Key             string   `json:"key,omitempty"`
	AudioKey        string   `json:"audioKey,omitempty"`
	InstanceID      string   `json:"instanceId,omitempty"`
	CoverImage      string   `json:"coverImage,omitempty"`
	Duration        float64  `json:"duration,omitempty"`
	Width           int      `json:"width,omitempty"`
	Height          int      `json:"height,omitempty"`
	Errmsg          string   `json:"errmsg,omitempty"`
	Status          string   `json:"status,omitempty"`
	Progress        string   `json:"progress,omitempty"`
	Action          string   `json:"action,omitempty"`
	Size            int      `json:"size,omitempty"`
	Name            string   `json:"name,omitempty"`
	ThumbnailPrefix string   `json:"thumbnailPrefix,omitempty"`
	ThumbnailKeys   []string `json:"thumbnailKeys,omitempty"`
	ThumbnailList   []string `json:"thumbnailList,omitempty"`
	SubtitleKeys    []string `json:"subtitleKeys,omitempty"`
	MediaID         string   `json:"mediaId,omitempty"`
	Offstandard     bool     `json:"offstandard,omitempty"`
}

type TimelineMaterial struct {
	Start     float64     `json:"start"`
	End       float64     `json:"end"`
	ShowStart float64     `json:"showStart"`
	ShowEnd   float64     `json:"showEnd"`
	Duration  float64     `json:"duration"`
	Xpos      float64     `json:"xpos"`
	Ypos      float64     `json:"ypos"`
	Width     float64     `json:"width"`
	Height    float64     `json:"height"`
	MediaInfo MediaInfo   `json:"mediaInfo"`
	Type      string      `json:"type"`
	UID       string      `json:"uid"`
	Name      string      `json:"name"`
	ExtInfo   interface{} `json:"extInfo"`
}

type ListByPageResponse struct {
	Data       []MatlibTaskGetResponse `json:"data"`
	PageNo     int                     `json:"pageNo"`
	PageSize   int                     `json:"pageSize"`
	TotalCount int                     `json:"totalCount"`
}

type MatlibTaskGetResponse struct {
	ID             int    `json:"id"`
	ResMaterialID  string `json:"resMaterialId"`
	UserID         string `json:"userId"`
	Title          string `json:"title"`
	Status         string `json:"status"`
	ErrorMessage   string `json:"errorMessage"`
	CoverURL       string `json:"coverUrl"`
	LastUpdateTime string `json:"lastUpdateTime"`
}

type DraftListRequest struct {
	Status    string `json:"status"`
	BeginTime string `json:"beginTime"`
	EndTime   string `json:"endTime"`
	PageNo    int    `json:"pageNo"`
	PageSize  int    `json:"pageSize"`
}

type VideoEditPollingResponse struct {
	Errorno int    `json:"errorno"`
	Errmsg  string `json:"errmsg"`
	Data    Data   `json:"data"`
}

type Data struct {
	EditID     int    `json:"editId"`
	EditStatus string `json:"editStatus"`
	Bucket     string `json:"bucket"`
	Key        string `json:"key"`
	CreateTime string `json:"createTime"`
	UpdateTime string `json:"updateTime"`
}

type VideoEditCreateRequest struct {
	Title        string      `json:"title"`
	Bucket       string      `json:"bucket"`
	KeyPath      string      `json:"keyPath"`
	TaskId       string      `json:"taskId"`
	Notification string      `json:"notification"`
	ExtInfo      interface{} `json:"extInfo"`
	Cmd          interface{} `json:"cmd"`
	Endpoint     string      `json:"endpoint"`
}

type VideoEditCreateResponse struct {
	Errorno int    `json:"errorno"`
	Errmsg  string `json:"errmsg"`
	Data    string `json:"data"`
	Ret     Ret    `json:"ret"`
	Ie      string `json:"ie"`
}
type Ret struct {
	EditID int    `json:"editId"`
	BosKey string `json:"bosKey"`
}
