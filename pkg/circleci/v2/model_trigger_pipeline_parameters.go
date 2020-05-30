// Copyright 2020 The ccictl Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package v2

import (
	json "github.com/goccy/go-json"
)

// TriggerPipelineParameters The information you can supply when triggering a pipeline.
type TriggerPipelineParameters struct {
	// The branch where the pipeline ran. The HEAD commit on this branch was used for the pipeline. Note that `branch` and `tag` are mutually exclusive.
	Branch *string `json:"branch,omitempty"`
	// The tag used by the pipeline. The commit that this tag points to was used for the pipeline. Note that `branch` and `tag` are mutually exclusive.
	Tag *string `json:"tag,omitempty"`
	// An object containing pipeline parameters and their values.
	Parameters *map[string]interface{} `json:"parameters,omitempty"`
}

// NewTriggerPipelineParameters instantiates a new TriggerPipelineParameters object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTriggerPipelineParameters() *TriggerPipelineParameters {
	this := TriggerPipelineParameters{}
	return &this
}

// NewTriggerPipelineParametersWithDefaults instantiates a new TriggerPipelineParameters object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTriggerPipelineParametersWithDefaults() *TriggerPipelineParameters {
	this := TriggerPipelineParameters{}
	return &this
}

// GetBranch returns the Branch field value if set, zero value otherwise.
func (o *TriggerPipelineParameters) GetBranch() string {
	if o == nil || o.Branch == nil {
		var ret string
		return ret
	}
	return *o.Branch
}

// GetBranchOk returns a tuple with the Branch field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TriggerPipelineParameters) GetBranchOk() (*string, bool) {
	if o == nil || o.Branch == nil {
		return nil, false
	}
	return o.Branch, true
}

// HasBranch returns a boolean if a field has been set.
func (o *TriggerPipelineParameters) HasBranch() bool {
	if o != nil && o.Branch != nil {
		return true
	}

	return false
}

// SetBranch gets a reference to the given string and assigns it to the Branch field.
func (o *TriggerPipelineParameters) SetBranch(v string) {
	o.Branch = &v
}

// GetTag returns the Tag field value if set, zero value otherwise.
func (o *TriggerPipelineParameters) GetTag() string {
	if o == nil || o.Tag == nil {
		var ret string
		return ret
	}
	return *o.Tag
}

// GetTagOk returns a tuple with the Tag field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TriggerPipelineParameters) GetTagOk() (*string, bool) {
	if o == nil || o.Tag == nil {
		return nil, false
	}
	return o.Tag, true
}

// HasTag returns a boolean if a field has been set.
func (o *TriggerPipelineParameters) HasTag() bool {
	if o != nil && o.Tag != nil {
		return true
	}

	return false
}

// SetTag gets a reference to the given string and assigns it to the Tag field.
func (o *TriggerPipelineParameters) SetTag(v string) {
	o.Tag = &v
}

// GetParameters returns the Parameters field value if set, zero value otherwise.
func (o *TriggerPipelineParameters) GetParameters() map[string]interface{} {
	if o == nil || o.Parameters == nil {
		var ret map[string]interface{}
		return ret
	}
	return *o.Parameters
}

// GetParametersOk returns a tuple with the Parameters field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TriggerPipelineParameters) GetParametersOk() (*map[string]interface{}, bool) {
	if o == nil || o.Parameters == nil {
		return nil, false
	}
	return o.Parameters, true
}

// HasParameters returns a boolean if a field has been set.
func (o *TriggerPipelineParameters) HasParameters() bool {
	if o != nil && o.Parameters != nil {
		return true
	}

	return false
}

// SetParameters gets a reference to the given map[string]interface{} and assigns it to the Parameters field.
func (o *TriggerPipelineParameters) SetParameters(v map[string]interface{}) {
	o.Parameters = &v
}

func (o TriggerPipelineParameters) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Branch != nil {
		toSerialize["branch"] = o.Branch
	}
	if o.Tag != nil {
		toSerialize["tag"] = o.Tag
	}
	if o.Parameters != nil {
		toSerialize["parameters"] = o.Parameters
	}
	return json.Marshal(toSerialize)
}

type NullableTriggerPipelineParameters struct {
	value *TriggerPipelineParameters
	isSet bool
}

func (v NullableTriggerPipelineParameters) Get() *TriggerPipelineParameters {
	return v.value
}

func (v *NullableTriggerPipelineParameters) Set(val *TriggerPipelineParameters) {
	v.value = val
	v.isSet = true
}

func (v NullableTriggerPipelineParameters) IsSet() bool {
	return v.isSet
}

func (v *NullableTriggerPipelineParameters) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTriggerPipelineParameters(val *TriggerPipelineParameters) *NullableTriggerPipelineParameters {
	return &NullableTriggerPipelineParameters{value: val, isSet: true}
}

func (v NullableTriggerPipelineParameters) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTriggerPipelineParameters) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
