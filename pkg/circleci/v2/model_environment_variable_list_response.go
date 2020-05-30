// Copyright 2020 The ccictl Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package v2

import (
	json "github.com/goccy/go-json"
)

// EnvironmentVariableListResponse struct for EnvironmentVariableListResponse
type EnvironmentVariableListResponse struct {
	Items []EnvironmentVariablePair `json:"items"`
	// A token to pass as a `page-token` query parameter to return the next page of results.
	NextPageToken string `json:"next_page_token"`
}

// NewEnvironmentVariableListResponse instantiates a new EnvironmentVariableListResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewEnvironmentVariableListResponse(items []EnvironmentVariablePair, nextPageToken string) *EnvironmentVariableListResponse {
	this := EnvironmentVariableListResponse{}
	this.Items = items
	this.NextPageToken = nextPageToken
	return &this
}

// NewEnvironmentVariableListResponseWithDefaults instantiates a new EnvironmentVariableListResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewEnvironmentVariableListResponseWithDefaults() *EnvironmentVariableListResponse {
	this := EnvironmentVariableListResponse{}
	return &this
}

// GetItems returns the Items field value
func (o *EnvironmentVariableListResponse) GetItems() []EnvironmentVariablePair {
	if o == nil {
		var ret []EnvironmentVariablePair
		return ret
	}

	return o.Items
}

// GetItemsOk returns a tuple with the Items field value
// and a boolean to check if the value has been set.
func (o *EnvironmentVariableListResponse) GetItemsOk() (*[]EnvironmentVariablePair, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Items, true
}

// SetItems sets field value
func (o *EnvironmentVariableListResponse) SetItems(v []EnvironmentVariablePair) {
	o.Items = v
}

// GetNextPageToken returns the NextPageToken field value
func (o *EnvironmentVariableListResponse) GetNextPageToken() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.NextPageToken
}

// GetNextPageTokenOk returns a tuple with the NextPageToken field value
// and a boolean to check if the value has been set.
func (o *EnvironmentVariableListResponse) GetNextPageTokenOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.NextPageToken, true
}

// SetNextPageToken sets field value
func (o *EnvironmentVariableListResponse) SetNextPageToken(v string) {
	o.NextPageToken = v
}

func (o EnvironmentVariableListResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["items"] = o.Items
	}
	if true {
		toSerialize["next_page_token"] = o.NextPageToken
	}
	return json.Marshal(toSerialize)
}

type NullableEnvironmentVariableListResponse struct {
	value *EnvironmentVariableListResponse
	isSet bool
}

func (v NullableEnvironmentVariableListResponse) Get() *EnvironmentVariableListResponse {
	return v.value
}

func (v *NullableEnvironmentVariableListResponse) Set(val *EnvironmentVariableListResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableEnvironmentVariableListResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableEnvironmentVariableListResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableEnvironmentVariableListResponse(val *EnvironmentVariableListResponse) *NullableEnvironmentVariableListResponse {
	return &NullableEnvironmentVariableListResponse{value: val, isSet: true}
}

func (v NullableEnvironmentVariableListResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableEnvironmentVariableListResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
