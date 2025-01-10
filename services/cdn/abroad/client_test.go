package abroad

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"runtime"
	"testing"
	"time"

	"github.com/baidubce/bce-sdk-go/model"
	"github.com/baidubce/bce-sdk-go/services/cdn/abroad/api"
	"github.com/baidubce/bce-sdk-go/util"
	"github.com/baidubce/bce-sdk-go/util/log"
)

var (
	testCli    *Client
	testDomain string
)

func TestMain(m *testing.M) {
	_, f, _, _ := runtime.Caller(0)
	confPath := filepath.Join(filepath.Dir(f), "config.json")
	fp, err := os.Open(confPath)
	if err != nil {
		fmt.Printf("TestMain terminated. cant found config json file: %+v\n", confPath)
		return
	}

	type Conf struct {
		AK       string `json:"AK"`
		SK       string `json:"SK"`
		Endpoint string `json:"Endpoint"`

		TestDomain string `json:"TestDomain"`
	}

	decoder := json.NewDecoder(fp)
	confObj := &Conf{}
	_ = decoder.Decode(confObj)
	if len(confObj.AK) == 0 || len(confObj.SK) == 0 || len(confObj.TestDomain) == 0 {
		fmt.Printf("TestMain terminated. config json file of AK/SK/TestDomain not given: %+v\n", confPath)
		return
	}

	testDomain = confObj.TestDomain
	testCli, _ = NewClient(confObj.AK, confObj.SK, confObj.Endpoint)
	log.SetLogLevel(log.DEBUG)
	m.Run()
}

func TestClient_GetDomainConfig(t *testing.T) {
	domainConfig, err := testCli.GetDomainConfig(testDomain)
	if err != nil {
		t.Fatalf("GetDomainConfig for %s failed: %s", testDomain, err)
	}
	t.Logf("GetDomainConfig for %s got: %+v", testDomain, domainConfig)
}

func TestClient_SetDomainOrigin(t *testing.T) {
	originPeers := []api.OriginPeer{
		{
			Type:   "DOMAIN",
			Addr:   "test.badiu.com",
			Backup: false,
		},
	}
	err := testCli.SetDomainOrigin(testDomain, originPeers)
	if err != nil {
		t.Fatalf("SetDomainOrigin for %s failed: %s", testDomain, err)
	}
	t.Logf("SetDomainOrigin successfully.")
}

func TestClient_SetCacheTTL(t *testing.T) {
	cacheTTls := []api.CacheTTL{
		{
			Type:           "path",
			Value:          "/",
			TTL:            3600,
			Weight:         100,
			OverrideOrigin: true,
		},
	}

	err := testCli.SetCacheTTL(testDomain, cacheTTls)
	if err != nil {
		t.Fatalf("SetCacheTTL for %s failed: %s", testDomain, err)
	}
	t.Logf("SetCacheTTL successfully.")
}

func TestClient_SetCacheFullUrl(t *testing.T) {
	err := testCli.SetCacheFullUrl(testDomain, true)
	if err != nil {
		t.Fatalf("SetCacheFullUrl for %s failed: %s", testDomain, err)
	}
	t.Logf("SetCacheFullUrl successfully.")
}

func TestClient_SetHostToOrigin(t *testing.T) {
	err := testCli.SetHostToOrigin(testDomain, "test.baidu.com")
	if err != nil {
		t.Fatalf("SetHostToOrigin for %s failed: %s", testDomain, err)
	}
	t.Logf("SetHostToOrigin successfully.")
}

func TestClient_SetRefererACL(t *testing.T) {
	err := testCli.SetRefererACL(testDomain, &api.RefererACL{
		BlackList:  []string{"bad.baidu.com"},
		AllowEmpty: true,
	})
	if err != nil {
		t.Fatalf("SetRefererACL for %s failed: %s", testDomain, err)
	}
	t.Logf("SetRefererACL successfully.")
}

func TestClient_SetIpACL(t *testing.T) {
	err := testCli.SetIpACL(testDomain, &api.IpACL{
		BlackList: []string{"220.181.38.148"},
	})
	if err != nil {
		t.Fatalf("SetIpACL for %s failed: %s", testDomain, err)
	}
	t.Logf("SetIpACL successfully.")
}

func TestClient_SetOriginProtocol(t *testing.T) {
	err := testCli.SetOriginProtocol(testDomain, api.HTTPSOrigin)
	if err != nil {
		t.Fatalf("SetOriginProtocol for %s failed: %s", testDomain, err)
	}
	t.Logf("SetOriginProtocol successfully.")
}

