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

// probe.go - the probe APIs definition supported by the VPC service

package vpc

import (
	"fmt"
	"strconv"

	"github.com/baidubce/bce-sdk-go/bce"
	"github.com/baidubce/bce-sdk-go/http"
)

// CreateProbe - create a new probe with the specified parameters
//
// PARAMS:
//     - args: the arguments to create a probe
// RETURNS:
//     - *CreateProbeResult: the ID of the probe newly created
//     - error: nil if success otherwise the specific error
func (c *Client) CreateProbe(args *CreateProbeArgs) (*CreateProbeResult, error) {
	if args == nil {
		return nil, fmt.Errorf("CreateProbeResult cannot be nil.")
	}
	result := &CreateProbeResult{}
	err := bce.NewRequestBuilder(c).
		WithURL(getURLForProbe()).
		WithMethod(http.POST).
		WithBody(args).
		WithQueryParamFilter("clientToken", args.ClientToken).
		WithResult(result).
		Do()

	return result, err
}

// UpdateProbe - update a probe with the specified parameters
//
// PARAMS:
//     - probeId: the ID of the probe to be updated
//     - args: the arguments to update the probe
// RETURNS:
//     - error: nil if success otherwise the specific error
func (c *Client) UpdateProbe(probeId string, args *UpdateProbeArgs) error {
	if probeId == "" {
		return fmt.Errorf("The probeId cannot be blank.")
	}
	if args == nil {
		return fmt.Errorf("UpdateProbeArgs cannot be nil.")
	}
	return bce.NewRequestBuilder(c).
		WithURL(getURLForProbeId(probeId)).
		WithMethod(http.PUT).
		WithBody(args).
		WithQueryParamFilter("clientToken", args.ClientToken).
		WithQueryParam("modifyAttribute", "").
		Do()
}

// DeleteProbe - delete a probe
// PARAMS:
//     - probeId: the ID of the probe to be deleted
//     - clientToken: the client token of the request
// RETURNS:
//     - error: nil if success otherwise the specific error
func (c *Client) DeleteProbe(probeId string, clientToken string) error {
	if probeId == "" {
		return fmt.Errorf("The probeId cannot be blank.")
	}

	return bce.NewRequestBuilder(c).
		WithURL(getURLForProbeId(probeId)).
		WithMethod(http.DELETE).
		WithQueryParamFilter("clientToken", clientToken).
		Do()
}

// ListProbes - list all probes of the user
//
// PARAMS:
//     - args: the arguments to list probes
// RETURNS:
//     - *ListProbesResult: the infromation of all probes
//     - error: nil if success otherwise the specific error
func (c *Client) ListProbes(args *ListProbesArgs) (*ListProbesResult, error) {
	if args == nil {
		args = &ListProbesArgs{}
	}
	if args.MaxKeys < 0 || args.MaxKeys > 1000 {
		return nil, fmt.Errorf("The field maxKeys is out of range [0, 1000]")
	} else if args.MaxKeys == 0 {
		args.MaxKeys = 1000
	}

	result := &ListProbesResult{}
	err := bce.NewRequestBuilder(c).
		WithURL(getURLForProbe()).
		WithMethod(http.GET).
		WithQueryParamFilter("marker", args.Marker).
		WithQueryParamFilter("maxKeys", strconv.Itoa(args.MaxKeys)).
		WithResult(result).
		Do()

	return result, err
}

// GetProbeDetail - get details of a probe
//
// PARAMS:
//     - probeId: the ID of the probe to get
// RETURNS:
//     - *Probe: the infromation of the probe
//     - error: nil if success otherwise the specific error
func (c *Client) GetProbeDetail(probeId string) (*Probe, error) {
	if probeId == "" {
		return nil, fmt.Errorf("The probeId cannot be blank.")
	}

	result := &Probe{}
	err := bce.NewRequestBuilder(c).
		WithURL(getURLForProbeId(probeId)).
		WithMethod(http.GET).
		WithResult(result).
		Do()

	return result, err
}
