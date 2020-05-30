// Copyright 2020 The ccictl Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package v2

import (
	json "github.com/goccy/go-json"
)

// CheckoutKeyInput struct for CheckoutKeyInput
type CheckoutKeyInput struct {
	// The type of checkout key to create. This may be either `deploy-key` or `user-key`.
	Type string `json:"type"`
}

// NewCheckoutKeyInput instantiates a new CheckoutKeyInput object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCheckoutKeyInput(type_ string) *CheckoutKeyInput {
	this := CheckoutKeyInput{}
	this.Type = type_
	return &this
}

// NewCheckoutKeyInputWithDefaults instantiates a new CheckoutKeyInput object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCheckoutKeyInputWithDefaults() *CheckoutKeyInput {
	this := CheckoutKeyInput{}
	return &this
}

// GetType returns the Type field value
func (o *CheckoutKeyInput) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *CheckoutKeyInput) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *CheckoutKeyInput) SetType(v string) {
	o.Type = v
}

func (o CheckoutKeyInput) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["type"] = o.Type
	}
	return json.Marshal(toSerialize)
}

type NullableCheckoutKeyInput struct {
	value *CheckoutKeyInput
	isSet bool
}

func (v NullableCheckoutKeyInput) Get() *CheckoutKeyInput {
	return v.value
}

func (v *NullableCheckoutKeyInput) Set(val *CheckoutKeyInput) {
	v.value = val
	v.isSet = true
}

func (v NullableCheckoutKeyInput) IsSet() bool {
	return v.isSet
}

func (v *NullableCheckoutKeyInput) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCheckoutKeyInput(val *CheckoutKeyInput) *NullableCheckoutKeyInput {
	return &NullableCheckoutKeyInput{value: val, isSet: true}
}

func (v NullableCheckoutKeyInput) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCheckoutKeyInput) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
