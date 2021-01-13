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

// client.go - define the client for Application LoadBalance service

// Package appblb defines the Application BLB services of BCE. The supported APIs are all defined in sub-package
package appblb

import "github.com/baidubce/bce-sdk-go/bce"

const (
	DEFAULT_SERVICE_DOMAIN = "blb." + bce.DEFAULT_REGION + ".baidubce.com"
	URI_PREFIX             = bce.URI_PREFIX + "v1"
	REQUEST_APPBLB_URL     = "/appblb"

	APP_SERVER_GROUP_URL      = "/appservergroup"
	APP_SERVER_GROUP_PORT_URL = "/appservergroupport"
	BLB_RS_URL                = "/blbrs"
	BLB_RS_MOUNT_URL          = "/blbrsmount"
	BLB_RS_UNMOUNT_URL        = "/blbrsunmount"

	APP_LISTENER_URL      = "/listener"
	APP_TCPLISTENER_URL   = "/TCPlistener"
	APP_UDPLISTENER_URL   = "/UDPlistener"
	APP_HTTPLISTENER_URL  = "/HTTPlistener"
	APP_HTTPSLISTENER_URL = "/HTTPSlistener"
	APP_SSLLISTENER_URL   = "/SSLlistener"

	POLICYS_URL = "/policys"

	APP_IP_GROUP_URL      = "/ipgroup"
	APP_IP_GROUP_BACKEND_POLICY_URL = "/ipgroup/backendpolicy"
	APP_IP_GROUP_MEMBER_URL = "/ipgroup/member"
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

func getAppBlbUri() string {
	return URI_PREFIX + REQUEST_APPBLB_URL
}

func getAppBlbUriWithId(id string) string {
	return URI_PREFIX + REQUEST_APPBLB_URL + "/" + id
}

func getAppServerGroupUri(id string) string {
	return URI_PREFIX + REQUEST_APPBLB_URL + "/" + id + APP_SERVER_GROUP_URL
}

func getAppServerGroupPortUri(id string) string {
	return URI_PREFIX + REQUEST_APPBLB_URL + "/" + id + APP_SERVER_GROUP_PORT_URL
}

func getBlbRsUri(id string) string {
	return URI_PREFIX + REQUEST_APPBLB_URL + "/" + id + BLB_RS_URL
}

func getBlbRsMountUri(id string) string {
	return URI_PREFIX + REQUEST_APPBLB_URL + "/" + id + BLB_RS_MOUNT_URL
}

func getBlbRsUnMountUri(id string) string {
	return URI_PREFIX + REQUEST_APPBLB_URL + "/" + id + BLB_RS_UNMOUNT_URL
}

func getAppListenerUri(id string) string {
	return URI_PREFIX + REQUEST_APPBLB_URL + "/" + id + APP_LISTENER_URL
}

func getAppTCPListenerUri(id string) string {
	return URI_PREFIX + REQUEST_APPBLB_URL + "/" + id + APP_TCPLISTENER_URL
}

func getAppUDPListenerUri(id string) string {
	return URI_PREFIX + REQUEST_APPBLB_URL + "/" + id + APP_UDPLISTENER_URL
}

func getAppHTTPListenerUri(id string) string {
	return URI_PREFIX + REQUEST_APPBLB_URL + "/" + id + APP_HTTPLISTENER_URL
}

func getAppHTTPSListenerUri(id string) string {
	return URI_PREFIX + REQUEST_APPBLB_URL + "/" + id + APP_HTTPSLISTENER_URL
}

func getAppSSLListenerUri(id string) string {
	return URI_PREFIX + REQUEST_APPBLB_URL + "/" + id + APP_SSLLISTENER_URL
}

func getPolicysUrl(id string) string {
	return URI_PREFIX + REQUEST_APPBLB_URL + "/" + id + POLICYS_URL
}

func getAppIpGroupUri(id string) string {
	return URI_PREFIX + REQUEST_APPBLB_URL + "/" + id + APP_IP_GROUP_URL
}

func getAppIpGroupBackendPolicyUri(id string) string {
	return URI_PREFIX + REQUEST_APPBLB_URL + "/" + id + APP_IP_GROUP_BACKEND_POLICY_URL
}

func getAppIpGroupMemberUri(id string) string {
	return URI_PREFIX + REQUEST_APPBLB_URL + "/" + id + APP_IP_GROUP_MEMBER_URL
}