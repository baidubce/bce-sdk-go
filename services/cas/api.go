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

// Package api defines all APIs supported by the CAS service of BCE.

package cas

import (
	"encoding/json"
	"errors"
	"github.com/baidubce/bce-sdk-go/bce"
	"github.com/baidubce/bce-sdk-go/http"
	"io"
	"strconv"
)

// GetSslList get user ssl list
func (c *Client) GetSslList(req *GetSslListReq) (*QuerySslListResp, error) {
	if req.PageNo < 0 {
		return nil, errors.New("page no must be greater than 0")
	}
	if req.PageSize < 0 {
		return nil, errors.New("page size must be greater than 0")
	}
	if req.PageSize > 1000 {
		return nil, errors.New("page size must be less than 1000")
	}

	resp := &QuerySslListResp{}
	err := bce.NewRequestBuilder(c).
		WithURL(joinPath("/query")).
		WithQueryParam("pageNo", strconv.Itoa(req.PageNo)).
		WithQueryParam("pageSize", strconv.Itoa(req.PageSize)).
		WithQueryParamFilter("brand", req.Brand).
		WithQueryParamFilter("certType", req.CertType).
		WithQueryParamFilter("status", req.Status).
		WithMethod(http.GET).
		WithResult(resp).
		Do()
	return resp, err
}

// CheckFreeSsl check user free ssl quota
func (c *Client) CheckFreeSsl() (*CheckFreeSslResp, error) {
	resp := &CheckFreeSslResp{}
	err := bce.NewRequestBuilder(c).
		WithURL(joinPath("/query/check")).
		WithMethod(http.GET).
		WithResult(resp).
		Do()
	return resp, err
}

// QuerySslPrice query ssl price
func (c *Client) QuerySslPrice(req *QuerySslPriceReq) (*QuerySslPriceResp, error) {
	if req.Brand == "" {
		return nil, errors.New("brand must not be empty")
	}
	if req.CertType == "" {
		return nil, errors.New("certType must not be empty")
	}
	if req.ProductType == "" {
		return nil, errors.New("productType must not be empty")
	}
	if req.OrderType == "" {
		return nil, errors.New("orderType must not be empty")
	}

	resp := &QuerySslPriceResp{}
	err := bce.NewRequestBuilder(c).
		WithURL(joinPath("/price")).
		WithMethod(http.POST).
		WithBody(req).
		WithResult(resp).
		Do()
	return resp, err
}

// CreateNewOrder create new ssl order
func (c *Client) CreateNewOrder(req *CreateNewOrderReq) (*CreateNewOrderResp, error) {
	if req.Brand == "" {
		return nil, errors.New("brand must not be empty")
	}
	if req.CertType == "" {
		return nil, errors.New("certType must not be empty")
	}
	if req.ProductType == "" {
		return nil, errors.New("productType must not be empty")
	}

	resp := &CreateNewOrderResp{}
	err := bce.NewRequestBuilder(c).
		WithURL(joinPath("/order")).
		WithMethod(http.POST).
		WithQueryParam("new", "").
		WithBody(req).
		WithResult(resp).
		Do()
	return resp, err
}

// ApplyCert apply certificate
func (c *Client) ApplyCert(req *ApplyCertReq, certId string) error {
	if certId == "" {
		return errors.New("certId must not be empty")
	}

	err := bce.NewRequestBuilder(c).
		WithURL(joinPath("/certs/" + certId)).
		WithMethod(http.POST).
		WithBody(req).
		Do()
	return err
}

// DownloadLetterTemplate Provide the function of downloading confirmation letter templates for certificates
// that require submission of confirmation letters
func (c *Client) DownloadLetterTemplate(certId string) (io.ReadCloser, error) {
	if certId == "" {
		return nil, errors.New("certId must not be empty")
	}

	// Build the request
	req := &bce.BceRequest{}
	req.SetUri(joinPath("/certs/" + certId + "/letter"))
	req.SetMethod(http.GET)
	// Send request and get response
	resp := &bce.BceResponse{}
	if err := c.SendRequest(req, resp); err != nil {
		return nil, err
	}
	if resp.IsFail() {
		return nil, resp.ServiceError()
	}
	return resp.Body(), nil
}

