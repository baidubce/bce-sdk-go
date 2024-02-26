package bcm

import (
	"errors"

	"github.com/baidubce/bce-sdk-go/bce"
	"github.com/baidubce/bce-sdk-go/http"
)

// 创建警报配置
func (c *Client) CreateCustomAlarmPolicy(config *CustomAlarmConfig) error {
	if config == nil {
		return errors.New("CustomAlarmConfig is must not empty")
	}
	if err := config.Validate(); err != nil {
		return err
	}
	err := bce.NewRequestBuilder(c).
		WithURL(CreateCustomAlarmConfig).
		WithMethod(http.POST).
		WithBody(config).
		Do()
	return err
}

// 编辑报警策略
func (c *Client) UpdateCustomAlarmPolicy(config *CustomAlarmConfig) error {
	if config == nil {
		return errors.New("CustomAlarmConfig is must not empty")
	}
	if err := config.Validate(); err != nil {
		return err
	}
	err := bce.NewRequestBuilder(c).
		WithURL(UpdateCustomAlarmConfig).
		WithMethod(http.PUT).
		WithBody(config).
		Do()
	return err
}

//  删除警报配置
func (c *Client) DeleteCustomAlarmPolicy(policys *AlarmPolicyBatchList) error {
	if policys == nil {
		return errors.New("AlarmPolicyBatchList and CustomEventAlarmList is must not empty")
	}

	if err := policys.Validate("CustomEventAlarmList"); err != nil {
		return err
	}
	err := bce.NewRequestBuilder(c).
		WithURL(BatchDeleteCustomAlarmPolicy).
		WithMethod(http.POST).
		WithBody(policys).
		Do()
	return err
}

//  列出警报配置
func (c *Client) ListCustomAlarmPolicy(params *ListCustomAlarmPolicyParams) (*ListCustomPolicyPageResultResponse, error) {
	if params == nil {
		return nil, errors.New("ListCustomAlarmPolicyParams is must not empty")
	}
	bceParams := params.ConvertToBceParams()
	var resp ListCustomPolicyPageResultResponse
	err := bce.NewRequestBuilder(c).
		WithURL(ListCustomAlarmPolicy).
		WithResult(&resp).
		WithMethod(http.GET).
		WithQueryParams(bceParams).
		Do()
	return &resp, err
}

// 策略详情
func (c *Client) DetailCustomAlarmConfig(params *DetailCustomAlarmPolicyParams) (*CustomAlarmConfig, error) {
	if params == nil {
		return nil, errors.New("DetailCustomAlarmConfigParams is must not empty")
	}
	bceParams := params.ConvertToBceParams()

	var resp CustomAlarmConfig
	err := bce.NewRequestBuilder(c).
		WithURL(DetailCustomAlarmPolicy).
		WithResult(&resp).
		WithMethod(http.GET).
		WithQueryParams(bceParams).
		Do()

	return &resp, err
}

// 关闭策略
func (c *Client) BlockCustomAlarmConfig(params *BlockCustomAlarmPolicyParams) error {
	if params == nil {
		return errors.New("BlockCustomAlarmConfigParams is must not empty")
	}
	bceParams := params.ConvertToBceParams()

	err := bce.NewRequestBuilder(c).
		WithURL(BlockCustomAlarmPolicy).
		WithMethod(http.POST).
		WithQueryParams(bceParams).
		Do()

	return err
}

//打开策略
func (c *Client) UnblockCustomAlarmConfig(params *UnblockCustomAlarmPolicyParams) error {
	if params == nil {
		return errors.New("UnblockCustomAlarmConfigParams is must not empty")
	}
	bceParams := params.ConvertToBceParams()

	err := bce.NewRequestBuilder(c).
		WithURL(UnblockCustomAlarmPolicy).
		WithMethod(http.POST).
		WithQueryParams(bceParams).
		Do()

	return err
}
