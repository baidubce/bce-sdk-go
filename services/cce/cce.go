/*
 * Copyright 2017 Baidu, Inc.
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

// cce.go - the cce APIs definition supported by the CCE service

package cce

import (
	"fmt"
	"strconv"

	"github.com/baidubce/bce-sdk-go/bce"
	"github.com/baidubce/bce-sdk-go/http"
)

// CreateCluster - create an CCE Cluster with the specific parameters
//
// PARAMS:
//     - args: the arguments to create a cce cluster
// RETURNS:
//     - *CreateClusterResult: the result of create cluster, contains new Cluster's uuid and order id
func (c *Client) CreateCluster(args *CreateClusterArgs) (*CreateClusterResult, error) {
	if args == nil {
		return nil, fmt.Errorf("please set create cluster argments")
	}

	result := &CreateClusterResult{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.POST).
		WithURL(getClusterUri()).
		WithBody(args).
		WithResult(result).
		Do()

	return result, err
}

// ListClusters - list CCE Clusters with the specific parameters
//
// PARAMS:
//     - args: the arguments to list cce cluster
// RETURNS:
//     - *ListClusterResult: the result of list cluster
func (c *Client) ListClusters(args *ListClusterArgs) (*ListClusterResult, error) {
	if args == nil {
		args = &ListClusterArgs{}
	}

	maxKeysStr := ""
	if args.MaxKeys != 0 {
		maxKeysStr = strconv.Itoa(args.MaxKeys)
	}

	result := &ListClusterResult{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.GET).
		WithURL(getClusterUri()).
		WithQueryParamFilter("maker", args.Marker).
		WithQueryParamFilter("maxKeys", maxKeysStr).
		WithQueryParamFilter("status", string(args.Status)).
		WithResult(result).
		Do()

	return result, err
}

// GetCluster - get a CCE Cluster with the specific cluster uuid
//
// PARAMS:
//     - args: the specific cluster uuid
// RETURNS:
//     - *GetClusterResult: the detail information about the CCE Cluster
func (c *Client) GetCluster(clusterUuid string) (*GetClusterResult, error) {
	result := &GetClusterResult{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.GET).
		WithURL(getClusterUriWithId(clusterUuid)).
		WithResult(result).
		Do()

	return result, err
}

// DeleteCluster - delete a CCE Cluster
//
// PARAMS:
//     - args: the arguments to delete a cce cluster
func (c *Client) DeleteCluster(args *DeleteClusterArgs) error {
	if args == nil || args.ClusterUuid == "" {
		return fmt.Errorf("please set delete cluster uuid")
	}

	return bce.NewRequestBuilder(c).
		WithMethod(http.DELETE).
		WithURL(getClusterUriWithId(args.ClusterUuid)).
		WithQueryParamFilter("deleteEipCds", strconv.FormatBool(args.DeleteEipCds)).
		WithQueryParamFilter("deleteSnap", strconv.FormatBool(args.DeleteSnap)).
		Do()
}

// ScalingUp - scaling up a CCE Cluster
//
// PARAMS:
//     - args: the arguments to create a cce cluster
// RETURNS:
//     - *ScalingUpResult: the result of scaling up cluster, contains new Cluster's uuid and order id
func (c *Client) ScalingUp(args *ScalingUpArgs) (*ScalingUpResult, error) {
	if args == nil {
		return nil, fmt.Errorf("please set scaling up cluster argments")
	}

	if args.ClusterUuid == "" {
		return nil, fmt.Errorf("please set scaling up clusterUuid")
	}

	result := &ScalingUpResult{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.POST).
		WithURL(getClusterUri()).
		WithQueryParam("scalingUp", "").
		WithBody(args).
		WithResult(result).
		Do()

	return result, err
}

// ScalingDown - scaling down a CCE Cluster
//
// PARAMS:
//     - args: the arguments to scaling down a cce cluster
func (c *Client) ScalingDown(args *ScalingDownArgs) error {
	if args == nil {
		return fmt.Errorf("please set scaling down cluster argments")
	}

	if args.ClusterUuid == "" {
		return fmt.Errorf("please set scaling down clusterUuid")
	}

	err := bce.NewRequestBuilder(c).
		WithMethod(http.POST).
		WithURL(getClusterUri()).
		WithQueryParam("scalingDown", "").
		WithQueryParamFilter("deleteEipCds", strconv.FormatBool(args.DeleteEipCds)).
		WithQueryParamFilter("deleteSnap", strconv.FormatBool(args.DeleteSnap)).
		WithBody(args).
		Do()

	return err
}

// ListNodes - list all nodes in CCE Cluster
//
// PARAMS:
//     - args: the arguments to list all nodes
// RETURNS:
//     - *ListNodeResult: the result of list nodes, contains a Cluster's nodes
func (c *Client) ListNodes(args *ListNodeArgs) (*ListNodeResult, error) {
	if args == nil {
		return nil, fmt.Errorf("please set list node argments")
	}

	if args.ClusterUuid == "" {
		return nil, fmt.Errorf("please set cluster uuid for list nodes")
	}

	maxKeysStr := ""
	if args.MaxKeys != 0 {
		maxKeysStr = strconv.Itoa(args.MaxKeys)
	}

	result := &ListNodeResult{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.GET).
		WithURL(getNodeUri()).
		WithQueryParamFilter("maker", args.Marker).
		WithQueryParamFilter("maxKeys", maxKeysStr).
		WithQueryParamFilter("clusterUuid", args.ClusterUuid).
		WithResult(result).
		Do()

	return result, err
}

// ShiftInNode - shift nodes into cluster
//
// PARAMS:
//     - args: the arguments about shift nodes into cce cluster
func (c *Client) ShiftInNode(args *ShiftInNodeArgs) error {
	if args == nil {
		return fmt.Errorf("please set shift in argments")
	}

	if args.ClusterUuid == "" {
		return fmt.Errorf("please set shift in cluster uuid")
	}

	return bce.NewRequestBuilder(c).
		WithMethod(http.POST).
		WithURL(getClusterExisteNodeUri()).
		WithQueryParamFilter("action", "shift_in").
		WithBody(args).
		Do()
}

// ShiftOutNode - shift nodes out from CCE Cluster
//
// PARAMS:
//     - args: the arguments about shift nodes out from cce cluster
func (c *Client) ShiftOutNode(args *ShiftOutNodeArgs) error {
	if args == nil {
		return fmt.Errorf("please set shift out argments")
	}

	if args.ClusterUuid == "" {
		return fmt.Errorf("please set shift out cluster uuid")
	}

	return bce.NewRequestBuilder(c).
		WithMethod(http.POST).
		WithURL(getClusterExisteNodeUri()).
		WithQueryParamFilter("action", "shift_out").
		WithBody(args).
		Do()
}

// ListExistedBccNode - list all bcc nodes which can shifted into CCE cluster
//
// PARAMS:
//     - args: the arguments to list bcc nodes
// RETURNS:
//     - *ListExistedNodeResult: the result of list nodes
func (c *Client) ListExistedBccNode(args *ListExistedNodeArgs) (*ListExistedNodeResult, error) {
	if args == nil {
		return nil, fmt.Errorf("please set list existed bcc node argments")
	}

	if args.ClusterUuid == "" {
		return nil, fmt.Errorf("please set list existed bcc node cluster uuid")
	}

	result := &ListExistedNodeResult{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.POST).
		WithURL(getClusterExisteNodeListUri()).
		WithBody(args).
		WithResult(result).
		Do()

	return result, err
}

// GetContainerNet - get container net in vpc
//
// PARAMS:
//     - args: the arguments to get args
// RETURNS:
//     - *GetContainerNetResult: the result of container net
func (c *Client) GetContainerNet(args *GetContainerNetArgs) (*GetContainerNetResult, error) {
	if args == nil {
		return nil, fmt.Errorf("please set container net argments")
	}

	result := &GetContainerNetResult{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.POST).
		WithURL(getClusterContainerNetUri()).
		WithBody(args).
		WithResult(result).
		Do()

	return result, err
}

// GetKubeConfig - get config file of CCE Cluster
//
// PARAMS:
//     - args: the arguments to get config file of a cce cluster
// RETURNS:
//     - *GetKubeConfigResult: the kubeconfig file data
func (c *Client) GetKubeConfig(args *GetKubeConfigArgs) (*GetKubeConfigResult, error) {
	if args == nil {
		return nil, fmt.Errorf("please set kube config argments")
	}

	if args.ClusterUuid == "" {
		return nil, fmt.Errorf("please set cluster uuid")
	}

	result := &GetKubeConfigResult{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.GET).
		WithURL(getClusterKubeConfigUri()).
		WithQueryParam("clusterUuid", args.ClusterUuid).
		WithQueryParamFilter("type", string(args.Type)).
		WithResult(result).
		Do()

	return result, err
}

// ListVersions - list all support kubernetes version
//
// RETURNS:
//     - *ListVersionsResult: all support kubernetes version list
func (c *Client) ListVersions() (*ListVersionsResult, error) {
	result := &ListVersionsResult{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.GET).
		WithURL(getClusterVersionsUri()).
		WithResult(result).
		Do()

	return result, err
}
