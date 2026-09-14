package api

import (
	"errors"
	nethttp "net/http"

	"github.com/baidubce/bce-sdk-go/bce"
)

// CreateCluster creates a BES cluster.
func CreateCluster(cli bce.Client, region string, request *CreateClusterRequest) (*CreateClusterResponse, error) {
	if cli == nil {
		return nil, ErrNilClient
	}
	if request == nil {
		return nil, errors.New("create cluster request should not be nil")
	}
	if request.Name == "" {
		return nil, errors.New("create cluster request name should not be empty")
	}
	if request.Password == "" {
		return nil, errors.New("create cluster request password should not be empty")
	}
	if request.SecurityGroupId == "" {
		return nil, errors.New("create cluster request securityGroupId should not be empty")
	}
	if request.SubnetUuid == "" {
		return nil, errors.New("create cluster request subnetUuid should not be empty")
	}
	if request.AvailableZone == "" {
		return nil, errors.New("create cluster request availableZone should not be empty")
	}
	if request.VpcId == "" {
		return nil, errors.New("create cluster request vpcId should not be empty")
	}
	if request.Version == "" {
		return nil, errors.New("create cluster request version should not be empty")
	}
	if len(request.Modules) == 0 {
		return nil, errors.New("create cluster request modules should not be empty")
	}
	for _, module := range request.Modules {
		if module.Type == "" {
			return nil, errors.New("create cluster request modules.type should not be empty")
		}
		if module.SlotType == "" {
			return nil, errors.New("create cluster request modules.slotType should not be empty")
		}
		if module.InstanceNum == 0 {
			return nil, errors.New("create cluster request modules.instanceNum should not be empty")
		}
	}
	if request.Billing == nil {
		return nil, errors.New("create cluster request billing should not be nil")
	}
	if request.Billing.PaymentType == "" {
		return nil, errors.New("create cluster request billing.paymentType should not be empty")
	}

	uri := URI_PREFIX_V2 + "/create"

	result := &CreateClusterResponse{}
	if err := createJSONRequest(cli, region, nethttp.MethodPost, uri, request, result); err != nil {
		return nil, err
	}
	return result, nil
}

// ListClusters lists BES clusters.
func ListClusters(cli bce.Client, region string, request *ListClustersRequest) (*ListClustersResponse, error) {
	if cli == nil {
		return nil, ErrNilClient
	}
	if request == nil {
		return nil, errors.New("list clusters request should not be nil")
	}
	if request.PageNo == 0 {
		return nil, errors.New("list clusters request pageNo should not be nil")
	}
	if request.PageSize == 0 {
		return nil, errors.New("list clusters request pageSize should not be nil")
	}

	uri := URI_V2 + "/list"

	result := &ListClustersResponse{}
	if err := createJSONRequest(cli, region, nethttp.MethodPost, uri, request, result); err != nil {
		return nil, err
	}
	return result, nil
}

// GetClusterDetail gets the detail of a BES cluster.
func GetClusterDetail(cli bce.Client, region string, request *GetClusterDetailRequest) (*GetClusterDetailResponse, error) {
	if cli == nil {
		return nil, ErrNilClient
	}
	if request == nil {
		return nil, errors.New("get cluster detail request should not be nil")
	}
	if request.ClusterId == "" {
		return nil, errors.New("get cluster detail request clusterId should not be empty")
	}

	uri := URI_V2 + "/detail"

	result := &GetClusterDetailResponse{}
	if err := createJSONRequest(cli, region, nethttp.MethodPost, uri, request, result); err != nil {
		return nil, err
	}
	return result, nil
}

// DeleteCluster deletes a BES cluster.
func DeleteCluster(cli bce.Client, region string, request *DeleteClusterRequest) (*DeleteClusterResponse, error) {
	if cli == nil {
		return nil, ErrNilClient
	}
	if request == nil {
		return nil, errors.New("delete cluster request should not be nil")
	}
	if request.ClusterId == "" {
		return nil, errors.New("delete cluster request clusterId should not be empty")
	}

	uri := URI_V2 + "/delete"

	result := &DeleteClusterResponse{}
	if err := createJSONRequest(cli, region, nethttp.MethodPost, uri, request, result); err != nil {
		return nil, err
	}
	return result, nil
}

