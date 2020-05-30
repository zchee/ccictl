// Copyright 2020 The ccictl Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package v1

import (
	"fmt"

	json "github.com/goccy/go-json"
)

// Status the model 'Status'
type Status string

// List of Status
const (
	STATUS_RETRIED             Status = "retried"
	STATUS_CANCELED            Status = "canceled"
	STATUS_INFRASTRUCTURE_FAIL Status = "infrastructure_fail"
	STATUS_TIMEDOUT            Status = "timedout"
	STATUS_NOT_RUN             Status = "not_run"
	STATUS_RUNNING             Status = "running"
	STATUS_FAILED              Status = "failed"
	STATUS_QUEUED              Status = "queued"
	STATUS_SCHEDULED           Status = "scheduled"
	STATUS_NOT_RUNNING         Status = "not_running"
	STATUS_NO_TESTS            Status = "no_tests"
	STATUS_FIXED               Status = "fixed"
	STATUS_SUCCESS             Status = "success"
)

func (v *Status) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := Status(value)
	for _, existing := range []Status{"retried", "canceled", "infrastructure_fail", "timedout", "not_run", "running", "failed", "queued", "scheduled", "not_running", "no_tests", "fixed", "success"} {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid Status", *v)
}

// Ptr returns reference to Status value
func (v Status) Ptr() *Status {
	return &v
}

type NullableStatus struct {
	value *Status
	isSet bool
}

func (v NullableStatus) Get() *Status {
	return v.value
}

func (v *NullableStatus) Set(val *Status) {
	v.value = val
	v.isSet = true
}

func (v NullableStatus) IsSet() bool {
	return v.isSet
}

func (v *NullableStatus) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableStatus(val *Status) *NullableStatus {
	return &NullableStatus{value: val, isSet: true}
}

func (v NullableStatus) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableStatus) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
