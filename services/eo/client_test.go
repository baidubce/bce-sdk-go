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
							Values:    []string{"1.test.com"},
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

// TestSetLimitConsumption - set usage-capping config for a site or a domain
func TestSetLimitConsumption(t *testing.T) {
	result, err := testCli.SetLimitConsumption(&api.LimitConsumption{
		Site:             testAuthoritySite,
		Object:           "site", // "site" 表示整站；域名则填具体域名
		Enabled:          false,
		Type:             "flow",
		Threshold:        "200GB",
		StatisticalCycle: 300,
		Action:           "stop",
	})

	data, _ := json.Marshal(result)
	t.Logf("setLimitConsumption: %s", string(data))
	checkClientErr(t, "SetLimitConsumption", err)
}

// TestGetLimitConsumption - query usage-capping rules of a site
func TestGetLimitConsumption(t *testing.T) {
	result, err := testCli.GetLimitConsumption(&api.LimitConsumptionQuery{
		Site:     testAuthoritySite,
		PageNo:   1,
		PageSize: 20,
	})

	data, _ := json.Marshal(result)
	t.Logf("getLimitConsumption: %s", string(data))
	checkClientErr(t, "GetLimitConsumption", err)
}

// TestDeleteLimitConsumption - delete usage-capping rule of a site or a domain
func TestDeleteLimitConsumption(t *testing.T) {
	result, err := testCli.DeleteLimitConsumption(&api.LimitConsumptionDelete{
		Site:   testAuthoritySite,
		Object: "site", // "site" 表示整站；域名则填具体域名
	})

	data, _ := json.Marshal(result)
	t.Logf("deleteLimitConsumption: %s", string(data))
	checkClientErr(t, "DeleteLimitConsumption", err)
}

// TestCreateOriginPool - create an origin pool for a site
func TestCreateOriginPool(t *testing.T) {
	originHost := "test.host.com"
	result, err := testCli.CreateOriginPool(testAuthoritySite, &api.OriginPoolConfig{
		Name: "test",
		OriginConfig: []api.OriginPoolItem{
			{
				Addr:            "your.bos.bucket1.name",
				Type:            "BUCKET",
				Host:            &originHost,
				Weight:          50,
				ThirdBucketAuth: &api.OriginPoolBucketAuth{AuthType: "bos", Enabled: false},
			},
			{
				Addr:   "1.2.3.4",
				Type:   "IP",
				Host:   &originHost,
				Weight: 50,
			},
		},
	})

	data, _ := json.Marshal(result)
	t.Logf("createOriginPool: %s", string(data))
	checkClientErr(t, "CreateOriginPool", err)
}

// TestQueryOriginPool - query the origin pools of a site
func TestQueryOriginPool(t *testing.T) {
	result, err := testCli.QueryOriginPool(&api.QueryOriginPoolRequest{
		Site:     testAuthoritySite,
		PageNo:   2,
		PageSize: 10,
	})

	data, _ := json.Marshal(result)
	t.Logf("queryOriginPool: %s", string(data))
	checkClientErr(t, "QueryOriginPool", err)
}

// TestUpdateOriginPool - update an origin pool of a site
func TestUpdateOriginPool(t *testing.T) {
	originHost := "test.host.com"
	result, err := testCli.UpdateOriginPool(testAuthoritySite, "geo_df9LiHmnB0dF", &api.OriginPoolConfig{
		Name: "test1217",
		OriginConfig: []api.OriginPoolItem{
			{
				Addr:            "your.bos.bucket1.name",
				Type:            "BUCKET",
				Host:            &originHost,
				Weight:          50,
				ThirdBucketAuth: &api.OriginPoolBucketAuth{AuthType: "bos", Enabled: false},
			},
			{
				Addr:   "1.2.3.4",
				Type:   "IP",
				Host:   &originHost,
				Weight: 50,
			},
		},
	})

	data, _ := json.Marshal(result)
	t.Logf("updateOriginPool: %s", string(data))
	checkClientErr(t, "UpdateOriginPool", err)
}

// TestDeleteOriginPool - delete an origin pool of a site
func TestDeleteOriginPool(t *testing.T) {
	result, err := testCli.DeleteOriginPool(testAuthoritySite, "geo_df9LiHmnB0dF")

	data, _ := json.Marshal(result)
	t.Logf("deleteOriginPool: %s", string(data))
	checkClientErr(t, "DeleteOriginPool", err)
}

