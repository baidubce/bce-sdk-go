package api

import (
	"encoding/json"
	"time"

	"github.com/baidubce/bce-sdk-go/bce"
	bcehttp "github.com/baidubce/bce-sdk-go/http"
)

// serviceRequest handles common service API request logic
func serviceRequest(cli bce.Client, uri string, method string, args interface{}, params map[string]interface{}) (map[string]interface{}, error) {
	// Validate args if it implements Validator
	if args != nil {
		if validator, ok := args.(Validator); ok {
			if err := validator.Validate(); err != nil {
				return nil, err
			}
		}
	}

	// Create request
	req := new(bce.BceRequest)
	req.SetUri(uri)
	req.SetMethod(method)

	// Set query parameters
	if params != nil {
		for key, value := range params {
			if valueStr, ok := value.(string); ok {
				req.SetParam(key, valueStr)
			}
		}
	}

	// Set request body for POST/PUT methods
	if (method == POST || method == PUT) && args != nil {
		argsBytes, err := json.Marshal(args)
		if err != nil {
			return nil, err
		}
		requestBody, err := bce.NewBodyFromBytes(argsBytes)
		if err != nil {
			return nil, err
		}
		req.SetHeader(bcehttp.CONTENT_TYPE, bce.DEFAULT_CONTENT_TYPE)
		req.SetBody(requestBody)
	}

	// Send request
	resp := &bce.BceResponse{}
	if err := cli.SendRequest(req, resp); err != nil {
		return nil, err
	}
	defer resp.Body().Close()
	if resp.IsFail() {
		return nil, resp.ServiceError()
	}

	// Parse JSON response
	var rawResponse struct {
		Success bool                   `json:"success"`
		Result  map[string]interface{} `json:"result"`
	}

	if err := resp.ParseJsonBody(&rawResponse); err != nil {
		return nil, err
	}

	return rawResponse.Result, nil
}

func GetService(cli bce.Client, args *GetServiceArgs) (*GetServiceResult, error) {
	params := map[string]interface{}{
		"serviceName": args.ServiceName,
	}

	resultMap, err := serviceRequest(cli, getServiceUri(), GET, args, params)
	if err != nil {
		return nil, err
	}

	service := convertMapToService(resultMap)
	return (*GetServiceResult)(service), nil
}

func CreateService(cli bce.Client, args *CreateServiceArgs) (*CreateServiceResult, error) {
	resultMap, err := serviceRequest(cli, serviceUri(), POST, args, nil)
	if err != nil {
		return nil, err
	}

	service := convertMapToService(resultMap)
	return (*CreateServiceResult)(service), nil
}

func UpdateService(cli bce.Client, args *UpdateServiceArgs) (*UpdateServiceResult, error) {
	params := map[string]interface{}{
		"serviceName": args.ServiceName,
	}

	resultMap, err := serviceRequest(cli, serviceUri(), PUT, args, params)
	if err != nil {
		return nil, err
	}

	service := convertMapToService(resultMap)
	return (*UpdateServiceResult)(service), nil
}

func DeleteService(cli bce.Client, args *DeleteServiceArgs) error {
	params := map[string]interface{}{
		"serviceName": args.ServiceName,
	}

	_, err := serviceRequest(cli, serviceUri(), DELETE, args, params)
	return err
}

func ListServices(cli bce.Client) (*ListServicesResult, error) {
	// Create request
	req := new(bce.BceRequest)
	req.SetUri(serviceUri())
	req.SetMethod(GET)

	// Send request
	resp := &bce.BceResponse{}
	if err := cli.SendRequest(req, resp); err != nil {
		return nil, err
	}
	defer resp.Body().Close()
	if resp.IsFail() {
		return nil, resp.ServiceError()
	}

	// Parse JSON response - ListServices has a nested result structure
	var rawResponse struct {
		Success bool `json:"success"`
		Result  struct {
			Result []map[string]interface{} `json:"result"`
		} `json:"result"`
	}

	if err := resp.ParseJsonBody(&rawResponse); err != nil {
		return nil, err
	}

	// Convert to Service objects
	services := make([]*Service, 0, len(rawResponse.Result.Result))
	for _, item := range rawResponse.Result.Result {
		service := convertMapToService(item)
		services = append(services, service)
	}

	return &ListServicesResult{Services: services}, nil
}

// Helper function to convert map[string]interface{} to Service
func convertMapToService(item map[string]interface{}) *Service {
	service := &Service{}

	// Manual conversion from map[string]interface{} to Service
	if uid, exists := item["Uid"]; exists {
		if uidStr, ok := uid.(string); ok {
			service.Uid = uidStr
		}
	}
	if serviceName, exists := item["ServiceName"]; exists {
		if serviceNameStr, ok := serviceName.(string); ok {
			service.ServiceName = serviceNameStr
		}
	}
	if serviceDesc, exists := item["ServiceDesc"]; exists {
		if serviceDescStr, ok := serviceDesc.(string); ok {
			service.ServiceDesc = &serviceDescStr
		}
	}
	if serviceConf, exists := item["ServiceConf"]; exists {
		if serviceConfStr, ok := serviceConf.(string); ok {
			service.ServiceConf = serviceConfStr
		}
	}
	if serviceConfig, exists := item["ServiceConfig"]; exists {
		if configMap, ok := serviceConfig.(map[string]interface{}); ok {
			service.ServiceConfig = configMap
		}
	}
	if region, exists := item["Region"]; exists {
		if regionStr, ok := region.(string); ok {
			service.Region = regionStr
		}
	}
	if status, exists := item["Status"]; exists {
		if statusFloat, ok := status.(float64); ok {
			service.Status = int(statusFloat)
		}
	}
	// Parse time fields
	if updatedAt, exists := item["UpdatedAt"]; exists {
		if updatedAtStr, ok := updatedAt.(string); ok {
			if t, err := time.Parse(time.RFC3339, updatedAtStr); err == nil {
				service.UpdatedAt = t
			}
		}
	}
	if createdAt, exists := item["CreatedAt"]; exists {
		if createdAtStr, ok := createdAt.(string); ok {
			if t, err := time.Parse(time.RFC3339, createdAtStr); err == nil {
				service.CreatedAt = t
			}
		}
	}

	return service
}
