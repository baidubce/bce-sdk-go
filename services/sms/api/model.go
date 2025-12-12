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

package api

// SendSmsArgs defines the data structure for sending a SMS request
type SendSmsArgs struct {
	Mobile        string                 `json:"mobile"`
	Template      string                 `json:"template"`
	SignatureId   string                 `json:"signatureId"`
	ContentVar    map[string]interface{} `json:"contentVar"`
	Custom        string                 `json:"custom,omitempty"`
	UserExtId     string                 `json:"userExtId,omitempty"`
	CallbackUrlId string                 `json:"merchantUrlId,omitempty"`
	ClientToken   string                 `json:"clientToken,omitempty"`
}

// SendSmsResult defines the data structure of the result of sending a SMS request
type SendSmsResult struct {
	Code      string            `json:"code"`
	RequestId string            `json:"requestId"`
	Message   string            `json:"message"`
	Data      []SendMessageItem `json:"data"`
}

type SendMessageItem struct {
	Code      string `json:"code"`
	Mobile    string `json:"mobile"`
	MessageId string `json:"messageId"`
	Message   string `json:"message"`
}

// CreateSignatureArgs defines the data structure for creating a signature
type CreateSignatureArgs struct {
	Content             string `json:"content"`
	ContentType         string `json:"contentType"`
	Description         string `json:"description,omitempty"`
	CountryType         string `json:"countryType"`
	SignatureFileBase64 string `json:"signatureFileBase64,omitempty"`
	SignatureFileFormat string `json:"signatureFileFormat,omitempty"`
}

// CreateSignatureResult defines the data structure of the result of creating a signature
type CreateSignatureResult struct {
	SignatureId string `json:"signatureId"`
	Status      string `json:"status"`
}

// DeleteSignatureArgs defines the input data structure for deleting a signature
type DeleteSignatureArgs struct {
	SignatureId string `json:"signatureId"`
}

// ModifySignatureArgs defines the input data structure for modifying parameters of a signature
type ModifySignatureArgs struct {
	SignatureId         string `json:"signatureId"`
	Content             string `json:"content"`
	ContentType         string `json:"contentType"`
	Description         string `json:"description,omitempty"`
	CountryType         string `json:"countryType"`
	SignatureFileBase64 string `json:"signatureFileBase64,omitempty"`
	SignatureFileFormat string `json:"signatureFileFormat,omitempty"`
}

// GetSignatureArgs defines the input data structure for Getting a signature
type GetSignatureArgs struct {
	SignatureId string `json:"signatureId"`
}

// GetSignatureResult defines the data structure of the result of getting a signature
type GetSignatureResult struct {
	SignatureId string `json:"signatureId"`
	UserId      string `json:"userId"`
	Content     string `json:"content"`
	ContentType string `json:"contentType"`
	Status      string `json:"status"`
	CountryType string `json:"countryType"`
	Review      string `json:"review"`
}

// CreateTemplateArgs defines the data structure for creating a template
type CreateTemplateArgs struct {
	Name        string `json:"name"`
	Content     string `json:"content"`
	SmsType     string `json:"smsType"`
	CountryType string `json:"countryType"`
	Description string `json:"description,omitempty"`
}

// CreateTemplateResult defines the data structure of the result of creating a template
type CreateTemplateResult struct {
	TemplateId string `json:"templateId"`
	Status     string `json:"status"`
}

// DeleteTemplateArgs defines the data structure for deleting a template
type DeleteTemplateArgs struct {
	TemplateId string `json:"templateId"`
}

// ModifyTemplateArgs defines the data structure for modifying a template
type ModifyTemplateArgs struct {
	TemplateId  string `json:"templateId"`
	Name        string `json:"name"`
	Content     string `json:"content"`
	SmsType     string `json:"smsType"`
	CountryType string `json:"countryType"`
	Description string `json:"description,omitempty"`
}

// GetTemplateArgs defines the data structure for getting a template
type GetTemplateArgs struct {
	TemplateId string `json:"templateId"`
}

// GetTemplateResult defines the data structure of the result of getting a template
type GetTemplateResult struct {
	TemplateId  string `json:"templateId"`
	UserId      string `json:"userId"`
	Name        string `json:"name"`
	Content     string `json:"content"`
	CountryType string `json:"countryType"`
	SmsType     string `json:"smsType"`
	Status      string `json:"status"`
	Description string `json:"description"`
	Review      string `json:"review"`
}

// UpdateQuotaRateArgs defines the data structure for updating quota and rate limit
type UpdateQuotaRateArgs struct {
	QuotaPerDay        int `json:"quotaPerDay"`
	QuotaPerMonth      int `json:"quotaPerMonth"`
	RateLimitPerDay    int `json:"rateLimitPerMobilePerSignByDay"`
	RateLimitPerHour   int `json:"rateLimitPerMobilePerSignByHour"`
	RateLimitPerMinute int `json:"rateLimitPerMobilePerSignByMinute"`
}

// QueryQuotaRateResult defines the data structure of querying the user's quota and rate limit
type QueryQuotaRateResult struct {
	QuotaPerDay          int    `json:"quotaPerDay"`
	QuotaRemainToday     int    `json:"quotaRemainToday"`
	QuotaPerMonth        int    `json:"quotaPerMonth"`
	QuotaRemainThisMonth int    `json:"quotaRemainThisMonth"`
	ApplyQuotaPerDay     int    `json:"applyQuotaPerDay"`
	ApplyQuotaPerMonth   int    `json:"applyQuotaPerMonth"`
	ApplyCheckStatus     string `json:"applyCheckStatus"`
	ApplyCheckReply      string `json:"checkReply"`
	RateLimitPerDay      int    `json:"rateLimitPerMobilePerSignByDay"`
	RateLimitPerHour     int    `json:"rateLimitPerMobilePerSignByHour"`
	RateLimitPerMinute   int    `json:"rateLimitPerMobilePerSignByMinute"`
	RateLimitWhitelist   bool   `json:"rateLimitWhitelist"`
}

