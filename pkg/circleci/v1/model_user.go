// Copyright 2020 The ccictl Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package v1

import (
	"time"

	json "github.com/goccy/go-json"
)

// User struct for User
type User struct {
	Admin               *bool                   `json:"admin,omitempty"`
	AllEmails           *[]string               `json:"all_emails,omitempty"`
	AnalyticsId         *string                 `json:"analytics_id,omitempty"`
	AvatarUrl           *string                 `json:"avatar_url,omitempty"`
	BasicEmailPrefs     *string                 `json:"basic_email_prefs,omitempty"`
	BitbucketAuthorized *bool                   `json:"bitbucket_authorized,omitempty"`
	Containers          *int32                  `json:"containers,omitempty"`
	CreatedAt           *time.Time              `json:"created_at,omitempty"`
	DaysLeftInTrial     *int32                  `json:"days_left_in_trial,omitempty"`
	DevAdmin            *bool                   `json:"dev_admin,omitempty"`
	EnrolledBetas       *[]string               `json:"enrolled_betas,omitempty"`
	GithubOauthScopes   *[]string               `json:"github_oauth_scopes,omitempty"`
	InBetaProgram       *bool                   `json:"in_beta_program,omitempty"`
	Login               *string                 `json:"login,omitempty"`
	Name                *string                 `json:"name,omitempty"`
	OrganizationPrefs   *map[string]interface{} `json:"organization_prefs,omitempty"`
	Parallelism         *int32                  `json:"parallelism,omitempty"`
	Projects            *map[string]interface{} `json:"projects,omitempty"`
	PusherId            *string                 `json:"pusher_id,omitempty"`
	SelectedEmail       *string                 `json:"selected_email,omitempty"`
	SignInCount         *int32                  `json:"sign_in_count,omitempty"`
	TrialEnd            *time.Time              `json:"trial_end,omitempty"`
}

// NewUser instantiates a new User object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewUser() *User {
	this := User{}
	return &this
}

// NewUserWithDefaults instantiates a new User object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewUserWithDefaults() *User {
	this := User{}
	return &this
}

// GetAdmin returns the Admin field value if set, zero value otherwise.
func (o *User) GetAdmin() bool {
	if o == nil || o.Admin == nil {
		var ret bool
		return ret
	}
	return *o.Admin
}

// GetAdminOk returns a tuple with the Admin field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *User) GetAdminOk() (*bool, bool) {
	if o == nil || o.Admin == nil {
		return nil, false
	}
	return o.Admin, true
}

// HasAdmin returns a boolean if a field has been set.
func (o *User) HasAdmin() bool {
	if o != nil && o.Admin != nil {
		return true
	}

	return false
}

// SetAdmin gets a reference to the given bool and assigns it to the Admin field.
func (o *User) SetAdmin(v bool) {
	o.Admin = &v
}

// GetAllEmails returns the AllEmails field value if set, zero value otherwise.
func (o *User) GetAllEmails() []string {
	if o == nil || o.AllEmails == nil {
		var ret []string
		return ret
	}
	return *o.AllEmails
}

// GetAllEmailsOk returns a tuple with the AllEmails field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *User) GetAllEmailsOk() (*[]string, bool) {
	if o == nil || o.AllEmails == nil {
		return nil, false
	}
	return o.AllEmails, true
}

// HasAllEmails returns a boolean if a field has been set.
func (o *User) HasAllEmails() bool {
	if o != nil && o.AllEmails != nil {
		return true
	}

	return false
}

// SetAllEmails gets a reference to the given []string and assigns it to the AllEmails field.
func (o *User) SetAllEmails(v []string) {
	o.AllEmails = &v
}

// GetAnalyticsId returns the AnalyticsId field value if set, zero value otherwise.
func (o *User) GetAnalyticsId() string {
	if o == nil || o.AnalyticsId == nil {
		var ret string
		return ret
	}
	return *o.AnalyticsId
}

