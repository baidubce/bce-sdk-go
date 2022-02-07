/*
 * Copyright 2022 Baidu, Inc.
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
 package kms

 import (
	 "github.com/baidubce/bce-sdk-go/bce"
	 "github.com/baidubce/bce-sdk-go/http"
 )
 
 // CancelKeyDeletion - cancel a CMK deletion
 //
 // PARAMS:
 //	   - req: CancelKeyDeletionReq
 // RETURNS:
 //     - error: nil if success otherwise the specific error
 func (c *Client) CancelKeyDeletion(req *CancelKeyDeletionReq) error {
 
	 params := make(map[string]string, 0)
	 params[action] = cancelKeyDeletion
 
	 return bce.NewRequestBuilder(c).
		 WithMethod(http.POST).
		 WithURL("/").
		 WithQueryParams(params).
		 WithHeader(http.CONTENT_TYPE, bce.DEFAULT_CONTENT_TYPE).
		 WithBody(req).
		 Do()
 }
 
 // CreateKey - create a CMK
 //
 // PARAMS:
 //	   - req: CreateKeyReq
 // RETURNS:
 //     - *CreateKeyRes:  create result if success otherwise nil
 //     - error: nil if success otherwise the specific error
 func (c *Client) CreateKey(req *CreateKeyReq) (*CreateKeyRes, error) {
	 // set default value
	 req.KeyUsage = defaultKeyUsage
 
	 if req.KeySpec == "" {
		 req.KeySpec = defaultKeySpec
	 }
 
	 if req.Origin == "" {
		 req.Origin = defaultOrigin
	 }
 
	 if req.ProtectedBy == "" {
		 req.ProtectedBy = defaultProtectedBy
	 }
 
	 params := make(map[string]string, 0)
	 params[action] = createKey
 
	 result := &CreateKeyRes{}
 
	 err := bce.NewRequestBuilder(c).
		 WithMethod(http.POST).
		 WithURL("/").
		 WithQueryParams(params).
		 WithHeader(http.CONTENT_TYPE, bce.DEFAULT_CONTENT_TYPE).
		 WithBody(req).
		 WithResult(result).
		 Do()
	 return result, err
 }
 
 // Decrypt - decrypt DecryptReq
 //
 // PARAMS:
 //	   - req: DecryptReq
 // RETURNS:
 //     - *DecryptRes: decrypt result if success otherwise nil
 //     - error: nil if success otherwise the specific error
 func (c *Client) Decrypt(req *DecryptReq) (*DecryptRes, error) {
	 params := make(map[string]string, 0)
	 params[action] = decrypt
 
	 result := &DecryptRes{}
 
	 err := bce.NewRequestBuilder(c).
		 WithMethod(http.POST).
		 WithURL("/").
		 WithQueryParams(params).
		 WithHeader(http.CONTENT_TYPE, bce.DEFAULT_CONTENT_TYPE).
		 WithBody(req).
		 WithResult(result).
		 Do()
	 return result, err
 }
 
 // DescribeKey - describe a CMK
 //
 // PARAMS:
 //	   - req: DescribeKeyReq
 // RETURNS:
 //     - *DescribeKeyRes: describe result if success otherwise nil
 //     - error: nil if success otherwise the specific error
 func (c *Client) DescribeKey(req *DescribeKeyReq) (*DescribeKeyRes, error) {
	 params := make(map[string]string, 0)
	 params[action] = describeKey
 
	 result := &DescribeKeyRes{}
 
	 err := bce.NewRequestBuilder(c).
		 WithMethod(http.POST).
		 WithURL("/").
		 WithQueryParams(params).
		 WithHeader(http.CONTENT_TYPE, bce.DEFAULT_CONTENT_TYPE).
		 WithBody(req).
		 WithResult(result).
		 Do()
	 return result, err
 }
 
 // DisableKey - disable a CMK
 //
 // PARAMS:
 //	   - req: DisableKeyReq
 // RETURNS:
 //     - *DisableKeyRes: disable a CMK result if success otherwise nil
 //     - error: nil if success otherwise the specific error
 func (c *Client) DisableKey(req *DisableKeyReq) (*DisableKeyRes, error) {
	 params := make(map[string]string, 0)
	 params[action] = disableKey
 
	 result := &DisableKeyRes{}
 
	 err := bce.NewRequestBuilder(c).
		 WithMethod(http.POST).
		 WithURL("/").
		 WithQueryParams(params).
		 WithHeader(http.CONTENT_TYPE, bce.DEFAULT_CONTENT_TYPE).
		 WithBody(req).
		 WithResult(result).
		 Do()
	 return result, err
 }
 
 // EnableKey - enable a CMK
 //
 // PARAMS:
 //	   - req: EnableKeyReq
 // RETURNS:
 //     - *EnableKeyRes: enable a CMK result if success otherwise nil
 //     - error: nil if success otherwise the specific error
 func (c *Client) EnableKey(req *EnableKeyReq) (*EnableKeyRes, error) {
	 params := make(map[string]string, 0)
	 params[action] = enableKey
 
	 result := &EnableKeyRes{}
 
	 err := bce.NewRequestBuilder(c).
		 WithMethod(http.POST).
		 WithURL("/").
		 WithQueryParams(params).
		 WithHeader(http.CONTENT_TYPE, bce.DEFAULT_CONTENT_TYPE).
		 WithBody(req).
		 WithResult(result).
		 Do()
	 return result, err
 }
 
 // Encrypt - encrypt a EncryptReq
 //
 // PARAMS:
 //	   - req: EncryptReq
 // RETURNS:
 //     - *EncryptRes: encrypt result if success otherwise nil
 //     - error: nil if success otherwise the specific error
 func (c *Client) Encrypt(req *EncryptReq) (*EncryptRes, error) {
	 params := make(map[string]string, 0)
	 params[action] = encrypt
	 result := &EncryptRes{}
	 err := bce.NewRequestBuilder(c).
		 WithMethod(http.POST).
		 WithURL("/").
		 WithQueryParams(params).
		 WithHeader(http.CONTENT_TYPE, bce.DEFAULT_CONTENT_TYPE).
		 WithBody(req).
		 WithResult(result).
		 Do()
	 return result, err
 }
 
 // GenerateDataKey - generate data key with specific CMK
 //
 // PARAMS:
 //	   - req: GenerateDataKeyReq
 // RETURNS:
 //     - *GenerateDataKeyRes: enable a CMK result if success otherwise nil
 //     - error: nil if success otherwise the specific error
 func (c *Client) GenerateDataKey(req *GenerateDataKeyReq) (*GenerateDataKeyRes, error) {
	 params := make(map[string]string, 0)
	 params[action] = generateDataKey
 
	 result := &GenerateDataKeyRes{}
 
	 err := bce.NewRequestBuilder(c).
		 WithMethod(http.POST).
		 WithURL("/").
		 WithQueryParams(params).
		 WithHeader(http.CONTENT_TYPE, bce.DEFAULT_CONTENT_TYPE).
		 WithBody(req).
		 WithResult(result).
		 Do()
	 return result, err
 }
 
 // ListKeys - list the CMK list in specific region
 //
 // PARAMS:
 //	   - req: ListKeysReq
 // RETURNS:
 //     - *ListKeysRes: list result if success otherwise nil
 //     - error: nil if success otherwise the specific error
 func (c *Client) ListKeys(req *ListKeysReq) (*ListKeysRes, error) {
	 params := make(map[string]string, 0)
	 params[action] = listKeys
 
	 result := &ListKeysRes{}
 
	 err := bce.NewRequestBuilder(c).
		 WithMethod(http.POST).
		 WithURL("/").
		 WithQueryParams(params).
		 WithHeader(http.CONTENT_TYPE, bce.DEFAULT_CONTENT_TYPE).
		 WithBody(req).
		 WithResult(result).
		 Do()
	 return result, err
 }
 
 // ScheduleKeyDeletion -
 //
 // PARAMS:
 //	   - req: ScheduleKeyDeletionReq
 // RETURNS:
 //     - *ScheduleKeyDeletionRes: result if success otherwise nil
 //     - error: nil if success otherwise the specific error
 func (c *Client) ScheduleKeyDeletion(req *ScheduleKeyDeletionReq) (*ScheduleKeyDeletionRes, error) {
	 if req.PendingWindowInDays < 7 || req.PendingWindowInDays > 30 {
		 req.PendingWindowInDays = defaultPendingWindowInDays
	 }
 
	 params := make(map[string]string, 0)
	 params[action] = scheduleKeyDeletion
 
	 result := &ScheduleKeyDeletionRes{}
 
	 err := bce.NewRequestBuilder(c).
		 WithMethod(http.POST).
		 WithURL("/").
		 WithQueryParams(params).
		 WithHeader(http.CONTENT_TYPE, bce.DEFAULT_CONTENT_TYPE).
		 WithBody(req).
		 WithResult(result).
		 Do()
	 return result, err
 }
 
 // GetParametersForImport -
 //
 // PARAMS:
 //	   - req: GetParametersForImportReq
 // RETURNS:
 //     - *GetParametersForImportRes: result if success otherwise nil
 //     - error: nil if success otherwise the specific error
 func (c *Client) GetParametersForImport(req *GetParametersForImportReq) (*GetParametersForImportRes, error) {
	 if req.WrappingAlgorithm == "" {
		 req.WrappingAlgorithm = defaultWrappingAlgorithm
	 }
 
	 if req.WrappingKeySpec == "" {
		 req.WrappingKeySpec = defaultWrappingKeySpec
	 }
 
	 if req.PublicKeyEncoding == "" {
		 req.PublicKeyEncoding = defaultPublicKeyEncoding
	 }
 
	 params := make(map[string]string, 0)
	 params[action] = getParametersForImport
 
	 result := &GetParametersForImportRes{}
 
	 err := bce.NewRequestBuilder(c).
		 WithMethod(http.POST).
		 WithURL("/").
		 WithQueryParams(params).
		 WithHeader(http.CONTENT_TYPE, bce.DEFAULT_CONTENT_TYPE).
		 WithBody(req).
		 WithResult(result).
		 Do()
	 return result, err
 }
 
 // ImportKey -
 //
 // PARAMS:
 //	   - req: ImportKeyReq
 // RETURNS:
 //     - *ImportKeyRes: result if success otherwise nil
 //     - error: nil if success otherwise the specific error
 func (c *Client) ImportKey(req *ImportKeyReq) (*ImportKeyRes, error) {
	 if req.KeyUsage == "" {
		 req.KeyUsage = defaultKeyUsage
	 }
 
	 params := make(map[string]string, 0)
	 params[action] = importKey
 
	 result := &ImportKeyRes{}
 
	 err := bce.NewRequestBuilder(c).
		 WithMethod(http.POST).
		 WithURL("/").
		 WithQueryParams(params).
		 WithHeader(http.CONTENT_TYPE, bce.DEFAULT_CONTENT_TYPE).
		 WithBody(req).
		 WithResult(result).
		 Do()
	 return result, err
 }
 
 // ImportAsymmetricKey -
 //
 // PARAMS:
 //	   - req: ImportAsymmetricKeyReq
 // RETURNS:
 //     - *ImportAsymmetricKeyRes: result if success otherwise nil
 //     - error: nil if success otherwise the specific error
 func (c *Client) ImportAsymmetricKey(req *ImportAsymmetricKeyReq) (*ImportAsymmetricKeyRes, error) {
 
	 params := make(map[string]string, 0)
	 params[action] = importAsymmetricKey
 
	 result := &ImportAsymmetricKeyRes{}
 
	 err := bce.NewRequestBuilder(c).
		 WithMethod(http.POST).
		 WithURL("/").
		 WithQueryParams(params).
		 WithHeader(http.CONTENT_TYPE, bce.DEFAULT_CONTENT_TYPE).
		 WithBody(req).
		 WithResult(result).
		 Do()
	 return result, err
 }
 