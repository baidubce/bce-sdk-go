package bill

import (
	"fmt"

	"github.com/baidubce/bce-sdk-go/bce"
	"github.com/baidubce/bce-sdk-go/http"
)

func (c *Client) GetResourceMonthBill(args *GetResourceMonthBillArgs) (list *GetResourceMonthBillResult, err error) {
	err = bce.NewRequestBuilder(c).
		WithMethod(http.GET).
		WithBody(args).
		WithURL(bce.URI_PREFIX+"v1/bill/resource/month").
		WithQueryParamFilter("month", args.Month).
		WithQueryParamFilter("productType", args.ProductType).
		WithQueryParamFilter("serviceType", args.ServiceType).
		WithQueryParamFilter("queryAccountId", args.QueryAccountId).
		WithQueryParamFilter("pageNo", fmt.Sprintf("%d", args.PageNo)).
		WithQueryParamFilter("pageSize", fmt.Sprintf("%d", args.PageSize)).
		WithResult(&list).
		Do()
	return
}
