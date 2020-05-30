// Copyright 2020 The ccictl Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package v2

import (
	json "github.com/goccy/go-json"
)

// JobDetailsMessages Message from CircleCI execution platform.
type JobDetailsMessages struct {
	// Message type.
	Type string `json:"type"`
	// Information describing message.
	Message string `json:"message"`
	// Value describing the reason for message to be added to the job.
	Reason *string `json:"reason,omitempty"`
}

// NewJobDetailsMessages instantiates a new JobDetailsMessages object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewJobDetailsMessages(type_ string, message string) *JobDetailsMessages {
	this := JobDetailsMessages{}
	this.Type = type_
	this.Message = message
	return &this
}

// NewJobDetailsMessagesWithDefaults instantiates a new JobDetailsMessages object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewJobDetailsMessagesWithDefaults() *JobDetailsMessages {
	this := JobDetailsMessages{}
	return &this
}

// GetType returns the Type field value
func (o *JobDetailsMessages) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *JobDetailsMessages) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *JobDetailsMessages) SetType(v string) {
	o.Type = v
}

// GetMessage returns the Message field value
func (o *JobDetailsMessages) GetMessage() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Message
}

// GetMessageOk returns a tuple with the Message field value
// and a boolean to check if the value has been set.
func (o *JobDetailsMessages) GetMessageOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Message, true
}

// SetMessage sets field value
func (o *JobDetailsMessages) SetMessage(v string) {
	o.Message = v
}

// GetReason returns the Reason field value if set, zero value otherwise.
func (o *JobDetailsMessages) GetReason() string {
	if o == nil || o.Reason == nil {
		var ret string
		return ret
	}
	return *o.Reason
}

// GetReasonOk returns a tuple with the Reason field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *JobDetailsMessages) GetReasonOk() (*string, bool) {
	if o == nil || o.Reason == nil {
		return nil, false
	}
	return o.Reason, true
}

// HasReason returns a boolean if a field has been set.
func (o *JobDetailsMessages) HasReason() bool {
	if o != nil && o.Reason != nil {
		return true
	}

	return false
}

// SetReason gets a reference to the given string and assigns it to the Reason field.
func (o *JobDetailsMessages) SetReason(v string) {
	o.Reason = &v
}

func (o JobDetailsMessages) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["type"] = o.Type
	}
	if true {
		toSerialize["message"] = o.Message
	}
	if o.Reason != nil {
		toSerialize["reason"] = o.Reason
	}
	return json.Marshal(toSerialize)
}

type NullableJobDetailsMessages struct {
	value *JobDetailsMessages
	isSet bool
}

func (v NullableJobDetailsMessages) Get() *JobDetailsMessages {
	return v.value
}

func (v *NullableJobDetailsMessages) Set(val *JobDetailsMessages) {
	v.value = val
	v.isSet = true
}

func (v NullableJobDetailsMessages) IsSet() bool {
	return v.isSet
}

func (v *NullableJobDetailsMessages) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableJobDetailsMessages(val *JobDetailsMessages) *NullableJobDetailsMessages {
	return &NullableJobDetailsMessages{value: val, isSet: true}
}

func (v NullableJobDetailsMessages) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableJobDetailsMessages) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
