// Copyright 2020 The ccictl Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package v1

import (
	"time"

	json "github.com/goccy/go-json"
)

// Build struct for Build
type Build struct {
	// commit message body
	Body            *string        `json:"body,omitempty"`
	Branch          *string        `json:"branch,omitempty"`
	BuildTimeMillis *int32         `json:"build_time_millis,omitempty"`
	BuildUrl        *string        `json:"build_url,omitempty"`
	CommitterEmail  *string        `json:"committer_email,omitempty"`
	CommitterName   *string        `json:"committer_name,omitempty"`
	Lifecycle       *Lifecycle     `json:"lifecycle,omitempty"`
	Previous        *PreviousBuild `json:"previous,omitempty"`
	// time build was queued
	QueuedAt *time.Time `json:"queued_at,omitempty"`
	Reponame *string    `json:"reponame,omitempty"`
	// time build started
	StartTime *time.Time `json:"start_time,omitempty"`
	// time build finished
	StopTime *time.Time `json:"stop_time,omitempty"`
	Subject  *string    `json:"subject,omitempty"`
	Username *string    `json:"username,omitempty"`
	VcsUrl   *string    `json:"vcs_url,omitempty"`
	// short string explaining the reason we built
	Why *string `json:"why,omitempty"`
}

// NewBuild instantiates a new Build object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBuild() *Build {
	this := Build{}
	return &this
}

// NewBuildWithDefaults instantiates a new Build object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBuildWithDefaults() *Build {
	this := Build{}
	return &this
}

// GetBody returns the Body field value if set, zero value otherwise.
func (o *Build) GetBody() string {
	if o == nil || o.Body == nil {
		var ret string
		return ret
	}
	return *o.Body
}

// GetBodyOk returns a tuple with the Body field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Build) GetBodyOk() (*string, bool) {
	if o == nil || o.Body == nil {
		return nil, false
	}
	return o.Body, true
}

// HasBody returns a boolean if a field has been set.
func (o *Build) HasBody() bool {
	if o != nil && o.Body != nil {
		return true
	}

	return false
}

// SetBody gets a reference to the given string and assigns it to the Body field.
func (o *Build) SetBody(v string) {
	o.Body = &v
}

// GetBranch returns the Branch field value if set, zero value otherwise.
func (o *Build) GetBranch() string {
	if o == nil || o.Branch == nil {
		var ret string
		return ret
	}
	return *o.Branch
}

// GetBranchOk returns a tuple with the Branch field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Build) GetBranchOk() (*string, bool) {
	if o == nil || o.Branch == nil {
		return nil, false
	}
	return o.Branch, true
}

// HasBranch returns a boolean if a field has been set.
func (o *Build) HasBranch() bool {
	if o != nil && o.Branch != nil {
		return true
	}

	return false
}

// SetBranch gets a reference to the given string and assigns it to the Branch field.
func (o *Build) SetBranch(v string) {
	o.Branch = &v
}

// GetBuildTimeMillis returns the BuildTimeMillis field value if set, zero value otherwise.
func (o *Build) GetBuildTimeMillis() int32 {
	if o == nil || o.BuildTimeMillis == nil {
		var ret int32
		return ret
	}
	return *o.BuildTimeMillis
}

// GetBuildTimeMillisOk returns a tuple with the BuildTimeMillis field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Build) GetBuildTimeMillisOk() (*int32, bool) {
	if o == nil || o.BuildTimeMillis == nil {
		return nil, false
	}
	return o.BuildTimeMillis, true
}

// HasBuildTimeMillis returns a boolean if a field has been set.
func (o *Build) HasBuildTimeMillis() bool {
	if o != nil && o.BuildTimeMillis != nil {
		return true
	}

	return false
}

// SetBuildTimeMillis gets a reference to the given int32 and assigns it to the BuildTimeMillis field.
func (o *Build) SetBuildTimeMillis(v int32) {
	o.BuildTimeMillis = &v
}

// GetBuildUrl returns the BuildUrl field value if set, zero value otherwise.
func (o *Build) GetBuildUrl() string {
	if o == nil || o.BuildUrl == nil {
		var ret string
		return ret
	}
	return *o.BuildUrl
}

