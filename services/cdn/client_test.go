package cdn

import (
	"encoding/json"
	"fmt"
	"testing"
	"time"

	"github.com/baidubce/bce-sdk-go/bce"
	"github.com/baidubce/bce-sdk-go/services/cdn/api"
	"github.com/baidubce/bce-sdk-go/util"
)

const (
	testAuthorityDomain = "your_valid_domain"
	testEndpoint        = "cdn.baidubce.com"
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

	if err := prepareForTest(testAuthorityDomain); err != nil {
		fmt.Printf("TestMain terminated, error:%s", err.Error())
		return
	}

	m.Run()
}

func prepareForTest(domain string) error {
	_, _ = testCli.CreateDomain(testAuthorityDomain, &api.OriginInit{
		Origin: []api.OriginPeer{
			{
				Peer: "1.2.3.4",
				Host: "1.2.3.4",
			},
		},
	})

	domainStatus, err := testCli.GetDomainStatus("ALL", "")
	if err != nil {
		return err
	}

	for _, item := range domainStatus {
		if item.Domain == domain {
			return nil
		}
	}

	return fmt.Errorf("prepare failed, invalid domain:%s", domain)
}

func checkClientErr(t *testing.T, funcName string, err error) {
	//time.Sleep(time.Second * 1)
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

////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
// Test function about operating domain.
////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
func TestListDomains(t *testing.T) {
	domains, _, err := testCli.ListDomains("")

	t.Logf("domains: %v", domains)
	checkClientErr(t, "ListDomains", err)
}

func TestGetDomainStatus(t *testing.T) {
	domainStatus, err := testCli.GetDomainStatus("ALL", "")

	t.Logf("domainStatus: %v", domainStatus)
	checkClientErr(t, "GetDomainStatus", err)
}

func TestIsValidDomain(t *testing.T) {
	domainValidInfo, err := testCli.IsValidDomain(testAuthorityDomain)

	t.Logf("domainValidInfo: %v", domainValidInfo)
	checkClientErr(t, "IsValidDomain", err)
}

func TestCreateDomain(t *testing.T) {
	domainCreatedInfo, err := testCli.CreateDomain(testAuthorityDomain, &api.OriginInit{
		Origin: []api.OriginPeer{
			{
				Peer: "1.2.3.4",
				Host: "1.2.3.4",
			},
		},
	})

	t.Logf("domainCreatedInfo: %v", domainCreatedInfo)
	checkClientErr(t, "CreateDomain", err)
}

func TestDisableDomain(t *testing.T) {
	err := testCli.DisableDomain(testAuthorityDomain)
	checkClientErr(t, "DisableDomain", err)
}

func TestEnableDomain(t *testing.T) {
	err := testCli.EnableDomain(testAuthorityDomain)
	checkClientErr(t, "EnableDomain", err)
}

// ignore delete
//func TestDeleteDomain(t *testing.T) {
//	err := testCli.DeleteDomain(testAuthorityDomain)
//	checkClientErr(t, "TestDeleteDomain", err)
//}

func TestGetIpInfo(t *testing.T) {
	ipInfo, err := testCli.GetIpInfo("1.2.3.4", "describeIp")
	t.Logf("ipInfo: %v", ipInfo)
	checkClientErr(t, "GetIpInfo", err)
}

////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
// Test function about CRUD domain config.
////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
func TestGetDomainConfig(t *testing.T) {
	domainConfig, err := testCli.GetDomainConfig(testAuthorityDomain)

	data, _ := json.Marshal(domainConfig)
	t.Logf("domainConfig: %s", string(data))
	checkClientErr(t, "GetDomainConfig", err)
}

func TestSetDomainOrigin(t *testing.T) {
	err := testCli.SetDomainOrigin(testAuthorityDomain, []api.OriginPeer{
		{
			Peer:      "1.1.1.1",
			Host:      "www.baidu.com",
			Backup:    true,
			Follow302: true,
		},
		{
			Peer:      "http://2.2.2.2",
			Host:      "www.baidu.com",
			Backup:    false,
			Follow302: true,
		},
	}, "www.baidu.com")

	checkClientErr(t, "SetDomainOrigin", err)
}

func TestSetDomainSeo(t *testing.T) {
	err := testCli.SetDomainSeo(testAuthorityDomain, &api.SeoSwitch{
		DirectlyOrigin: "ON",
		PushRecord:     "OFF",
	})

	checkClientErr(t, "SetDomainSeo", err)
}

func TestGetDomainSeo(t *testing.T) {
	seoSwitch, err := testCli.GetDomainSeo(testAuthorityDomain)

	data, _ := json.Marshal(seoSwitch)
	t.Logf("seoSwitch: %s", string(data))
	checkClientErr(t, "GetDomainSeo", err)
}

func TestGetCacheTTL(t *testing.T) {
	cacheTTls, err := testCli.GetCacheTTL(testAuthorityDomain)

	data, _ := json.Marshal(cacheTTls)
	t.Logf("cacheTTls: %s", string(data))
	checkClientErr(t, "GetCacheTTL", err)
}

func TestSetCacheTTL(t *testing.T) {
	err := testCli.SetCacheTTL(testAuthorityDomain, []api.CacheTTL{
		{
			Type:   "suffix",
			Value:  ".jpg",
			TTL:    420000,
			Weight: 30,
		},
		{
			Type:  "suffix",
			Value: ".mp4",
			TTL:   10000,
		},
	})

	checkClientErr(t, "SetCacheTTL", err)
}

func TestSetRefererACL(t *testing.T) {
	// set white referer list
	err := testCli.SetRefererACL(testAuthorityDomain, nil, []string{
		"a.bbbbbb.c",
		"*.baidu.com.*",
	}, true)

	checkClientErr(t, "SetRefererACL", err)

	// set black referer list
	err = testCli.SetRefererACL(testAuthorityDomain, []string{
		"a.b.c",
		"*.xxxxx.com.*",
	}, nil, true)
	checkClientErr(t, "SetRefererACL", err)
}

func TestGetRefererACL(t *testing.T) {
	refererACL, err := testCli.GetRefererACL(testAuthorityDomain)
	data, _ := json.Marshal(refererACL)
	t.Logf("refererACL: %s", string(data))
	checkClientErr(t, "GetRefererACL", err)
}

func TestSetIpACL(t *testing.T) {
	err := testCli.SetIpACL(testAuthorityDomain, []string{
		"5.5.5.5",
		"6.6.6.6",
	}, nil)

	checkClientErr(t, "SetIpACL", err)

	err = testCli.SetIpACL(testAuthorityDomain, nil, []string{
		"1.2.3.4/24",
	})

	checkClientErr(t, "SetIpACL", err)
}

func TestGetIpACL(t *testing.T) {
	ipACL, err := testCli.GetIpACL(testAuthorityDomain)
	data, _ := json.Marshal(ipACL)
	t.Logf("ipACL: %s", string(data))
	checkClientErr(t, "GetIpACL", err)
}

func TestSetLimitRate(t *testing.T) {
	err := testCli.SetLimitRate(testAuthorityDomain, 1024)
	checkClientErr(t, "SetLimitRate", err)
}

func TestSetDomainHttps(t *testing.T) {
	err := testCli.SetDomainHttps(testAuthorityDomain, &api.HTTPSConfig{
		Enabled:          true,
		CertId:           "ssl-xxxxxx",
		Http2Enabled:     true,
		HttpRedirect:     true,
		HttpRedirectCode: 301,
	})

	checkClientErr(t, "SetDomainHttps", err)
}

func TestSetDomainRequestAuth(t *testing.T) {
	err := testCli.SetDomainRequestAuth(testAuthorityDomain, &api.RequestAuth{
		Type:    "c",
		Key1:    "secretekey1",
		Key2:    "secretekey2",
		Timeout: 300,
		WhiteList: []string{
			"/crossdomain.xml",
		},
		SignArg: "sign",
		TimeArg: "t",
	})

	checkClientErr(t, "SetDomainRequestAuth", err)
}

func TestSetFollowProtocol(t *testing.T) {
	err := testCli.SetFollowProtocol(testAuthorityDomain, true)
	checkClientErr(t, "SetFollowProtocol", err)
}

func TestSetHttpHeader(t *testing.T) {
	err := testCli.SetHttpHeader(testAuthorityDomain, []api.HttpHeader{
		{
			Type:   "origin",
			Header: "x-auth-cn",
			Value:  "xxxxxxxxx",
			Action: "remove",
		},
		{
			Type:   "response",
			Header: "content-type",
			Value:  "application/octet-stream",
			Action: "add",
		},
	})

	checkClientErr(t, "SetHttpHeader", err)
}

func TestGetHttpHerder(t *testing.T) {
	headers, err := testCli.GetHttpHeader(testAuthorityDomain)

	data, _ := json.Marshal(headers)
	t.Logf("headers: %s", string(data))
	checkClientErr(t, "GetHttpHeader", err)
}

func TestSetErrorPage(t *testing.T) {
	err := testCli.SetErrorPage(testAuthorityDomain, []api.ErrorPage{
		{
			Code:         510,
			RedirectCode: 302,
			Url:          "/customer_404.html",
		},
		{
			Code: 403,
			Url:  "/custom_403.html",
		},
	})

	checkClientErr(t, "SetErrorPage", err)
}

func TestGetErrorPage(t *testing.T) {
	errorPages, err := testCli.GetErrorPage(testAuthorityDomain)

	data, _ := json.Marshal(errorPages)
	t.Logf("errorPages: %s", string(data))
	checkClientErr(t, "GetErrorPage", err)
}

func TestSetMediaDrag(t *testing.T) {
	err := testCli.SetMediaDrag(testAuthorityDomain, &api.MediaDragConf{
		Mp4: &api.MediaCfg{
			DragMode: "second",
			FileSuffix: []string{
				"mp4",
				"m4a",
				"m4z",
			},
			StartArgName: "startIndex",
		},
		Flv: &api.MediaCfg{
			DragMode:   "byteAV",
			FileSuffix: []string{},
		},
	})

	checkClientErr(t, "SetMediaDrag", err)
}

func TestGetMediaDrag(t *testing.T) {
	mediaDragConf, err := testCli.GetMediaDrag(testAuthorityDomain)

	data, _ := json.Marshal(mediaDragConf)
	t.Logf("mediaDragConf: %s", string(data))
	checkClientErr(t, "GetMediaDrag", err)
}

func TestSetFileTrim(t *testing.T) {
	err := testCli.SetFileTrim(testAuthorityDomain, true)
	checkClientErr(t, "SetFileTrim", err)
}

func TestGetFileTrim(t *testing.T) {
	fileTrim, err := testCli.GetFileTrim(testAuthorityDomain)

	t.Logf("fileTrim: %v", fileTrim)
	checkClientErr(t, "GetFiletrim", err)
}

func TestSetMobileAccess(t *testing.T) {
	err := testCli.SetMobileAccess(testAuthorityDomain, true)
	checkClientErr(t, "SetMobileAccess", err)
}

func TestGetMobileAccess(t *testing.T) {
	distinguishClient, err := testCli.GetMobileAccess(testAuthorityDomain)

	t.Logf("distinguishClient: %v", distinguishClient)
	checkClientErr(t, "GetMobileAccess", err)
}

func TestSetClientIp(t *testing.T) {
	err := testCli.SetClientIp(testAuthorityDomain, &api.ClientIp{
		Enabled: true,
		Name:    "X-Real-IP",
	})

	checkClientErr(t, "SetClientIp", err)
}

func TestGetClientIp(t *testing.T) {
	clientIp, err := testCli.GetClientIp(testAuthorityDomain)

	t.Logf("clientIp: %v", clientIp)
	checkClientErr(t, "GetClientIp", err)
}

func TestSetAccessLimit(t *testing.T) {
	err := testCli.SetAccessLimit(testAuthorityDomain, &api.AccessLimit{
		Enabled: true,
		Limit:   200,
	})

	checkClientErr(t, "SetAccessLimit", err)
}

func TestGetAccessLimit(t *testing.T) {
	accessLimit, err := testCli.GetAccessLimit(testAuthorityDomain)

	t.Logf("accessLimit: %v", accessLimit)
	checkClientErr(t, "GetAccessLimit", err)
}

func TestSetCacheUrlArgs(t *testing.T) {
	err := testCli.SetCacheUrlArgs(testAuthorityDomain, &api.CacheUrlArgs{
		CacheFullUrl: false,
		CacheUrlArgs: []string{"1", "2"},
	})

	checkClientErr(t, "SetCacheUrlArgs", err)
}

func TestGetCacheUrlArgs(t *testing.T) {
	cacheUrlArgs, err := testCli.GetCacheUrlArgs(testAuthorityDomain)

	data, _ := json.Marshal(cacheUrlArgs)
	t.Logf("cacheUrlArgs: %s", string(data))
	checkClientErr(t, "GetCacheUrlArgs", err)
}

func TestSetCors(t *testing.T) {
	err := testCli.SetCors(testAuthorityDomain, true, []string{
		"http://www.baidu.com",
		"http://*.bce.com",
	})

	checkClientErr(t, "SetCors", err)
}

func TestGetCors(t *testing.T) {
	cors, err := testCli.GetCors(testAuthorityDomain)

	data, _ := json.Marshal(cors)
	t.Logf("cors: %s", string(data))
	checkClientErr(t, "GetCors", err)
}

func TestSetRangeSwitch(t *testing.T) {
	err := testCli.SetRangeSwitch(testAuthorityDomain, false)

	checkClientErr(t, "SetRangeSwitch", err)
}

func TestGetRangeSwitch(t *testing.T) {
	rangeSwitch, err := testCli.GetRangeSwitch(testAuthorityDomain)

	t.Logf("rangeSwitch: %+v", rangeSwitch)
	checkClientErr(t, "GetRangeSwitch", err)
}

func TestSetContentEncoding(t *testing.T) {
	err := testCli.SetContentEncoding(testAuthorityDomain, true, "br")
	checkClientErr(t, "SetContentEncoding", err)
}

func TestGetContentEncoding(t *testing.T) {
	contentEncoding, err := testCli.GetContentEncoding(testAuthorityDomain)
	t.Logf("contentEncoding: %+v", contentEncoding)
	checkClientErr(t, "GetContentEncoding", err)
}

////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
// Test function about purge/prefetch.
////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
func TestPurge(t *testing.T) {
	purgedId, err := testCli.Purge([]api.PurgeTask{
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

func TestGetPurgedStatus(t *testing.T) {
	purgedStatus, err := testCli.GetPurgedStatus(nil)

	data, _ := json.Marshal(purgedStatus)
	t.Logf("purgedStatus: %s", string(data))
	checkClientErr(t, "GetPurgedStatus", err)
}

func TestPrefetch(t *testing.T) {
	prefetchId, err := testCli.Prefetch([]api.PrefetchTask{
		{
			Url: "http://my.domain.com/path/to/purge/1.data",
		},
	})

	t.Logf("prefetchId: %s", string(prefetchId))
	checkClientErr(t, "Prefetch", err)
}

func TestGetPrefetchStatus(t *testing.T) {
	prefetchStatus, err := testCli.GetPrefetchStatus(nil)

	data, _ := json.Marshal(prefetchStatus)
	t.Logf("prefetchStatus: %s", string(data))
	checkClientErr(t, "GetPrefetchStatus", err)
}

func TestGetQuota(t *testing.T) {
	quotaDetail, err := testCli.GetQuota()

	data, _ := json.Marshal(quotaDetail)
	t.Logf("quotaDetail: %s", string(data))
	checkClientErr(t, "GetQuota", err)
}

func TestGetCacheOpRecords(t *testing.T) {
	recordDetails, err := testCli.GetCacheOpRecords(&api.CRecordQueryData{
		StartTime: "2019-08-12T12:00:00Z",
		EndTime:   "2019-08-14T12:00:00Z",
	})
	data, _ := json.Marshal(recordDetails)
	t.Logf("GetCacheOpRecords: %s", string(data))
	checkClientErr(t, "GetCacheOpRecords", err)
}

////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
// Test function about DSA.
////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
func TestEnableDsa(t *testing.T) {
	err := testCli.EnableDsa()
	checkClientErr(t, "EnableDsa", err)
}

func TestDisableDsa(t *testing.T) {
	err := testCli.DisableDsa()
	checkClientErr(t, "DisableDsa", err)
}

func TestListDsaDomains(t *testing.T) {
	dsaDomains, err := testCli.ListDsaDomains()
	data, _ := json.Marshal(dsaDomains)
	fmt.Println(string(data))
	checkClientErr(t, "ListDsaDomains", err)
}

func TestSetDsaConfig(t *testing.T) {
	err := testCli.SetDsaConfig(testAuthorityDomain, &api.DSAConfig{
		Enabled: true,
		Rules: []api.DSARule{
			{
				Type:  "suffix",
				Value: ".mp4;.jpg;.php",
			},
			{
				Type:  "path",
				Value: "/path",
			},
			{
				Type:  "exactPath",
				Value: "/path/to/file.mp4",
			},
		},
		Comment: "test",
	})

	checkClientErr(t, "SetDsaConfig", err)
}

////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
// Test function about log.
////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
func TestGetDomainLog(t *testing.T) {
	endTs := time.Now().Unix()
	startTs := endTs - 24*60*60
	endTime := util.FormatISO8601Date(endTs)
	startTime := util.FormatISO8601Date(startTs)
	domainLogs, err := testCli.GetDomainLog(testAuthorityDomain, api.TimeInterval{
		StartTime: startTime,
		EndTime:   endTime,
	})

	data, _ := json.Marshal(domainLogs)
	t.Logf("domainLogs: %s", string(data))
	checkClientErr(t, "GetDomainLog", err)
}

func TestGetMultiDomainLog(t *testing.T) {
	endTs := time.Now().Unix()
	startTs := endTs - 24*60*60
	endTime := util.FormatISO8601Date(endTs)
	startTime := util.FormatISO8601Date(startTs)

	domainLogs, err := testCli.GetMultiDomainLog(&api.LogQueryData{
		TimeInterval: api.TimeInterval{
			StartTime: startTime,
			EndTime:   endTime,
		},
		Type:    1,
		Domains: []string{"1.baidu.com", "2.baidu.com"},
	})

	data, _ := json.Marshal(domainLogs)
	t.Logf("domainLogs: %s", string(data))
	checkClientErr(t, "GetMultiDomainLog", err)
}

////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
// Test function about query statistics.
////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
func TestGetAvgSpeed(t *testing.T) {
	queryCondition := &api.QueryCondition{}
	avgSpeedDetails, err := testCli.GetAvgSpeed(queryCondition)

	data, _ := json.Marshal(avgSpeedDetails)
	t.Logf("avgSpeedDetails: %s", string(data))
	checkClientErr(t, "GetAvgSpeed", err)
}

func TestGetAvgSpeedByRegion(t *testing.T) {
	queryCondition := &api.QueryCondition{}
	avgSpeedDetails, err := testCli.GetAvgSpeedByRegion(queryCondition, "beijing", "")

	data, _ := json.Marshal(avgSpeedDetails)
	t.Logf("avgSpeedDetails: %s", string(data))
	checkClientErr(t, "GetAvgSpeedByRegion", err)
}

func TestGetPv(t *testing.T) {
	queryCondition := &api.QueryCondition{}
	pvDetails, err := testCli.GetPv(queryCondition, "all")

	data, _ := json.Marshal(pvDetails)
	t.Logf("pvDetails: %s", string(data))
	checkClientErr(t, "GetPv", err)
}

func TestGetSrcPv(t *testing.T) {
	queryCondition := &api.QueryCondition{}
	pvDetails, err := testCli.GetSrcPv(queryCondition)

	data, _ := json.Marshal(pvDetails)
	t.Logf("pvDetails: %s", string(data))
	checkClientErr(t, "GetSrcPv", err)
}

func TestGetPvInRegion(t *testing.T) {
	queryCondition := &api.QueryCondition{}
	pvRegionDetails, err := testCli.GetPvByRegion(queryCondition, "beijing", "")

	data, _ := json.Marshal(pvRegionDetails)
	t.Logf("pvRegionDetails: %s", string(data))
	checkClientErr(t, "GetPvByRegion", err)
}

func TestGetUv(t *testing.T) {
	queryCondition := &api.QueryCondition{}
	uvDetails, err := testCli.GetUv(queryCondition)

	data, _ := json.Marshal(uvDetails)
	t.Logf("uvDetails: %s", string(data))
	checkClientErr(t, "GetUv", err)
}

func TestGetFlow(t *testing.T) {
	queryCondition := &api.QueryCondition{
		StartTime: "2019-06-16T16:00:00Z",
		EndTime:   "2019-06-19T16:00:00Z",
		Period:    86400,
		KeyType:   0,
		Key:       []string{testAuthorityDomain},
		GroupBy:   "key",
	}
	flowDetails, err := testCli.GetFlow(queryCondition, "all")

	data, _ := json.Marshal(flowDetails)
	t.Logf("flowDetails: %s", string(data))
	checkClientErr(t, "GetFlow", err)
}

func TestGetFlowByProtocol(t *testing.T) {
	queryCondition := &api.QueryCondition{}
	flowDetails, err := testCli.GetFlowByProtocol(queryCondition, "all")

	data, _ := json.Marshal(flowDetails)
	t.Logf("flowDetails: %s", string(data))
	checkClientErr(t, "GetFlowByProtocol", err)
}

func TestGetFlowByRegion(t *testing.T) {
	queryCondition := &api.QueryCondition{}
	flowRegionDetails, err := testCli.GetFlowByRegion(queryCondition, "beijing", "")

	data, _ := json.Marshal(flowRegionDetails)
	t.Logf("flowRegionDetails: %s", string(data))
	checkClientErr(t, "GetFlowByRegion", err)
}

func TestGetSrcFlow(t *testing.T) {
	queryCondition := &api.QueryCondition{}
	flowDetails, err := testCli.GetSrcFlow(queryCondition)

	data, _ := json.Marshal(flowDetails)
	t.Logf("flowDetails: %s", string(data))
	checkClientErr(t, "GetFlowByRegion", err)
}

func TestGetRealHit(t *testing.T) {
	queryCondition := &api.QueryCondition{}
	hitDetails, err := testCli.GetRealHit(queryCondition)

	data, _ := json.Marshal(hitDetails)
	t.Logf("hitDetails: %s", string(data))
	checkClientErr(t, "GetRealHit", err)
}

func TestGetPvHit(t *testing.T) {
	queryCondition := &api.QueryCondition{}
	hitDetails, err := testCli.GetPvHit(queryCondition)

	data, _ := json.Marshal(hitDetails)
	t.Logf("hitDetails: %s", string(data))
	checkClientErr(t, "GetPvHit", err)
}

func TestGetHttpCode(t *testing.T) {
	queryCondition := &api.QueryCondition{}
	httpCodeDetails, err := testCli.GetHttpCode(queryCondition)

	data, _ := json.Marshal(httpCodeDetails)
	t.Logf("httpCodeDetails: %s", string(data))
	checkClientErr(t, "GetHttpCode", err)
}

func TestGetSrcHttpCode(t *testing.T) {
	queryCondition := &api.QueryCondition{}
	httpCodeDetails, err := testCli.GetSrcHttpCode(queryCondition)

	data, _ := json.Marshal(httpCodeDetails)
	t.Logf("httpCodeDetails: %s", string(data))
	checkClientErr(t, "GetSrcHttpCode", err)
}

func TestGetHttpCodeByRegion(t *testing.T) {
	queryCondition := &api.QueryCondition{}
	httpCodeDetails, err := testCli.GetHttpCodeByRegion(queryCondition, "beijing", "")

	data, _ := json.Marshal(httpCodeDetails)
	t.Logf("httpCodeDetails: %s", string(data))
	checkClientErr(t, "GetHttpCodeByRegion", err)
}

func TestGetTopNUrls(t *testing.T) {
	queryCondition := &api.QueryCondition{}
	topNUrls, err := testCli.GetTopNUrls(queryCondition, "200")

	data, _ := json.Marshal(topNUrls)
	t.Logf("topNUrls: %s", string(data))
	checkClientErr(t, "GetTopNUrls", err)
}

func TestGetTopNReferers(t *testing.T) {
	queryCondition := &api.QueryCondition{}
	topNReferers, err := testCli.GetTopNReferers(queryCondition, "")

	data, _ := json.Marshal(topNReferers)
	t.Logf("topNReferers: %s", string(data))
	checkClientErr(t, "GetTopNReferers", err)
}

func TestGetTopNDomains(t *testing.T) {
	queryCondition := &api.QueryCondition{}
	topNDomains, err := testCli.GetTopNDomains(queryCondition, "")

	data, _ := json.Marshal(topNDomains)
	t.Logf("topNDomains: %s", string(data))
	checkClientErr(t, "GetTopNDomains", err)
}

func TestGetError(t *testing.T) {
	queryCondition := &api.QueryCondition{}
	errorDetails, err := testCli.GetError(queryCondition)

	data, _ := json.Marshal(errorDetails)
	t.Logf("errorDetails: %s", string(data))
	checkClientErr(t, "GetErrorCount", err)
}

func TestGetPeak95Bandwidth(t *testing.T) {
	peak95Time, peak95Band, err := testCli.GetPeak95Bandwidth(
		"2020-05-01T00:00:00Z", "2020-05-10T00:00:00Z", []string{"www.test.com"}, nil)
	t.Logf("peak95Time %s, peak95Band %d", peak95Time, peak95Band)
	checkClientErr(t, "TestGetPeak95Bandwidth", err)
}
