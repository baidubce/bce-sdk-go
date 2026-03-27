package api

import (
	"encoding/json"

	"github.com/baidubce/bce-sdk-go/bce"
	"github.com/baidubce/bce-sdk-go/http"
)

func PostMigration(cli bce.Client, args *PostMigrationArgs) (*PostMigrationResult, error) {
	if args == nil {
		return nil, bce.NewBceClientError("PostMigrationArgs is nil")
	}
	req := &bce.BceRequest{}
	req.SetMethod(http.POST)
	req.SetUri("v1")
	req.SetParam("migration", "")
	resp := &bce.BceResponse{}
	// encrypt and base64 encode ak/sk
	err := MarkAuthentication(&(args.SourceConfig.MigrationConfigCommon))
	if err != nil {
		return nil, err
	}
	err = MarkAuthentication(&(args.DestinationConfig.MigrationConfigCommon))
	if err != nil {
		return nil, err
	}

	jsonBody, err := json.Marshal(args)
	if err != nil {
		return nil, err
	}
	body, err := bce.NewBodyFromBytes(jsonBody)
	if err != nil {
		return nil, err
	}
	req.SetBody(body)
	if err := SendRequest(cli, req, resp); err != nil {
		return nil, err
	}
	if resp.IsFail() {
		return nil, resp.ServiceError()
	}
	result := &PostMigrationResult{}
	if err := resp.ParseJsonBody(result); err != nil {
		return nil, err
	}
	if result.Code != "" {
		return result, bce.NewBceServiceError(result.Code, result.Message, result.RequestId, 500)
	}
	return result, nil
}

func PostMigrationFromList(cli bce.Client, args *PostMigrationFromListArgs) (*PostMigrationResult, error) {
	if args == nil {
		return nil, bce.NewBceClientError("PostMigrationFromListArgs is nil")
	}
	req := &bce.BceRequest{}
	req.SetMethod(http.POST)
	req.SetUri("v1")
	req.SetParam("migrationFromList", "")
	resp := &bce.BceResponse{}
	// encrypt and base64 encode ak/sk
	err := MarkAuthentication(&(args.SourceConfig.MigrationConfigCommon))
	if err != nil {
		return nil, err
	}
	err = MarkAuthentication(&(args.DestinationConfig.MigrationConfigCommon))
	if err != nil {
		return nil, err
	}

	jsonBody, err := json.Marshal(args)
	if err != nil {
		return nil, err
	}
	body, err := bce.NewBodyFromBytes(jsonBody)
	if err != nil {
		return nil, err
	}
	req.SetBody(body)
	if err := SendRequest(cli, req, resp); err != nil {
		return nil, err
	}
	if resp.IsFail() {
		return nil, resp.ServiceError()
	}
	result := &PostMigrationResult{}
	if err := resp.ParseJsonBody(result); err != nil {
		return nil, err
	}
	if result.Code != "" {
		return result, bce.NewBceServiceError(result.Code, result.Message, result.RequestId, 500)
	}
	return result, nil
}

func GetMigration(cli bce.Client, taskId string) (*GetMigrationInfo, error) {
	req := &bce.BceRequest{}
	req.SetMethod(http.GET)
	req.SetUri("v1")
	req.SetParam("migration", "")
	req.SetParam("taskId", taskId)
	resp := &bce.BceResponse{}
	if err := SendRequest(cli, req, resp); err != nil {
		return nil, err
	}
	if resp.IsFail() {
		return nil, resp.ServiceError()
	}
	result := &GetMigrationInfo{}
	if err := resp.ParseJsonBody(result); err != nil {
		return nil, err
	}
	if result.Code != "" {
		return result, bce.NewBceServiceError(result.Code, result.Message, result.RequestId, 500)
	}
	return result, nil
}

func ListMigration(cli bce.Client) (*ListMigrationInfo, error) {
	req := &bce.BceRequest{}
	req.SetMethod(http.GET)
	req.SetUri("v1")
	req.SetParam("migrationList", "")
	resp := &bce.BceResponse{}
	if err := SendRequest(cli, req, resp); err != nil {
		return nil, err
	}
	if resp.IsFail() {
		return nil, resp.ServiceError()
	}
	result := &ListMigrationInfo{}
	if err := resp.ParseJsonBody(result); err != nil {
		return nil, err
	}
	if result.Code != "" {
		return result, bce.NewBceServiceError(result.Code, result.Message, result.RequestId, 500)
	}
	return result, nil
}

