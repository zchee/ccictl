// Copyright 2020 The ccictl Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package v1

import (
	"time"

	json "github.com/goccy/go-json"
)

// CommitDetail struct for CommitDetail
type CommitDetail struct {
	AuthorDate     *time.Time `json:"author_date,omitempty"`
	AuthorEmail    *string    `json:"author_email,omitempty"`
	AuthorLogin    *string    `json:"author_login,omitempty"`
	AuthorName     *string    `json:"author_name,omitempty"`
	Body           *string    `json:"body,omitempty"`
	Commit         *string    `json:"commit,omitempty"`
	CommitUrl      *string    `json:"commit_url,omitempty"`
	CommitterDate  *time.Time `json:"committer_date,omitempty"`
	CommitterEmail *string    `json:"committer_email,omitempty"`
	CommitterLogin *string    `json:"committer_login,omitempty"`
	CommitterName  *string    `json:"committer_name,omitempty"`
	Subject        *string    `json:"subject,omitempty"`
}

// NewCommitDetail instantiates a new CommitDetail object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewCommitDetail() *CommitDetail {
	this := CommitDetail{}
	return &this
}

// NewCommitDetailWithDefaults instantiates a new CommitDetail object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewCommitDetailWithDefaults() *CommitDetail {
	this := CommitDetail{}
	return &this
}

// GetAuthorDate returns the AuthorDate field value if set, zero value otherwise.
func (o *CommitDetail) GetAuthorDate() time.Time {
	if o == nil || o.AuthorDate == nil {
		var ret time.Time
		return ret
	}
	return *o.AuthorDate
}

// GetAuthorDateOk returns a tuple with the AuthorDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CommitDetail) GetAuthorDateOk() (*time.Time, bool) {
	if o == nil || o.AuthorDate == nil {
		return nil, false
	}
	return o.AuthorDate, true
}

// HasAuthorDate returns a boolean if a field has been set.
func (o *CommitDetail) HasAuthorDate() bool {
	if o != nil && o.AuthorDate != nil {
		return true
	}

	return false
}

// SetAuthorDate gets a reference to the given time.Time and assigns it to the AuthorDate field.
func (o *CommitDetail) SetAuthorDate(v time.Time) {
	o.AuthorDate = &v
}

// GetAuthorEmail returns the AuthorEmail field value if set, zero value otherwise.
func (o *CommitDetail) GetAuthorEmail() string {
	if o == nil || o.AuthorEmail == nil {
		var ret string
		return ret
	}
	return *o.AuthorEmail
}

// GetAuthorEmailOk returns a tuple with the AuthorEmail field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CommitDetail) GetAuthorEmailOk() (*string, bool) {
	if o == nil || o.AuthorEmail == nil {
		return nil, false
	}
	return o.AuthorEmail, true
}

// HasAuthorEmail returns a boolean if a field has been set.
func (o *CommitDetail) HasAuthorEmail() bool {
	if o != nil && o.AuthorEmail != nil {
		return true
	}

	return false
}

// SetAuthorEmail gets a reference to the given string and assigns it to the AuthorEmail field.
func (o *CommitDetail) SetAuthorEmail(v string) {
	o.AuthorEmail = &v
}

// GetAuthorLogin returns the AuthorLogin field value if set, zero value otherwise.
func (o *CommitDetail) GetAuthorLogin() string {
	if o == nil || o.AuthorLogin == nil {
		var ret string
		return ret
	}
	return *o.AuthorLogin
}

// GetAuthorLoginOk returns a tuple with the AuthorLogin field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CommitDetail) GetAuthorLoginOk() (*string, bool) {
	if o == nil || o.AuthorLogin == nil {
		return nil, false
	}
	return o.AuthorLogin, true
}

// HasAuthorLogin returns a boolean if a field has been set.
func (o *CommitDetail) HasAuthorLogin() bool {
	if o != nil && o.AuthorLogin != nil {
		return true
	}

	return false
}

// SetAuthorLogin gets a reference to the given string and assigns it to the AuthorLogin field.
func (o *CommitDetail) SetAuthorLogin(v string) {
	o.AuthorLogin = &v
}

// GetAuthorName returns the AuthorName field value if set, zero value otherwise.
func (o *CommitDetail) GetAuthorName() string {
	if o == nil || o.AuthorName == nil {
		var ret string
		return ret
	}
	return *o.AuthorName
}

// GetAuthorNameOk returns a tuple with the AuthorName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CommitDetail) GetAuthorNameOk() (*string, bool) {
	if o == nil || o.AuthorName == nil {
		return nil, false
	}
	return o.AuthorName, true
}

