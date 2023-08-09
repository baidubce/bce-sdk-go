/*
 * Copyright 2023 Baidu, Inc.
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

package afd

// SyncArgs def
type SyncArgs struct {
	SC    string `json:"sc"`
	TS    string `json:"ts"`
	M     string `json:"m"`
	IP    string `json:"ip"`
	App   string `json:"app"`
	AppID string `json:"appid"`
	AID   string `json:"aid"`
	EV    string `json:"ev"`

	SV       string `json:"sv,omitempty"`
	Email    string `json:"email,omitempty"`
	Username string `json:"username,omitempty"`

	LR           string `json:"lr,omitempty"`
	FR           string `json:"fr,omitempty"`
	RTS          string `json:"rts,omitempty"`
	RIP          string `json:"rip,omitempty"`
	LoginType    string `json:"logintype,omitempty"`
	LoginActType string `json:"loginactype,omitempty"`

	Z         string `json:"z,omitempty"`
	I         string `json:"i,omitempty"`
	MAC       string `json:"mac,omitempty"`
	IDFA      string `json:"idfa,omitempty"`
	IDFV      string `json:"idfv,omitempty"`
	UserID    string `json:"userid,omitempty"`
	Ver       string `json:"ver,omitempty"`
	Model     string `json:"model,omitempty"`
	UA        string `json:"ua,omitempty"`
	BSSID     string `json:"bssid,omitempty"`
	SSID      string `json:"ssid,omitempty"`
	InviteRID string `json:"inviterid,omitempty"`
	LikeDID   string `json:"likedid,omitempty"`
	Cash      string `json:"cash,omitempty"`
	CT        string `json:"ct,omitempty"`
	LAL       string `json:"lal,omitempty"`
	CSR       string `json:"csr,omitempty"`
	Referer   string `json:"referer,omitempty"`
	Net       string `json:"net,omitempty"`
	JT        string `json:"jt,omitempty"`
	JSEnv     string `json:"js_env,omitempty"`

	Header map[string]string `json:"header,omitempty"`
	Extra  map[string]string `json:"extra,omitempty"`
}

// SyncResponse def
type SyncResponse struct {
	RequestID     string `json:"request_id"`
	ReturnCode    string `json:"ret_code"`
	RerurnMessage string `json:"ret_msg"`
	ReturnData    struct {
		Level string   `json:"level"`
		Tags  []string `json:"t"`
	} `json:"ret_data,omitempty"`
}

// FactorArgs def
type FactorArgs struct {
	Z     string `json:"z"`
	App   string `json:"app"`
	JSEnv string `json:"js_env,omitempty"`
	JT    string `json:"jt,omitempty"`
}

// FactorResponse def
type FactorResponse struct {
	RequestID     string `json:"request_id"`
	ReturnCode    string `json:"ret_code"`
	RerurnMessage string `json:"ret_msg"`
	ReturnData    struct {
		X    string   `json:"x"`
		Tags []string `json:"t"`
	} `json:"ret_data,omitempty"`
	JID   string `json:"jid,omitempty"`
	JTags string `json:"jtag,omitempty"`
}
