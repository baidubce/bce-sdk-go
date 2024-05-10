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

// model.go - definitions of the request arguments and results data structure model

package cas

var (
	certStatuses = map[string]string{
		"APPLY_PENDING":     "待申请",
		"APPLYING":          "申请中",
		"NORMAL":            "申请成功",
		"FAILED":            "申请失败",
		"RENEW_PENDING":     "待续费",
		"REPLACING":         "重新颁发",
		"REPLACING_FAILED":  "重新颁发失败",
		"REPLACING_NORMAL":  "重新颁发成功",
		"REPLACING_CONFIRM": "重新颁发待确认",
		"REPLACE_CA":        "重颁发信息确认后，需要提交信息至CA",
		"EXPIRED":           "已过期",
		"DELETED":           "已删除",
	}

	certTypes = map[string]string{
		"DV":    "DV(域名型)",
		"OV":    "OV(企业型)",
		"EV":    "EV(增强型)",
		"OVPRO": "企业型专业版（OV PRO）证书",
		"EVPRO": "增强型专业版（EV PRO)证书",
	}

	certBrands = map[string]string{
		"SECURESITE": "SECURESITE",
		"GEOTRUST":   "GEOTRUST",
		"GLOBALSIGN": "GLOBALSIGN",
		"CFCA":       "CFCA",
		"TRUSTASIA":  "TRUSTASIA",
		"BAIDUTRUST": "BAIDUTRUST",
	}

	productTypes = map[string]string{
		"SINGLE":       "单域名",
		"MULTI":        "多域名版",
		"WILDCARD":     "通配符域名版",
		"SINGLE_PRO":   "单域名专业版",
		"MULTI_PRO":    "多域名专业版",
		"WILDCARD_PRO": "通配符域名专业版",
	}

	orderTypes = map[string]string{
		"NEW":   "购买",
		"RENEW": "续费",
	}

	encryptionTypes = map[string]string{
		"RSA":   "RSA",
		"ECDSA": "ECDSA",
	}

	encryptionStrengths = map[string]string{
		"ECDSA_PRIME256V1": "ECDSA",
		"RSA_2048":         "RSA",
		"RSA_4096":         "RSA",
	}

	authTypes = map[string]string{
		"DNS":   "DNS",
		"FILE":  "FILE",
		"HTTP":  "HTTP",
		"CNAME": "CNAME",
	}

	CertFileTypes = map[string]string{
		"PEM":         "PEM",
		"PEM_APACHE":  "PEM_APACHE",
		"PEM_NGINX":   "PEM_NGINX",
		"PEM_HAPROXY": "PEM_HAPROXY",
		"JKS_TOMCAT":  "JKS_TOMCAT",
		"JKS":         "JKS",
		"PKCS12":      "PKCS12",
	}
)

type GetSslListReq struct {
	PageNo   int    `json:"pageNo,omitempty"`
	PageSize int    `json:"pageSize,omitempty"`
	Brand    string `json:"brand,omitempty"`
	CertType string `json:"certType,omitempty"`
	Status   string `json:"status,omitempty"`
}

type QuerySslListResp struct {
	TotalCount int            `json:"totalCount,omitempty"`
	Result     []QuerySslList `json:"result,omitempty"`
}

type QuerySslList struct {
	CertType       string `json:"certType,omitempty"`
	ProductType    string `json:"productType,omitempty"`
	ExpireTime     string `json:"expireTime,omitempty"`
	CreateTime     string `json:"createTime,omitempty"`
	Duration       int    `json:"duration,omitempty"`
	ProductId      string `json:"productId,omitempty"`
	Brand          string `json:"brand,omitempty"`
	DomainNumber   int    `json:"domainNumber,omitempty"`
	WildcardNumber int    `json:"wildcardNumber,omitempty"`
	Status         string `json:"status,omitempty"`
	DomainName     string `json:"domainName,omitempty"`
}

type CheckFreeSslResp struct {
	FreeCount        int  `json:"freeCount,omitempty"`
	EnablePurchaseDV bool `json:"enablePurchaseDV,omitempty"`
}

type QuerySslPriceReq struct {
	CertType       string `json:"certType,omitempty"`
	ProductType    string `json:"productType,omitempty"`
	OrderType      string `json:"orderType,omitempty"`
	Brand          string `json:"brand,omitempty"`
	DomainNumber   int    `json:"domainNumber,omitempty"`
	WildcardNumber int    `json:"wildcardNumber,omitempty"`
	PurchaseLength int    `json:"purchaseLength,omitempty"`
}

type QuerySslPriceResp struct {
	Price string `json:"price"`
}