// GetAnalyticsIdOk returns a tuple with the AnalyticsId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *User) GetAnalyticsIdOk() (*string, bool) {
	if o == nil || o.AnalyticsId == nil {
		return nil, false
	}
	return o.AnalyticsId, true
}

// HasAnalyticsId returns a boolean if a field has been set.
func (o *User) HasAnalyticsId() bool {
	if o != nil && o.AnalyticsId != nil {
		return true
	}

	return false
}

// SetAnalyticsId gets a reference to the given string and assigns it to the AnalyticsId field.
func (o *User) SetAnalyticsId(v string) {
	o.AnalyticsId = &v
}

// GetAvatarUrl returns the AvatarUrl field value if set, zero value otherwise.
func (o *User) GetAvatarUrl() string {
	if o == nil || o.AvatarUrl == nil {
		var ret string
		return ret
	}
	return *o.AvatarUrl
}

// GetAvatarUrlOk returns a tuple with the AvatarUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *User) GetAvatarUrlOk() (*string, bool) {
	if o == nil || o.AvatarUrl == nil {
		return nil, false
	}
	return o.AvatarUrl, true
}

// HasAvatarUrl returns a boolean if a field has been set.
func (o *User) HasAvatarUrl() bool {
	if o != nil && o.AvatarUrl != nil {
		return true
	}

	return false
}

// SetAvatarUrl gets a reference to the given string and assigns it to the AvatarUrl field.
func (o *User) SetAvatarUrl(v string) {
	o.AvatarUrl = &v
}

// GetBasicEmailPrefs returns the BasicEmailPrefs field value if set, zero value otherwise.
func (o *User) GetBasicEmailPrefs() string {
	if o == nil || o.BasicEmailPrefs == nil {
		var ret string
		return ret
	}
	return *o.BasicEmailPrefs
}

// GetBasicEmailPrefsOk returns a tuple with the BasicEmailPrefs field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *User) GetBasicEmailPrefsOk() (*string, bool) {
	if o == nil || o.BasicEmailPrefs == nil {
		return nil, false
	}
	return o.BasicEmailPrefs, true
}

// HasBasicEmailPrefs returns a boolean if a field has been set.
func (o *User) HasBasicEmailPrefs() bool {
	if o != nil && o.BasicEmailPrefs != nil {
		return true
	}

	return false
}

// SetBasicEmailPrefs gets a reference to the given string and assigns it to the BasicEmailPrefs field.
func (o *User) SetBasicEmailPrefs(v string) {
	o.BasicEmailPrefs = &v
}

// GetBitbucketAuthorized returns the BitbucketAuthorized field value if set, zero value otherwise.
func (o *User) GetBitbucketAuthorized() bool {
	if o == nil || o.BitbucketAuthorized == nil {
		var ret bool
		return ret
	}
	return *o.BitbucketAuthorized
}

// GetBitbucketAuthorizedOk returns a tuple with the BitbucketAuthorized field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *User) GetBitbucketAuthorizedOk() (*bool, bool) {
	if o == nil || o.BitbucketAuthorized == nil {
		return nil, false
	}
	return o.BitbucketAuthorized, true
}

// HasBitbucketAuthorized returns a boolean if a field has been set.
func (o *User) HasBitbucketAuthorized() bool {
	if o != nil && o.BitbucketAuthorized != nil {
		return true
	}

	return false
}

// SetBitbucketAuthorized gets a reference to the given bool and assigns it to the BitbucketAuthorized field.
func (o *User) SetBitbucketAuthorized(v bool) {
	o.BitbucketAuthorized = &v
}

// GetContainers returns the Containers field value if set, zero value otherwise.
func (o *User) GetContainers() int32 {
	if o == nil || o.Containers == nil {
		var ret int32
		return ret
	}
	return *o.Containers
}

// GetContainersOk returns a tuple with the Containers field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *User) GetContainersOk() (*int32, bool) {
	if o == nil || o.Containers == nil {
		return nil, false
	}
	return o.Containers, true
}

