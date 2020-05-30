// Copyright 2020 The ccictl Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package v2

import (
	json "github.com/goccy/go-json"
)

// JobDetailsOrganization Information about an organization.
type JobDetailsOrganization struct {
	// The name of the organization.
	Name string `json:"name"`
}

// NewJobDetailsOrganization instantiates a new JobDetailsOrganization object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewJobDetailsOrganization(name string) *JobDetailsOrganization {
	this := JobDetailsOrganization{}
	this.Name = name
	return &this
}

// NewJobDetailsOrganizationWithDefaults instantiates a new JobDetailsOrganization object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewJobDetailsOrganizationWithDefaults() *JobDetailsOrganization {
	this := JobDetailsOrganization{}
	return &this
}

// GetName returns the Name field value
func (o *JobDetailsOrganization) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *JobDetailsOrganization) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *JobDetailsOrganization) SetName(v string) {
	o.Name = v
}

func (o JobDetailsOrganization) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["name"] = o.Name
	}
	return json.Marshal(toSerialize)
}

type NullableJobDetailsOrganization struct {
	value *JobDetailsOrganization
	isSet bool
}

func (v NullableJobDetailsOrganization) Get() *JobDetailsOrganization {
	return v.value
}

func (v *NullableJobDetailsOrganization) Set(val *JobDetailsOrganization) {
	v.value = val
	v.isSet = true
}

func (v NullableJobDetailsOrganization) IsSet() bool {
	return v.isSet
}

func (v *NullableJobDetailsOrganization) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableJobDetailsOrganization(val *JobDetailsOrganization) *NullableJobDetailsOrganization {
	return &NullableJobDetailsOrganization{value: val, isSet: true}
}

func (v NullableJobDetailsOrganization) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableJobDetailsOrganization) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
