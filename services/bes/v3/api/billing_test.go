package api

import (
	"errors"
	nethttp "net/http"
	"testing"
)

func TestConvertToPrepay(t *testing.T) {
	cli, server := newTestClient(t, nethttp.HandlerFunc(func(w nethttp.ResponseWriter, r *nethttp.Request) {
		defer r.Body.Close()

		if r.Method != nethttp.MethodPost {
			t.Fatalf("unexpected method: %s", r.Method)
		}
		if r.URL.Path != "/v3/clusters/orders/to-prepay" {
			t.Fatalf("unexpected path: %s", r.URL.Path)
		}
		if got := r.Header.Get("X-Region"); got != "sh" {
			t.Fatalf("unexpected region: %s", got)
		}

		body := readJSONBody(t, r)
		if got := body["clusterId"]; got != "search-xxxx" {
			t.Fatalf("unexpected clusterId: %v", got)
		}
		if got := body["timeLength"]; got != float64(1) {
			t.Fatalf("unexpected timeLength: %v", got)
		}
		if got := body["timeUnit"]; got != "MONTH" {
			t.Fatalf("unexpected timeUnit: %v", got)
		}

		writeJSONResponse(t, w, map[string]interface{}{
			"orderId": "order-to-prepay-example",
		})
	}))
	defer server.Close()

	result, err := ConvertToPrepay(cli, testRegion, &ConvertToPrepayRequest{
		Request:    Request{Region: "sh"},
		ClusterId:  "search-xxxx",
		TimeLength: 1,
		TimeUnit:   "MONTH",
	})
	if err != nil {
		t.Fatalf("convert to prepay failed: %v", err)
	}
	if result.OrderId != "order-to-prepay-example" {
		t.Fatalf("unexpected orderId: %s", result.OrderId)
	}
}

// TestConvertToPostpay leaves Request.Region unset on purpose so the region argument passed by the
// caller is the one that ends up on the wire.
func TestConvertToPostpay(t *testing.T) {
	cli, server := newTestClient(t, nethttp.HandlerFunc(func(w nethttp.ResponseWriter, r *nethttp.Request) {
		defer r.Body.Close()

		if r.Method != nethttp.MethodPost {
			t.Fatalf("unexpected method: %s", r.Method)
		}
		if r.URL.Path != "/v3/clusters/orders/to-postpay" {
			t.Fatalf("unexpected path: %s", r.URL.Path)
		}
		if got := r.Header.Get("X-Region"); got != testRegion {
			t.Fatalf("unexpected region: %s", got)
		}

		body := readJSONBody(t, r)
		if got := body["clusterId"]; got != "search-xxxx" {
			t.Fatalf("unexpected clusterId: %v", got)
		}

		writeJSONResponse(t, w, map[string]interface{}{
			"orderId": "order-to-postpay-example",
		})
	}))
	defer server.Close()

	result, err := ConvertToPostpay(cli, testRegion, &ConvertToPostpayRequest{ClusterId: "search-xxxx"})
	if err != nil {
		t.Fatalf("convert to postpay failed: %v", err)
	}
	if result.OrderId != "order-to-postpay-example" {
		t.Fatalf("unexpected orderId: %s", result.OrderId)
	}
}

func TestCancelConvertToPostpay(t *testing.T) {
	cli, server := newTestClient(t, nethttp.HandlerFunc(func(w nethttp.ResponseWriter, r *nethttp.Request) {
		defer r.Body.Close()

		if r.Method != nethttp.MethodPost {
			t.Fatalf("unexpected method: %s", r.Method)
		}
		if r.URL.Path != "/v3/clusters/orders/to-postpay/cancel" {
			t.Fatalf("unexpected path: %s", r.URL.Path)
		}
		if got := r.Header.Get("X-Region"); got != "gz" {
			t.Fatalf("unexpected region: %s", got)
		}

		body := readJSONBody(t, r)
		if got := body["clusterId"]; got != "search-xxxx" {
			t.Fatalf("unexpected clusterId: %v", got)
		}

		writeJSONResponse(t, w, map[string]interface{}{
			"orderId": "order-to-postpay-example",
		})
	}))
	defer server.Close()

	result, err := CancelConvertToPostpay(cli, testRegion, &CancelConvertToPostpayRequest{
		Request:   Request{Region: "gz"},
		ClusterId: "search-xxxx",
	})
	if err != nil {
		t.Fatalf("cancel convert to postpay failed: %v", err)
	}
	if result.OrderId != "order-to-postpay-example" {
		t.Fatalf("unexpected orderId: %s", result.OrderId)
	}
}

