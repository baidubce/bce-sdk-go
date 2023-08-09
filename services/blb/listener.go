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

// listener.go - the Normal BLB Listener APIs definition supported by the BLB service

package blb

import (
	"fmt"
	"strconv"

	"github.com/baidubce/bce-sdk-go/bce"
	"github.com/baidubce/bce-sdk-go/http"
)

// CreateTCPListener - create a TCP Listener
//
// PARAMS:
//   - blbId: LoadBalancer's ID
//   - args: parameters to create TCP Listener
//
// RETURNS:
//   - error: nil if ok otherwise the specific error
func (c *Client) CreateTCPListener(blbId string, args *CreateTCPListenerArgs) error {
	if args == nil {
		return fmt.Errorf("unset args")
	}

	if args.ListenerPort == 0 {
		return fmt.Errorf("unsupport listener port")
	}

	if args.BackendPort == 0 {
		return fmt.Errorf("unsupport backend port")
	}

	if len(args.Scheduler) == 0 {
		return fmt.Errorf("unset scheduler")
	}

	return bce.NewRequestBuilder(c).
		WithMethod(http.POST).
		WithURL(getTCPListenerUri(blbId)).
		WithQueryParamFilter("clientToken", args.ClientToken).
		WithBody(args).
		Do()
}

// CreateUDPListener - create a UDP Listener
//
// PARAMS:
//   - blbId: LoadBalancer's ID
//   - args: parameters to create UDP Listener
//
// RETURNS:
//   - error: nil if ok otherwise the specific error
func (c *Client) CreateUDPListener(blbId string, args *CreateUDPListenerArgs) error {
	if args == nil {
		return fmt.Errorf("unset args")
	}

	if args.ListenerPort == 0 {
		return fmt.Errorf("unsupport listener port")
	}

	if args.BackendPort == 0 {
		return fmt.Errorf("unsupport backend port")
	}

	if len(args.Scheduler) == 0 {
		return fmt.Errorf("unset scheduler")
	}

	if len(args.HealthCheckString) == 0 {
		return fmt.Errorf("unset healthCheckString")
	}

	return bce.NewRequestBuilder(c).
		WithMethod(http.POST).
		WithURL(getUDPListenerUri(blbId)).
		WithQueryParamFilter("clientToken", args.ClientToken).
		WithBody(args).
		Do()
}

// CreateHTTPListener - create a HTTP Listener
//
// PARAMS:
//   - blbId: LoadBalancer's ID
//   - args: parameters to create HTTP Listener
//
// RETURNS:
//   - error: nil if ok otherwise the specific error
func (c *Client) CreateHTTPListener(blbId string, args *CreateHTTPListenerArgs) error {
	if args == nil {
		return fmt.Errorf("unset args")
	}

	if args.ListenerPort == 0 {
		return fmt.Errorf("unsupport listener port")
	}

	if args.BackendPort == 0 {
		return fmt.Errorf("unsupport backend port")
	}

	if len(args.Scheduler) == 0 {
		return fmt.Errorf("unset scheduler")
	}

	return bce.NewRequestBuilder(c).
		WithMethod(http.POST).
		WithURL(getHTTPListenerUri(blbId)).
		WithQueryParamFilter("clientToken", args.ClientToken).
		WithBody(args).
		Do()
}

// CreateHTTPSListener - create a HTTPS Listener
//
// PARAMS:
//   - blbId: LoadBalancer's ID
//   - args: parameters to create HTTPS Listener
//
// RETURNS:
//   - error: nil if ok otherwise the specific error
func (c *Client) CreateHTTPSListener(blbId string, args *CreateHTTPSListenerArgs) error {
	if args == nil {
		return fmt.Errorf("unset args")
	}

	if args.ListenerPort == 0 {
		return fmt.Errorf("unsupport listener port")
	}

	if args.BackendPort == 0 {
		return fmt.Errorf("unsupport backend port")
	}

	if len(args.Scheduler) == 0 {
		return fmt.Errorf("unset scheduler")
	}

	if len(args.CertIds) == 0 {
		return fmt.Errorf("unset certIds")
	}

	return bce.NewRequestBuilder(c).
		WithMethod(http.POST).
		WithURL(getHTTPSListenerUri(blbId)).
		WithQueryParamFilter("clientToken", args.ClientToken).
		WithBody(args).
		Do()
}

