// Copyright 2020 The ccictl Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package v2

import (
	json "github.com/goccy/go-json"
)

// JobDetailsPipeline Info about a pipeline the job is a part of.
type JobDetailsPipeline struct {
	// The unique ID of the pipeline.
	Id string `json:"id"`
}

// NewJobDetailsPipeline instantiates a new JobDetailsPipeline object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewJobDetailsPipeline(id string) *JobDetailsPipeline {
	this := JobDetailsPipeline{}
	this.Id = id
	return &this
}

// NewJobDetailsPipelineWithDefaults instantiates a new JobDetailsPipeline object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewJobDetailsPipelineWithDefaults() *JobDetailsPipeline {
	this := JobDetailsPipeline{}
	return &this
}

// GetId returns the Id field value
func (o *JobDetailsPipeline) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *JobDetailsPipeline) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *JobDetailsPipeline) SetId(v string) {
	o.Id = v
}

func (o JobDetailsPipeline) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["id"] = o.Id
	}
	return json.Marshal(toSerialize)
}

type NullableJobDetailsPipeline struct {
	value *JobDetailsPipeline
	isSet bool
}

func (v NullableJobDetailsPipeline) Get() *JobDetailsPipeline {
	return v.value
}

func (v *NullableJobDetailsPipeline) Set(val *JobDetailsPipeline) {
	v.value = val
	v.isSet = true
}

func (v NullableJobDetailsPipeline) IsSet() bool {
	return v.isSet
}

func (v *NullableJobDetailsPipeline) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableJobDetailsPipeline(val *JobDetailsPipeline) *NullableJobDetailsPipeline {
	return &NullableJobDetailsPipeline{value: val, isSet: true}
}

func (v NullableJobDetailsPipeline) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableJobDetailsPipeline) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
