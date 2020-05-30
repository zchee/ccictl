// Copyright 2020 The ccictl Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package v2

import (
	json "github.com/goccy/go-json"
)

// PipelineConfig The configuration strings for the pipeline.
type PipelineConfig struct {
	// The source configuration for the pipeline, before any config compilation has been performed. If there is no config, then this field will be empty.
	Source string `json:"source"`
	// The compiled configuration for the pipeline, after all orb expansion has been performed. If there were errors processing the pipeline's configuration, then this field may be empty.
	Compiled string `json:"compiled"`
}

// NewPipelineConfig instantiates a new PipelineConfig object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPipelineConfig(source string, compiled string) *PipelineConfig {
	this := PipelineConfig{}
	this.Source = source
	this.Compiled = compiled
	return &this
}

// NewPipelineConfigWithDefaults instantiates a new PipelineConfig object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPipelineConfigWithDefaults() *PipelineConfig {
	this := PipelineConfig{}
	return &this
}

// GetSource returns the Source field value
func (o *PipelineConfig) GetSource() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Source
}

// GetSourceOk returns a tuple with the Source field value
// and a boolean to check if the value has been set.
func (o *PipelineConfig) GetSourceOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Source, true
}

// SetSource sets field value
func (o *PipelineConfig) SetSource(v string) {
	o.Source = v
}

// GetCompiled returns the Compiled field value
func (o *PipelineConfig) GetCompiled() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Compiled
}

// GetCompiledOk returns a tuple with the Compiled field value
// and a boolean to check if the value has been set.
func (o *PipelineConfig) GetCompiledOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Compiled, true
}

// SetCompiled sets field value
func (o *PipelineConfig) SetCompiled(v string) {
	o.Compiled = v
}

func (o PipelineConfig) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["source"] = o.Source
	}
	if true {
		toSerialize["compiled"] = o.Compiled
	}
	return json.Marshal(toSerialize)
}

type NullablePipelineConfig struct {
	value *PipelineConfig
	isSet bool
}

func (v NullablePipelineConfig) Get() *PipelineConfig {
	return v.value
}

func (v *NullablePipelineConfig) Set(val *PipelineConfig) {
	v.value = val
	v.isSet = true
}

func (v NullablePipelineConfig) IsSet() bool {
	return v.isSet
}

func (v *NullablePipelineConfig) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePipelineConfig(val *PipelineConfig) *NullablePipelineConfig {
	return &NullablePipelineConfig{value: val, isSet: true}
}

func (v NullablePipelineConfig) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePipelineConfig) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
