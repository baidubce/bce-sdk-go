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

// peerconn.go - the peer connection APIs definition supported by the VPC service

package vpc

import (
	"fmt"
	"strconv"

	"github.com/baidubce/bce-sdk-go/bce"
	"github.com/baidubce/bce-sdk-go/http"
)

// CreatePeerConn - create a new peer connection with the specific parameters
//
// PARAMS:
//   - args: the arguments to create peer connection
//
// RETURNS:
//   - *CreatePeerConnResult: the id of peer connection newly created
//   - error: nil if success otherwise the specific error
func (c *Client) CreatePeerConn(args *CreatePeerConnArgs) (*CreatePeerConnResult, error) {
	if args == nil {
		return nil, fmt.Errorf("The createPeerConnArgs cannot be nil.")
	}

	result := &CreatePeerConnResult{}
	err := bce.NewRequestBuilder(c).
		WithURL(getURLForPeerConn()).
		WithMethod(http.POST).
		WithBody(args).
		WithQueryParamFilter("clientToken", args.ClientToken).
		WithResult(result).
		Do()

	return result, err
}

// ListPeerConn - list all peer connections with the specific parameters
//
// PARAMS:
//   - args: the arguments to list peer connections
//
// RETURNS:
//   - *ListPeerConnsResult: the result of the peer connection list
//   - error: nil if success otherwise the specific error
func (c *Client) ListPeerConn(args *ListPeerConnsArgs) (*ListPeerConnsResult, error) {
	if args == nil {
		args = &ListPeerConnsArgs{}
	}
	if args.MaxKeys == 0 {
		args.MaxKeys = 1000
	}

	result := &ListPeerConnsResult{}
	err := bce.NewRequestBuilder(c).
		WithURL(getURLForPeerConn()).
		WithMethod(http.GET).
		WithQueryParamFilter("vpcId", args.VpcId).
		WithQueryParamFilter("marker", args.Marker).
		WithQueryParamFilter("maxKeys", strconv.Itoa(args.MaxKeys)).
		WithResult(result).
		Do()

	return result, err
}

// GetPeerConnDetail - get details for the specific peer connection
//
// PARAMS:
//   - peerConnId: the id of the specific peer connection
//   - role: the role of the specific peer connection, which can be initiator or acceptor
//
// RETURNS:
//   - *PeerConn: the result of the specfic peer connection details
//   - error: nil if success otherwise the specific error
func (c *Client) GetPeerConnDetail(peerConnId string, role PeerConnRoleType) (*PeerConn, error) {
	result := &PeerConn{}
	err := bce.NewRequestBuilder(c).
		WithURL(getURLForPeerConnId(peerConnId)).
		WithMethod(http.GET).
		WithQueryParamFilter("role", string(role)).
		WithResult(result).
		Do()

	return result, err
}

// UpdatePeerConn - update the specific peer connection
//
// PARAMS:
//   - peerConnId: the id of the specific peer connection
//   - args: the arguments to update the specific peer connection
//
// RETURNS:
//   - error: nil if success otherwise the specific error
func (c *Client) UpdatePeerConn(peerConnId string, args *UpdatePeerConnArgs) error {
	if args == nil {
		return fmt.Errorf("The updatePeerConnArgs cannot be nil.")
	}

	return bce.NewRequestBuilder(c).
		WithURL(getURLForPeerConnId(peerConnId)).
		WithMethod(http.PUT).
		WithBody(args).
		Do()
}

// AcceptPeerConnApply - accept the specific peer connection
//
// PARAMS:
//   - peerConnId: the id of the specific peer connection
//   - clientToken: the idempotent token
//
// RETURNS:
//   - error: nil if success otherwise the specific error
func (c *Client) AcceptPeerConnApply(peerConnId, clientToken string) error {
	return bce.NewRequestBuilder(c).
		WithURL(getURLForPeerConnId(peerConnId)).
		WithMethod(http.PUT).
		WithQueryParam("accept", "").
		WithQueryParamFilter("clientToken", clientToken).
		Do()
}

