package bbc

import (
	"fmt"

	"github.com/baidubce/bce-sdk-go/bce"
	"github.com/baidubce/bce-sdk-go/http"
)

// ListOperationLog -- xx
func (c *Client) ListOperationLog(args *ListOperationLogArgs) (ret *ListOperationLogResult, err error) {
	err = bce.NewRequestBuilder(c).
		WithMethod(http.GET).
		WithURL(getURLforOperationLog(1)).
		WithQueryParamFilter("marker", args.Marker).
		WithQueryParamFilter("maxKeys", fmt.Sprintf("%d", args.MaxKeys)).
		WithQueryParamFilter("startTime", args.StartTime).
		WithQueryParamFilter("endTime", args.EndTime).
		WithResult(&ret).
		Do()
	return
}