// UploadLetter upload confirmation letter
func (c *Client) UploadLetter(body *bce.Body, certId string) error {
	if certId == "" {
		return errors.New("certId must not be empty")
	}

	// Build the request
	req := &bce.BceRequest{}
	req.SetUri(joinPath("/certs/" + certId + "/letter"))
	req.SetMethod(http.PUT)
	req.SetBody(body)
	// Send request and get response
	resp := &bce.BceResponse{}
	if err := c.SendRequest(req, resp); err != nil {
		return err
	}
	if resp.IsFail() {
		return resp.ServiceError()
	}
	return nil
}

// CancelCertApplication cancel ssl application
func (c *Client) CancelCertApplication(certId string) error {
	if certId == "" {
		return errors.New("certId must not be empty")
	}

	err := bce.NewRequestBuilder(c).
		WithURL(joinPath("/certs/"+certId)).
		WithQueryParam("cancel", "").
		WithMethod(http.DELETE).
		Do()
	return err
}

// DeleteCertApplication cancel ssl application, Used to delete failed or expired certificates and release
// free DV certificate quotas.
func (c *Client) DeleteCertApplication(certId string) error {
	if certId == "" {
		return errors.New("certId must not be empty")
	}

	err := bce.NewRequestBuilder(c).
		WithURL(joinPath("/certs/"+certId)).
		WithQueryParam("delete", "").
		WithMethod(http.DELETE).
		Do()
	return err
}

// DownloadCert download certificate,  Used to download certificates that have been successfully issued.
// Do not support downloading reissued certificates
func (c *Client) DownloadCert(body *DownloadCertReq, certId string) (io.ReadCloser, error) {
	if certId == "" {
		return nil, errors.New("certId must not be empty")
	}
	if body.Format == "" {
		return nil, errors.New("format must not be empty")
	}
	if body.FilePassword == "" {
		return nil, errors.New("filePassword must not be empty")
	}

	// Build the request
	req := &bce.BceRequest{}
	req.SetUri(joinPath("/certs/" + certId))
	req.SetMethod(http.POST)
	req.SetParam("download", "")

	if bytes, err := json.Marshal(body); err != nil {
		return nil, err
	} else if bodyI, err := bce.NewBodyFromBytes(bytes); err != nil {
		return nil, err
	} else {
		req.SetBody(bodyI)
	}
	// Send request and get response
	resp := &bce.BceResponse{}
	if err := c.SendRequest(req, resp); err != nil {
		return nil, err
	}
	if resp.IsFail() {
		return nil, resp.ServiceError()
	}
	return resp.Body(), nil
}

// GetCertDetail get certificate detail
func (c *Client) GetCertDetail(certId string) (*CertDetailResp, error) {
	if certId == "" {
		return nil, errors.New("certId must not be empty")
	}
	req := &ProductIdReq{certId}
	resp := &CertDetailResp{}
	err := bce.NewRequestBuilder(c).
		WithURL(joinPath("/certs/" + certId + "/detail")).
		WithMethod(http.POST).
		WithBody(req).
		WithResult(resp).
		Do()
	return resp, err
}

// GetCertPki get PKI
func (c *Client) GetCertPki(certId string) (*PkiResp, error) {
	if certId == "" {
		return nil, errors.New("certId must not be empty")
	}
	resp := &PkiResp{}
	err := bce.NewRequestBuilder(c).
		WithURL(joinPath("/certs/" + certId + "/pki")).
		WithMethod(http.GET).
		WithResult(resp).
		Do()
	return resp, err
}

// GetCertContact get certificate contact
func (c *Client) GetCertContact(certId string) (*CertContactResp, error) {
	if certId == "" {
		return nil, errors.New("certId must not be empty")
	}
	resp := &CertContactResp{}
	err := bce.NewRequestBuilder(c).
		WithURL(joinPath("/certs/" + certId + "/contact")).
		WithMethod(http.GET).
		WithResult(resp).
		Do()
	return resp, err
}

// ChangeCertUser change certificate user
func (c *Client) ChangeCertUser(req *ChangeCertUserReq, newUserId string) error {
	if newUserId == "" {
		return errors.New("newUserId must not be empty")
	}
	err := bce.NewRequestBuilder(c).
		WithURL(joinPath("/trans/batch/" + newUserId)).
		WithMethod(http.POST).
		WithBody(req).
		Do()
	return err
}

func joinPath(suffix string) string {
	return UriPrefix + suffix
}