// HasContainers returns a boolean if a field has been set.
func (o *User) HasContainers() bool {
	if o != nil && o.Containers != nil {
		return true
	}

	return false
}

// SetContainers gets a reference to the given int32 and assigns it to the Containers field.
func (o *User) SetContainers(v int32) {
	o.Containers = &v
}

// GetCreatedAt returns the CreatedAt field value if set, zero value otherwise.
func (o *User) GetCreatedAt() time.Time {
	if o == nil || o.CreatedAt == nil {
		var ret time.Time
		return ret
	}
	return *o.CreatedAt
}

// GetCreatedAtOk returns a tuple with the CreatedAt field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *User) GetCreatedAtOk() (*time.Time, bool) {
	if o == nil || o.CreatedAt == nil {
		return nil, false
	}
	return o.CreatedAt, true
}

// HasCreatedAt returns a boolean if a field has been set.
func (o *User) HasCreatedAt() bool {
	if o != nil && o.CreatedAt != nil {
		return true
	}

	return false
}

// SetCreatedAt gets a reference to the given time.Time and assigns it to the CreatedAt field.
func (o *User) SetCreatedAt(v time.Time) {
	o.CreatedAt = &v
}

// GetDaysLeftInTrial returns the DaysLeftInTrial field value if set, zero value otherwise.
func (o *User) GetDaysLeftInTrial() int32 {
	if o == nil || o.DaysLeftInTrial == nil {
		var ret int32
		return ret
	}
	return *o.DaysLeftInTrial
}

// GetDaysLeftInTrialOk returns a tuple with the DaysLeftInTrial field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *User) GetDaysLeftInTrialOk() (*int32, bool) {
	if o == nil || o.DaysLeftInTrial == nil {
		return nil, false
	}
	return o.DaysLeftInTrial, true
}

// HasDaysLeftInTrial returns a boolean if a field has been set.
func (o *User) HasDaysLeftInTrial() bool {
	if o != nil && o.DaysLeftInTrial != nil {
		return true
	}

	return false
}

// SetDaysLeftInTrial gets a reference to the given int32 and assigns it to the DaysLeftInTrial field.
func (o *User) SetDaysLeftInTrial(v int32) {
	o.DaysLeftInTrial = &v
}

// GetDevAdmin returns the DevAdmin field value if set, zero value otherwise.
func (o *User) GetDevAdmin() bool {
	if o == nil || o.DevAdmin == nil {
		var ret bool
		return ret
	}
	return *o.DevAdmin
}

// GetDevAdminOk returns a tuple with the DevAdmin field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *User) GetDevAdminOk() (*bool, bool) {
	if o == nil || o.DevAdmin == nil {
		return nil, false
	}
	return o.DevAdmin, true
}

// HasDevAdmin returns a boolean if a field has been set.
func (o *User) HasDevAdmin() bool {
	if o != nil && o.DevAdmin != nil {
		return true
	}

	return false
}

// SetDevAdmin gets a reference to the given bool and assigns it to the DevAdmin field.
func (o *User) SetDevAdmin(v bool) {
	o.DevAdmin = &v
}

// GetEnrolledBetas returns the EnrolledBetas field value if set, zero value otherwise.
func (o *User) GetEnrolledBetas() []string {
	if o == nil || o.EnrolledBetas == nil {
		var ret []string
		return ret
	}
	return *o.EnrolledBetas
}

// GetEnrolledBetasOk returns a tuple with the EnrolledBetas field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *User) GetEnrolledBetasOk() (*[]string, bool) {
	if o == nil || o.EnrolledBetas == nil {
		return nil, false
	}
	return o.EnrolledBetas, true
}

// HasEnrolledBetas returns a boolean if a field has been set.
func (o *User) HasEnrolledBetas() bool {
	if o != nil && o.EnrolledBetas != nil {
		return true
	}

	return false
}

