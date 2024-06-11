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

// cert.go - the certificate APIs definition supported by the Cert service
package cert

import (
	"fmt"

	"github.com/baidubce/bce-sdk-go/bce"
	"github.com/baidubce/bce-sdk-go/http"
)

// CreateCert - create a cert with the specific parameters
//
// PARAMS:
//   - args: the arguments to create a cert
//
// RETURNS:
//   - *CreateCertResult: the result of create Cert, contains new Cert's ID
//   - error: nil if success otherwise the specific error
func (c *Client) CreateCert(args *CreateCertArgs) (*CreateCertResult, error) {
	if args == nil {
		return nil, fmt.Errorf("unset args")
	}

	if args.CertName == "" {
		return nil, fmt.Errorf("unset CertName")
	}

	if args.CertServerData == "" {
		return nil, fmt.Errorf("unset CertServerData")
	}

	if args.CertPrivateData == "" {
		return nil, fmt.Errorf("unset CertPrivateData")
	}

	result := &CreateCertResult{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.POST).
		WithURL(getCertUri()).
		WithBody(args).
		WithResult(result).
		Do()

	return result, err
}

// UpdateCertName - update a cert's name
//
// PARAMS:
//   - id: the specific cert's ID
//   - args: the arguments to update a cert's name
//
// RETURNS:
//   - error: nil if success otherwise the specific error
func (c *Client) UpdateCertName(id string, args *UpdateCertNameArgs) error {
	if args == nil || args.CertName == "" {
		return fmt.Errorf("unset CertName")
	}

	return bce.NewRequestBuilder(c).
		WithMethod(http.PUT).
		WithURL(getCertUriWithId(id)).
		WithQueryParam("certName", "").
		WithBody(args).
		Do()
}

// ListCerts - list all certs
//
// RETURNS:
//   - *ListCertResult: the result of list all certs, contains all certs' meta
//   - error: nil if success otherwise the specific error
func (c *Client) ListCerts() (*ListCertResult, error) {
	result := &ListCertResult{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.GET).
		WithURL(getCertUri()).
		WithResult(result).
		Do()

	return result, err
}

// ListCertDetail - list all certs detail
//
// RETURNS:
//   - *ListCertDetailResult: the result of list all certs detail, contains all certs' meta detail
//   - error: nil if success otherwise the specific error
func (c *Client) ListCertDetail() (*ListCertDetailResult, error) {
	result := &ListCertDetailResult{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.GET).
		WithURL(getListCertDetailUri()).
		WithResult(result).
		Do()

	return result, err
}

// GetCertMeta - get a specific cert's meta
//
// PARAMS:
//   - id: the specific cert's ID
//
// RETURNS:
//   - *CertificateMeta: the specific cert's meta with
//   - error: nil if success otherwise the specific error
func (c *Client) GetCertMeta(id string) (*CertificateMeta, error) {
	result := &CertificateMeta{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.GET).
		WithURL(getCertUriWithId(id)).
		WithResult(result).
		Do()

	return result, err
}

// GetCertDetail - get a specific cert's meta detail
//
// PARAMS:
//   - id: the specific cert's ID
//
// RETURNS:
//   - *CertificateDetailMeta: the specific cert's meta detail
//   - error: nil if success otherwise the specific error
func (c *Client) GetCertDetail(id string) (*CertificateDetailMeta, error) {
	result := &CertificateDetailMeta{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.GET).
		WithURL(getCertDetailUriWithId(id)).
		WithResult(result).
		Do()

	return result, err
}

// DeleteCert - delete a specific cert
//
// PARAMS:
//   - id: the specific cert's ID
//
// RETURNS:
//   - error: nil if success otherwise the specific error
func (c *Client) DeleteCert(id string) error {
	return bce.NewRequestBuilder(c).
		WithMethod(http.DELETE).
		WithURL(getCertUriWithId(id)).
		Do()
}

// UpdateCertData - update a specific cert's data, include update key
//
// PARAMS:
//   - id: the specific cert's ID
//   - args: the arguments to update a specific cert
//
// RETURNS:
//   - error: nil if success otherwise the specific error
func (c *Client) UpdateCertData(id string, args *UpdateCertDataArgs) error {
	if args == nil {
		return fmt.Errorf("unset args")
	}

	if args.CertName == "" {
		return fmt.Errorf("unset CertName")
	}

	if args.CertServerData == "" {
		return fmt.Errorf("unset CertServerData")
	}

	if args.CertPrivateData == "" {
		return fmt.Errorf("unset CertPrivateData")
	}

	err := bce.NewRequestBuilder(c).
		WithMethod(http.PUT).
		WithURL(getCertUriWithId(id)).
		WithQueryParam("certData", "").
		WithBody(args).
		Do()

	return err
}
