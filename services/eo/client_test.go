package eo

import (
	"encoding/json"
	"fmt"
	"testing"

	"github.com/baidubce/bce-sdk-go/bce"
	"github.com/baidubce/bce-sdk-go/services/eo/api"
)

const (
	testAuthoritySite   = "your_valid_site"
	testAuthorityDomain = "your_valid_domain"
	testEndpoint        = "geo.baidubce.com"
	testAK              = "your_access_key_id"
	testSK              = "your_secret_key_id"

	// set testConfigOk true for unit test
	testConfigOk = false
)

var testCli *Client

func TestMain(m *testing.M) {
	if !testConfigOk {
		fmt.Printf("TestMain terminated, please check testing config")
		return
	}

	var err error
	testCli, err = NewClient(testAK, testSK, testEndpoint)
	if err != nil {
		fmt.Printf("TestMain terminated, err:%+v\n", err)
		return
	}

	m.Run()
}

func checkClientErr(t *testing.T, funcName string, err error) {
	if funcName == "" {
		t.Fatalf(`error param when called checkClientErr, the funcName is ""`)
	}

	if !testConfigOk {
		t.Logf("Configuration did not complete initialization\n")
		return
	}

	if err == nil {
		return
	}

	e, ok := err.(*bce.BceServiceError)
	if !ok {
		t.Fatalf("%s: %v\n", funcName, err)
		return
	}

	// `AccessDenied` indicates unauthorized AK/SK.
	// `InvalidArgument` indicates sending the error params to server.
	// `NotFound` indicates using error method.
	if e.Code == "AccessDenied" || e.Code == "InvalidArgument" || e.Code == "NotFound" {
		t.Fatalf("%s: %v\n", funcName, err)
	}

	// we do not judge the errors in business logic.
	t.Logf("%s: UT is ok, but there is a logic error:\n%s", funcName, err.Error())
}

// //////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
// Test functions about purge and prefetch.
// //////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

func TestPurge(t *testing.T) {
	purgedId, err := testCli.Purge(testAuthoritySite, []api.PurgeTask{
		{
			Url: "http://my.domain.com/path/to/purge/2.data",
		},
		{
			Url:  "http://my.domain.com/path/to/purege/html/",
			Type: "directory",
		},
	})

	t.Logf("purgedId: %s", string(purgedId))
	checkClientErr(t, "Purge", err)
}

// TestGetPurgedStatusByTimeRangeAndType - query purge status by site and time range and optional type filter
func TestGetPurgedStatusByTimeRangeAndType(t *testing.T) {
	purgedStatus, err := testCli.GetPurgedStatus(&api.PurgeStatusQueryData{
		Site:      testAuthoritySite,
		StartTime: "2026-05-01T00:00:00Z",
		EndTime:   "2026-05-31T23:59:59Z",
		Type:      "file", // optional: filter by purge type, can be "file" or "directory"
	})

	data, _ := json.Marshal(purgedStatus)
	t.Logf("purgedStatus : %s", string(data))
	checkClientErr(t, "GetPurgedStatusByTimeRangeAndType", err)
}

// TestGetPurgedStatusById - query purge status by site and task id
func TestGetPurgedStatusById(t *testing.T) {
	purgedStatus, err := testCli.GetPurgedStatus(&api.PurgeStatusQueryData{
		Site: testAuthoritySite,
		Id:   "your_purge_task_id",
	})

	data, _ := json.Marshal(purgedStatus)
	t.Logf("purgedStatus : %s", string(data))
	checkClientErr(t, "GetPurgedStatusById", err)
}

func TestPrefetch(t *testing.T) {
	prefetchId, err := testCli.Prefetch(testAuthoritySite, []api.PrefetchTask{
		{
			Url: "http://my.domain.com/path/to/prefetch/1.data",
		},
	})

	t.Logf("prefetchId: %s", string(prefetchId))
	checkClientErr(t, "Prefetch", err)
}