func TestRenewCluster(t *testing.T) {
	cli, server := newTestClient(t, nethttp.HandlerFunc(func(w nethttp.ResponseWriter, r *nethttp.Request) {
		defer r.Body.Close()

		if r.Method != nethttp.MethodPost {
			t.Fatalf("unexpected method: %s", r.Method)
		}
		if r.URL.Path != "/v3/order/renew/confirm" {
			t.Fatalf("unexpected path: %s", r.URL.Path)
		}
		if got := r.Header.Get("X-Region"); got != "sh" {
			t.Fatalf("unexpected region: %s", got)
		}

		body := readJSONBody(t, r)
		if got := body["clusterId"]; got != "search-xxxx" {
			t.Fatalf("unexpected clusterId: %v", got)
		}
		if got := body["timeLength"]; got != float64(1) {
			t.Fatalf("unexpected timeLength: %v", got)
		}
		if got := body["timeUnit"]; got != "MONTH" {
			t.Fatalf("unexpected timeUnit: %v", got)
		}

		writeJSONResponse(t, w, map[string]interface{}{
			"orderId": "order-renew-example",
		})
	}))
	defer server.Close()

	result, err := RenewCluster(cli, testRegion, &RenewClusterRequest{
		Request:    Request{Region: "sh"},
		ClusterId:  "search-xxxx",
		TimeLength: 1,
		TimeUnit:   "MONTH",
	})
	if err != nil {
		t.Fatalf("renew cluster failed: %v", err)
	}
	if result.OrderId != "order-renew-example" {
		t.Fatalf("unexpected orderId: %s", result.OrderId)
	}
}

func TestQueryConfigPrice(t *testing.T) {
	cli, server := newTestClient(t, nethttp.HandlerFunc(func(w nethttp.ResponseWriter, r *nethttp.Request) {
		defer r.Body.Close()

		if r.Method != nethttp.MethodPost {
			t.Fatalf("unexpected method: %s", r.Method)
		}
		if r.URL.Path != "/v3/orders/prices" {
			t.Fatalf("unexpected path: %s", r.URL.Path)
		}
		if got := r.Header.Get("X-Region"); got != testRegion {
			t.Fatalf("unexpected region: %s", got)
		}

		body := readJSONBody(t, r)
		if got := body["payment"]; got != PAYMENT_PREPAID {
			t.Fatalf("unexpected payment: %v", got)
		}
		if got := body["timeLength"]; got != float64(3) {
			t.Fatalf("unexpected timeLength: %v", got)
		}
		if got := body["timeUnit"]; got != "MONTH" {
			t.Fatalf("unexpected timeUnit: %v", got)
		}
		nodeSpecs, ok := body["nodeSpecs"].([]interface{})
		if !ok || len(nodeSpecs) != 2 {
			t.Fatalf("unexpected nodeSpecs: %v", body["nodeSpecs"])
		}

		writeJSONResponse(t, w, map[string]interface{}{
			"price":        0.022076928,
			"catalogPrice": 0.044153856,
			"nodeTypePrices": []map[string]interface{}{
				{
					"type":             "data",
					"nodePrice":        0.016580586,
					"nodeCatalogPrice": 0.033161172,
					"diskPrice":        0.0017325,
					"diskCatalogPrice": 0.003465,
				},
				{
					"type":             "dashboards",
					"nodePrice":        0.003763842,
					"nodeCatalogPrice": 0.007527684,
				},
			},
		})
	}))
	defer server.Close()

	result, err := QueryConfigPrice(cli, testRegion, &QueryConfigPriceRequest{
		Payment:    PAYMENT_PREPAID,
		TimeLength: 3,
		TimeUnit:   "MONTH",
		NodeSpecs: []NodeSpec{
			{Type: "data", NodeType: "opensearch.c2m8", NodeCount: intPtr(3), DiskType: "enhanced_ssd_pl1", DiskSize: intPtr(50), DiskCount: intPtr(1)},
			{Type: "dashboards", NodeType: "opensearch.c2m4", NodeCount: intPtr(1)},
		},
	})
	if err != nil {
		t.Fatalf("query config price failed: %v", err)
	}
	if result.Price != 0.022076928 {
		t.Fatalf("unexpected price: %v", result.Price)
	}
	if len(result.NodeTypePrices) != 2 || result.NodeTypePrices[0].Type != "data" {
		t.Fatalf("unexpected nodeTypePrices: %+v", result.NodeTypePrices)
	}
}

