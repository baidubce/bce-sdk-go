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
	"encoding/json"
	"reflect"
	"testing"
)

func marshalObject(t *testing.T, value interface{}) map[string]interface{} {
	t.Helper()
	data, err := json.Marshal(value)
	if err != nil {
		t.Fatalf("json.Marshal returned error: %v", err)
	}
	result := map[string]interface{}{}
	if err := json.Unmarshal(data, &result); err != nil {
		t.Fatalf("json.Unmarshal returned error: %v", err)
	}
	return result
}

func TestNullableNestedRequestObjectsAreOmitted(t *testing.T) {
	cluster := marshalObject(t, CreateClusterRequest{Name: "serverless", Type: TypeServerless})
	if _, ok := cluster["provisioned"]; ok {
		t.Fatalf("serverless request unexpectedly contains provisioned: %v", cluster)
	}

	enabled := false
	storage := marshalObject(t, UpdateStoragePolicyRequest{StoragePolicyEnabled: &enabled})
	if _, ok := storage["storagePolicy"]; ok {
		t.Fatalf("storage update unexpectedly contains storagePolicy: %v", storage)
	}
}

func TestMessageFieldsDistinguishAbsentAndEmpty(t *testing.T) {
	absent := marshalObject(t, SendTopicMessageRequest{})
	if _, ok := absent["key"]; ok {
		t.Fatalf("unset key unexpectedly serialized: %v", absent)
	}
	if _, ok := absent["value"]; ok {
		t.Fatalf("unset value unexpectedly serialized: %v", absent)
	}

	empty := ""
	explicit := marshalObject(t, SendTopicMessageRequest{Key: &empty, Value: &empty})
	if explicit["key"] != "" || explicit["value"] != "" {
		t.Fatalf("explicit empty message fields were not preserved: %v", explicit)
	}

	var record QueryTopicRecord
	if err := json.Unmarshal([]byte(`{"key":null,"value":""}`), &record); err != nil {
		t.Fatalf("json.Unmarshal returned error: %v", err)
	}
	if record.Key != nil {
		t.Fatalf("null key = %#v, want nil", record.Key)
	}
	if record.Value == nil || *record.Value != "" {
		t.Fatalf("empty value = %#v, want pointer to empty string", record.Value)
	}
}

func TestExplicitEmptyUpdateCollectionsArePreserved(t *testing.T) {
	unsetAccess := marshalObject(t, UpdateAccessConfigRequest{})
	if _, ok := unsetAccess["authentications"]; ok {
		t.Fatalf("nil authentications unexpectedly serialized: %v", unsetAccess)
	}
	access := marshalObject(t, UpdateAccessConfigRequest{Authentications: []Authentication{}})
	if values, ok := access["authentications"].([]interface{}); !ok || len(values) != 0 {
		t.Fatalf("authentications = %#v, want explicit empty array", access["authentications"])
	}

	unsetUser := marshalObject(t, CreateUserRequest{})
	if _, ok := unsetUser["saslMechanisms"]; ok {
		t.Fatalf("nil saslMechanisms unexpectedly serialized: %v", unsetUser)
	}
	user := marshalObject(t, CreateUserRequest{SASLMechanisms: []string{}})
	if values, ok := user["saslMechanisms"].([]interface{}); !ok || len(values) != 0 {
		t.Fatalf("saslMechanisms = %#v, want explicit empty array", user["saslMechanisms"])
	}
}

func TestAllRequestCollectionsPreserveExplicitEmptyValues(t *testing.T) {
	tests := []struct {
		name  string
		value interface{}
		field string
	}{
		{"acl operations", CreateAclRequest{Operations: []string{}}, "operations"},
		{"cluster tags", CreateClusterRequest{Tags: []Tag{}}, "tags"},
		{"expand coupons", ExpandBrokerDiskCapacityRequest{CouponIds: []string{}}, "couponIds"},
		{"increase coupons", IncreaseBrokerCountRequest{CouponIds: []string{}}, "couponIds"},
		{"migration coupons", MigrateClusterAzRequest{CouponIds: []string{}}, "couponIds"},
		{"migration zones", MigrateClusterAzRequest{LogicalZones: []string{}}, "logicalZones"},
		{"migration subnets", MigrateClusterAzRequest{Subnets: []string{}}, "subnets"},
		{"resize coupons", ResizeClusterEipBandwidthRequest{CouponIds: []string{}}, "couponIds"},
		{"eip authentication", SwitchClusterEipRequest{AuthenticationMode: []AuthMode{}}, "authenticationMode"},
		{"eip coupons", SwitchClusterEipRequest{CouponIds: []string{}}, "couponIds"},
		{"intranet authentication", SwitchClusterIntranetIpRequest{AuthenticationMode: []AuthMode{}}, "authenticationMode"},
		{"access authentication", UpdateAccessConfigRequest{Authentications: []Authentication{}}, "authentications"},
		{"broker coupons", UpdateBrokerNodeTypeRequest{CouponIds: []string{}}, "couponIds"},
		{"maintenance periods", UpdateMaintenanceDurationRequest{MaintenancePeriods: []MaintainPeriod{}}, "maintenancePeriods"},
		{"security groups", UpdateSecurityGroupRequest{SecurityGroupIds: []string{}}, "securityGroupIds"},
		{"config context", CreateClusterConfigRequest{Context: map[string]string{}}, "context"},
		{"revision context", CreateClusterConfigRevisionRequest{Context: map[string]string{}}, "context"},
		{"consumer partitions", ResetConsumerGroupRequest{Partitions: []int{}}, "partitions"},
		{"topic configs", CreateTopicRequest{OtherConfigs: map[string]string{}}, "otherConfigs"},
		{"updated topic configs", UpdateTopicRequest{OtherConfigs: map[string]string{}}, "otherConfigs"},
		{"user mechanisms", CreateUserRequest{SASLMechanisms: []string{}}, "saslMechanisms"},
		{"reset mechanisms", ResetUserPasswordRequest{SASLMechanisms: []string{}}, "saslMechanisms"},
		{"billing coupons", Billing{CouponIds: []string{}}, "couponIds"},
		{"provisioned subnets", Provisioned{Subnets: []Subnet{}}, "subnets"},
		{"provisioned zones", Provisioned{LogicalZones: []string{}}, "logicalZones"},
		{"provisioned authentications", Provisioned{Authentications: []Authentication{}}, "authentications"},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			object := marshalObject(t, test.value)
			if collection, ok := object[test.field]; !ok || collection == nil {
				t.Fatalf("%s = %#v, want explicit empty collection in %v", test.field, collection, object)
			}
		})
	}
}

