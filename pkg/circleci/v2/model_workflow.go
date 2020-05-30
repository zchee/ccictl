// Copyright 2020 The ccictl Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package v2

import (
	"time"

	json "github.com/goccy/go-json"
)

// Workflow A workflow
type Workflow struct {
	// The ID of the pipeline this workflow belongs to.
	PipelineId string  `json:"pipeline_id"`
	CanceledBy *string `json:"canceled_by,omitempty"`
	// The unique ID of the workflow.
	Id string `json:"id"`
	// The name of the workflow.
	Name string `json:"name"`
	// The project-slug for the pipeline this workflow belongs to.
	ProjectSlug string  `json:"project_slug"`
	ErroredBy   *string `json:"errored_by,omitempty"`
	// The current status of the workflow.
	Status    string `json:"status"`
	StartedBy string `json:"started_by"`
	// The number of the pipeline this workflow belongs to.
	PipelineNumber int64 `json:"pipeline_number"`
	// The date and time the workflow was created.
	CreatedAt time.Time `json:"created_at"`
	// The date and time the workflow stopped.
	StoppedAt time.Time `json:"stopped_at"`
}

// NewWorkflow instantiates a new Workflow object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWorkflow(pipelineId string, id string, name string, projectSlug string, status string, startedBy string, pipelineNumber int64, createdAt time.Time, stoppedAt time.Time) *Workflow {
	this := Workflow{}
	this.PipelineId = pipelineId
	this.Id = id
	this.Name = name
	this.ProjectSlug = projectSlug
	this.Status = status
	this.StartedBy = startedBy
	this.PipelineNumber = pipelineNumber
	this.CreatedAt = createdAt
	this.StoppedAt = stoppedAt
	return &this
}

// NewWorkflowWithDefaults instantiates a new Workflow object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWorkflowWithDefaults() *Workflow {
	this := Workflow{}
	return &this
}

// GetPipelineId returns the PipelineId field value
func (o *Workflow) GetPipelineId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.PipelineId
}

// GetPipelineIdOk returns a tuple with the PipelineId field value
// and a boolean to check if the value has been set.
func (o *Workflow) GetPipelineIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PipelineId, true
}

// SetPipelineId sets field value
func (o *Workflow) SetPipelineId(v string) {
	o.PipelineId = v
}

// GetCanceledBy returns the CanceledBy field value if set, zero value otherwise.
func (o *Workflow) GetCanceledBy() string {
	if o == nil || o.CanceledBy == nil {
		var ret string
		return ret
	}
	return *o.CanceledBy
}

// GetCanceledByOk returns a tuple with the CanceledBy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Workflow) GetCanceledByOk() (*string, bool) {
	if o == nil || o.CanceledBy == nil {
		return nil, false
	}
	return o.CanceledBy, true
}

// HasCanceledBy returns a boolean if a field has been set.
func (o *Workflow) HasCanceledBy() bool {
	if o != nil && o.CanceledBy != nil {
		return true
	}

	return false
}

// SetCanceledBy gets a reference to the given string and assigns it to the CanceledBy field.
func (o *Workflow) SetCanceledBy(v string) {
	o.CanceledBy = &v
}

// GetId returns the Id field value
func (o *Workflow) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *Workflow) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *Workflow) SetId(v string) {
	o.Id = v
}

// GetName returns the Name field value
func (o *Workflow) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *Workflow) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *Workflow) SetName(v string) {
	o.Name = v
}

// GetProjectSlug returns the ProjectSlug field value
func (o *Workflow) GetProjectSlug() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ProjectSlug
}

// GetProjectSlugOk returns a tuple with the ProjectSlug field value
// and a boolean to check if the value has been set.
func (o *Workflow) GetProjectSlugOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ProjectSlug, true
}

// SetProjectSlug sets field value
func (o *Workflow) SetProjectSlug(v string) {
	o.ProjectSlug = v
}

// GetErroredBy returns the ErroredBy field value if set, zero value otherwise.
func (o *Workflow) GetErroredBy() string {
	if o == nil || o.ErroredBy == nil {
		var ret string
		return ret
	}
	return *o.ErroredBy
}

