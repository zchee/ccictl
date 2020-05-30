// Copyright 2020 The ccictl Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package v2

import (
	json "github.com/goccy/go-json"
)

// ProjectVcsInfo Information about the VCS that hosts the project source code.
type ProjectVcsInfo struct {
	// URL to the repository hosting the project's code
	VcsUrl string `json:"vcs_url"`
	// The VCS provider
	Provider      string `json:"provider"`
	DefaultBranch string `json:"default_branch"`
}

// NewProjectVcsInfo instantiates a new ProjectVcsInfo object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewProjectVcsInfo(vcsUrl string, provider string, defaultBranch string) *ProjectVcsInfo {
	this := ProjectVcsInfo{}
	this.VcsUrl = vcsUrl
	this.Provider = provider
	this.DefaultBranch = defaultBranch
	return &this
}

// NewProjectVcsInfoWithDefaults instantiates a new ProjectVcsInfo object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewProjectVcsInfoWithDefaults() *ProjectVcsInfo {
	this := ProjectVcsInfo{}
	return &this
}

// GetVcsUrl returns the VcsUrl field value
func (o *ProjectVcsInfo) GetVcsUrl() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.VcsUrl
}

// GetVcsUrlOk returns a tuple with the VcsUrl field value
// and a boolean to check if the value has been set.
func (o *ProjectVcsInfo) GetVcsUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.VcsUrl, true
}

// SetVcsUrl sets field value
func (o *ProjectVcsInfo) SetVcsUrl(v string) {
	o.VcsUrl = v
}

// GetProvider returns the Provider field value
func (o *ProjectVcsInfo) GetProvider() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Provider
}

// GetProviderOk returns a tuple with the Provider field value
// and a boolean to check if the value has been set.
func (o *ProjectVcsInfo) GetProviderOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Provider, true
}

// SetProvider sets field value
func (o *ProjectVcsInfo) SetProvider(v string) {
	o.Provider = v
}

// GetDefaultBranch returns the DefaultBranch field value
func (o *ProjectVcsInfo) GetDefaultBranch() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.DefaultBranch
}

// GetDefaultBranchOk returns a tuple with the DefaultBranch field value
// and a boolean to check if the value has been set.
func (o *ProjectVcsInfo) GetDefaultBranchOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.DefaultBranch, true
}

// SetDefaultBranch sets field value
func (o *ProjectVcsInfo) SetDefaultBranch(v string) {
	o.DefaultBranch = v
}

func (o ProjectVcsInfo) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["vcs_url"] = o.VcsUrl
	}
	if true {
		toSerialize["provider"] = o.Provider
	}
	if true {
		toSerialize["default_branch"] = o.DefaultBranch
	}
	return json.Marshal(toSerialize)
}

type NullableProjectVcsInfo struct {
	value *ProjectVcsInfo
	isSet bool
}

func (v NullableProjectVcsInfo) Get() *ProjectVcsInfo {
	return v.value
}

func (v *NullableProjectVcsInfo) Set(val *ProjectVcsInfo) {
	v.value = val
	v.isSet = true
}

func (v NullableProjectVcsInfo) IsSet() bool {
	return v.isSet
}

func (v *NullableProjectVcsInfo) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableProjectVcsInfo(val *ProjectVcsInfo) *NullableProjectVcsInfo {
	return &NullableProjectVcsInfo{value: val, isSet: true}
}

func (v NullableProjectVcsInfo) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableProjectVcsInfo) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