func TestClient_GetQuota(t *testing.T) {
	quota, err := testCli.GetQuota()
	if err != nil {
		t.Fatalf("GetQuota failed: %s", err)
	}
	t.Logf("GetQuota successfully: %+v", quota)
}

func TestClient_Purge(t *testing.T) {
	rawurl := fmt.Sprintf("http://%s/test/index.html", testDomain)
	purgedId, err := testCli.Purge([]api.PurgeTask{
		{
			Url: rawurl,
		},
	})
	if err != nil {
		t.Fatalf("Purge for %s failed: %s", rawurl, err)
	}
	t.Logf("Purge successfully: %+v", purgedId)
}

func TestClient_GetPurgedStatus(t *testing.T) {
	purgeId := "eJwFwUkNADAIBEBFJJRr4VcrlAT_EjpjypLquIHNOL5UM0s2_aiRQhZyhhlaLB_9OQr4"
	details, err := testCli.GetPurgedStatus(&api.CStatusQueryData{
		Id: purgeId,
	})
	if err != nil {
		t.Fatalf("GetPurgedStatus for %s failed: %s", purgeId, err)
	}
	t.Logf("GetPurgedStatus successfully: %+v", details)
}

func TestClient_Prefetch(t *testing.T) {
	rawurl := fmt.Sprintf("http://%s/test/index.html", testDomain)
	prefetchId, err := testCli.Prefetch([]api.PrefetchTask{
		{
			Url: rawurl,
		},
	})
	if err != nil {
		t.Fatalf("Prefetch for %s failed: %s", rawurl, err)
	}
	t.Logf("Prefetch successfully: %+v", prefetchId)
}

func TestClient_GetPrefetchStatus(t *testing.T) {
	prefetchId := "eJwFwcERACAIA7CJuBMK0qeriOL-I5g4hhER67EuNCHVR8UTW2hQ6SrazGBf-_4AC1U="
	details, err := testCli.GetPrefetchStatus(&api.CStatusQueryData{
		Id: prefetchId,
	})
	if err != nil {
		t.Fatalf("GetPrefetchStatus for %s failed: %s", prefetchId, err)
	}
	t.Logf("GetPrefetchStatus successfully: %+v", details)
}

func TestClient_GetDomainLog(t *testing.T) {
	endTs := time.Now().Unix()
	startTs := endTs - 24*60*60
	endTime := util.FormatISO8601Date(endTs)
	startTime := util.FormatISO8601Date(startTs)

	logEntries, err := testCli.GetDomainLog(testDomain, api.TimeInterval{
		StartTime: startTime,
		EndTime:   endTime,
	})
	if err != nil {
		t.Fatalf("GetDomainLog failed: %s", err)
	}
	t.Logf("GetDomainLog successfully: %+v", logEntries)
}

func TestClient_GetFlow(t *testing.T) {
	var details []api.FlowDetail
	var err error

	// 查询账户整体1小时纬度的整体带宽
	details, err = testCli.GetFlow(
		api.QueryStatByTimeRange("2024-04-15T00:00:00Z", "2024-04-16T00:00:00Z"),
		api.QueryStatByPeriod(api.Period3600))
	if err != nil {
		t.Fatalf("GetFlow failed: %s", err)
	}
	t.Logf("GetFlow successfully: %+v", details)

	// 查询特定域名在巴西的带宽
	details, err = testCli.GetFlow(
		api.QueryStatByCountry("BR"), // BR 是巴西的 GEC 代码
		api.QueryStatByTimeRange("2024-04-15T00:00:00Z", "2024-04-16T00:00:00Z"),
		api.QueryStatByDomains([]string{testDomain}),
		api.QueryStatByPeriod(api.Period3600))
	if err != nil {
		t.Fatalf("GetFlow failed: %s", err)
	}
	t.Logf("GetFlow successfully: %+v", details)
}

