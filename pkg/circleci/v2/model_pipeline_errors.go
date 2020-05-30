// Copyright 2020 The ccictl Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package v2

import (
	json "github.com/goccy/go-json"
)

// PipelineErrors An error with a type and message.
type PipelineErrors struct {
	// The type of error.
	Type string `json:"type"`
	// A human-readable error message.
	Message string `json:"message"`
}

// NewPipelineErrors instantiates a new PipelineErrors object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPipelineErrors(type_ string, message string) *PipelineErrors {
	this := PipelineErrors{}
	this.Type = type_
	this.Message = message
	return &this
}

// NewPipelineErrorsWithDefaults instantiates a new PipelineErrors object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPipelineErrorsWithDefaults() *PipelineErrors {
	this := PipelineErrors{}
	return &this
}

// GetType returns the Type field value
func (o *PipelineErrors) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *PipelineErrors) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *PipelineErrors) SetType(v string) {
	o.Type = v
}

// GetMessage returns the Message field value
func (o *PipelineErrors) GetMessage() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Message
}

// GetMessageOk returns a tuple with the Message field value
// and a boolean to check if the value has been set.
func (o *PipelineErrors) GetMessageOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Message, true
}

// SetMessage sets field value
func (o *PipelineErrors) SetMessage(v string) {
	o.Message = v
}

func (o PipelineErrors) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["type"] = o.Type
	}
	if true {
		toSerialize["message"] = o.Message
	}
	return json.Marshal(toSerialize)
}

type NullablePipelineErrors struct {
	value *PipelineErrors
	isSet bool
}

func (v NullablePipelineErrors) Get() *PipelineErrors {
	return v.value
}

func (v *NullablePipelineErrors) Set(val *PipelineErrors) {
	v.value = val
	v.isSet = true
}

func (v NullablePipelineErrors) IsSet() bool {
	return v.isSet
}

func (v *NullablePipelineErrors) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePipelineErrors(val *PipelineErrors) *NullablePipelineErrors {
	return &NullablePipelineErrors{value: val, isSet: true}
}

func (v NullablePipelineErrors) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePipelineErrors) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
