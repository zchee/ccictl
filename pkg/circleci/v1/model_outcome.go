// Copyright 2020 The ccictl Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package v1

import (
	"fmt"

	json "github.com/goccy/go-json"
)

// Outcome the model 'Outcome'
type Outcome string

// List of Outcome
const (
	OUTCOME_CANCELED            Outcome = "canceled"
	OUTCOME_INFRASTRUCTURE_FAIL Outcome = "infrastructure_fail"
	OUTCOME_TIMEDOUT            Outcome = "timedout"
	OUTCOME_FAILED              Outcome = "failed"
	OUTCOME_NO_TESTS            Outcome = "no_tests"
	OUTCOME_SUCCESS             Outcome = "success"
)

func (v *Outcome) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := Outcome(value)
	for _, existing := range []Outcome{"canceled", "infrastructure_fail", "timedout", "failed", "no_tests", "success"} {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid Outcome", *v)
}

// Ptr returns reference to Outcome value
func (v Outcome) Ptr() *Outcome {
	return &v
}

type NullableOutcome struct {
	value *Outcome
	isSet bool
}

func (v NullableOutcome) Get() *Outcome {
	return v.value
}

func (v *NullableOutcome) Set(val *Outcome) {
	v.value = val
	v.isSet = true
}

func (v NullableOutcome) IsSet() bool {
	return v.isSet
}

func (v *NullableOutcome) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableOutcome(val *Outcome) *NullableOutcome {
	return &NullableOutcome{value: val, isSet: true}
}

func (v NullableOutcome) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableOutcome) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
