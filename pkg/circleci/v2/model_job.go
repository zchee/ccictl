// Copyright 2020 The ccictl Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package v2

import (
	"time"

	json "github.com/goccy/go-json"
)

// Job Job
type Job struct {
	// The unique ID of the user.
	CanceledBy *string `json:"canceled_by,omitempty"`
	// A sequence of the unique job IDs for the jobs that this job depends upon in the workflow.
	Dependencies []string `json:"dependencies"`
	// The number of the job.
	JobNumber *int64 `json:"job_number,omitempty"`
	// The unique ID of the job.
	Id string `json:"id"`
	// The date and time the job started.
	StartedAt time.Time `json:"started_at"`
	// The name of the job.
	Name string `json:"name"`
	// The unique ID of the user.
	ApprovedBy *string `json:"approved_by,omitempty"`
	// The project-slug for the job.
	ProjectSlug string `json:"project_slug"`
	// The current status of the job.
	Status interface{} `json:"status"`
	// The type of job.
	Type string `json:"type"`
	// The time when the job stopped.
	StoppedAt *time.Time `json:"stopped_at,omitempty"`
	// The unique ID of the job.
	ApprovalRequestId *string `json:"approval_request_id,omitempty"`
}

// NewJob instantiates a new Job object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewJob(dependencies []string, id string, startedAt time.Time, name string, projectSlug string, status interface{}, type_ string) *Job {
	this := Job{}
	this.Dependencies = dependencies
	this.Id = id
	this.StartedAt = startedAt
	this.Name = name
	this.ProjectSlug = projectSlug
	this.Status = status
	this.Type = type_
	return &this
}

// NewJobWithDefaults instantiates a new Job object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewJobWithDefaults() *Job {
	this := Job{}
	return &this
}

// GetCanceledBy returns the CanceledBy field value if set, zero value otherwise.
func (o *Job) GetCanceledBy() string {
	if o == nil || o.CanceledBy == nil {
		var ret string
		return ret
	}
	return *o.CanceledBy
}

// GetCanceledByOk returns a tuple with the CanceledBy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Job) GetCanceledByOk() (*string, bool) {
	if o == nil || o.CanceledBy == nil {
		return nil, false
	}
	return o.CanceledBy, true
}

// HasCanceledBy returns a boolean if a field has been set.
func (o *Job) HasCanceledBy() bool {
	if o != nil && o.CanceledBy != nil {
		return true
	}

	return false
}

// SetCanceledBy gets a reference to the given string and assigns it to the CanceledBy field.
func (o *Job) SetCanceledBy(v string) {
	o.CanceledBy = &v
}

// GetDependencies returns the Dependencies field value
func (o *Job) GetDependencies() []string {
	if o == nil {
		var ret []string
		return ret
	}

	return o.Dependencies
}

// GetDependenciesOk returns a tuple with the Dependencies field value
// and a boolean to check if the value has been set.
func (o *Job) GetDependenciesOk() (*[]string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Dependencies, true
}

// SetDependencies sets field value
func (o *Job) SetDependencies(v []string) {
	o.Dependencies = v
}

// GetJobNumber returns the JobNumber field value if set, zero value otherwise.
func (o *Job) GetJobNumber() int64 {
	if o == nil || o.JobNumber == nil {
		var ret int64
		return ret
	}
	return *o.JobNumber
}

// GetJobNumberOk returns a tuple with the JobNumber field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Job) GetJobNumberOk() (*int64, bool) {
	if o == nil || o.JobNumber == nil {
		return nil, false
	}
	return o.JobNumber, true
}

// HasJobNumber returns a boolean if a field has been set.
func (o *Job) HasJobNumber() bool {
	if o != nil && o.JobNumber != nil {
		return true
	}

	return false
}

// SetJobNumber gets a reference to the given int64 and assigns it to the JobNumber field.
func (o *Job) SetJobNumber(v int64) {
	o.JobNumber = &v
}

// GetId returns the Id field value
func (o *Job) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *Job) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *Job) SetId(v string) {
	o.Id = v
}

// GetStartedAt returns the StartedAt field value
func (o *Job) GetStartedAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.StartedAt
}

// GetStartedAtOk returns a tuple with the StartedAt field value
// and a boolean to check if the value has been set.
func (o *Job) GetStartedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.StartedAt, true
}

// SetStartedAt sets field value
func (o *Job) SetStartedAt(v time.Time) {
	o.StartedAt = v
}

// GetName returns the Name field value
func (o *Job) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *Job) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *Job) SetName(v string) {
	o.Name = v
}

// GetApprovedBy returns the ApprovedBy field value if set, zero value otherwise.
func (o *Job) GetApprovedBy() string {
	if o == nil || o.ApprovedBy == nil {
		var ret string
		return ret
	}
	return *o.ApprovedBy
}

