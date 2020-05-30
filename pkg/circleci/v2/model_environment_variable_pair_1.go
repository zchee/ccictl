// Copyright 2020 The ccictl Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package v2

import (
	json "github.com/goccy/go-json"
)

// EnvironmentVariablePair1 struct for EnvironmentVariablePair1
type EnvironmentVariablePair1 struct {
	// The name of the environment variable.
	Name string `json:"name"`
	// The value of the environment variable.
	Value string `json:"value"`
}

// NewEnvironmentVariablePair1 instantiates a new EnvironmentVariablePair1 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEnvironmentVariablePair1(name string, value string) *EnvironmentVariablePair1 {
	this := EnvironmentVariablePair1{}
	this.Name = name
	this.Value = value
	return &this
}

// NewEnvironmentVariablePair1WithDefaults instantiates a new EnvironmentVariablePair1 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEnvironmentVariablePair1WithDefaults() *EnvironmentVariablePair1 {
	this := EnvironmentVariablePair1{}
	return &this
}

// GetName returns the Name field value
func (o *EnvironmentVariablePair1) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *EnvironmentVariablePair1) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *EnvironmentVariablePair1) SetName(v string) {
	o.Name = v
}

// GetValue returns the Value field value
func (o *EnvironmentVariablePair1) GetValue() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Value
}

// GetValueOk returns a tuple with the Value field value
// and a boolean to check if the value has been set.
func (o *EnvironmentVariablePair1) GetValueOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Value, true
}

// SetValue sets field value
func (o *EnvironmentVariablePair1) SetValue(v string) {
	o.Value = v
}

func (o EnvironmentVariablePair1) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["name"] = o.Name
	}
	if true {
		toSerialize["value"] = o.Value
	}
	return json.Marshal(toSerialize)
}

type NullableEnvironmentVariablePair1 struct {
	value *EnvironmentVariablePair1
	isSet bool
}

func (v NullableEnvironmentVariablePair1) Get() *EnvironmentVariablePair1 {
	return v.value
}

func (v *NullableEnvironmentVariablePair1) Set(val *EnvironmentVariablePair1) {
	v.value = val
	v.isSet = true
}

func (v NullableEnvironmentVariablePair1) IsSet() bool {
	return v.isSet
}

func (v *NullableEnvironmentVariablePair1) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEnvironmentVariablePair1(val *EnvironmentVariablePair1) *NullableEnvironmentVariablePair1 {
	return &NullableEnvironmentVariablePair1{value: val, isSet: true}
}

func (v NullableEnvironmentVariablePair1) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEnvironmentVariablePair1) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
