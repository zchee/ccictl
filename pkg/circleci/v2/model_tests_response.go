// Copyright 2020 The ccictl Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package v2

import (
	json "github.com/goccy/go-json"
)

// TestsResponse struct for TestsResponse
type TestsResponse struct {
	Items []TestsResponseItems `json:"items"`
	// A token to pass as a `page-token` query parameter to return the next page of results.
	NextPageToken string `json:"next_page_token"`
}

// NewTestsResponse instantiates a new TestsResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTestsResponse(items []TestsResponseItems, nextPageToken string) *TestsResponse {
	this := TestsResponse{}
	this.Items = items
	this.NextPageToken = nextPageToken
	return &this
}

// NewTestsResponseWithDefaults instantiates a new TestsResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTestsResponseWithDefaults() *TestsResponse {
	this := TestsResponse{}
	return &this
}

// GetItems returns the Items field value
func (o *TestsResponse) GetItems() []TestsResponseItems {
	if o == nil {
		var ret []TestsResponseItems
		return ret
	}

	return o.Items
}

// GetItemsOk returns a tuple with the Items field value
// and a boolean to check if the value has been set.
func (o *TestsResponse) GetItemsOk() (*[]TestsResponseItems, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Items, true
}

// SetItems sets field value
func (o *TestsResponse) SetItems(v []TestsResponseItems) {
	o.Items = v
}

// GetNextPageToken returns the NextPageToken field value
func (o *TestsResponse) GetNextPageToken() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.NextPageToken
}

// GetNextPageTokenOk returns a tuple with the NextPageToken field value
// and a boolean to check if the value has been set.
func (o *TestsResponse) GetNextPageTokenOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.NextPageToken, true
}

// SetNextPageToken sets field value
func (o *TestsResponse) SetNextPageToken(v string) {
	o.NextPageToken = v
}

func (o TestsResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["items"] = o.Items
	}
	if true {
		toSerialize["next_page_token"] = o.NextPageToken
	}
	return json.Marshal(toSerialize)
}

type NullableTestsResponse struct {
	value *TestsResponse
	isSet bool
}

func (v NullableTestsResponse) Get() *TestsResponse {
	return v.value
}

func (v *NullableTestsResponse) Set(val *TestsResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableTestsResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableTestsResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTestsResponse(val *TestsResponse) *NullableTestsResponse {
	return &NullableTestsResponse{value: val, isSet: true}
}

func (v NullableTestsResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTestsResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