// TestGetPrefetchStatusByTimeRange - query prefetch status by site and time range
func TestGetPrefetchStatusByTimeRange(t *testing.T) {
	prefetchStatus, err := testCli.GetPrefetchStatus(&api.PrefetchStatusQueryData{
		Site:      testAuthoritySite,
		StartTime: "2026-05-01T00:00:00Z",
		EndTime:   "2026-05-31T23:59:59Z",
	})

	data, _ := json.Marshal(prefetchStatus)
	t.Logf("prefetchStatus : %s", string(data))
	checkClientErr(t, "GetPrefetchStatusByTimeRange", err)
}

func TestGetPrefetchStatusById(t *testing.T) {
	prefetchStatus, err := testCli.GetPrefetchStatus(&api.PrefetchStatusQueryData{
		Site: testAuthoritySite,
		Id:   "your_prefetch_task_id",
	})

	data, _ := json.Marshal(prefetchStatus)
	t.Logf("prefetchStatus : %s", string(data))
	checkClientErr(t, "GetPrefetchStatusById", err)
}

// //////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
// Test functions about offline log.
// //////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

// TestGetOfflineLog - query offline log download urls by site, time range and domain list
func TestGetOfflineLog(t *testing.T) {
	logResult, err := testCli.GetOfflineLog(&api.LogQueryData{
		Site:       testAuthoritySite,
		StartTime:  "2026-06-01T17:00:00Z",
		EndTime:    "2026-06-02T11:03:19Z",
		DomainList: []string{testAuthorityDomain},
		PageNo:     1,
		PageSize:   20,
	})

	data, _ := json.Marshal(logResult)
	t.Logf("offlineLog: %s", string(data))
	checkClientErr(t, "GetOfflineLog", err)
}

// //////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
// Test functions about stat metrics.
// //////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

// TestGetStatTime - type 1: 时间打点（无 group），查询带宽打点
func TestGetStatTime(t *testing.T) {
	result, err := testCli.GetStatTime(&api.StatTimeQueryData{
		Metric:    []string{"sum_bps", "upstream_bps", "download_bps"},
		StartTime: "2026-06-02T16:00:00Z",
		EndTime:   "2026-06-03T02:03:14Z",
		Filter: []api.StatFilterItem{
			{
				Key:       "host",
				Operation: "equal",
				Value:     []string{testAuthorityDomain},
			},
		},
	})

	data, _ := json.Marshal(result)
	t.Logf("stat time series: %s", string(data))
	checkClientErr(t, "GetStatTime", err)
}

// TestGetStatPeak - type 2: 峰值（无 group），查询峰值带宽
func TestGetStatPeak(t *testing.T) {
	result, err := testCli.GetStatPeak(&api.StatPeakQueryData{
		Metric:    []string{"sum_bps", "upstream_bps", "download_bps"},
		StartTime: "2026-06-02T16:00:00Z",
		EndTime:   "2026-06-03T02:03:14Z",
		Filter: []api.StatFilterItem{
			{
				Key:       "host",
				Operation: "equal",
				Value:     []string{testAuthorityDomain},
			},
		},
	})

	data, _ := json.Marshal(result)
	t.Logf("stat peak: %s", string(data))
	checkClientErr(t, "GetStatPeak", err)
}

// TestGetStatTimeByGroup - type 3: 时间打点带 group 参数，按状态码分组查 PV 打点
func TestGetStatTimeByGroup(t *testing.T) {
	result, err := testCli.GetStatTimeByGroup(&api.StatTimeByGroupQueryData{
		Metric:    []string{"pv"},
		StartTime: "2026-06-02T16:00:00Z",
		EndTime:   "2026-06-03T02:03:14Z",
		Group:     []string{"code"},
		Filter: []api.StatFilterItem{
			{
				Key:       "site",
				Operation: "equal",
				Value:     []string{testAuthoritySite},
			},
		},
	})

	data, _ := json.Marshal(result)
	t.Logf("stat time by group: %s", string(data))
	checkClientErr(t, "GetStatTimeByGroup", err)
}

