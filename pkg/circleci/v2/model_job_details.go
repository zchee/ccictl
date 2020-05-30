// Copyright 2020 The ccictl Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package v2

import (
	"time"

	json "github.com/goccy/go-json"
)

// JobDetails Job Details
type JobDetails struct {
	// URL of the job in CircleCI Web UI.
	WebUrl  string            `json:"web_url"`
	Project JobDetailsProject `json:"project"`
	// Info about parallels runs and their status.
	ParallelRuns []JobDetailsParallelRuns `json:"parallel_runs"`
	// The date and time the job started.
	StartedAt      time.Time                `json:"started_at"`
	LatestWorkflow JobDetailsLatestWorkflow `json:"latest_workflow"`
	// The name of the job.
	Name     string             `json:"name"`
	Executor JobDetailsExecutor `json:"executor"`
	// A number of parallel runs the job has.
	Parallelism int64 `json:"parallelism"`
	// The current status of the job.
	Status interface{} `json:"status"`
	// The number of the job.
	Number   int64              `json:"number"`
	Pipeline JobDetailsPipeline `json:"pipeline"`
	// Duration of a job in milliseconds.
	Duration int64 `json:"duration"`
	// The time when the job was created.
	CreatedAt time.Time `json:"created_at"`
	// Messages from CircleCI execution platform.
	Messages []JobDetailsMessages `json:"messages"`
	// List of contexts used by the job.
	Contexts     []JobDetailsContexts   `json:"contexts"`
	Organization JobDetailsOrganization `json:"organization"`
	// The time when the job was placed in a queue.
	QueuedAt time.Time `json:"queued_at"`
	// The time when the job stopped.
	StoppedAt *time.Time `json:"stopped_at,omitempty"`
}

// NewJobDetails instantiates a new JobDetails object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewJobDetails(webUrl string, project JobDetailsProject, parallelRuns []JobDetailsParallelRuns, startedAt time.Time, latestWorkflow JobDetailsLatestWorkflow, name string, executor JobDetailsExecutor, parallelism int64, status interface{}, number int64, pipeline JobDetailsPipeline, duration int64, createdAt time.Time, messages []JobDetailsMessages, contexts []JobDetailsContexts, organization JobDetailsOrganization, queuedAt time.Time) *JobDetails {
	this := JobDetails{}
	this.WebUrl = webUrl
	this.Project = project
	this.ParallelRuns = parallelRuns
	this.StartedAt = startedAt
	this.LatestWorkflow = latestWorkflow
	this.Name = name
	this.Executor = executor
	this.Parallelism = parallelism
	this.Status = status
	this.Number = number
	this.Pipeline = pipeline
	this.Duration = duration
	this.CreatedAt = createdAt
	this.Messages = messages
	this.Contexts = contexts
	this.Organization = organization
	this.QueuedAt = queuedAt
	return &this
}

// NewJobDetailsWithDefaults instantiates a new JobDetails object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewJobDetailsWithDefaults() *JobDetails {
	this := JobDetails{}
	return &this
}

// GetWebUrl returns the WebUrl field value
func (o *JobDetails) GetWebUrl() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.WebUrl
}

// GetWebUrlOk returns a tuple with the WebUrl field value
// and a boolean to check if the value has been set.
func (o *JobDetails) GetWebUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.WebUrl, true
}

// SetWebUrl sets field value
func (o *JobDetails) SetWebUrl(v string) {
	o.WebUrl = v
}

// GetProject returns the Project field value
func (o *JobDetails) GetProject() JobDetailsProject {
	if o == nil {
		var ret JobDetailsProject
		return ret
	}

	return o.Project
}

// GetProjectOk returns a tuple with the Project field value
// and a boolean to check if the value has been set.
func (o *JobDetails) GetProjectOk() (*JobDetailsProject, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Project, true
}

// SetProject sets field value
func (o *JobDetails) SetProject(v JobDetailsProject) {
	o.Project = v
}

// GetParallelRuns returns the ParallelRuns field value
func (o *JobDetails) GetParallelRuns() []JobDetailsParallelRuns {
	if o == nil {
		var ret []JobDetailsParallelRuns
		return ret
	}

	return o.ParallelRuns
}

// GetParallelRunsOk returns a tuple with the ParallelRuns field value
// and a boolean to check if the value has been set.
func (o *JobDetails) GetParallelRunsOk() (*[]JobDetailsParallelRuns, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ParallelRuns, true
}

// SetParallelRuns sets field value
func (o *JobDetails) SetParallelRuns(v []JobDetailsParallelRuns) {
	o.ParallelRuns = v
}

// GetStartedAt returns the StartedAt field value
func (o *JobDetails) GetStartedAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.StartedAt
}

