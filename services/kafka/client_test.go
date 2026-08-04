/*
 * Copyright 2026 Baidu, Inc.
 *
 * Licensed under the Apache License, Version 2.0 (the "License"); you may not use this file
 * except in compliance with the License. You may obtain a copy of the License at
 *
 * http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software distributed under the
 * License is distributed on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND,
 * either express or implied. See the License for the specific language governing permissions
 * and limitations under the License.
 */

package kafka

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"strings"
	"testing"
	"time"

	"github.com/baidubce/bce-sdk-go/auth"
	"github.com/baidubce/bce-sdk-go/bce"
	"github.com/baidubce/bce-sdk-go/util"
)

func newMockKafkaClient(t *testing.T, responseBody string) *Client {
	t.Helper()
	httpClient := util.NewMockHTTPClient(
		util.SetStatusCode(http.StatusOK),
		util.SetStatusMsg(http.StatusText(http.StatusOK)),
		util.SetRespBody(responseBody),
	)
	client, err := NewClientWithConfig(&KafkaClientConfiguration{
		Endpoint: "192.168.1.1:8080",
		Credentials: &auth.BceCredentials{
			AccessKeyId:     "test-ak",
			SecretAccessKey: "1234567890abcdef-extra",
		},
		HTTPClient: httpClient,
		Retry:      bce.NewNoRetryPolicy(),
	})
	if err != nil {
		t.Fatalf("NewClientWithConfig returned error: %v", err)
	}
	return client
}

func kafkaInt(value int) *int       { return &value }
func kafkaInt64(value int64) *int64 { return &value }
func kafkaBool(value bool) *bool    { return &value }

func TestEncryptPasswordWithSecret(t *testing.T) {
	got, err := encryptPasswordWithSecret("password", "1234567890abcdef-extra")
	if err != nil {
		t.Fatalf("encryptPasswordWithSecret returned error: %v", err)
	}
	const want = "AE603C650117BBA793D852295CBAF8CB"
	if got != want {
		t.Fatalf("encryptPasswordWithSecret() = %q, want %q", got, want)
	}
}

func TestEncryptPasswordWithShortSecret(t *testing.T) {
	_, err := encryptPasswordWithSecret("password", "short")
	if err == nil {
		t.Fatal("encryptPasswordWithSecret() expected error for short secret")
	}
}

func TestCheckNotNilWithTypedNil(t *testing.T) {
	var revisionID *int
	if err := checkNotNil(revisionID, revisionIDKey); err == nil {
		t.Fatal("checkNotNil() expected error for typed nil pointer")
	}
}

func TestNewClientWithConfigAppliesKafkaDefaults(t *testing.T) {
	credentials, err := auth.NewBceCredentials("ak", "sk")
	if err != nil {
		t.Fatalf("NewBceCredentials returned error: %v", err)
	}
	config := &KafkaClientConfiguration{
		Credentials:               credentials,
		Endpoint:                  "custom-endpoint",
		ConnectionTimeoutInMillis: 321,
	}

	client, err := NewClientWithConfig(config)
	if err != nil {
		t.Fatalf("NewClientWithConfig returned error: %v", err)
	}

	if client.Config.Endpoint != "custom-endpoint" {
		t.Fatalf("Endpoint = %q, want %q", client.Config.Endpoint, "custom-endpoint")
	}
	if client.Config.Region != bce.DEFAULT_REGION {
		t.Fatalf("Region = %q, want %q", client.Config.Region, bce.DEFAULT_REGION)
	}
	if client.Config.UserAgent != bce.DEFAULT_USER_AGENT {
		t.Fatalf("UserAgent = %q, want %q", client.Config.UserAgent, bce.DEFAULT_USER_AGENT)
	}
	if client.Config.ConnectionTimeoutInMillis != 321 {
		t.Fatalf("ConnectionTimeoutInMillis = %d, want %d",
			client.Config.ConnectionTimeoutInMillis, 321)
	}
	if client.Config.Credentials != credentials {
		t.Fatal("Credentials were not taken from KafkaClientConfiguration")
	}
	if client.Config.SignOption == nil {
		t.Fatal("SignOption should be initialized")
	}
	if client.Config.Retry == nil {
		t.Fatal("Retry policy should be initialized")
	}
}

func TestNewClientWithConfigUsesKafkaSignHeaders(t *testing.T) {
	client, err := NewClientWithConfig(&KafkaClientConfiguration{})
	if err != nil {
		t.Fatalf("NewClientWithConfig returned error: %v", err)
	}
	for _, header := range []string{"host", "x-bce-date"} {
		if _, ok := client.Config.SignOption.HeadersToSign[header]; !ok {
			t.Errorf("HeadersToSign missing %q: %v", header, client.Config.SignOption.HeadersToSign)
		}
	}
	if client.Config.SignOption.ExpireSeconds != auth.DEFAULT_EXPIRE_SECONDS {
		t.Fatalf("ExpireSeconds = %d, want %d", client.Config.SignOption.ExpireSeconds, auth.DEFAULT_EXPIRE_SECONDS)
	}
}

func TestNewClientWithConfigBuildsEndpointFromRegion(t *testing.T) {
	client, err := NewClientWithConfig(&KafkaClientConfiguration{Region: "gz"})
	if err != nil {
		t.Fatalf("NewClientWithConfig returned error: %v", err)
	}

	if client.Config.Endpoint != "kafka.gz.baidubce.com" {
		t.Fatalf("Endpoint = %q, want %q", client.Config.Endpoint, "kafka.gz.baidubce.com")
	}
}

func TestNewClientWithConfigUsesJavaTimeoutDefaults(t *testing.T) {
	client, err := NewClientWithConfig(&KafkaClientConfiguration{})
	if err != nil {
		t.Fatalf("NewClientWithConfig returned error: %v", err)
	}

	if client.Config.ConnectionTimeoutInMillis != 50*1000 {
		t.Fatalf("ConnectionTimeoutInMillis = %d, want %d",
			client.Config.ConnectionTimeoutInMillis, 50*1000)
	}
	if client.Config.DialTimeout == nil || *client.Config.DialTimeout != 50*time.Second {
		t.Fatalf("DialTimeout = %v, want %s", client.Config.DialTimeout, 50*time.Second)
	}
	if client.Config.ReadTimeout == nil || *client.Config.ReadTimeout != 50*time.Second {
		t.Fatalf("ReadTimeout = %v, want %s", client.Config.ReadTimeout, 50*time.Second)
	}
}

func TestNewClientWithNilConfigReturnsError(t *testing.T) {
	if _, err := NewClientWithConfig(nil); err == nil {
		t.Fatal("NewClientWithConfig(nil) expected error")
	}
}

func TestNewClientWithConfigUsesConfiguredHTTPClient(t *testing.T) {
	credentials, err := auth.NewBceCredentials("ak", "sk")
	if err != nil {
		t.Fatalf("NewBceCredentials returned error: %v", err)
	}
	customHTTPClient := &http.Client{}

	client, err := NewClientWithConfig(&KafkaClientConfiguration{
		Credentials: credentials,
		HTTPClient:  customHTTPClient,
	})
	if err != nil {
		t.Fatalf("NewClientWithConfig returned error: %v", err)
	}
	if client.HTTPClient != customHTTPClient {
		t.Fatalf("HTTPClient = %p, want configured client %p", client.HTTPClient, customHTTPClient)
	}
}

func TestNewClientWithConfigUsesTransportAndRateLimitConfiguration(t *testing.T) {
	timeout := 7 * time.Second
	rateLimit := int64(128)

	client, err := NewClientWithConfig(&KafkaClientConfiguration{
		HTTPClientTimeout: &timeout,
		UploadRatelimit:   &rateLimit,
		DownloadRatelimit: &rateLimit,
	})
	if err != nil {
		t.Fatalf("NewClientWithConfig returned error: %v", err)
	}
	if client.HTTPClient == nil {
		t.Fatal("HTTPClient should be initialized from KafkaClientConfiguration")
	}
	if client.HTTPClient.Timeout != timeout {
		t.Fatalf("HTTPClient.Timeout = %s, want %s", client.HTTPClient.Timeout, timeout)
	}
	if client.Config.ConnectionTimeoutInMillis != 50*1000 {
		t.Fatalf("ConnectionTimeoutInMillis = %d, want %d",
			client.Config.ConnectionTimeoutInMillis, 50*1000)
	}
	if client.RateLimiters[bce.RateLimiterSlotTx] == nil {
		t.Fatal("upload rate limiter should be initialized")
	}
	if client.RateLimiters[bce.RateLimiterSlotRx] == nil {
		t.Fatal("download rate limiter should be initialized")
	}
}