// billingInt64Ptr is local to the billing tests because BillingConfig.ExpirationTime is the only
// *int64 field exercised here; the shared testhelper only provides boolPtr and intPtr.
func billingInt64Ptr(v int64) *int64 { return &v }

// TestQueryClusterMarginPrice populates every optional request field (resizeType, billingConfig and
// its nested optional fields, deploy set flag) so the full payload is exercised.
func TestQueryClusterMarginPrice(t *testing.T) {
	cli, server := newTestClient(t, nethttp.HandlerFunc(func(w nethttp.ResponseWriter, r *nethttp.Request) {
		defer r.Body.Close()

		if r.Method != nethttp.MethodPost {
			t.Fatalf("unexpected method: %s", r.Method)
		}
		if r.URL.Path != "/v3/clusters/margin-prices" {
			t.Fatalf("unexpected path: %s", r.URL.Path)
		}
		if got := r.Header.Get("X-Region"); got != "sh" {
			t.Fatalf("unexpected region: %s", got)
		}

		body := readJSONBody(t, r)
		if got := body["clusterId"]; got != "search-5e810SYNgCoBW5wOBkfY" {
			t.Fatalf("unexpected clusterId: %v", got)
		}
		if got := body["mode"]; got != RESIZE_MODE_BLUE_GREEN {
			t.Fatalf("unexpected mode: %v", got)
		}
		if got := body["resizeType"]; got != float64(1) {
			t.Fatalf("unexpected resizeType: %v", got)
		}
		if got := body["vpcId"]; got != "vpc-pekgs85bpt95" {
			t.Fatalf("unexpected vpcId: %v", got)
		}
		if got := body["subnetId"]; got != "sbn-gcv4r388upuq" {
			t.Fatalf("unexpected subnetId: %v", got)
		}
		if got := body["enableDeploySet"]; got != false {
			t.Fatalf("unexpected enableDeploySet: %v", got)
		}
		if zones, ok := body["logicalZones"].([]interface{}); !ok || len(zones) != 1 || zones[0] != "cn-bj-a" {
			t.Fatalf("unexpected logicalZones: %v", body["logicalZones"])
		}
		if nodeSpecs, ok := body["nodeSpecs"].([]interface{}); !ok || len(nodeSpecs) != 2 {
			t.Fatalf("unexpected nodeSpecs: %v", body["nodeSpecs"])
		}
		billingConfig, ok := body["billingConfig"].(map[string]interface{})
		if !ok {
			t.Fatalf("unexpected billingConfig: %v", body["billingConfig"])
		}
		if got := billingConfig["payment"]; got != PAYMENT_POSTPAID {
			t.Fatalf("unexpected billingConfig payment: %v", got)
		}
		if got := billingConfig["timeLength"]; got != float64(1) {
			t.Fatalf("unexpected billingConfig timeLength: %v", got)
		}
		if got := billingConfig["timeUnit"]; got != "MONTH" {
			t.Fatalf("unexpected billingConfig timeUnit: %v", got)
		}
		if got := billingConfig["expirationTime"]; got != float64(1767225600000) {
			t.Fatalf("unexpected billingConfig expirationTime: %v", got)
		}
		if got := billingConfig["isAutoPay"]; got != true {
			t.Fatalf("unexpected billingConfig isAutoPay: %v", got)
		}

		writeJSONResponse(t, w, map[string]interface{}{
			"currentComponentPrices": map[string]interface{}{
				"Postpaid": map[string]interface{}{"price": 0.022076928, "catalogPrice": 0.044153856, "discountRate": 0.5},
			},
			"updateComponentPrices": map[string]interface{}{
				"Postpaid": map[string]interface{}{"price": 0.037800048, "catalogPrice": 0.075600096, "discountRate": 0.5},
			},
			"marginComponentPrices": map[string]interface{}{
				"Postpaid": map[string]interface{}{"price": 0.01572312, "catalogPrice": 0.03144624, "discountRate": 0.5},
			},
			"expirationTime": "2026-12-31T16:00:00Z",
		})
	}))
	defer server.Close()

	result, err := QueryClusterMarginPrice(cli, testRegion, newFullMarginPriceRequest())
	if err != nil {
		t.Fatalf("query cluster margin price failed: %v", err)
	}
	if len(result.CurrentComponentPrices) != 1 || len(result.UpdateComponentPrices) != 1 {
		t.Fatalf("unexpected component prices: %+v", result)
	}
	if got := result.MarginComponentPrices["Postpaid"].Price; got != 0.01572312 {
		t.Fatalf("unexpected margin price: %v", got)
	}
	if result.ExpirationTime != "2026-12-31T16:00:00Z" {
		t.Fatalf("unexpected expirationTime: %s", result.ExpirationTime)
	}
}