// StartCluster starts a BES cluster.
func StartCluster(cli bce.Client, region string, request *StartClusterRequest) (*StartClusterResponse, error) {
	if cli == nil {
		return nil, ErrNilClient
	}
	if request == nil {
		return nil, errors.New("start cluster request should not be nil")
	}
	if request.ClusterId == "" {
		return nil, errors.New("start cluster request clusterId should not be empty")
	}

	uri := URI_PREFIX_V2 + "/start"

	result := &StartClusterResponse{}
	if err := createJSONRequest(cli, region, nethttp.MethodPost, uri, request, result); err != nil {
		return nil, err
	}
	return result, nil
}

// StopCluster stops a BES cluster.
func StopCluster(cli bce.Client, region string, request *StopClusterRequest) (*StopClusterResponse, error) {
	if cli == nil {
		return nil, ErrNilClient
	}
	if request == nil {
		return nil, errors.New("stop cluster request should not be nil")
	}
	if request.ClusterId == "" {
		return nil, errors.New("stop cluster request clusterId should not be empty")
	}

	uri := URI_PREFIX_V2 + "/stop"

	result := &StopClusterResponse{}
	if err := createJSONRequest(cli, region, nethttp.MethodPost, uri, request, result); err != nil {
		return nil, err
	}
	return result, nil
}

// RestartCluster restarts a BES cluster.
func RestartCluster(cli bce.Client, region string, request *RestartClusterRequest) (*RestartClusterResponse, error) {
	if cli == nil {
		return nil, ErrNilClient
	}
	if request == nil {
		return nil, errors.New("restart cluster request should not be nil")
	}
	if request.ClusterId == "" {
		return nil, errors.New("restart cluster request clusterId should not be empty")
	}

	uri := URI_PREFIX_V2 + "/restart"

	result := &RestartClusterResponse{}
	if err := createJSONRequest(cli, region, nethttp.MethodPost, uri, request, result); err != nil {
		return nil, err
	}
	return result, nil
}

// ResizeCluster resizes a BES cluster.
func ResizeCluster(cli bce.Client, region string, request *ResizeClusterRequest) (*ResizeClusterResponse, error) {
	if cli == nil {
		return nil, ErrNilClient
	}
	if request == nil {
		return nil, errors.New("resize cluster request should not be nil")
	}
	if request.ClusterId == "" {
		return nil, errors.New("resize cluster request clusterId should not be empty")
	}
	if len(request.Modules) == 0 {
		return nil, errors.New("resize cluster request modules should not be empty")
	}
	if request.PaymentType == "" {
		return nil, errors.New("resize cluster request paymentType should not be empty")
	}
	if request.ResizeMode == "" {
		return nil, errors.New("resize cluster request resizeMode should not be empty")
	}

	uri := URI_PREFIX_V2 + "/resize"

	result := &ResizeClusterResponse{}
	if err := createJSONRequest(cli, region, nethttp.MethodPost, uri, request, result); err != nil {
		return nil, err
	}
	return result, nil
}

// AddClusterModule adds a new node type to a BES cluster.
func AddClusterModule(cli bce.Client, region string, request *AddClusterModuleRequest) (*AddClusterModuleResponse, error) {
	if cli == nil {
		return nil, ErrNilClient
	}
	if request == nil {
		return nil, errors.New("add cluster module request should not be nil")
	}
	if request.ClusterId == "" {
		return nil, errors.New("add cluster module request clusterId should not be empty")
	}
	if len(request.Modules) == 0 {
		return nil, errors.New("add cluster module request modules should not be empty")
	}
	if request.ResizeMode == "" {
		return nil, errors.New("add cluster module request resizeMode should not be empty")
	}
	if request.PaymentType == "" {
		return nil, errors.New("add cluster module request paymentType should not be empty")
	}

	uri := URI_PREFIX_V2 + "/addModule"

	result := &AddClusterModuleResponse{}
	if err := createJSONRequest(cli, region, nethttp.MethodPost, uri, request, result); err != nil {
		return nil, err
	}
	return result, nil
}