type roundTripFunc func(*http.Request) (*http.Response, error)

func (fn roundTripFunc) RoundTrip(request *http.Request) (*http.Response, error) {
	return fn(request)
}

func TestListUsersAndListAclsMatchJavaOverloadSemantics(t *testing.T) {
	transport := roundTripFunc(func(request *http.Request) (*http.Response, error) {
		if request.Method != http.MethodGet {
			t.Errorf("method = %s, want GET", request.Method)
		}
		var body string
		switch request.URL.Path {
		case "/v2/clusters/cluster-1/users":
			if len(request.URL.Query()) != 0 {
				t.Errorf("user list query = %v, want empty query", request.URL.Query())
			}
			body = `{"users":[{"username":"alice"}]}`
		case "/v2/clusters/cluster-1/acls":
			query := request.URL.Query()
			if query.Get(usernameKey) != "alice" || query.Get(resourceTypeKey) != "TOPIC" {
				t.Errorf("ACL list query = %v, want username and resourceType filters", query)
			}
			if _, ok := query[patternTypeKey]; ok {
				t.Errorf("ACL list query unexpectedly contains empty patternType: %v", query)
			}
			body = `{"acls":[{"username":"alice","resourceType":"TOPIC"}]}`
		default:
			t.Fatalf("unexpected request path %q", request.URL.Path)
		}
		return &http.Response{
			StatusCode: http.StatusOK,
			Status:     "200 OK",
			Header:     make(http.Header),
			Body:       io.NopCloser(bytes.NewBufferString(body)),
			Request:    request,
		}, nil
	})

	client, err := NewClientWithConfig(&KafkaClientConfiguration{
		Endpoint:   "http://kafka.test",
		HTTPClient: &http.Client{Transport: transport},
	})
	if err != nil {
		t.Fatalf("NewClientWithConfig returned error: %v", err)
	}

	users, err := client.ListUsers("cluster-1")
	if err != nil {
		t.Fatalf("ListUsers(string) returned error: %v", err)
	}
	if len(users.Users) != 1 || users.Users[0].Username != "alice" {
		t.Fatalf("ListUsers(string) = %#v, want alice", users.Users)
	}

	acls, err := client.ListAcls(&ListAclRequest{
		ClusterID:    "cluster-1",
		Username:     "alice",
		ResourceType: "TOPIC",
	})
	if err != nil {
		t.Fatalf("ListAcls(ListAclRequest) returned error: %v", err)
	}
	if len(acls.Acls) != 1 || acls.Acls[0].Username != "alice" {
		t.Fatalf("ListAcls(ListAclRequest) = %#v, want alice", acls.Acls)
	}
}

func TestCreateUserUsesClientCredentialsForSigningAndEncryption(t *testing.T) {
	clientCredentials := &auth.BceCredentials{
		AccessKeyId:     "client-ak",
		SecretAccessKey: "client-secret-1234",
	}
	wantPassword, err := encryptPasswordWithSecret("password", clientCredentials.SecretAccessKey)
	if err != nil {
		t.Fatalf("encryptPasswordWithSecret returned error: %v", err)
	}

	transport := roundTripFunc(func(request *http.Request) (*http.Response, error) {
		if authorization := request.Header.Get("Authorization"); !bytes.Contains([]byte(authorization), []byte("/client-ak/")) {
			t.Errorf("Authorization = %q, want client access key", authorization)
		}
		var payload map[string]interface{}
		if err := json.NewDecoder(request.Body).Decode(&payload); err != nil {
			t.Fatalf("decode request body: %v", err)
		}
		if payload["password"] != wantPassword {
			t.Errorf("encrypted password = %v, want %q", payload["password"], wantPassword)
		}
		if _, ok := payload["requestCredentials"]; ok {
			t.Errorf("request credentials leaked into JSON: %v", payload)
		}
		return &http.Response{
			StatusCode: http.StatusOK,
			Status:     "200 OK",
			Header:     make(http.Header),
			Body:       io.NopCloser(bytes.NewBufferString(`{"username":"alice"}`)),
			Request:    request,
		}, nil
	})
	client, err := NewClientWithConfig(&KafkaClientConfiguration{
		Endpoint:    "http://kafka.test",
		Credentials: clientCredentials,
		HTTPClient:  &http.Client{Transport: transport},
	})
	if err != nil {
		t.Fatalf("NewClientWithConfig returned error: %v", err)
	}
	request := &CreateUserRequest{
		ClusterID: "cluster-1",
		Username:  "alice",
		Password:  "password",
	}
	response, err := client.CreateUser(request)
	if err != nil {
		t.Fatalf("CreateUser returned error: %v", err)
	}
	if response.Username != "alice" {
		t.Fatalf("Username = %q, want alice", response.Username)
	}
}

func TestListClustersPreservesEmptyTagValue(t *testing.T) {
	transport := roundTripFunc(func(request *http.Request) (*http.Response, error) {
		query := request.URL.Query()
		if query.Get(tagKeyKey) != "environment" {
			t.Errorf("tagKey = %q, want environment", query.Get(tagKeyKey))
		}
		values, ok := query[tagValueKey]
		if !ok || len(values) != 1 || values[0] != "" {
			t.Errorf("tagValue = %#v, want one explicit empty value", values)
		}
		return &http.Response{
			StatusCode: http.StatusOK,
			Status:     "200 OK",
			Header:     make(http.Header),
			Body:       io.NopCloser(bytes.NewBufferString(`{"clusters":[]}`)),
			Request:    request,
		}, nil
	})
	client, err := NewClientWithConfig(&KafkaClientConfiguration{
		Endpoint:   "http://kafka.test",
		HTTPClient: &http.Client{Transport: transport},
	})
	if err != nil {
		t.Fatalf("NewClientWithConfig returned error: %v", err)
	}
	empty := ""
	if _, err := client.ListClusters(&ListClustersRequest{TagKey: "environment", TagValue: &empty}); err != nil {
		t.Fatalf("ListClusters returned error: %v", err)
	}
	if _, err := client.ListClusters(&ListClustersRequest{TagKey: "environment"}); err == nil {
		t.Fatal("ListClusters expected error for nil tagValue")
	}
	value := "prod"
	if _, err := client.ListClusters(&ListClustersRequest{TagValue: &value}); err == nil {
		t.Fatal("ListClusters expected error for tagValue without tagKey")
	}
}

func TestListTopicPartitionsMatchesJavaPaginationSemantics(t *testing.T) {
	attempt := 0
	transport := roundTripFunc(func(request *http.Request) (*http.Response, error) {
		attempt++
		query := request.URL.Query()
		switch attempt {
		case 1:
			if query.Get(pageNoKey) != "1" || query.Get(pageSizeKey) != "10" {
				t.Errorf("pagination query = %v, want pageNo=1&pageSize=10", query)
			}
		case 2:
			if query.Get(pageNoKey) != "-1" || query.Get(pageSizeKey) != "20" {
				t.Errorf("pagination query = %v, want pageNo=-1&pageSize=20", query)
			}
		case 3:
			if query.Get(pageNoKey) != "0" {
				t.Errorf("pageNo = %q, want explicit 0", query.Get(pageNoKey))
			}
			if _, ok := query[pageSizeKey]; ok {
				t.Errorf("nonpositive pageSize unexpectedly sent: %v", query)
			}
		}
		return &http.Response{
			StatusCode: http.StatusOK,
			Status:     "200 OK",
			Header:     make(http.Header),
			Body:       io.NopCloser(bytes.NewBufferString(`{"partitions":[]}`)),
			Request:    request,
		}, nil
	})
	client, err := NewClientWithConfig(&KafkaClientConfiguration{
		Endpoint:   "http://kafka.test",
		HTTPClient: &http.Client{Transport: transport},
	})
	if err != nil {
		t.Fatalf("NewClientWithConfig returned error: %v", err)
	}
	base := ListTopicPartitionsRequest{ClusterID: "cluster-1", TopicName: "topic-1"}
	if _, err := client.ListTopicPartitions(&base); err != nil {
		t.Fatalf("ListTopicPartitions returned error: %v", err)
	}
	pageNo := -1
	pageSize := 20
	base.PageNo = &pageNo
	base.PageSize = &pageSize
	if _, err := client.ListTopicPartitions(&base); err != nil {
		t.Fatalf("ListTopicPartitions returned error: %v", err)
	}
	pageNo = 0
	pageSize = 0
	if _, err := client.ListTopicPartitions(&base); err != nil {
		t.Fatalf("ListTopicPartitions returned error: %v", err)
	}
}

