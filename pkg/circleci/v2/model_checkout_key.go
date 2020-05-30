// Copyright 2020 The ccictl Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package v2

import (
	"time"

	json "github.com/goccy/go-json"
)

// CheckoutKey struct for CheckoutKey
type CheckoutKey struct {
	// A public SSH key.
	PublicKey string `json:"public-key"`
	// The type of checkout key. This may be either `deploy-key` or `github-user-key`.
	Type string `json:"type"`
	// An SSH key fingerprint.
	Fingerprint string `json:"fingerprint"`
	// A boolean value that indicates if this key is preferred.
	Preferred bool `json:"preferred"`
	// The date and time the checkout key was created.
	CreatedAt time.Time `json:"created-at"`
}

// NewCheckoutKey instantiates a new CheckoutKey object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCheckoutKey(publicKey string, type_ string, fingerprint string, preferred bool, createdAt time.Time) *CheckoutKey {
	this := CheckoutKey{}
	this.PublicKey = publicKey
	this.Type = type_
	this.Fingerprint = fingerprint
	this.Preferred = preferred
	this.CreatedAt = createdAt
	return &this
}

// NewCheckoutKeyWithDefaults instantiates a new CheckoutKey object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCheckoutKeyWithDefaults() *CheckoutKey {
	this := CheckoutKey{}
	return &this
}

// GetPublicKey returns the PublicKey field value
func (o *CheckoutKey) GetPublicKey() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.PublicKey
}

// GetPublicKeyOk returns a tuple with the PublicKey field value
// and a boolean to check if the value has been set.
func (o *CheckoutKey) GetPublicKeyOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PublicKey, true
}

// SetPublicKey sets field value
func (o *CheckoutKey) SetPublicKey(v string) {
	o.PublicKey = v
}

// GetType returns the Type field value
func (o *CheckoutKey) GetType() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Type
}

// GetTypeOk returns a tuple with the Type field value
// and a boolean to check if the value has been set.
func (o *CheckoutKey) GetTypeOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Type, true
}

// SetType sets field value
func (o *CheckoutKey) SetType(v string) {
	o.Type = v
}

// GetFingerprint returns the Fingerprint field value
func (o *CheckoutKey) GetFingerprint() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Fingerprint
}

// GetFingerprintOk returns a tuple with the Fingerprint field value
// and a boolean to check if the value has been set.
func (o *CheckoutKey) GetFingerprintOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Fingerprint, true
}

// SetFingerprint sets field value
func (o *CheckoutKey) SetFingerprint(v string) {
	o.Fingerprint = v
}

// GetPreferred returns the Preferred field value
func (o *CheckoutKey) GetPreferred() bool {
	if o == nil {
		var ret bool
		return ret
	}

	return o.Preferred
}

// GetPreferredOk returns a tuple with the Preferred field value
// and a boolean to check if the value has been set.
func (o *CheckoutKey) GetPreferredOk() (*bool, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Preferred, true
}

// SetPreferred sets field value
func (o *CheckoutKey) SetPreferred(v bool) {
	o.Preferred = v
}

// GetCreatedAt returns the CreatedAt field value
func (o *CheckoutKey) GetCreatedAt() time.Time {
	if o == nil {
		var ret time.Time
		return ret
	}

	return o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value
// and a boolean to check if the value has been set.
func (o *CheckoutKey) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return &o.CreatedAt, true
}

// SetCreatedAt sets field value
func (o *CheckoutKey) SetCreatedAt(v time.Time) {
	o.CreatedAt = v
}

func (o CheckoutKey) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if true {
		toSerialize["public-key"] = o.PublicKey
	}
	if true {
		toSerialize["type"] = o.Type
	}
	if true {
		toSerialize["fingerprint"] = o.Fingerprint
	}
	if true {
		toSerialize["preferred"] = o.Preferred
	}
	if true {
		toSerialize["created-at"] = o.CreatedAt
	}
	return json.Marshal(toSerialize)
}

type NullableCheckoutKey struct {
	value *CheckoutKey
	isSet bool
}

func (v NullableCheckoutKey) Get() *CheckoutKey {
	return v.value
}

func (v *NullableCheckoutKey) Set(val *CheckoutKey) {
	v.value = val
	v.isSet = true
}

func (v NullableCheckoutKey) IsSet() bool {
	return v.isSet
}

func (v *NullableCheckoutKey) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCheckoutKey(val *CheckoutKey) *NullableCheckoutKey {
	return &NullableCheckoutKey{value: val, isSet: true}
}

func (v NullableCheckoutKey) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCheckoutKey) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
