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

// listener.go - the Application BLB Listener APIs definition supported by the APPBLB service

package appblb

import (
	"fmt"
	"strconv"

	"github.com/baidubce/bce-sdk-go/bce"
	"github.com/baidubce/bce-sdk-go/http"
)

// CreateAppTCPListener - create a TCP Listener
//
// PARAMS:
//     - blbId: LoadBalancer's ID
//     - args: parameters to create TCP Listener
// RETURNS:
//     - error: nil if ok otherwise the specific error
func (c *Client) CreateAppTCPListener(blbId string, args *CreateAppTCPListenerArgs) error {
	if args == nil {
		return fmt.Errorf("unset args")
	}

	if args.ListenerPort == 0 {
		return fmt.Errorf("unsupport listener port")
	}

	if len(args.Scheduler) == 0 {
		return fmt.Errorf("unset scheduler")
	}

	return bce.NewRequestBuilder(c).
		WithMethod(http.POST).
		WithURL(getAppTCPListenerUri(blbId)).
		WithQueryParamFilter("clientToken", args.ClientToken).
		WithBody(args).
		Do()
}

// CreateAppUDPListener - create a UDP Listener
//
// PARAMS:
//     - blbId: LoadBalancer's ID
//     - args: parameters to create UDP Listener
// RETURNS:
//     - error: nil if ok otherwise the specific error
func (c *Client) CreateAppUDPListener(blbId string, args *CreateAppUDPListenerArgs) error {
	if args == nil {
		return fmt.Errorf("unset args")
	}

	if args.ListenerPort == 0 {
		return fmt.Errorf("unsupport listener port")
	}

	if len(args.Scheduler) == 0 {
		return fmt.Errorf("unset scheduler")
	}

	return bce.NewRequestBuilder(c).
		WithMethod(http.POST).
		WithURL(getAppUDPListenerUri(blbId)).
		WithQueryParamFilter("clientToken", args.ClientToken).
		WithBody(args).
		Do()
}

// CreateAppHTTPListener - create a HTTP Listener
//
// PARAMS:
//     - blbId: LoadBalancer's ID
//     - args: parameters to create HTTP Listener
// RETURNS:
//     - error: nil if ok otherwise the specific error
func (c *Client) CreateAppHTTPListener(blbId string, args *CreateAppHTTPListenerArgs) error {
	if args == nil {
		return fmt.Errorf("unset args")
	}

	if args.ListenerPort == 0 {
		return fmt.Errorf("unsupport listener port")
	}

	if len(args.Scheduler) == 0 {
		return fmt.Errorf("unset scheduler")
	}

	return bce.NewRequestBuilder(c).
		WithMethod(http.POST).
		WithURL(getAppHTTPListenerUri(blbId)).
		WithQueryParamFilter("clientToken", args.ClientToken).
		WithBody(args).
		Do()
}

// CreateAppHTTPSListener - create a HTTPS Listener
//
// PARAMS:
//     - blbId: LoadBalancer's ID
//     - args: parameters to create HTTPS Listener
// RETURNS:
//     - error: nil if ok otherwise the specific error
func (c *Client) CreateAppHTTPSListener(blbId string, args *CreateAppHTTPSListenerArgs) error {
	if args == nil {
		return fmt.Errorf("unset args")
	}

	if args.ListenerPort == 0 {
		return fmt.Errorf("unsupport listener port")
	}

	if len(args.Scheduler) == 0 {
		return fmt.Errorf("unset scheduler")
	}

	if len(args.CertIds) == 0 {
		return fmt.Errorf("unset certIds")
	}

	return bce.NewRequestBuilder(c).
		WithMethod(http.POST).
		WithURL(getAppHTTPSListenerUri(blbId)).
		WithQueryParamFilter("clientToken", args.ClientToken).
		WithBody(args).
		Do()
}

// CreateAppSSLListener - create a SSL Listener
//
// PARAMS:
//     - blbId: LoadBalancer's ID
//     - args: parameters to create SSL Listener
// RETURNS:
//     - error: nil if ok otherwise the specific error
func (c *Client) CreateAppSSLListener(blbId string, args *CreateAppSSLListenerArgs) error {
	if args == nil {
		return fmt.Errorf("unset args")
	}

	if args.ListenerPort == 0 {
		return fmt.Errorf("unsupport listener port")
	}

	if len(args.Scheduler) == 0 {
		return fmt.Errorf("unset scheduler")
	}

	if len(args.CertIds) == 0 {
		return fmt.Errorf("unset certIds")
	}

	return bce.NewRequestBuilder(c).
		WithMethod(http.POST).
		WithURL(getAppSSLListenerUri(blbId)).
		WithQueryParamFilter("clientToken", args.ClientToken).
		WithBody(args).
		Do()
}