func TestNewClientConvenienceConstructorsUseConfigPath(t *testing.T) {
	client, err := NewClient("ak", "sk", "")
	if err != nil {
		t.Fatalf("NewClient returned error: %v", err)
	}
	if client.Config.Endpoint != DEFAULT_ENDPOINT {
		t.Fatalf("Endpoint = %q, want %q", client.Config.Endpoint, DEFAULT_ENDPOINT)
	}
	if client.Config.Credentials.AccessKeyId != "ak" || client.Config.Credentials.SecretAccessKey != "sk" {
		t.Fatalf("Credentials = %#v, want ak/sk", client.Config.Credentials)
	}

	stsClient, err := NewClientWithSTS("ak", "sk", "token", "")
	if err != nil {
		t.Fatalf("NewClientWithSTS returned error: %v", err)
	}
	if stsClient.Config.Credentials.SessionToken != "token" {
		t.Fatalf("SessionToken = %q, want token", stsClient.Config.Credentials.SessionToken)
	}
}

func TestClient_CreateCluster(t *testing.T) {
	client := newMockKafkaClient(t, `{"clusterId":"cluster-1"}`)
	result, err := client.CreateCluster(&CreateClusterRequest{Name: "cluster", Type: TypeServerless})
	if err != nil {
		t.Fatalf("CreateCluster returned error: %v", err)
	}
	if result.ClusterID != "cluster-1" {
		t.Fatalf("ClusterID = %q, want cluster-1", result.ClusterID)
	}
	if _, err := client.CreateCluster(nil); err == nil {
		t.Fatal("CreateCluster(nil) expected error")
	}
}

func TestClient_DeleteCluster(t *testing.T) {
	client := newMockKafkaClient(t, `{}`)
	if _, err := client.DeleteCluster(&DeleteClusterRequest{ClusterID: "cluster-1"}); err != nil {
		t.Fatalf("DeleteCluster returned error: %v", err)
	}
}

func TestClient_ListClusters(t *testing.T) {
	client := newMockKafkaClient(t, `{"clusters":[{"clusterId":"cluster-1","name":"cluster"}]}`)
	tagValue := "prod"
	result, err := client.ListClusters(&ListClustersRequest{
		ListRequest: ListRequest{Marker: "next", MaxKeys: 20},
		ClusterName: "cluster", State: "ACTIVE", Mode: "HA", KafkaVersion: "3.6",
		Payment: "Postpaid", TagKey: "env", TagValue: &tagValue,
	})
	if err != nil {
		t.Fatalf("ListClusters returned error: %v", err)
	}
	if len(result.Clusters) != 1 || result.Clusters[0].ClusterID != "cluster-1" {
		t.Fatalf("Clusters = %#v, want cluster-1", result.Clusters)
	}
}

func TestClient_GetClusterDetail(t *testing.T) {
	client := newMockKafkaClient(t, `{"cluster":{"clusterId":"cluster-1","name":"cluster"}}`)
	result, err := client.GetClusterDetail(&GetClusterDetailRequest{ClusterID: "cluster-1"})
	if err != nil {
		t.Fatalf("GetClusterDetail returned error: %v", err)
	}
	if result.Cluster == nil || result.Cluster.ClusterID != "cluster-1" {
		t.Fatalf("Cluster = %#v, want cluster-1", result.Cluster)
	}
}

func TestClient_GetClusterDeletion(t *testing.T) {
	client := newMockKafkaClient(t, `{}`)
	if _, err := client.GetClusterDeletion(&GetClusterDeletionRequest{ClusterID: "cluster-1"}); err != nil {
		t.Fatalf("GetClusterDeletion returned error: %v", err)
	}
}

func TestClient_GetClusterAccessEndpoints(t *testing.T) {
	client := newMockKafkaClient(t, `{}`)
	if _, err := client.GetClusterAccessEndpoints(&GetClusterAccessEndpointsRequest{ClusterID: "cluster-1"}); err != nil {
		t.Fatalf("GetClusterAccessEndpoints returned error: %v", err)
	}
}

func TestClient_GetClusterNodes(t *testing.T) {
	client := newMockKafkaClient(t, `{"nodes":[{"nodeId":"node-1"}]}`)
	result, err := client.GetClusterNodes(&GetClusterNodesRequest{
		ListRequest: ListRequest{Marker: "next", MaxKeys: 20}, ClusterID: "cluster-1", State: "ACTIVE",
	})
	if err != nil {
		t.Fatalf("GetClusterNodes returned error: %v", err)
	}
	if len(result.Nodes) != 1 || result.Nodes[0].NodeID != "node-1" {
		t.Fatalf("Nodes = %#v, want node-1", result.Nodes)
	}
}

func TestClient_GetClusterConfigurations(t *testing.T) {
	client := newMockKafkaClient(t, `{}`)
	if _, err := client.GetClusterConfigurations(&GetClusterConfigurationsRequest{ClusterID: "cluster-1"}); err != nil {
		t.Fatalf("GetClusterConfigurations returned error: %v", err)
	}
}

func TestClient_IncreaseBrokerCount(t *testing.T) {
	client := newMockKafkaClient(t, `{}`)
	request := &IncreaseBrokerCountRequest{ClusterID: "cluster-1", NumberOfBrokerNodes: kafkaInt(6)}
	if _, err := client.IncreaseBrokerCount(request); err != nil {
		t.Fatalf("IncreaseBrokerCount returned error: %v", err)
	}
}

func TestClient_DecreaseBrokerCount(t *testing.T) {
	client := newMockKafkaClient(t, `{}`)
	request := &DecreaseBrokerCountRequest{ClusterID: "cluster-1", NumberOfBrokerNodes: kafkaInt(3)}
	if _, err := client.DecreaseBrokerCount(request); err != nil {
		t.Fatalf("DecreaseBrokerCount returned error: %v", err)
	}
}

func TestClient_MigrateClusterAz(t *testing.T) {
	client := newMockKafkaClient(t, `{}`)
	request := &MigrateClusterAzRequest{ClusterID: "cluster-1", ResizeType: kafkaInt(1)}
	if _, err := client.MigrateClusterAz(request); err != nil {
		t.Fatalf("MigrateClusterAz returned error: %v", err)
	}
}

func TestClient_UnifyClusterEndpoint(t *testing.T) {
	client := newMockKafkaClient(t, `{}`)
	request := &UnifyClusterEndpointRequest{ClusterID: "cluster-1", ActionID: "action-1"}
	if _, err := client.UnifyClusterEndpoint(request); err != nil {
		t.Fatalf("UnifyClusterEndpoint returned error: %v", err)
	}
}

func TestClient_UpdateBrokerNodeType(t *testing.T) {
	client := newMockKafkaClient(t, `{}`)
	request := &UpdateBrokerNodeTypeRequest{ClusterID: "cluster-1", NodeType: "kafka.c2.medium"}
	if _, err := client.UpdateBrokerNodeType(request); err != nil {
		t.Fatalf("UpdateBrokerNodeType returned error: %v", err)
	}
}