// GetBuildUrlOk returns a tuple with the BuildUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Build) GetBuildUrlOk() (*string, bool) {
	if o == nil || o.BuildUrl == nil {
		return nil, false
	}
	return o.BuildUrl, true
}

// HasBuildUrl returns a boolean if a field has been set.
func (o *Build) HasBuildUrl() bool {
	if o != nil && o.BuildUrl != nil {
		return true
	}

	return false
}

// SetBuildUrl gets a reference to the given string and assigns it to the BuildUrl field.
func (o *Build) SetBuildUrl(v string) {
	o.BuildUrl = &v
}

// GetCommitterEmail returns the CommitterEmail field value if set, zero value otherwise.
func (o *Build) GetCommitterEmail() string {
	if o == nil || o.CommitterEmail == nil {
		var ret string
		return ret
	}
	return *o.CommitterEmail
}

// GetCommitterEmailOk returns a tuple with the CommitterEmail field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Build) GetCommitterEmailOk() (*string, bool) {
	if o == nil || o.CommitterEmail == nil {
		return nil, false
	}
	return o.CommitterEmail, true
}

// HasCommitterEmail returns a boolean if a field has been set.
func (o *Build) HasCommitterEmail() bool {
	if o != nil && o.CommitterEmail != nil {
		return true
	}

	return false
}

// SetCommitterEmail gets a reference to the given string and assigns it to the CommitterEmail field.
func (o *Build) SetCommitterEmail(v string) {
	o.CommitterEmail = &v
}

// GetCommitterName returns the CommitterName field value if set, zero value otherwise.
func (o *Build) GetCommitterName() string {
	if o == nil || o.CommitterName == nil {
		var ret string
		return ret
	}
	return *o.CommitterName
}

// GetCommitterNameOk returns a tuple with the CommitterName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Build) GetCommitterNameOk() (*string, bool) {
	if o == nil || o.CommitterName == nil {
		return nil, false
	}
	return o.CommitterName, true
}

// HasCommitterName returns a boolean if a field has been set.
func (o *Build) HasCommitterName() bool {
	if o != nil && o.CommitterName != nil {
		return true
	}

	return false
}

// SetCommitterName gets a reference to the given string and assigns it to the CommitterName field.
func (o *Build) SetCommitterName(v string) {
	o.CommitterName = &v
}

// GetLifecycle returns the Lifecycle field value if set, zero value otherwise.
func (o *Build) GetLifecycle() Lifecycle {
	if o == nil || o.Lifecycle == nil {
		var ret Lifecycle
		return ret
	}
	return *o.Lifecycle
}

// GetLifecycleOk returns a tuple with the Lifecycle field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Build) GetLifecycleOk() (*Lifecycle, bool) {
	if o == nil || o.Lifecycle == nil {
		return nil, false
	}
	return o.Lifecycle, true
}

// HasLifecycle returns a boolean if a field has been set.
func (o *Build) HasLifecycle() bool {
	if o != nil && o.Lifecycle != nil {
		return true
	}

	return false
}

// SetLifecycle gets a reference to the given Lifecycle and assigns it to the Lifecycle field.
func (o *Build) SetLifecycle(v Lifecycle) {
	o.Lifecycle = &v
}

// GetPrevious returns the Previous field value if set, zero value otherwise.
func (o *Build) GetPrevious() PreviousBuild {
	if o == nil || o.Previous == nil {
		var ret PreviousBuild
		return ret
	}
	return *o.Previous
}

// GetPreviousOk returns a tuple with the Previous field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Build) GetPreviousOk() (*PreviousBuild, bool) {
	if o == nil || o.Previous == nil {
		return nil, false
	}
	return o.Previous, true
}

// HasPrevious returns a boolean if a field has been set.
func (o *Build) HasPrevious() bool {
	if o != nil && o.Previous != nil {
		return true
	}

	return false
}

// SetPrevious gets a reference to the given PreviousBuild and assigns it to the Previous field.
func (o *Build) SetPrevious(v PreviousBuild) {
	o.Previous = &v
}

