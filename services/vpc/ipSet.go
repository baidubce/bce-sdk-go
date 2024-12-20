package vpc

import (
	"fmt"
	"github.com/baidubce/bce-sdk-go/bce"
	"github.com/baidubce/bce-sdk-go/http"
)

func (c *Client) CreateIpSet(args *CreateIpSetArgs) (*CreateIpSetResult, error) {
	if args == nil {
		return nil, fmt.Errorf("The CreateIpSetArgs cannot be nil.")
	}
	result := &CreateIpSetResult{}
	err := bce.NewRequestBuilder(c).
		WithURL(getURLForIpSet()).
		WithMethod(http.POST).
		WithBody(args).
		WithQueryParamFilter("clientToken", args.ClientToken).
		WithResult(result).
		Do()
	return result, err
}
