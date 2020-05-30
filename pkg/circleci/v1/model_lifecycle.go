// Copyright 2020 The ccictl Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package v1

import (
	"fmt"

	json "github.com/goccy/go-json"
)

// Lifecycle the model 'Lifecycle'
type Lifecycle string

// List of Lifecycle
const (
	LIFECYCLE_QUEUED      Lifecycle = "queued"
	LIFECYCLE_SCHEDULED   Lifecycle = "scheduled"
	LIFECYCLE_NOT_RUN     Lifecycle = "not_run"
	LIFECYCLE_NOT_RUNNING Lifecycle = "not_running"
	LIFECYCLE_RUNNING     Lifecycle = "running"
	LIFECYCLE_FINISHED    Lifecycle = "finished"
)

func (v *Lifecycle) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := Lifecycle(value)
	for _, existing := range []Lifecycle{"queued", "scheduled", "not_run", "not_running", "running", "finished"} {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid Lifecycle", *v)
}

// Ptr returns reference to Lifecycle value
func (v Lifecycle) Ptr() *Lifecycle {
	return &v
}

type NullableLifecycle struct {
	value *Lifecycle
	isSet bool
}

func (v NullableLifecycle) Get() *Lifecycle {
	return v.value
}

func (v *NullableLifecycle) Set(val *Lifecycle) {
	v.value = val
	v.isSet = true
}

func (v NullableLifecycle) IsSet() bool {
	return v.isSet
}

func (v *NullableLifecycle) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableLifecycle(val *Lifecycle) *NullableLifecycle {
	return &NullableLifecycle{value: val, isSet: true}
}

func (v NullableLifecycle) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableLifecycle) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