// CreateAppSSLListener - create a SSL Listener
//
// PARAMS:
//   - blbId: LoadBalancer's ID
//   - args: parameters to create SSL Listener
//
// RETURNS:
//   - error: nil if ok otherwise the specific error
func (c *Client) CreateSSLListener(blbId string, args *CreateSSLListenerArgs) error {
	if args == nil {
		return fmt.Errorf("unset args")
	}

	if args.ListenerPort == 0 {
		return fmt.Errorf("unsupport listener port")
	}

	if args.BackendPort == 0 {
		return fmt.Errorf("unsupport backend port")
	}

	if len(args.Scheduler) == 0 {
		return fmt.Errorf("unset scheduler")
	}

	if len(args.CertIds) == 0 {
		return fmt.Errorf("unset certIds")
	}

	return bce.NewRequestBuilder(c).
		WithMethod(http.POST).
		WithURL(getSSLListenerUri(blbId)).
		WithQueryParamFilter("clientToken", args.ClientToken).
		WithBody(args).
		Do()
}

// UpdateTCPListener - update a TCP Listener
//
// PARAMS:
//   - blbId: LoadBalancer's ID
//   - args: parameters to update TCP Listener
//
// RETURNS:
//   - error: nil if ok otherwise the specific error
func (c *Client) UpdateTCPListener(blbId string, args *UpdateTCPListenerArgs) error {
	if args == nil || args.ListenerPort == 0 {
		return fmt.Errorf("unset listener port")
	}

	return bce.NewRequestBuilder(c).
		WithMethod(http.PUT).
		WithURL(getTCPListenerUri(blbId)).
		WithQueryParam("listenerPort", strconv.Itoa(int(args.ListenerPort))).
		WithQueryParamFilter("clientToken", args.ClientToken).
		WithBody(args).
		Do()
}

// UpdateUDPListener - update a UDP Listener
//
// PARAMS:
//   - blbId: LoadBalancer's ID
//   - args: parameters to update UDP Listener
//
// RETURNS:
//   - error: nil if ok otherwise the specific error
func (c *Client) UpdateUDPListener(blbId string, args *UpdateUDPListenerArgs) error {
	if args == nil || args.ListenerPort == 0 {
		return fmt.Errorf("unset listener port")
	}

	return bce.NewRequestBuilder(c).
		WithMethod(http.PUT).
		WithURL(getUDPListenerUri(blbId)).
		WithQueryParam("listenerPort", strconv.Itoa(int(args.ListenerPort))).
		WithQueryParamFilter("clientToken", args.ClientToken).
		WithBody(args).
		Do()
}

// UpdateHTTPListener - update a HTTP Listener
//
// PARAMS:
//   - blbId: LoadBalancer's ID
//   - args: parameters to update HTTP Listener
//
// RETURNS:
//   - error: nil if ok otherwise the specific error
func (c *Client) UpdateHTTPListener(blbId string, args *UpdateHTTPListenerArgs) error {
	if args == nil || args.ListenerPort == 0 {
		return fmt.Errorf("unset listener port")
	}

	return bce.NewRequestBuilder(c).
		WithMethod(http.PUT).
		WithURL(getHTTPListenerUri(blbId)).
		WithQueryParam("listenerPort", strconv.Itoa(int(args.ListenerPort))).
		WithQueryParamFilter("clientToken", args.ClientToken).
		WithBody(args).
		Do()
}

// UpdateHTTPSListener - update a HTTPS Listener
//
// PARAMS:
//   - blbId: LoadBalancer's ID
//   - args: parameters to update HTTPS Listener
//
// RETURNS:
//   - error: nil if ok otherwise the specific error
func (c *Client) UpdateHTTPSListener(blbId string, args *UpdateHTTPSListenerArgs) error {
	if args == nil || args.ListenerPort == 0 {
		return fmt.Errorf("unset listener port")
	}

	return bce.NewRequestBuilder(c).
		WithMethod(http.PUT).
		WithURL(getHTTPSListenerUri(blbId)).
		WithQueryParam("listenerPort", strconv.Itoa(int(args.ListenerPort))).
		WithQueryParamFilter("clientToken", args.ClientToken).
		WithBody(args).
		Do()
}