// GetApprovedByOk returns a tuple with the ApprovedBy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Job) GetApprovedByOk() (*string, bool) {
	if o == nil || o.ApprovedBy == nil {
		return nil, false
	}
	return o.ApprovedBy, true
}

// HasApprovedBy returns a boolean if a field has been set.
func (o *Job) HasApprovedBy() bool {
	if o != nil && o.ApprovedBy != nil {
		return true
	}

	return false
}

// SetApprovedBy gets a reference to the given string and assigns it to the ApprovedBy field.
func (o *Job) SetApprovedBy(v string) {
	o.ApprovedBy = &v
}

// GetProjectSlug returns the ProjectSlug field value
func (o *Job) GetProjectSlug() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ProjectSlug
}

// GetProjectSlugOk returns a tuple with the ProjectSlug field value
// and a boolean to check if the value has been set.
func (o *Job) GetProjectSlugOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ProjectSlug, true
}

// SetProjectSlug sets field value
func (o *Job) SetProjectSlug(v string) {
	o.ProjectSlug = v
}

// GetStatus returns the Status field value
// If the value is explicit nil, the zero value for interface{} will be returned
func (o *Job) GetStatus() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}

	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *Job) GetStatusOk() (*interface{}, bool) {
	if o == nil || o.Status == nil {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value
func (o *Job) SetStatus(v interface{}) {
	o.Status = v
}

// GetType returns the Type field value
func (o *Job) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *Job) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *Job) SetType(v string) {
	o.Type = v
}

// GetStoppedAt returns the StoppedAt field value if set, zero value otherwise.
func (o *Job) GetStoppedAt() time.Time {
	if o == nil || o.StoppedAt == nil {
		var ret time.Time
		return ret
	}
	return *o.StoppedAt
}

// GetStoppedAtOk returns a tuple with the StoppedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Job) GetStoppedAtOk() (*time.Time, bool) {
	if o == nil || o.StoppedAt == nil {
		return nil, false
	}
	return o.StoppedAt, true
}

// HasStoppedAt returns a boolean if a field has been set.
func (o *Job) HasStoppedAt() bool {
	if o != nil && o.StoppedAt != nil {
		return true
	}

	return false
}

// SetStoppedAt gets a reference to the given time.Time and assigns it to the StoppedAt field.
func (o *Job) SetStoppedAt(v time.Time) {
	o.StoppedAt = &v
}

// GetApprovalRequestId returns the ApprovalRequestId field value if set, zero value otherwise.
func (o *Job) GetApprovalRequestId() string {
	if o == nil || o.ApprovalRequestId == nil {
		var ret string
		return ret
	}
	return *o.ApprovalRequestId
}

// GetApprovalRequestIdOk returns a tuple with the ApprovalRequestId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Job) GetApprovalRequestIdOk() (*string, bool) {
	if o == nil || o.ApprovalRequestId == nil {
		return nil, false
	}
	return o.ApprovalRequestId, true
}

// HasApprovalRequestId returns a boolean if a field has been set.
func (o *Job) HasApprovalRequestId() bool {
	if o != nil && o.ApprovalRequestId != nil {
		return true
	}

	return false
}

// SetApprovalRequestId gets a reference to the given string and assigns it to the ApprovalRequestId field.
func (o *Job) SetApprovalRequestId(v string) {
	o.ApprovalRequestId = &v
}

func (o Job) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.CanceledBy != nil {
		toSerialize["canceled_by"] = o.CanceledBy
	}
	if true {
		toSerialize["dependencies"] = o.Dependencies
	}
	if o.JobNumber != nil {
		toSerialize["job_number"] = o.JobNumber
	}
	if true {
		toSerialize["id"] = o.Id
	}
	if true {
		toSerialize["started_at"] = o.StartedAt
	}
	if true {
		toSerialize["name"] = o.Name
	}
	if o.ApprovedBy != nil {
		toSerialize["approved_by"] = o.ApprovedBy
	}
	if true {
		toSerialize["project_slug"] = o.ProjectSlug
	}
	if o.Status != nil {
		toSerialize["status"] = o.Status
	}
	if true {
		toSerialize["type"] = o.Type
	}
	if o.StoppedAt != nil {
		toSerialize["stopped_at"] = o.StoppedAt
	}
	if o.ApprovalRequestId != nil {
		toSerialize["approval_request_id"] = o.ApprovalRequestId
	}
	return json.Marshal(toSerialize)
}

type NullableJob struct {
	value *Job
	isSet bool
}

func (v NullableJob) Get() *Job {
	return v.value
}

func (v *NullableJob) Set(val *Job) {
	v.value = val
	v.isSet = true
}

func (v NullableJob) IsSet() bool {
	return v.isSet
}

func (v *NullableJob) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableJob(val *Job) *NullableJob {
	return &NullableJob{value: val, isSet: true}
}

func (v NullableJob) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableJob) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
