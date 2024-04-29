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

// vpn.go - the vpn gateway APIs definition supported by the VPN service

package vpn

import (
	"fmt"
	"github.com/baidubce/bce-sdk-go/bce"
	"github.com/baidubce/bce-sdk-go/http"
	"strconv"
)

// CreateVPNGateway - create a new vpn gateway
//
// PARAMS:
//    - args: the arguments to create vpn gateway
// RETURNS:
//    - *CreateVpnGatewayResult: the id of the vpn gateway newly created
//    - error: nil if success otherwise the specific error

func (c *Client) CreateVpnGateway(args *CreateVpnGatewayArgs) (*CreateVpnGatewayResult, error) {
	if args == nil {
		return nil, fmt.Errorf("The createVpnGatewayArgs cannot be nil.")
	}

	result := &CreateVpnGatewayResult{}
	err := bce.NewRequestBuilder(c).
		WithURL(getURLForVPN()).
		WithMethod(http.POST).
		WithBody(args).
		WithQueryParamFilter("clientToken", args.ClientToken).
		WithResult(result).
		Do()

	return result, err
}

//
// ListVpn - list all vpn gateways with the specific parameters
// PARAMS:
//    - args: the arguments to list vpn gateways
// RETURNS:
//    - *ListVpnGatewayResult: the result of vpn gateway list, excluding tags.
//    - error: nil if success otherwise the specific error

func (c *Client) ListVpnGateway(args *ListVpnGatewayArgs) (*ListVpnGatewayResult, error) {
	if args == nil {
		return nil, fmt.Errorf("The listVpnGatewayArgs cannot be nil.")
	}
	if args.MaxKeys == 0 {
		args.MaxKeys = 1000
	}

	result := &ListVpnGatewayResult{}
	err := bce.NewRequestBuilder(c).
		WithURL(getURLForVPN()).
		WithMethod(http.GET).
		WithQueryParam("vpcId", args.VpcId).
		WithQueryParamFilter("eip", args.Eip).
		WithQueryParamFilter("marker", args.Marker).
		WithQueryParamFilter("maxKeys", strconv.Itoa(args.MaxKeys)).
		WithResult(result).
		Do()

	return result, err
}

// DeleteVpnGateway - delete the specific vpn gateway
//
// PARAMS:
//    - vpnId: the id of the specific vpn gateway
//    - clientToken: the idempotent token
// RETURNS:
//    - error: nil if success otherwise the specific error

func (c *Client) DeleteVpn(vpnId, clientToken string) error {
	return bce.NewRequestBuilder(c).
		WithURL(getURLForVPNId(vpnId)).
		WithMethod(http.DELETE).
		WithQueryParamFilter("clientToken", clientToken).
		Do()
}

// DeleteVpnGateway - delete the specific vpn gateway
//
// PARAMS:
//     - vpnId: the id of the specific vpn gateway
//     - clientToken: the idempotent token
// RETURNS:
//     - error: nil if success otherwise the specific error
func (c *Client) DeleteVpnGateway(vpcId, clientToken string) error {
	return bce.NewRequestBuilder(c).
		WithURL(getURLForVPNId(vpcId)).
		WithMethod(http.DELETE).
		WithQueryParamFilter("clientToken", clientToken).
		Do()
}

// GetVpnGatewayDetail - get details of the specific vpn gateway
//
// PARAMS:
//     - vpnId: the id of the specified vpn
// RETURNS:
//     - *VPN: the result of the specific vpn gateway details
//     - error: nil if success otherwise the specific error
func (c *Client) GetVpnGatewayDetail(vpnId string) (*VPN, error) {
	result := &VPN{}
	err := bce.NewRequestBuilder(c).
		WithURL(getURLForVPNId(vpnId)).
		WithMethod(http.GET).
		WithResult(result).
		Do()

	return result, err
}