func TestClient_GetPv(t *testing.T) {
	var details []api.PvDetail
	var err error

	// 查询账户整体1小时纬度的整体PV
	details, err = testCli.GetPv(
		api.QueryStatByTimeRange("2024-04-15T00:00:00Z", "2024-04-16T00:00:00Z"),
		api.QueryStatByPeriod(api.Period3600))
	if err != nil {
		t.Fatalf("GetPv failed: %s", err)
	}
	t.Logf("GetPv successfully: %+v", details)

	// 查询特定域名在巴西的PV
	details, err = testCli.GetPv(
		api.QueryStatByCountry("BR"), // BR 是巴西的 GEC 代码
		api.QueryStatByTimeRange("2024-04-15T00:00:00Z", "2024-04-16T00:00:00Z"),
		api.QueryStatByDomains([]string{testDomain}),
		api.QueryStatByPeriod(api.Period3600))
	if err != nil {
		t.Fatalf("GetPv failed: %s", err)
	}
	t.Logf("GetPv successfully: %+v", details)
}

func TestClient_GetSrcFlow(t *testing.T) {
	var details []api.FlowDetail
	var err error

	// 查询账户整体1小时纬度的整体回源带宽
	details, err = testCli.GetSrcFlow(
		api.QueryStatByTimeRange("2024-04-15T00:00:00Z", "2024-04-16T00:00:00Z"),
		api.QueryStatByPeriod(api.Period3600))
	if err != nil {
		t.Fatalf("GetSrcFlow failed: %s", err)
	}
	t.Logf("GetSrcFlow successfully: %+v", details)

	// 查询特定域名的回源带宽
	details, err = testCli.GetSrcFlow(
		api.QueryStatByTimeRange("2024-04-15T00:00:00Z", "2024-04-16T00:00:00Z"),
		api.QueryStatByDomains([]string{testDomain}),
		api.QueryStatByPeriod(api.Period3600))
	if err != nil {
		t.Fatalf("GetSrcFlow failed: %s", err)
	}
	t.Logf("GetSrcFlow successfully: %+v", details)
}

func TestClient_GetHttpCode(t *testing.T) {
	var details []api.HttpCodeDetail
	var err error

	// 查询账户整体1小时纬度的整体状态码分布详情
	details, err = testCli.GetHttpCode(
		api.QueryStatByTimeRange("2024-04-15T00:00:00Z", "2024-04-16T00:00:00Z"),
		api.QueryStatByPeriod(api.Period3600))
	if err != nil {
		t.Fatalf("GetHttpCode failed: %s", err)
	}
	t.Logf("GetHttpCode successfully: %+v", details)

	// 查询特定域名的状态码分布详情
	details, err = testCli.GetHttpCode(
		api.QueryStatByTimeRange("2024-04-15T00:00:00Z", "2024-04-16T00:00:00Z"),
		api.QueryStatByDomains([]string{testDomain}),
		api.QueryStatByPeriod(api.Period3600))
	if err != nil {
		t.Fatalf("GetHttpCode failed: %s", err)
	}
	t.Logf("GetHttpCode successfully: %+v", details)
}

func TestClient_GethitRate(t *testing.T) {
	var details []api.HitDetail
	var err error

	// 查询账户整体1小时纬度的整体流量命中率详情
	details, err = testCli.GetRealHit(
		api.QueryStatByTimeRange("2024-04-15T00:00:00Z", "2024-04-16T00:00:00Z"),
		api.QueryStatByPeriod(api.Period3600))
	if err != nil {
		t.Fatalf("GetRealHit failed: %s", err)
	}
	t.Logf("GetRealHit successfully: %+v", details)

	// 查询特定域名的流量命中率详情
	details, err = testCli.GetRealHit(
		api.QueryStatByTimeRange("2024-04-15T00:00:00Z", "2024-04-16T00:00:00Z"),
		api.QueryStatByDomains([]string{testDomain}),
		api.QueryStatByPeriod(api.Period3600))
	if err != nil {
		t.Fatalf("GetRealHit failed: %s", err)
	}
	t.Logf("GetRealHit successfully: %+v", details)
}

func TestClient_SetHTTPSConfigWithOptions(t *testing.T) {
	var err error

	// 开启 HTTPS
	var certId = "cert-4xkhw3m73hxs"
	err = testCli.SetHTTPSConfigWithOptions(testDomain, true,
		api.HTTPSConfigCertID(certId),    // 必选
		api.HTTPSConfigRedirectWith301(), // 可选
		api.HTTPSConfigEnableH2(),        // 可选
	)
	if err != nil {
		t.Fatalf("SetHTTPSConfigWithOptions enable HTTPS failed: %s", err)
	}
	t.Logf("SetHTTPSConfigWithOptions enable HTTPS successfully")

	// 关闭 HTTPS
	err = testCli.SetHTTPSConfigWithOptions(testDomain, false)
	if err != nil {
		t.Fatalf("SetHTTPSConfigWithOptions disable HTTPS failed: %s", err)
	}
	t.Logf("SetHTTPSConfigWithOptions disable HTTPS successfully")
}

