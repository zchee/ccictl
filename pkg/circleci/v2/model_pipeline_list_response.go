// Copyright 2020 The ccictl Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package v2

import (
	json "github.com/goccy/go-json"
)

// PipelineListResponse List of pipelines
type PipelineListResponse struct {
	Items []Pipeline `json:"items"`
	// A token to pass as a `page-token` query parameter to return the next page of results.
	NextPageToken string `json:"next_page_token"`
}

// NewPipelineListResponse instantiates a new PipelineListResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPipelineListResponse(items []Pipeline, nextPageToken string) *PipelineListResponse {
	this := PipelineListResponse{}
	this.Items = items
	this.NextPageToken = nextPageToken
	return &this
}

// NewPipelineListResponseWithDefaults instantiates a new PipelineListResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPipelineListResponseWithDefaults() *PipelineListResponse {
	this := PipelineListResponse{}
	return &this
}

// GetItems returns the Items field value
func (o *PipelineListResponse) GetItems() []Pipeline {
	if o == nil {
		var ret []Pipeline
		return ret
	}

	return o.Items
}

// GetItemsOk returns a tuple with the Items field value
// and a boolean to check if the value has been set.
func (o *PipelineListResponse) GetItemsOk() (*[]Pipeline, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Items, true
}

// SetItems sets field value
func (o *PipelineListResponse) SetItems(v []Pipeline) {
	o.Items = v
}

// GetNextPageToken returns the NextPageToken field value
func (o *PipelineListResponse) GetNextPageToken() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.NextPageToken
}

// GetNextPageTokenOk returns a tuple with the NextPageToken field value
// and a boolean to check if the value has been set.
func (o *PipelineListResponse) GetNextPageTokenOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.NextPageToken, true
}

// SetNextPageToken sets field value
func (o *PipelineListResponse) SetNextPageToken(v string) {
	o.NextPageToken = v
}

func (o PipelineListResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["items"] = o.Items
	}
	if true {
		toSerialize["next_page_token"] = o.NextPageToken
	}
	return json.Marshal(toSerialize)
}

type NullablePipelineListResponse struct {
	value *PipelineListResponse
	isSet bool
}

func (v NullablePipelineListResponse) Get() *PipelineListResponse {
	return v.value
}

func (v *NullablePipelineListResponse) Set(val *PipelineListResponse) {
	v.value = val
	v.isSet = true
}

func (v NullablePipelineListResponse) IsSet() bool {
	return v.isSet
}

func (v *NullablePipelineListResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePipelineListResponse(val *PipelineListResponse) *NullablePipelineListResponse {
	return &NullablePipelineListResponse{value: val, isSet: true}
}

func (v NullablePipelineListResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePipelineListResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