// GetQueuedAt returns the QueuedAt field value if set, zero value otherwise.
func (o *Build) GetQueuedAt() time.Time {
	if o == nil || o.QueuedAt == nil {
		var ret time.Time
		return ret
	}
	return *o.QueuedAt
}

// GetQueuedAtOk returns a tuple with the QueuedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Build) GetQueuedAtOk() (*time.Time, bool) {
	if o == nil || o.QueuedAt == nil {
		return nil, false
	}
	return o.QueuedAt, true
}

// HasQueuedAt returns a boolean if a field has been set.
func (o *Build) HasQueuedAt() bool {
	if o != nil && o.QueuedAt != nil {
		return true
	}

	return false
}

// SetQueuedAt gets a reference to the given time.Time and assigns it to the QueuedAt field.
func (o *Build) SetQueuedAt(v time.Time) {
	o.QueuedAt = &v
}

// GetReponame returns the Reponame field value if set, zero value otherwise.
func (o *Build) GetReponame() string {
	if o == nil || o.Reponame == nil {
		var ret string
		return ret
	}
	return *o.Reponame
}

// GetReponameOk returns a tuple with the Reponame field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Build) GetReponameOk() (*string, bool) {
	if o == nil || o.Reponame == nil {
		return nil, false
	}
	return o.Reponame, true
}

// HasReponame returns a boolean if a field has been set.
func (o *Build) HasReponame() bool {
	if o != nil && o.Reponame != nil {
		return true
	}

	return false
}

// SetReponame gets a reference to the given string and assigns it to the Reponame field.
func (o *Build) SetReponame(v string) {
	o.Reponame = &v
}

// GetStartTime returns the StartTime field value if set, zero value otherwise.
func (o *Build) GetStartTime() time.Time {
	if o == nil || o.StartTime == nil {
		var ret time.Time
		return ret
	}
	return *o.StartTime
}

// GetStartTimeOk returns a tuple with the StartTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Build) GetStartTimeOk() (*time.Time, bool) {
	if o == nil || o.StartTime == nil {
		return nil, false
	}
	return o.StartTime, true
}

// HasStartTime returns a boolean if a field has been set.
func (o *Build) HasStartTime() bool {
	if o != nil && o.StartTime != nil {
		return true
	}

	return false
}

// SetStartTime gets a reference to the given time.Time and assigns it to the StartTime field.
func (o *Build) SetStartTime(v time.Time) {
	o.StartTime = &v
}

// GetStopTime returns the StopTime field value if set, zero value otherwise.
func (o *Build) GetStopTime() time.Time {
	if o == nil || o.StopTime == nil {
		var ret time.Time
		return ret
	}
	return *o.StopTime
}

// GetStopTimeOk returns a tuple with the StopTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Build) GetStopTimeOk() (*time.Time, bool) {
	if o == nil || o.StopTime == nil {
		return nil, false
	}
	return o.StopTime, true
}

// HasStopTime returns a boolean if a field has been set.
func (o *Build) HasStopTime() bool {
	if o != nil && o.StopTime != nil {
		return true
	}

	return false
}

// SetStopTime gets a reference to the given time.Time and assigns it to the StopTime field.
func (o *Build) SetStopTime(v time.Time) {
	o.StopTime = &v
}

// GetSubject returns the Subject field value if set, zero value otherwise.
func (o *Build) GetSubject() string {
	if o == nil || o.Subject == nil {
		var ret string
		return ret
	}
	return *o.Subject
}

// GetSubjectOk returns a tuple with the Subject field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Build) GetSubjectOk() (*string, bool) {
	if o == nil || o.Subject == nil {
		return nil, false
	}
	return o.Subject, true
}

// HasSubject returns a boolean if a field has been set.
func (o *Build) HasSubject() bool {
	if o != nil && o.Subject != nil {
		return true
	}

	return false
}

// SetSubject gets a reference to the given string and assigns it to the Subject field.
func (o *Build) SetSubject(v string) {
	o.Subject = &v
}

// GetUsername returns the Username field value if set, zero value otherwise.
func (o *Build) GetUsername() string {
	if o == nil || o.Username == nil {
		var ret string
		return ret
	}
	return *o.Username
}

