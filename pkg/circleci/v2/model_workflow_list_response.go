// Copyright 2020 The ccictl Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package v2

import (
	json "github.com/goccy/go-json"
)

// WorkflowListResponse A list of workflows and associated pagination token.
type WorkflowListResponse struct {
	// A list of workflows.
	Items []Workflow `json:"items"`
	// A token to pass as a `page-token` query parameter to return the next page of results.
	NextPageToken string `json:"next_page_token"`
}

// NewWorkflowListResponse instantiates a new WorkflowListResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWorkflowListResponse(items []Workflow, nextPageToken string) *WorkflowListResponse {
	this := WorkflowListResponse{}
	this.Items = items
	this.NextPageToken = nextPageToken
	return &this
}

// NewWorkflowListResponseWithDefaults instantiates a new WorkflowListResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWorkflowListResponseWithDefaults() *WorkflowListResponse {
	this := WorkflowListResponse{}
	return &this
}

// GetItems returns the Items field value
func (o *WorkflowListResponse) GetItems() []Workflow {
	if o == nil {
		var ret []Workflow
		return ret
	}

	return o.Items
}

// GetItemsOk returns a tuple with the Items field value
// and a boolean to check if the value has been set.
func (o *WorkflowListResponse) GetItemsOk() (*[]Workflow, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Items, true
}

// SetItems sets field value
func (o *WorkflowListResponse) SetItems(v []Workflow) {
	o.Items = v
}

// GetNextPageToken returns the NextPageToken field value
func (o *WorkflowListResponse) GetNextPageToken() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.NextPageToken
}

// GetNextPageTokenOk returns a tuple with the NextPageToken field value
// and a boolean to check if the value has been set.
func (o *WorkflowListResponse) GetNextPageTokenOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.NextPageToken, true
}

// SetNextPageToken sets field value
func (o *WorkflowListResponse) SetNextPageToken(v string) {
	o.NextPageToken = v
}

func (o WorkflowListResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["items"] = o.Items
	}
	if true {
		toSerialize["next_page_token"] = o.NextPageToken
	}
	return json.Marshal(toSerialize)
}

type NullableWorkflowListResponse struct {
	value *WorkflowListResponse
	isSet bool
}

func (v NullableWorkflowListResponse) Get() *WorkflowListResponse {
	return v.value
}

func (v *NullableWorkflowListResponse) Set(val *WorkflowListResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableWorkflowListResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableWorkflowListResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWorkflowListResponse(val *WorkflowListResponse) *NullableWorkflowListResponse {
	return &NullableWorkflowListResponse{value: val, isSet: true}
}

func (v NullableWorkflowListResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWorkflowListResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