// UpdateSSLListener - update a SSL Listener
//
// PARAMS:
//   - blbId: LoadBalancer's ID
//   - args: parameters to update SSL Listener
//
// RETURNS:
//   - error: nil if ok otherwise the specific error
func (c *Client) UpdateSSLListener(blbId string, args *UpdateSSLListenerArgs) error {
	if args == nil || args.ListenerPort == 0 {
		return fmt.Errorf("unset listener port")
	}

	return bce.NewRequestBuilder(c).
		WithMethod(http.PUT).
		WithURL(getSSLListenerUri(blbId)).
		WithQueryParam("listenerPort", strconv.Itoa(int(args.ListenerPort))).
		WithQueryParamFilter("clientToken", args.ClientToken).
		WithBody(args).
		Do()
}

// DescribeTCPListeners - describe all TCP Listeners
//
// PARAMS:
//   - blbId: LoadBalancer's ID
//   - args: parameters to describe all TCP Listeners
//
// RETURNS:
//   - *DescribeTCPListenersResult: the result of describe all TCP Listeners
//   - error: nil if ok otherwise the specific error
func (c *Client) DescribeTCPListeners(blbId string, args *DescribeListenerArgs) (*DescribeTCPListenersResult, error) {
	if args == nil {
		args = &DescribeListenerArgs{}
	}

	if args.MaxKeys <= 0 || args.MaxKeys > 1000 {
		args.MaxKeys = 1000
	}

	result := &DescribeTCPListenersResult{}
	request := bce.NewRequestBuilder(c).
		WithMethod(http.GET).
		WithURL(getTCPListenerUri(blbId)).
		WithQueryParamFilter("marker", args.Marker).
		WithQueryParam("maxKeys", strconv.Itoa(args.MaxKeys)).
		WithResult(result)

	if args.ListenerPort != 0 {
		request.WithQueryParam("listenerPort", strconv.Itoa(int(args.ListenerPort)))
	}

	err := request.Do()
	return result, err
}

// DescribeUDPListeners - describe all UDP Listeners
//
// PARAMS:
//   - blbId: LoadBalancer's ID
//   - args: parameters to describe all UDP Listeners
//
// RETURNS:
//   - *DescribeUDPListenersResult: the result of describe all UDP Listeners
//   - error: nil if ok otherwise the specific error
func (c *Client) DescribeUDPListeners(blbId string, args *DescribeListenerArgs) (*DescribeUDPListenersResult, error) {
	if args == nil {
		args = &DescribeListenerArgs{}
	}

	if args.MaxKeys <= 0 || args.MaxKeys > 1000 {
		args.MaxKeys = 1000
	}

	result := &DescribeUDPListenersResult{}
	request := bce.NewRequestBuilder(c).
		WithMethod(http.GET).
		WithURL(getUDPListenerUri(blbId)).
		WithQueryParamFilter("marker", args.Marker).
		WithQueryParam("maxKeys", strconv.Itoa(args.MaxKeys)).
		WithResult(result)

	if args.ListenerPort != 0 {
		request.WithQueryParam("listenerPort", strconv.Itoa(int(args.ListenerPort)))
	}

	err := request.Do()
	return result, err
}

// DescribeHTTPListeners - describe all HTTP Listeners
//
// PARAMS:
//   - blbId: LoadBalancer's ID
//   - args: parameters to describe all HTTP Listeners
//
// RETURNS:
//   - *DescribeHTTPListenersResult: the result of describe all HTTP Listeners
//   - error: nil if ok otherwise the specific error
func (c *Client) DescribeHTTPListeners(blbId string, args *DescribeListenerArgs) (*DescribeHTTPListenersResult, error) {
	if args == nil {
		args = &DescribeListenerArgs{}
	}

	if args.MaxKeys <= 0 || args.MaxKeys > 1000 {
		args.MaxKeys = 1000
	}

	result := &DescribeHTTPListenersResult{}
	request := bce.NewRequestBuilder(c).
		WithMethod(http.GET).
		WithURL(getHTTPListenerUri(blbId)).
		WithQueryParamFilter("marker", args.Marker).
		WithQueryParam("maxKeys", strconv.Itoa(args.MaxKeys)).
		WithResult(result)

	if args.ListenerPort != 0 {
		request.WithQueryParam("listenerPort", strconv.Itoa(int(args.ListenerPort)))
	}

	err := request.Do()
	return result, err
}