// GetUsernameOk returns a tuple with the Username field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Build) GetUsernameOk() (*string, bool) {
	if o == nil || o.Username == nil {
		return nil, false
	}
	return o.Username, true
}

// HasUsername returns a boolean if a field has been set.
func (o *Build) HasUsername() bool {
	if o != nil && o.Username != nil {
		return true
	}

	return false
}

// SetUsername gets a reference to the given string and assigns it to the Username field.
func (o *Build) SetUsername(v string) {
	o.Username = &v
}

// GetVcsUrl returns the VcsUrl field value if set, zero value otherwise.
func (o *Build) GetVcsUrl() string {
	if o == nil || o.VcsUrl == nil {
		var ret string
		return ret
	}
	return *o.VcsUrl
}

// GetVcsUrlOk returns a tuple with the VcsUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Build) GetVcsUrlOk() (*string, bool) {
	if o == nil || o.VcsUrl == nil {
		return nil, false
	}
	return o.VcsUrl, true
}

// HasVcsUrl returns a boolean if a field has been set.
func (o *Build) HasVcsUrl() bool {
	if o != nil && o.VcsUrl != nil {
		return true
	}

	return false
}

// SetVcsUrl gets a reference to the given string and assigns it to the VcsUrl field.
func (o *Build) SetVcsUrl(v string) {
	o.VcsUrl = &v
}

// GetWhy returns the Why field value if set, zero value otherwise.
func (o *Build) GetWhy() string {
	if o == nil || o.Why == nil {
		var ret string
		return ret
	}
	return *o.Why
}

// GetWhyOk returns a tuple with the Why field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Build) GetWhyOk() (*string, bool) {
	if o == nil || o.Why == nil {
		return nil, false
	}
	return o.Why, true
}

// HasWhy returns a boolean if a field has been set.
func (o *Build) HasWhy() bool {
	if o != nil && o.Why != nil {
		return true
	}

	return false
}

// SetWhy gets a reference to the given string and assigns it to the Why field.
func (o *Build) SetWhy(v string) {
	o.Why = &v
}

func (o Build) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Body != nil {
		toSerialize["body"] = o.Body
	}
	if o.Branch != nil {
		toSerialize["branch"] = o.Branch
	}
	if o.BuildTimeMillis != nil {
		toSerialize["build_time_millis"] = o.BuildTimeMillis
	}
	if o.BuildUrl != nil {
		toSerialize["build_url"] = o.BuildUrl
	}
	if o.CommitterEmail != nil {
		toSerialize["committer_email"] = o.CommitterEmail
	}
	if o.CommitterName != nil {
		toSerialize["committer_name"] = o.CommitterName
	}
	if o.Lifecycle != nil {
		toSerialize["lifecycle"] = o.Lifecycle
	}
	if o.Previous != nil {
		toSerialize["previous"] = o.Previous
	}
	if o.QueuedAt != nil {
		toSerialize["queued_at"] = o.QueuedAt
	}
	if o.Reponame != nil {
		toSerialize["reponame"] = o.Reponame
	}
	if o.StartTime != nil {
		toSerialize["start_time"] = o.StartTime
	}
	if o.StopTime != nil {
		toSerialize["stop_time"] = o.StopTime
	}
	if o.Subject != nil {
		toSerialize["subject"] = o.Subject
	}
	if o.Username != nil {
		toSerialize["username"] = o.Username
	}
	if o.VcsUrl != nil {
		toSerialize["vcs_url"] = o.VcsUrl
	}
	if o.Why != nil {
		toSerialize["why"] = o.Why
	}
	return json.Marshal(toSerialize)
}

type NullableBuild struct {
	value *Build
	isSet bool
}

func (v NullableBuild) Get() *Build {
	return v.value
}

func (v *NullableBuild) Set(val *Build) {
	v.value = val
	v.isSet = true
}

func (v NullableBuild) IsSet() bool {
	return v.isSet
}

func (v *NullableBuild) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBuild(val *Build) *NullableBuild {
	return &NullableBuild{value: val, isSet: true}
}

func (v NullableBuild) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBuild) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
