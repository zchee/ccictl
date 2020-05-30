// Copyright 2020 The ccictl Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package v2

import (
	"time"

	json "github.com/goccy/go-json"
)

// InlineResponse2003Items struct for InlineResponse2003Items
type InlineResponse2003Items struct {
	// The unique ID of the job.
	Id string `json:"id"`
	// The date and time the job started.
	StartedAt time.Time `json:"started_at"`
	// The time when the job stopped.
	StoppedAt time.Time `json:"stopped_at"`
	// Job status.
	Status string `json:"status"`
	// The number of credits used during execution
	CreditsUsed int64 `json:"credits_used"`
}

// NewInlineResponse2003Items instantiates a new InlineResponse2003Items object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineResponse2003Items(id string, startedAt time.Time, stoppedAt time.Time, status string, creditsUsed int64) *InlineResponse2003Items {
	this := InlineResponse2003Items{}
	this.Id = id
	this.StartedAt = startedAt
	this.StoppedAt = stoppedAt
	this.Status = status
	this.CreditsUsed = creditsUsed
	return &this
}

// NewInlineResponse2003ItemsWithDefaults instantiates a new InlineResponse2003Items object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineResponse2003ItemsWithDefaults() *InlineResponse2003Items {
	this := InlineResponse2003Items{}
	return &this
}

// GetId returns the Id field value
func (o *InlineResponse2003Items) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *InlineResponse2003Items) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *InlineResponse2003Items) SetId(v string) {
	o.Id = v
}

// GetStartedAt returns the StartedAt field value
func (o *InlineResponse2003Items) GetStartedAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.StartedAt
}

// GetStartedAtOk returns a tuple with the StartedAt field value
// and a boolean to check if the value has been set.
func (o *InlineResponse2003Items) GetStartedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.StartedAt, true
}

// SetStartedAt sets field value
func (o *InlineResponse2003Items) SetStartedAt(v time.Time) {
	o.StartedAt = v
}

// GetStoppedAt returns the StoppedAt field value
func (o *InlineResponse2003Items) GetStoppedAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.StoppedAt
}

// GetStoppedAtOk returns a tuple with the StoppedAt field value
// and a boolean to check if the value has been set.
func (o *InlineResponse2003Items) GetStoppedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.StoppedAt, true
}

// SetStoppedAt sets field value
func (o *InlineResponse2003Items) SetStoppedAt(v time.Time) {
	o.StoppedAt = v
}

// GetStatus returns the Status field value
func (o *InlineResponse2003Items) GetStatus() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *InlineResponse2003Items) GetStatusOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value
func (o *InlineResponse2003Items) SetStatus(v string) {
	o.Status = v
}

// GetCreditsUsed returns the CreditsUsed field value
func (o *InlineResponse2003Items) GetCreditsUsed() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.CreditsUsed
}

// GetCreditsUsedOk returns a tuple with the CreditsUsed field value
// and a boolean to check if the value has been set.
func (o *InlineResponse2003Items) GetCreditsUsedOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreditsUsed, true
}

// SetCreditsUsed sets field value
func (o *InlineResponse2003Items) SetCreditsUsed(v int64) {
	o.CreditsUsed = v
}

func (o InlineResponse2003Items) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["id"] = o.Id
	}
	if true {
		toSerialize["started_at"] = o.StartedAt
	}
	if true {
		toSerialize["stopped_at"] = o.StoppedAt
	}
	if true {
		toSerialize["status"] = o.Status
	}
	if true {
		toSerialize["credits_used"] = o.CreditsUsed
	}
	return json.Marshal(toSerialize)
}

type NullableInlineResponse2003Items struct {
	value *InlineResponse2003Items
	isSet bool
}

func (v NullableInlineResponse2003Items) Get() *InlineResponse2003Items {
	return v.value
}

func (v *NullableInlineResponse2003Items) Set(val *InlineResponse2003Items) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineResponse2003Items) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineResponse2003Items) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineResponse2003Items(val *InlineResponse2003Items) *NullableInlineResponse2003Items {
	return &NullableInlineResponse2003Items{value: val, isSet: true}
}

func (v NullableInlineResponse2003Items) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineResponse2003Items) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