func TestClient_SetHTTPSConfig(t *testing.T) {
	var err error

	// 开启 HTTPS
	var certId = "cert-4xkhw3m73hxs"
	err = testCli.SetHTTPSConfig(testDomain, &api.HTTPSConfig{
		Enabled:      true,
		CertId:       certId,
		HttpRedirect: true,
		Http2Enabled: true,
	})
	if err != nil {
		t.Fatalf("SetHTTPSConfig enable HTTPS failed: %s", err)
	}
	t.Logf("SetHTTPSConfig enable HTTPS successfully")

	// 关闭 HTTPS
	err = testCli.SetHTTPSConfig(testDomain, &api.HTTPSConfig{
		Enabled: false,
	})
	if err != nil {
		t.Fatalf("SetHTTPSConfig disable HTTPS failed: %s", err)
	}
	t.Logf("SetHTTPSConfig disable HTTPS successfully")
}

func TestClient_ListDomains(t *testing.T) {
	domains, _, err := testCli.ListDomains("")
	if err != nil {
		t.Fatalf("ListDomains failed: %s", err)
	}
	t.Logf("ListDomains success: %v", domains)
}

func TestClient_ListDomainsInfo(t *testing.T) {
	domainsInfo, _, err := testCli.ListDomainInfos("")
	if err != nil {
		t.Fatalf("ListDomainsInfo failed: %s", err)
	}
	t.Logf("ListDomainsInfo success: %v", domainsInfo)
}

func TestClient_CreateDomainWithOptions(t *testing.T) {
	info, err := testCli.CreateDomainWithOptions(testDomain, []api.OriginPeer{
		{
			Type:   "IP",
			Backup: false,
			Addr:   "1.1.1.1",
		},
		{
			Type:   "IP",
			Backup: true,
			Addr:   "2.2.2.2",
		},
	}, CreateDomainWithTags([]model.TagModel{
		{
			TagKey:   "abroad",
			TagValue: "test",
		},
	}))
	if err != nil {
		t.Fatalf("CreateDomainWithOptions for %s failed: %s", testDomain, err)
	}

	t.Logf("CreateDomainWithOptions for %s success: %+v", testDomain, info)
}

func TestClient_CreateDomainForDynamic(t *testing.T) {
	info, err := testCli.CreateDomainWithOptions(testDomain, []api.OriginPeer{
		{
			Type:   "IP",
			Backup: false,
			Addr:   "1.6.7.8",
		},
	}, CreateDomainWithForm(api.DynamicDomainForm))
	if err != nil {
		t.Fatalf("CreateDomainForDynamic for %s failed: %s", testDomain, err)
	}

	t.Logf("CreateDomainForDynamic for %s success: %+v", testDomain, info)
}

func TestClient_EnableDomain(t *testing.T) {
	err := testCli.EnableDomain(testDomain)
	if err != nil {
		t.Fatalf("EnableDomain for %s failed: %s", testDomain, err)
	}
	t.Logf("EnableDomain for %s success", testDomain)
}

func TestClient_DisableDomain(t *testing.T) {
	err := testCli.DisableDomain(testDomain)
	if err != nil {
		t.Fatalf("DisableDomain for %s failed: %s", testDomain, err)
	}
	t.Logf("DisableDomain for %s success", testDomain)
}

func TestClient_DeleteDomain(t *testing.T) {
	err := testCli.DeleteDomain(testDomain)
	if err != nil {
		t.Fatalf("DeleteDomain for %s failed: %s", testDomain, err)
	}
	t.Logf("DeleteDomain for %s success", testDomain)
}

func TestSetTags(t *testing.T) {
	err := testCli.SetTags(testDomain, []model.TagModel{
		{
			TagKey:   "abroad",
			TagValue: "test",
		},
	})
	if err != nil {
		t.Fatalf("SetTags for %s failed: %s", testDomain, err)
	}
}

func TestGetTags(t *testing.T) {
	tags, err := testCli.GetTags(testDomain)
	if err != nil {
		t.Fatalf("GetTags for %s failed: %s", testDomain, err)
	}
	t.Logf("tags: %+v", tags)
}