// GetErroredByOk returns a tuple with the ErroredBy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Workflow) GetErroredByOk() (*string, bool) {
	if o == nil || o.ErroredBy == nil {
		return nil, false
	}
	return o.ErroredBy, true
}

// HasErroredBy returns a boolean if a field has been set.
func (o *Workflow) HasErroredBy() bool {
	if o != nil && o.ErroredBy != nil {
		return true
	}

	return false
}

// SetErroredBy gets a reference to the given string and assigns it to the ErroredBy field.
func (o *Workflow) SetErroredBy(v string) {
	o.ErroredBy = &v
}

// GetStatus returns the Status field value
func (o *Workflow) GetStatus() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
func (o *Workflow) GetStatusOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value
func (o *Workflow) SetStatus(v string) {
	o.Status = v
}

// GetStartedBy returns the StartedBy field value
func (o *Workflow) GetStartedBy() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.StartedBy
}

// GetStartedByOk returns a tuple with the StartedBy field value
// and a boolean to check if the value has been set.
func (o *Workflow) GetStartedByOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.StartedBy, true
}

// SetStartedBy sets field value
func (o *Workflow) SetStartedBy(v string) {
	o.StartedBy = v
}

// GetPipelineNumber returns the PipelineNumber field value
func (o *Workflow) GetPipelineNumber() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.PipelineNumber
}

// GetPipelineNumberOk returns a tuple with the PipelineNumber field value
// and a boolean to check if the value has been set.
func (o *Workflow) GetPipelineNumberOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PipelineNumber, true
}

// SetPipelineNumber sets field value
func (o *Workflow) SetPipelineNumber(v int64) {
	o.PipelineNumber = v
}

// GetCreatedAt returns the CreatedAt field value
func (o *Workflow) GetCreatedAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value
// and a boolean to check if the value has been set.
func (o *Workflow) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreatedAt, true
}

// SetCreatedAt sets field value
func (o *Workflow) SetCreatedAt(v time.Time) {
	o.CreatedAt = v
}

// GetStoppedAt returns the StoppedAt field value
func (o *Workflow) GetStoppedAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.StoppedAt
}

// GetStoppedAtOk returns a tuple with the StoppedAt field value
// and a boolean to check if the value has been set.
func (o *Workflow) GetStoppedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.StoppedAt, true
}

// SetStoppedAt sets field value
func (o *Workflow) SetStoppedAt(v time.Time) {
	o.StoppedAt = v
}

func (o Workflow) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["pipeline_id"] = o.PipelineId
	}
	if o.CanceledBy != nil {
		toSerialize["canceled_by"] = o.CanceledBy
	}
	if true {
		toSerialize["id"] = o.Id
	}
	if true {
		toSerialize["name"] = o.Name
	}
	if true {
		toSerialize["project_slug"] = o.ProjectSlug
	}
	if o.ErroredBy != nil {
		toSerialize["errored_by"] = o.ErroredBy
	}
	if true {
		toSerialize["status"] = o.Status
	}
	if true {
		toSerialize["started_by"] = o.StartedBy
	}
	if true {
		toSerialize["pipeline_number"] = o.PipelineNumber
	}
	if true {
		toSerialize["created_at"] = o.CreatedAt
	}
	if true {
		toSerialize["stopped_at"] = o.StoppedAt
	}
	return json.Marshal(toSerialize)
}

type NullableWorkflow struct {
	value *Workflow
	isSet bool
}

func (v NullableWorkflow) Get() *Workflow {
	return v.value
}

func (v *NullableWorkflow) Set(val *Workflow) {
	v.value = val
	v.isSet = true
}

func (v NullableWorkflow) IsSet() bool {
	return v.isSet
}

func (v *NullableWorkflow) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWorkflow(val *Workflow) *NullableWorkflow {
	return &NullableWorkflow{value: val, isSet: true}
}

func (v NullableWorkflow) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWorkflow) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
