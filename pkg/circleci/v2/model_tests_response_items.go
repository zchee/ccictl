// Copyright 2020 The ccictl Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package v2

import (
	json "github.com/goccy/go-json"
)

// TestsResponseItems struct for TestsResponseItems
type TestsResponseItems struct {
	// The failure message associated with the test.
	Message string `json:"message"`
	// The file in which the test is defined.
	File string `json:"file"`
	// Indication of whether the test succeeded.
	Result string `json:"result"`
	// The name of the test.
	Name string `json:"name"`
	// The programmatic location of the test.
	Classname string `json:"classname"`
}

// NewTestsResponseItems instantiates a new TestsResponseItems object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTestsResponseItems(message string, file string, result string, name string, classname string) *TestsResponseItems {
	this := TestsResponseItems{}
	this.Message = message
	this.File = file
	this.Result = result
	this.Name = name
	this.Classname = classname
	return &this
}

// NewTestsResponseItemsWithDefaults instantiates a new TestsResponseItems object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTestsResponseItemsWithDefaults() *TestsResponseItems {
	this := TestsResponseItems{}
	return &this
}

// GetMessage returns the Message field value
func (o *TestsResponseItems) GetMessage() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Message
}

// GetMessageOk returns a tuple with the Message field value
// and a boolean to check if the value has been set.
func (o *TestsResponseItems) GetMessageOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Message, true
}

// SetMessage sets field value
func (o *TestsResponseItems) SetMessage(v string) {
	o.Message = v
}

// GetFile returns the File field value
func (o *TestsResponseItems) GetFile() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.File
}

// GetFileOk returns a tuple with the File field value
// and a boolean to check if the value has been set.
func (o *TestsResponseItems) GetFileOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.File, true
}

// SetFile sets field value
func (o *TestsResponseItems) SetFile(v string) {
	o.File = v
}

// GetResult returns the Result field value
func (o *TestsResponseItems) GetResult() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Result
}

// GetResultOk returns a tuple with the Result field value
// and a boolean to check if the value has been set.
func (o *TestsResponseItems) GetResultOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Result, true
}

// SetResult sets field value
func (o *TestsResponseItems) SetResult(v string) {
	o.Result = v
}

// GetName returns the Name field value
func (o *TestsResponseItems) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *TestsResponseItems) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *TestsResponseItems) SetName(v string) {
	o.Name = v
}

// GetClassname returns the Classname field value
func (o *TestsResponseItems) GetClassname() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Classname
}

// GetClassnameOk returns a tuple with the Classname field value
// and a boolean to check if the value has been set.
func (o *TestsResponseItems) GetClassnameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Classname, true
}

// SetClassname sets field value
func (o *TestsResponseItems) SetClassname(v string) {
	o.Classname = v
}

func (o TestsResponseItems) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["message"] = o.Message
	}
	if true {
		toSerialize["file"] = o.File
	}
	if true {
		toSerialize["result"] = o.Result
	}
	if true {
		toSerialize["name"] = o.Name
	}
	if true {
		toSerialize["classname"] = o.Classname
	}
	return json.Marshal(toSerialize)
}

type NullableTestsResponseItems struct {
	value *TestsResponseItems
	isSet bool
}

func (v NullableTestsResponseItems) Get() *TestsResponseItems {
	return v.value
}

func (v *NullableTestsResponseItems) Set(val *TestsResponseItems) {
	v.value = val
	v.isSet = true
}

func (v NullableTestsResponseItems) IsSet() bool {
	return v.isSet
}

func (v *NullableTestsResponseItems) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTestsResponseItems(val *TestsResponseItems) *NullableTestsResponseItems {
	return &NullableTestsResponseItems{value: val, isSet: true}
}

func (v NullableTestsResponseItems) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTestsResponseItems) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