// CreateMobileBlackArgs defines the data structure for creating a mobileBlack
type CreateMobileBlackArgs struct {
	Type           string `json:"type"`
	SmsType        string `json:"smsType"`
	SignatureIdStr string `json:"signatureIdStr"`
	Phone          string `json:"phone"`
	CountryType    string `json:"countryType"`
}

// DeleteMobileBlackArgs defines the data structure for deleting mobileBlack by phones
type DeleteMobileBlackArgs struct {
	Phones string `json:"phones"`
}

// GetMobileBlackArgs defines the data structure for get mobileBlackList
// startTime、endTime format is yyyy-MM-dd
type GetMobileBlackArgs struct {
	Phone          string
	CountryType    string
	SmsType        string
	SignatureIdStr string
	StartTime      string
	EndTime        string
	PageNo         string
	PageSize       string
}

// GetMobileBlackResult defines the data structure for get mobileBlackList
type GetMobileBlackResult struct {
	TotalCount int                 `json:"totalCount"`
	PageNo     int                 `json:"pageNo"`
	PageSize   int                 `json:"pageSize"`
	BlackLists []MobileBlackDetail `json:"blacklists"`
}

// MobileBlackDetail defines the data structure for mobileBlackList detail
type MobileBlackDetail struct {
	Phone          string `json:"phone"`
	CountryType    string `json:"countryType"`
	Type           string `json:"type"`
	SmsType        string `json:"smsType"`
	SignatureIdStr string `json:"signatureIdStr"`
	UpdateDate     string `json:"updateDate"`
}

// ListStatisticsArgs defines the request data structure of ListStatistics
type ListStatisticsArgs struct {
	SmsType      string `json:"smsType"`
	SignatureId  string `json:"signatureId"`
	TemplateCode string `json:"TemplateCode"`
	CountryType  string `json:"countryType"` // available values: "domestic", "international"
	StartTime    string `json:"startTime"`   // format: "yyyy-MM-dd"
	EndTime      string `json:"endTime"`     // format: "yyyy-MM-dd"
}

// ListStatisticsResponse defines the response data structure of ListStatistics
type ListStatisticsResponse struct {
	StatisticsResults []StatisticsResult `json:"statisticsResults"`
}

// StatisticsResult defines the detail of ListStatisticsResponse
type StatisticsResult struct {
	Datetime                  string `json:"datetime"`
	CountryAlpha2Code         string `json:"countryAlpha2Code"`
	SubmitCount               string `json:"submitCount"`
	SubmitLongCount           string `json:"submitLongCount"`
	ResponseSuccessCount      string `json:"responseSuccessCount"`
	ResponseSuccessProportion string `json:"responseSuccessProportion"`
	DeliverSuccessCount       string `json:"deliverSuccessCount"`
	DeliverSuccessLongCount   string `json:"deliverSuccessLongCount"`
	DeliverSuccessProportion  string `json:"deliverSuccessProportion"`
	DeliverFailureCount       string `json:"deliverFailureCount"`
	DeliverFailureProportion  string `json:"deliverFailureProportion"`
	ReceiptProportion         string `json:"receiptProportion"`
	UnknownCount              string `json:"unknownCount"`
	UnknownProportion         string `json:"unknownProportion"`
	ResponseTimeoutCount      string `json:"responseTimeoutCount"`
	UnknownErrorCount         string `json:"unknownErrorCount"`
	NotExistCount             string `json:"notExistCount"`
	SignatureOrTemplateCount  string `json:"signatureOrTemplateCount"`
	AbnormalCount             string `json:"abnormalCount"`
	OverclockingCount         string `json:"overclockingCount"`
	OtherErrorCount           string `json:"otherErrorCount"`
	BlacklistCount            string `json:"blacklistCount"`
	RouteErrorCount           string `json:"routeErrorCount"`
	IssueFailureCount         string `json:"issueFailureCount"`
	ParameterErrorCount       string `json:"parameterErrorCount"`
	IllegalWordCount          string `json:"illegalWordCount"`
	AnomalyCount              string `json:"anomalyCount"`
}

// GetPrepaidPackageArgs defines the request data structure of PrepaidPackages
type GetPrepaidPackageArgs struct {
	UserID        string `json:"userId"`
	CountryType   string `json:"countryType"`   // available values: "domestic", "international"， "global"
	PackageStatus string `json:"packageStatus"` // available values: "RUNNING", "EXPIRED", "USED_UP", "DESTROYED"
	PageNo        string `json:"pageNo"`
	PageSize      string `json:"pageSize"`
}

// GetPrepaidPackageResponse defines the response data structure of PrepaidPackages
type GetPrepaidPackageResponse struct {
	PrepaidPackages []PrepaidPackage `json:"prepaidPackages"`
	TotalCount      int              `json:"totalCount"`
}

type PrepaidPackage struct {
	PackageId         string  `json:"packageId"`
	Name              string  `json:"name"`
	CountryType       string  `json:"countryType"`
	Capacity          float64 `json:"capacity"`
	RemainingCapacity float64 `json:"remainingCapacity"`
	PackageStatus     string  `json:"packageStatus"`
	PurchaseDate      string  `json:"purchaseDate"`
	ExpireDate        string  `json:"expireDate"`
}
