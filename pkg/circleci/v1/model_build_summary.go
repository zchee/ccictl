// Copyright 2020 The ccictl Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package v1

import (
	"time"

	json "github.com/goccy/go-json"
)

// BuildSummary struct for BuildSummary
type BuildSummary struct {
	AddedAt     *time.Time `json:"added_at,omitempty"`
	BuildNum    *int32     `json:"build_num,omitempty"`
	Outcome     *Outcome   `json:"outcome,omitempty"`
	PushedAt    *time.Time `json:"pushed_at,omitempty"`
	Status      *Status    `json:"status,omitempty"`
	VcsRevision *string    `json:"vcs_revision,omitempty"`
}

// NewBuildSummary instantiates a new BuildSummary object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBuildSummary() *BuildSummary {
	this := BuildSummary{}
	return &this
}

// NewBuildSummaryWithDefaults instantiates a new BuildSummary object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBuildSummaryWithDefaults() *BuildSummary {
	this := BuildSummary{}
	return &this
}

// GetAddedAt returns the AddedAt field value if set, zero value otherwise.
func (o *BuildSummary) GetAddedAt() time.Time {
	if o == nil || o.AddedAt == nil {
		var ret time.Time
		return ret
	}
	return *o.AddedAt
}

// GetAddedAtOk returns a tuple with the AddedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BuildSummary) GetAddedAtOk() (*time.Time, bool) {
	if o == nil || o.AddedAt == nil {
		return nil, false
	}
	return o.AddedAt, true
}

// HasAddedAt returns a boolean if a field has been set.
func (o *BuildSummary) HasAddedAt() bool {
	if o != nil && o.AddedAt != nil {
		return true
	}

	return false
}

// SetAddedAt gets a reference to the given time.Time and assigns it to the AddedAt field.
func (o *BuildSummary) SetAddedAt(v time.Time) {
	o.AddedAt = &v
}

// GetBuildNum returns the BuildNum field value if set, zero value otherwise.
func (o *BuildSummary) GetBuildNum() int32 {
	if o == nil || o.BuildNum == nil {
		var ret int32
		return ret
	}
	return *o.BuildNum
}

// GetBuildNumOk returns a tuple with the BuildNum field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BuildSummary) GetBuildNumOk() (*int32, bool) {
	if o == nil || o.BuildNum == nil {
		return nil, false
	}
	return o.BuildNum, true
}

// HasBuildNum returns a boolean if a field has been set.
func (o *BuildSummary) HasBuildNum() bool {
	if o != nil && o.BuildNum != nil {
		return true
	}

	return false
}

// SetBuildNum gets a reference to the given int32 and assigns it to the BuildNum field.
func (o *BuildSummary) SetBuildNum(v int32) {
	o.BuildNum = &v
}

// GetOutcome returns the Outcome field value if set, zero value otherwise.
func (o *BuildSummary) GetOutcome() Outcome {
	if o == nil || o.Outcome == nil {
		var ret Outcome
		return ret
	}
	return *o.Outcome
}

// GetOutcomeOk returns a tuple with the Outcome field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BuildSummary) GetOutcomeOk() (*Outcome, bool) {
	if o == nil || o.Outcome == nil {
		return nil, false
	}
	return o.Outcome, true
}

// HasOutcome returns a boolean if a field has been set.
func (o *BuildSummary) HasOutcome() bool {
	if o != nil && o.Outcome != nil {
		return true
	}

	return false
}

// SetOutcome gets a reference to the given Outcome and assigns it to the Outcome field.
func (o *BuildSummary) SetOutcome(v Outcome) {
	o.Outcome = &v
}

// GetPushedAt returns the PushedAt field value if set, zero value otherwise.
func (o *BuildSummary) GetPushedAt() time.Time {
	if o == nil || o.PushedAt == nil {
		var ret time.Time
		return ret
	}
	return *o.PushedAt
}

// GetPushedAtOk returns a tuple with the PushedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BuildSummary) GetPushedAtOk() (*time.Time, bool) {
	if o == nil || o.PushedAt == nil {
		return nil, false
	}
	return o.PushedAt, true
}

// HasPushedAt returns a boolean if a field has been set.
func (o *BuildSummary) HasPushedAt() bool {
	if o != nil && o.PushedAt != nil {
		return true
	}

	return false
}

// SetPushedAt gets a reference to the given time.Time and assigns it to the PushedAt field.
func (o *BuildSummary) SetPushedAt(v time.Time) {
	o.PushedAt = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *BuildSummary) GetStatus() Status {
	if o == nil || o.Status == nil {
		var ret Status
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BuildSummary) GetStatusOk() (*Status, bool) {
	if o == nil || o.Status == nil {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *BuildSummary) HasStatus() bool {
	if o != nil && o.Status != nil {
		return true
	}

	return false
}

// SetStatus gets a reference to the given Status and assigns it to the Status field.
func (o *BuildSummary) SetStatus(v Status) {
	o.Status = &v
}

// GetVcsRevision returns the VcsRevision field value if set, zero value otherwise.
func (o *BuildSummary) GetVcsRevision() string {
	if o == nil || o.VcsRevision == nil {
		var ret string
		return ret
	}
	return *o.VcsRevision
}

// GetVcsRevisionOk returns a tuple with the VcsRevision field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BuildSummary) GetVcsRevisionOk() (*string, bool) {
	if o == nil || o.VcsRevision == nil {
		return nil, false
	}
	return o.VcsRevision, true
}

// HasVcsRevision returns a boolean if a field has been set.
func (o *BuildSummary) HasVcsRevision() bool {
	if o != nil && o.VcsRevision != nil {
		return true
	}

	return false
}

// SetVcsRevision gets a reference to the given string and assigns it to the VcsRevision field.
func (o *BuildSummary) SetVcsRevision(v string) {
	o.VcsRevision = &v
}

func (o BuildSummary) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.AddedAt != nil {
		toSerialize["added_at"] = o.AddedAt
	}
	if o.BuildNum != nil {
		toSerialize["build_num"] = o.BuildNum
	}
	if o.Outcome != nil {
		toSerialize["outcome"] = o.Outcome
	}
	if o.PushedAt != nil {
		toSerialize["pushed_at"] = o.PushedAt
	}
	if o.Status != nil {
		toSerialize["status"] = o.Status
	}
	if o.VcsRevision != nil {
		toSerialize["vcs_revision"] = o.VcsRevision
	}
	return json.Marshal(toSerialize)
}

type NullableBuildSummary struct {
	value *BuildSummary
	isSet bool
}

func (v NullableBuildSummary) Get() *BuildSummary {
	return v.value
}

func (v *NullableBuildSummary) Set(val *BuildSummary) {
	v.value = val
	v.isSet = true
}

func (v NullableBuildSummary) IsSet() bool {
	return v.isSet
}

func (v *NullableBuildSummary) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBuildSummary(val *BuildSummary) *NullableBuildSummary {
	return &NullableBuildSummary{value: val, isSet: true}
}

func (v NullableBuildSummary) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBuildSummary) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
