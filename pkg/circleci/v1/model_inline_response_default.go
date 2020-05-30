// Copyright 2020 The ccictl Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package v1

import (
	json "github.com/goccy/go-json"
)

// InlineResponseDefault struct for InlineResponseDefault
type InlineResponseDefault struct {
	Message *string `json:"message,omitempty"`
}

// NewInlineResponseDefault instantiates a new InlineResponseDefault object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineResponseDefault() *InlineResponseDefault {
	this := InlineResponseDefault{}
	return &this
}

// NewInlineResponseDefaultWithDefaults instantiates a new InlineResponseDefault object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineResponseDefaultWithDefaults() *InlineResponseDefault {
	this := InlineResponseDefault{}
	return &this
}

// GetMessage returns the Message field value if set, zero value otherwise.
func (o *InlineResponseDefault) GetMessage() string {
	if o == nil || o.Message == nil {
		var ret string
		return ret
	}
	return *o.Message
}

// GetMessageOk returns a tuple with the Message field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineResponseDefault) GetMessageOk() (*string, bool) {
	if o == nil || o.Message == nil {
		return nil, false
	}
	return o.Message, true
}

// HasMessage returns a boolean if a field has been set.
func (o *InlineResponseDefault) HasMessage() bool {
	if o != nil && o.Message != nil {
		return true
	}

	return false
}

// SetMessage gets a reference to the given string and assigns it to the Message field.
func (o *InlineResponseDefault) SetMessage(v string) {
	o.Message = &v
}

func (o InlineResponseDefault) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Message != nil {
		toSerialize["message"] = o.Message
	}
	return json.Marshal(toSerialize)
}

type NullableInlineResponseDefault struct {
	value *InlineResponseDefault
	isSet bool
}

func (v NullableInlineResponseDefault) Get() *InlineResponseDefault {
	return v.value
}

func (v *NullableInlineResponseDefault) Set(val *InlineResponseDefault) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineResponseDefault) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineResponseDefault) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineResponseDefault(val *InlineResponseDefault) *NullableInlineResponseDefault {
	return &NullableInlineResponseDefault{value: val, isSet: true}
}

func (v NullableInlineResponseDefault) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineResponseDefault) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
