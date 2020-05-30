// Copyright 2020 The ccictl Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package v2

import (
	json "github.com/goccy/go-json"
)

// Artifact An artifact
type Artifact struct {
	// The artifact path.
	Path string `json:"path"`
	// The index of the node that stored the artifact.
	NodeIndex int64 `json:"node_index"`
	// The URL to download the artifact contents.
	Url string `json:"url"`
}

// NewArtifact instantiates a new Artifact object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewArtifact(path string, nodeIndex int64, url string) *Artifact {
	this := Artifact{}
	this.Path = path
	this.NodeIndex = nodeIndex
	this.Url = url
	return &this
}

// NewArtifactWithDefaults instantiates a new Artifact object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewArtifactWithDefaults() *Artifact {
	this := Artifact{}
	return &this
}

// GetPath returns the Path field value
func (o *Artifact) GetPath() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Path
}

// GetPathOk returns a tuple with the Path field value
// and a boolean to check if the value has been set.
func (o *Artifact) GetPathOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Path, true
}

// SetPath sets field value
func (o *Artifact) SetPath(v string) {
	o.Path = v
}

// GetNodeIndex returns the NodeIndex field value
func (o *Artifact) GetNodeIndex() int64 {
	if o == nil {
		var ret int64
		return ret
	}

	return o.NodeIndex
}

// GetNodeIndexOk returns a tuple with the NodeIndex field value
// and a boolean to check if the value has been set.
func (o *Artifact) GetNodeIndexOk() (*int64, bool) {
	if o == nil {
		return nil, false
	}
	return &o.NodeIndex, true
}

// SetNodeIndex sets field value
func (o *Artifact) SetNodeIndex(v int64) {
	o.NodeIndex = v
}

// GetUrl returns the Url field value
func (o *Artifact) GetUrl() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Url
}

// GetUrlOk returns a tuple with the Url field value
// and a boolean to check if the value has been set.
func (o *Artifact) GetUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Url, true
}

// SetUrl sets field value
func (o *Artifact) SetUrl(v string) {
	o.Url = v
}

func (o Artifact) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["path"] = o.Path
	}
	if true {
		toSerialize["node_index"] = o.NodeIndex
	}
	if true {
		toSerialize["url"] = o.Url
	}
	return json.Marshal(toSerialize)
}

type NullableArtifact struct {
	value *Artifact
	isSet bool
}

func (v NullableArtifact) Get() *Artifact {
	return v.value
}

func (v *NullableArtifact) Set(val *Artifact) {
	v.value = val
	v.isSet = true
}

func (v NullableArtifact) IsSet() bool {
	return v.isSet
}

func (v *NullableArtifact) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableArtifact(val *Artifact) *NullableArtifact {
	return &NullableArtifact{value: val, isSet: true}
}

func (v NullableArtifact) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableArtifact) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