// HasAuthorName returns a boolean if a field has been set.
func (o *CommitDetail) HasAuthorName() bool {
	if o != nil && o.AuthorName != nil {
		return true
	}

	return false
}

// SetAuthorName gets a reference to the given string and assigns it to the AuthorName field.
func (o *CommitDetail) SetAuthorName(v string) {
	o.AuthorName = &v
}

// GetBody returns the Body field value if set, zero value otherwise.
func (o *CommitDetail) GetBody() string {
	if o == nil || o.Body == nil {
		var ret string
		return ret
	}
	return *o.Body
}

// GetBodyOk returns a tuple with the Body field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CommitDetail) GetBodyOk() (*string, bool) {
	if o == nil || o.Body == nil {
		return nil, false
	}
	return o.Body, true
}

// HasBody returns a boolean if a field has been set.
func (o *CommitDetail) HasBody() bool {
	if o != nil && o.Body != nil {
		return true
	}

	return false
}

// SetBody gets a reference to the given string and assigns it to the Body field.
func (o *CommitDetail) SetBody(v string) {
	o.Body = &v
}

// GetCommit returns the Commit field value if set, zero value otherwise.
func (o *CommitDetail) GetCommit() string {
	if o == nil || o.Commit == nil {
		var ret string
		return ret
	}
	return *o.Commit
}

// GetCommitOk returns a tuple with the Commit field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CommitDetail) GetCommitOk() (*string, bool) {
	if o == nil || o.Commit == nil {
		return nil, false
	}
	return o.Commit, true
}

// HasCommit returns a boolean if a field has been set.
func (o *CommitDetail) HasCommit() bool {
	if o != nil && o.Commit != nil {
		return true
	}

	return false
}

// SetCommit gets a reference to the given string and assigns it to the Commit field.
func (o *CommitDetail) SetCommit(v string) {
	o.Commit = &v
}

// GetCommitUrl returns the CommitUrl field value if set, zero value otherwise.
func (o *CommitDetail) GetCommitUrl() string {
	if o == nil || o.CommitUrl == nil {
		var ret string
		return ret
	}
	return *o.CommitUrl
}

// GetCommitUrlOk returns a tuple with the CommitUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CommitDetail) GetCommitUrlOk() (*string, bool) {
	if o == nil || o.CommitUrl == nil {
		return nil, false
	}
	return o.CommitUrl, true
}

// HasCommitUrl returns a boolean if a field has been set.
func (o *CommitDetail) HasCommitUrl() bool {
	if o != nil && o.CommitUrl != nil {
		return true
	}

	return false
}

// SetCommitUrl gets a reference to the given string and assigns it to the CommitUrl field.
func (o *CommitDetail) SetCommitUrl(v string) {
	o.CommitUrl = &v
}

// GetCommitterDate returns the CommitterDate field value if set, zero value otherwise.
func (o *CommitDetail) GetCommitterDate() time.Time {
	if o == nil || o.CommitterDate == nil {
		var ret time.Time
		return ret
	}
	return *o.CommitterDate
}

// GetCommitterDateOk returns a tuple with the CommitterDate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CommitDetail) GetCommitterDateOk() (*time.Time, bool) {
	if o == nil || o.CommitterDate == nil {
		return nil, false
	}
	return o.CommitterDate, true
}

// HasCommitterDate returns a boolean if a field has been set.
func (o *CommitDetail) HasCommitterDate() bool {
	if o != nil && o.CommitterDate != nil {
		return true
	}

	return false
}

// SetCommitterDate gets a reference to the given time.Time and assigns it to the CommitterDate field.
func (o *CommitDetail) SetCommitterDate(v time.Time) {
	o.CommitterDate = &v
}

// GetCommitterEmail returns the CommitterEmail field value if set, zero value otherwise.
func (o *CommitDetail) GetCommitterEmail() string {
	if o == nil || o.CommitterEmail == nil {
		var ret string
		return ret
	}
	return *o.CommitterEmail
}

// GetCommitterEmailOk returns a tuple with the CommitterEmail field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CommitDetail) GetCommitterEmailOk() (*string, bool) {
	if o == nil || o.CommitterEmail == nil {
		return nil, false
	}
	return o.CommitterEmail, true
}

// HasCommitterEmail returns a boolean if a field has been set.
func (o *CommitDetail) HasCommitterEmail() bool {
	if o != nil && o.CommitterEmail != nil {
		return true
	}

	return false
}