// UpdateVpnGateway - update the specified vpn gateway
//
// PARAMS:
//     - vpnId: the id of the specific vpn gateway
//     - args: the arguments to update vpn gateway
// RETURNS:
//     - error: nil if success otherwise the specific error
func (c *Client) UpdateVpnGateway(vpnId string, args *UpdateVpnGatewayArgs) error {
	if args == nil {
		return fmt.Errorf("The updateVpnGatewayArgs cannot be nil.")
	}

	return bce.NewRequestBuilder(c).
		WithURL(getURLForVPNId(vpnId)).
		WithMethod(http.PUT).
		WithBody(args).
		WithQueryParam("modifyAttribute", "").
		WithQueryParamFilter("clientToken", args.ClientToken).
		Do()
}

// BindEip - bind eip for the specific vpn gateway
//
// PARAMS:
//     - vpnId: the id of the specific vpn gateway
//     - args: the arguments to bind eip
// RETURNS:
//     - error: nil if success otherwise the specific error
func (c *Client) BindEip(vpnId string, args *BindEipArgs) error {
	if args == nil {
		return fmt.Errorf("The bindEipArgs cannot be nil.")
	}
	return bce.NewRequestBuilder(c).
		WithURL(getURLForVPNId(vpnId)).
		WithMethod(http.PUT).
		WithBody(args).
		WithQueryParamFilter("clientToken", args.ClientToken).
		WithQueryParam("bind", "").
		Do()
}

// UnBindEips - unbind eip for the specific vpn gateway
//
// PARAMS:
//     - vpnId: the id of the specific vpn gateway
// RETURNS:
//     - error: nil if success otherwise the specific error
func (c *Client) UnBindEip(vpnId, clientToken string) error {
	return bce.NewRequestBuilder(c).
		WithURL(getURLForVPNId(vpnId)).
		WithMethod(http.PUT).
		WithQueryParamFilter("clientToken", clientToken).
		WithQueryParam("unbind", "").
		Do()
}

// RenewVpnGateway - renew vpn gateway with the specific parameters
//
// PARAMS:
//     - vpnId: the id of the specific vpn gateway
//     - args: the arguments to renew vpn gateway
// RETURNS:
//     - error: nil if success otherwise the specific error
func (c *Client) RenewVpnGateway(vpnId string, args *RenewVpnGatewayArgs) error {
	if args == nil {
		return fmt.Errorf("The renewVpnGatewayArgs cannot be nil.")
	}

	return bce.NewRequestBuilder(c).
		WithURL(getURLForVPNId(vpnId)).
		WithMethod(http.PUT).
		WithBody(args).
		WithQueryParamFilter("clientToken", args.ClientToken).
		WithQueryParam("purchaseReserved", "").
		Do()
}

// CreateVpnConn - create vpnconn with the specific parameters
//
// PARAMS:
//     - vpnId: the id of the specific vpn gateway
//     - args: the arguments to create vpnconn
// RETURNS:
//     - error: nil if success otherwise the specific error
func (c *Client) CreateVpnConn(args *CreateVpnConnArgs) (*CreateVpnConnResult, error) {
	if args == nil {
		return nil, fmt.Errorf("The CreateVpnConnArgs cannot be nil.")
	}
	result := &CreateVpnConnResult{}
	err := bce.NewRequestBuilder(c).
		WithURL(getURLForVPNId(args.VpnId) + "/vpnconn").
		WithMethod(http.POST).
		WithBody(args).
		WithResult(result).
		Do()
	return result, err
}

// UpdateVpnConn - create vpnconn with the specific parameters
//
// PARAMS:
//     - args: the arguments to update vpnconn
// RETURNS:
//     - error: nil if success otherwise the specific error
func (c *Client) UpdateVpnConn(args *UpdateVpnConnArgs) error {
	return bce.NewRequestBuilder(c).
		WithURL(getURLForVpnConnId(args.VpnConnId)).
		WithMethod(http.PUT).
		WithBody(args.UpdateVpnconn).
		Do()
}

