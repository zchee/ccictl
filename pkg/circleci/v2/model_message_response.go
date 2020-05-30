// Copyright 2020 The ccictl Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package v2

import (
	json "github.com/goccy/go-json"
)

// MessageResponse message response
type MessageResponse struct {
	// A human-readable message
	Message string `json:"message"`
}

// NewMessageResponse instantiates a new MessageResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewMessageResponse(message string) *MessageResponse {
	this := MessageResponse{}
	this.Message = message
	return &this
}

// NewMessageResponseWithDefaults instantiates a new MessageResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewMessageResponseWithDefaults() *MessageResponse {
	this := MessageResponse{}
	return &this
}

// GetMessage returns the Message field value
func (o *MessageResponse) GetMessage() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Message
}

// GetMessageOk returns a tuple with the Message field value
// and a boolean to check if the value has been set.
func (o *MessageResponse) GetMessageOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Message, true
}

// SetMessage sets field value
func (o *MessageResponse) SetMessage(v string) {
	o.Message = v
}

func (o MessageResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["message"] = o.Message
	}
	return json.Marshal(toSerialize)
}

type NullableMessageResponse struct {
	value *MessageResponse
	isSet bool
}

func (v NullableMessageResponse) Get() *MessageResponse {
	return v.value
}

func (v *NullableMessageResponse) Set(val *MessageResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableMessageResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableMessageResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMessageResponse(val *MessageResponse) *NullableMessageResponse {
	return &NullableMessageResponse{value: val, isSet: true}
}

func (v NullableMessageResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMessageResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