func TestClient_ExpandBrokerDiskCapacity(t *testing.T) {
	client := newMockKafkaClient(t, `{}`)
	request := &ExpandBrokerDiskCapacityRequest{ClusterID: "cluster-1", StorageSize: kafkaInt64(100)}
	if _, err := client.ExpandBrokerDiskCapacity(request); err != nil {
		t.Fatalf("ExpandBrokerDiskCapacity returned error: %v", err)
	}
}

func TestClient_UpdateAccessConfig(t *testing.T) {
	client := newMockKafkaClient(t, `{}`)
	request := &UpdateAccessConfigRequest{ClusterID: "cluster-1", ACLEnabled: kafkaBool(true)}
	if _, err := client.UpdateAccessConfig(request); err != nil {
		t.Fatalf("UpdateAccessConfig returned error: %v", err)
	}
}

func TestClient_StartCluster(t *testing.T) {
	client := newMockKafkaClient(t, `{}`)
	if _, err := client.StartCluster(&StartClusterRequest{ClusterID: "cluster-1"}); err != nil {
		t.Fatalf("StartCluster returned error: %v", err)
	}
}

func TestClient_StopCluster(t *testing.T) {
	client := newMockKafkaClient(t, `{}`)
	if _, err := client.StopCluster(&StopClusterRequest{ClusterID: "cluster-1"}); err != nil {
		t.Fatalf("StopCluster returned error: %v", err)
	}
}

func TestClient_ResizeClusterEipBandwidth(t *testing.T) {
	client := newMockKafkaClient(t, `{}`)
	request := &ResizeClusterEipBandwidthRequest{ClusterID: "cluster-1", PublicIPBandwidth: kafkaInt(10)}
	if _, err := client.ResizeClusterEipBandwidth(request); err != nil {
		t.Fatalf("ResizeClusterEipBandwidth returned error: %v", err)
	}
}

func TestClient_SwitchClusterEip(t *testing.T) {
	client := newMockKafkaClient(t, `{}`)
	request := &SwitchClusterEipRequest{ClusterID: "cluster-1", PublicIPEnabled: kafkaBool(true)}
	if _, err := client.SwitchClusterEip(request); err != nil {
		t.Fatalf("SwitchClusterEip returned error: %v", err)
	}
}

func TestClient_UpdateStoragePolicy(t *testing.T) {
	client := newMockKafkaClient(t, `{}`)
	request := &UpdateStoragePolicyRequest{ClusterID: "cluster-1", StoragePolicyEnabled: kafkaBool(false)}
	if _, err := client.UpdateStoragePolicy(request); err != nil {
		t.Fatalf("UpdateStoragePolicy returned error: %v", err)
	}
}

func TestClient_UpdateKafkaConfig(t *testing.T) {
	client := newMockKafkaClient(t, `{}`)
	request := &UpdateKafkaConfigRequest{ClusterID: "cluster-1", ConfigID: "config-1", RevisionID: kafkaInt(1)}
	if _, err := client.UpdateKafkaConfig(request); err != nil {
		t.Fatalf("UpdateKafkaConfig returned error: %v", err)
	}
}

func TestClient_UpdateSecurityGroup(t *testing.T) {
	client := newMockKafkaClient(t, `{}`)
	request := &UpdateSecurityGroupRequest{ClusterID: "cluster-1", SecurityGroupIds: []string{"sg-1"}}
	if _, err := client.UpdateSecurityGroup(request); err != nil {
		t.Fatalf("UpdateSecurityGroup returned error: %v", err)
	}
}

func TestClient_UpdateMaintenanceDuration(t *testing.T) {
	client := newMockKafkaClient(t, `{}`)
	request := &UpdateMaintenanceDurationRequest{
		ClusterID: "cluster-1", MaintenancePeriods: []MaintainPeriod{MaintainPeriodMonday},
	}
	if _, err := client.UpdateMaintenanceDuration(request); err != nil {
		t.Fatalf("UpdateMaintenanceDuration returned error: %v", err)
	}
}

func TestClient_SwitchClusterIntranetIp(t *testing.T) {
	client := newMockKafkaClient(t, `{}`)
	request := &SwitchClusterIntranetIpRequest{ClusterID: "cluster-1", IntranetIPEnabled: kafkaBool(true)}
	if _, err := client.SwitchClusterIntranetIp(request); err != nil {
		t.Fatalf("SwitchClusterIntranetIp returned error: %v", err)
	}
}

func TestClient_GetClusterCurrentController(t *testing.T) {
	client := newMockKafkaClient(t, `{}`)
	request := &GetClusterCurrentControllerRequest{ClusterID: "cluster-1"}
	if _, err := client.GetClusterCurrentController(request); err != nil {
		t.Fatalf("GetClusterCurrentController returned error: %v", err)
	}
}

func TestClient_GetClusterHistoryController(t *testing.T) {
	client := newMockKafkaClient(t, `{}`)
	request := &GetClusterHistoryControllerRequest{ClusterID: "cluster-1"}
	if _, err := client.GetClusterHistoryController(request); err != nil {
		t.Fatalf("GetClusterHistoryController returned error: %v", err)
	}
}

func TestClient_RestartCluster(t *testing.T) {
	client := newMockKafkaClient(t, `{}`)
	if _, err := client.RestartCluster(&RestartClusterRequest{ClusterID: "cluster-1"}); err != nil {
		t.Fatalf("RestartCluster returned error: %v", err)
	}
}

func TestClient_RestartBroker(t *testing.T) {
	client := newMockKafkaClient(t, `{}`)
	request := &RestartBrokerRequest{ClusterID: "cluster-1", NodeID: "node-1"}
	if _, err := client.RestartBroker(request); err != nil {
		t.Fatalf("RestartBroker returned error: %v", err)
	}
}

func TestClient_SwitchClusterAdvertisedIp(t *testing.T) {
	client := newMockKafkaClient(t, `{}`)
	request := &SwitchClusterAdvertisedIpRequest{ClusterID: "cluster-1", AdvertisedIPEnabled: kafkaBool(true)}
	if _, err := client.SwitchClusterAdvertisedIp(request); err != nil {
		t.Fatalf("SwitchClusterAdvertisedIp returned error: %v", err)
	}
}

func TestClient_SwitchClusterDomain(t *testing.T) {
	client := newMockKafkaClient(t, `{}`)
	request := &SwitchClusterDomainRequest{ClusterID: "cluster-1", DomainEnabled: kafkaBool(true)}
	if _, err := client.SwitchClusterDomain(request); err != nil {
		t.Fatalf("SwitchClusterDomain returned error: %v", err)
	}
}

func TestClient_GetZkPassword(t *testing.T) {
	client := newMockKafkaClient(t, `{"password":"password"}`)
	result, err := client.GetZkPassword(&GetZkPasswordRequest{ClusterID: "cluster-1"})
	if err != nil {
		t.Fatalf("GetZkPassword returned error: %v", err)
	}
	if result.Password != "password" {
		t.Fatalf("Password = %q, want password", result.Password)
	}
}

func TestClient_CreateClusterConfig(t *testing.T) {
	client := newMockKafkaClient(t, `{"configId":"config-1"}`)
	request := &CreateClusterConfigRequest{Name: "config", Context: map[string]string{}}
	result, err := client.CreateClusterConfig(request)
	if err != nil {
		t.Fatalf("CreateClusterConfig returned error: %v", err)
	}
	if result.ConfigID != "config-1" {
		t.Fatalf("ConfigID = %q, want config-1", result.ConfigID)
	}
}

func TestClient_ListClusterConfigs(t *testing.T) {
	client := newMockKafkaClient(t, `{}`)
	request := &ListClusterConfigsRequest{
		ListRequest: ListRequest{Marker: "next", MaxKeys: 20}, ConfigName: "config", State: "ACTIVE",
	}
	if _, err := client.ListClusterConfigs(request); err != nil {
		t.Fatalf("ListClusterConfigs returned error: %v", err)
	}
}

func TestClient_DeleteClusterConfig(t *testing.T) {
	client := newMockKafkaClient(t, `{}`)
	if _, err := client.DeleteClusterConfig(&DeleteClusterConfigRequest{ConfigID: "config-1"}); err != nil {
		t.Fatalf("DeleteClusterConfig returned error: %v", err)
	}
}

