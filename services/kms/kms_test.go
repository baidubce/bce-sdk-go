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
	 "encoding/base64"
	 "fmt"
	 "testing"
 
	 "github.com/stretchr/testify/assert"
 )
 
 var kmsClientFake = &Client{}
 
 // Test data
 var myKeyID = "myfakekeyid"
 var myPlanDelKeyID = "0bc173c1-f443-1f7b-5356-f4847de582a9"
 
 func TestMain(m *testing.M) {
	 ak := "b0c95858b721483f81deddab291cd904"
	 sk := "0d6455a067f948969d2bef585743aa4d"
	 cli, err := NewClient(ak, sk, "bj")
	 if err != nil {
		 fmt.Println(err)
		 return
	 }
 
	 kmsClientFake = cli
	 m.Run()
 }
 
 func TestClient_CancelKeyDeletion_WithInvalidReq(t *testing.T) {
	 err := kmsClientFake.CancelKeyDeletion(&CancelKeyDeletionReq{})
	 assert.NotNil(t, err)
 }
 
 func TestClient_CancelKeyDeletion_WithValidReq(t *testing.T) {
	 err := kmsClientFake.CancelKeyDeletion(&CancelKeyDeletionReq{KeyID: myPlanDelKeyID})
	 assert.Nil(t, err)
 }
 
 func TestClient_CreateKey_WithInvalidReq(t *testing.T) {
	 _, err := kmsClientFake.CreateKey(&CreateKeyReq{})
	 assert.NotNil(t, err)
 }
 
 func TestClient_CreateKey_WithValidReq(t *testing.T) {
	 _, err := kmsClientFake.CreateKey(&CreateKeyReq{Description: "bce test", KeyUsage: "bce test"})
	 assert.Nil(t, err)
 }
 
 func TestClient_Decrypt_WithInvalidReq(t *testing.T) {
	 _, err := kmsClientFake.Decrypt(&DecryptReq{})
	 assert.NotNil(t, err)
 }
 
 func TestClient_Decrypt_WithValidReq(t *testing.T) {
	 str := base64.StdEncoding.EncodeToString([]byte("1"))
 
	 enRes, err := kmsClientFake.Encrypt(&EncryptReq{KeyID: myKeyID, Plaintext: str})
	 assert.Nil(t, err)
 
	 deRes, err := kmsClientFake.Decrypt(&DecryptReq{Ciphertext: enRes.Ciphertext})
	 assert.Nil(t, err)
 
	 assert.EqualValues(t, str, deRes.Plaintext)
 }
 
 func TestClient_DescribeKey_WithInvalidReq(t *testing.T) {
	 _, err := kmsClientFake.DescribeKey(&DescribeKeyReq{})
	 assert.NotNil(t, err)
 }
 
 func TestClient_DescribeKey_WithValidReq(t *testing.T) {
	 _, err := kmsClientFake.DescribeKey(&DescribeKeyReq{KeyID: myKeyID})
	 assert.Nil(t, err)
 }
 
 func TestClient_DisableKey_WithInvalidReq(t *testing.T) {
	 _, err := kmsClientFake.DisableKey(&DisableKeyReq{})
	 assert.NotNil(t, err)
 }
 
 func TestClient_DisableKey_WithValidReq(t *testing.T) {
	 _, err := kmsClientFake.DisableKey(&DisableKeyReq{KeyID: myPlanDelKeyID})
	 assert.Nil(t, err)
 }
 
 func TestClient_EnableKey_WithInvalidReq(t *testing.T) {
	 _, err := kmsClientFake.EnableKey(&EnableKeyReq{})
	 assert.NotNil(t, err)
 }
 
 func TestClient_EnableKey_WithValidReq(t *testing.T) {
	 _, err := kmsClientFake.EnableKey(&EnableKeyReq{KeyID: myPlanDelKeyID})
	 assert.Nil(t, err)
 }
 
 func TestClient_Encrypt_WithInvalidReq(t *testing.T) {
	 _, err := kmsClientFake.Encrypt(&EncryptReq{})
	 assert.NotNil(t, err)
 }
 
 func TestClient_Encrypt_WithValidReq(t *testing.T) {
	 str := base64.StdEncoding.EncodeToString([]byte("1"))
 
	 _, err := kmsClientFake.Encrypt(&EncryptReq{KeyID: myKeyID, Plaintext: str})
	 assert.Nil(t, err)
 }
 
 func TestClient_GenerateDataKey_WithInvalidReq(t *testing.T) {
	 _, err := kmsClientFake.GenerateDataKey(&GenerateDataKeyReq{})
	 assert.NotNil(t, err)
 }
 
 func TestClient_GenerateDataKey_WithValidReq(t *testing.T) {
	 _, err := kmsClientFake.GenerateDataKey(&GenerateDataKeyReq{KeyID: myPlanDelKeyID, NumberOfBytes: 4096})
	 assert.Nil(t, err)
 }
 
 func TestClient_ListKeys_WithInvalidReq(t *testing.T) {
	 _, err := kmsClientFake.ListKeys(&ListKeysReq{})
	 assert.NotNil(t, err)
 }
 
 func TestClient_ListKeys_WithValidReq(t *testing.T) {
	 _, err := kmsClientFake.ListKeys(&ListKeysReq{Limit: 5, Marker: ""})
	 assert.Nil(t, err)
 }
 
 func TestClient_ScheduleKeyDeletion_WithInvalidReq(t *testing.T) {
	 _, err := kmsClientFake.ScheduleKeyDeletion(&ScheduleKeyDeletionReq{})
	 assert.NotNil(t, err)
 }
 
 func TestClient_ScheduleKeyDeletion_WithValidReq(t *testing.T) {
	 _, err := kmsClientFake.ScheduleKeyDeletion(&ScheduleKeyDeletionReq{KeyID: myPlanDelKeyID, PendingWindowInDays: 7})
	 assert.Nil(t, err)
 }
 
 func TestClient_GetParametersForImport_WithInvalidReq(t *testing.T) {
	 _, err := kmsClientFake.GetParametersForImport(&GetParametersForImportReq{})
	 assert.NotNil(t, err)
 }
 
 func TestClient_GetParametersForImport_WithValidReq(t *testing.T) {
	 _, err := kmsClientFake.GetParametersForImport(&GetParametersForImportReq{KeyID: myPlanDelKeyID})
	 assert.Nil(t, err)
 }
 
 func TestClient_ImportKey_WithInvalidReq(t *testing.T) {
	 _, err := kmsClientFake.ImportKey(&ImportKeyReq{})
	 assert.NotNil(t, err)
 }
 
 func TestClient_ImportKey_WithValidReq(t *testing.T) {
	 _, err := kmsClientFake.ImportKey(&ImportKeyReq{KeyID: myPlanDelKeyID})
	 assert.Nil(t, err)
 }
 
 func TestClient_ImportAsymmetricKey_WithInvalidReq(t *testing.T) {
	 _, err := kmsClientFake.ImportAsymmetricKey(&ImportAsymmetricKeyReq{})
	 assert.NotNil(t, err)
 }
 
 func TestClient_ImportAsymmetricKey_WithValidReq(t *testing.T) {
	 _, err := kmsClientFake.ImportAsymmetricKey(&ImportAsymmetricKeyReq{KeyID: myPlanDelKeyID})
	 assert.Nil(t, err)
 }
 