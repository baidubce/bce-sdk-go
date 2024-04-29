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

// tag.go - the resmanager APIs definition supported by the resmanager service
package tag

import (
	"fmt"
	"github.com/baidubce/bce-sdk-go/bce"
	"github.com/baidubce/bce-sdk-go/http"
	"strconv"
)

// CreateTags https://cloud.baidu.com/doc/TAG/s/Okbrb3ral
func (c *Client) CreateTags(args *TagsRequest) error {
	if args == nil {
		return fmt.Errorf("unset args")
	}
	if len(args.Tags) == 0 {
		return fmt.Errorf("unset tags")
	}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.PUT).
		WithURL(tagBaseUri()).
		WithQueryParam("create", "").
		WithBody(args).
		WithHeader(http.CONTENT_TYPE, bce.DEFAULT_CONTENT_TYPE).
		Do()
	return err
}

// DeleteTags https://cloud.baidu.com/doc/TAG/s/Xkbrb3rhr
func (c *Client) DeleteTags(args *TagsRequest) error {
	if args == nil {
		return fmt.Errorf("unset args")
	}
	if len(args.Tags) == 0 {
		return fmt.Errorf("unset tags")
	}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.PUT).
		WithURL(tagBaseUri()).
		WithQueryParam("delete", "").
		WithBody(args).
		WithHeader(http.CONTENT_TYPE, bce.DEFAULT_CONTENT_TYPE).
		Do()
	return err
}

// TagsResources https://cloud.baidu.com/doc/TAG/s/Bkbrb3roy
func (c *Client) TagsResources(tagKey string, tagValue string, region string, resourceType string) (*TagsAssociationWithResources, error) {
	result := &TagsAssociationWithResources{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.GET).
		WithURL(tagsResourcesUri()).
		WithQueryParamFilter("tagKey", tagKey).
		WithQueryParamFilter("tagValue", tagValue).
		WithQueryParamFilter("region", region).
		WithQueryParamFilter("resourceType", resourceType).
		WithHeader(http.CONTENT_TYPE, bce.DEFAULT_CONTENT_TYPE).
		WithResult(result).
		Do()
	return result, err
}

// UserTagList https://cloud.baidu.com/doc/TAG/s/Ukbrb3r3d
func (c *Client) UserTagList(tagKey, tagValue string) (*TagsResult, error) {

	result := &TagsResult{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.GET).
		WithURL(tagBaseUri()).
		WithQueryParamFilter("tagKey", tagKey).
		WithQueryParamFilter("tagValue", tagValue).
		WithHeader(http.CONTENT_TYPE, bce.DEFAULT_CONTENT_TYPE).
		WithResult(result).
		Do()
	return result, err
}

// CreateAssociationsByTag https://cloud.baidu.com/doc/TAG/s/rkm1yqvhz
func (c *Client) CreateAssociationsByTag(args *CreateAssociationsByTagRequest) error {
	if args == nil {
		return fmt.Errorf("unset args")
	}
	if len(args.TagKey) == 0 {
		return fmt.Errorf("unset tagKey")
	}
	if len(args.TagValue) == 0 {
		return fmt.Errorf("unset tagValue")
	}
	if len(args.ServiceType) == 0 {
		return fmt.Errorf("unset serviceType")
	}
	if len(args.Resource) == 0 {
		return fmt.Errorf("unset resource")
	}

	err := bce.NewRequestBuilder(c).
		WithMethod(http.POST).
		WithURL(resourceBaseUri()).
		WithBody(args).
		WithHeader(http.CONTENT_TYPE, bce.DEFAULT_CONTENT_TYPE).
		Do()
	return err
}

// DeleteAssociationsByTag https://cloud.baidu.com/doc/TAG/s/7km1zf2j2
func (c *Client) DeleteAssociationsByTag(args *DeleteAssociationsByTagRequest) error {
	if args == nil {
		return fmt.Errorf("unset args")
	}
	if len(args.TagKey) == 0 {
		return fmt.Errorf("unset tagKey")
	}
	if len(args.TagValue) == 0 {
		return fmt.Errorf("unset tagValue")
	}
	if len(args.ServiceType) == 0 {
		return fmt.Errorf("unset serviceType")
	}
	if len(args.Resource) == 0 {
		return fmt.Errorf("unset resource")
	}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.PUT).
		WithURL(resourceBaseUri()).
		WithBody(args).
		WithHeader(http.CONTENT_TYPE, bce.DEFAULT_CONTENT_TYPE).
		Do()
	return err
}

// DeleteTagAssociation https://cloud.baidu.com/doc/TAG/s/wkm1sfady
func (c *Client) DeleteTagAssociation(args *DeleteTagAssociationRequest) error {
	if args == nil {
		return fmt.Errorf("unset args")
	}

	if args.Resource == nil {
		return fmt.Errorf("unset resource")
	}

	err := bce.NewRequestBuilder(c).
		WithMethod(http.PUT).
		WithURL(tagBaseUri()).
		WithQueryParam("deleteTagAssociation", "").
		WithBody(args).
		WithHeader(http.CONTENT_TYPE, bce.DEFAULT_CONTENT_TYPE).
		Do()
	return err
}

// CreateAndAssign https://cloud.baidu.com/doc/TAG/s/Wkm1t1xca
func (c *Client) CreateAndAssign(args *CreateAndAssignTagRequest) error {
	if args == nil {
		return fmt.Errorf("unset args")
	}
	if len(args.ResourceWithTag) == 0 {
		return fmt.Errorf("unset resources")
	}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.POST).
		WithURL(tagBaseUri()).
		WithQueryParam("createAndAssign", "").
		WithBody(args).
		WithHeader(http.CONTENT_TYPE, bce.DEFAULT_CONTENT_TYPE).
		Do()
	return err
}

// QueryFullList https://cloud.baidu.com/doc/TAG/s/Ekm1tunjn
func (c *Client) QueryFullList(strongAssociation bool, args *FullTagListRequest) (*TagAssociationFulls, error) {
	if args == nil {
		return nil, fmt.Errorf("unset args")
	}
	result := &TagAssociationFulls{}
	err := bce.NewRequestBuilder(c).
		WithMethod(http.POST).
		WithURL(queryFullListUri()).
		WithQueryParamFilter("strongAssociation", strconv.FormatBool(strongAssociation)).
		WithBody(args).
		WithHeader(http.CONTENT_TYPE, bce.DEFAULT_CONTENT_TYPE).
		WithResult(result).
		Do()
	return result, err
}