func TestClient_GetClusterConfig(t *testing.T) {
	client := newMockKafkaClient(t, `{"config":{"configId":"config-1"}}`)
	result, err := client.GetClusterConfig(&GetClusterConfigRequest{ConfigID: "config-1"})
	if err != nil {
		t.Fatalf("GetClusterConfig returned error: %v", err)
	}
	if result.Config == nil || result.Config.ConfigID != "config-1" {
		t.Fatalf("Config = %#v, want config-1", result.Config)
	}
}

func TestClient_CreateClusterConfigRevision(t *testing.T) {
	client := newMockKafkaClient(t, `{}`)
	request := &CreateClusterConfigRevisionRequest{ConfigID: "config-1", Context: map[string]string{}}
	if _, err := client.CreateClusterConfigRevision(request); err != nil {
		t.Fatalf("CreateClusterConfigRevision returned error: %v", err)
	}
}

func TestClient_ListClusterConfigRevisions(t *testing.T) {
	client := newMockKafkaClient(t, `{}`)
	request := &ListClusterConfigRevisionsRequest{ConfigID: "config-1", State: "ACTIVE"}
	if _, err := client.ListClusterConfigRevisions(request); err != nil {
		t.Fatalf("ListClusterConfigRevisions returned error: %v", err)
	}
}

func TestClient_GetClusterConfigRevision(t *testing.T) {
	client := newMockKafkaClient(t, `{}`)
	request := &GetClusterConfigRevisionRequest{ConfigID: "config-1", RevisionID: kafkaInt(2)}
	if _, err := client.GetClusterConfigRevision(request); err != nil {
		t.Fatalf("GetClusterConfigRevision returned error: %v", err)
	}
	request.RevisionID = nil
	if _, err := client.GetClusterConfigRevision(request); err == nil {
		t.Fatal("GetClusterConfigRevision expected error for nil revisionId")
	}
}

func TestClient_UpdateTopic(t *testing.T) {
	client := newMockKafkaClient(t, `{}`)
	request := &UpdateTopicRequest{ClusterID: "cluster-1", TopicName: "topic-1", PartitionNum: "3"}
	if _, err := client.UpdateTopic(request); err != nil {
		t.Fatalf("UpdateTopic returned error: %v", err)
	}
	request.PartitionNum = ""
	if _, err := client.UpdateTopic(request); err == nil {
		t.Fatal("UpdateTopic expected error when no update field is set")
	}
}

func TestClient_GetSubscribedGroupDetail(t *testing.T) {
	client := newMockKafkaClient(t, `{}`)
	request := &GetSubscribedGroupDetailRequest{ClusterID: "cluster-1", TopicName: "topic-1", GroupName: "group-1"}
	if _, err := client.GetSubscribedGroupDetail(request); err != nil {
		t.Fatalf("GetSubscribedGroupDetail returned error: %v", err)
	}
}

func TestClient_ListTopicPartitions(t *testing.T) {
	client := newMockKafkaClient(t, `{"partitions":[{"partitionId":0}]}`)
	request := &ListTopicPartitionsRequest{
		PageListRequest: PageListRequest{PageNo: kafkaInt(2), PageSize: kafkaInt(20)},
		ClusterID:       "cluster-1",
		TopicName:       "topic-1",
	}
	result, err := client.ListTopicPartitions(request)
	if err != nil {
		t.Fatalf("ListTopicPartitions returned error: %v", err)
	}
	if len(result.Partitions) != 1 {
		t.Fatalf("Partitions = %#v, want one partition", result.Partitions)
	}
}

func TestClient_GetTopicPartitionDetail(t *testing.T) {
	client := newMockKafkaClient(t, `{}`)
	request := &GetTopicPartitionDetailRequest{ClusterID: "cluster-1", TopicName: "topic-1", PartitionID: "0"}
	if _, err := client.GetTopicPartitionDetail(request); err != nil {
		t.Fatalf("GetTopicPartitionDetail returned error: %v", err)
	}
}

func TestClient_ListTopic(t *testing.T) {
	client := newMockKafkaClient(t, `{"topics":[{"topicName":"topic-1"}]}`)
	result, err := client.ListTopic(&ListTopicRequest{ClusterID: "cluster-1", TopicName: "topic-1"})
	if err != nil {
		t.Fatalf("ListTopic returned error: %v", err)
	}
	if len(result.Topics) != 1 || result.Topics[0].TopicName != "topic-1" {
		t.Fatalf("Topics = %#v, want topic-1", result.Topics)
	}
}

func TestClient_GetTopicDetail(t *testing.T) {
	client := newMockKafkaClient(t, `{"topic":{"topicName":"topic-1"}}`)
	result, err := client.GetTopicDetail(&GetTopicDetailRequest{ClusterID: "cluster-1", TopicName: "topic-1"})
	if err != nil {
		t.Fatalf("GetTopicDetail returned error: %v", err)
	}
	if result.Topic == nil || result.Topic.TopicName != "topic-1" {
		t.Fatalf("Topic = %#v, want topic-1", result.Topic)
	}
}

func TestClient_CreateTopic(t *testing.T) {
	client := newMockKafkaClient(t, `{"topicName":"topic-1"}`)
	request := &CreateTopicRequest{ClusterID: "cluster-1", TopicName: "topic-1", PartitionNum: 3, ReplicationFactor: 3}
	result, err := client.CreateTopic(request)
	if err != nil {
		t.Fatalf("CreateTopic returned error: %v", err)
	}
	if result.TopicName != "topic-1" {
		t.Fatalf("TopicName = %q, want topic-1", result.TopicName)
	}
}

func TestClient_ListSubscribedGroups(t *testing.T) {
	client := newMockKafkaClient(t, `{}`)
	request := &ListSubscribedGroupsRequest{ClusterID: "cluster-1", TopicName: "topic-1"}
	if _, err := client.ListSubscribedGroups(request); err != nil {
		t.Fatalf("ListSubscribedGroups returned error: %v", err)
	}
}

func TestClient_DeleteTopic(t *testing.T) {
	client := newMockKafkaClient(t, `{}`)
	request := &DeleteTopicRequest{ClusterID: "cluster-1", TopicName: "topic-1"}
	if _, err := client.DeleteTopic(request); err != nil {
		t.Fatalf("DeleteTopic returned error: %v", err)
	}
}

func TestClient_GetTopicPartitionOverview(t *testing.T) {
	client := newMockKafkaClient(t, `{}`)
	request := &GetTopicPartitionOverviewRequest{ClusterID: "cluster-1", TopicName: "topic-1"}
	if _, err := client.GetTopicPartitionOverview(request); err != nil {
		t.Fatalf("GetTopicPartitionOverview returned error: %v", err)
	}
}

func TestClient_GetSubscribedGroupOverview(t *testing.T) {
	client := newMockKafkaClient(t, `{}`)
	request := &GetSubscribedGroupOverviewRequest{ClusterID: "cluster-1", TopicName: "topic-1"}
	if _, err := client.GetSubscribedGroupOverview(request); err != nil {
		t.Fatalf("GetSubscribedGroupOverview returned error: %v", err)
	}
}

func TestClient_ListTopicConfigOptions(t *testing.T) {
	client := newMockKafkaClient(t, `{}`)
	if _, err := client.ListTopicConfigOptions(); err != nil {
		t.Fatalf("ListTopicConfigOptions returned error: %v", err)
	}
}

func TestClient_SendTopicMessage(t *testing.T) {
	client := newMockKafkaClient(t, `{}`)
	message := "message"
	request := &SendTopicMessageRequest{
		ClusterID: "cluster-1", TopicName: "topic-1", PartitionID: kafkaInt(1), Value: &message,
	}
	if _, err := client.SendTopicMessage(request); err != nil {
		t.Fatalf("SendTopicMessage returned error: %v", err)
	}
	request.Value = nil
	if _, err := client.SendTopicMessage(request); err == nil {
		t.Fatal("SendTopicMessage expected error for nil value")
	}
}