// newFullMarginPriceRequest returns a request with every field populated, so validation tests can
// clear exactly one field per case.
func newFullMarginPriceRequest() *QueryClusterMarginPriceRequest {
	return &QueryClusterMarginPriceRequest{
		Request:    Request{Region: "sh"},
		ClusterId:  "search-5e810SYNgCoBW5wOBkfY",
		ResizeType: intPtr(1),
		BillingConfig: &BillingConfig{
			Payment:        PAYMENT_POSTPAID,
			TimeLength:     intPtr(1),
			TimeUnit:       "MONTH",
			ExpirationTime: billingInt64Ptr(1767225600000),
			IsAutoPay:      boolPtr(true),
		},
		LogicalZones:    []string{"cn-bj-a"},
		VpcId:           "vpc-pekgs85bpt95",
		SubnetId:        "sbn-gcv4r388upuq",
		EnableDeploySet: boolPtr(false),
		NodeSpecs: []NodeSpec{
			{Type: "data", NodeType: "opensearch.c4m16", NodeCount: intPtr(3), DiskType: "enhanced_ssd_pl1", DiskSize: intPtr(50), DiskCount: intPtr(1)},
			{Type: "dashboards", NodeType: "opensearch.c2m4", NodeCount: intPtr(1)},
		},
		Mode: RESIZE_MODE_BLUE_GREEN,
	}
}

func TestBillingAPIRejectsNilClient(t *testing.T) {
	cases := map[string]func() error{
		"ConvertToPrepay": func() error {
			_, err := ConvertToPrepay(nil, testRegion, &ConvertToPrepayRequest{})
			return err
		},
		"ConvertToPostpay": func() error {
			_, err := ConvertToPostpay(nil, testRegion, &ConvertToPostpayRequest{})
			return err
		},
		"CancelConvertToPostpay": func() error {
			_, err := CancelConvertToPostpay(nil, testRegion, &CancelConvertToPostpayRequest{})
			return err
		},
		"RenewCluster": func() error {
			_, err := RenewCluster(nil, testRegion, &RenewClusterRequest{})
			return err
		},
		"QueryConfigPrice": func() error {
			_, err := QueryConfigPrice(nil, testRegion, &QueryConfigPriceRequest{})
			return err
		},
		"QueryClusterMarginPrice": func() error {
			_, err := QueryClusterMarginPrice(nil, testRegion, &QueryClusterMarginPriceRequest{})
			return err
		},
	}
	for name, call := range cases {
		if err := call(); !errors.Is(err, ErrNilClient) {
			t.Fatalf("%s: got %v, want ErrNilClient", name, err)
		}
	}
}

func TestBillingAPIRejectsNilRequest(t *testing.T) {
	cli, server := newTestClient(t, nethttp.HandlerFunc(func(w nethttp.ResponseWriter, r *nethttp.Request) {
		t.Fatal("request should not be sent")
	}))
	defer server.Close()

	cases := map[string]func() error{
		"ConvertToPrepay":         func() error { _, err := ConvertToPrepay(cli, testRegion, nil); return err },
		"ConvertToPostpay":        func() error { _, err := ConvertToPostpay(cli, testRegion, nil); return err },
		"CancelConvertToPostpay":  func() error { _, err := CancelConvertToPostpay(cli, testRegion, nil); return err },
		"RenewCluster":            func() error { _, err := RenewCluster(cli, testRegion, nil); return err },
		"QueryConfigPrice":        func() error { _, err := QueryConfigPrice(cli, testRegion, nil); return err },
		"QueryClusterMarginPrice": func() error { _, err := QueryClusterMarginPrice(cli, testRegion, nil); return err },
	}
	for name, call := range cases {
		if err := call(); err == nil {
			t.Fatalf("%s: expected error for nil request", name)
		}
	}
}

