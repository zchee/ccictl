// Copyright 2020 The ccictl Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package v2

import (
	json "github.com/goccy/go-json"
)

// InlineResponse2003 Paginated recent job runs.
type InlineResponse2003 struct {
	// Recent job runs.
	Items []InlineResponse2003Items `json:"items"`
	// A token to pass as a `page-token` query parameter to return the next page of results.
	NextPageToken string `json:"next_page_token"`
}

// NewInlineResponse2003 instantiates a new InlineResponse2003 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineResponse2003(items []InlineResponse2003Items, nextPageToken string) *InlineResponse2003 {
	this := InlineResponse2003{}
	this.Items = items
	this.NextPageToken = nextPageToken
	return &this
}

// NewInlineResponse2003WithDefaults instantiates a new InlineResponse2003 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineResponse2003WithDefaults() *InlineResponse2003 {
	this := InlineResponse2003{}
	return &this
}

// GetItems returns the Items field value
func (o *InlineResponse2003) GetItems() []InlineResponse2003Items {
	if o == nil {
		var ret []InlineResponse2003Items
		return ret
	}

	return o.Items
}

// GetItemsOk returns a tuple with the Items field value
// and a boolean to check if the value has been set.
func (o *InlineResponse2003) GetItemsOk() (*[]InlineResponse2003Items, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Items, true
}

// SetItems sets field value
func (o *InlineResponse2003) SetItems(v []InlineResponse2003Items) {
	o.Items = v
}

// GetNextPageToken returns the NextPageToken field value
func (o *InlineResponse2003) GetNextPageToken() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.NextPageToken
}

// GetNextPageTokenOk returns a tuple with the NextPageToken field value
// and a boolean to check if the value has been set.
func (o *InlineResponse2003) GetNextPageTokenOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.NextPageToken, true
}

// SetNextPageToken sets field value
func (o *InlineResponse2003) SetNextPageToken(v string) {
	o.NextPageToken = v
}

func (o InlineResponse2003) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["items"] = o.Items
	}
	if true {
		toSerialize["next_page_token"] = o.NextPageToken
	}
	return json.Marshal(toSerialize)
}

type NullableInlineResponse2003 struct {
	value *InlineResponse2003
	isSet bool
}

func (v NullableInlineResponse2003) Get() *InlineResponse2003 {
	return v.value
}

func (v *NullableInlineResponse2003) Set(val *InlineResponse2003) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineResponse2003) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineResponse2003) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineResponse2003(val *InlineResponse2003) *NullableInlineResponse2003 {
	return &NullableInlineResponse2003{value: val, isSet: true}
}

func (v NullableInlineResponse2003) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineResponse2003) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