func TestClient_QueryTopicMessagesByStartTime(t *testing.T) {
	client := newMockKafkaClient(t, `{}`)
	request := &QueryTopicMessagesByStartTimeRequest{
		ClusterID: "cluster-1", TopicName: "topic-1", PartitionID: kafkaInt(1), StartTime: 100,
	}
	if _, err := client.QueryTopicMessagesByStartTime(request); err != nil {
		t.Fatalf("QueryTopicMessagesByStartTime returned error: %v", err)
	}
}

func TestClient_QueryTopicMessagesByStartOffset(t *testing.T) {
	client := newMockKafkaClient(t, `{}`)
	request := &QueryTopicMessagesByStartOffsetRequest{
		ClusterID: "cluster-1", TopicName: "topic-1", PartitionID: 1, StartOffset: 2,
	}
	if _, err := client.QueryTopicMessagesByStartOffset(request); err != nil {
		t.Fatalf("QueryTopicMessagesByStartOffset returned error: %v", err)
	}
}

func TestClient_ListSubscribedTopics(t *testing.T) {
	client := newMockKafkaClient(t, `{}`)
	request := &ListSubscribedTopicsRequest{ClusterID: "cluster-1", GroupName: "group-1"}
	if _, err := client.ListSubscribedTopics(request); err != nil {
		t.Fatalf("ListSubscribedTopics returned error: %v", err)
	}
}

func TestClient_ListConsumerGroup(t *testing.T) {
	client := newMockKafkaClient(t, `{"groups":[{"groupName":"group-1"}]}`)
	result, err := client.ListConsumerGroup(&ListConsumerGroupRequest{ClusterID: "cluster-1", GroupName: "group-1"})
	if err != nil {
		t.Fatalf("ListConsumerGroup returned error: %v", err)
	}
	if len(result.Groups) != 1 || result.Groups[0].GroupName != "group-1" {
		t.Fatalf("Groups = %#v, want group-1", result.Groups)
	}
}

func TestClient_DeleteConsumerGroup(t *testing.T) {
	client := newMockKafkaClient(t, `{}`)
	request := &DeleteConsumerGroupRequest{ClusterID: "cluster-1", GroupName: "group-1"}
	if _, err := client.DeleteConsumerGroup(request); err != nil {
		t.Fatalf("DeleteConsumerGroup returned error: %v", err)
	}
}

func TestClient_ResetConsumerGroup(t *testing.T) {
	client := newMockKafkaClient(t, `{}`)
	request := &ResetConsumerGroupRequest{
		ClusterID: "cluster-1", GroupName: "group-1", TopicName: "topic-1",
		Partitions: []int{0}, ResetStrategy: "EARLIEST",
	}
	if _, err := client.ResetConsumerGroup(request); err != nil {
		t.Fatalf("ResetConsumerGroup returned error: %v", err)
	}
	request.Partitions = nil
	if _, err := client.ResetConsumerGroup(request); err == nil {
		t.Fatal("ResetConsumerGroup expected error for empty partitions")
	}
}

func TestClient_GetSubscribedTopicOverview(t *testing.T) {
	client := newMockKafkaClient(t, `{}`)
	request := &GetSubscribedTopicOverviewRequest{ClusterID: "cluster-1", GroupName: "group-1"}
	if _, err := client.GetSubscribedTopicOverview(request); err != nil {
		t.Fatalf("GetSubscribedTopicOverview returned error: %v", err)
	}
}

func TestClient_CreateUser(t *testing.T) {
	client := newMockKafkaClient(t, `{"username":"alice"}`)
	request := &CreateUserRequest{ClusterID: "cluster-1", Username: "alice", Password: "Password9!"}
	result, err := client.CreateUser(request)
	if err != nil {
		t.Fatalf("CreateUser returned error: %v", err)
	}
	if result.Username != "alice" {
		t.Fatalf("Username = %q, want alice", result.Username)
	}
}

func TestClient_DeleteUser(t *testing.T) {
	client := newMockKafkaClient(t, `{}`)
	request := &DeleteUserRequest{ClusterID: "cluster-1", Username: "alice"}
	if _, err := client.DeleteUser(request); err != nil {
		t.Fatalf("DeleteUser returned error: %v", err)
	}
}

func TestClient_ResetUserPassword(t *testing.T) {
	client := newMockKafkaClient(t, `{}`)
	request := &ResetUserPasswordRequest{ClusterID: "cluster-1", Username: "alice", Password: "Password9!"}
	if _, err := client.ResetUserPassword(request); err != nil {
		t.Fatalf("ResetUserPassword returned error: %v", err)
	}
}

func TestClient_ListUsers(t *testing.T) {
	client := newMockKafkaClient(t, `{"users":[{"username":"alice"}]}`)
	result, err := client.ListUsers(ListUsersRequest{ClusterID: "cluster-1"})
	if err != nil {
		t.Fatalf("ListUsers returned error: %v", err)
	}
	if len(result.Users) != 1 || result.Users[0].Username != "alice" {
		t.Fatalf("Users = %#v, want alice", result.Users)
	}
	if _, err := client.ListUsers(1); err == nil {
		t.Fatal("ListUsers expected error for unsupported request type")
	}
}

func TestClient_CreateAcl(t *testing.T) {
	client := newMockKafkaClient(t, `{}`)
	request := &CreateAclRequest{
		ClusterID: "cluster-1", Username: "alice", PatternType: "LITERAL",
		ResourceType: "TOPIC", ResourceName: "topic-1", Operations: []string{"READ"},
	}
	if _, err := client.CreateAcl(request); err != nil {
		t.Fatalf("CreateAcl returned error: %v", err)
	}
	request.Operations = nil
	if _, err := client.CreateAcl(request); err == nil {
		t.Fatal("CreateAcl expected error for empty operations")
	}
}

func TestClient_DeleteAcl(t *testing.T) {
	client := newMockKafkaClient(t, `{}`)
	request := &DeleteAclRequest{
		ClusterID: "cluster-1", Username: "alice", PatternType: "LITERAL",
		ResourceType: "TOPIC", ResourceName: "topic-1", Operation: "READ",
	}
	if _, err := client.DeleteAcl(request); err != nil {
		t.Fatalf("DeleteAcl returned error: %v", err)
	}
}

func TestClient_ListAcls(t *testing.T) {
	client := newMockKafkaClient(t, `{"acls":[{"username":"alice","resourceType":"TOPIC"}]}`)
	request := ListAclRequest{
		ClusterID: "cluster-1", Username: "alice", PatternType: "LITERAL",
		ResourceType: "TOPIC", ResourceName: "topic-1",
	}
	result, err := client.ListAcls(request)
	if err != nil {
		t.Fatalf("ListAcls returned error: %v", err)
	}
	if len(result.Acls) != 1 || result.Acls[0].Username != "alice" {
		t.Fatalf("Acls = %#v, want alice", result.Acls)
	}
	if _, err := client.ListAcls(1); err == nil {
		t.Fatal("ListAcls expected error for unsupported request type")
	}
}

func TestClient_ListJobs(t *testing.T) {
	client := newMockKafkaClient(t, `{"jobs":[{"actionId":"action-1"}]}`)
	request := &ListJobsRequest{
		ListRequest: ListRequest{Marker: "next", MaxKeys: 20}, ClusterID: "cluster-1", Name: "restart",
	}
	result, err := client.ListJobs(request)
	if err != nil {
		t.Fatalf("ListJobs returned error: %v", err)
	}
	if len(result.Jobs) != 1 || result.Jobs[0].ActionID != "action-1" {
		t.Fatalf("Jobs = %#v, want action-1", result.Jobs)
	}
}

func TestClient_GetJob(t *testing.T) {
	client := newMockKafkaClient(t, `{"job":{"actionId":"action-1"}}`)
	request := &GetJobDetailRequest{ClusterID: "cluster-1", ActionID: "action-1"}
	result, err := client.GetJob(request)
	if err != nil {
		t.Fatalf("GetJob returned error: %v", err)
	}
	if result.Job == nil || result.Job.ActionID != "action-1" {
		t.Fatalf("Job = %#v, want action-1", result.Job)
	}
}

