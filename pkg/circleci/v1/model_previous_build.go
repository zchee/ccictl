// Copyright 2020 The ccictl Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package v1

import (
	json "github.com/goccy/go-json"
)

// PreviousBuild previous build
type PreviousBuild struct {
	BuildNum        *int32  `json:"build_num,omitempty"`
	BuildTimeMillis *int32  `json:"build_time_millis,omitempty"`
	Status          *Status `json:"status,omitempty"`
}

// NewPreviousBuild instantiates a new PreviousBuild object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPreviousBuild() *PreviousBuild {
	this := PreviousBuild{}
	return &this
}

// NewPreviousBuildWithDefaults instantiates a new PreviousBuild object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPreviousBuildWithDefaults() *PreviousBuild {
	this := PreviousBuild{}
	return &this
}

// GetBuildNum returns the BuildNum field value if set, zero value otherwise.
func (o *PreviousBuild) GetBuildNum() int32 {
	if o == nil || o.BuildNum == nil {
		var ret int32
		return ret
	}
	return *o.BuildNum
}

// GetBuildNumOk returns a tuple with the BuildNum field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PreviousBuild) GetBuildNumOk() (*int32, bool) {
	if o == nil || o.BuildNum == nil {
		return nil, false
	}
	return o.BuildNum, true
}

// HasBuildNum returns a boolean if a field has been set.
func (o *PreviousBuild) HasBuildNum() bool {
	if o != nil && o.BuildNum != nil {
		return true
	}

	return false
}

// SetBuildNum gets a reference to the given int32 and assigns it to the BuildNum field.
func (o *PreviousBuild) SetBuildNum(v int32) {
	o.BuildNum = &v
}

// GetBuildTimeMillis returns the BuildTimeMillis field value if set, zero value otherwise.
func (o *PreviousBuild) GetBuildTimeMillis() int32 {
	if o == nil || o.BuildTimeMillis == nil {
		var ret int32
		return ret
	}
	return *o.BuildTimeMillis
}

// GetBuildTimeMillisOk returns a tuple with the BuildTimeMillis field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PreviousBuild) GetBuildTimeMillisOk() (*int32, bool) {
	if o == nil || o.BuildTimeMillis == nil {
		return nil, false
	}
	return o.BuildTimeMillis, true
}

// HasBuildTimeMillis returns a boolean if a field has been set.
func (o *PreviousBuild) HasBuildTimeMillis() bool {
	if o != nil && o.BuildTimeMillis != nil {
		return true
	}

	return false
}

// SetBuildTimeMillis gets a reference to the given int32 and assigns it to the BuildTimeMillis field.
func (o *PreviousBuild) SetBuildTimeMillis(v int32) {
	o.BuildTimeMillis = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *PreviousBuild) GetStatus() Status {
	if o == nil || o.Status == nil {
		var ret Status
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PreviousBuild) GetStatusOk() (*Status, bool) {
	if o == nil || o.Status == nil {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *PreviousBuild) HasStatus() bool {
	if o != nil && o.Status != nil {
		return true
	}

	return false
}

// SetStatus gets a reference to the given Status and assigns it to the Status field.
func (o *PreviousBuild) SetStatus(v Status) {
	o.Status = &v
}

func (o PreviousBuild) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.BuildNum != nil {
		toSerialize["build_num"] = o.BuildNum
	}
	if o.BuildTimeMillis != nil {
		toSerialize["build_time_millis"] = o.BuildTimeMillis
	}
	if o.Status != nil {
		toSerialize["status"] = o.Status
	}
	return json.Marshal(toSerialize)
}

type NullablePreviousBuild struct {
	value *PreviousBuild
	isSet bool
}

func (v NullablePreviousBuild) Get() *PreviousBuild {
	return v.value
}

func (v *NullablePreviousBuild) Set(val *PreviousBuild) {
	v.value = val
	v.isSet = true
}

func (v NullablePreviousBuild) IsSet() bool {
	return v.isSet
}

func (v *NullablePreviousBuild) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePreviousBuild(val *PreviousBuild) *NullablePreviousBuild {
	return &NullablePreviousBuild{value: val, isSet: true}
}

func (v NullablePreviousBuild) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePreviousBuild) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
