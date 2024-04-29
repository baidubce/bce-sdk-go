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

package tag

type Tag struct {
	TagKey   string `json:"tagKey"`
	TagValue string `json:"tagValue"`
}

type TagsRequest struct {
	Tags []Tag `json:"tags"`
}

type TagsResult struct {
	Tags []Tag `json:"tags"`
}

type Resource struct {
	ResourceId  string `json:"resourceId"`
	ServiceType string `json:"serviceType"`
	Region      string `json:"region"`
}

type TagAssociationFull struct {
	TagAssociation
	TagKey          string `json:"tagKey"`
	TagValue        string `json:"tagValue"`
	AssociationType int    `json:"associationType"`
}

type TagAssociation struct {
	AccountId    string `json:"accountId"`
	ResourceId   string `json:"resourceId"`
	ResourceUuid string `json:"resourceUuid"`
	Region       string `json:"region"`
	ServiceType  string `json:"serviceType"`
	TagId        int64  `json:"tagId"`
}

type TagResource struct {
	ResourceType string `json:"resourceType"`
	ResourceId   string `json:"resourceId"`
	Region       string `json:"region"`
}

type ResourceWithTag struct {
	ResourceId      string `json:"resourceId"`
	ResourceUuid    string `json:"resourceUuid"`
	ServiceType     string `json:"serviceType"`
	Region          string `json:"region"`
	Tags            []Tag  `json:"tags"`
	AssociationType string `json:"associationType"`
}

type TagAssociationWithResources struct {
	Tag
	Resources []TagResource `json:"resources"`
}

type TagsAssociationWithResources struct {
	TagResources []TagAssociationWithResources `json:"tagResources"`
}

type DeleteAssociationsByTagRequest struct {
	TagKey      string     `json:"tagKey"`
	TagValue    string     `json:"tagValue"`
	ServiceType string     `json:"serviceType"`
	Resource    []Resource `json:"resources"`
}

type DeleteTagAssociationRequest struct {
	Resource *Resource `json:"resource"`
}

type CreateAssociationsByTagRequest struct {
	TagKey      string     `json:"tagKey"`
	TagValue    string     `json:"tagValue"`
	ServiceType string     `json:"serviceType"`
	Resource    []Resource `json:"resources"`
}

type CreateAndAssignTagRequest struct {
	ResourceWithTag []ResourceWithTag `json:"resources"`
}

type FullTagListRequest struct {
	TagKey       string   `json:"tagKey,omitempty"`
	TagValue     string   `json:"tagValue,omitempty"`
	Regions      []string `json:"regions,omitempty"`
	ServiceTypes []string `json:"serviceTypes,omitempty"`
	ResourceIds  []string `json:"resourceIds,omitempty"`
}

type TagAssociationFulls struct {
	TagAssociationFull []TagAssociationFull `json:"tagAssociationFulls"`
}