// UpdateAppTCPListener - update a TCP Listener
//
// PARAMS:
//     - blbId: LoadBalancer's ID
//     - args: parameters to update TCP Listener
// RETURNS:
//     - error: nil if ok otherwise the specific error
func (c *Client) UpdateAppTCPListener(blbId string, args *UpdateAppTCPListenerArgs) error {
	if args == nil || args.ListenerPort == 0 {
		return fmt.Errorf("unset listener port")
	}

	return bce.NewRequestBuilder(c).
		WithMethod(http.PUT).
		WithURL(getAppTCPListenerUri(blbId)).
		WithQueryParam("listenerPort", strconv.Itoa(int(args.ListenerPort))).
		WithQueryParamFilter("clientToken", args.ClientToken).
		WithBody(args).
		Do()
}

// UpdateAppUDPListener - update a UDP Listener
//
// PARAMS:
//     - blbId: LoadBalancer's ID
//     - args: parameters to update UDP Listener
// RETURNS:
//     - error: nil if ok otherwise the specific error
func (c *Client) UpdateAppUDPListener(blbId string, args *UpdateAppUDPListenerArgs) error {
	if args == nil || args.ListenerPort == 0 {
		return fmt.Errorf("unset listener port")
	}

	if len(args.Scheduler) == 0 {
		return fmt.Errorf("unset scheduler")
	}

	return bce.NewRequestBuilder(c).
		WithMethod(http.PUT).
		WithURL(getAppUDPListenerUri(blbId)).
		WithQueryParam("listenerPort", strconv.Itoa(int(args.ListenerPort))).
		WithQueryParamFilter("clientToken", args.ClientToken).
		WithBody(args).
		Do()
}

// UpdateAppHTTPListener - update a HTTP Listener
//
// PARAMS:
//     - blbId: LoadBalancer's ID
//     - args: parameters to update HTTP Listener
// RETURNS:
//     - error: nil if ok otherwise the specific error
func (c *Client) UpdateAppHTTPListener(blbId string, args *UpdateAppHTTPListenerArgs) error {
	if args == nil || args.ListenerPort == 0 {
		return fmt.Errorf("unset listener port")
	}

	if len(args.Scheduler) == 0 {
		return fmt.Errorf("unset scheduler")
	}

	return bce.NewRequestBuilder(c).
		WithMethod(http.PUT).
		WithURL(getAppHTTPListenerUri(blbId)).
		WithQueryParam("listenerPort", strconv.Itoa(int(args.ListenerPort))).
		WithQueryParamFilter("clientToken", args.ClientToken).
		WithBody(args).
		Do()
}

// UpdateAppHTTPSListener - update a HTTPS Listener
//
// PARAMS:
//     - blbId: LoadBalancer's ID
//     - args: parameters to update HTTPS Listener
// RETURNS:
//     - error: nil if ok otherwise the specific error
func (c *Client) UpdateAppHTTPSListener(blbId string, args *UpdateAppHTTPSListenerArgs) error {
	if args == nil || args.ListenerPort == 0 {
		return fmt.Errorf("unset listener port")
	}

	if len(args.Scheduler) == 0 {
		return fmt.Errorf("unset scheduler")
	}

	if len(args.CertIds) == 0 {
		return fmt.Errorf("unset certIds")
	}

	return bce.NewRequestBuilder(c).
		WithMethod(http.PUT).
		WithURL(getAppHTTPSListenerUri(blbId)).
		WithQueryParam("listenerPort", strconv.Itoa(int(args.ListenerPort))).
		WithQueryParamFilter("clientToken", args.ClientToken).
		WithBody(args).
		Do()
}

// UpdateAppSSLListener - update a SSL Listener
//
// PARAMS:
//     - blbId: LoadBalancer's ID
//     - args: parameters to update SSL Listener
// RETURNS:
//     - error: nil if ok otherwise the specific error
func (c *Client) UpdateAppSSLListener(blbId string, args *UpdateAppSSLListenerArgs) error {
	if args == nil || args.ListenerPort == 0 {
		return fmt.Errorf("unset listener port")
	}

	if len(args.Scheduler) == 0 {
		return fmt.Errorf("unset scheduler")
	}

	if len(args.CertIds) == 0 {
		return fmt.Errorf("unset certIds")
	}

	return bce.NewRequestBuilder(c).
		WithMethod(http.PUT).
		WithURL(getAppSSLListenerUri(blbId)).
		WithQueryParam("listenerPort", strconv.Itoa(int(args.ListenerPort))).
		WithQueryParamFilter("clientToken", args.ClientToken).
		WithBody(args).
		Do()
}

