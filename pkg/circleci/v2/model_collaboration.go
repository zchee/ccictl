// Copyright 2020 The ccictl Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package v2

import (
	json "github.com/goccy/go-json"
)

// Collaboration struct for Collaboration
type Collaboration struct {
	// The VCS provider
	VcsType string `json:"vcs-type"`
	// The name of the organization
	Name string `json:"name"`
	// URL to the user's avatar on the VCS
	AvatarUrl string `json:"avatar_url"`
}

// NewCollaboration instantiates a new Collaboration object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCollaboration(vcsType string, name string, avatarUrl string) *Collaboration {
	this := Collaboration{}
	this.VcsType = vcsType
	this.Name = name
	this.AvatarUrl = avatarUrl
	return &this
}

// NewCollaborationWithDefaults instantiates a new Collaboration object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCollaborationWithDefaults() *Collaboration {
	this := Collaboration{}
	return &this
}

// GetVcsType returns the VcsType field value
func (o *Collaboration) GetVcsType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.VcsType
}

// GetVcsTypeOk returns a tuple with the VcsType field value
// and a boolean to check if the value has been set.
func (o *Collaboration) GetVcsTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.VcsType, true
}

// SetVcsType sets field value
func (o *Collaboration) SetVcsType(v string) {
	o.VcsType = v
}

// GetName returns the Name field value
func (o *Collaboration) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *Collaboration) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *Collaboration) SetName(v string) {
	o.Name = v
}

// GetAvatarUrl returns the AvatarUrl field value
func (o *Collaboration) GetAvatarUrl() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AvatarUrl
}

// GetAvatarUrlOk returns a tuple with the AvatarUrl field value
// and a boolean to check if the value has been set.
func (o *Collaboration) GetAvatarUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AvatarUrl, true
}

// SetAvatarUrl sets field value
func (o *Collaboration) SetAvatarUrl(v string) {
	o.AvatarUrl = v
}

func (o Collaboration) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["vcs-type"] = o.VcsType
	}
	if true {
		toSerialize["name"] = o.Name
	}
	if true {
		toSerialize["avatar_url"] = o.AvatarUrl
	}
	return json.Marshal(toSerialize)
}

type NullableCollaboration struct {
	value *Collaboration
	isSet bool
}

func (v NullableCollaboration) Get() *Collaboration {
	return v.value
}

func (v *NullableCollaboration) Set(val *Collaboration) {
	v.value = val
	v.isSet = true
}

func (v NullableCollaboration) IsSet() bool {
	return v.isSet
}

func (v *NullableCollaboration) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCollaboration(val *Collaboration) *NullableCollaboration {
	return &NullableCollaboration{value: val, isSet: true}
}

func (v NullableCollaboration) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCollaboration) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
