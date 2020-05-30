// Copyright 2020 The ccictl Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package v2

import (
	json "github.com/goccy/go-json"
)

// ArtifactListResponse struct for ArtifactListResponse
type ArtifactListResponse struct {
	Items []Artifact `json:"items"`
	// A token to pass as a `page-token` query parameter to return the next page of results.
	NextPageToken string `json:"next_page_token"`
}

// NewArtifactListResponse instantiates a new ArtifactListResponse object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewArtifactListResponse(items []Artifact, nextPageToken string) *ArtifactListResponse {
	this := ArtifactListResponse{}
	this.Items = items
	this.NextPageToken = nextPageToken
	return &this
}

// NewArtifactListResponseWithDefaults instantiates a new ArtifactListResponse object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewArtifactListResponseWithDefaults() *ArtifactListResponse {
	this := ArtifactListResponse{}
	return &this
}

// GetItems returns the Items field value
func (o *ArtifactListResponse) GetItems() []Artifact {
	if o == nil {
		var ret []Artifact
		return ret
	}

	return o.Items
}

// GetItemsOk returns a tuple with the Items field value
// and a boolean to check if the value has been set.
func (o *ArtifactListResponse) GetItemsOk() (*[]Artifact, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Items, true
}

// SetItems sets field value
func (o *ArtifactListResponse) SetItems(v []Artifact) {
	o.Items = v
}

// GetNextPageToken returns the NextPageToken field value
func (o *ArtifactListResponse) GetNextPageToken() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.NextPageToken
}

// GetNextPageTokenOk returns a tuple with the NextPageToken field value
// and a boolean to check if the value has been set.
func (o *ArtifactListResponse) GetNextPageTokenOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.NextPageToken, true
}

// SetNextPageToken sets field value
func (o *ArtifactListResponse) SetNextPageToken(v string) {
	o.NextPageToken = v
}

func (o ArtifactListResponse) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["items"] = o.Items
	}
	if true {
		toSerialize["next_page_token"] = o.NextPageToken
	}
	return json.Marshal(toSerialize)
}

type NullableArtifactListResponse struct {
	value *ArtifactListResponse
	isSet bool
}

func (v NullableArtifactListResponse) Get() *ArtifactListResponse {
	return v.value
}

func (v *NullableArtifactListResponse) Set(val *ArtifactListResponse) {
	v.value = val
	v.isSet = true
}

func (v NullableArtifactListResponse) IsSet() bool {
	return v.isSet
}

func (v *NullableArtifactListResponse) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableArtifactListResponse(val *ArtifactListResponse) *NullableArtifactListResponse {
	return &NullableArtifactListResponse{value: val, isSet: true}
}

func (v NullableArtifactListResponse) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableArtifactListResponse) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
