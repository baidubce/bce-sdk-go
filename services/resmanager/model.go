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

package resmanager

type Binding struct {
	ResourceId     string `json:"resourceId"`
	ResourceType   string `json:"resourceType"`
	ResourceRegion string `json:"resourceRegion"`
	GroupId        string `json:"groupId"`
}

type BindResourceToGroupArgs struct {
	Bindings []Binding `json:"bindings"`
}

type Group struct {
	Name       string `json:"name"`
	Extra      string `json:"extra"`
	ParentUUID string `json:"parentUuid"`
	GroupId    string `json:"groupId"`
}

type ResGroup struct {
	AccountID string  `json:"accountId"`
	UserID    string  `json:"userId"`
	Name      string  `json:"name"`
	Type      string  `json:"type"`
	Region    string  `json:"region"`
	ID        string  `json:"id"`
	UUID      string  `json:"uuid"`
	Summary   string  `json:"summary"`
	URL       string  `json:"url"`
	Groups    []Group `json:"groups"`
}

type BindResourceResult struct {
	Status    string     `json:"status"`
	ResGroups []ResGroup `json:"resGroups"`
}

type ChangeResourceGroupArgs struct {
	MoveResModels []MoveResModel `json:"moveResModels"`
}

type MoveResModel struct {
	TargetGroupId   string          `json:"targetGroupId"`
	OldGroupResInfo OldGroupResInfo `json:"oldGroupResInfo"`
}

type OldGroupResInfo struct {
	ResourceId     string `json:"reourceId"`
	ResourceType   string `json:"resourceType"`
	ResourceRegion string `json:"resourceRegion"`
	GroupId        string `json:"groupId"`
}

type GroupTree struct {
	ParentID string      `json:"parentId"`
	GroupID  string      `json:"groupId"`
	Name     string      `json:"name"`
	Extra    string      `json:"extra"`
	Children []GroupTree `json:"children"`
}

type GroupList struct {
	GroupTrees []GroupTree `json:"groups"`
}