// SetEnrolledBetas gets a reference to the given []string and assigns it to the EnrolledBetas field.
func (o *User) SetEnrolledBetas(v []string) {
	o.EnrolledBetas = &v
}

// GetGithubOauthScopes returns the GithubOauthScopes field value if set, zero value otherwise.
func (o *User) GetGithubOauthScopes() []string {
	if o == nil || o.GithubOauthScopes == nil {
		var ret []string
		return ret
	}
	return *o.GithubOauthScopes
}

// GetGithubOauthScopesOk returns a tuple with the GithubOauthScopes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *User) GetGithubOauthScopesOk() (*[]string, bool) {
	if o == nil || o.GithubOauthScopes == nil {
		return nil, false
	}
	return o.GithubOauthScopes, true
}

// HasGithubOauthScopes returns a boolean if a field has been set.
func (o *User) HasGithubOauthScopes() bool {
	if o != nil && o.GithubOauthScopes != nil {
		return true
	}

	return false
}

// SetGithubOauthScopes gets a reference to the given []string and assigns it to the GithubOauthScopes field.
func (o *User) SetGithubOauthScopes(v []string) {
	o.GithubOauthScopes = &v
}

// GetInBetaProgram returns the InBetaProgram field value if set, zero value otherwise.
func (o *User) GetInBetaProgram() bool {
	if o == nil || o.InBetaProgram == nil {
		var ret bool
		return ret
	}
	return *o.InBetaProgram
}

// GetInBetaProgramOk returns a tuple with the InBetaProgram field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *User) GetInBetaProgramOk() (*bool, bool) {
	if o == nil || o.InBetaProgram == nil {
		return nil, false
	}
	return o.InBetaProgram, true
}

// HasInBetaProgram returns a boolean if a field has been set.
func (o *User) HasInBetaProgram() bool {
	if o != nil && o.InBetaProgram != nil {
		return true
	}

	return false
}

// SetInBetaProgram gets a reference to the given bool and assigns it to the InBetaProgram field.
func (o *User) SetInBetaProgram(v bool) {
	o.InBetaProgram = &v
}

// GetLogin returns the Login field value if set, zero value otherwise.
func (o *User) GetLogin() string {
	if o == nil || o.Login == nil {
		var ret string
		return ret
	}
	return *o.Login
}

// GetLoginOk returns a tuple with the Login field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *User) GetLoginOk() (*string, bool) {
	if o == nil || o.Login == nil {
		return nil, false
	}
	return o.Login, true
}

// HasLogin returns a boolean if a field has been set.
func (o *User) HasLogin() bool {
	if o != nil && o.Login != nil {
		return true
	}

	return false
}

// SetLogin gets a reference to the given string and assigns it to the Login field.
func (o *User) SetLogin(v string) {
	o.Login = &v
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *User) GetName() string {
	if o == nil || o.Name == nil {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *User) GetNameOk() (*string, bool) {
	if o == nil || o.Name == nil {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *User) HasName() bool {
	if o != nil && o.Name != nil {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *User) SetName(v string) {
	o.Name = &v
}

// GetOrganizationPrefs returns the OrganizationPrefs field value if set, zero value otherwise.
func (o *User) GetOrganizationPrefs() map[string]interface{} {
	if o == nil || o.OrganizationPrefs == nil {
		var ret map[string]interface{}
		return ret
	}
	return *o.OrganizationPrefs
}

// GetOrganizationPrefsOk returns a tuple with the OrganizationPrefs field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *User) GetOrganizationPrefsOk() (*map[string]interface{}, bool) {
	if o == nil || o.OrganizationPrefs == nil {
		return nil, false
	}
	return o.OrganizationPrefs, true
}

// HasOrganizationPrefs returns a boolean if a field has been set.
func (o *User) HasOrganizationPrefs() bool {
	if o != nil && o.OrganizationPrefs != nil {
		return true
	}

	return false
}

// SetOrganizationPrefs gets a reference to the given map[string]interface{} and assigns it to the OrganizationPrefs field.
func (o *User) SetOrganizationPrefs(v map[string]interface{}) {
	o.OrganizationPrefs = &v
}

// GetParallelism returns the Parallelism field value if set, zero value otherwise.
func (o *User) GetParallelism() int32 {
	if o == nil || o.Parallelism == nil {
		var ret int32
		return ret
	}
	return *o.Parallelism
}

// GetParallelismOk returns a tuple with the Parallelism field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *User) GetParallelismOk() (*int32, bool) {
	if o == nil || o.Parallelism == nil {
		return nil, false
	}
	return o.Parallelism, true
}

// HasParallelism returns a boolean if a field has been set.
func (o *User) HasParallelism() bool {
	if o != nil && o.Parallelism != nil {
		return true
	}

	return false
}

// SetParallelism gets a reference to the given int32 and assigns it to the Parallelism field.
func (o *User) SetParallelism(v int32) {
	o.Parallelism = &v
}

// GetProjects returns the Projects field value if set, zero value otherwise.
func (o *User) GetProjects() map[string]interface{} {
	if o == nil || o.Projects == nil {
		var ret map[string]interface{}
		return ret
	}
	return *o.Projects
}

// GetProjectsOk returns a tuple with the Projects field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *User) GetProjectsOk() (*map[string]interface{}, bool) {
	if o == nil || o.Projects == nil {
		return nil, false
	}
	return o.Projects, true
}

