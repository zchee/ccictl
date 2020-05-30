// Copyright 2020 The ccictl Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package v1

import (
	json "github.com/goccy/go-json"
)

// Tests struct for Tests
type Tests struct {
	Tests *[]TestsTests `json:"tests,omitempty"`
}

// NewTests instantiates a new Tests object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTests() *Tests {
	this := Tests{}
	return &this
}

// NewTestsWithDefaults instantiates a new Tests object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTestsWithDefaults() *Tests {
	this := Tests{}
	return &this
}

// GetTests returns the Tests field value if set, zero value otherwise.
func (o *Tests) GetTests() []TestsTests {
	if o == nil || o.Tests == nil {
		var ret []TestsTests
		return ret
	}
	return *o.Tests
}

// GetTestsOk returns a tuple with the Tests field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Tests) GetTestsOk() (*[]TestsTests, bool) {
	if o == nil || o.Tests == nil {
		return nil, false
	}
	return o.Tests, true
}

// HasTests returns a boolean if a field has been set.
func (o *Tests) HasTests() bool {
	if o != nil && o.Tests != nil {
		return true
	}

	return false
}

// SetTests gets a reference to the given []TestsTests and assigns it to the Tests field.
func (o *Tests) SetTests(v []TestsTests) {
	o.Tests = &v
}

func (o Tests) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Tests != nil {
		toSerialize["tests"] = o.Tests
	}
	return json.Marshal(toSerialize)
}

type NullableTests struct {
	value *Tests
	isSet bool
}

func (v NullableTests) Get() *Tests {
	return v.value
}

func (v *NullableTests) Set(val *Tests) {
	v.value = val
	v.isSet = true
}

func (v NullableTests) IsSet() bool {
	return v.isSet
}

func (v *NullableTests) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTests(val *Tests) *NullableTests {
	return &NullableTests{value: val, isSet: true}
}

func (v NullableTests) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTests) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