// RejectPeerConnApply - reject the specific peer connection
//
// PARAMS:
//   - peerConnId: the id of the specific peer connection
//   - clientToken: the idempotent token
//
// RETURNS:
//   - error: nil if success otherwise the specific error
func (c *Client) RejectPeerConnApply(peerConnId, clientToken string) error {
	return bce.NewRequestBuilder(c).
		WithURL(getURLForPeerConnId(peerConnId)).
		WithMethod(http.PUT).
		WithQueryParam("reject", "").
		WithQueryParamFilter("clientToken", clientToken).
		Do()
}

// DeletePeerConn - delete the specific peer connection
//
// PARAMS:
//   - peerConnId: the id of the specific peer connection
//   - clientToken: the idempotent token
//
// RETURNS:
//   - error: nil if success otherwise the specific error
func (c *Client) DeletePeerConn(peerConnId, clientToken string) error {
	return bce.NewRequestBuilder(c).
		WithURL(getURLForPeerConnId(peerConnId)).
		WithMethod(http.DELETE).
		WithQueryParamFilter("clientToken", clientToken).
		Do()
}

// ResizePeerConn - resize the specific peer connection
//
// PARAMS:
//   - peerConnId: the id of the specific peer connection
//   - args: the arguments to resize the peer connection
//
// RETURNS:
//   - error: nil if success otherwise the specific error
func (c *Client) ResizePeerConn(peerConnId string, args *ResizePeerConnArgs) error {
	if args == nil {
		return fmt.Errorf("The resizePeerConnArgs cannot be nil.")
	}

	return bce.NewRequestBuilder(c).
		WithURL(getURLForPeerConnId(peerConnId)).
		WithMethod(http.PUT).
		WithBody(args).
		WithQueryParam("resize", "").
		WithQueryParamFilter("clientToken", args.ClientToken).
		Do()
}

// RenewPeerConn - renew the specific peer connection
//
// PARAMS:
//   - peerConnId: the id of the specific peer connection
//   - args: the arguments to renew the peer connection
//
// RETURNS:
//   - error: nil if success otherwise the specific error
func (c *Client) RenewPeerConn(peerConnId string, args *RenewPeerConnArgs) error {
	if args == nil {
		return fmt.Errorf("The renewPeerConnArgs cannot be nil.")
	}

	return bce.NewRequestBuilder(c).
		WithURL(getURLForPeerConnId(peerConnId)).
		WithMethod(http.PUT).
		WithBody(args).
		WithQueryParam("purchaseReserved", "").
		WithQueryParamFilter("clientToken", args.ClientToken).
		Do()
}

// OpenPeerConnSyncDNS - open the dns synchronization for the given peer connection
//
// PARAMS:
//   - peerConnId: the id of the specific peer connection
//   - args: the arguments to open dns synchronization
//
// RETURNS:
//   - error: nil if success otherwise the specific error
func (c *Client) OpenPeerConnSyncDNS(peerConnId string, args *PeerConnSyncDNSArgs) error {
	if args == nil {
		return fmt.Errorf("The peerConnSyncDNS cannot be nil.")
	}

	return bce.NewRequestBuilder(c).
		WithURL(getURLForPeerConnId(peerConnId)).
		WithMethod(http.PUT).
		WithQueryParam("open", "").
		WithQueryParamFilter("role", string(args.Role)).
		WithQueryParamFilter("clientToken", args.ClientToken).
		WithBody(args).
		Do()
}

// ClosePeerConnSyncDNS - close the dns synchronization for the given peer connection
//
// PARAMS:
//   - peerConnId: the id of the specific peer connection
//   - args: the arguments to close dns synchronization
//
// RETURNS:
//   - error: nil if success otherwise the specific error
func (c *Client) ClosePeerConnSyncDNS(peerConnId string, args *PeerConnSyncDNSArgs) error {
	if args == nil {
		return fmt.Errorf("The peerConnSyncDNS cannot be nil.")
	}

	return bce.NewRequestBuilder(c).
		WithURL(getURLForPeerConnId(peerConnId)).
		WithMethod(http.PUT).
		WithQueryParam("close", "").
		WithQueryParamFilter("role", string(args.Role)).
		WithQueryParamFilter("clientToken", args.ClientToken).
		Do()
}