// ResetClusterPassword resets the admin password of a BES cluster.
func ResetClusterPassword(cli bce.Client, region string, request *ResetClusterPasswordRequest) (*ResetClusterPasswordResponse, error) {
	if cli == nil {
		return nil, ErrNilClient
	}
	if request == nil {
		return nil, errors.New("reset cluster password request should not be nil")
	}
	if request.ClusterId == "" {
		return nil, errors.New("reset cluster password request clusterId should not be empty")
	}
	if request.NewPassword == "" {
		return nil, errors.New("reset cluster password request newPassword should not be empty")
	}
	if request.ConfirmPassword == "" {
		return nil, errors.New("reset cluster password request confirmPassword should not be empty")
	}

	uri := URI_PREFIX_V2 + "/password/reset"

	result := &ResetClusterPasswordResponse{}
	if err := createJSONRequest(cli, region, nethttp.MethodPost, uri, request, result); err != nil {
		return nil, err
	}
	return result, nil
}

// ToggleClusterHTTPS enables or disables HTTPS for a BES cluster.
func ToggleClusterHTTPS(cli bce.Client, region string, request *ToggleClusterHTTPSRequest) (*ToggleClusterHTTPSResponse, error) {
	if cli == nil {
		return nil, ErrNilClient
	}
	if request == nil {
		return nil, errors.New("toggle cluster https request should not be nil")
	}
	if request.ClusterId == "" {
		return nil, errors.New("toggle cluster https request clusterId should not be empty")
	}

	uri := URI_V2 + "/esProtocolTransform"

	result := &ToggleClusterHTTPSResponse{}
	if err := createJSONRequest(cli, region, nethttp.MethodPost, uri, request, result); err != nil {
		return nil, err
	}
	return result, nil
}

// BindClusterEIP binds an EIP to a cluster module.
func BindClusterEIP(cli bce.Client, region string, request *BindClusterEIPRequest) (*BindClusterEIPResponse, error) {
	if cli == nil {
		return nil, ErrNilClient
	}
	if request == nil {
		return nil, errors.New("bind cluster eip request should not be nil")
	}
	if request.InstanceId == "" {
		return nil, errors.New("bind cluster eip request instanceId should not be empty")
	}
	if request.InstanceType == "" {
		return nil, errors.New("bind cluster eip request instanceType should not be empty")
	}
	if request.Eip == "" {
		return nil, errors.New("bind cluster eip request eip should not be empty")
	}

	uri := URI_PREFIX_V2 + "/eip/bind"

	result := &BindClusterEIPResponse{}
	if err := createJSONRequest(cli, region, nethttp.MethodPost, uri, request, result); err != nil {
		return nil, err
	}
	return result, nil
}

// UnbindClusterEIP unbinds an EIP from a cluster module.
func UnbindClusterEIP(cli bce.Client, region string, request *UnbindClusterEIPRequest) (*UnbindClusterEIPResponse, error) {
	if cli == nil {
		return nil, ErrNilClient
	}
	if request == nil {
		return nil, errors.New("unbind cluster eip request should not be nil")
	}
	if request.DeployId == "" {
		return nil, errors.New("unbind cluster eip request deployId should not be empty")
	}
	if request.ModuleTemplateName == "" {
		return nil, errors.New("unbind cluster eip request moduleTemplateName should not be empty")
	}

	uri := URI_PREFIX_V2 + "/eip/unbind"

	result := &UnbindClusterEIPResponse{}
	if err := createJSONRequest(cli, region, nethttp.MethodPost, uri, request, result); err != nil {
		return nil, err
	}
	return result, nil
}

// ToggleClusterMonitor enables or disables Grafana monitor for a BES cluster.
func ToggleClusterMonitor(cli bce.Client, region string, request *ToggleClusterMonitorRequest) (*ToggleClusterMonitorResponse, error) {
	if cli == nil {
		return nil, ErrNilClient
	}
	if request == nil {
		return nil, errors.New("toggle cluster monitor request should not be nil")
	}
	if request.ClusterId == "" {
		return nil, errors.New("toggle cluster monitor request clusterId should not be empty")
	}
	if request.EnableMonitor == nil {
		return nil, errors.New("toggle cluster monitor request enableMonitor should not be nil")
	}

	uri := URI_PREFIX_V2 + "/monitor_state"

	result := &ToggleClusterMonitorResponse{}
	if err := createJSONRequest(cli, region, nethttp.MethodPost, uri, request, result); err != nil {
		return nil, err
	}
	return result, nil
}

