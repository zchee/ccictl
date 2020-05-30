// Copyright 2020 The ccictl Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package v1

import (
	json "github.com/goccy/go-json"
)

// ProjectFeatureFlags struct for ProjectFeatureFlags
type ProjectFeatureFlags struct {
	BuildForkPrs    *bool `json:"build-fork-prs,omitempty"`
	Junit           *bool `json:"junit,omitempty"`
	Oss             *bool `json:"oss,omitempty"`
	Osx             *bool `json:"osx,omitempty"`
	SetGithubStatus *bool `json:"set-github-status,omitempty"`
	TrustyBeta      *bool `json:"trusty-beta,omitempty"`
}

// NewProjectFeatureFlags instantiates a new ProjectFeatureFlags object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewProjectFeatureFlags() *ProjectFeatureFlags {
	this := ProjectFeatureFlags{}
	return &this
}

// NewProjectFeatureFlagsWithDefaults instantiates a new ProjectFeatureFlags object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewProjectFeatureFlagsWithDefaults() *ProjectFeatureFlags {
	this := ProjectFeatureFlags{}
	return &this
}

// GetBuildForkPrs returns the BuildForkPrs field value if set, zero value otherwise.
func (o *ProjectFeatureFlags) GetBuildForkPrs() bool {
	if o == nil || o.BuildForkPrs == nil {
		var ret bool
		return ret
	}
	return *o.BuildForkPrs
}

// GetBuildForkPrsOk returns a tuple with the BuildForkPrs field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProjectFeatureFlags) GetBuildForkPrsOk() (*bool, bool) {
	if o == nil || o.BuildForkPrs == nil {
		return nil, false
	}
	return o.BuildForkPrs, true
}

// HasBuildForkPrs returns a boolean if a field has been set.
func (o *ProjectFeatureFlags) HasBuildForkPrs() bool {
	if o != nil && o.BuildForkPrs != nil {
		return true
	}

	return false
}

// SetBuildForkPrs gets a reference to the given bool and assigns it to the BuildForkPrs field.
func (o *ProjectFeatureFlags) SetBuildForkPrs(v bool) {
	o.BuildForkPrs = &v
}

// GetJunit returns the Junit field value if set, zero value otherwise.
func (o *ProjectFeatureFlags) GetJunit() bool {
	if o == nil || o.Junit == nil {
		var ret bool
		return ret
	}
	return *o.Junit
}

// GetJunitOk returns a tuple with the Junit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProjectFeatureFlags) GetJunitOk() (*bool, bool) {
	if o == nil || o.Junit == nil {
		return nil, false
	}
	return o.Junit, true
}

// HasJunit returns a boolean if a field has been set.
func (o *ProjectFeatureFlags) HasJunit() bool {
	if o != nil && o.Junit != nil {
		return true
	}

	return false
}

// SetJunit gets a reference to the given bool and assigns it to the Junit field.
func (o *ProjectFeatureFlags) SetJunit(v bool) {
	o.Junit = &v
}

// GetOss returns the Oss field value if set, zero value otherwise.
func (o *ProjectFeatureFlags) GetOss() bool {
	if o == nil || o.Oss == nil {
		var ret bool
		return ret
	}
	return *o.Oss
}

// GetOssOk returns a tuple with the Oss field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProjectFeatureFlags) GetOssOk() (*bool, bool) {
	if o == nil || o.Oss == nil {
		return nil, false
	}
	return o.Oss, true
}

// HasOss returns a boolean if a field has been set.
func (o *ProjectFeatureFlags) HasOss() bool {
	if o != nil && o.Oss != nil {
		return true
	}

	return false
}

// SetOss gets a reference to the given bool and assigns it to the Oss field.
func (o *ProjectFeatureFlags) SetOss(v bool) {
	o.Oss = &v
}

// GetOsx returns the Osx field value if set, zero value otherwise.
func (o *ProjectFeatureFlags) GetOsx() bool {
	if o == nil || o.Osx == nil {
		var ret bool
		return ret
	}
	return *o.Osx
}

