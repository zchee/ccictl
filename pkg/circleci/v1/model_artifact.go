// Copyright 2020 The ccictl Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package v1

import (
	json "github.com/goccy/go-json"
)

// Artifact struct for Artifact
type Artifact struct {
	NodeIndex  *int32  `json:"node_index,omitempty"`
	Path       *string `json:"path,omitempty"`
	PrettyPath *string `json:"pretty_path,omitempty"`
	Url        *string `json:"url,omitempty"`
}

// NewArtifact instantiates a new Artifact object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewArtifact() *Artifact {
	this := Artifact{}
	return &this
}

// NewArtifactWithDefaults instantiates a new Artifact object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewArtifactWithDefaults() *Artifact {
	this := Artifact{}
	return &this
}

// GetNodeIndex returns the NodeIndex field value if set, zero value otherwise.
func (o *Artifact) GetNodeIndex() int32 {
	if o == nil || o.NodeIndex == nil {
		var ret int32
		return ret
	}
	return *o.NodeIndex
}

// GetNodeIndexOk returns a tuple with the NodeIndex field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Artifact) GetNodeIndexOk() (*int32, bool) {
	if o == nil || o.NodeIndex == nil {
		return nil, false
	}
	return o.NodeIndex, true
}

// HasNodeIndex returns a boolean if a field has been set.
func (o *Artifact) HasNodeIndex() bool {
	if o != nil && o.NodeIndex != nil {
		return true
	}

	return false
}

// SetNodeIndex gets a reference to the given int32 and assigns it to the NodeIndex field.
func (o *Artifact) SetNodeIndex(v int32) {
	o.NodeIndex = &v
}

// GetPath returns the Path field value if set, zero value otherwise.
func (o *Artifact) GetPath() string {
	if o == nil || o.Path == nil {
		var ret string
		return ret
	}
	return *o.Path
}

// GetPathOk returns a tuple with the Path field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Artifact) GetPathOk() (*string, bool) {
	if o == nil || o.Path == nil {
		return nil, false
	}
	return o.Path, true
}

// HasPath returns a boolean if a field has been set.
func (o *Artifact) HasPath() bool {
	if o != nil && o.Path != nil {
		return true
	}

	return false
}

// SetPath gets a reference to the given string and assigns it to the Path field.
func (o *Artifact) SetPath(v string) {
	o.Path = &v
}

// GetPrettyPath returns the PrettyPath field value if set, zero value otherwise.
func (o *Artifact) GetPrettyPath() string {
	if o == nil || o.PrettyPath == nil {
		var ret string
		return ret
	}
	return *o.PrettyPath
}

// GetPrettyPathOk returns a tuple with the PrettyPath field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Artifact) GetPrettyPathOk() (*string, bool) {
	if o == nil || o.PrettyPath == nil {
		return nil, false
	}
	return o.PrettyPath, true
}

// HasPrettyPath returns a boolean if a field has been set.
func (o *Artifact) HasPrettyPath() bool {
	if o != nil && o.PrettyPath != nil {
		return true
	}

	return false
}

// SetPrettyPath gets a reference to the given string and assigns it to the PrettyPath field.
func (o *Artifact) SetPrettyPath(v string) {
	o.PrettyPath = &v
}

// GetUrl returns the Url field value if set, zero value otherwise.
func (o *Artifact) GetUrl() string {
	if o == nil || o.Url == nil {
		var ret string
		return ret
	}
	return *o.Url
}

// GetUrlOk returns a tuple with the Url field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Artifact) GetUrlOk() (*string, bool) {
	if o == nil || o.Url == nil {
		return nil, false
	}
	return o.Url, true
}

// HasUrl returns a boolean if a field has been set.
func (o *Artifact) HasUrl() bool {
	if o != nil && o.Url != nil {
		return true
	}

	return false
}

// SetUrl gets a reference to the given string and assigns it to the Url field.
func (o *Artifact) SetUrl(v string) {
	o.Url = &v
}

func (o Artifact) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.NodeIndex != nil {
		toSerialize["node_index"] = o.NodeIndex
	}
	if o.Path != nil {
		toSerialize["path"] = o.Path
	}
	if o.PrettyPath != nil {
		toSerialize["pretty_path"] = o.PrettyPath
	}
	if o.Url != nil {
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
