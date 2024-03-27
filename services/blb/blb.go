/*
 * Copyright 2020 Baidu, Inc.
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

// blb.go - the Normal BLB APIs definition supported by the BLB service

package blb

import (
	"fmt"
	"strconv"

	"github.com/baidubce/bce-sdk-go/bce"
	"github.com/baidubce/bce-sdk-go/http"
)

// CreateLoadBalancer - create a LoadBalancer
//
// PARAMS:
//     - args: parameters to create LoadBalancer
// RETURNS:
//     - *CreateLoadBalancerResult: the result of create LoadBalancer, contains new LoadBalancer's ID
//     - error: nil if ok otherwise the specific error
func (c *Client) CreateLoadBalancer(args *CreateLoadBalancerArgs) (*CreateLoadBalancerResult, error) {
	if args == nil || len(args.SubnetId) == 0 {
		return nil, fmt.Errorf("unset subnet id")
	}

	if len(args.VpcId) == 0 {
		return nil, fmt.Errorf("unset vpc id")
	}

	result := &CreateLoadBalancerResult{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.POST).
		WithURL(getBlbUri()).
		WithQueryParamFilter("clientToken", args.ClientToken).
		WithBody(args).
		WithResult(result).
		Do()

	return result, err
}

// UpdateLoadBalancer - update a LoadBalancer
//
// PARAMS:
//     - blbId: LoadBalancer's ID
//     - args: parameters to update LoadBalancer
// RETURNS:
//     - error: nil if ok otherwise the specific error
func (c *Client) UpdateLoadBalancer(blbId string, args *UpdateLoadBalancerArgs) error {
	if args == nil {
		args = &UpdateLoadBalancerArgs{}
	}

	return bce.NewRequestBuilder(c).
		WithMethod(http.PUT).
		WithURL(getBlbUriWithId(blbId)).
		WithQueryParamFilter("clientToken", args.ClientToken).
		WithBody(args).
		Do()
}

// DescribeLoadBalancers - describe all LoadBalancers
//
// PARAMS:
//     - args: parameters to describe all LoadBalancers
// RETURNS:
//     - *DescribeLoadBalancersResult: the result all LoadBalancers's detail
//     - error: nil if ok otherwise the specific error
func (c *Client) DescribeLoadBalancers(args *DescribeLoadBalancersArgs) (*DescribeLoadBalancersResult, error) {
	if args == nil {
		args = &DescribeLoadBalancersArgs{}
	}

	if args.MaxKeys > 1000 || args.MaxKeys <= 0 {
		args.MaxKeys = 1000
	}

	result := &DescribeLoadBalancersResult{}
	request := bce.NewRequestBuilder(c).
		WithMethod(http.GET).
		WithURL(getBlbUri()).
		WithQueryParamFilter("address", args.Address).
		WithQueryParamFilter("name", args.Name).
		WithQueryParamFilter("blbId", args.BlbId).
		WithQueryParamFilter("bccId", args.BccId).
		WithQueryParamFilter("marker", args.Marker).
		WithQueryParamFilter("maxKeys", strconv.Itoa(args.MaxKeys)).
		WithQueryParamFilter("type", args.Type).
		WithResult(result)

	if args.ExactlyMatch {
		request.WithQueryParam("exactlyMatch", "true")
	}

	err := request.Do()
	return result, err
}

// DescribeLoadBalancerDetail - describe a LoadBalancer
//
// PARAMS:
//     - blbId: describe LoadBalancer's ID
// RETURNS:
//     - *DescribeLoadBalancerDetailResult: the result LoadBalancer detail
//     - error: nil if ok otherwise the specific error
func (c *Client) DescribeLoadBalancerDetail(blbId string) (*DescribeLoadBalancerDetailResult, error) {
	result := &DescribeLoadBalancerDetailResult{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.GET).
		WithURL(getBlbUriWithId(blbId)).
		WithResult(result).
		Do()

	return result, err
}

// DeleteLoadBalancer - delete a LoadBalancer
//
// PARAMS:
//     - blbId: parameters to delete LoadBalancer
// RETURNS:
//     - error: nil if ok otherwise the specific error
func (c *Client) DeleteLoadBalancer(blbId string) error {
	return bce.NewRequestBuilder(c).
		WithMethod(http.DELETE).
		WithURL(getBlbUriWithId(blbId)).
		Do()
}

// DescribeLbClusterDetail - describe a LoadBalancer cluster
//
// PARAMS:
//     - clusterId: describe LoadBalancer cluster's ID
// RETURNS:
//     - *DescribeLbClusterDetailResult: the result LoadBalancer cluster detail
//     - error: nil if ok otherwise the specific error
func (c *Client) DescribeLbClusterDetail(clusterId string) (*DescribeLbClusterDetailResult, error) {
	result := &DescribeLbClusterDetailResult{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.GET).
		WithURL(getBlbClusterUriWithId(clusterId)).
		WithResult(result).
		Do()

	return result, err
}

// DescribeLbClusters - describe all LoadBalancerClusters
//
// PARAMS:
//     - args: parameters to describe all LoadBalancerClusters
// RETURNS:
//     - *DescribeLbClustersResult: the result all LoadBalancerClusters's detail
//     - error: nil if ok otherwise the specific error
func (c *Client) DescribeLbClusters(args *DescribeLbClustersArgs) (*DescribeLbClustersResult, error) {
	if args == nil {
		args = &DescribeLbClustersArgs{}
	}

	if args.MaxKeys > 1000 || args.MaxKeys <= 0 {
		args.MaxKeys = 1000
	}

	result := &DescribeLbClustersResult{}
	request := bce.NewRequestBuilder(c).
		WithMethod(http.GET).
		WithURL(getBlbClusterUri()).
		WithQueryParamFilter("clusterName", args.ClusterName).
		WithQueryParamFilter("clusterId", args.ClusterId).
		WithQueryParamFilter("marker", args.Marker).
		WithQueryParamFilter("maxKeys", strconv.Itoa(args.MaxKeys)).
		WithResult(result)

	if args.ExactlyMatch {
		request.WithQueryParam("exactlyMatch", "true")
	}

	err := request.Do()

	return result, err
}

// UpdateLoadBalancerAcl - update the specified LoadBalancer to support the acl feature
//
// PARAMS:
//     - blbId: LoadBalancer's ID
//     - args: parameters to update LoadBalancer acl
// RETURNS:
//     - error: nil if ok otherwise the specific error
func (c *Client) UpdateLoadBalancerAcl(blbId string, args *UpdateLoadBalancerAclArgs) error {
	if args == nil || args.SupportAcl == nil {
		args = &UpdateLoadBalancerAclArgs{}
	}

	return bce.NewRequestBuilder(c).
		WithMethod(http.PUT).
		WithURL(getBlbAclUriWithId(blbId)).
		WithQueryParamFilter("clientToken", args.ClientToken).
		WithBody(args).
		Do()
}

// StartLoadBalancerAutoRenew - start the specified LoadBalancer to support autoRenew
//
// PARAMS:
//     - blbId: LoadBalancer's ID
//     - args: parameters to start LoadBalancer autoRenew
// RETURNS:
//     - error: nil if ok otherwise the specific error
func (c *Client) StartLoadBalancerAutoRenew(blbId string, args *StartLoadBalancerAutoRenewArgs) error {
	if args == nil {
		return fmt.Errorf("please set blb auto renew argments")
	}

	return bce.NewRequestBuilder(c).
		WithMethod(http.PUT).
		WithURL(getBlbAutoRenewUriWithId(blbId)).
		WithQueryParamFilter("clientToken", args.ClientToken).
		WithBody(args).
		Do()
}