func TestJavaModelDefaultsAreAppliedDuringSerialization(t *testing.T) {
	billing := marshalObject(t, Billing{})
	if billing["timeUnit"] != "month" || billing["autoRenewTimeUnit"] != "month" || billing["isAutoPay"] != true {
		t.Fatalf("billing defaults = %v", billing)
	}

	storage := marshalObject(t, StorageMeta{})
	if storage["numberOfDisk"] != float64(1) {
		t.Fatalf("numberOfDisk = %v, want 1", storage["numberOfDisk"])
	}

	config := marshalObject(t, ConfigMeta{})
	context, ok := config["context"].(map[string]interface{})
	if !ok || len(context) != 0 {
		t.Fatalf("context = %#v, want initialized empty object", config["context"])
	}
}

func TestJavaCompatibleModelConstructors(t *testing.T) {
	billing := NewBilling()
	if billing.TimeUnit != "month" || billing.AutoRenewTimeUnit != "month" || billing.IsAutoPay == nil || !*billing.IsAutoPay {
		t.Fatalf("NewBilling() = %#v", billing)
	}
	storage := NewStorageMeta()
	if storage.NumberOfDisk != 1 {
		t.Fatalf("NewStorageMeta().NumberOfDisk = %d, want 1", storage.NumberOfDisk)
	}
	config := NewConfigMeta()
	if config.Context == nil || len(config.Context) != 0 {
		t.Fatalf("NewConfigMeta().Context = %#v, want initialized empty map", config.Context)
	}
}

func TestNullableNestedResponsesRemainNilWhenAbsent(t *testing.T) {
	tests := []struct {
		name     string
		response interface{}
		field    string
	}{
		{"controller", &GetClusterCurrentControllerResponse{}, "Controller"},
		{"cluster", &GetClusterDetailResponse{}, "Cluster"},
		{"config", &GetClusterConfigResponse{}, "Config"},
		{"revision", &GetClusterConfigRevisionResponse{}, "Revision"},
		{"job", &GetJobDetailResponse{}, "Job"},
		{"operation", &GetOperationDetailResponse{}, "Operation"},
		{"created quota", &CreateQuotaResponse{}, "Quota"},
		{"updated quota", &UpdateQuotaResponse{}, "Quota"},
		{"topic", &GetTopicDetailResponse{}, "Topic"},
		{"partition", &GetTopicPartitionDetailResponse{}, "Partition"},
		{"message", &SendTopicMessageResponse{}, "Message"},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			if err := json.Unmarshal([]byte(`{}`), test.response); err != nil {
				t.Fatalf("json.Unmarshal returned error: %v", err)
			}
			field := reflect.ValueOf(test.response).Elem().FieldByName(test.field)
			if field.Kind() != reflect.Ptr {
				t.Fatalf("%s kind = %s, want pointer", test.field, field.Kind())
			}
			if !field.IsNil() {
				t.Fatalf("%s = %#v, want nil", test.field, field.Interface())
			}
		})
	}
}

func TestGetClusterNodesResponseUnmarshalsNumericBrokerID(t *testing.T) {
	var response GetClusterNodesResponse
	if err := json.Unmarshal([]byte(`{"nodes":[{"brokerId":1,"nodeId":"node-1","state":"ACTIVE"}]}`), &response); err != nil {
		t.Fatalf("json.Unmarshal returned error: %v", err)
	}
	if len(response.Nodes) != 1 {
		t.Fatalf("len(Nodes) = %d, want 1", len(response.Nodes))
	}
	if response.Nodes[0].BrokerID == nil || *response.Nodes[0].BrokerID != 1 {
		t.Fatalf("BrokerID = %#v, want pointer to 1", response.Nodes[0].BrokerID)
	}

	encoded, err := json.Marshal(response)
	if err != nil {
		t.Fatalf("json.Marshal returned error: %v", err)
	}
	var object map[string]interface{}
	if err := json.Unmarshal(encoded, &object); err != nil {
		t.Fatalf("json.Unmarshal encoded response returned error: %v", err)
	}
	nodes, ok := object["nodes"].([]interface{})
	if !ok || len(nodes) != 1 {
		t.Fatalf("encoded nodes = %#v, want one node", object["nodes"])
	}
	node, ok := nodes[0].(map[string]interface{})
	if !ok || node["brokerId"] != float64(1) {
		t.Fatalf("encoded brokerId = %#v, want JSON number 1", node["brokerId"])
	}
}
