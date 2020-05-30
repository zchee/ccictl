// Copyright 2020 The ccictl Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package v2

import (
	json "github.com/goccy/go-json"
)

// PipelineTriggerActor The user who triggered the Pipeline.
type PipelineTriggerActor struct {
	// The login information for the user on the VCS.
	Login string `json:"login"`
	// URL to the user's avatar on the VCS
	AvatarUrl string `json:"avatar_url"`
}

// NewPipelineTriggerActor instantiates a new PipelineTriggerActor object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPipelineTriggerActor(login string, avatarUrl string) *PipelineTriggerActor {
	this := PipelineTriggerActor{}
	this.Login = login
	this.AvatarUrl = avatarUrl
	return &this
}

// NewPipelineTriggerActorWithDefaults instantiates a new PipelineTriggerActor object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPipelineTriggerActorWithDefaults() *PipelineTriggerActor {
	this := PipelineTriggerActor{}
	return &this
}

// GetLogin returns the Login field value
func (o *PipelineTriggerActor) GetLogin() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Login
}

// GetLoginOk returns a tuple with the Login field value
// and a boolean to check if the value has been set.
func (o *PipelineTriggerActor) GetLoginOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Login, true
}

// SetLogin sets field value
func (o *PipelineTriggerActor) SetLogin(v string) {
	o.Login = v
}

// GetAvatarUrl returns the AvatarUrl field value
func (o *PipelineTriggerActor) GetAvatarUrl() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.AvatarUrl
}

// GetAvatarUrlOk returns a tuple with the AvatarUrl field value
// and a boolean to check if the value has been set.
func (o *PipelineTriggerActor) GetAvatarUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.AvatarUrl, true
}

// SetAvatarUrl sets field value
func (o *PipelineTriggerActor) SetAvatarUrl(v string) {
	o.AvatarUrl = v
}

func (o PipelineTriggerActor) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["login"] = o.Login
	}
	if true {
		toSerialize["avatar_url"] = o.AvatarUrl
	}
	return json.Marshal(toSerialize)
}

type NullablePipelineTriggerActor struct {
	value *PipelineTriggerActor
	isSet bool
}

func (v NullablePipelineTriggerActor) Get() *PipelineTriggerActor {
	return v.value
}

func (v *NullablePipelineTriggerActor) Set(val *PipelineTriggerActor) {
	v.value = val
	v.isSet = true
}

func (v NullablePipelineTriggerActor) IsSet() bool {
	return v.isSet
}

func (v *NullablePipelineTriggerActor) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePipelineTriggerActor(val *PipelineTriggerActor) *NullablePipelineTriggerActor {
	return &NullablePipelineTriggerActor{value: val, isSet: true}
}

func (v NullablePipelineTriggerActor) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePipelineTriggerActor) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