// DescribeHTTPSListeners - describe all HTTPS Listeners
//
// PARAMS:
//   - blbId: LoadBalancer's ID
//   - args: parameters to describe all HTTPS Listeners
//
// RETURNS:
//   - *DescribeHTTPSListenersResult: the result of describe all HTTPS Listeners
//   - error: nil if ok otherwise the specific error
func (c *Client) DescribeHTTPSListeners(blbId string, args *DescribeListenerArgs) (*DescribeHTTPSListenersResult, error) {
	if args == nil {
		args = &DescribeListenerArgs{}
	}

	if args.MaxKeys <= 0 || args.MaxKeys > 1000 {
		args.MaxKeys = 1000
	}

	result := &DescribeHTTPSListenersResult{}
	request := bce.NewRequestBuilder(c).
		WithMethod(http.GET).
		WithURL(getHTTPSListenerUri(blbId)).
		WithQueryParamFilter("marker", args.Marker).
		WithQueryParam("maxKeys", strconv.Itoa(args.MaxKeys)).
		WithResult(result)

	if args.ListenerPort != 0 {
		request.WithQueryParam("listenerPort", strconv.Itoa(int(args.ListenerPort)))
	}

	err := request.Do()
	return result, err
}

// DescribeSSLListeners - describe all SSL Listeners
//
// PARAMS:
//   - blbId: LoadBalancer's ID
//   - args: parameters to describe all SSL Listeners
//
// RETURNS:
//   - *DescribeSSLListenersResult: the result of describe all SSL Listeners
//   - error: nil if ok otherwise the specific error
func (c *Client) DescribeSSLListeners(blbId string, args *DescribeListenerArgs) (*DescribeSSLListenersResult, error) {
	if args == nil {
		args = &DescribeListenerArgs{}
	}

	if args.MaxKeys <= 0 || args.MaxKeys > 1000 {
		args.MaxKeys = 1000
	}

	result := &DescribeSSLListenersResult{}
	request := bce.NewRequestBuilder(c).
		WithMethod(http.GET).
		WithURL(getSSLListenerUri(blbId)).
		WithQueryParamFilter("marker", args.Marker).
		WithQueryParam("maxKeys", strconv.Itoa(args.MaxKeys)).
		WithResult(result)

	if args.ListenerPort != 0 {
		request.WithQueryParam("listenerPort", strconv.Itoa(int(args.ListenerPort)))
	}

	err := request.Do()
	return result, err
}

// DescribeAllListeners - describe all Listeners
//
// PARAMS:
//   - blbId: LoadBalancer's ID
//   - args: parameters to describe all Listeners
//
// RETURNS:
//   - *DescribeAllListenersResult: the result of describe all Listeners
//   - error: nil if ok otherwise the specific error
func (c *Client) DescribeAllListeners(blbId string, args *DescribeListenerArgs) (*DescribeAllListenersResult, error) {
	if args == nil {
		args = &DescribeListenerArgs{}
	}

	if args.MaxKeys <= 0 || args.MaxKeys > 1000 {
		args.MaxKeys = 1000
	}

	result := &DescribeAllListenersResult{}
	request := bce.NewRequestBuilder(c).
		WithMethod(http.GET).
		WithURL(getListenerUri(blbId)).
		WithQueryParamFilter("marker", args.Marker).
		WithQueryParam("maxKeys", strconv.Itoa(args.MaxKeys)).
		WithResult(result)

	if args.ListenerPort != 0 {
		request.WithQueryParam("listenerPort", strconv.Itoa(int(args.ListenerPort)))
	}

	err := request.Do()
	return result, err
}

// DeleteListeners - delete Listeners
//
// PARAMS:
//   - blbId: LoadBalancer's ID
//   - args: parameters to delete Listeners, a listener port list
//
// RETURNS:
//   - error: nil if ok otherwise the specific error
func (c *Client) DeleteListeners(blbId string, args *DeleteListenersArgs) error {
	if args == nil {
		return fmt.Errorf("unset args")
	}

	if len(args.PortList) == 0 && len(args.PortTypeList) == 0 {
		return fmt.Errorf("unset port list")
	}

	return bce.NewRequestBuilder(c).
		WithMethod(http.PUT).
		WithURL(getListenerUri(blbId)).
		WithQueryParamFilter("clientToken", args.ClientToken).
		WithQueryParam("batchdelete", "").
		WithBody(args).
		Do()
}