// TestCreateLoadBalance - create a load balance for a site
func TestCreateLoadBalance(t *testing.T) {
	result, err := testCli.CreateLoadBalance(&api.LoadBalanceRequest{
		SiteId:       1,
		Name:         "test",
		DispatchMode: "priority_mode", // 主备调度；加权轮询用 "weight_mode"
		SessionMode:  "IP",            // 会话保持，或 "OFF"
		DetectConfig: &api.DetectConfig{
			HealthCheckType:   "ICMP",
			ProbeInterval:     10,
			ProbeMinSuccesses: 5,
			ProbeMaxFails:     5,
			ProbeTimeout:      3,
		},
		OriginPoolIds:        []string{"geo_Wcizr850oayz"},
		FallbackOriginPoolId: "geo_xGen5sexyI9v",
	})

	data, _ := json.Marshal(result)
	t.Logf("createLoadBalance: %s", string(data))
	checkClientErr(t, "CreateLoadBalance", err)
}

// TestQueryLoadBalance - query the load balances of a site
func TestQueryLoadBalance(t *testing.T) {
	result, err := testCli.QueryLoadBalance(&api.QueryLoadBalanceRequest{
		SiteId: 1,
		//FetchAll: true, // 全量查询；分页查询传 PageNo/PageSize
		PageNo:   1,
		PageSize: 10,
	})

	data, _ := json.Marshal(result)
	t.Logf("queryLoadBalance: %s", string(data))
	checkClientErr(t, "QueryLoadBalance", err)
}

// TestUpdateLoadBalance - update a load balance of a site
func TestUpdateLoadBalance(t *testing.T) {
	result, err := testCli.UpdateLoadBalance("eo_lb_wn1XUeil7aCg", &api.LoadBalanceRequest{
		SiteId:       1,
		Name:         "test",
		DispatchMode: "weight_mode", // 加权轮询；主备调度用 "priority_mode"
		SessionMode:  "OFF",
		DetectConfig: &api.DetectConfig{
			ProbeInterval:     100,
			ProbeMinSuccesses: 5,
			ProbeMaxFails:     5,
			ProbeSucceedCode:  "2xx",
			ProbeUrl:          "/test",
			ProbeTimeout:      3,
			ProbeMethod:       "GET",
		},
		OriginPoolIds:        []string{"geo_z7IqfgPN1uYL"},
		Weights:              []int{10},
		FallbackOriginPoolId: "geo_xnPe2TVp8d8B",
	})

	data, _ := json.Marshal(result)
	t.Logf("updateLoadBalance: %s", string(data))
	checkClientErr(t, "UpdateLoadBalance", err)
}

// TestDeleteLoadBalance - delete a load balance of a site
func TestDeleteLoadBalance(t *testing.T) {
	result, err := testCli.DeleteLoadBalance("eo_lb_wn1XUeil7aCg", 933)

	data, _ := json.Marshal(result)
	t.Logf("deleteLoadBalance: %s", string(data))
	checkClientErr(t, "DeleteLoadBalance", err)
}

// TestCreateDomain - create an accelerated domain for a site
func TestCreateDomain(t *testing.T) {
	originHost := "test.eo.com"
	httpPort := 80

	t.Run("IP源站", func(t *testing.T) {
		result, err := testCli.CreateDomain(&api.CreateDomainRequest{
			SiteId: 1,
			Domain: "a.test.com",
			OriginConfig: []api.DomainOriginConfig{
				{
					Addr:             "1.2.3.4",
					Type:             "IP",
					UpstreamProtocol: "http",
					HttpPort:         &httpPort,
					Host:             &originHost,
				},
			},
		})
		data, _ := json.Marshal(result)
		t.Logf("createDomain(IP): %s", string(data))
		checkClientErr(t, "CreateDomain(IP)", err)
	})

	t.Run("源站池", func(t *testing.T) {
		result, err := testCli.CreateDomain(&api.CreateDomainRequest{
			SiteId:           1,
			Domain:           "b.test.com",
			OriginPoolId:     "geo_1qrUW3nppJIy",
			UpstreamProtocol: "http",
			HttpPort:         &httpPort,
		})
		data, _ := json.Marshal(result)
		t.Logf("createDomain(originPool): %s", string(data))
		checkClientErr(t, "CreateDomain(originPool)", err)
	})

	t.Run("负载均衡", func(t *testing.T) {
		result, err := testCli.CreateDomain(&api.CreateDomainRequest{
			SiteId:           1,
			Domain:           "c.test.com",
			LoadBalanceId:    "eo_lb_30RjuhwJ7anZ",
			UpstreamProtocol: "http",
			HttpPort:         &httpPort,
		})
		data, _ := json.Marshal(result)
		t.Logf("createDomain(loadBalance): %s", string(data))
		checkClientErr(t, "CreateDomain(loadBalance)", err)
	})
}