// ListVpnConn - list vpnconn with the specific vpnId
//
// PARAMS:
//     - vpnId:the id you want to list vpnconn
// RETURNS:
//     - *ListVpnConnResult: the result of vpn gateway'conn list
//     - error: nil if success otherwise the specific error
func (c *Client) ListVpnConn(vpnId string) (*ListVpnConnResult, error) {
	result := &ListVpnConnResult{}
	err := bce.NewRequestBuilder(c).
		WithURL(getURLForVpnConn() + "/" + vpnId).
		WithMethod(http.GET).
		WithResult(result).
		Do()
	return result, err
}

// DeleteVpnConn - delete the specific vpnconn
//
// PARAMS:
//     - vpnConnId: the id of the specific vpnconn
//     - clientToken: the idempotent token
// RETURNS:
//     - error: nil if success otherwise the specific error
func (c *Client) DeleteVpnConn(vpnConnId, clientToken string) error {
	return bce.NewRequestBuilder(c).
		WithURL(getURLForVpnConnId(vpnConnId)).
		WithMethod(http.DELETE).
		WithQueryParamFilter("clientToken", clientToken).
		Do()
}

// CreateSslVpnServer - create ssl-vpn server with the specific parameters
//
// PARAMS:
//     - args: the arguments to create ssl-vpn server
// RETURNS:
//     - error: nil if success otherwise the specific error
func (c *Client) CreateSslVpnServer(args *CreateSslVpnServerArgs) (*CreateSslVpnServerResult, error) {
	if args == nil {
		return nil, fmt.Errorf("the CreateSslVpnServerArgs cannot be nil")
	}

	if args.InterfaceType == nil {
		defaultInterfaceType := "tap"
		args.InterfaceType = &defaultInterfaceType
	}
	result := &CreateSslVpnServerResult{}
	err := bce.NewRequestBuilder(c).
		WithURL(getURLForSslVpnServerByVpnId(args.VpnId)).
		WithMethod(http.POST).
		WithQueryParamFilter("clientToken", args.ClientToken).
		WithBody(args).
		WithResult(result).
		Do()
	return result, err
}

// UpdateSslVpnServer - update ssl-vpn server with the specific parameters
//
// PARAMS:
//     - args: the arguments to update ssl ssl-vpn server
// RETURNS:
//     - error: nil if success otherwise the specific error
func (c *Client) UpdateSslVpnServer(args *UpdateSslVpnServerArgs) error {
	if args == nil {
		return fmt.Errorf("the UpdateSslVpnServerArgs cannot be nil")
	}
	return bce.NewRequestBuilder(c).
		WithURL(getURLForSslVpnServerByVpnId(args.VpnId)+"/"+args.SslVpnServerId).
		WithMethod(http.PUT).
		WithQueryParamFilter("clientToken", args.ClientToken).
		WithBody(args.UpdateSslVpnServer).
		Do()
}

// GetSslVpnServer - get details of the specific ssl-vpn server
//
// PARAMS:
//     - vpnId: the id of the specified vpn
// RETURNS:
//     - *SslVpnServer: the result of the specific ssl-vpn server details
//     - error: nil if success otherwise the specific error
func (c *Client) GetSslVpnServer(vpnId, clientToken string) (*SslVpnServer, error) {
	result := &SslVpnServer{}
	err := bce.NewRequestBuilder(c).
		WithURL(getURLForSslVpnServerByVpnId(vpnId)).
		WithMethod(http.GET).
		WithQueryParamFilter("clientToken", clientToken).
		WithResult(result).
		Do()

	return result, err
}

// DeleteSslVpnServer - delete ssl-vpn server with the specific vpnId
//
// PARAMS:
//     - vpnId: the id of the specific vpn gateway
//	   - sslVpnServerId: the id of the specific ssl-vpn server
//     - clientToken: the idempotent token
// RETURNS:
//     - error: nil if success otherwise the specific error
func (c *Client) DeleteSslVpnServer(vpnId, sslVpnServerId, clientToken string) error {
	return bce.NewRequestBuilder(c).
		WithURL(getURLForSslVpnServerByVpnId(vpnId)+"/"+sslVpnServerId).
		WithMethod(http.DELETE).
		WithQueryParamFilter("clientToken", clientToken).
		Do()
}

