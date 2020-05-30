// Copyright 2020 The ccictl Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package v2

import (
	json "github.com/goccy/go-json"
)

// WorkflowJobListResponse struct for WorkflowJobListResponse
type WorkflowJobListResponse struct {
	Items []Job `json:"items"`
	// A token to pass as a `page-token` query parameter to return the next page of results.
	NextPageToken string `json:"next_page_token"`
}

// NewWorkflowJobListResponse instantiates a new WorkflowJobListResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWorkflowJobListResponse(items []Job, nextPageToken string) *WorkflowJobListResponse {
	this := WorkflowJobListResponse{}
	this.Items = items
	this.NextPageToken = nextPageToken
	return &this
}

// NewWorkflowJobListResponseWithDefaults instantiates a new WorkflowJobListResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWorkflowJobListResponseWithDefaults() *WorkflowJobListResponse {
	this := WorkflowJobListResponse{}
	return &this
}

// GetItems returns the Items field value
func (o *WorkflowJobListResponse) GetItems() []Job {
	if o == nil {
		var ret []Job
		return ret
	}

	return o.Items
}

// GetItemsOk returns a tuple with the Items field value
// and a boolean to check if the value has been set.
func (o *WorkflowJobListResponse) GetItemsOk() (*[]Job, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Items, true
}

// SetItems sets field value
func (o *WorkflowJobListResponse) SetItems(v []Job) {
	o.Items = v
}

// GetNextPageToken returns the NextPageToken field value
func (o *WorkflowJobListResponse) GetNextPageToken() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.NextPageToken
}

// GetNextPageTokenOk returns a tuple with the NextPageToken field value
// and a boolean to check if the value has been set.
func (o *WorkflowJobListResponse) GetNextPageTokenOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.NextPageToken, true
}

// SetNextPageToken sets field value
func (o *WorkflowJobListResponse) SetNextPageToken(v string) {
	o.NextPageToken = v
}

func (o WorkflowJobListResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["items"] = o.Items
	}
	if true {
		toSerialize["next_page_token"] = o.NextPageToken
	}
	return json.Marshal(toSerialize)
}

type NullableWorkflowJobListResponse struct {
	value *WorkflowJobListResponse
	isSet bool
}

func (v NullableWorkflowJobListResponse) Get() *WorkflowJobListResponse {
	return v.value
}

func (v *NullableWorkflowJobListResponse) Set(val *WorkflowJobListResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableWorkflowJobListResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableWorkflowJobListResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWorkflowJobListResponse(val *WorkflowJobListResponse) *NullableWorkflowJobListResponse {
	return &NullableWorkflowJobListResponse{value: val, isSet: true}
}

func (v NullableWorkflowJobListResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWorkflowJobListResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
