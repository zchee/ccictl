// Copyright 2020 The ccictl Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package v2

import (
	json "github.com/goccy/go-json"
)

// PipelineVcsCommit The latest commit in the pipeline.
type PipelineVcsCommit struct {
	// The subject of the commit message.
	Subject string `json:"subject"`
	// The body of the commit message.
	Body string `json:"body"`
}

// NewPipelineVcsCommit instantiates a new PipelineVcsCommit object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPipelineVcsCommit(subject string, body string) *PipelineVcsCommit {
	this := PipelineVcsCommit{}
	this.Subject = subject
	this.Body = body
	return &this
}

// NewPipelineVcsCommitWithDefaults instantiates a new PipelineVcsCommit object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPipelineVcsCommitWithDefaults() *PipelineVcsCommit {
	this := PipelineVcsCommit{}
	return &this
}

// GetSubject returns the Subject field value
func (o *PipelineVcsCommit) GetSubject() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Subject
}

// GetSubjectOk returns a tuple with the Subject field value
// and a boolean to check if the value has been set.
func (o *PipelineVcsCommit) GetSubjectOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Subject, true
}

// SetSubject sets field value
func (o *PipelineVcsCommit) SetSubject(v string) {
	o.Subject = v
}

// GetBody returns the Body field value
func (o *PipelineVcsCommit) GetBody() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Body
}

// GetBodyOk returns a tuple with the Body field value
// and a boolean to check if the value has been set.
func (o *PipelineVcsCommit) GetBodyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Body, true
}

// SetBody sets field value
func (o *PipelineVcsCommit) SetBody(v string) {
	o.Body = v
}

func (o PipelineVcsCommit) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["subject"] = o.Subject
	}
	if true {
		toSerialize["body"] = o.Body
	}
	return json.Marshal(toSerialize)
}

type NullablePipelineVcsCommit struct {
	value *PipelineVcsCommit
	isSet bool
}

func (v NullablePipelineVcsCommit) Get() *PipelineVcsCommit {
	return v.value
}

func (v *NullablePipelineVcsCommit) Set(val *PipelineVcsCommit) {
	v.value = val
	v.isSet = true
}

func (v NullablePipelineVcsCommit) IsSet() bool {
	return v.isSet
}

func (v *NullablePipelineVcsCommit) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePipelineVcsCommit(val *PipelineVcsCommit) *NullablePipelineVcsCommit {
	return &NullablePipelineVcsCommit{value: val, isSet: true}
}

func (v NullablePipelineVcsCommit) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePipelineVcsCommit) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