// GetStartedAtOk returns a tuple with the StartedAt field value
// and a boolean to check if the value has been set.
func (o *JobDetails) GetStartedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.StartedAt, true
}

// SetStartedAt sets field value
func (o *JobDetails) SetStartedAt(v time.Time) {
	o.StartedAt = v
}

// GetLatestWorkflow returns the LatestWorkflow field value
func (o *JobDetails) GetLatestWorkflow() JobDetailsLatestWorkflow {
	if o == nil {
		var ret JobDetailsLatestWorkflow
		return ret
	}

	return o.LatestWorkflow
}

// GetLatestWorkflowOk returns a tuple with the LatestWorkflow field value
// and a boolean to check if the value has been set.
func (o *JobDetails) GetLatestWorkflowOk() (*JobDetailsLatestWorkflow, bool) {
	if o == nil {
		return nil, false
	}
	return &o.LatestWorkflow, true
}

// SetLatestWorkflow sets field value
func (o *JobDetails) SetLatestWorkflow(v JobDetailsLatestWorkflow) {
	o.LatestWorkflow = v
}

// GetName returns the Name field value
func (o *JobDetails) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *JobDetails) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *JobDetails) SetName(v string) {
	o.Name = v
}

// GetExecutor returns the Executor field value
func (o *JobDetails) GetExecutor() JobDetailsExecutor {
	if o == nil {
		var ret JobDetailsExecutor
		return ret
	}

	return o.Executor
}

// GetExecutorOk returns a tuple with the Executor field value
// and a boolean to check if the value has been set.
func (o *JobDetails) GetExecutorOk() (*JobDetailsExecutor, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Executor, true
}

// SetExecutor sets field value
func (o *JobDetails) SetExecutor(v JobDetailsExecutor) {
	o.Executor = v
}

// GetParallelism returns the Parallelism field value
func (o *JobDetails) GetParallelism() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.Parallelism
}

// GetParallelismOk returns a tuple with the Parallelism field value
// and a boolean to check if the value has been set.
func (o *JobDetails) GetParallelismOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Parallelism, true
}

// SetParallelism sets field value
func (o *JobDetails) SetParallelism(v int64) {
	o.Parallelism = v
}

// GetStatus returns the Status field value
// If the value is explicit nil, the zero value for interface{} will be returned
func (o *JobDetails) GetStatus() interface{} {
	if o == nil {
		var ret interface{}
		return ret
	}

	return o.Status
}

// GetStatusOk returns a tuple with the Status field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *JobDetails) GetStatusOk() (*interface{}, bool) {
	if o == nil || o.Status == nil {
		return nil, false
	}
	return &o.Status, true
}

// SetStatus sets field value
func (o *JobDetails) SetStatus(v interface{}) {
	o.Status = v
}

// GetNumber returns the Number field value
func (o *JobDetails) GetNumber() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.Number
}

// GetNumberOk returns a tuple with the Number field value
// and a boolean to check if the value has been set.
func (o *JobDetails) GetNumberOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Number, true
}

// SetNumber sets field value
func (o *JobDetails) SetNumber(v int64) {
	o.Number = v
}

// GetPipeline returns the Pipeline field value
func (o *JobDetails) GetPipeline() JobDetailsPipeline {
	if o == nil {
		var ret JobDetailsPipeline
		return ret
	}

	return o.Pipeline
}

// GetPipelineOk returns a tuple with the Pipeline field value
// and a boolean to check if the value has been set.
func (o *JobDetails) GetPipelineOk() (*JobDetailsPipeline, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Pipeline, true
}

// SetPipeline sets field value
func (o *JobDetails) SetPipeline(v JobDetailsPipeline) {
	o.Pipeline = v
}

// GetDuration returns the Duration field value
func (o *JobDetails) GetDuration() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.Duration
}

// GetDurationOk returns a tuple with the Duration field value
// and a boolean to check if the value has been set.
func (o *JobDetails) GetDurationOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Duration, true
}

// SetDuration sets field value
func (o *JobDetails) SetDuration(v int64) {
	o.Duration = v
}

// GetCreatedAt returns the CreatedAt field value
func (o *JobDetails) GetCreatedAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value
// and a boolean to check if the value has been set.
func (o *JobDetails) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreatedAt, true
}

// SetCreatedAt sets field value
func (o *JobDetails) SetCreatedAt(v time.Time) {
	o.CreatedAt = v
}

// GetMessages returns the Messages field value
func (o *JobDetails) GetMessages() []JobDetailsMessages {
	if o == nil {
		var ret []JobDetailsMessages
		return ret
	}

	return o.Messages
}