// SetCommitterEmail gets a reference to the given string and assigns it to the CommitterEmail field.
func (o *CommitDetail) SetCommitterEmail(v string) {
	o.CommitterEmail = &v
}

// GetCommitterLogin returns the CommitterLogin field value if set, zero value otherwise.
func (o *CommitDetail) GetCommitterLogin() string {
	if o == nil || o.CommitterLogin == nil {
		var ret string
		return ret
	}
	return *o.CommitterLogin
}

// GetCommitterLoginOk returns a tuple with the CommitterLogin field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CommitDetail) GetCommitterLoginOk() (*string, bool) {
	if o == nil || o.CommitterLogin == nil {
		return nil, false
	}
	return o.CommitterLogin, true
}

// HasCommitterLogin returns a boolean if a field has been set.
func (o *CommitDetail) HasCommitterLogin() bool {
	if o != nil && o.CommitterLogin != nil {
		return true
	}

	return false
}

// SetCommitterLogin gets a reference to the given string and assigns it to the CommitterLogin field.
func (o *CommitDetail) SetCommitterLogin(v string) {
	o.CommitterLogin = &v
}

// GetCommitterName returns the CommitterName field value if set, zero value otherwise.
func (o *CommitDetail) GetCommitterName() string {
	if o == nil || o.CommitterName == nil {
		var ret string
		return ret
	}
	return *o.CommitterName
}

// GetCommitterNameOk returns a tuple with the CommitterName field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CommitDetail) GetCommitterNameOk() (*string, bool) {
	if o == nil || o.CommitterName == nil {
		return nil, false
	}
	return o.CommitterName, true
}

// HasCommitterName returns a boolean if a field has been set.
func (o *CommitDetail) HasCommitterName() bool {
	if o != nil && o.CommitterName != nil {
		return true
	}

	return false
}

// SetCommitterName gets a reference to the given string and assigns it to the CommitterName field.
func (o *CommitDetail) SetCommitterName(v string) {
	o.CommitterName = &v
}

// GetSubject returns the Subject field value if set, zero value otherwise.
func (o *CommitDetail) GetSubject() string {
	if o == nil || o.Subject == nil {
		var ret string
		return ret
	}
	return *o.Subject
}

// GetSubjectOk returns a tuple with the Subject field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *CommitDetail) GetSubjectOk() (*string, bool) {
	if o == nil || o.Subject == nil {
		return nil, false
	}
	return o.Subject, true
}

// HasSubject returns a boolean if a field has been set.
func (o *CommitDetail) HasSubject() bool {
	if o != nil && o.Subject != nil {
		return true
	}

	return false
}

// SetSubject gets a reference to the given string and assigns it to the Subject field.
func (o *CommitDetail) SetSubject(v string) {
	o.Subject = &v
}

func (o CommitDetail) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.AuthorDate != nil {
		toSerialize["author_date"] = o.AuthorDate
	}
	if o.AuthorEmail != nil {
		toSerialize["author_email"] = o.AuthorEmail
	}
	if o.AuthorLogin != nil {
		toSerialize["author_login"] = o.AuthorLogin
	}
	if o.AuthorName != nil {
		toSerialize["author_name"] = o.AuthorName
	}
	if o.Body != nil {
		toSerialize["body"] = o.Body
	}
	if o.Commit != nil {
		toSerialize["commit"] = o.Commit
	}
	if o.CommitUrl != nil {
		toSerialize["commit_url"] = o.CommitUrl
	}
	if o.CommitterDate != nil {
		toSerialize["committer_date"] = o.CommitterDate
	}
	if o.CommitterEmail != nil {
		toSerialize["committer_email"] = o.CommitterEmail
	}
	if o.CommitterLogin != nil {
		toSerialize["committer_login"] = o.CommitterLogin
	}
	if o.CommitterName != nil {
		toSerialize["committer_name"] = o.CommitterName
	}
	if o.Subject != nil {
		toSerialize["subject"] = o.Subject
	}
	return json.Marshal(toSerialize)
}

type NullableCommitDetail struct {
	value *CommitDetail
	isSet bool
}

func (v NullableCommitDetail) Get() *CommitDetail {
	return v.value
}

func (v *NullableCommitDetail) Set(val *CommitDetail) {
	v.value = val
	v.isSet = true
}

func (v NullableCommitDetail) IsSet() bool {
	return v.isSet
}

func (v *NullableCommitDetail) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCommitDetail(val *CommitDetail) *NullableCommitDetail {
	return &NullableCommitDetail{value: val, isSet: true}
}

func (v NullableCommitDetail) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCommitDetail) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