// TestGetStatSumByGroup - type 4: 聚合带 group 参数，按状态码分组查 PV 总和
func TestGetStatSumByGroup(t *testing.T) {
	result, err := testCli.GetStatSumByGroup(&api.StatSumByGroupQueryData{
		Metric:    []string{"pv"},
		StartTime: "2026-06-02T16:00:00Z",
		EndTime:   "2026-06-03T02:03:14Z",
		Group:     []string{"code"},
		Filter: []api.StatFilterItem{
			{
				Key:       "site",
				Operation: "equal",
				Value:     []string{testAuthoritySite},
			},
		},
	})

	data, _ := json.Marshal(result)
	t.Logf("stat sum by group: %s", string(data))

	checkClientErr(t, "GetStatSumByGroup", err)
}

// TestGetStatTopByGroup - type 5: TOP 带 group 参数，按 host 分组查 PV TOP
func TestGetStatTopByGroup(t *testing.T) {
	result, err := testCli.GetStatTopByGroup(&api.StatTopByGroupQueryData{
		Metric:    []string{"pv"},
		StartTime: "2026-06-02T16:00:00Z",
		EndTime:   "2026-06-03T02:03:14Z",
		Group:     []string{"host"},
		Filter: []api.StatFilterItem{
			{
				Key:       "site",
				Operation: "equal",
				Value:     []string{testAuthoritySite},
			},
		},
		Limit: &api.StatLimit{PageSize: 100},
	})

	data, _ := json.Marshal(result)
	t.Logf("stat top by group: %s", string(data))
	checkClientErr(t, "GetStatTopByGroup", err)
}

// //////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
// Test functions about site config.
// //////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