// HasProjects returns a boolean if a field has been set.
func (o *User) HasProjects() bool {
	if o != nil && o.Projects != nil {
		return true
	}

	return false
}

// SetProjects gets a reference to the given map[string]interface{} and assigns it to the Projects field.
func (o *User) SetProjects(v map[string]interface{}) {
	o.Projects = &v
}

// GetPusherId returns the PusherId field value if set, zero value otherwise.
func (o *User) GetPusherId() string {
	if o == nil || o.PusherId == nil {
		var ret string
		return ret
	}
	return *o.PusherId
}

// GetPusherIdOk returns a tuple with the PusherId field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *User) GetPusherIdOk() (*string, bool) {
	if o == nil || o.PusherId == nil {
		return nil, false
	}
	return o.PusherId, true
}

// HasPusherId returns a boolean if a field has been set.
func (o *User) HasPusherId() bool {
	if o != nil && o.PusherId != nil {
		return true
	}

	return false
}

// SetPusherId gets a reference to the given string and assigns it to the PusherId field.
func (o *User) SetPusherId(v string) {
	o.PusherId = &v
}

// GetSelectedEmail returns the SelectedEmail field value if set, zero value otherwise.
func (o *User) GetSelectedEmail() string {
	if o == nil || o.SelectedEmail == nil {
		var ret string
		return ret
	}
	return *o.SelectedEmail
}

// GetSelectedEmailOk returns a tuple with the SelectedEmail field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *User) GetSelectedEmailOk() (*string, bool) {
	if o == nil || o.SelectedEmail == nil {
		return nil, false
	}
	return o.SelectedEmail, true
}

// HasSelectedEmail returns a boolean if a field has been set.
func (o *User) HasSelectedEmail() bool {
	if o != nil && o.SelectedEmail != nil {
		return true
	}

	return false
}

// SetSelectedEmail gets a reference to the given string and assigns it to the SelectedEmail field.
func (o *User) SetSelectedEmail(v string) {
	o.SelectedEmail = &v
}

// GetSignInCount returns the SignInCount field value if set, zero value otherwise.
func (o *User) GetSignInCount() int32 {
	if o == nil || o.SignInCount == nil {
		var ret int32
		return ret
	}
	return *o.SignInCount
}

// GetSignInCountOk returns a tuple with the SignInCount field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *User) GetSignInCountOk() (*int32, bool) {
	if o == nil || o.SignInCount == nil {
		return nil, false
	}
	return o.SignInCount, true
}

// HasSignInCount returns a boolean if a field has been set.
func (o *User) HasSignInCount() bool {
	if o != nil && o.SignInCount != nil {
		return true
	}

	return false
}

