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

 const (
	 action string = "action"
 
	 cancelKeyDeletion      string = "CancelKeyDeletion"
	 createKey              string = "CreateKey"
	 decrypt                string = "Decrypt"
	 describeKey            string = "DescribeKey"
	 disableKey             string = "DisableKey"
	 enableKey              string = "EnableKey"
	 encrypt                string = "Encrypt"
	 generateDataKey        string = "GenerateDataKey"
	 listKeys               string = "ListKeys"
	 scheduleKeyDeletion    string = "ScheduleKeyDeletion"
	 getParametersForImport string = "GetParametersForImport"
	 importKey              string = "ImportKey"
	 importAsymmetricKey    string = "ImportAsymmetricKey"
 
	 defaultOrigin              = "BAIDU_KMS"
	 defaultProtectedBy         = "SOFTWARE"
	 defaultKeySpec             = "BAIDU_AES_256"
	 defaultKeyUsage            = "ENCRYPT_DECRYPT"
	 defaultPendingWindowInDays = 30
	 defaultWrappingAlgorithm   = "RAW_HEX"
	 defaultWrappingKeySpec     = "RSAES_PKCS1_V1_5"
	 defaultPublicKeyEncoding   = "RSA_2048"
 )
 
 // KeyMetadata key metadata
 type KeyMetadata struct {
	 KeyID        string `json:"keyId"`
	 CreationDate string `json:"creationDate"`
	 KeyState     string `json:"keyState"`
	 Description  string `json:"description"`
	 DeletionDate string `json:"deletionDate"`
	 KeyUsage     string `json:"keyUsage"`
	 Region       string `json:"region"`
 }
 
 // KeyListEntry key list entry
 type KeyListEntry struct {
	 KeyID string `json:"keyId"`
 }
 
 // EncryptedRsaKey encrypted rsa key
 type EncryptedRsaKey struct {
	 PublicKeyDer  string `json:"publicKeyDer"`
	 EncryptedD    string `json:"encryptedD"`
	 EncryptedP    string `json:"encryptedP"`
	 EncryptedQ    string `json:"encryptedQ"`
	 EncryptedDp   string `json:"encryptedDp"`
	 EncryptedDq   string `json:"encryptedDq"`
	 EncryptedQinv string `json:"encryptedQinv"`
 }
 
 // EncryptReq encrypt request
 type EncryptReq struct {
	 KeyID     string `json:"keyId"`
	 Plaintext string `json:"plaintext"`
 }
 
 // EncryptRes encrypt response
 type EncryptRes struct {
	 KeyID      string `json:"keyId"`
	 Ciphertext string `json:"ciphertext"`
 }
 
 // DecryptReq decrypt request
 type DecryptReq struct {
	 Ciphertext string `json:"ciphertext"`
 }
 
 // DecryptRes decrypt response
 type DecryptRes struct {
	 KeyID     string `json:"keyId"`
	 Plaintext string `json:"plaintext"`
 }
 
 // CancelKeyDeletionReq cancel key deletion request
 type CancelKeyDeletionReq struct {
	 KeyID string `json:"keyId"`
 }
 
 // CreateKeyReq create key request
 type CreateKeyReq struct {
	 Description string `json:"description"`
	 // must be ENCRYPT_DECRYPT
	 KeyUsage string `json:"keyUsage"`
	 // default BAIDU_AES_256, one of BAIDU_AES_256 / AES_128 / AES_256 / RSA_1024 / RSA_2048 / RSA_4096
	 KeySpec string `json:"keySpec"`
	 // default SOFTWARE, one of SOFTWARE / HSM
	 ProtectedBy string `json:"protectedBy"`
	 // default BAIDU_KMS, one of EXTERNAL / BAIDU_KMS
	 Origin string `json:"origin"`
 }
 
 // CreateKeyRes create key response
 type CreateKeyRes struct {
	 KeyMetadata
 }
 
 // DescribeKeyReq describe key request
 type DescribeKeyReq struct {
	 KeyID string `json:"keyId"`
 }
 
 // DescribeKeyRes describe key response
 type DescribeKeyRes struct {
	 KeyMetadata
 }
 
 // DisableKeyReq disable key request
 type DisableKeyReq struct {
	 KeyID string `json:"keyId"`
 }
 
 // DisableKeyRes disable key response
 type DisableKeyRes struct {
 }
 
 // EnableKeyReq enable key request
 type EnableKeyReq struct {
	 KeyID string `json:"keyId"`
 }
 
 // EnableKeyRes enable key response
 type EnableKeyRes struct {
 }
 
 // GenerateDataKeyReq generate data key request
 type GenerateDataKeyReq struct {
	 KeyID         string `json:"keyId"`
	 KeySpec       string `json:"keySpec"`
	 NumberOfBytes int    `json:"numberOfBytes"`
 }
 
 // GenerateDataKeyRes generate data key response
 type GenerateDataKeyRes struct {
	 Ciphertext string `json:"ciphertext"`
	 KeyID      string `json:"keyId"`
	 Plaintext  string `json:"plaintext"`
 }
 
 // ListKeysReq list keys request
 type ListKeysReq struct {
	 Limit  int    `json:"limit"`
	 Marker string `json:"marker"`
 }
 
 // ListKeysRes list keys response
 type ListKeysRes struct {
	 Keys       []KeyListEntry `json:"keys"`
	 NextMarker string         `json:"nextMarker"`
	 Truncated  bool           `json:"truncated"`
 }
 
 // ScheduleKeyDeletionReq schedule key deletion request
 type ScheduleKeyDeletionReq struct {
	 KeyID string `json:"keyId"`
	 // 7 to 30, default 30
	 PendingWindowInDays int `json:"pendingWindowInDays"`
 }
 
 // ScheduleKeyDeletionRes schedule key deletion response
 type ScheduleKeyDeletionRes struct {
	 KeyID        string `json:"keyId"`
	 DeletionDate string `json:"deletionDate"`
 }
 
 // GetParametersForImportReq get parameters for import request
 type GetParametersForImportReq struct {
	 KeyID string `json:"keyId"`
	 // default RAW_HEX, one of RAW_HEX / BASE64 / PEM
	 WrappingAlgorithm string `json:"wrappingAlgorithm"`
	 // default RSAES_PKCS1_V1_5
	 WrappingKeySpec string `json:"wrappingKeySpec"`
	 // default RSA_2048
	 PublicKeyEncoding string `json:"publicKeyEncoding"`
 }
 
 // GetParametersForImportRes get parameters for import response
 type GetParametersForImportRes struct {
	 KeyID string `json:"keyId"`
	 // 其他参数文档有偏差
 }
 
 // ImportKeyReq import key request
 type ImportKeyReq struct {
	 KeyID        string `json:"keyId"`
	 ImportToken  string `json:"importToken"`
	 EncryptedKey string `json:"encryptedKey"`
	 KeySpec      string `json:"keySpec"`
	 // default ENCRYPT_DECRYPT
	 KeyUsage string `json:"keyUsage"`
 }
 
 // ImportKeyRes import key response
 type ImportKeyRes struct {
 }
 
 // ImportAsymmetricKeyReq import asymmetric key request
 type ImportAsymmetricKeyReq struct {
	 KeyID                     string          `json:"keyId"`
	 ImportToken               string          `json:"importToken"`
	 AsymmetricKeySpec         string          `json:"asymmetricKeySpec"`
	 AsymmetricKeyUsage        string          `json:"asymmetricKeyUsage"`
	 EncryptedKeyEncryptionKey string          `json:"encryptedKeyEncryptionKey"`
	 EncryptedRsaKey           EncryptedRsaKey `json:"encryptedRsaKey"`
 }
 
 // ImportAsymmetricKeyRes import asymmetric key response
 type ImportAsymmetricKeyRes struct {
 }
 