// TestQueryDomainConfig - query the configuration of a domain
func TestQueryDomainConfig(t *testing.T) {
	result, err := testCli.QueryDomainConfig("1.test.com")

	data, _ := json.Marshal(result)
	t.Logf("queryDomainConfig: %s", string(data))
	checkClientErr(t, "QueryDomainConfig", err)
}

// TestUpdateDomainConfig - update the configuration of a domain
func TestUpdateDomainConfig(t *testing.T) {
	originHost := "test.host.com"
	httpPort := 80
	result, err := testCli.UpdateDomainConfig(&api.UpdateDomainConfigRequest{
		Site:   "test.com",
		Domain: "1.test.com",
		Config: &api.DomainConfigUpdate{
			OriginConfig: []api.DomainOriginConfig{
				{
					Addr:             "1.2.3.5",
					Type:             "IP",
					Host:             &originHost,
					HttpPort:         &httpPort,
					UpstreamProtocol: "http",
				},
			},
			Ipv6Dispatch: &api.Ipv6Dispatch{Enable: "ON"},
		},
	})

	data, _ := json.Marshal(result)
	t.Logf("updateDomainConfig: %s", string(data))
	checkClientErr(t, "UpdateDomainConfig", err)
}

// TestQueryDomains - query the accelerated domain list
func TestQueryDomains(t *testing.T) {
	result, err := testCli.QueryDomains(&api.QueryDomainsRequest{
		SiteId:   1,
		PageNo:   1,
		PageSize: 10,
	})

	data, _ := json.Marshal(result)
	t.Logf("queryDomains: %s", string(data))
	checkClientErr(t, "QueryDomains", err)
}

// TestStartDomain - enable one or more accelerated domains
func TestStartDomain(t *testing.T) {
	t.Run("启用单个域名", func(t *testing.T) {
		result, err := testCli.StartDomain(&api.DomainBatchRequest{
			Domain: "1.test.com",
		})
		data, _ := json.Marshal(result)
		t.Logf("startDomain(single): %s", string(data))
		checkClientErr(t, "StartDomain(single)", err)
	})

	t.Run("批量启用域名", func(t *testing.T) {
		result, err := testCli.StartDomain(&api.DomainBatchRequest{
			Domains: []string{"1.test.com", "2.test.com"},
		})
		data, _ := json.Marshal(result)
		t.Logf("startDomain(batch): %s", string(data))
		checkClientErr(t, "StartDomain(batch)", err)
	})
}

// TestStopDomain - disable one or more accelerated domains
func TestStopDomain(t *testing.T) {
	t.Run("停用单个域名", func(t *testing.T) {
		result, err := testCli.StopDomain(&api.DomainBatchRequest{
			Domain: "1.test.com",
		})
		data, _ := json.Marshal(result)
		t.Logf("stopDomain(single): %s", string(data))
		checkClientErr(t, "StopDomain(single)", err)
	})

	t.Run("批量停用域名", func(t *testing.T) {
		result, err := testCli.StopDomain(&api.DomainBatchRequest{
			Domains: []string{"1.test.com", "2.test.com"},
		})
		data, _ := json.Marshal(result)
		t.Logf("stopDomain(batch): %s", string(data))
		checkClientErr(t, "StopDomain(batch)", err)
	})
}

// TestDeleteDomain - delete one or more accelerated domains
func TestDeleteDomain(t *testing.T) {
	t.Run("删除单个域名", func(t *testing.T) {
		result, err := testCli.DeleteDomain(&api.DomainBatchRequest{
			Domain: "1.test.com",
		})
		data, _ := json.Marshal(result)
		t.Logf("deleteDomain(single): %s", string(data))
		checkClientErr(t, "DeleteDomain(single)", err)
	})

	t.Run("批量删除域名", func(t *testing.T) {
		result, err := testCli.DeleteDomain(&api.DomainBatchRequest{
			Domains: []string{"1.test.com", "2.test.com"},
		})
		data, _ := json.Marshal(result)
		t.Logf("deleteDomain(batch): %s", string(data))
		checkClientErr(t, "DeleteDomain(batch)", err)
	})
}