// GetClusterTasks gets the operation history of a BES cluster.
func GetClusterTasks(cli bce.Client, region string, request *GetClusterTasksRequest) (*GetClusterTasksResponse, error) {
	if cli == nil {
		return nil, ErrNilClient
	}
	if request == nil {
		return nil, errors.New("get cluster tasks request should not be nil")
	}
	if request.ClusterId == "" {
		return nil, errors.New("get cluster tasks request clusterId should not be empty")
	}

	uri := URI_PREFIX_V2 + "/cluster_tasks"

	result := &GetClusterTasksResponse{}
	if err := createJSONRequest(cli, region, nethttp.MethodPost, uri, request, result); err != nil {
		return nil, err
	}
	return result, nil
}

// GetClusterDataSizeTendency gets the data size tendency of a BES cluster.
func GetClusterDataSizeTendency(cli bce.Client, region string, request *GetClusterDataSizeTendencyRequest) (*GetClusterDataSizeTendencyResponse, error) {
	if cli == nil {
		return nil, ErrNilClient
	}
	if request == nil {
		return nil, errors.New("get cluster data size tendency request should not be nil")
	}
	if request.ClusterId == "" {
		return nil, errors.New("get cluster data size tendency request clusterId should not be empty")
	}
	if request.DatePattern == "" {
		return nil, errors.New("get cluster data size tendency request datePattern should not be empty")
	}
	if request.Times == 0 {
		return nil, errors.New("get cluster data size tendency request times should not be empty")
	}
	if request.TimeUnit == "" {
		return nil, errors.New("get cluster data size tendency request timeUnit should not be empty")
	}

	uri := URI_PREFIX_V2 + "/data_size_tendency"

	result := &GetClusterDataSizeTendencyResponse{}
	if err := createJSONRequest(cli, region, nethttp.MethodPost, uri, request, result); err != nil {
		return nil, err
	}
	return result, nil
}

// ListAvailableCoupons lists the user's available coupons.
func ListAvailableCoupons(cli bce.Client, region string) (*ListAvailableCouponsResponse, error) {
	if cli == nil {
		return nil, ErrNilClient
	}

	uri := URI_PREFIX_V2 + "/coupon/avail"

	result := &ListAvailableCouponsResponse{}
	if err := createJSONRequest(cli, region, nethttp.MethodPost, uri, struct{}{}, result); err != nil {
		return nil, err
	}
	return result, nil
}

// AssessClusterSource runs an intelligent capacity assessment for a cluster plan.
func AssessClusterSource(cli bce.Client, region string, request *AssessClusterSourceRequest) (*AssessClusterSourceResponse, error) {
	if cli == nil {
		return nil, ErrNilClient
	}
	if request == nil {
		return nil, errors.New("assess cluster source request should not be nil")
	}
	if request.UseType == "" {
		return nil, errors.New("assess cluster source request useType should not be empty")
	}
	if request.OriginDataSize == 0 {
		return nil, errors.New("assess cluster source request originDataSize should not be empty")
	}
	if request.OriginDataSizeUnit == "" {
		return nil, errors.New("assess cluster source request originDataSizeUnit should not be empty")
	}
	if request.DataAdd == 0 {
		return nil, errors.New("assess cluster source request dataAdd should not be empty")
	}
	if request.DataAddUnit == "" {
		return nil, errors.New("assess cluster source request dataAddUnit should not be empty")
	}
	if request.StorageDays == 0 {
		return nil, errors.New("assess cluster source request storageDays should not be empty")
	}
	if request.WriteThroughput == 0 {
		return nil, errors.New("assess cluster source request writeThroughput should not be empty")
	}
	if request.Replica == "" {
		return nil, errors.New("assess cluster source request replica should not be empty")
	}

	uri := URI_PREFIX_V2 + "/source_assess"

	result := &AssessClusterSourceResponse{}
	if err := createJSONRequest(cli, region, nethttp.MethodPost, uri, request, result); err != nil {
		return nil, err
	}
	return result, nil
}
