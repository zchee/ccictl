// Copyright 2020 The ccictl Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package v1

import (
	"time"

	json "github.com/goccy/go-json"
)

// BuildDetail previous build
type BuildDetail struct {
	AllCommitDetails        *[]CommitDetail `json:"all_commit_details,omitempty"`
	JobName                 *string         `json:"job_name,omitempty"`
	PreviousSuccessfulBuild *PreviousBuild  `json:"previous_successful_build,omitempty"`
	UsageQueuedAt           *time.Time      `json:"usage_queued_at,omitempty"`
	User                    *User           `json:"user,omitempty"`
}

// NewBuildDetail instantiates a new BuildDetail object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBuildDetail() *BuildDetail {
	this := BuildDetail{}
	return &this
}

// NewBuildDetailWithDefaults instantiates a new BuildDetail object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBuildDetailWithDefaults() *BuildDetail {
	this := BuildDetail{}
	return &this
}

// GetAllCommitDetails returns the AllCommitDetails field value if set, zero value otherwise.
func (o *BuildDetail) GetAllCommitDetails() []CommitDetail {
	if o == nil || o.AllCommitDetails == nil {
		var ret []CommitDetail
		return ret
	}
	return *o.AllCommitDetails
}

// GetAllCommitDetailsOk returns a tuple with the AllCommitDetails field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BuildDetail) GetAllCommitDetailsOk() (*[]CommitDetail, bool) {
	if o == nil || o.AllCommitDetails == nil {
		return nil, false
	}
	return o.AllCommitDetails, true
}

// HasAllCommitDetails returns a boolean if a field has been set.
func (o *BuildDetail) HasAllCommitDetails() bool {
	if o != nil && o.AllCommitDetails != nil {
		return true
	}

	return false
}

// SetAllCommitDetails gets a reference to the given []CommitDetail and assigns it to the AllCommitDetails field.
func (o *BuildDetail) SetAllCommitDetails(v []CommitDetail) {
	o.AllCommitDetails = &v
}

// GetJobName returns the JobName field value if set, zero value otherwise.
func (o *BuildDetail) GetJobName() string {
	if o == nil || o.JobName == nil {
		var ret string
		return ret
	}
	return *o.JobName
}

// GetJobNameOk returns a tuple with the JobName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BuildDetail) GetJobNameOk() (*string, bool) {
	if o == nil || o.JobName == nil {
		return nil, false
	}
	return o.JobName, true
}

// HasJobName returns a boolean if a field has been set.
func (o *BuildDetail) HasJobName() bool {
	if o != nil && o.JobName != nil {
		return true
	}

	return false
}

// SetJobName gets a reference to the given string and assigns it to the JobName field.
func (o *BuildDetail) SetJobName(v string) {
	o.JobName = &v
}

// GetPreviousSuccessfulBuild returns the PreviousSuccessfulBuild field value if set, zero value otherwise.
func (o *BuildDetail) GetPreviousSuccessfulBuild() PreviousBuild {
	if o == nil || o.PreviousSuccessfulBuild == nil {
		var ret PreviousBuild
		return ret
	}
	return *o.PreviousSuccessfulBuild
}

// GetPreviousSuccessfulBuildOk returns a tuple with the PreviousSuccessfulBuild field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BuildDetail) GetPreviousSuccessfulBuildOk() (*PreviousBuild, bool) {
	if o == nil || o.PreviousSuccessfulBuild == nil {
		return nil, false
	}
	return o.PreviousSuccessfulBuild, true
}

// HasPreviousSuccessfulBuild returns a boolean if a field has been set.
func (o *BuildDetail) HasPreviousSuccessfulBuild() bool {
	if o != nil && o.PreviousSuccessfulBuild != nil {
		return true
	}

	return false
}

// SetPreviousSuccessfulBuild gets a reference to the given PreviousBuild and assigns it to the PreviousSuccessfulBuild field.
func (o *BuildDetail) SetPreviousSuccessfulBuild(v PreviousBuild) {
	o.PreviousSuccessfulBuild = &v
}

// GetUsageQueuedAt returns the UsageQueuedAt field value if set, zero value otherwise.
func (o *BuildDetail) GetUsageQueuedAt() time.Time {
	if o == nil || o.UsageQueuedAt == nil {
		var ret time.Time
		return ret
	}
	return *o.UsageQueuedAt
}

// GetUsageQueuedAtOk returns a tuple with the UsageQueuedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BuildDetail) GetUsageQueuedAtOk() (*time.Time, bool) {
	if o == nil || o.UsageQueuedAt == nil {
		return nil, false
	}
	return o.UsageQueuedAt, true
}

// HasUsageQueuedAt returns a boolean if a field has been set.
func (o *BuildDetail) HasUsageQueuedAt() bool {
	if o != nil && o.UsageQueuedAt != nil {
		return true
	}

	return false
}

// SetUsageQueuedAt gets a reference to the given time.Time and assigns it to the UsageQueuedAt field.
func (o *BuildDetail) SetUsageQueuedAt(v time.Time) {
	o.UsageQueuedAt = &v
}

// GetUser returns the User field value if set, zero value otherwise.
func (o *BuildDetail) GetUser() User {
	if o == nil || o.User == nil {
		var ret User
		return ret
	}
	return *o.User
}

// GetUserOk returns a tuple with the User field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *BuildDetail) GetUserOk() (*User, bool) {
	if o == nil || o.User == nil {
		return nil, false
	}
	return o.User, true
}

// HasUser returns a boolean if a field has been set.
func (o *BuildDetail) HasUser() bool {
	if o != nil && o.User != nil {
		return true
	}

	return false
}

// SetUser gets a reference to the given User and assigns it to the User field.
func (o *BuildDetail) SetUser(v User) {
	o.User = &v
}

func (o BuildDetail) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.AllCommitDetails != nil {
		toSerialize["all_commit_details"] = o.AllCommitDetails
	}
	if o.JobName != nil {
		toSerialize["job_name"] = o.JobName
	}
	if o.PreviousSuccessfulBuild != nil {
		toSerialize["previous_successful_build"] = o.PreviousSuccessfulBuild
	}
	if o.UsageQueuedAt != nil {
		toSerialize["usage_queued_at"] = o.UsageQueuedAt
	}
	if o.User != nil {
		toSerialize["user"] = o.User
	}
	return json.Marshal(toSerialize)
}

type NullableBuildDetail struct {
	value *BuildDetail
	isSet bool
}

func (v NullableBuildDetail) Get() *BuildDetail {
	return v.value
}

func (v *NullableBuildDetail) Set(val *BuildDetail) {
	v.value = val
	v.isSet = true
}

func (v NullableBuildDetail) IsSet() bool {
	return v.isSet
}

func (v *NullableBuildDetail) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBuildDetail(val *BuildDetail) *NullableBuildDetail {
	return &NullableBuildDetail{value: val, isSet: true}
}

func (v NullableBuildDetail) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBuildDetail) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
