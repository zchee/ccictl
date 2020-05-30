// Copyright 2020 The ccictl Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package v2

import (
	"time"

	json "github.com/goccy/go-json"
)

// PipelineLight NOTE: The definition of pipeline is subject to change.
type PipelineLight struct {
	// The unique ID of the pipeline.
	Id string `json:"id"`
	// The current state of the pipeline.
	State string `json:"state"`
	// The number of the pipeline.
	Number int64 `json:"number"`
	// The date and time the pipeline was created.
	CreatedAt time.Time `json:"created_at"`
}

// NewPipelineLight instantiates a new PipelineLight object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPipelineLight(id string, state string, number int64, createdAt time.Time) *PipelineLight {
	this := PipelineLight{}
	this.Id = id
	this.State = state
	this.Number = number
	this.CreatedAt = createdAt
	return &this
}

// NewPipelineLightWithDefaults instantiates a new PipelineLight object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPipelineLightWithDefaults() *PipelineLight {
	this := PipelineLight{}
	return &this
}

// GetId returns the Id field value
func (o *PipelineLight) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *PipelineLight) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *PipelineLight) SetId(v string) {
	o.Id = v
}

// GetState returns the State field value
func (o *PipelineLight) GetState() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.State
}

// GetStateOk returns a tuple with the State field value
// and a boolean to check if the value has been set.
func (o *PipelineLight) GetStateOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.State, true
}

// SetState sets field value
func (o *PipelineLight) SetState(v string) {
	o.State = v
}

// GetNumber returns the Number field value
func (o *PipelineLight) GetNumber() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.Number
}

// GetNumberOk returns a tuple with the Number field value
// and a boolean to check if the value has been set.
func (o *PipelineLight) GetNumberOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Number, true
}

// SetNumber sets field value
func (o *PipelineLight) SetNumber(v int64) {
	o.Number = v
}

// GetCreatedAt returns the CreatedAt field value
func (o *PipelineLight) GetCreatedAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value
// and a boolean to check if the value has been set.
func (o *PipelineLight) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreatedAt, true
}

// SetCreatedAt sets field value
func (o *PipelineLight) SetCreatedAt(v time.Time) {
	o.CreatedAt = v
}

func (o PipelineLight) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["id"] = o.Id
	}
	if true {
		toSerialize["state"] = o.State
	}
	if true {
		toSerialize["number"] = o.Number
	}
	if true {
		toSerialize["created_at"] = o.CreatedAt
	}
	return json.Marshal(toSerialize)
}

type NullablePipelineLight struct {
	value *PipelineLight
	isSet bool
}

func (v NullablePipelineLight) Get() *PipelineLight {
	return v.value
}

func (v *NullablePipelineLight) Set(val *PipelineLight) {
	v.value = val
	v.isSet = true
}

func (v NullablePipelineLight) IsSet() bool {
	return v.isSet
}

func (v *NullablePipelineLight) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePipelineLight(val *PipelineLight) *NullablePipelineLight {
	return &NullablePipelineLight{value: val, isSet: true}
}

func (v NullablePipelineLight) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePipelineLight) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