// DescribeAppTCPListeners - describe all TCP Listeners
//
// PARAMS:
//     - blbId: LoadBalancer's ID
//     - args: parameters to describe all TCP Listeners
// RETURNS:
//     - *DescribeAppTCPListenersResult: the result of describe all TCP Listeners
//     - error: nil if ok otherwise the specific error
func (c *Client) DescribeAppTCPListeners(blbId string, args *DescribeAppListenerArgs) (*DescribeAppTCPListenersResult, error) {
	if args == nil {
		args = &DescribeAppListenerArgs{}
	}

	if args.MaxKeys <= 0 || args.MaxKeys > 1000 {
		args.MaxKeys = 1000
	}

	result := &DescribeAppTCPListenersResult{}
	request := bce.NewRequestBuilder(c).
		WithMethod(http.GET).
		WithURL(getAppTCPListenerUri(blbId)).
		WithQueryParamFilter("marker", args.Marker).
		WithQueryParam("maxKeys", strconv.Itoa(args.MaxKeys)).
		WithResult(result)

	if args.ListenerPort != 0 {
		request.WithQueryParam("listenerPort", strconv.Itoa(int(args.ListenerPort)))
	}

	err := request.Do()
	return result, err
}

// DescribeAppUDPListeners - describe all UDP Listeners
//
// PARAMS:
//     - blbId: LoadBalancer's ID
//     - args: parameters to describe all UDP Listeners
// RETURNS:
//     - *DescribeAppUDPListenersResult: the result of describe all UDP Listeners
//     - error: nil if ok otherwise the specific error
func (c *Client) DescribeAppUDPListeners(blbId string, args *DescribeAppListenerArgs) (*DescribeAppUDPListenersResult, error) {
	if args == nil {
		args = &DescribeAppListenerArgs{}
	}

	if args.MaxKeys <= 0 || args.MaxKeys > 1000 {
		args.MaxKeys = 1000
	}

	result := &DescribeAppUDPListenersResult{}
	request := bce.NewRequestBuilder(c).
		WithMethod(http.GET).
		WithURL(getAppUDPListenerUri(blbId)).
		WithQueryParamFilter("marker", args.Marker).
		WithQueryParam("maxKeys", strconv.Itoa(args.MaxKeys)).
		WithResult(result)

	if args.ListenerPort != 0 {
		request.WithQueryParam("listenerPort", strconv.Itoa(int(args.ListenerPort)))
	}

	err := request.Do()
	return result, err
}

// DescribeAppHTTPListeners - describe all HTTP Listeners
//
// PARAMS:
//     - blbId: LoadBalancer's ID
//     - args: parameters to describe all HTTP Listeners
// RETURNS:
//     - *DescribeAppHTTPListenersResult: the result of describe all HTTP Listeners
//     - error: nil if ok otherwise the specific error
func (c *Client) DescribeAppHTTPListeners(blbId string, args *DescribeAppListenerArgs) (*DescribeAppHTTPListenersResult, error) {
	if args == nil {
		args = &DescribeAppListenerArgs{}
	}

	if args.MaxKeys <= 0 || args.MaxKeys > 1000 {
		args.MaxKeys = 1000
	}

	result := &DescribeAppHTTPListenersResult{}
	request := bce.NewRequestBuilder(c).
		WithMethod(http.GET).
		WithURL(getAppHTTPListenerUri(blbId)).
		WithQueryParamFilter("marker", args.Marker).
		WithQueryParam("maxKeys", strconv.Itoa(args.MaxKeys)).
		WithResult(result)

	if args.ListenerPort != 0 {
		request.WithQueryParam("listenerPort", strconv.Itoa(int(args.ListenerPort)))
	}

	err := request.Do()
	return result, err
}

// DescribeAppHTTPSListeners - describe all HTTPS Listeners
//
// PARAMS:
//     - blbId: LoadBalancer's ID
//     - args: parameters to describe all HTTPS Listeners
// RETURNS:
//     - *DescribeAppHTTPSListenersResult: the result of describe all HTTPS Listeners
//     - error: nil if ok otherwise the specific error
func (c *Client) DescribeAppHTTPSListeners(blbId string, args *DescribeAppListenerArgs) (*DescribeAppHTTPSListenersResult, error) {
	if args == nil {
		args = &DescribeAppListenerArgs{}
	}

	if args.MaxKeys <= 0 || args.MaxKeys > 1000 {
		args.MaxKeys = 1000
	}

	result := &DescribeAppHTTPSListenersResult{}
	request := bce.NewRequestBuilder(c).
		WithMethod(http.GET).
		WithURL(getAppHTTPSListenerUri(blbId)).
		WithQueryParamFilter("marker", args.Marker).
		WithQueryParam("maxKeys", strconv.Itoa(args.MaxKeys)).
		WithResult(result)

	if args.ListenerPort != 0 {
		request.WithQueryParam("listenerPort", strconv.Itoa(int(args.ListenerPort)))
	}

	err := request.Do()
	return result, err
}