// TestSetSiteConfig - set site config; SiteConfig is a unified container,
// any subset of fields can be combined in a single request.
func TestSetSiteConfig(t *testing.T) {
	query := false
	cacheKeyIgnoreCase := false
	offlineMode := "ON"
	httpToHttpsEnabled := "ON"
	httpToHttpsCode := "302"
	hstsMaxAge := 20
	hstsIncludeSubDomains := true
	hstsPreload := false
	http2Disable := "OFF"
	http3Enable := false
	clientMaxBodySize := "500m"
	compress := "ON"
	isa := "ON"
	grpcOrigin := "ON"
	http2Origin := "ON"
	ruleCacheKeyQuery := false
	ruleCacheKeyIgnoreCase := false
	clientCacheTtlMode := "custom"
	clientCacheTtl := 150
	webSocketEnabled := true
	webSocketTimeout := 10
	sslMode := "strong"
	sslCipherList := "TLS_ECDHE_ECDSA_WITH_CHACHA20_POLY1305_SHA256:TLS_ECDHE_RSA_WITH_CHACHA20_POLY1305_SHA256:" +
		"TLS_ECDHE_ECDSA_WITH_AES_128_GCM_SHA256:TLS_ECDHE_RSA_WITH_AES_128_GCM_SHA256:" +
		"TLS_ECDHE_ECDSA_WITH_AES_256_GCM_SHA384:TLS_ECDHE_RSA_WITH_AES_256_GCM_SHA384"
	sslModePtr := &sslMode
	sslCipherListPtr := &sslCipherList
	ocsp := "ON"
	realIpEnabled := true
	realIpName := "True-Client-Ip"
	originHttpPort := 80
	originHttpsPort := 443
	originHost := "1.test.com"

	result, err := testCli.SetSiteConfig(testAuthoritySite, &api.SiteConfig{
		// 节点缓存配置
		CacheTtl: &[]api.CacheTtl{
			{
				Value:          "/",
				Weight:         100,
				OverrideOrigin: true,
				Ttl:            1,
				Type:           "path",
			},
		},

		// 查询字符串配置
		CacheKey: &api.CacheKey{
			Query:       &query,
			IncludeArgs: &[]string{"test1"},
			IgnoreCase:  &cacheKeyIgnoreCase,
		},

		// 离线模式配置
		OfflineMode: &offlineMode,

		// 强制HTTPS配置
		HttpToHttpsEnabled: &httpToHttpsEnabled,
		HttpToHttpsCode:    &httpToHttpsCode,

		// HSTS配置
		Hsts: &api.HSTS{
			MaxAge:            &hstsMaxAge,
			IncludeSubDomains: &hstsIncludeSubDomains,
			Preload:           &hstsPreload,
		},

		// HTTP2配置
		Http2Disable: &http2Disable,

		// HTTP3配置
		Http3: &api.HTTP3{
			Enable: &http3Enable,
		},

		// 最大上传大小配置
		ClientMaxBodySize: &clientMaxBodySize,

		// 页面压缩配置
		Compress:            &compress,
		CompressMethodArray: &[]string{"gzip", "br"},

		// 智能加速配置
		Isa: &isa,

		// 状态码缓存配置
		CacheCodeTtl: &[]api.CacheCodeTtl{
			{Value: "404", Weight: 100, OverrideOrigin: true, Ttl: 10, Type: "code"},
			{Value: "400", Weight: 100, OverrideOrigin: true, Ttl: 10, Type: "code"},
		},

		// gRPC回源配置
		GrpcOrigin: &grpcOrigin,

		// HTTP2回源配置
		Http2Origin: &http2Origin,

		// 规则引擎配置
		PageRules: &[]api.PageRule{
			{
				Name:   "test",
				Status: "ON",
				Rules: [][]api.Rule{
					{
						{
							MatchFrom: "host",
							Operator:  "inValues",
							Values:    []string{"1.ly.com"},
						},
					},
				},
				Config: &api.RuleConfig{
					CacheTtl: &[]api.CacheTtl{
						{Value: "/", Weight: 100, OverrideOrigin: true, Ttl: 3600, Type: "path"},
					},
					CacheKey: &api.RuleCacheKey{
						Query:       &ruleCacheKeyQuery,
						IncludeArgs: &[]string{"v"},
						IgnoreCase:  &ruleCacheKeyIgnoreCase,
						Headers:     &[]string{"Accept-Language"},
						Cookie:      &[]string{"uid"},
					},
					Http3: &api.HTTP3{Enable: &http3Enable},

					// 源站配置
					OriginConfig: &[]api.OriginItem{
						{
							Addr:             "1.2.3.4",
							Type:             "IP",
							UpstreamProtocol: "*",
							HttpPort:         &originHttpPort,
							HttpsPort:        &originHttpsPort,
							Host:             &originHost,
						},
					},
				},
			},
		},

		// 回源请求头配置
		OriginRequest: &api.OriginRequest{
			AddHeaders: &api.HeaderMap{
				"test":  "${uri}",
				"test2": "11",
			},
			RemoveHeaders: &api.HeaderMap{
				"Content-Type": "",
			},
		},

		// 节点响应头配置
		ClientResponse: &api.ClientResponse{
			AddHeaders: &api.HeaderMap{
				"test":  "${uri}",
				"test2": "11",
			},
			RemoveHeaders: &api.HeaderMap{
				"Content-Type": "",
			},
		},

		// 浏览器缓存TTL配置
		ClientCacheTtl: &api.ClientCacheTtl{
			Mode: &clientCacheTtlMode,
			Ttl:  &clientCacheTtl,
		},

		// webSocket配置
		WebSocket: &api.WebSocket{
			Enabled: &webSocketEnabled,
			Timeout: &webSocketTimeout,
		},

		// SSL/TLS安全配置
		SslProtocols:  &[]string{"TLSv1.2", "TLSv1.3"},
		SslMode:       &sslModePtr,
		SslCipherList: &sslCipherListPtr,

		// OCSP配置
		Ocsp: &ocsp,

		// 用户ip获取配置
		RealIp: &api.RealIp{
			Enabled: &realIpEnabled,
			Name:    &realIpName,
		},
	})

	data, _ := json.Marshal(result)
	t.Logf("setSiteConfig: %s", string(data))
	checkClientErr(t, "SetSiteConfig", err)
}

// TestGetSiteConfig - query site config
func TestGetSiteConfig(t *testing.T) {
	cfg, err := testCli.GetSiteConfig(testAuthoritySite)

	// 单测里若只关注某项，可单独打印
	//data, _ := json.Marshal(cfg.CacheKey)
	//if err != nil {
	//	t.Logf("err:%+v\n", err)
	//	return
	//}
	//t.Logf("cacheKey: %s", string(data))

	data, _ := json.Marshal(cfg)
	t.Logf("siteConfig: %s", string(data))
	checkClientErr(t, "GetSiteConfig", err)
}