// GetOsxOk returns a tuple with the Osx field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProjectFeatureFlags) GetOsxOk() (*bool, bool) {
	if o == nil || o.Osx == nil {
		return nil, false
	}
	return o.Osx, true
}

// HasOsx returns a boolean if a field has been set.
func (o *ProjectFeatureFlags) HasOsx() bool {
	if o != nil && o.Osx != nil {
		return true
	}

	return false
}

// SetOsx gets a reference to the given bool and assigns it to the Osx field.
func (o *ProjectFeatureFlags) SetOsx(v bool) {
	o.Osx = &v
}

// GetSetGithubStatus returns the SetGithubStatus field value if set, zero value otherwise.
func (o *ProjectFeatureFlags) GetSetGithubStatus() bool {
	if o == nil || o.SetGithubStatus == nil {
		var ret bool
		return ret
	}
	return *o.SetGithubStatus
}

// GetSetGithubStatusOk returns a tuple with the SetGithubStatus field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProjectFeatureFlags) GetSetGithubStatusOk() (*bool, bool) {
	if o == nil || o.SetGithubStatus == nil {
		return nil, false
	}
	return o.SetGithubStatus, true
}

// HasSetGithubStatus returns a boolean if a field has been set.
func (o *ProjectFeatureFlags) HasSetGithubStatus() bool {
	if o != nil && o.SetGithubStatus != nil {
		return true
	}

	return false
}

// SetSetGithubStatus gets a reference to the given bool and assigns it to the SetGithubStatus field.
func (o *ProjectFeatureFlags) SetSetGithubStatus(v bool) {
	o.SetGithubStatus = &v
}

// GetTrustyBeta returns the TrustyBeta field value if set, zero value otherwise.
func (o *ProjectFeatureFlags) GetTrustyBeta() bool {
	if o == nil || o.TrustyBeta == nil {
		var ret bool
		return ret
	}
	return *o.TrustyBeta
}

// GetTrustyBetaOk returns a tuple with the TrustyBeta field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *ProjectFeatureFlags) GetTrustyBetaOk() (*bool, bool) {
	if o == nil || o.TrustyBeta == nil {
		return nil, false
	}
	return o.TrustyBeta, true
}

// HasTrustyBeta returns a boolean if a field has been set.
func (o *ProjectFeatureFlags) HasTrustyBeta() bool {
	if o != nil && o.TrustyBeta != nil {
		return true
	}

	return false
}

// SetTrustyBeta gets a reference to the given bool and assigns it to the TrustyBeta field.
func (o *ProjectFeatureFlags) SetTrustyBeta(v bool) {
	o.TrustyBeta = &v
}

func (o ProjectFeatureFlags) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.BuildForkPrs != nil {
		toSerialize["build-fork-prs"] = o.BuildForkPrs
	}
	if o.Junit != nil {
		toSerialize["junit"] = o.Junit
	}
	if o.Oss != nil {
		toSerialize["oss"] = o.Oss
	}
	if o.Osx != nil {
		toSerialize["osx"] = o.Osx
	}
	if o.SetGithubStatus != nil {
		toSerialize["set-github-status"] = o.SetGithubStatus
	}
	if o.TrustyBeta != nil {
		toSerialize["trusty-beta"] = o.TrustyBeta
	}
	return json.Marshal(toSerialize)
}

type NullableProjectFeatureFlags struct {
	value *ProjectFeatureFlags
	isSet bool
}

func (v NullableProjectFeatureFlags) Get() *ProjectFeatureFlags {
	return v.value
}

func (v *NullableProjectFeatureFlags) Set(val *ProjectFeatureFlags) {
	v.value = val
	v.isSet = true
}

func (v NullableProjectFeatureFlags) IsSet() bool {
	return v.isSet
}

func (v *NullableProjectFeatureFlags) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableProjectFeatureFlags(val *ProjectFeatureFlags) *NullableProjectFeatureFlags {
	return &NullableProjectFeatureFlags{value: val, isSet: true}
}

func (v NullableProjectFeatureFlags) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableProjectFeatureFlags) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
