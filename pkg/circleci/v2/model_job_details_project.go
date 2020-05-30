// Copyright 2020 The ccictl Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package v2

import (
	json "github.com/goccy/go-json"
)

// JobDetailsProject Information about a project.
type JobDetailsProject struct {
	// Project slug in the form `vcs-slug/org-name/repo-name`. The `/` characters may be URL-escaped.
	Slug string `json:"slug"`
	// The name of the project
	Name string `json:"name"`
	// URL to the repository hosting the project's code
	ExternalUrl string `json:"external_url"`
}

// NewJobDetailsProject instantiates a new JobDetailsProject object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewJobDetailsProject(slug string, name string, externalUrl string) *JobDetailsProject {
	this := JobDetailsProject{}
	this.Slug = slug
	this.Name = name
	this.ExternalUrl = externalUrl
	return &this
}

// NewJobDetailsProjectWithDefaults instantiates a new JobDetailsProject object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewJobDetailsProjectWithDefaults() *JobDetailsProject {
	this := JobDetailsProject{}
	return &this
}

// GetSlug returns the Slug field value
func (o *JobDetailsProject) GetSlug() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Slug
}

// GetSlugOk returns a tuple with the Slug field value
// and a boolean to check if the value has been set.
func (o *JobDetailsProject) GetSlugOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Slug, true
}

// SetSlug sets field value
func (o *JobDetailsProject) SetSlug(v string) {
	o.Slug = v
}

// GetName returns the Name field value
func (o *JobDetailsProject) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *JobDetailsProject) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *JobDetailsProject) SetName(v string) {
	o.Name = v
}

// GetExternalUrl returns the ExternalUrl field value
func (o *JobDetailsProject) GetExternalUrl() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ExternalUrl
}

// GetExternalUrlOk returns a tuple with the ExternalUrl field value
// and a boolean to check if the value has been set.
func (o *JobDetailsProject) GetExternalUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ExternalUrl, true
}

// SetExternalUrl sets field value
func (o *JobDetailsProject) SetExternalUrl(v string) {
	o.ExternalUrl = v
}

func (o JobDetailsProject) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["slug"] = o.Slug
	}
	if true {
		toSerialize["name"] = o.Name
	}
	if true {
		toSerialize["external_url"] = o.ExternalUrl
	}
	return json.Marshal(toSerialize)
}

type NullableJobDetailsProject struct {
	value *JobDetailsProject
	isSet bool
}

func (v NullableJobDetailsProject) Get() *JobDetailsProject {
	return v.value
}

func (v *NullableJobDetailsProject) Set(val *JobDetailsProject) {
	v.value = val
	v.isSet = true
}

func (v NullableJobDetailsProject) IsSet() bool {
	return v.isSet
}

func (v *NullableJobDetailsProject) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableJobDetailsProject(val *JobDetailsProject) *NullableJobDetailsProject {
	return &NullableJobDetailsProject{value: val, isSet: true}
}

func (v NullableJobDetailsProject) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableJobDetailsProject) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