func TestClient_GetOperation(t *testing.T) {
	client := newMockKafkaClient(t, `{}`)
	request := &GetOperationDetailRequest{ClusterID: "cluster-1", ActionID: "action-1", OperationID: "operation-1"}
	if _, err := client.GetOperation(request); err != nil {
		t.Fatalf("GetOperation returned error: %v", err)
	}
}

func TestClient_StartJob(t *testing.T) {
	client := newMockKafkaClient(t, `{}`)
	request := &StartJobRequest{ClusterID: "cluster-1", ActionID: "action-1"}
	if _, err := client.StartJob(request); err != nil {
		t.Fatalf("StartJob returned error: %v", err)
	}
}

func TestClient_CancelJob(t *testing.T) {
	client := newMockKafkaClient(t, `{}`)
	request := &CancelJobRequest{ClusterID: "cluster-1", ActionID: "action-1"}
	if _, err := client.CancelJob(request); err != nil {
		t.Fatalf("CancelJob returned error: %v", err)
	}
}

func TestClient_SuspendJob(t *testing.T) {
	client := newMockKafkaClient(t, `{}`)
	request := &SuspendJobRequest{ClusterID: "cluster-1", ActionID: "action-1"}
	if _, err := client.SuspendJob(request); err != nil {
		t.Fatalf("SuspendJob returned error: %v", err)
	}
}

func TestClient_ResumeJob(t *testing.T) {
	client := newMockKafkaClient(t, `{}`)
	request := &ResumeJobRequest{ClusterID: "cluster-1", ActionID: "action-1"}
	if _, err := client.ResumeJob(request); err != nil {
		t.Fatalf("ResumeJob returned error: %v", err)
	}
}

func TestClient_ListQuotas(t *testing.T) {
	client := newMockKafkaClient(t, `{"quotas":[{"username":"alice"}]}`)
	request := &ListQuotasRequest{ClusterID: "cluster-1", EntityType: "USER"}
	result, err := client.ListQuotas(request)
	if err != nil {
		t.Fatalf("ListQuotas returned error: %v", err)
	}
	if len(result.Quotas) != 1 || result.Quotas[0].Username != "alice" {
		t.Fatalf("Quotas = %#v, want alice", result.Quotas)
	}
}

func TestClient_CreateQuota(t *testing.T) {
	client := newMockKafkaClient(t, `{}`)
	request := &CreateQuotaRequest{ClusterID: "cluster-1", Username: "alice", ProducerByteRate: kafkaInt64(1024)}
	if _, err := client.CreateQuota(request); err != nil {
		t.Fatalf("CreateQuota returned error: %v", err)
	}
}

func TestClient_UpdateQuota(t *testing.T) {
	client := newMockKafkaClient(t, `{}`)
	request := &UpdateQuotaRequest{ClusterID: "cluster-1", Username: "alice", ProducerByteRate: kafkaInt64(2048)}
	if _, err := client.UpdateQuota(request); err != nil {
		t.Fatalf("UpdateQuota returned error: %v", err)
	}
}

func TestClient_DeleteQuota(t *testing.T) {
	client := newMockKafkaClient(t, `{}`)
	request := &DeleteQuotaRequest{
		ClusterID: "cluster-1", Username: "alice", UserDefault: kafkaBool(false),
		ClientID: "client-1", ClientDefault: kafkaBool(true),
	}
	if _, err := client.DeleteQuota(request); err != nil {
		t.Fatalf("DeleteQuota returned error: %v", err)
	}
}

func TestClientAPIsRejectNilRequests(t *testing.T) {
	client := &Client{}
	tests := []struct {
		name string
		call func() error
	}{
		{"CreateCluster", func() error { _, err := client.CreateCluster(nil); return err }},
		{"DeleteCluster", func() error { _, err := client.DeleteCluster(nil); return err }},
		{"ListClusters", func() error { _, err := client.ListClusters(nil); return err }},
		{"GetClusterDetail", func() error { _, err := client.GetClusterDetail(nil); return err }},
		{"GetClusterDeletion", func() error { _, err := client.GetClusterDeletion(nil); return err }},
		{"GetClusterAccessEndpoints", func() error { _, err := client.GetClusterAccessEndpoints(nil); return err }},
		{"GetClusterNodes", func() error { _, err := client.GetClusterNodes(nil); return err }},
		{"GetClusterConfigurations", func() error { _, err := client.GetClusterConfigurations(nil); return err }},
		{"IncreaseBrokerCount", func() error { _, err := client.IncreaseBrokerCount(nil); return err }},
		{"DecreaseBrokerCount", func() error { _, err := client.DecreaseBrokerCount(nil); return err }},
		{"MigrateClusterAz", func() error { _, err := client.MigrateClusterAz(nil); return err }},
		{"UnifyClusterEndpoint", func() error { _, err := client.UnifyClusterEndpoint(nil); return err }},
		{"UpdateBrokerNodeType", func() error { _, err := client.UpdateBrokerNodeType(nil); return err }},
		{"ExpandBrokerDiskCapacity", func() error { _, err := client.ExpandBrokerDiskCapacity(nil); return err }},
		{"UpdateAccessConfig", func() error { _, err := client.UpdateAccessConfig(nil); return err }},
		{"StartCluster", func() error { _, err := client.StartCluster(nil); return err }},
		{"StopCluster", func() error { _, err := client.StopCluster(nil); return err }},
		{"ResizeClusterEipBandwidth", func() error { _, err := client.ResizeClusterEipBandwidth(nil); return err }},
		{"SwitchClusterEip", func() error { _, err := client.SwitchClusterEip(nil); return err }},
		{"UpdateStoragePolicy", func() error { _, err := client.UpdateStoragePolicy(nil); return err }},
		{"UpdateKafkaConfig", func() error { _, err := client.UpdateKafkaConfig(nil); return err }},
		{"UpdateSecurityGroup", func() error { _, err := client.UpdateSecurityGroup(nil); return err }},
		{"UpdateMaintenanceDuration", func() error { _, err := client.UpdateMaintenanceDuration(nil); return err }},
		{"SwitchClusterIntranetIp", func() error { _, err := client.SwitchClusterIntranetIp(nil); return err }},
		{"GetClusterCurrentController", func() error { _, err := client.GetClusterCurrentController(nil); return err }},
		{"GetClusterHistoryController", func() error { _, err := client.GetClusterHistoryController(nil); return err }},
		{"RestartCluster", func() error { _, err := client.RestartCluster(nil); return err }},
		{"RestartBroker", func() error { _, err := client.RestartBroker(nil); return err }},
		{"SwitchClusterAdvertisedIp", func() error { _, err := client.SwitchClusterAdvertisedIp(nil); return err }},
		{"SwitchClusterDomain", func() error { _, err := client.SwitchClusterDomain(nil); return err }},
		{"GetZkPassword", func() error { _, err := client.GetZkPassword(nil); return err }},
		{"CreateClusterConfig", func() error { _, err := client.CreateClusterConfig(nil); return err }},
		{"ListClusterConfigs", func() error { _, err := client.ListClusterConfigs(nil); return err }},
		{"DeleteClusterConfig", func() error { _, err := client.DeleteClusterConfig(nil); return err }},
		{"GetClusterConfig", func() error { _, err := client.GetClusterConfig(nil); return err }},
		{"CreateClusterConfigRevision", func() error { _, err := client.CreateClusterConfigRevision(nil); return err }},
		{"ListClusterConfigRevisions", func() error { _, err := client.ListClusterConfigRevisions(nil); return err }},
		{"GetClusterConfigRevision", func() error { _, err := client.GetClusterConfigRevision(nil); return err }},
		{"UpdateTopic", func() error { _, err := client.UpdateTopic(nil); return err }},
		{"GetSubscribedGroupDetail", func() error { _, err := client.GetSubscribedGroupDetail(nil); return err }},
		{"ListTopicPartitions", func() error { _, err := client.ListTopicPartitions(nil); return err }},
		{"GetTopicPartitionDetail", func() error { _, err := client.GetTopicPartitionDetail(nil); return err }},
		{"ListTopic", func() error { _, err := client.ListTopic(nil); return err }},
		{"GetTopicDetail", func() error { _, err := client.GetTopicDetail(nil); return err }},
		{"CreateTopic", func() error { _, err := client.CreateTopic(nil); return err }},
		{"ListSubscribedGroups", func() error { _, err := client.ListSubscribedGroups(nil); return err }},
		{"DeleteTopic", func() error { _, err := client.DeleteTopic(nil); return err }},
		{"GetTopicPartitionOverview", func() error { _, err := client.GetTopicPartitionOverview(nil); return err }},
		{"GetSubscribedGroupOverview", func() error { _, err := client.GetSubscribedGroupOverview(nil); return err }},
		{"SendTopicMessage", func() error { _, err := client.SendTopicMessage(nil); return err }},
		{"QueryTopicMessagesByStartTime", func() error { _, err := client.QueryTopicMessagesByStartTime(nil); return err }},
		{"QueryTopicMessagesByStartOffset", func() error { _, err := client.QueryTopicMessagesByStartOffset(nil); return err }},
		{"ListSubscribedTopics", func() error { _, err := client.ListSubscribedTopics(nil); return err }},
		{"ListConsumerGroup", func() error { _, err := client.ListConsumerGroup(nil); return err }},
		{"DeleteConsumerGroup", func() error { _, err := client.DeleteConsumerGroup(nil); return err }},
		{"ResetConsumerGroup", func() error { _, err := client.ResetConsumerGroup(nil); return err }},
		{"GetSubscribedTopicOverview", func() error { _, err := client.GetSubscribedTopicOverview(nil); return err }},
		{"CreateUser", func() error { _, err := client.CreateUser(nil); return err }},
		{"DeleteUser", func() error { _, err := client.DeleteUser(nil); return err }},
		{"ResetUserPassword", func() error { _, err := client.ResetUserPassword(nil); return err }},
		{"ListUsers", func() error { _, err := client.ListUsers((*ListUsersRequest)(nil)); return err }},
		{"CreateAcl", func() error { _, err := client.CreateAcl(nil); return err }},
		{"DeleteAcl", func() error { _, err := client.DeleteAcl(nil); return err }},
		{"ListAcls", func() error { _, err := client.ListAcls((*ListAclRequest)(nil)); return err }},
		{"ListJobs", func() error { _, err := client.ListJobs(nil); return err }},
		{"GetJob", func() error { _, err := client.GetJob(nil); return err }},
		{"GetOperation", func() error { _, err := client.GetOperation(nil); return err }},
		{"StartJob", func() error { _, err := client.StartJob(nil); return err }},
		{"CancelJob", func() error { _, err := client.CancelJob(nil); return err }},
		{"SuspendJob", func() error { _, err := client.SuspendJob(nil); return err }},
		{"ResumeJob", func() error { _, err := client.ResumeJob(nil); return err }},
		{"ListQuotas", func() error { _, err := client.ListQuotas(nil); return err }},
		{"CreateQuota", func() error { _, err := client.CreateQuota(nil); return err }},
		{"UpdateQuota", func() error { _, err := client.UpdateQuota(nil); return err }},
		{"DeleteQuota", func() error { _, err := client.DeleteQuota(nil); return err }},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			if err := test.call(); err != errRequestNil {
				t.Fatalf("error = %v, want %v", err, errRequestNil)
			}
		})
	}
}

