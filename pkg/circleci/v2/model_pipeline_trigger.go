// Copyright 2020 The ccictl Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package v2

import (
	"time"

	json "github.com/goccy/go-json"
)

// PipelineTrigger A summary of the trigger.
type PipelineTrigger struct {
	// The type of trigger.
	Type string `json:"type"`
	// The date and time the trigger was received.
	ReceivedAt time.Time            `json:"received_at"`
	Actor      PipelineTriggerActor `json:"actor"`
}

// NewPipelineTrigger instantiates a new PipelineTrigger object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPipelineTrigger(type_ string, receivedAt time.Time, actor PipelineTriggerActor) *PipelineTrigger {
	this := PipelineTrigger{}
	this.Type = type_
	this.ReceivedAt = receivedAt
	this.Actor = actor
	return &this
}

// NewPipelineTriggerWithDefaults instantiates a new PipelineTrigger object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPipelineTriggerWithDefaults() *PipelineTrigger {
	this := PipelineTrigger{}
	return &this
}

// GetType returns the Type field value
func (o *PipelineTrigger) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *PipelineTrigger) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *PipelineTrigger) SetType(v string) {
	o.Type = v
}

// GetReceivedAt returns the ReceivedAt field value
func (o *PipelineTrigger) GetReceivedAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.ReceivedAt
}

// GetReceivedAtOk returns a tuple with the ReceivedAt field value
// and a boolean to check if the value has been set.
func (o *PipelineTrigger) GetReceivedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ReceivedAt, true
}

// SetReceivedAt sets field value
func (o *PipelineTrigger) SetReceivedAt(v time.Time) {
	o.ReceivedAt = v
}

// GetActor returns the Actor field value
func (o *PipelineTrigger) GetActor() PipelineTriggerActor {
	if o == nil {
		var ret PipelineTriggerActor
		return ret
	}

	return o.Actor
}

// GetActorOk returns a tuple with the Actor field value
// and a boolean to check if the value has been set.
func (o *PipelineTrigger) GetActorOk() (*PipelineTriggerActor, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Actor, true
}

// SetActor sets field value
func (o *PipelineTrigger) SetActor(v PipelineTriggerActor) {
	o.Actor = v
}

func (o PipelineTrigger) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["type"] = o.Type
	}
	if true {
		toSerialize["received_at"] = o.ReceivedAt
	}
	if true {
		toSerialize["actor"] = o.Actor
	}
	return json.Marshal(toSerialize)
}

type NullablePipelineTrigger struct {
	value *PipelineTrigger
	isSet bool
}

func (v NullablePipelineTrigger) Get() *PipelineTrigger {
	return v.value
}

func (v *NullablePipelineTrigger) Set(val *PipelineTrigger) {
	v.value = val
	v.isSet = true
}

func (v NullablePipelineTrigger) IsSet() bool {
	return v.isSet
}

func (v *NullablePipelineTrigger) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePipelineTrigger(val *PipelineTrigger) *NullablePipelineTrigger {
	return &NullablePipelineTrigger{value: val, isSet: true}
}

func (v NullablePipelineTrigger) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePipelineTrigger) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
