package bvw

import (
	"github.com/baidubce/bce-sdk-go/auth"
	"github.com/baidubce/bce-sdk-go/bce"
	"github.com/baidubce/bce-sdk-go/services/bvw/api"
)

type Client struct {
	*bce.BceClient
}

const DEFAULT_SERVICE_DOMAIN = "bvw.bj.baidubce.com"

// NewClient make BVW service client with defualt configuration
// endPoint value only bj can be chosed and enPoint value defualt bj
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

// upload video /audio/picture to MaterialLibaray
func (cli *Client) UploadMaterial(request *api.MatlibUploadRequest) (*api.MatlibUploadResponse, error) {
	return api.UploadMaterial(cli, request)
}

// search material
func (cli *Client) SearchMaterial(materialSearchRequest *api.MaterialSearchRequest) (*api.MaterialSearchResponse, error) {
	return api.SearchMaterial(cli, materialSearchRequest)
}

// get material by materialId
func (cli *Client) GetMaterial(id string) (*api.MaterialGetResponse, error) {
	return api.GetMaterial(cli, id)
}

// delete material by materialId
func (cli *Client) DeleteMaterial(id string) error {
	return api.DeleteMaterial(cli, id)
}

// upload preset materials to the media library
func (cli *Client) UploadMaterialPreset(fileType string, request *api.MatlibUploadRequest) (
	*api.MaterialPresetUploadResponse, error) {
	return api.UploadMaterialPreset(cli, fileType, request)
}

// search preset materials
func (cli Client) SearchMaterialPreset(request *api.MaterialPresetSearchRequest) (*api.MaterialPresetSearchResponse, error) {
	return api.SearchMaterialPreset(cli, request)
}

// get preset matertials by id
func (cli Client) GetMaterialPreset(id string) (*api.MaterialPresetGetResponse, error) {
	return api.GetMaterialPreset(cli, id)
}

// delete preset matertials by id
func (cli *Client) DeleteMaterialPreset(id string) error {
	return api.DeleteMaterialPreset(cli, id)
}

// create matlib config
func (cli *Client) CreateMatlibConfig(request *api.MatlibConfigBaseRequest) (*bce.BceResponse, error) {
	return api.CreateMatlibConfig(cli, request)
}

// get matlib config
func (cli *Client) GetMatlibConfig() (*api.MatlibConfigGetResponse, error) {
	return api.GetMatlibConfig(cli)
}

// update matlib config
func (cli *Client) UpdateMatlibConfig(request *api.MatlibConfigUpdateRequest) error {
	return api.UpdateMatlibConfig(cli, request)
}

// create edit draft
func (cli *Client) CreateDraft(request *api.CreateDraftRequest) (*api.MatlibTaskResponse, error) {
	return api.CreateDraft(cli, request)
}

// get draft and timeline with matlib task id
func (cli *Client) GetSingleDraft(id int) (*api.GetDraftResponse, error) {
	return api.GetSingleDraft(cli, id)
}

// get draft list
func (cli *Client) GetDraftList(request *api.DraftListRequest) (*api.ListByPageResponse, error) {
	return api.GetDraftList(cli, request)
}

// update draft by id
func (cli *Client) UpdateDraft(id int, request *api.MatlibTaskRequest) error {
	return api.UpdateDraft(cli, id, request)
}

// query video edit job status
func (cli *Client) PollingVideoEdit(id int) (*api.VideoEditPollingResponse, error) {
	return api.PollingVideoEdit(cli, id)
}

// create edit job
func (cli *Client) CreateVideoEdit(request *api.VideoEditCreateRequest) (*api.VideoEditCreateResponse, error) {
	return api.CreateVideoEdit(cli, request)
}