// GetMessagesOk returns a tuple with the Messages field value
// and a boolean to check if the value has been set.
func (o *JobDetails) GetMessagesOk() (*[]JobDetailsMessages, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Messages, true
}

// SetMessages sets field value
func (o *JobDetails) SetMessages(v []JobDetailsMessages) {
	o.Messages = v
}

// GetContexts returns the Contexts field value
func (o *JobDetails) GetContexts() []JobDetailsContexts {
	if o == nil {
		var ret []JobDetailsContexts
		return ret
	}

	return o.Contexts
}

// GetContextsOk returns a tuple with the Contexts field value
// and a boolean to check if the value has been set.
func (o *JobDetails) GetContextsOk() (*[]JobDetailsContexts, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Contexts, true
}

// SetContexts sets field value
func (o *JobDetails) SetContexts(v []JobDetailsContexts) {
	o.Contexts = v
}

// GetOrganization returns the Organization field value
func (o *JobDetails) GetOrganization() JobDetailsOrganization {
	if o == nil {
		var ret JobDetailsOrganization
		return ret
	}

	return o.Organization
}

// GetOrganizationOk returns a tuple with the Organization field value
// and a boolean to check if the value has been set.
func (o *JobDetails) GetOrganizationOk() (*JobDetailsOrganization, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Organization, true
}

// SetOrganization sets field value
func (o *JobDetails) SetOrganization(v JobDetailsOrganization) {
	o.Organization = v
}

// GetQueuedAt returns the QueuedAt field value
func (o *JobDetails) GetQueuedAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.QueuedAt
}

// GetQueuedAtOk returns a tuple with the QueuedAt field value
// and a boolean to check if the value has been set.
func (o *JobDetails) GetQueuedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.QueuedAt, true
}

// SetQueuedAt sets field value
func (o *JobDetails) SetQueuedAt(v time.Time) {
	o.QueuedAt = v
}

// GetStoppedAt returns the StoppedAt field value if set, zero value otherwise.
func (o *JobDetails) GetStoppedAt() time.Time {
	if o == nil || o.StoppedAt == nil {
		var ret time.Time
		return ret
	}
	return *o.StoppedAt
}

// GetStoppedAtOk returns a tuple with the StoppedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *JobDetails) GetStoppedAtOk() (*time.Time, bool) {
	if o == nil || o.StoppedAt == nil {
		return nil, false
	}
	return o.StoppedAt, true
}

// HasStoppedAt returns a boolean if a field has been set.
func (o *JobDetails) HasStoppedAt() bool {
	if o != nil && o.StoppedAt != nil {
		return true
	}

	return false
}

// SetStoppedAt gets a reference to the given time.Time and assigns it to the StoppedAt field.
func (o *JobDetails) SetStoppedAt(v time.Time) {
	o.StoppedAt = &v
}

func (o JobDetails) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["web_url"] = o.WebUrl
	}
	if true {
		toSerialize["project"] = o.Project
	}
	if true {
		toSerialize["parallel_runs"] = o.ParallelRuns
	}
	if true {
		toSerialize["started_at"] = o.StartedAt
	}
	if true {
		toSerialize["latest_workflow"] = o.LatestWorkflow
	}
	if true {
		toSerialize["name"] = o.Name
	}
	if true {
		toSerialize["executor"] = o.Executor
	}
	if true {
		toSerialize["parallelism"] = o.Parallelism
	}
	if o.Status != nil {
		toSerialize["status"] = o.Status
	}
	if true {
		toSerialize["number"] = o.Number
	}
	if true {
		toSerialize["pipeline"] = o.Pipeline
	}
	if true {
		toSerialize["duration"] = o.Duration
	}
	if true {
		toSerialize["created_at"] = o.CreatedAt
	}
	if true {
		toSerialize["messages"] = o.Messages
	}
	if true {
		toSerialize["contexts"] = o.Contexts
	}
	if true {
		toSerialize["organization"] = o.Organization
	}
	if true {
		toSerialize["queued_at"] = o.QueuedAt
	}
	if o.StoppedAt != nil {
		toSerialize["stopped_at"] = o.StoppedAt
	}
	return json.Marshal(toSerialize)
}

type NullableJobDetails struct {
	value *JobDetails
	isSet bool
}

func (v NullableJobDetails) Get() *JobDetails {
	return v.value
}

func (v *NullableJobDetails) Set(val *JobDetails) {
	v.value = val
	v.isSet = true
}

func (v NullableJobDetails) IsSet() bool {
	return v.isSet
}

func (v *NullableJobDetails) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableJobDetails(val *JobDetails) *NullableJobDetails {
	return &NullableJobDetails{value: val, isSet: true}
}

func (v NullableJobDetails) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableJobDetails) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
