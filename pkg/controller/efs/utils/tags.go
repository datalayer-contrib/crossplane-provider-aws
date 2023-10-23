/*
Copyright 2021 The Crossplane Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package utils

import (
	"sort"
	"strings"

	svcsdk "github.com/aws/aws-sdk-go/service/efs"
	"github.com/aws/aws-sdk-go/service/efs/efsiface"
	"github.com/crossplane/crossplane-runtime/pkg/resource"
	"github.com/pkg/errors"

	svcapitypes "github.com/crossplane-contrib/provider-aws/apis/efs/v1alpha1"
	"github.com/crossplane-contrib/provider-aws/pkg/utils/pointer"
)

const (
	errListTagsForResource = "cannot list tags"
	errRemoveTags          = "cannot remove tags"
	errCreateTags          = "cannot create tags"
)

// AreTagsUpToDate for spec and resourceID
func AreTagsUpToDate(client efsiface.EFSAPI, spec []*svcapitypes.Tag, resourceID *string) (bool, error) {
	current, err := ListTagsForResource(client, resourceID)
	if err != nil {
		return false, err
	}

	add, remove := DiffTags(spec, current)

	return len(add) == 0 && len(remove) == 0, nil
}

// UpdateTagsForResource with resourceID
func UpdateTagsForResource(client efsiface.EFSAPI, spec []*svcapitypes.Tag, resourceID *string) error {
	current, err := ListTagsForResource(client, resourceID)
	if err != nil {
		return err
	}

	add, remove := DiffTags(spec, current)
	if len(remove) != 0 {
		if _, err := client.UntagResource(&svcsdk.UntagResourceInput{
			ResourceId: resourceID,
			TagKeys:    remove,
		}); err != nil {
			return errors.Wrap(err, errRemoveTags)
		}
	}
	if len(add) != 0 {
		if _, err := client.TagResource(&svcsdk.TagResourceInput{
			ResourceId: resourceID,
			Tags:       add,
		}); err != nil {
			return errors.Wrap(err, errCreateTags)
		}
	}

	return nil
}

// ListTagsForResource for the given resource
func ListTagsForResource(client efsiface.EFSAPI, resourceID *string) ([]*svcsdk.Tag, error) {
	req := &svcsdk.ListTagsForResourceInput{
		ResourceId: resourceID,
	}

	resp, err := client.ListTagsForResource(req)
	if err != nil {
		return nil, errors.Wrap(err, errListTagsForResource)
	}

	return resp.Tags, nil
}

// DiffTags between spec and current
func DiffTags(spec []*svcapitypes.Tag, current []*svcsdk.Tag) (addTags []*svcsdk.Tag, removeTags []*string) {
	currentMap := make(map[string]string, len(current))
	for _, t := range current {
		currentMap[pointer.StringValue(t.Key)] = pointer.StringValue(t.Value)
	}

	specMap := make(map[string]string, len(spec))
	for _, t := range spec {
		key := pointer.StringValue(t.Key)

		// Ignore "aws:" internal tags since they cannot be added or removed.
		if strings.HasPrefix(key, "aws:") {
			continue
		}

		val := pointer.StringValue(t.Value)
		specMap[key] = pointer.StringValue(t.Value)

		if currentVal, exists := currentMap[key]; exists {
			if currentVal != val {
				removeTags = append(removeTags, t.Key)
				addTags = append(addTags, &svcsdk.Tag{
					Key:   pointer.String(key),
					Value: pointer.String(val),
				})
			}
		} else {
			addTags = append(addTags, &svcsdk.Tag{
				Key:   pointer.String(key),
				Value: pointer.String(val),
			})
		}
	}

	for _, t := range current {
		key := pointer.StringValue(t.Key)

		// Ignore "aws:" internal tags since they cannot be added or removed.
		if strings.HasPrefix(key, "aws:") {
			continue
		}

		if _, exists := specMap[key]; !exists {
			removeTags = append(removeTags, pointer.String(key))
		}
	}

	return addTags, removeTags
}

// AddExternalTags to spec if they don't exist
func AddExternalTags(mg resource.Managed, spec []*svcapitypes.Tag) []*svcapitypes.Tag {
	tagMap := make(map[string]struct{}, len(spec))
	for _, t := range spec {
		tagMap[pointer.StringValue(t.Key)] = struct{}{}
	}

	tags := spec
	for _, t := range GetExternalTags(mg) {
		if _, exists := tagMap[pointer.StringValue(t.Key)]; !exists {
			tags = append(tags, t)
		}
	}

	return tags
}

// GetExternalTags is a wrapper around resource.GetExternalTags to return a sorted array instead of a map
func GetExternalTags(mg resource.Managed) []*svcapitypes.Tag {
	externalTags := []*svcapitypes.Tag{}
	for k, v := range resource.GetExternalTags(mg) {
		externalTags = append(externalTags, &svcapitypes.Tag{Key: pointer.String(k), Value: pointer.String(v)})
	}

	sort.Slice(externalTags, func(i, j int) bool {
		return pointer.StringValue(externalTags[i].Key) > pointer.StringValue(externalTags[j].Key)
	})

	return externalTags
}