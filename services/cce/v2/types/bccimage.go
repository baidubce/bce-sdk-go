// Copyright 2019 Baidu Inc. All rights reserved
// Use of this source code is governed by a CCE
// license that can be found in the LICENSE file.
/*
modification history
--------------------
2020/07/28 16:26:00, by jichao04@baidu.com, create
*/

package types

import bccapi "github.com/baidubce/bce-sdk-go/services/bcc/api"

const (
	ImageTypeService bccapi.ImageType = "service"
)

// OSType 操作系统类型
type OSType string

const (
	// OSTypeLinux linux
	OSTypeLinux OSType = "linux"
	// OSTypeWindows windows
	OSTypeWindows OSType = "windows"
)

// OSName 操作系统名字
type OSName string

const (
	// OSNameCentOS centos
	OSNameCentOS OSName = "CentOS"
	// OSNameUbuntu ubuntu
	OSNameUbuntu OSName = "Ubuntu"
	// OSNameWindows windows
	OSNameWindows OSName = "Windows Server"
	// OSNameDebian debian
	OSNameDebian OSName = "Debian"
	// OSNameOpensuse opensuse
	OSNameOpensuse OSName = "opensuse"
)

