// Copyright 2020 The ccictl Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package v1

import (
	json "github.com/goccy/go-json"
)

// TestsTests struct for TestsTests
type TestsTests struct {
	Classname *string  `json:"classname,omitempty"`
	File      *string  `json:"file,omitempty"`
	Message   *string  `json:"message,omitempty"`
	Name      *string  `json:"name,omitempty"`
	Result    *Status  `json:"result,omitempty"`
	RunTime   *float32 `json:"run_time,omitempty"`
	Source    *string  `json:"source,omitempty"`
}

// NewTestsTests instantiates a new TestsTests object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewTestsTests() *TestsTests {
	this := TestsTests{}
	return &this
}

// NewTestsTestsWithDefaults instantiates a new TestsTests object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewTestsTestsWithDefaults() *TestsTests {
	this := TestsTests{}
	return &this
}

// GetClassname returns the Classname field value if set, zero value otherwise.
func (o *TestsTests) GetClassname() string {
	if o == nil || o.Classname == nil {
		var ret string
		return ret
	}
	return *o.Classname
}

// GetClassnameOk returns a tuple with the Classname field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TestsTests) GetClassnameOk() (*string, bool) {
	if o == nil || o.Classname == nil {
		return nil, false
	}
	return o.Classname, true
}

// HasClassname returns a boolean if a field has been set.
func (o *TestsTests) HasClassname() bool {
	if o != nil && o.Classname != nil {
		return true
	}

	return false
}

// SetClassname gets a reference to the given string and assigns it to the Classname field.
func (o *TestsTests) SetClassname(v string) {
	o.Classname = &v
}

// GetFile returns the File field value if set, zero value otherwise.
func (o *TestsTests) GetFile() string {
	if o == nil || o.File == nil {
		var ret string
		return ret
	}
	return *o.File
}

// GetFileOk returns a tuple with the File field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TestsTests) GetFileOk() (*string, bool) {
	if o == nil || o.File == nil {
		return nil, false
	}
	return o.File, true
}

// HasFile returns a boolean if a field has been set.
func (o *TestsTests) HasFile() bool {
	if o != nil && o.File != nil {
		return true
	}

	return false
}

// SetFile gets a reference to the given string and assigns it to the File field.
func (o *TestsTests) SetFile(v string) {
	o.File = &v
}

// GetMessage returns the Message field value if set, zero value otherwise.
func (o *TestsTests) GetMessage() string {
	if o == nil || o.Message == nil {
		var ret string
		return ret
	}
	return *o.Message
}

// GetMessageOk returns a tuple with the Message field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TestsTests) GetMessageOk() (*string, bool) {
	if o == nil || o.Message == nil {
		return nil, false
	}
	return o.Message, true
}

// HasMessage returns a boolean if a field has been set.
func (o *TestsTests) HasMessage() bool {
	if o != nil && o.Message != nil {
		return true
	}

	return false
}

// SetMessage gets a reference to the given string and assigns it to the Message field.
func (o *TestsTests) SetMessage(v string) {
	o.Message = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *TestsTests) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TestsTests) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *TestsTests) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *TestsTests) SetName(v string) {
	o.Name = &v
}

// GetResult returns the Result field value if set, zero value otherwise.
func (o *TestsTests) GetResult() Status {
	if o == nil || o.Result == nil {
		var ret Status
		return ret
	}
	return *o.Result
}

// GetResultOk returns a tuple with the Result field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TestsTests) GetResultOk() (*Status, bool) {
	if o == nil || o.Result == nil {
		return nil, false
	}
	return o.Result, true
}

// HasResult returns a boolean if a field has been set.
func (o *TestsTests) HasResult() bool {
	if o != nil && o.Result != nil {
		return true
	}

	return false
}

// SetResult gets a reference to the given Status and assigns it to the Result field.
func (o *TestsTests) SetResult(v Status) {
	o.Result = &v
}

// GetRunTime returns the RunTime field value if set, zero value otherwise.
func (o *TestsTests) GetRunTime() float32 {
	if o == nil || o.RunTime == nil {
		var ret float32
		return ret
	}
	return *o.RunTime
}

// GetRunTimeOk returns a tuple with the RunTime field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TestsTests) GetRunTimeOk() (*float32, bool) {
	if o == nil || o.RunTime == nil {
		return nil, false
	}
	return o.RunTime, true
}

// HasRunTime returns a boolean if a field has been set.
func (o *TestsTests) HasRunTime() bool {
	if o != nil && o.RunTime != nil {
		return true
	}

	return false
}

// SetRunTime gets a reference to the given float32 and assigns it to the RunTime field.
func (o *TestsTests) SetRunTime(v float32) {
	o.RunTime = &v
}

// GetSource returns the Source field value if set, zero value otherwise.
func (o *TestsTests) GetSource() string {
	if o == nil || o.Source == nil {
		var ret string
		return ret
	}
	return *o.Source
}

// GetSourceOk returns a tuple with the Source field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *TestsTests) GetSourceOk() (*string, bool) {
	if o == nil || o.Source == nil {
		return nil, false
	}
	return o.Source, true
}

// HasSource returns a boolean if a field has been set.
func (o *TestsTests) HasSource() bool {
	if o != nil && o.Source != nil {
		return true
	}

	return false
}

// SetSource gets a reference to the given string and assigns it to the Source field.
func (o *TestsTests) SetSource(v string) {
	o.Source = &v
}

func (o TestsTests) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Classname != nil {
		toSerialize["classname"] = o.Classname
	}
	if o.File != nil {
		toSerialize["file"] = o.File
	}
	if o.Message != nil {
		toSerialize["message"] = o.Message
	}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.Result != nil {
		toSerialize["result"] = o.Result
	}
	if o.RunTime != nil {
		toSerialize["run_time"] = o.RunTime
	}
	if o.Source != nil {
		toSerialize["source"] = o.Source
	}
	return json.Marshal(toSerialize)
}

type NullableTestsTests struct {
	value *TestsTests
	isSet bool
}

func (v NullableTestsTests) Get() *TestsTests {
	return v.value
}

func (v *NullableTestsTests) Set(val *TestsTests) {
	v.value = val
	v.isSet = true
}

func (v NullableTestsTests) IsSet() bool {
	return v.isSet
}

func (v *NullableTestsTests) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableTestsTests(val *TestsTests) *NullableTestsTests {
	return &NullableTestsTests{value: val, isSet: true}
}

func (v NullableTestsTests) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableTestsTests) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
