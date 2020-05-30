// Copyright 2020 The ccictl Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package v1

import (
	"fmt"

	json "github.com/goccy/go-json"
)

// Scope the model 'Scope'
type Scope string

// List of Scope
const (
	SCOPE_WRITE_SETTINGS Scope = "write-settings"
	SCOPE_VIEW_BUILDS    Scope = "view-builds"
	SCOPE_READ_SETTINGS  Scope = "read-settings"
	SCOPE_TRIGGER_BUILDS Scope = "trigger-builds"
	SCOPE_ALL            Scope = "all"
	SCOPE_STATUS         Scope = "status"
	SCOPE_NONE           Scope = "none"
)

func (v *Scope) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := Scope(value)
	for _, existing := range []Scope{"write-settings", "view-builds", "read-settings", "trigger-builds", "all", "status", "none"} {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid Scope", *v)
}

// Ptr returns reference to Scope value
func (v Scope) Ptr() *Scope {
	return &v
}

type NullableScope struct {
	value *Scope
	isSet bool
}

func (v NullableScope) Get() *Scope {
	return v.value
}

func (v *NullableScope) Set(val *Scope) {
	v.value = val
	v.isSet = true
}

func (v NullableScope) IsSet() bool {
	return v.isSet
}

func (v *NullableScope) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableScope(val *Scope) *NullableScope {
	return &NullableScope{value: val, isSet: true}
}

func (v NullableScope) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableScope) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
