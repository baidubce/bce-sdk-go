/*
 * Copyright 2021 Baidu, Inc.
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

package api

import (
	"github.com/baidubce/bce-sdk-go/bce"
	"github.com/baidubce/bce-sdk-go/http"
)

func CreatePolicy(cli bce.Client, body *bce.Body) (*CreatePolicyResult, error) {
	req := &bce.BceRequest{}
	req.SetUri(URI_PREFIX + URI_POLICY)
	req.SetMethod(http.POST)
	req.SetBody(body)
	req.SetHeader(http.CONTENT_TYPE, bce.DEFAULT_CONTENT_TYPE)

	resp := &bce.BceResponse{}
	if err := cli.SendRequest(req, resp); err != nil {
		return nil, err
	}
	if resp.IsFail() {
		return nil, resp.ServiceError()
	}
	jsonBody := &CreatePolicyResult{}
	if err := resp.ParseJsonBody(jsonBody); err != nil {
		return nil, err
	}
	defer func() { resp.Body().Close() }()
	return jsonBody, nil
}

func GetPolicy(cli bce.Client, name, policyType string) (*GetPolicyResult, error) {
	req := &bce.BceRequest{}
	req.SetUri(getPolicyUri(name))
	if policyType != "" {
		req.SetParam("policyType", policyType)
	}
	req.SetMethod(http.GET)

	resp := &bce.BceResponse{}
	if err := cli.SendRequest(req, resp); err != nil {
		return nil, err
	}
	if resp.IsFail() {
		return nil, resp.ServiceError()
	}
	jsonBody := &GetPolicyResult{}
	if err := resp.ParseJsonBody(jsonBody); err != nil {
		return nil, err
	}
	defer func() { resp.Body().Close() }()
	return jsonBody, nil
}

func DeletePolicy(cli bce.Client, name string) error {
	req := &bce.BceRequest{}
	req.SetUri(getPolicyUri(name))
	req.SetMethod(http.DELETE)

	resp := &bce.BceResponse{}
	if err := cli.SendRequest(req, resp); err != nil {
		return err
	}
	if resp.IsFail() {
		return resp.ServiceError()
	}
	defer func() { resp.Body().Close() }()
	return nil
}

func ListPolicy(cli bce.Client, nameFilter, policyType string) (*ListPolicyResult, error) {
	req := &bce.BceRequest{}
	req.SetUri(URI_PREFIX + URI_POLICY)
	if nameFilter != "" {
		req.SetParam("nameFilter", nameFilter)
	}
	if policyType != "" {
		req.SetParam("policyType", policyType)
	}
	req.SetMethod(http.GET)

	resp := &bce.BceResponse{}
	if err := cli.SendRequest(req, resp); err != nil {
		return nil, err
	}
	if resp.IsFail() {
		return nil, resp.ServiceError()
	}
	jsonBody := &ListPolicyResult{}
	if err := resp.ParseJsonBody(jsonBody); err != nil {
		return nil, err
	}
	defer func() { resp.Body().Close() }()
	return jsonBody, nil
}

func AttachPolicyToUser(cli bce.Client, args *AttachPolicyToUserArgs) error {
	req := &bce.BceRequest{}
	req.SetUri(getUserUri(args.UserName) + URI_POLICY + "/" + args.PolicyName)
	if args.PolicyType != "" {
		req.SetParam("policyType", args.PolicyType)
	}
	req.SetMethod(http.PUT)

	resp := &bce.BceResponse{}
	if err := cli.SendRequest(req, resp); err != nil {
		return err
	}
	if resp.IsFail() {
		return resp.ServiceError()
	}
	defer func() { resp.Body().Close() }()
	return nil
}

func DetachPolicyFromUser(cli bce.Client, args *DetachPolicyFromUserArgs) error {
	req := &bce.BceRequest{}
	req.SetUri(getUserUri(args.UserName) + URI_POLICY + "/" + args.PolicyName)
	if args.PolicyType != "" {
		req.SetParam("policyType", args.PolicyType)
	}
	req.SetMethod(http.DELETE)

	resp := &bce.BceResponse{}
	if err := cli.SendRequest(req, resp); err != nil {
		return err
	}
	if resp.IsFail() {
		return resp.ServiceError()
	}
	defer func() { resp.Body().Close() }()
	return nil
}

func ListUserAttachedPolicies(cli bce.Client, name string) (*ListPolicyResult, error) {
	req := &bce.BceRequest{}
	req.SetUri(getUserUri(name) + URI_POLICY)
	req.SetMethod(http.GET)

	resp := &bce.BceResponse{}
	if err := cli.SendRequest(req, resp); err != nil {
		return nil, err
	}
	if resp.IsFail() {
		return nil, resp.ServiceError()
	}
	jsonBody := &ListPolicyResult{}
	if err := resp.ParseJsonBody(jsonBody); err != nil {
		return nil, err
	}
	defer func() { resp.Body().Close() }()
	return jsonBody, nil
}

func AttachPolicyToGroup(cli bce.Client, args *AttachPolicyToGroupArgs) error {
	req := &bce.BceRequest{}
	req.SetUri(getGroupUri(args.GroupName) + URI_POLICY + "/" + args.PolicyName)
	if args.PolicyType != "" {
		req.SetParam("policyType", args.PolicyType)
	}
	req.SetMethod(http.PUT)

	resp := &bce.BceResponse{}
	if err := cli.SendRequest(req, resp); err != nil {
		return err
	}
	if resp.IsFail() {
		return resp.ServiceError()
	}
	defer func() { resp.Body().Close() }()
	return nil
}

func DetachPolicyFromGroup(cli bce.Client, args *DetachPolicyFromGroupArgs) error {
	req := &bce.BceRequest{}
	req.SetUri(getGroupUri(args.GroupName) + URI_POLICY + "/" + args.PolicyName)
	if args.PolicyType != "" {
		req.SetParam("policyType", args.PolicyType)
	}
	req.SetMethod(http.DELETE)

	resp := &bce.BceResponse{}
	if err := cli.SendRequest(req, resp); err != nil {
		return err
	}
	if resp.IsFail() {
		return resp.ServiceError()
	}
	defer func() { resp.Body().Close() }()
	return nil
}

func ListGroupAttachedPolicies(cli bce.Client, name string) (*ListPolicyResult, error) {
	req := &bce.BceRequest{}
	req.SetUri(getGroupUri(name) + URI_POLICY)
	req.SetMethod(http.GET)

	resp := &bce.BceResponse{}
	if err := cli.SendRequest(req, resp); err != nil {
		return nil, err
	}
	if resp.IsFail() {
		return nil, resp.ServiceError()
	}
	jsonBody := &ListPolicyResult{}
	if err := resp.ParseJsonBody(jsonBody); err != nil {
		return nil, err
	}
	defer func() { resp.Body().Close() }()
	return jsonBody, nil
}

func AttachPolicyToRole(cli bce.Client, args *AttachPolicyToRoleArgs) error {
	req := &bce.BceRequest{}
	req.SetUri(getRoleUri(args.RoleName) + URI_POLICY + "/" + args.PolicyName)
	req.SetMethod(http.PUT)
	if args.PolicyType != "" {
		req.SetParam("policyType", args.PolicyType)
	}

	resp := &bce.BceResponse{}
	if err := cli.SendRequest(req, resp); err != nil {
		return err
	}
	if resp.IsFail() {
		return resp.ServiceError()
	}
	defer func() { resp.Body().Close() }()
	return nil
}

func DetachPolicyFromRole(cli bce.Client, args *DetachPolicyToRoleArgs) error {
	req := &bce.BceRequest{}
	req.SetUri(getRoleUri(args.RoleName) + URI_POLICY + "/" + args.PolicyName)
	req.SetMethod(http.DELETE)
	if args.PolicyType != "" {
		req.SetParam("policyType", args.PolicyType)
	}

	resp := &bce.BceResponse{}
	if err := cli.SendRequest(req, resp); err != nil {
		return err
	}
	if resp.IsFail() {
		return resp.ServiceError()
	}
	defer func() { resp.Body().Close() }()
	return nil
}

func ListRoleAttachedPolicies(cli bce.Client, roleName string) (*ListPolicyResult, error) {
	req := &bce.BceRequest{}
	req.SetUri(getRoleUri(roleName) + URI_POLICY)
	req.SetMethod(http.GET)

	resp := &bce.BceResponse{}
	if err := cli.SendRequest(req, resp); err != nil {
		return nil, err
	}
	if resp.IsFail() {
		return nil, resp.ServiceError()
	}
	jsonBody := &ListPolicyResult{}
	if err := resp.ParseJsonBody(jsonBody); err != nil {
		return nil, err
	}
	defer func() { resp.Body().Close() }()
	return jsonBody, nil
}

func ListPolicyAttachedEntities(cli bce.Client, policyId string) (*ListPolicyAttachedEntityResult, error) {
	req := &bce.BceRequest{}
	req.SetUri(getPolicyUri(policyId) + URI_ENTITY)
	req.SetMethod(http.GET)

	resp := &bce.BceResponse{}
	if err := cli.SendRequest(req, resp); err != nil {
		return nil, err
	}
	if resp.IsFail() {
		return nil, resp.ServiceError()
	}
	jsonBody := &ListPolicyAttachedEntityResult{}
	if err := resp.ParseJsonBody(jsonBody); err != nil {
		return nil, err
	}
	defer func() { resp.Body().Close() }()
	return jsonBody, nil
}

func getPolicyUri(name string) string {
	return URI_PREFIX + URI_POLICY + "/" + name
}
