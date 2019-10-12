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

// util.go - define the utilities for api package of BCC service
package api

import (
	"encoding/hex"
	"fmt"

	"github.com/baidubce/bce-sdk-go/bce"
	"github.com/baidubce/bce-sdk-go/util/crypto"
)

const (
	URI_PREFIX = bce.URI_PREFIX + "v2"

	REQUEST_INSTANCE_URI = "/instance"
	REQUEST_VNC_SUFFIX   = "/vnc"

	REQUEST_VOLUME_URI           = "/volume"
	REQUEST_SECURITYGROUP_URI    = "/securityGroup"
	REQUEST_IMAGE_URI            = "/image"
	REQUEST_IMAGE_SHAREDUSER_URI = "/sharedUsers"
	REQUEST_IMAGE_OS_URI         = "/os"

	REQUEST_SNAPSHOT_URI = "/snapshot"
	REQUEST_ASP_URI      = "/asp"
	REQUEST_SPEC_URI     = "/instance/spec"
	REQUEST_ZONE_URI     = "/zone"

	REQUEST_SUBNET_URI = "/subnet"
)

func getInstanceUri() string {
	return URI_PREFIX + REQUEST_INSTANCE_URI
}

func getInstanceUriWithId(id string) string {
	return URI_PREFIX + REQUEST_INSTANCE_URI + "/" + id
}

func getInstanceVNCUri(id string) string {
	return URI_PREFIX + REQUEST_INSTANCE_URI + "/" + id + REQUEST_VNC_SUFFIX
}

func Aes128EncryptUseSecreteKey(sk string, data string) (string, error) {
	if len(sk) < 16 {
		return "", fmt.Errorf("error secrete key")
	}

	crypted, err := crypto.EBCEncrypto([]byte(sk[:16]), []byte(data))
	if err != nil {
		return "", err
	}

	return hex.EncodeToString(crypted), nil
}

func getVolumeUri() string {
	return URI_PREFIX + REQUEST_VOLUME_URI
}

func getVolumeUriWithId(id string) string {
	return URI_PREFIX + REQUEST_VOLUME_URI + "/" + id
}

func getSecurityGroupUri() string {
	return URI_PREFIX + REQUEST_SECURITYGROUP_URI
}

func getSecurityGroupUriWithId(id string) string {
	return URI_PREFIX + REQUEST_SECURITYGROUP_URI + "/" + id
}

func getImageUri() string {
	return URI_PREFIX + REQUEST_IMAGE_URI
}

func getImageUriWithId(id string) string {
	return URI_PREFIX + REQUEST_IMAGE_URI + "/" + id
}

func getImageSharedUserUri(id string) string {
	return URI_PREFIX + REQUEST_IMAGE_URI + "/" + id + REQUEST_IMAGE_SHAREDUSER_URI
}

func getImageOsUri() string {
	return URI_PREFIX + REQUEST_IMAGE_URI + REQUEST_IMAGE_OS_URI
}

func getSnapshotUri() string {
	return URI_PREFIX + REQUEST_SNAPSHOT_URI
}

func getSnapshotUriWithId(id string) string {
	return URI_PREFIX + REQUEST_SNAPSHOT_URI + "/" + id
}

func getASPUri() string {
	return URI_PREFIX + REQUEST_ASP_URI
}

func getASPUriWithId(id string) string {
	return URI_PREFIX + REQUEST_ASP_URI + "/" + id
}

func getSpecUri() string {
	return URI_PREFIX + REQUEST_SPEC_URI
}

func getZoneUri() string {
	return URI_PREFIX + REQUEST_ZONE_URI
}

func getChangeSubnetUri() string {
	return URI_PREFIX + REQUEST_SUBNET_URI + "/changeSubnet"
}
