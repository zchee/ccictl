// Copyright 2020 The ccictl Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package v1

import (
	json "github.com/goccy/go-json"
)

// Envvar struct for Envvar
type Envvar struct {
	Name  *string `json:"name,omitempty"`
	Value *string `json:"value,omitempty"`
}

// NewEnvvar instantiates a new Envvar object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEnvvar() *Envvar {
	this := Envvar{}
	return &this
}

// NewEnvvarWithDefaults instantiates a new Envvar object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEnvvarWithDefaults() *Envvar {
	this := Envvar{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *Envvar) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Envvar) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *Envvar) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *Envvar) SetName(v string) {
	o.Name = &v
}

// GetValue returns the Value field value if set, zero value otherwise.
func (o *Envvar) GetValue() string {
	if o == nil || o.Value == nil {
		var ret string
		return ret
	}
	return *o.Value
}

// GetValueOk returns a tuple with the Value field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Envvar) GetValueOk() (*string, bool) {
	if o == nil || o.Value == nil {
		return nil, false
	}
	return o.Value, true
}

// HasValue returns a boolean if a field has been set.
func (o *Envvar) HasValue() bool {
	if o != nil && o.Value != nil {
		return true
	}

	return false
}

// SetValue gets a reference to the given string and assigns it to the Value field.
func (o *Envvar) SetValue(v string) {
	o.Value = &v
}

func (o Envvar) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.Value != nil {
		toSerialize["value"] = o.Value
	}
	return json.Marshal(toSerialize)
}

type NullableEnvvar struct {
	value *Envvar
	isSet bool
}

func (v NullableEnvvar) Get() *Envvar {
	return v.value
}

func (v *NullableEnvvar) Set(val *Envvar) {
	v.value = val
	v.isSet = true
}

func (v NullableEnvvar) IsSet() bool {
	return v.isSet
}

func (v *NullableEnvvar) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEnvvar(val *Envvar) *NullableEnvvar {
	return &NullableEnvvar{value: val, isSet: true}
}

func (v NullableEnvvar) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEnvvar) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
