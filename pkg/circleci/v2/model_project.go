// Copyright 2020 The ccictl Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package v2

import (
	json "github.com/goccy/go-json"
)

// Project NOTE: The definition of Project is subject to change.
type Project struct {
	// Project slug in the form `vcs-slug/org-name/repo-name`. The `/` characters may be URL-escaped.
	Slug string `json:"slug"`
	// The name of the project
	Name string `json:"name"`
	// The name of the organization the project belongs to
	OrganizationName string         `json:"organization_name"`
	VcsInfo          ProjectVcsInfo `json:"vcs_info"`
}

// NewProject instantiates a new Project object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewProject(slug string, name string, organizationName string, vcsInfo ProjectVcsInfo) *Project {
	this := Project{}
	this.Slug = slug
	this.Name = name
	this.OrganizationName = organizationName
	this.VcsInfo = vcsInfo
	return &this
}

// NewProjectWithDefaults instantiates a new Project object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewProjectWithDefaults() *Project {
	this := Project{}
	return &this
}

// GetSlug returns the Slug field value
func (o *Project) GetSlug() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Slug
}

// GetSlugOk returns a tuple with the Slug field value
// and a boolean to check if the value has been set.
func (o *Project) GetSlugOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Slug, true
}

// SetSlug sets field value
func (o *Project) SetSlug(v string) {
	o.Slug = v
}

// GetName returns the Name field value
func (o *Project) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *Project) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *Project) SetName(v string) {
	o.Name = v
}

// GetOrganizationName returns the OrganizationName field value
func (o *Project) GetOrganizationName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.OrganizationName
}

// GetOrganizationNameOk returns a tuple with the OrganizationName field value
// and a boolean to check if the value has been set.
func (o *Project) GetOrganizationNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.OrganizationName, true
}

// SetOrganizationName sets field value
func (o *Project) SetOrganizationName(v string) {
	o.OrganizationName = v
}

// GetVcsInfo returns the VcsInfo field value
func (o *Project) GetVcsInfo() ProjectVcsInfo {
	if o == nil {
		var ret ProjectVcsInfo
		return ret
	}

	return o.VcsInfo
}

// GetVcsInfoOk returns a tuple with the VcsInfo field value
// and a boolean to check if the value has been set.
func (o *Project) GetVcsInfoOk() (*ProjectVcsInfo, bool) {
	if o == nil {
		return nil, false
	}
	return &o.VcsInfo, true
}

// SetVcsInfo sets field value
func (o *Project) SetVcsInfo(v ProjectVcsInfo) {
	o.VcsInfo = v
}

func (o Project) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["slug"] = o.Slug
	}
	if true {
		toSerialize["name"] = o.Name
	}
	if true {
		toSerialize["organization_name"] = o.OrganizationName
	}
	if true {
		toSerialize["vcs_info"] = o.VcsInfo
	}
	return json.Marshal(toSerialize)
}

type NullableProject struct {
	value *Project
	isSet bool
}

func (v NullableProject) Get() *Project {
	return v.value
}

func (v *NullableProject) Set(val *Project) {
	v.value = val
	v.isSet = true
}

func (v NullableProject) IsSet() bool {
	return v.isSet
}

func (v *NullableProject) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableProject(val *Project) *NullableProject {
	return &NullableProject{value: val, isSet: true}
}

func (v NullableProject) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableProject) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
