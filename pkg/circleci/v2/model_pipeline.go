// Copyright 2020 The ccictl Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package v2

import (
	"time"

	json "github.com/goccy/go-json"
)

// Pipeline NOTE: The definition of pipeline is subject to change.
type Pipeline struct {
	// The unique ID of the pipeline.
	Id string `json:"id"`
	// A sequence of errors that have occurred within the pipeline.
	Errors []PipelineErrors `json:"errors"`
	// The project-slug for the pipeline.
	ProjectSlug string `json:"project_slug"`
	// The date and time the pipeline was last updated.
	UpdatedAt *time.Time `json:"updated_at,omitempty"`
	// The number of the pipeline.
	Number int64 `json:"number"`
	// The current state of the pipeline.
	State string `json:"state"`
	// The date and time the pipeline was created.
	CreatedAt time.Time       `json:"created_at"`
	Trigger   PipelineTrigger `json:"trigger"`
	Vcs       *PipelineVcs    `json:"vcs,omitempty"`
}

// NewPipeline instantiates a new Pipeline object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPipeline(id string, errors []PipelineErrors, projectSlug string, number int64, state string, createdAt time.Time, trigger PipelineTrigger) *Pipeline {
	this := Pipeline{}
	this.Id = id
	this.Errors = errors
	this.ProjectSlug = projectSlug
	this.Number = number
	this.State = state
	this.CreatedAt = createdAt
	this.Trigger = trigger
	return &this
}

// NewPipelineWithDefaults instantiates a new Pipeline object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPipelineWithDefaults() *Pipeline {
	this := Pipeline{}
	return &this
}

// GetId returns the Id field value
func (o *Pipeline) GetId() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *Pipeline) GetIdOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *Pipeline) SetId(v string) {
	o.Id = v
}

// GetErrors returns the Errors field value
func (o *Pipeline) GetErrors() []PipelineErrors {
	if o == nil {
		var ret []PipelineErrors
		return ret
	}

	return o.Errors
}

// GetErrorsOk returns a tuple with the Errors field value
// and a boolean to check if the value has been set.
func (o *Pipeline) GetErrorsOk() (*[]PipelineErrors, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Errors, true
}

// SetErrors sets field value
func (o *Pipeline) SetErrors(v []PipelineErrors) {
	o.Errors = v
}

// GetProjectSlug returns the ProjectSlug field value
func (o *Pipeline) GetProjectSlug() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ProjectSlug
}

// GetProjectSlugOk returns a tuple with the ProjectSlug field value
// and a boolean to check if the value has been set.
func (o *Pipeline) GetProjectSlugOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ProjectSlug, true
}

// SetProjectSlug sets field value
func (o *Pipeline) SetProjectSlug(v string) {
	o.ProjectSlug = v
}

// GetUpdatedAt returns the UpdatedAt field value if set, zero value otherwise.
func (o *Pipeline) GetUpdatedAt() time.Time {
	if o == nil || o.UpdatedAt == nil {
		var ret time.Time
		return ret
	}
	return *o.UpdatedAt
}

// GetUpdatedAtOk returns a tuple with the UpdatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Pipeline) GetUpdatedAtOk() (*time.Time, bool) {
	if o == nil || o.UpdatedAt == nil {
		return nil, false
	}
	return o.UpdatedAt, true
}

// HasUpdatedAt returns a boolean if a field has been set.
func (o *Pipeline) HasUpdatedAt() bool {
	if o != nil && o.UpdatedAt != nil {
		return true
	}

	return false
}

// SetUpdatedAt gets a reference to the given time.Time and assigns it to the UpdatedAt field.
func (o *Pipeline) SetUpdatedAt(v time.Time) {
	o.UpdatedAt = &v
}

// GetNumber returns the Number field value
func (o *Pipeline) GetNumber() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.Number
}

// GetNumberOk returns a tuple with the Number field value
// and a boolean to check if the value has been set.
func (o *Pipeline) GetNumberOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Number, true
}

// SetNumber sets field value
func (o *Pipeline) SetNumber(v int64) {
	o.Number = v
}

// GetState returns the State field value
func (o *Pipeline) GetState() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.State
}

// GetStateOk returns a tuple with the State field value
// and a boolean to check if the value has been set.
func (o *Pipeline) GetStateOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.State, true
}

// SetState sets field value
func (o *Pipeline) SetState(v string) {
	o.State = v
}

// GetCreatedAt returns the CreatedAt field value
func (o *Pipeline) GetCreatedAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value
// and a boolean to check if the value has been set.
func (o *Pipeline) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreatedAt, true
}

// SetCreatedAt sets field value
func (o *Pipeline) SetCreatedAt(v time.Time) {
	o.CreatedAt = v
}

// GetTrigger returns the Trigger field value
func (o *Pipeline) GetTrigger() PipelineTrigger {
	if o == nil {
		var ret PipelineTrigger
		return ret
	}

	return o.Trigger
}

// GetTriggerOk returns a tuple with the Trigger field value
// and a boolean to check if the value has been set.
func (o *Pipeline) GetTriggerOk() (*PipelineTrigger, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Trigger, true
}

// SetTrigger sets field value
func (o *Pipeline) SetTrigger(v PipelineTrigger) {
	o.Trigger = v
}

// GetVcs returns the Vcs field value if set, zero value otherwise.
func (o *Pipeline) GetVcs() PipelineVcs {
	if o == nil || o.Vcs == nil {
		var ret PipelineVcs
		return ret
	}
	return *o.Vcs
}

// GetVcsOk returns a tuple with the Vcs field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Pipeline) GetVcsOk() (*PipelineVcs, bool) {
	if o == nil || o.Vcs == nil {
		return nil, false
	}
	return o.Vcs, true
}

// HasVcs returns a boolean if a field has been set.
func (o *Pipeline) HasVcs() bool {
	if o != nil && o.Vcs != nil {
		return true
	}

	return false
}

// SetVcs gets a reference to the given PipelineVcs and assigns it to the Vcs field.
func (o *Pipeline) SetVcs(v PipelineVcs) {
	o.Vcs = &v
}

func (o Pipeline) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["id"] = o.Id
	}
	if true {
		toSerialize["errors"] = o.Errors
	}
	if true {
		toSerialize["project_slug"] = o.ProjectSlug
	}
	if o.UpdatedAt != nil {
		toSerialize["updated_at"] = o.UpdatedAt
	}
	if true {
		toSerialize["number"] = o.Number
	}
	if true {
		toSerialize["state"] = o.State
	}
	if true {
		toSerialize["created_at"] = o.CreatedAt
	}
	if true {
		toSerialize["trigger"] = o.Trigger
	}
	if o.Vcs != nil {
		toSerialize["vcs"] = o.Vcs
	}
	return json.Marshal(toSerialize)
}

type NullablePipeline struct {
	value *Pipeline
	isSet bool
}

func (v NullablePipeline) Get() *Pipeline {
	return v.value
}

func (v *NullablePipeline) Set(val *Pipeline) {
	v.value = val
	v.isSet = true
}

func (v NullablePipeline) IsSet() bool {
	return v.isSet
}

func (v *NullablePipeline) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePipeline(val *Pipeline) *NullablePipeline {
	return &NullablePipeline{value: val, isSet: true}
}

func (v NullablePipeline) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePipeline) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
