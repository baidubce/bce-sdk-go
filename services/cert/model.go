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

// model.go - definitions of the request arguments and results data structure model

package cert

type CreateCertArgs struct {
	CertName        string `json:"certName"`
	CertServerData  string `json:"certServerData"`
	CertPrivateData string `json:"certPrivateData"`
	CertLinkData    string `json:"certLinkData,omitempty"`
	CertType        int    `json:"certType,omitempty"`
}

type CreateCertResult struct {
	CertName string `json:"certName"`
	CertId   string `json:"certId"`
}

type UpdateCertNameArgs struct {
	CertName string `json:"certName"`
}

type CertificateMeta struct {
	CertId          string `json:"certId"`
	CertName        string `json:"certName"`
	CertCommonName  string `json:"certCommonName"`
	CertFingerprint string `json:"certFingerprint"`
	CertStartTime   string `json:"certStartTime"`
	CertStopTime    string `json:"certStopTime"`
	CertCreateTime  string `json:"certCreateTime"`
	CertUpdateTime  string `json:"certUpdateTime"`
	CertType        int    `json:"certType"`
}

type ListCertResult struct {
	Certs []CertificateMeta `json:"certs"`
}

type UpdateCertDataArgs struct {
	CertName        string `json:"certName"`
	CertServerData  string `json:"certServerData"`
	CertPrivateData string `json:"certPrivateData"`
	CertLinkData    string `json:"certLinkData,omitempty"`
	CertType        int    `json:"certType,omitempty"`
}
