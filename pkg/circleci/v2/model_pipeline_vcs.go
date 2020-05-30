// Copyright 2020 The ccictl Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package v2

import (
	json "github.com/goccy/go-json"
)

// PipelineVcs VCS information for the pipeline.
type PipelineVcs struct {
	// Name of the VCS provider (e.g. GitHub, Bitbucket).
	ProviderName string `json:"provider_name"`
	// URL for the repository where the trigger originated. For fork-PR pipelines, this is the URL to the fork. For other pipelines the `origin_` and `target_repository_url`s will be the same.
	OriginRepositoryUrl string `json:"origin_repository_url"`
	// URL for the repository the trigger targets (i.e. the repository where the PR will be merged). For fork-PR pipelines, this is the URL to the parent repo. For other pipelines, the `origin_` and `target_repository_url`s will be the same.
	TargetRepositoryUrl string `json:"target_repository_url"`
	// The code revision the pipeline ran.
	Revision string `json:"revision"`
	// The branch where the pipeline ran. The HEAD commit on this branch was used for the pipeline. Note that `branch` and `tag` are mutually exclusive.
	Branch *string `json:"branch,omitempty"`
	// The tag used by the pipeline. The commit that this tag points to was used for the pipeline. Note that `branch` and `tag` are mutually exclusive.
	Tag    *string            `json:"tag,omitempty"`
	Commit *PipelineVcsCommit `json:"commit,omitempty"`
}

// NewPipelineVcs instantiates a new PipelineVcs object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPipelineVcs(providerName string, originRepositoryUrl string, targetRepositoryUrl string, revision string) *PipelineVcs {
	this := PipelineVcs{}
	this.ProviderName = providerName
	this.OriginRepositoryUrl = originRepositoryUrl
	this.TargetRepositoryUrl = targetRepositoryUrl
	this.Revision = revision
	return &this
}

// NewPipelineVcsWithDefaults instantiates a new PipelineVcs object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPipelineVcsWithDefaults() *PipelineVcs {
	this := PipelineVcs{}
	return &this
}

// GetProviderName returns the ProviderName field value
func (o *PipelineVcs) GetProviderName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.ProviderName
}

// GetProviderNameOk returns a tuple with the ProviderName field value
// and a boolean to check if the value has been set.
func (o *PipelineVcs) GetProviderNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.ProviderName, true
}

// SetProviderName sets field value
func (o *PipelineVcs) SetProviderName(v string) {
	o.ProviderName = v
}

// GetOriginRepositoryUrl returns the OriginRepositoryUrl field value
func (o *PipelineVcs) GetOriginRepositoryUrl() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.OriginRepositoryUrl
}

// GetOriginRepositoryUrlOk returns a tuple with the OriginRepositoryUrl field value
// and a boolean to check if the value has been set.
func (o *PipelineVcs) GetOriginRepositoryUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.OriginRepositoryUrl, true
}

// SetOriginRepositoryUrl sets field value
func (o *PipelineVcs) SetOriginRepositoryUrl(v string) {
	o.OriginRepositoryUrl = v
}

// GetTargetRepositoryUrl returns the TargetRepositoryUrl field value
func (o *PipelineVcs) GetTargetRepositoryUrl() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.TargetRepositoryUrl
}

// GetTargetRepositoryUrlOk returns a tuple with the TargetRepositoryUrl field value
// and a boolean to check if the value has been set.
func (o *PipelineVcs) GetTargetRepositoryUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.TargetRepositoryUrl, true
}

// SetTargetRepositoryUrl sets field value
func (o *PipelineVcs) SetTargetRepositoryUrl(v string) {
	o.TargetRepositoryUrl = v
}

// GetRevision returns the Revision field value
func (o *PipelineVcs) GetRevision() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Revision
}

// GetRevisionOk returns a tuple with the Revision field value
// and a boolean to check if the value has been set.
func (o *PipelineVcs) GetRevisionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Revision, true
}

// SetRevision sets field value
func (o *PipelineVcs) SetRevision(v string) {
	o.Revision = v
}

// GetBranch returns the Branch field value if set, zero value otherwise.
func (o *PipelineVcs) GetBranch() string {
	if o == nil || o.Branch == nil {
		var ret string
		return ret
	}
	return *o.Branch
}

// GetBranchOk returns a tuple with the Branch field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PipelineVcs) GetBranchOk() (*string, bool) {
	if o == nil || o.Branch == nil {
		return nil, false
	}
	return o.Branch, true
}

// HasBranch returns a boolean if a field has been set.
func (o *PipelineVcs) HasBranch() bool {
	if o != nil && o.Branch != nil {
		return true
	}

	return false
}

// SetBranch gets a reference to the given string and assigns it to the Branch field.
func (o *PipelineVcs) SetBranch(v string) {
	o.Branch = &v
}

// GetTag returns the Tag field value if set, zero value otherwise.
func (o *PipelineVcs) GetTag() string {
	if o == nil || o.Tag == nil {
		var ret string
		return ret
	}
	return *o.Tag
}

// GetTagOk returns a tuple with the Tag field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PipelineVcs) GetTagOk() (*string, bool) {
	if o == nil || o.Tag == nil {
		return nil, false
	}
	return o.Tag, true
}

// HasTag returns a boolean if a field has been set.
func (o *PipelineVcs) HasTag() bool {
	if o != nil && o.Tag != nil {
		return true
	}

	return false
}

// SetTag gets a reference to the given string and assigns it to the Tag field.
func (o *PipelineVcs) SetTag(v string) {
	o.Tag = &v
}

// GetCommit returns the Commit field value if set, zero value otherwise.
func (o *PipelineVcs) GetCommit() PipelineVcsCommit {
	if o == nil || o.Commit == nil {
		var ret PipelineVcsCommit
		return ret
	}
	return *o.Commit
}

// GetCommitOk returns a tuple with the Commit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PipelineVcs) GetCommitOk() (*PipelineVcsCommit, bool) {
	if o == nil || o.Commit == nil {
		return nil, false
	}
	return o.Commit, true
}

// HasCommit returns a boolean if a field has been set.
func (o *PipelineVcs) HasCommit() bool {
	if o != nil && o.Commit != nil {
		return true
	}

	return false
}

// SetCommit gets a reference to the given PipelineVcsCommit and assigns it to the Commit field.
func (o *PipelineVcs) SetCommit(v PipelineVcsCommit) {
	o.Commit = &v
}

func (o PipelineVcs) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["provider_name"] = o.ProviderName
	}
	if true {
		toSerialize["origin_repository_url"] = o.OriginRepositoryUrl
	}
	if true {
		toSerialize["target_repository_url"] = o.TargetRepositoryUrl
	}
	if true {
		toSerialize["revision"] = o.Revision
	}
	if o.Branch != nil {
		toSerialize["branch"] = o.Branch
	}
	if o.Tag != nil {
		toSerialize["tag"] = o.Tag
	}
	if o.Commit != nil {
		toSerialize["commit"] = o.Commit
	}
	return json.Marshal(toSerialize)
}

type NullablePipelineVcs struct {
	value *PipelineVcs
	isSet bool
}

func (v NullablePipelineVcs) Get() *PipelineVcs {
	return v.value
}

func (v *NullablePipelineVcs) Set(val *PipelineVcs) {
	v.value = val
	v.isSet = true
}

func (v NullablePipelineVcs) IsSet() bool {
	return v.isSet
}

func (v *NullablePipelineVcs) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePipelineVcs(val *PipelineVcs) *NullablePipelineVcs {
	return &NullablePipelineVcs{value: val, isSet: true}
}

func (v NullablePipelineVcs) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePipelineVcs) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