func TestBillingAPIRejectsMissingRequiredFields(t *testing.T) {
	cli, server := newTestClient(t, nethttp.HandlerFunc(func(w nethttp.ResponseWriter, r *nethttp.Request) {
		t.Fatal("request should not be sent")
	}))
	defer server.Close()

	cases := map[string]func() error{
		"ConvertToPrepay/clusterId": func() error {
			_, err := ConvertToPrepay(cli, testRegion, &ConvertToPrepayRequest{TimeLength: 1})
			return err
		},
		"ConvertToPrepay/timeLength": func() error {
			_, err := ConvertToPrepay(cli, testRegion, &ConvertToPrepayRequest{ClusterId: "search-xxxx"})
			return err
		},
		"ConvertToPostpay/clusterId": func() error {
			_, err := ConvertToPostpay(cli, testRegion, &ConvertToPostpayRequest{})
			return err
		},
		"CancelConvertToPostpay/clusterId": func() error {
			_, err := CancelConvertToPostpay(cli, testRegion, &CancelConvertToPostpayRequest{})
			return err
		},
		"RenewCluster/clusterId": func() error {
			_, err := RenewCluster(cli, testRegion, &RenewClusterRequest{TimeLength: 1})
			return err
		},
		"RenewCluster/timeLength": func() error {
			_, err := RenewCluster(cli, testRegion, &RenewClusterRequest{ClusterId: "search-xxxx"})
			return err
		},
		"QueryConfigPrice/payment": func() error {
			_, err := QueryConfigPrice(cli, testRegion, &QueryConfigPriceRequest{
				NodeSpecs: []NodeSpec{{Type: "data", NodeType: "opensearch.c2m8"}},
			})
			return err
		},
		"QueryConfigPrice/nodeSpecs": func() error {
			_, err := QueryConfigPrice(cli, testRegion, &QueryConfigPriceRequest{Payment: PAYMENT_POSTPAID})
			return err
		},
	}
	for name, call := range cases {
		if err := call(); err == nil {
			t.Fatalf("%s: expected validation error", name)
		}
	}
}

func TestQueryClusterMarginPriceRejectsMissingRequiredFields(t *testing.T) {
	cli, server := newTestClient(t, nethttp.HandlerFunc(func(w nethttp.ResponseWriter, r *nethttp.Request) {
		t.Fatal("request should not be sent")
	}))
	defer server.Close()

	cases := map[string]func(*QueryClusterMarginPriceRequest){
		"clusterId": func(r *QueryClusterMarginPriceRequest) { r.ClusterId = "" },
		"nodeSpecs": func(r *QueryClusterMarginPriceRequest) { r.NodeSpecs = nil },
		"mode":      func(r *QueryClusterMarginPriceRequest) { r.Mode = "" },
	}
	for name, clear := range cases {
		request := newFullMarginPriceRequest()
		clear(request)
		if _, err := QueryClusterMarginPrice(cli, testRegion, request); err == nil {
			t.Fatalf("%s: expected validation error", name)
		}
	}
}

// TestBillingAPIPropagatesServerError uses 400 rather than 500 on purpose: the helper mirrors the
// production retry policy, which retries 5xx with backoff and would make this test take seconds.
func TestBillingAPIPropagatesServerError(t *testing.T) {
	cli, server := newTestClient(t, nethttp.HandlerFunc(func(w nethttp.ResponseWriter, r *nethttp.Request) {
		defer r.Body.Close()
		w.WriteHeader(nethttp.StatusBadRequest)
		writeJSONResponse(t, w, map[string]interface{}{
			"code":      "InvalidParameter",
			"message":   "invalid parameter",
			"requestId": "req-xxxx",
		})
	}))
	defer server.Close()

	cases := map[string]func() error{
		"ConvertToPrepay": func() error {
			_, err := ConvertToPrepay(cli, testRegion, &ConvertToPrepayRequest{ClusterId: "search-xxxx", TimeLength: 1})
			return err
		},
		"ConvertToPostpay": func() error {
			_, err := ConvertToPostpay(cli, testRegion, &ConvertToPostpayRequest{ClusterId: "search-xxxx"})
			return err
		},
		"CancelConvertToPostpay": func() error {
			_, err := CancelConvertToPostpay(cli, testRegion, &CancelConvertToPostpayRequest{ClusterId: "search-xxxx"})
			return err
		},
		"RenewCluster": func() error {
			_, err := RenewCluster(cli, testRegion, &RenewClusterRequest{ClusterId: "search-xxxx", TimeLength: 1})
			return err
		},
		"QueryConfigPrice": func() error {
			_, err := QueryConfigPrice(cli, testRegion, &QueryConfigPriceRequest{
				Payment:   PAYMENT_POSTPAID,
				NodeSpecs: []NodeSpec{{Type: "data", NodeType: "opensearch.c2m8"}},
			})
			return err
		},
		"QueryClusterMarginPrice": func() error {
			_, err := QueryClusterMarginPrice(cli, testRegion, newFullMarginPriceRequest())
			return err
		},
	}
	for name, call := range cases {
		if err := call(); err == nil {
			t.Fatalf("%s: expected error for 400 response", name)
		}
	}
}