// BatchCreateSslVpnUser - batch create ssl-vpn user with the specific parameters
//
// PARAMS:
//     - args: the arguments to create ssl-vpn users
// RETURNS:
//     - error: nil if success otherwise the specific error
func (c *Client) BatchCreateSslVpnUser(args *BatchCreateSslVpnUserArgs) (*BatchCreateSslVpnUserResult, error) {
	if args == nil {
		return nil, fmt.Errorf("the BatchCreateSslVpnUserArgs cannot be nil")
	}
	result := &BatchCreateSslVpnUserResult{}
	err := bce.NewRequestBuilder(c).
		WithURL(getURLForSslVpnUserByVpnId(args.VpnId)).
		WithMethod(http.POST).
		WithQueryParamFilter("clientToken", args.ClientToken).
		WithBody(args).
		WithResult(result).
		Do()
	return result, err
}

// UpdateSslVpnUser - update ssl-vpn user with the specific parameters
//
// PARAMS:
//     - args: the arguments to update the specific ssl-vpn user
// RETURNS:
//     - error: nil if success otherwise the specific error
func (c *Client) UpdateSslVpnUser(args *UpdateSslVpnUserArgs) error {
	if args == nil {
		return fmt.Errorf("the UpdateSslVpnUserArgs cannot be nil")
	}
	if len(args.SslVpnUser.Password) == 0 && args.SslVpnUser.Description == nil {
		return fmt.Errorf("both password and description are nil")
	}
	return bce.NewRequestBuilder(c).
		WithURL(getURLForSslVpnUserByVpnId(args.VpnId)+"/"+args.UserId).
		WithMethod(http.PUT).
		WithQueryParamFilter("clientToken", args.ClientToken).
		WithBody(args.SslVpnUser).
		Do()
}

// ListSslVpnUser - list ssl-vpn user with the specific vpnId
//
// PARAMS:
//     - args: the arguments to list ssl-vpn users
// RETURNS:
//     - *ListSslVpnUserResult: the result of vpn ssl-vpn user list contains page infos
//     - error: nil if success otherwise the specific error
func (c *Client) ListSslVpnUser(args *ListSslVpnUserArgs) (*ListSslVpnUserResult, error) {
	if args == nil {
		args = &ListSslVpnUserArgs{}
	}
	if args.MaxKeys < 0 || args.MaxKeys > 1000 {
		return nil, fmt.Errorf("the field maxKeys is out of range [0, 1000]")
	}
	result := &ListSslVpnUserResult{}
	builder := bce.NewRequestBuilder(c).
		WithURL(getURLForSslVpnUserByVpnId(args.VpnId)).
		WithMethod(http.GET).
		WithResult(result).
		WithQueryParamFilter("marker", args.Marker)
	if args.MaxKeys != 0 {
		builder.WithQueryParamFilter("maxKeys", strconv.Itoa(args.MaxKeys))
	}
	err := builder.Do()
	return result, err
}

// DeleteSslVpnUser - delete ssl-vpn user with the specific vpnId and userId
//
// PARAMS:
//     - vpnId: the id of the specific vpn gateway
//     - userId: the id of the specific ssl-vpn user
//     - clientToken: the idempotent token
// RETURNS:
//     - error: nil if success otherwise the specific error
func (c *Client) DeleteSslVpnUser(vpnId, userId, clientToken string) error {
	return bce.NewRequestBuilder(c).
		WithURL(getURLForSslVpnUserByVpnId(vpnId)+"/"+userId).
		WithMethod(http.DELETE).
		WithQueryParamFilter("clientToken", clientToken).
		Do()

}
