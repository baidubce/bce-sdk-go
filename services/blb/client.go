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

// client.go - define the client for Application LoadBalance service

// Package blb defines the Normal BLB services of BCE. The supported APIs are all defined in sub-package
package blb

import "github.com/baidubce/bce-sdk-go/bce"

const (
	DEFAULT_SERVICE_DOMAIN = "blb." + bce.DEFAULT_REGION + ".baidubce.com"
	URI_PREFIX             = bce.URI_PREFIX + "v1"
	REQUEST_BLB_URL        = "/blb"

	LISTENER_URL      = "/listener"
	TCPLISTENER_URL   = "/TCPlistener"
	UDPLISTENER_URL   = "/UDPlistener"
	HTTPLISTENER_URL  = "/HTTPlistener"
	HTTPSLISTENER_URL = "/HTTPSlistener"
	SSLLISTENER_URL   = "/SSLlistener"

	BACKENDSERVER_URL = "/backendserver"

	REQUEST_BLB_CLUSTER_URL       = "/blbcluster"
	SECURITY_GROUP_URL            = "/securitygroup"
	ENTERPRISE_SECURITY_GROUP_URL = "/enterprise/securitygroup"

	ActionBlbToPostpaid       = "TO_POSTPAY"
	ActionBLbCancelToPostpaid = "CANCEL_TO_POSTPAY"
	ActionBlbToPrepaid        = "TO_PREPAY"
	ActionBlbResize           = "RESIZE"
)

// Client of APPBLB service is a kind of BceClient, so derived from BceClient
type Client struct {
	*bce.BceClient
}

func NewClient(ak, sk, endPoint string) (*Client, error) {
	if endPoint == "" {
		endPoint = DEFAULT_SERVICE_DOMAIN
	}
	client, err := bce.NewBceClientWithAkSk(ak, sk, endPoint)
	if err != nil {
		return nil, err
	}
	return &Client{client}, nil
}

func getBlbUri() string {
	return URI_PREFIX + REQUEST_BLB_URL
}

func getBlbUriWithId(id string) string {
	return URI_PREFIX + REQUEST_BLB_URL + "/" + id
}

func getBlbAclUriWithId(id string) string {
	return URI_PREFIX + REQUEST_BLB_URL + "/acl/" + id
}

func getBlbAutoRenewUriWithId(id string) string {
	return URI_PREFIX + REQUEST_BLB_URL + "/autoRenew/" + id
}

func getBlbRefundUriWithId(id string) string {
	return URI_PREFIX + REQUEST_BLB_URL + "/refund/" + id
}

func getBLbChargeUriWithId(id string) string {
	return URI_PREFIX + REQUEST_BLB_URL + "/" + id + "/charge"
}

func getListenerUri(id string) string {
	return URI_PREFIX + REQUEST_BLB_URL + "/" + id + LISTENER_URL
}

func getTCPListenerUri(id string) string {
	return URI_PREFIX + REQUEST_BLB_URL + "/" + id + TCPLISTENER_URL
}

func getUDPListenerUri(id string) string {
	return URI_PREFIX + REQUEST_BLB_URL + "/" + id + UDPLISTENER_URL
}

func getHTTPListenerUri(id string) string {
	return URI_PREFIX + REQUEST_BLB_URL + "/" + id + HTTPLISTENER_URL
}

func getHTTPSListenerUri(id string) string {
	return URI_PREFIX + REQUEST_BLB_URL + "/" + id + HTTPSLISTENER_URL
}

func getSSLListenerUri(id string) string {
	return URI_PREFIX + REQUEST_BLB_URL + "/" + id + SSLLISTENER_URL
}

func getBackendServerUri(id string) string {
	return URI_PREFIX + REQUEST_BLB_URL + "/" + id + BACKENDSERVER_URL
}

func getBlbClusterUri() string {
	return URI_PREFIX + REQUEST_BLB_CLUSTER_URL
}

func getBlbClusterUriWithId(id string) string {
	return URI_PREFIX + REQUEST_BLB_CLUSTER_URL + "/" + id
}

func getSecurityGroupUri(id string) string {
	return URI_PREFIX + REQUEST_BLB_URL + "/" + id + SECURITY_GROUP_URL
}

func getEnterpriseSecurityGroupUri(id string) string {
	return URI_PREFIX + REQUEST_BLB_URL + "/" + id + ENTERPRISE_SECURITY_GROUP_URL
}
