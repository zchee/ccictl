// Copyright 2020 The ccictl Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package v1

import (
	json "github.com/goccy/go-json"
)

// InlineObject2 struct for InlineObject2
type InlineObject2 struct {
	// Additional environment variables to inject into the build environment. A map of names to values.
	BuildParameters *map[string]interface{} `json:"build_parameters,omitempty"`
	// The number of containers to use to run the build. Default is null and the project default is used.
	Parallel *string `json:"parallel,omitempty"`
	// The specific revision to build. Default is null and the head of the branch is used. Cannot be used with tag parameter.
	Revision *string `json:"revision,omitempty"`
}

// NewInlineObject2 instantiates a new InlineObject2 object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewInlineObject2() *InlineObject2 {
	this := InlineObject2{}
	return &this
}

// NewInlineObject2WithDefaults instantiates a new InlineObject2 object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewInlineObject2WithDefaults() *InlineObject2 {
	this := InlineObject2{}
	return &this
}

// GetBuildParameters returns the BuildParameters field value if set, zero value otherwise.
func (o *InlineObject2) GetBuildParameters() map[string]interface{} {
	if o == nil || o.BuildParameters == nil {
		var ret map[string]interface{}
		return ret
	}
	return *o.BuildParameters
}

// GetBuildParametersOk returns a tuple with the BuildParameters field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject2) GetBuildParametersOk() (*map[string]interface{}, bool) {
	if o == nil || o.BuildParameters == nil {
		return nil, false
	}
	return o.BuildParameters, true
}

// HasBuildParameters returns a boolean if a field has been set.
func (o *InlineObject2) HasBuildParameters() bool {
	if o != nil && o.BuildParameters != nil {
		return true
	}

	return false
}

// SetBuildParameters gets a reference to the given map[string]interface{} and assigns it to the BuildParameters field.
func (o *InlineObject2) SetBuildParameters(v map[string]interface{}) {
	o.BuildParameters = &v
}

// GetParallel returns the Parallel field value if set, zero value otherwise.
func (o *InlineObject2) GetParallel() string {
	if o == nil || o.Parallel == nil {
		var ret string
		return ret
	}
	return *o.Parallel
}

// GetParallelOk returns a tuple with the Parallel field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject2) GetParallelOk() (*string, bool) {
	if o == nil || o.Parallel == nil {
		return nil, false
	}
	return o.Parallel, true
}

// HasParallel returns a boolean if a field has been set.
func (o *InlineObject2) HasParallel() bool {
	if o != nil && o.Parallel != nil {
		return true
	}

	return false
}

// SetParallel gets a reference to the given string and assigns it to the Parallel field.
func (o *InlineObject2) SetParallel(v string) {
	o.Parallel = &v
}

// GetRevision returns the Revision field value if set, zero value otherwise.
func (o *InlineObject2) GetRevision() string {
	if o == nil || o.Revision == nil {
		var ret string
		return ret
	}
	return *o.Revision
}

// GetRevisionOk returns a tuple with the Revision field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *InlineObject2) GetRevisionOk() (*string, bool) {
	if o == nil || o.Revision == nil {
		return nil, false
	}
	return o.Revision, true
}

// HasRevision returns a boolean if a field has been set.
func (o *InlineObject2) HasRevision() bool {
	if o != nil && o.Revision != nil {
		return true
	}

	return false
}

// SetRevision gets a reference to the given string and assigns it to the Revision field.
func (o *InlineObject2) SetRevision(v string) {
	o.Revision = &v
}

func (o InlineObject2) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.BuildParameters != nil {
		toSerialize["build_parameters"] = o.BuildParameters
	}
	if o.Parallel != nil {
		toSerialize["parallel"] = o.Parallel
	}
	if o.Revision != nil {
		toSerialize["revision"] = o.Revision
	}
	return json.Marshal(toSerialize)
}

type NullableInlineObject2 struct {
	value *InlineObject2
	isSet bool
}

func (v NullableInlineObject2) Get() *InlineObject2 {
	return v.value
}

func (v *NullableInlineObject2) Set(val *InlineObject2) {
	v.value = val
	v.isSet = true
}

func (v NullableInlineObject2) IsSet() bool {
	return v.isSet
}

func (v *NullableInlineObject2) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableInlineObject2(val *InlineObject2) *NullableInlineObject2 {
	return &NullableInlineObject2{value: val, isSet: true}
}

func (v NullableInlineObject2) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableInlineObject2) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
