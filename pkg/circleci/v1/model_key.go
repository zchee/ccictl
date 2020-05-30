// Copyright 2020 The ccictl Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package v1

import (
	"time"

	json "github.com/goccy/go-json"
)

// Key struct for Key
type Key struct {
	Fingerprint *string `json:"fingerprint,omitempty"`
	Preferred   *bool   `json:"preferred,omitempty"`
	PublicKey   *string `json:"public_key,omitempty"`
	// when the key was issued
	Time *time.Time `json:"time,omitempty"`
	// can be \"deploy-key\" or \"github-user-key\"
	Type *string `json:"type,omitempty"`
}

// NewKey instantiates a new Key object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewKey() *Key {
	this := Key{}
	return &this
}

// NewKeyWithDefaults instantiates a new Key object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewKeyWithDefaults() *Key {
	this := Key{}
	return &this
}

// GetFingerprint returns the Fingerprint field value if set, zero value otherwise.
func (o *Key) GetFingerprint() string {
	if o == nil || o.Fingerprint == nil {
		var ret string
		return ret
	}
	return *o.Fingerprint
}

// GetFingerprintOk returns a tuple with the Fingerprint field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Key) GetFingerprintOk() (*string, bool) {
	if o == nil || o.Fingerprint == nil {
		return nil, false
	}
	return o.Fingerprint, true
}

// HasFingerprint returns a boolean if a field has been set.
func (o *Key) HasFingerprint() bool {
	if o != nil && o.Fingerprint != nil {
		return true
	}

	return false
}

// SetFingerprint gets a reference to the given string and assigns it to the Fingerprint field.
func (o *Key) SetFingerprint(v string) {
	o.Fingerprint = &v
}

// GetPreferred returns the Preferred field value if set, zero value otherwise.
func (o *Key) GetPreferred() bool {
	if o == nil || o.Preferred == nil {
		var ret bool
		return ret
	}
	return *o.Preferred
}

// GetPreferredOk returns a tuple with the Preferred field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Key) GetPreferredOk() (*bool, bool) {
	if o == nil || o.Preferred == nil {
		return nil, false
	}
	return o.Preferred, true
}

// HasPreferred returns a boolean if a field has been set.
func (o *Key) HasPreferred() bool {
	if o != nil && o.Preferred != nil {
		return true
	}

	return false
}

// SetPreferred gets a reference to the given bool and assigns it to the Preferred field.
func (o *Key) SetPreferred(v bool) {
	o.Preferred = &v
}

// GetPublicKey returns the PublicKey field value if set, zero value otherwise.
func (o *Key) GetPublicKey() string {
	if o == nil || o.PublicKey == nil {
		var ret string
		return ret
	}
	return *o.PublicKey
}

// GetPublicKeyOk returns a tuple with the PublicKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Key) GetPublicKeyOk() (*string, bool) {
	if o == nil || o.PublicKey == nil {
		return nil, false
	}
	return o.PublicKey, true
}

// HasPublicKey returns a boolean if a field has been set.
func (o *Key) HasPublicKey() bool {
	if o != nil && o.PublicKey != nil {
		return true
	}

	return false
}

// SetPublicKey gets a reference to the given string and assigns it to the PublicKey field.
func (o *Key) SetPublicKey(v string) {
	o.PublicKey = &v
}

// GetTime returns the Time field value if set, zero value otherwise.
func (o *Key) GetTime() time.Time {
	if o == nil || o.Time == nil {
		var ret time.Time
		return ret
	}
	return *o.Time
}

// GetTimeOk returns a tuple with the Time field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Key) GetTimeOk() (*time.Time, bool) {
	if o == nil || o.Time == nil {
		return nil, false
	}
	return o.Time, true
}

// HasTime returns a boolean if a field has been set.
func (o *Key) HasTime() bool {
	if o != nil && o.Time != nil {
		return true
	}

	return false
}

// SetTime gets a reference to the given time.Time and assigns it to the Time field.
func (o *Key) SetTime(v time.Time) {
	o.Time = &v
}

// GetType returns the Type field value if set, zero value otherwise.
func (o *Key) GetType() string {
	if o == nil || o.Type == nil {
		var ret string
		return ret
	}
	return *o.Type
}

// GetTypeOk returns a tuple with the Type field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Key) GetTypeOk() (*string, bool) {
	if o == nil || o.Type == nil {
		return nil, false
	}
	return o.Type, true
}

// HasType returns a boolean if a field has been set.
func (o *Key) HasType() bool {
	if o != nil && o.Type != nil {
		return true
	}

	return false
}

// SetType gets a reference to the given string and assigns it to the Type field.
func (o *Key) SetType(v string) {
	o.Type = &v
}

func (o Key) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Fingerprint != nil {
		toSerialize["fingerprint"] = o.Fingerprint
	}
	if o.Preferred != nil {
		toSerialize["preferred"] = o.Preferred
	}
	if o.PublicKey != nil {
		toSerialize["public_key"] = o.PublicKey
	}
	if o.Time != nil {
		toSerialize["time"] = o.Time
	}
	if o.Type != nil {
		toSerialize["type"] = o.Type
	}
	return json.Marshal(toSerialize)
}

type NullableKey struct {
	value *Key
	isSet bool
}

func (v NullableKey) Get() *Key {
	return v.value
}

func (v *NullableKey) Set(val *Key) {
	v.value = val
	v.isSet = true
}

func (v NullableKey) IsSet() bool {
	return v.isSet
}

func (v *NullableKey) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableKey(val *Key) *NullableKey {
	return &NullableKey{value: val, isSet: true}
}

func (v NullableKey) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableKey) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
