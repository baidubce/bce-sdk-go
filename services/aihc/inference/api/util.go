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

// util.go - define the utilities for api package of aihc inference service

package api

import "github.com/baidubce/bce-sdk-go/bce"

const (
	API      = "/api"
	AIHC_POM = "/aihcpom"

	URI_PREFIXV1 = bce.URI_PREFIX + "v1"

	TYPE_RESPOOL = "/respool"
	TYPE_APP     = "/app"
	TYPE_POD     = "/pod"

	REQUEST_CREATE       = "/create"
	REQUEST_LIST         = "/list"
	REQUEST_STATS        = "/stats"
	REQUEST_DETAILS      = "/details"
	REQUEST_DETAIL       = "/detail"
	REQUEST_UPDATE       = "/update"
	REQUEST_SCALE        = "/scale"
	REQUEST_PUBACCESS    = "/pubaccess"
	REQUEST_LISTCHANGE   = "/listchange"
	REQUEST_CHANGEDETAIL = "/changedetail"
	REQUEST_DELETE       = "/delete"
	REQUEST_BLOCK        = "/block"
	REQUEST_LISTBRIEF    = "/listbrief"
)

func createAppUri() string {
	return API + URI_PREFIXV1 + AIHC_POM + TYPE_APP + REQUEST_CREATE
}

func listAppUri() string {
	return API + URI_PREFIXV1 + AIHC_POM + TYPE_APP + REQUEST_LIST
}

func listAppStatUri() string {
	return API + URI_PREFIXV1 + AIHC_POM + TYPE_APP + REQUEST_STATS
}

func appDetailsUri() string {
	return API + URI_PREFIXV1 + AIHC_POM + TYPE_APP + REQUEST_DETAILS
}

func updateAppUri() string {
	return API + URI_PREFIXV1 + AIHC_POM + TYPE_APP + REQUEST_UPDATE
}

func scaleAppUri() string {
	return API + URI_PREFIXV1 + AIHC_POM + TYPE_APP + REQUEST_SCALE
}

func pubAccessUri() string {
	return API + URI_PREFIXV1 + AIHC_POM + TYPE_APP + REQUEST_PUBACCESS
}

func listChangeUri() string {
	return API + URI_PREFIXV1 + AIHC_POM + TYPE_APP + REQUEST_LISTCHANGE
}

func changeDetailUri() string {
	return API + URI_PREFIXV1 + AIHC_POM + TYPE_APP + REQUEST_CHANGEDETAIL
}

func deleteAppUri() string {
	return API + URI_PREFIXV1 + AIHC_POM + TYPE_APP + REQUEST_DELETE
}

func listPodUri() string {
	return API + URI_PREFIXV1 + AIHC_POM + TYPE_POD + REQUEST_LIST
}

func blockPodUri() string {
	return API + URI_PREFIXV1 + AIHC_POM + TYPE_POD + REQUEST_BLOCK
}

func deletePodUri() string {
	return API + URI_PREFIXV1 + AIHC_POM + TYPE_POD + REQUEST_DELETE
}

func listBriefResPoolUri() string {
	return API + URI_PREFIXV1 + AIHC_POM + TYPE_RESPOOL + REQUEST_LISTBRIEF
}

func resPoolDetailUri() string {
	return API + URI_PREFIXV1 + AIHC_POM + TYPE_RESPOOL + REQUEST_DETAIL
}