// SetSignInCount gets a reference to the given int32 and assigns it to the SignInCount field.
func (o *User) SetSignInCount(v int32) {
	o.SignInCount = &v
}

// GetTrialEnd returns the TrialEnd field value if set, zero value otherwise.
func (o *User) GetTrialEnd() time.Time {
	if o == nil || o.TrialEnd == nil {
		var ret time.Time
		return ret
	}
	return *o.TrialEnd
}

// GetTrialEndOk returns a tuple with the TrialEnd field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *User) GetTrialEndOk() (*time.Time, bool) {
	if o == nil || o.TrialEnd == nil {
		return nil, false
	}
	return o.TrialEnd, true
}

// HasTrialEnd returns a boolean if a field has been set.
func (o *User) HasTrialEnd() bool {
	if o != nil && o.TrialEnd != nil {
		return true
	}

	return false
}

// SetTrialEnd gets a reference to the given time.Time and assigns it to the TrialEnd field.
func (o *User) SetTrialEnd(v time.Time) {
	o.TrialEnd = &v
}

func (o User) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Admin != nil {
		toSerialize["admin"] = o.Admin
	}
	if o.AllEmails != nil {
		toSerialize["all_emails"] = o.AllEmails
	}
	if o.AnalyticsId != nil {
		toSerialize["analytics_id"] = o.AnalyticsId
	}
	if o.AvatarUrl != nil {
		toSerialize["avatar_url"] = o.AvatarUrl
	}
	if o.BasicEmailPrefs != nil {
		toSerialize["basic_email_prefs"] = o.BasicEmailPrefs
	}
	if o.BitbucketAuthorized != nil {
		toSerialize["bitbucket_authorized"] = o.BitbucketAuthorized
	}
	if o.Containers != nil {
		toSerialize["containers"] = o.Containers
	}
	if o.CreatedAt != nil {
		toSerialize["created_at"] = o.CreatedAt
	}
	if o.DaysLeftInTrial != nil {
		toSerialize["days_left_in_trial"] = o.DaysLeftInTrial
	}
	if o.DevAdmin != nil {
		toSerialize["dev_admin"] = o.DevAdmin
	}
	if o.EnrolledBetas != nil {
		toSerialize["enrolled_betas"] = o.EnrolledBetas
	}
	if o.GithubOauthScopes != nil {
		toSerialize["github_oauth_scopes"] = o.GithubOauthScopes
	}
	if o.InBetaProgram != nil {
		toSerialize["in_beta_program"] = o.InBetaProgram
	}
	if o.Login != nil {
		toSerialize["login"] = o.Login
	}
	if o.Name != nil {
		toSerialize["name"] = o.Name
	}
	if o.OrganizationPrefs != nil {
		toSerialize["organization_prefs"] = o.OrganizationPrefs
	}
	if o.Parallelism != nil {
		toSerialize["parallelism"] = o.Parallelism
	}
	if o.Projects != nil {
		toSerialize["projects"] = o.Projects
	}
	if o.PusherId != nil {
		toSerialize["pusher_id"] = o.PusherId
	}
	if o.SelectedEmail != nil {
		toSerialize["selected_email"] = o.SelectedEmail
	}
	if o.SignInCount != nil {
		toSerialize["sign_in_count"] = o.SignInCount
	}
	if o.TrialEnd != nil {
		toSerialize["trial_end"] = o.TrialEnd
	}
	return json.Marshal(toSerialize)
}

type NullableUser struct {
	value *User
	isSet bool
}

func (v NullableUser) Get() *User {
	return v.value
}

func (v *NullableUser) Set(val *User) {
	v.value = val
	v.isSet = true
}

func (v NullableUser) IsSet() bool {
	return v.isSet
}

func (v *NullableUser) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableUser(val *User) *NullableUser {
	return &NullableUser{value: val, isSet: true}
}

func (v NullableUser) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableUser) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