func TestClientAPIsValidateRequiredFieldsAndBounds(t *testing.T) {
	client := &Client{}
	empty := ""
	tests := []struct {
		name    string
		wantErr string
		call    func() error
	}{
		{"cluster name", "name", func() error { _, err := client.CreateCluster(&CreateClusterRequest{}); return err }},
		{"broker node", "nodeId", func() error {
			_, err := client.RestartBroker(&RestartBrokerRequest{ClusterID: "cluster-1"})
			return err
		}},
		{"topic name", "topicName", func() error { _, err := client.UpdateTopic(&UpdateTopicRequest{ClusterID: "cluster-1"}); return err }},
		{"group name", "groupName", func() error {
			_, err := client.GetSubscribedGroupDetail(&GetSubscribedGroupDetailRequest{ClusterID: "cluster-1", TopicName: "topic-1"})
			return err
		}},
		{"partition id", "partitionId", func() error {
			_, err := client.GetTopicPartitionDetail(&GetTopicPartitionDetailRequest{ClusterID: "cluster-1", TopicName: "topic-1"})
			return err
		}},
		{"empty message", "value", func() error {
			_, err := client.SendTopicMessage(&SendTopicMessageRequest{ClusterID: "cluster-1", TopicName: "topic-1", Value: &empty})
			return err
		}},
		{"start time", "startTime", func() error {
			_, err := client.QueryTopicMessagesByStartTime(&QueryTopicMessagesByStartTimeRequest{ClusterID: "cluster-1", TopicName: "topic-1"})
			return err
		}},
		{"negative partition", "partitionId", func() error {
			_, err := client.QueryTopicMessagesByStartOffset(&QueryTopicMessagesByStartOffsetRequest{ClusterID: "cluster-1", TopicName: "topic-1", PartitionID: -1})
			return err
		}},
		{"negative offset", "startOffset", func() error {
			_, err := client.QueryTopicMessagesByStartOffset(&QueryTopicMessagesByStartOffsetRequest{ClusterID: "cluster-1", TopicName: "topic-1", StartOffset: -1})
			return err
		}},
		{"reset strategy", "resetStrategy", func() error {
			_, err := client.ResetConsumerGroup(&ResetConsumerGroupRequest{
				ClusterID: "cluster-1", GroupName: "group-1", TopicName: "topic-1", Partitions: []int{0},
			})
			return err
		}},
		{"user password", "password", func() error {
			_, err := client.CreateUser(&CreateUserRequest{ClusterID: "cluster-1", Username: "alice"})
			return err
		}},
		{"reset password", "password", func() error {
			_, err := client.ResetUserPassword(&ResetUserPasswordRequest{ClusterID: "cluster-1", Username: "alice"})
			return err
		}},
		{"ACL resource", "resourceName", func() error {
			_, err := client.CreateAcl(&CreateAclRequest{
				ClusterID: "cluster-1", Username: "alice", PatternType: "LITERAL", ResourceType: "TOPIC",
			})
			return err
		}},
		{"ACL operation", "operation", func() error {
			_, err := client.DeleteAcl(&DeleteAclRequest{
				ClusterID: "cluster-1", Username: "alice", PatternType: "LITERAL", ResourceType: "TOPIC", ResourceName: "topic-1",
			})
			return err
		}},
		{"job action", "actionId", func() error { _, err := client.GetJob(&GetJobDetailRequest{ClusterID: "cluster-1"}); return err }},
		{"job operation", "operationId", func() error {
			_, err := client.GetOperation(&GetOperationDetailRequest{ClusterID: "cluster-1", ActionID: "action-1"})
			return err
		}},
		{"list users cluster", "clusterId", func() error { _, err := client.ListUsers(""); return err }},
		{"list ACLs cluster", "clusterId", func() error { _, err := client.ListAcls(""); return err }},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			err := test.call()
			if err == nil || !strings.Contains(err.Error(), test.wantErr) {
				t.Fatalf("error = %v, want error containing %q", err, test.wantErr)
			}
		})
	}
}

func TestClientConstructorsAndPasswordEncryptionRejectInvalidCredentials(t *testing.T) {
	if _, err := NewClient("", "sk", ""); err == nil {
		t.Fatal("NewClient expected error for empty access key")
	}
	if _, err := NewClientWithSTS("ak", "sk", "", ""); err == nil {
		t.Fatal("NewClientWithSTS expected error for empty session token")
	}

	var nilClient *Client
	if _, err := nilClient.encryptPassword("password"); err == nil {
		t.Fatal("encryptPassword expected error for nil client")
	}
	client, err := NewClientWithConfig(&KafkaClientConfiguration{})
	if err != nil {
		t.Fatalf("NewClientWithConfig returned error: %v", err)
	}
	if _, err := client.encryptPassword("password"); err == nil {
		t.Fatal("encryptPassword expected error for missing credentials")
	}
}
