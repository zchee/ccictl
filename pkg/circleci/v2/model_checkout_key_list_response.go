// Copyright 2020 The ccictl Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package v2

import (
	json "github.com/goccy/go-json"
)

// CheckoutKeyListResponse struct for CheckoutKeyListResponse
type CheckoutKeyListResponse struct {
	Items []CheckoutKey `json:"items"`
	// A token to pass as a `page-token` query parameter to return the next page of results.
	NextPageToken string `json:"next_page_token"`
}

// NewCheckoutKeyListResponse instantiates a new CheckoutKeyListResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCheckoutKeyListResponse(items []CheckoutKey, nextPageToken string) *CheckoutKeyListResponse {
	this := CheckoutKeyListResponse{}
	this.Items = items
	this.NextPageToken = nextPageToken
	return &this
}

// NewCheckoutKeyListResponseWithDefaults instantiates a new CheckoutKeyListResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCheckoutKeyListResponseWithDefaults() *CheckoutKeyListResponse {
	this := CheckoutKeyListResponse{}
	return &this
}

// GetItems returns the Items field value
func (o *CheckoutKeyListResponse) GetItems() []CheckoutKey {
	if o == nil {
		var ret []CheckoutKey
		return ret
	}

	return o.Items
}

// GetItemsOk returns a tuple with the Items field value
// and a boolean to check if the value has been set.
func (o *CheckoutKeyListResponse) GetItemsOk() (*[]CheckoutKey, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Items, true
}

// SetItems sets field value
func (o *CheckoutKeyListResponse) SetItems(v []CheckoutKey) {
	o.Items = v
}

// GetNextPageToken returns the NextPageToken field value
func (o *CheckoutKeyListResponse) GetNextPageToken() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.NextPageToken
}

// GetNextPageTokenOk returns a tuple with the NextPageToken field value
// and a boolean to check if the value has been set.
func (o *CheckoutKeyListResponse) GetNextPageTokenOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.NextPageToken, true
}

// SetNextPageToken sets field value
func (o *CheckoutKeyListResponse) SetNextPageToken(v string) {
	o.NextPageToken = v
}

func (o CheckoutKeyListResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["items"] = o.Items
	}
	if true {
		toSerialize["next_page_token"] = o.NextPageToken
	}
	return json.Marshal(toSerialize)
}

type NullableCheckoutKeyListResponse struct {
	value *CheckoutKeyListResponse
	isSet bool
}

func (v NullableCheckoutKeyListResponse) Get() *CheckoutKeyListResponse {
	return v.value
}

func (v *NullableCheckoutKeyListResponse) Set(val *CheckoutKeyListResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableCheckoutKeyListResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableCheckoutKeyListResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCheckoutKeyListResponse(val *CheckoutKeyListResponse) *NullableCheckoutKeyListResponse {
	return &NullableCheckoutKeyListResponse{value: val, isSet: true}
}

func (v NullableCheckoutKeyListResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCheckoutKeyListResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}