type CreateNewOrderReq struct {
	CertType       string `json:"certType,omitempty"`
	ProductType    string `json:"productType,omitempty"`
	Brand          string `json:"brand,omitempty"`
	DomainNumber   int    `json:"domainNumber,omitempty"`
	WildcardNumber int    `json:"wildcardNumber,omitempty"`
	PurchaseLength int    `json:"purchaseLength,omitempty"`
}

type CreateNewOrderResp struct {
	BceOrderId string   `json:"bceOrderId,omitempty"`
	CertIds    []string `json:"certIds,omitempty"`
}

type ApplyCertReq struct {
	Company         string   `json:"company,omitempty"`
	Address         string   `json:"address,omitempty"`
	PostalCode      string   `json:"postalCode,omitempty"`
	Region          Region   `json:"region,omitempty"`
	Password        string   `json:"password,omitempty"`
	Algorithm       string   `json:"algorithm,omitempty"`
	Strength        string   `json:"strength,omitempty"`
	Crs             string   `json:"csr,omitempty"`
	Domain          string   `json:"domain,omitempty"`
	VerifyMode      string   `json:"verifyMode,omitempty"`
	MultiDomain     []string `json:"multiDomain,omitempty"`
	Department      string   `json:"department,omitempty"`
	CompanyPhone    string   `json:"companyPhone,omitempty"`
	OrderGivenName  string   `json:"orderGivenName,omitempty"`
	OrderFamilyName string   `json:"orderFamilyName,omitempty"`
	OrderPosition   string   `json:"orderPosition,omitempty"`
	OrderEmail      string   `json:"orderEmail,omitempty"`
	OrderPhone      string   `json:"orderPhone,omitempty"`
	TechGivenName   string   `json:"techGivenName,omitempty"`
	TechFamilyName  string   `json:"techFamilyName,omitempty"`
	TechPosition    string   `json:"techPosition,omitempty"`
	TechEmail       string   `json:"techEmail,omitempty"`
	TechPhone       string   `json:"techPhone,omitempty"`
}

type Region struct {
	Province string `json:"province,omitempty"`
	City     string `json:"city,omitempty"`
	Country  string `json:"country,omitempty"`
}

type DownloadCertReq struct {
	Format        string `json:"format,omitempty"`
	OrderPassword string `json:"orderPassword,omitempty"`
	FilePassword  string `json:"filePassword,omitempty"`
}

type CertDetailResp struct {
	ProductName       string   `json:"productName,omitempty"`
	CertType          string   `json:"certType,omitempty"`
	ProductType       string   `json:"productType,omitempty"`
	ApplyTime         string   `json:"applyTime,omitempty"`
	DownloadSupported bool     `json:"downloadSupported,omitempty"`
	BindDomains       []string `json:"bindDomains,omitempty"`
	Duration          int      `json:"duration,omitempty"`
	FromBaidu         bool     `json:"fromBaidu,omitempty"`
	ProductId         string   `json:"productId,omitempty"`
	Brand             string   `json:"brand,omitempty"`
	DomainNumber      int      `json:"domainNumber,omitempty"`
	WildcardNumber    int      `json:"wildcardNumber,omitempty"`
	ProcessStatus     string   `json:"processStatus,omitempty"`
	DomainName        string   `json:"domainName,omitempty"`
}

type ProductIdReq struct {
	ProductId string `json:"productId,omitempty"`
}

type PkiResp struct {
	Company    string `json:"company,omitempty"`
	Address    string `json:"address,omitempty"`
	Region     Region `json:"region,omitempty"`
	Algorithm  string `json:"algorithm,omitempty"`
	Strength   string `json:"strength,omitempty"`
	CsrPem     string `json:"csrPem,omitempty"`
	CertPem    string `json:"certPem,omitempty"`
	CertCaPem  string `json:"certCaPem,omitempty"`
	ExpireTime string `json:"expireTime,omitempty"`
	DomainName string `json:"domainName,omitempty"`
	Department string `json:"department,omitempty"`
	StartTime  string `json:"startTime,omitempty"`
}

type CertContactResp struct {
	Company       string `json:"company,omitempty"`
	Address       string `json:"address,omitempty"`
	PostalCode    string `json:"postalCode,omitempty"`
	Region        Region `json:"region,omitempty"`
	Department    string `json:"department,omitempty"`
	CompanyPhone  string `json:"companyPhone,omitempty"`
	OrderName     string `json:"orderName,omitempty"`
	OrderPosition string `json:"orderPosition,omitempty"`
	OrderEmail    string `json:"orderEmail,omitempty"`
	OrderPhone    string `json:"orderPhone,omitempty"`
	TechName      string `json:"techName,omitempty"`
	TechPosition  string `json:"techPosition,omitempty"`
	TechEmail     string `json:"techEmail,omitempty"`
	TechPhone     string `json:"techPhone,omitempty"`
}

type ChangeCertUserReq struct {
	Params []string `json:"params,omitempty"`
}