func GetMigrationResult(cli bce.Client, taskId string) (*MigrationResult, error) {
	req := &bce.BceRequest{}
	req.SetMethod(http.GET)
	req.SetUri("v1")
	req.SetParam("migrationResult", "")
	req.SetParam("taskId", taskId)
	resp := &bce.BceResponse{}
	if err := SendRequest(cli, req, resp); err != nil {
		return nil, err
	}
	if resp.IsFail() {
		return nil, resp.ServiceError()
	}
	result := &MigrationResult{}
	if err := resp.ParseJsonBody(result); err != nil {
		return nil, err
	}
	if result.Code != "" {
		return result, bce.NewBceServiceError(result.Code, result.Message, result.RequestId, 500)
	}
	return result, nil
}

func PauseMigration(cli bce.Client, taskId string) (*MigrationResultCommon, error) {
	req := &bce.BceRequest{}
	req.SetMethod(http.POST)
	req.SetUri("v1")
	req.SetParam("pauseMigration", "")
	req.SetParam("taskId", taskId)
	resp := &bce.BceResponse{}
	if err := SendRequest(cli, req, resp); err != nil {
		return nil, err
	}
	if resp.IsFail() {
		return nil, resp.ServiceError()
	}
	result := &MigrationResultCommon{}
	if err := resp.ParseJsonBody(result); err != nil {
		return nil, err
	}
	if result.Code != "" {
		return result, bce.NewBceServiceError(result.Code, result.Message, result.RequestId, 500)
	}
	return result, nil
}

func ResumeMigration(cli bce.Client, taskId string) (*MigrationResultCommon, error) {
	req := &bce.BceRequest{}
	req.SetMethod(http.POST)
	req.SetUri("v1")
	req.SetParam("resumeMigration", "")
	req.SetParam("taskId", taskId)
	resp := &bce.BceResponse{}
	if err := SendRequest(cli, req, resp); err != nil {
		return nil, err
	}
	if resp.IsFail() {
		return nil, resp.ServiceError()
	}
	result := &MigrationResultCommon{}
	if err := resp.ParseJsonBody(result); err != nil {
		return nil, err
	}
	if result.Code != "" {
		return result, bce.NewBceServiceError(result.Code, result.Message, result.RequestId, 500)
	}
	return result, nil
}

func RetryMigration(cli bce.Client, taskId string) (*MigrationResultCommon, error) {
	req := &bce.BceRequest{}
	req.SetMethod(http.POST)
	req.SetUri("v1")
	req.SetParam("retryMigration", "")
	req.SetParam("taskId", taskId)
	resp := &bce.BceResponse{}
	if err := SendRequest(cli, req, resp); err != nil {
		return nil, err
	}
	if resp.IsFail() {
		return nil, resp.ServiceError()
	}
	result := &MigrationResultCommon{}
	if err := resp.ParseJsonBody(result); err != nil {
		return nil, err
	}
	if result.Code != "" {
		return result, bce.NewBceServiceError(result.Code, result.Message, result.RequestId, 500)
	}
	return result, nil
}

func DeleteMigration(cli bce.Client, taskId string) (*MigrationResultCommon, error) {
	req := &bce.BceRequest{}
	req.SetMethod(http.DELETE)
	req.SetUri("v1")
	req.SetParam("migration", "")
	req.SetParam("taskId", taskId)
	resp := &bce.BceResponse{}
	if err := SendRequest(cli, req, resp); err != nil {
		return nil, err
	}
	if resp.IsFail() {
		return nil, resp.ServiceError()
	}
	result := &MigrationResultCommon{}
	if err := resp.ParseJsonBody(result); err != nil {
		return nil, err
	}
	if result.Code != "" {
		return result, bce.NewBceServiceError(result.Code, result.Message, result.RequestId, 500)
	}
	return result, nil
}