// DescribeAppSSLListeners - describe all SSL Listeners
//
// PARAMS:
//     - blbId: LoadBalancer's ID
//     - args: parameters to describe all SSL Listeners
// RETURNS:
//     - *DescribeAppSSLListenersResult: the result of describe all SSL Listeners
//     - error: nil if ok otherwise the specific error
func (c *Client) DescribeAppSSLListeners(blbId string, args *DescribeAppListenerArgs) (*DescribeAppSSLListenersResult, error) {
	if args == nil {
		args = &DescribeAppListenerArgs{}
	}

	if args.MaxKeys <= 0 || args.MaxKeys > 1000 {
		args.MaxKeys = 1000
	}

	result := &DescribeAppSSLListenersResult{}
	request := bce.NewRequestBuilder(c).
		WithMethod(http.GET).
		WithURL(getAppSSLListenerUri(blbId)).
		WithQueryParamFilter("marker", args.Marker).
		WithQueryParam("maxKeys", strconv.Itoa(args.MaxKeys)).
		WithResult(result)

	if args.ListenerPort != 0 {
		request.WithQueryParam("listenerPort", strconv.Itoa(int(args.ListenerPort)))
	}

	err := request.Do()
	return result, err
}

// DeleteAppListeners - delete Listeners
//
// PARAMS:
//     - blbId: LoadBalancer's ID
//     - args: parameters to delete Listeners, a listener port list
// RETURNS:
//     - error: nil if ok otherwise the specific error
func (c *Client) DeleteAppListeners(blbId string, args *DeleteAppListenersArgs) error {
	if args == nil {
		return fmt.Errorf("unset args")
	}

	if len(args.PortList) == 0 {
		return fmt.Errorf("unset port list")
	}

	return bce.NewRequestBuilder(c).
		WithMethod(http.PUT).
		WithURL(getAppListenerUri(blbId)).
		WithQueryParamFilter("clientToken", args.ClientToken).
		WithQueryParam("batchdelete", "").
		WithBody(args).
		Do()
}

// CreatePolicys - create a policy bind with Listener
//
// PARAMS:
//     - blbId: LoadBalancer's ID
//     - args: parameters to create a policy
// RETURNS:
//     - error: nil if ok otherwise the specific error
func (c *Client) CreatePolicys(blbId string, args *CreatePolicysArgs) error {
	if args == nil {
		return fmt.Errorf("unset args")
	}

	if args.ListenerPort == 0 {
		return fmt.Errorf("unset listen port")
	}

	if len(args.AppPolicyVos) == 0 {
		return fmt.Errorf("unset App Policy")
	}

	return bce.NewRequestBuilder(c).
		WithMethod(http.POST).
		WithURL(getPolicysUrl(blbId)).
		WithQueryParamFilter("clientToken", args.ClientToken).
		WithBody(args).
		Do()
}

// DescribePolicys - descirbe a policy
//
// PARAMS:
//     - blbId: LoadBalancer's ID
//     - args: parameters to create a policy
// RETURNS:
//     - error: nil if ok otherwise the specific error
func (c *Client) DescribePolicys(blbId string, args *DescribePolicysArgs) (*DescribePolicysResult, error) {
	if args == nil {
		return nil, fmt.Errorf("unset args")
	}

	if args.Port == 0 {
		return nil, fmt.Errorf("unset port")
	}

	if args.MaxKeys > 1000 || args.MaxKeys <= 0 {
		args.MaxKeys = 1000
	}

	result := &DescribePolicysResult{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.GET).
		WithURL(getPolicysUrl(blbId)).
		WithQueryParam("port", strconv.Itoa(int(args.Port))).
		WithQueryParamFilter("marker", args.Marker).
		WithQueryParamFilter("maxKeys", strconv.Itoa(args.MaxKeys)).
		WithResult(result).
		Do()

	return result, err
}

// DeletePolicys - delete a policy
//
// PARAMS:
//     - blbId: LoadBalancer's ID
//     - args: parameters to delete a policy
// RETURNS:
//     - error: nil if ok otherwise the specific error
func (c *Client) DeletePolicys(blbId string, args *DeletePolicysArgs) error {
	if args == nil {
		return fmt.Errorf("unset args")
	}

	if args.Port == 0 {
		return fmt.Errorf("unset port")
	}

	if len(args.PolicyIdList) == 0 {
		return fmt.Errorf("unset policy id list")
	}

	return bce.NewRequestBuilder(c).
		WithMethod(http.PUT).
		WithURL(getPolicysUrl(blbId)).
		WithQueryParamFilter("clientToken", args.ClientToken).
		WithQueryParam("batchdelete", "").
		WithBody(args).
		Do()
}
