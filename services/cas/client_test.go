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

package cas

import (
	"encoding/json"
	"github.com/baidubce/bce-sdk-go/bce"
	"github.com/baidubce/bce-sdk-go/util/log"
	"io"
	"os"
	"path/filepath"
	"reflect"
	"runtime"
	"testing"
)

var (
	casClient *Client
	casConf   *Conf
	certId    string
)

type Conf struct {
	AK         string `json:"AK"`
	SK         string `json:"SK"`
	Endpoint   string `json:"Endpoint"`
	InstanceId string `json:"InstanceId"`
	UserId     string `json:"UserID"`
}

func init() {
	_, f, _, _ := runtime.Caller(0)
	conf := filepath.Join(filepath.Dir(f), "data/config.json")
	fp, err := os.Open(conf)
	if err != nil {
		log.Fatal("config json file of ak/sk not given:", conf)
		os.Exit(1)
	}
	decoder := json.NewDecoder(fp)
	casConf = &Conf{}
	_ = decoder.Decode(casConf)

	casClient, _ = NewClient(casConf.AK, casConf.SK, casConf.Endpoint)
	log.SetLogLevel(log.WARN)

	certId = ""
}

// ExpectEqual is the helper function for test each case
func ExpectEqual(alert func(format string, args ...interface{}),
	expected interface{}, actual interface{}) bool {
	expectedValue, actualValue := reflect.ValueOf(expected), reflect.ValueOf(actual)
	equal := false
	switch {
	case expected == nil && actual == nil:
		return true
	case expected != nil && actual == nil:
		equal = expectedValue.IsNil()
	case expected == nil && actual != nil:
		equal = actualValue.IsNil()
	default:
		if actualType := reflect.TypeOf(actual); actualType != nil {
			if expectedValue.IsValid() && expectedValue.Type().ConvertibleTo(actualType) {
				equal = reflect.DeepEqual(expectedValue.Convert(actualType).Interface(), actual)
			}
		}
	}
	if !equal {
		_, file, line, _ := runtime.Caller(1)
		alert("%s:%d: missmatch, expect %v but %v", file, line, expected, actual)
		return false
	}
	return true
}

func TestGetSslList(t *testing.T) {
	req := &GetSslListReq{
		PageSize: 10,
		PageNo:   1,
	}
	res, err := casClient.GetSslList(req)
	ExpectEqual(t.Errorf, err, nil)
	t.Logf("%+v", res)
}

func TestCheckFreeSsl(t *testing.T) {
	res, err := casClient.CheckFreeSsl()
	ExpectEqual(t.Errorf, err, nil)
	t.Logf("%+v", res)
}

func TestQuerySslPriceTest(t *testing.T) {
	req := &QuerySslPriceReq{
		OrderType:      "NEW",
		CertType:       "DV",
		ProductType:    "SINGLE",
		Brand:          "TRUSTASIA",
		DomainNumber:   1,
		WildcardNumber: 0,
		PurchaseLength: 1,
	}
	res, err := casClient.QuerySslPrice(req)
	ExpectEqual(t.Errorf, err, nil)
	t.Logf("%+v", res)
}

func TestCreateNewOrder(t *testing.T) {
	req := &CreateNewOrderReq{
		CertType:       "DV",
		ProductType:    "SINGLE",
		Brand:          "TRUSTASIA",
		DomainNumber:   1,
		WildcardNumber: 0,
		PurchaseLength: 1,
	}
	res, err := casClient.CreateNewOrder(req)
	ExpectEqual(t.Errorf, err, nil)
	t.Logf("%+v", res)
	//
}

func TestApplyCert(t *testing.T) {
	req := &ApplyCertReq{
		Company:    "com",
		Address:    "addr",
		PostalCode: "111000",
		Region: Region{
			Province: "110000",
			City:     "110100",
			Country:  "中国",
		},
		Algorithm:       "RSA",
		Strength:        "RSA_4096",
		Domain:          "domain.com",
		Password:        "abc",
		VerifyMode:      "DNS",
		MultiDomain:     []string{},
		Department:      "cas team",
		CompanyPhone:    "18500000000",
		OrderGivenName:  "hua",
		OrderFamilyName: "li",
		OrderPosition:   "director",
		OrderEmail:      "test@test.com",
		OrderPhone:      "18500000002",
		TechGivenName:   "an",
		TechFamilyName:  "li",
		TechPosition:    "tech director",
		TechEmail:       "tech@test.com",
		TechPhone:       "18500000001",
	}
	err := casClient.ApplyCert(req, certId)
	ExpectEqual(t.Errorf, err, nil)
}

func TestDownloadLetterTemplate(t *testing.T) {
	body, err := casClient.DownloadLetterTemplate(certId)
	ExpectEqual(t.Errorf, err, nil)
	defer func(body io.ReadCloser) {
		err := body.Close()
		if err != nil {
			ExpectEqual(t.Errorf, err, nil)
		}
	}(body)
	// download certificate to local dir
	// Zip decompression password is req.FilePassword
	file, err := os.Create("data/confirmation.doc")
	ExpectEqual(t.Errorf, err, nil)
	written, err := io.Copy(file, body)
	ExpectEqual(t.Errorf, err, nil)
	t.Logf("%+v", written)
}

func TestUploadLetter(t *testing.T) {
	body, _ := bce.NewBodyFromString("hello")
	body, _ = bce.NewBodyFromFile("data/confirmation.doc")
	err := casClient.UploadLetter(body, certId)
	ExpectEqual(t.Errorf, err, nil)
}

func TestCancelCertApplication(t *testing.T) {
	err := casClient.CancelCertApplication(certId)
	ExpectEqual(t.Errorf, err, nil)
}

func TestDeleteCertApplication(t *testing.T) {
	err := casClient.DeleteCertApplication(certId)
	ExpectEqual(t.Errorf, err, nil)
}

func TestDownloadCert(t *testing.T) {
	req := &DownloadCertReq{
		Format:       "PEM",
		FilePassword: "aaa",
	}
	body, err := casClient.DownloadCert(req, certId)
	ExpectEqual(t.Errorf, err, nil)

	defer func(body io.ReadCloser) {
		err := body.Close()
		if err != nil {
			ExpectEqual(t.Errorf, err, nil)
		}
	}(body)
	// download certificate to local dir
	// Zip decompression password is req.FilePassword
	file, err := os.Create("data/downloaded_cert.zip")
	ExpectEqual(t.Errorf, err, nil)
	written, err := io.Copy(file, body)
	ExpectEqual(t.Errorf, err, nil)
	t.Logf("%+v", written)
}

func TestGetCertDetail(t *testing.T) {
	detail, err := casClient.GetCertDetail(certId)
	ExpectEqual(t.Errorf, err, nil)
	t.Logf("%+v", detail)
}

func TestGetCertPki(t *testing.T) {
	pki, err := casClient.GetCertPki(certId)
	ExpectEqual(t.Errorf, err, nil)
	t.Logf("%+v", pki)
}

func TestGetCertContact(t *testing.T) {
	contact, err := casClient.GetCertContact(certId)
	ExpectEqual(t.Errorf, err, nil)
	t.Logf("%+v", contact)
}

func TestChangeCertUser(t *testing.T) {
	req := &ChangeCertUserReq{
		Params: []string{
			certId,
		},
	}
	err := casClient.ChangeCertUser(req, "newUserId")
	ExpectEqual(t.Errorf, err, nil)
}
