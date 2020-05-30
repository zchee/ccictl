// Copyright 2020 The ccictl Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package v2

import (
	"time"

	json "github.com/goccy/go-json"
)

// InlineResponse2001Items struct for InlineResponse2001Items
type InlineResponse2001Items struct {
	// The unique ID of the workflow.
	Id string `json:"id"`
	// The duration in seconds of a run.
	Duration int64 `json:"duration"`
	// The date and time the workflow was created.
	CreatedAt time.Time `json:"created_at"`
	// The date and time the workflow stopped.
	StoppedAt time.Time `json:"stopped_at"`
	// The number of credits used during execution
	CreditsUsed int64 `json:"credits_used"`
	// Workflow status.
	Status string `json:"status"`
}

// NewInlineResponse2001Items instantiates a new InlineResponse2001Items object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineResponse2001Items(id string, duration int64, createdAt time.Time, stoppedAt time.Time, creditsUsed int64, status string) *InlineResponse2001Items {
	this := InlineResponse2001Items{}
	this.Id = id
	this.Duration = duration
	this.CreatedAt = createdAt
	this.StoppedAt = stoppedAt
	this.CreditsUsed = creditsUsed
	this.Status = status
	return &this
}

// NewInlineResponse2001ItemsWithDefaults instantiates a new InlineResponse2001Items object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineResponse2001ItemsWithDefaults() *InlineResponse2001Items {
	this := InlineResponse2001Items{}
	return &this
}

// GetId returns the Id field value
func (o *InlineResponse2001Items) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *InlineResponse2001Items) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *InlineResponse2001Items) SetId(v string) {
	o.Id = v
}

// GetDuration returns the Duration field value
func (o *InlineResponse2001Items) GetDuration() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.Duration
}

// GetDurationOk returns a tuple with the Duration field value
// and a boolean to check if the value has been set.
func (o *InlineResponse2001Items) GetDurationOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Duration, true
}

// SetDuration sets field value
func (o *InlineResponse2001Items) SetDuration(v int64) {
	o.Duration = v
}

// GetCreatedAt returns the CreatedAt field value
func (o *InlineResponse2001Items) GetCreatedAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value
// and a boolean to check if the value has been set.
func (o *InlineResponse2001Items) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreatedAt, true
}

// SetCreatedAt sets field value
func (o *InlineResponse2001Items) SetCreatedAt(v time.Time) {
	o.CreatedAt = v
}

// GetStoppedAt returns the StoppedAt field value
func (o *InlineResponse2001Items) GetStoppedAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.StoppedAt
}

// GetStoppedAtOk returns a tuple with the StoppedAt field value
// and a boolean to check if the value has been set.
func (o *InlineResponse2001Items) GetStoppedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.StoppedAt, true
}

// SetStoppedAt sets field value
func (o *InlineResponse2001Items) SetStoppedAt(v time.Time) {
	o.StoppedAt = v
}

// GetCreditsUsed returns the CreditsUsed field value
func (o *InlineResponse2001Items) GetCreditsUsed() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.CreditsUsed
}

// GetCreditsUsedOk returns a tuple with the CreditsUsed field value
// and a boolean to check if the value has been set.
func (o *InlineResponse2001Items) GetCreditsUsedOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreditsUsed, true
}

// SetCreditsUsed sets field value
func (o *InlineResponse2001Items) SetCreditsUsed(v int64) {
	o.CreditsUsed = v
}

// GetStatus returns the Status field value
func (o *InlineResponse2001Items) GetStatus() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *InlineResponse2001Items) GetStatusOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value
func (o *InlineResponse2001Items) SetStatus(v string) {
	o.Status = v
}

func (o InlineResponse2001Items) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["id"] = o.Id
	}
	if true {
		toSerialize["duration"] = o.Duration
	}
	if true {
		toSerialize["created_at"] = o.CreatedAt
	}
	if true {
		toSerialize["stopped_at"] = o.StoppedAt
	}
	if true {
		toSerialize["credits_used"] = o.CreditsUsed
	}
	if true {
		toSerialize["status"] = o.Status
	}
	return json.Marshal(toSerialize)
}

type NullableInlineResponse2001Items struct {
	value *InlineResponse2001Items
	isSet bool
}

func (v NullableInlineResponse2001Items) Get() *InlineResponse2001Items {
	return v.value
}

func (v *NullableInlineResponse2001Items) Set(val *InlineResponse2001Items) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineResponse2001Items) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineResponse2001Items) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineResponse2001Items(val *InlineResponse2001Items) *NullableInlineResponse2001Items {
	return &NullableInlineResponse2001Items{value: val, isSet: true}
}

func (v NullableInlineResponse2001Items) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineResponse2001Items) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
