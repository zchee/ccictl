// Copyright 2020 The ccictl Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package v2

import (
	json "github.com/goccy/go-json"
)

// EnvironmentVariablePair struct for EnvironmentVariablePair
type EnvironmentVariablePair struct {
	// The name of the environment variable.
	Name string `json:"name"`
	// The value of the environment variable.
	Value string `json:"value"`
}

// NewEnvironmentVariablePair instantiates a new EnvironmentVariablePair object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEnvironmentVariablePair(name string, value string) *EnvironmentVariablePair {
	this := EnvironmentVariablePair{}
	this.Name = name
	this.Value = value
	return &this
}

// NewEnvironmentVariablePairWithDefaults instantiates a new EnvironmentVariablePair object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEnvironmentVariablePairWithDefaults() *EnvironmentVariablePair {
	this := EnvironmentVariablePair{}
	return &this
}

// GetName returns the Name field value
func (o *EnvironmentVariablePair) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *EnvironmentVariablePair) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *EnvironmentVariablePair) SetName(v string) {
	o.Name = v
}

// GetValue returns the Value field value
func (o *EnvironmentVariablePair) GetValue() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Value
}

// GetValueOk returns a tuple with the Value field value
// and a boolean to check if the value has been set.
func (o *EnvironmentVariablePair) GetValueOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Value, true
}

// SetValue sets field value
func (o *EnvironmentVariablePair) SetValue(v string) {
	o.Value = v
}

func (o EnvironmentVariablePair) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["name"] = o.Name
	}
	if true {
		toSerialize["value"] = o.Value
	}
	return json.Marshal(toSerialize)
}

type NullableEnvironmentVariablePair struct {
	value *EnvironmentVariablePair
	isSet bool
}

func (v NullableEnvironmentVariablePair) Get() *EnvironmentVariablePair {
	return v.value
}

func (v *NullableEnvironmentVariablePair) Set(val *EnvironmentVariablePair) {
	v.value = val
	v.isSet = true
}

func (v NullableEnvironmentVariablePair) IsSet() bool {
	return v.isSet
}

func (v *NullableEnvironmentVariablePair) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEnvironmentVariablePair(val *EnvironmentVariablePair) *NullableEnvironmentVariablePair {
	return &NullableEnvironmentVariablePair{value: val, isSet: true}
}

func (v NullableEnvironmentVariablePair) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEnvironmentVariablePair) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
