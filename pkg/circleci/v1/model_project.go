// Copyright 2020 The ccictl Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package v1

import (
	json "github.com/goccy/go-json"
)

// Project struct for Project
type Project struct {
	Aws                  *map[string]interface{} `json:"aws,omitempty"`
	Branches             *map[string]interface{} `json:"branches,omitempty"`
	CampfireNotifyPrefs  *map[string]interface{} `json:"campfire_notify_prefs,omitempty"`
	CampfireRoom         *map[string]interface{} `json:"campfire_room,omitempty"`
	CampfireSubdomain    *map[string]interface{} `json:"campfire_subdomain,omitempty"`
	CampfireToken        *map[string]interface{} `json:"campfire_token,omitempty"`
	Compile              *string                 `json:"compile,omitempty"`
	DefaultBranch        *string                 `json:"default_branch,omitempty"`
	Dependencies         *string                 `json:"dependencies,omitempty"`
	Extra                *string                 `json:"extra,omitempty"`
	FeatureFlags         *ProjectFeatureFlags    `json:"feature_flags,omitempty"`
	FlowdockApiToken     *map[string]interface{} `json:"flowdock_api_token,omitempty"`
	Followed             *bool                   `json:"followed,omitempty"`
	HasUsableKey         *bool                   `json:"has_usable_key,omitempty"`
	HerokuDeployUser     *map[string]interface{} `json:"heroku_deploy_user,omitempty"`
	HipchatApiToken      *map[string]interface{} `json:"hipchat_api_token,omitempty"`
	HipchatNotify        *map[string]interface{} `json:"hipchat_notify,omitempty"`
	HipchatRoom          *map[string]interface{} `json:"hipchat_room,omitempty"`
	IrcChannel           *map[string]interface{} `json:"irc_channel,omitempty"`
	IrcKeyword           *map[string]interface{} `json:"irc_keyword,omitempty"`
	IrcNotifyPrefs       *map[string]interface{} `json:"irc_notify_prefs,omitempty"`
	IrcPassword          *map[string]interface{} `json:"irc_password,omitempty"`
	IrcServer            *map[string]interface{} `json:"irc_server,omitempty"`
	IrcUsername          *map[string]interface{} `json:"irc_username,omitempty"`
	Language             *string                 `json:"language,omitempty"`
	Oss                  *bool                   `json:"oss,omitempty"`
	Parallel             *int32                  `json:"parallel,omitempty"`
	Reponame             *string                 `json:"reponame,omitempty"`
	Scopes               *[]Scope                `json:"scopes,omitempty"`
	Setup                *string                 `json:"setup,omitempty"`
	SlackApiToken        *map[string]interface{} `json:"slack_api_token,omitempty"`
	SlackChannel         *map[string]interface{} `json:"slack_channel,omitempty"`
	SlackChannelOverride *map[string]interface{} `json:"slack_channel_override,omitempty"`
	SlackNotifyPrefs     *map[string]interface{} `json:"slack_notify_prefs,omitempty"`
	SlackSubdomain       *map[string]interface{} `json:"slack_subdomain,omitempty"`
	SlackWebhookUrl      *string                 `json:"slack_webhook_url,omitempty"`
	SshKeys              *[]string               `json:"ssh_keys,omitempty"`
	Test                 *string                 `json:"test,omitempty"`
	Username             *string                 `json:"username,omitempty"`
	VcsType              *string                 `json:"vcs_type,omitempty"`
	VcsUrl               *string                 `json:"vcs_url,omitempty"`
}

// NewProject instantiates a new Project object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewProject() *Project {
	this := Project{}
	return &this
}

// NewProjectWithDefaults instantiates a new Project object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewProjectWithDefaults() *Project {
	this := Project{}
	return &this
}

// GetAws returns the Aws field value if set, zero value otherwise.
func (o *Project) GetAws() map[string]interface{} {
	if o == nil || o.Aws == nil {
		var ret map[string]interface{}
		return ret
	}
	return *o.Aws
}

// GetAwsOk returns a tuple with the Aws field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Project) GetAwsOk() (*map[string]interface{}, bool) {
	if o == nil || o.Aws == nil {
		return nil, false
	}
	return o.Aws, true
}

// HasAws returns a boolean if a field has been set.
func (o *Project) HasAws() bool {
	if o != nil && o.Aws != nil {
		return true
	}

	return false
}

// SetAws gets a reference to the given map[string]interface{} and assigns it to the Aws field.
func (o *Project) SetAws(v map[string]interface{}) {
	o.Aws = &v
}

// GetBranches returns the Branches field value if set, zero value otherwise.
func (o *Project) GetBranches() map[string]interface{} {
	if o == nil || o.Branches == nil {
		var ret map[string]interface{}
		return ret
	}
	return *o.Branches
}

// GetBranchesOk returns a tuple with the Branches field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Project) GetBranchesOk() (*map[string]interface{}, bool) {
	if o == nil || o.Branches == nil {
		return nil, false
	}
	return o.Branches, true
}

// HasBranches returns a boolean if a field has been set.
func (o *Project) HasBranches() bool {
	if o != nil && o.Branches != nil {
		return true
	}

	return false
}

// SetBranches gets a reference to the given map[string]interface{} and assigns it to the Branches field.
func (o *Project) SetBranches(v map[string]interface{}) {
	o.Branches = &v
}

// GetCampfireNotifyPrefs returns the CampfireNotifyPrefs field value if set, zero value otherwise.
func (o *Project) GetCampfireNotifyPrefs() map[string]interface{} {
	if o == nil || o.CampfireNotifyPrefs == nil {
		var ret map[string]interface{}
		return ret
	}
	return *o.CampfireNotifyPrefs
}

// GetCampfireNotifyPrefsOk returns a tuple with the CampfireNotifyPrefs field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Project) GetCampfireNotifyPrefsOk() (*map[string]interface{}, bool) {
	if o == nil || o.CampfireNotifyPrefs == nil {
		return nil, false
	}
	return o.CampfireNotifyPrefs, true
}

// HasCampfireNotifyPrefs returns a boolean if a field has been set.
func (o *Project) HasCampfireNotifyPrefs() bool {
	if o != nil && o.CampfireNotifyPrefs != nil {
		return true
	}

	return false
}

// SetCampfireNotifyPrefs gets a reference to the given map[string]interface{} and assigns it to the CampfireNotifyPrefs field.
func (o *Project) SetCampfireNotifyPrefs(v map[string]interface{}) {
	o.CampfireNotifyPrefs = &v
}

// GetCampfireRoom returns the CampfireRoom field value if set, zero value otherwise.
func (o *Project) GetCampfireRoom() map[string]interface{} {
	if o == nil || o.CampfireRoom == nil {
		var ret map[string]interface{}
		return ret
	}
	return *o.CampfireRoom
}

// GetCampfireRoomOk returns a tuple with the CampfireRoom field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Project) GetCampfireRoomOk() (*map[string]interface{}, bool) {
	if o == nil || o.CampfireRoom == nil {
		return nil, false
	}
	return o.CampfireRoom, true
}

// HasCampfireRoom returns a boolean if a field has been set.
func (o *Project) HasCampfireRoom() bool {
	if o != nil && o.CampfireRoom != nil {
		return true
	}

	return false
}

// SetCampfireRoom gets a reference to the given map[string]interface{} and assigns it to the CampfireRoom field.
func (o *Project) SetCampfireRoom(v map[string]interface{}) {
	o.CampfireRoom = &v
}

// GetCampfireSubdomain returns the CampfireSubdomain field value if set, zero value otherwise.
func (o *Project) GetCampfireSubdomain() map[string]interface{} {
	if o == nil || o.CampfireSubdomain == nil {
		var ret map[string]interface{}
		return ret
	}
	return *o.CampfireSubdomain
}

// GetCampfireSubdomainOk returns a tuple with the CampfireSubdomain field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Project) GetCampfireSubdomainOk() (*map[string]interface{}, bool) {
	if o == nil || o.CampfireSubdomain == nil {
		return nil, false
	}
	return o.CampfireSubdomain, true
}

// HasCampfireSubdomain returns a boolean if a field has been set.
func (o *Project) HasCampfireSubdomain() bool {
	if o != nil && o.CampfireSubdomain != nil {
		return true
	}

	return false
}

// SetCampfireSubdomain gets a reference to the given map[string]interface{} and assigns it to the CampfireSubdomain field.
func (o *Project) SetCampfireSubdomain(v map[string]interface{}) {
	o.CampfireSubdomain = &v
}

// GetCampfireToken returns the CampfireToken field value if set, zero value otherwise.
func (o *Project) GetCampfireToken() map[string]interface{} {
	if o == nil || o.CampfireToken == nil {
		var ret map[string]interface{}
		return ret
	}
	return *o.CampfireToken
}

// GetCampfireTokenOk returns a tuple with the CampfireToken field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Project) GetCampfireTokenOk() (*map[string]interface{}, bool) {
	if o == nil || o.CampfireToken == nil {
		return nil, false
	}
	return o.CampfireToken, true
}

// HasCampfireToken returns a boolean if a field has been set.
func (o *Project) HasCampfireToken() bool {
	if o != nil && o.CampfireToken != nil {
		return true
	}

	return false
}

// SetCampfireToken gets a reference to the given map[string]interface{} and assigns it to the CampfireToken field.
func (o *Project) SetCampfireToken(v map[string]interface{}) {
	o.CampfireToken = &v
}

// GetCompile returns the Compile field value if set, zero value otherwise.
func (o *Project) GetCompile() string {
	if o == nil || o.Compile == nil {
		var ret string
		return ret
	}
	return *o.Compile
}

// GetCompileOk returns a tuple with the Compile field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Project) GetCompileOk() (*string, bool) {
	if o == nil || o.Compile == nil {
		return nil, false
	}
	return o.Compile, true
}

// HasCompile returns a boolean if a field has been set.
func (o *Project) HasCompile() bool {
	if o != nil && o.Compile != nil {
		return true
	}

	return false
}

// SetCompile gets a reference to the given string and assigns it to the Compile field.
func (o *Project) SetCompile(v string) {
	o.Compile = &v
}

// GetDefaultBranch returns the DefaultBranch field value if set, zero value otherwise.
func (o *Project) GetDefaultBranch() string {
	if o == nil || o.DefaultBranch == nil {
		var ret string
		return ret
	}
	return *o.DefaultBranch
}

// GetDefaultBranchOk returns a tuple with the DefaultBranch field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Project) GetDefaultBranchOk() (*string, bool) {
	if o == nil || o.DefaultBranch == nil {
		return nil, false
	}
	return o.DefaultBranch, true
}

// HasDefaultBranch returns a boolean if a field has been set.
func (o *Project) HasDefaultBranch() bool {
	if o != nil && o.DefaultBranch != nil {
		return true
	}

	return false
}

// SetDefaultBranch gets a reference to the given string and assigns it to the DefaultBranch field.
func (o *Project) SetDefaultBranch(v string) {
	o.DefaultBranch = &v
}

// GetDependencies returns the Dependencies field value if set, zero value otherwise.
func (o *Project) GetDependencies() string {
	if o == nil || o.Dependencies == nil {
		var ret string
		return ret
	}
	return *o.Dependencies
}

// GetDependenciesOk returns a tuple with the Dependencies field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Project) GetDependenciesOk() (*string, bool) {
	if o == nil || o.Dependencies == nil {
		return nil, false
	}
	return o.Dependencies, true
}

// HasDependencies returns a boolean if a field has been set.
func (o *Project) HasDependencies() bool {
	if o != nil && o.Dependencies != nil {
		return true
	}

	return false
}

// SetDependencies gets a reference to the given string and assigns it to the Dependencies field.
func (o *Project) SetDependencies(v string) {
	o.Dependencies = &v
}

// GetExtra returns the Extra field value if set, zero value otherwise.
func (o *Project) GetExtra() string {
	if o == nil || o.Extra == nil {
		var ret string
		return ret
	}
	return *o.Extra
}

// GetExtraOk returns a tuple with the Extra field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Project) GetExtraOk() (*string, bool) {
	if o == nil || o.Extra == nil {
		return nil, false
	}
	return o.Extra, true
}

// HasExtra returns a boolean if a field has been set.
func (o *Project) HasExtra() bool {
	if o != nil && o.Extra != nil {
		return true
	}

	return false
}

// SetExtra gets a reference to the given string and assigns it to the Extra field.
func (o *Project) SetExtra(v string) {
	o.Extra = &v
}

// GetFeatureFlags returns the FeatureFlags field value if set, zero value otherwise.
func (o *Project) GetFeatureFlags() ProjectFeatureFlags {
	if o == nil || o.FeatureFlags == nil {
		var ret ProjectFeatureFlags
		return ret
	}
	return *o.FeatureFlags
}

// GetFeatureFlagsOk returns a tuple with the FeatureFlags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Project) GetFeatureFlagsOk() (*ProjectFeatureFlags, bool) {
	if o == nil || o.FeatureFlags == nil {
		return nil, false
	}
	return o.FeatureFlags, true
}

// HasFeatureFlags returns a boolean if a field has been set.
func (o *Project) HasFeatureFlags() bool {
	if o != nil && o.FeatureFlags != nil {
		return true
	}

	return false
}

// SetFeatureFlags gets a reference to the given ProjectFeatureFlags and assigns it to the FeatureFlags field.
func (o *Project) SetFeatureFlags(v ProjectFeatureFlags) {
	o.FeatureFlags = &v
}

// GetFlowdockApiToken returns the FlowdockApiToken field value if set, zero value otherwise.
func (o *Project) GetFlowdockApiToken() map[string]interface{} {
	if o == nil || o.FlowdockApiToken == nil {
		var ret map[string]interface{}
		return ret
	}
	return *o.FlowdockApiToken
}

// GetFlowdockApiTokenOk returns a tuple with the FlowdockApiToken field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Project) GetFlowdockApiTokenOk() (*map[string]interface{}, bool) {
	if o == nil || o.FlowdockApiToken == nil {
		return nil, false
	}
	return o.FlowdockApiToken, true
}

// HasFlowdockApiToken returns a boolean if a field has been set.
func (o *Project) HasFlowdockApiToken() bool {
	if o != nil && o.FlowdockApiToken != nil {
		return true
	}

	return false
}

// SetFlowdockApiToken gets a reference to the given map[string]interface{} and assigns it to the FlowdockApiToken field.
func (o *Project) SetFlowdockApiToken(v map[string]interface{}) {
	o.FlowdockApiToken = &v
}

// GetFollowed returns the Followed field value if set, zero value otherwise.
func (o *Project) GetFollowed() bool {
	if o == nil || o.Followed == nil {
		var ret bool
		return ret
	}
	return *o.Followed
}

// GetFollowedOk returns a tuple with the Followed field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Project) GetFollowedOk() (*bool, bool) {
	if o == nil || o.Followed == nil {
		return nil, false
	}
	return o.Followed, true
}

// HasFollowed returns a boolean if a field has been set.
func (o *Project) HasFollowed() bool {
	if o != nil && o.Followed != nil {
		return true
	}

	return false
}

// SetFollowed gets a reference to the given bool and assigns it to the Followed field.
func (o *Project) SetFollowed(v bool) {
	o.Followed = &v
}

// GetHasUsableKey returns the HasUsableKey field value if set, zero value otherwise.
func (o *Project) GetHasUsableKey() bool {
	if o == nil || o.HasUsableKey == nil {
		var ret bool
		return ret
	}
	return *o.HasUsableKey
}

// GetHasUsableKeyOk returns a tuple with the HasUsableKey field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Project) GetHasUsableKeyOk() (*bool, bool) {
	if o == nil || o.HasUsableKey == nil {
		return nil, false
	}
	return o.HasUsableKey, true
}

// HasHasUsableKey returns a boolean if a field has been set.
func (o *Project) HasHasUsableKey() bool {
	if o != nil && o.HasUsableKey != nil {
		return true
	}

	return false
}

// SetHasUsableKey gets a reference to the given bool and assigns it to the HasUsableKey field.
func (o *Project) SetHasUsableKey(v bool) {
	o.HasUsableKey = &v
}

// GetHerokuDeployUser returns the HerokuDeployUser field value if set, zero value otherwise.
func (o *Project) GetHerokuDeployUser() map[string]interface{} {
	if o == nil || o.HerokuDeployUser == nil {
		var ret map[string]interface{}
		return ret
	}
	return *o.HerokuDeployUser
}

// GetHerokuDeployUserOk returns a tuple with the HerokuDeployUser field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Project) GetHerokuDeployUserOk() (*map[string]interface{}, bool) {
	if o == nil || o.HerokuDeployUser == nil {
		return nil, false
	}
	return o.HerokuDeployUser, true
}

// HasHerokuDeployUser returns a boolean if a field has been set.
func (o *Project) HasHerokuDeployUser() bool {
	if o != nil && o.HerokuDeployUser != nil {
		return true
	}

	return false
}

// SetHerokuDeployUser gets a reference to the given map[string]interface{} and assigns it to the HerokuDeployUser field.
func (o *Project) SetHerokuDeployUser(v map[string]interface{}) {
	o.HerokuDeployUser = &v
}

// GetHipchatApiToken returns the HipchatApiToken field value if set, zero value otherwise.
func (o *Project) GetHipchatApiToken() map[string]interface{} {
	if o == nil || o.HipchatApiToken == nil {
		var ret map[string]interface{}
		return ret
	}
	return *o.HipchatApiToken
}

// GetHipchatApiTokenOk returns a tuple with the HipchatApiToken field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Project) GetHipchatApiTokenOk() (*map[string]interface{}, bool) {
	if o == nil || o.HipchatApiToken == nil {
		return nil, false
	}
	return o.HipchatApiToken, true
}

// HasHipchatApiToken returns a boolean if a field has been set.
func (o *Project) HasHipchatApiToken() bool {
	if o != nil && o.HipchatApiToken != nil {
		return true
	}

	return false
}

// SetHipchatApiToken gets a reference to the given map[string]interface{} and assigns it to the HipchatApiToken field.
func (o *Project) SetHipchatApiToken(v map[string]interface{}) {
	o.HipchatApiToken = &v
}

// GetHipchatNotify returns the HipchatNotify field value if set, zero value otherwise.
func (o *Project) GetHipchatNotify() map[string]interface{} {
	if o == nil || o.HipchatNotify == nil {
		var ret map[string]interface{}
		return ret
	}
	return *o.HipchatNotify
}

// GetHipchatNotifyOk returns a tuple with the HipchatNotify field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Project) GetHipchatNotifyOk() (*map[string]interface{}, bool) {
	if o == nil || o.HipchatNotify == nil {
		return nil, false
	}
	return o.HipchatNotify, true
}

// HasHipchatNotify returns a boolean if a field has been set.
func (o *Project) HasHipchatNotify() bool {
	if o != nil && o.HipchatNotify != nil {
		return true
	}

	return false
}

// SetHipchatNotify gets a reference to the given map[string]interface{} and assigns it to the HipchatNotify field.
func (o *Project) SetHipchatNotify(v map[string]interface{}) {
	o.HipchatNotify = &v
}

// GetHipchatRoom returns the HipchatRoom field value if set, zero value otherwise.
func (o *Project) GetHipchatRoom() map[string]interface{} {
	if o == nil || o.HipchatRoom == nil {
		var ret map[string]interface{}
		return ret
	}
	return *o.HipchatRoom
}

// GetHipchatRoomOk returns a tuple with the HipchatRoom field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Project) GetHipchatRoomOk() (*map[string]interface{}, bool) {
	if o == nil || o.HipchatRoom == nil {
		return nil, false
	}
	return o.HipchatRoom, true
}

// HasHipchatRoom returns a boolean if a field has been set.
func (o *Project) HasHipchatRoom() bool {
	if o != nil && o.HipchatRoom != nil {
		return true
	}

	return false
}

// SetHipchatRoom gets a reference to the given map[string]interface{} and assigns it to the HipchatRoom field.
func (o *Project) SetHipchatRoom(v map[string]interface{}) {
	o.HipchatRoom = &v
}

// GetIrcChannel returns the IrcChannel field value if set, zero value otherwise.
func (o *Project) GetIrcChannel() map[string]interface{} {
	if o == nil || o.IrcChannel == nil {
		var ret map[string]interface{}
		return ret
	}
	return *o.IrcChannel
}

// GetIrcChannelOk returns a tuple with the IrcChannel field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Project) GetIrcChannelOk() (*map[string]interface{}, bool) {
	if o == nil || o.IrcChannel == nil {
		return nil, false
	}
	return o.IrcChannel, true
}

// HasIrcChannel returns a boolean if a field has been set.
func (o *Project) HasIrcChannel() bool {
	if o != nil && o.IrcChannel != nil {
		return true
	}

	return false
}

// SetIrcChannel gets a reference to the given map[string]interface{} and assigns it to the IrcChannel field.
func (o *Project) SetIrcChannel(v map[string]interface{}) {
	o.IrcChannel = &v
}

// GetIrcKeyword returns the IrcKeyword field value if set, zero value otherwise.
func (o *Project) GetIrcKeyword() map[string]interface{} {
	if o == nil || o.IrcKeyword == nil {
		var ret map[string]interface{}
		return ret
	}
	return *o.IrcKeyword
}

// GetIrcKeywordOk returns a tuple with the IrcKeyword field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Project) GetIrcKeywordOk() (*map[string]interface{}, bool) {
	if o == nil || o.IrcKeyword == nil {
		return nil, false
	}
	return o.IrcKeyword, true
}

// HasIrcKeyword returns a boolean if a field has been set.
func (o *Project) HasIrcKeyword() bool {
	if o != nil && o.IrcKeyword != nil {
		return true
	}

	return false
}

// SetIrcKeyword gets a reference to the given map[string]interface{} and assigns it to the IrcKeyword field.
func (o *Project) SetIrcKeyword(v map[string]interface{}) {
	o.IrcKeyword = &v
}

// GetIrcNotifyPrefs returns the IrcNotifyPrefs field value if set, zero value otherwise.
func (o *Project) GetIrcNotifyPrefs() map[string]interface{} {
	if o == nil || o.IrcNotifyPrefs == nil {
		var ret map[string]interface{}
		return ret
	}
	return *o.IrcNotifyPrefs
}

// GetIrcNotifyPrefsOk returns a tuple with the IrcNotifyPrefs field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Project) GetIrcNotifyPrefsOk() (*map[string]interface{}, bool) {
	if o == nil || o.IrcNotifyPrefs == nil {
		return nil, false
	}
	return o.IrcNotifyPrefs, true
}

// HasIrcNotifyPrefs returns a boolean if a field has been set.
func (o *Project) HasIrcNotifyPrefs() bool {
	if o != nil && o.IrcNotifyPrefs != nil {
		return true
	}

	return false
}

// SetIrcNotifyPrefs gets a reference to the given map[string]interface{} and assigns it to the IrcNotifyPrefs field.
func (o *Project) SetIrcNotifyPrefs(v map[string]interface{}) {
	o.IrcNotifyPrefs = &v
}

// GetIrcPassword returns the IrcPassword field value if set, zero value otherwise.
func (o *Project) GetIrcPassword() map[string]interface{} {
	if o == nil || o.IrcPassword == nil {
		var ret map[string]interface{}
		return ret
	}
	return *o.IrcPassword
}

// GetIrcPasswordOk returns a tuple with the IrcPassword field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Project) GetIrcPasswordOk() (*map[string]interface{}, bool) {
	if o == nil || o.IrcPassword == nil {
		return nil, false
	}
	return o.IrcPassword, true
}

// HasIrcPassword returns a boolean if a field has been set.
func (o *Project) HasIrcPassword() bool {
	if o != nil && o.IrcPassword != nil {
		return true
	}

	return false
}

// SetIrcPassword gets a reference to the given map[string]interface{} and assigns it to the IrcPassword field.
func (o *Project) SetIrcPassword(v map[string]interface{}) {
	o.IrcPassword = &v
}

// GetIrcServer returns the IrcServer field value if set, zero value otherwise.
func (o *Project) GetIrcServer() map[string]interface{} {
	if o == nil || o.IrcServer == nil {
		var ret map[string]interface{}
		return ret
	}
	return *o.IrcServer
}

// GetIrcServerOk returns a tuple with the IrcServer field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Project) GetIrcServerOk() (*map[string]interface{}, bool) {
	if o == nil || o.IrcServer == nil {
		return nil, false
	}
	return o.IrcServer, true
}

// HasIrcServer returns a boolean if a field has been set.
func (o *Project) HasIrcServer() bool {
	if o != nil && o.IrcServer != nil {
		return true
	}

	return false
}

// SetIrcServer gets a reference to the given map[string]interface{} and assigns it to the IrcServer field.
func (o *Project) SetIrcServer(v map[string]interface{}) {
	o.IrcServer = &v
}

// GetIrcUsername returns the IrcUsername field value if set, zero value otherwise.
func (o *Project) GetIrcUsername() map[string]interface{} {
	if o == nil || o.IrcUsername == nil {
		var ret map[string]interface{}
		return ret
	}
	return *o.IrcUsername
}

// GetIrcUsernameOk returns a tuple with the IrcUsername field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Project) GetIrcUsernameOk() (*map[string]interface{}, bool) {
	if o == nil || o.IrcUsername == nil {
		return nil, false
	}
	return o.IrcUsername, true
}

// HasIrcUsername returns a boolean if a field has been set.
func (o *Project) HasIrcUsername() bool {
	if o != nil && o.IrcUsername != nil {
		return true
	}

	return false
}

// SetIrcUsername gets a reference to the given map[string]interface{} and assigns it to the IrcUsername field.
func (o *Project) SetIrcUsername(v map[string]interface{}) {
	o.IrcUsername = &v
}

// GetLanguage returns the Language field value if set, zero value otherwise.
func (o *Project) GetLanguage() string {
	if o == nil || o.Language == nil {
		var ret string
		return ret
	}
	return *o.Language
}

// GetLanguageOk returns a tuple with the Language field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Project) GetLanguageOk() (*string, bool) {
	if o == nil || o.Language == nil {
		return nil, false
	}
	return o.Language, true
}

// HasLanguage returns a boolean if a field has been set.
func (o *Project) HasLanguage() bool {
	if o != nil && o.Language != nil {
		return true
	}

	return false
}

// SetLanguage gets a reference to the given string and assigns it to the Language field.
func (o *Project) SetLanguage(v string) {
	o.Language = &v
}

// GetOss returns the Oss field value if set, zero value otherwise.
func (o *Project) GetOss() bool {
	if o == nil || o.Oss == nil {
		var ret bool
		return ret
	}
	return *o.Oss
}

// GetOssOk returns a tuple with the Oss field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Project) GetOssOk() (*bool, bool) {
	if o == nil || o.Oss == nil {
		return nil, false
	}
	return o.Oss, true
}

// HasOss returns a boolean if a field has been set.
func (o *Project) HasOss() bool {
	if o != nil && o.Oss != nil {
		return true
	}

	return false
}

// SetOss gets a reference to the given bool and assigns it to the Oss field.
func (o *Project) SetOss(v bool) {
	o.Oss = &v
}

// GetParallel returns the Parallel field value if set, zero value otherwise.
func (o *Project) GetParallel() int32 {
	if o == nil || o.Parallel == nil {
		var ret int32
		return ret
	}
	return *o.Parallel
}

// GetParallelOk returns a tuple with the Parallel field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Project) GetParallelOk() (*int32, bool) {
	if o == nil || o.Parallel == nil {
		return nil, false
	}
	return o.Parallel, true
}

// HasParallel returns a boolean if a field has been set.
func (o *Project) HasParallel() bool {
	if o != nil && o.Parallel != nil {
		return true
	}

	return false
}

// SetParallel gets a reference to the given int32 and assigns it to the Parallel field.
func (o *Project) SetParallel(v int32) {
	o.Parallel = &v
}

// GetReponame returns the Reponame field value if set, zero value otherwise.
func (o *Project) GetReponame() string {
	if o == nil || o.Reponame == nil {
		var ret string
		return ret
	}
	return *o.Reponame
}

// GetReponameOk returns a tuple with the Reponame field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Project) GetReponameOk() (*string, bool) {
	if o == nil || o.Reponame == nil {
		return nil, false
	}
	return o.Reponame, true
}

// HasReponame returns a boolean if a field has been set.
func (o *Project) HasReponame() bool {
	if o != nil && o.Reponame != nil {
		return true
	}

	return false
}

// SetReponame gets a reference to the given string and assigns it to the Reponame field.
func (o *Project) SetReponame(v string) {
	o.Reponame = &v
}

// GetScopes returns the Scopes field value if set, zero value otherwise.
func (o *Project) GetScopes() []Scope {
	if o == nil || o.Scopes == nil {
		var ret []Scope
		return ret
	}
	return *o.Scopes
}

// GetScopesOk returns a tuple with the Scopes field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Project) GetScopesOk() (*[]Scope, bool) {
	if o == nil || o.Scopes == nil {
		return nil, false
	}
	return o.Scopes, true
}

// HasScopes returns a boolean if a field has been set.
func (o *Project) HasScopes() bool {
	if o != nil && o.Scopes != nil {
		return true
	}

	return false
}

// SetScopes gets a reference to the given []Scope and assigns it to the Scopes field.
func (o *Project) SetScopes(v []Scope) {
	o.Scopes = &v
}

// GetSetup returns the Setup field value if set, zero value otherwise.
func (o *Project) GetSetup() string {
	if o == nil || o.Setup == nil {
		var ret string
		return ret
	}
	return *o.Setup
}

// GetSetupOk returns a tuple with the Setup field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Project) GetSetupOk() (*string, bool) {
	if o == nil || o.Setup == nil {
		return nil, false
	}
	return o.Setup, true
}

// HasSetup returns a boolean if a field has been set.
func (o *Project) HasSetup() bool {
	if o != nil && o.Setup != nil {
		return true
	}

	return false
}

// SetSetup gets a reference to the given string and assigns it to the Setup field.
func (o *Project) SetSetup(v string) {
	o.Setup = &v
}

// GetSlackApiToken returns the SlackApiToken field value if set, zero value otherwise.
func (o *Project) GetSlackApiToken() map[string]interface{} {
	if o == nil || o.SlackApiToken == nil {
		var ret map[string]interface{}
		return ret
	}
	return *o.SlackApiToken
}

// GetSlackApiTokenOk returns a tuple with the SlackApiToken field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Project) GetSlackApiTokenOk() (*map[string]interface{}, bool) {
	if o == nil || o.SlackApiToken == nil {
		return nil, false
	}
	return o.SlackApiToken, true
}

// HasSlackApiToken returns a boolean if a field has been set.
func (o *Project) HasSlackApiToken() bool {
	if o != nil && o.SlackApiToken != nil {
		return true
	}

	return false
}

// SetSlackApiToken gets a reference to the given map[string]interface{} and assigns it to the SlackApiToken field.
func (o *Project) SetSlackApiToken(v map[string]interface{}) {
	o.SlackApiToken = &v
}

// GetSlackChannel returns the SlackChannel field value if set, zero value otherwise.
func (o *Project) GetSlackChannel() map[string]interface{} {
	if o == nil || o.SlackChannel == nil {
		var ret map[string]interface{}
		return ret
	}
	return *o.SlackChannel
}

// GetSlackChannelOk returns a tuple with the SlackChannel field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Project) GetSlackChannelOk() (*map[string]interface{}, bool) {
	if o == nil || o.SlackChannel == nil {
		return nil, false
	}
	return o.SlackChannel, true
}

// HasSlackChannel returns a boolean if a field has been set.
func (o *Project) HasSlackChannel() bool {
	if o != nil && o.SlackChannel != nil {
		return true
	}

	return false
}

// SetSlackChannel gets a reference to the given map[string]interface{} and assigns it to the SlackChannel field.
func (o *Project) SetSlackChannel(v map[string]interface{}) {
	o.SlackChannel = &v
}

// GetSlackChannelOverride returns the SlackChannelOverride field value if set, zero value otherwise.
func (o *Project) GetSlackChannelOverride() map[string]interface{} {
	if o == nil || o.SlackChannelOverride == nil {
		var ret map[string]interface{}
		return ret
	}
	return *o.SlackChannelOverride
}

// GetSlackChannelOverrideOk returns a tuple with the SlackChannelOverride field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Project) GetSlackChannelOverrideOk() (*map[string]interface{}, bool) {
	if o == nil || o.SlackChannelOverride == nil {
		return nil, false
	}
	return o.SlackChannelOverride, true
}

// HasSlackChannelOverride returns a boolean if a field has been set.
func (o *Project) HasSlackChannelOverride() bool {
	if o != nil && o.SlackChannelOverride != nil {
		return true
	}

	return false
}

// SetSlackChannelOverride gets a reference to the given map[string]interface{} and assigns it to the SlackChannelOverride field.
func (o *Project) SetSlackChannelOverride(v map[string]interface{}) {
	o.SlackChannelOverride = &v
}

// GetSlackNotifyPrefs returns the SlackNotifyPrefs field value if set, zero value otherwise.
func (o *Project) GetSlackNotifyPrefs() map[string]interface{} {
	if o == nil || o.SlackNotifyPrefs == nil {
		var ret map[string]interface{}
		return ret
	}
	return *o.SlackNotifyPrefs
}

// GetSlackNotifyPrefsOk returns a tuple with the SlackNotifyPrefs field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Project) GetSlackNotifyPrefsOk() (*map[string]interface{}, bool) {
	if o == nil || o.SlackNotifyPrefs == nil {
		return nil, false
	}
	return o.SlackNotifyPrefs, true
}

// HasSlackNotifyPrefs returns a boolean if a field has been set.
func (o *Project) HasSlackNotifyPrefs() bool {
	if o != nil && o.SlackNotifyPrefs != nil {
		return true
	}

	return false
}

// SetSlackNotifyPrefs gets a reference to the given map[string]interface{} and assigns it to the SlackNotifyPrefs field.
func (o *Project) SetSlackNotifyPrefs(v map[string]interface{}) {
	o.SlackNotifyPrefs = &v
}

// GetSlackSubdomain returns the SlackSubdomain field value if set, zero value otherwise.
func (o *Project) GetSlackSubdomain() map[string]interface{} {
	if o == nil || o.SlackSubdomain == nil {
		var ret map[string]interface{}
		return ret
	}
	return *o.SlackSubdomain
}

// GetSlackSubdomainOk returns a tuple with the SlackSubdomain field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Project) GetSlackSubdomainOk() (*map[string]interface{}, bool) {
	if o == nil || o.SlackSubdomain == nil {
		return nil, false
	}
	return o.SlackSubdomain, true
}

// HasSlackSubdomain returns a boolean if a field has been set.
func (o *Project) HasSlackSubdomain() bool {
	if o != nil && o.SlackSubdomain != nil {
		return true
	}

	return false
}

// SetSlackSubdomain gets a reference to the given map[string]interface{} and assigns it to the SlackSubdomain field.
func (o *Project) SetSlackSubdomain(v map[string]interface{}) {
	o.SlackSubdomain = &v
}

// GetSlackWebhookUrl returns the SlackWebhookUrl field value if set, zero value otherwise.
func (o *Project) GetSlackWebhookUrl() string {
	if o == nil || o.SlackWebhookUrl == nil {
		var ret string
		return ret
	}
	return *o.SlackWebhookUrl
}

// GetSlackWebhookUrlOk returns a tuple with the SlackWebhookUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Project) GetSlackWebhookUrlOk() (*string, bool) {
	if o == nil || o.SlackWebhookUrl == nil {
		return nil, false
	}
	return o.SlackWebhookUrl, true
}

// HasSlackWebhookUrl returns a boolean if a field has been set.
func (o *Project) HasSlackWebhookUrl() bool {
	if o != nil && o.SlackWebhookUrl != nil {
		return true
	}

	return false
}

// SetSlackWebhookUrl gets a reference to the given string and assigns it to the SlackWebhookUrl field.
func (o *Project) SetSlackWebhookUrl(v string) {
	o.SlackWebhookUrl = &v
}

// GetSshKeys returns the SshKeys field value if set, zero value otherwise.
func (o *Project) GetSshKeys() []string {
	if o == nil || o.SshKeys == nil {
		var ret []string
		return ret
	}
	return *o.SshKeys
}

// GetSshKeysOk returns a tuple with the SshKeys field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Project) GetSshKeysOk() (*[]string, bool) {
	if o == nil || o.SshKeys == nil {
		return nil, false
	}
	return o.SshKeys, true
}

// HasSshKeys returns a boolean if a field has been set.
func (o *Project) HasSshKeys() bool {
	if o != nil && o.SshKeys != nil {
		return true
	}

	return false
}

// SetSshKeys gets a reference to the given []string and assigns it to the SshKeys field.
func (o *Project) SetSshKeys(v []string) {
	o.SshKeys = &v
}

// GetTest returns the Test field value if set, zero value otherwise.
func (o *Project) GetTest() string {
	if o == nil || o.Test == nil {
		var ret string
		return ret
	}
	return *o.Test
}

// GetTestOk returns a tuple with the Test field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Project) GetTestOk() (*string, bool) {
	if o == nil || o.Test == nil {
		return nil, false
	}
	return o.Test, true
}

// HasTest returns a boolean if a field has been set.
func (o *Project) HasTest() bool {
	if o != nil && o.Test != nil {
		return true
	}

	return false
}

// SetTest gets a reference to the given string and assigns it to the Test field.
func (o *Project) SetTest(v string) {
	o.Test = &v
}

// GetUsername returns the Username field value if set, zero value otherwise.
func (o *Project) GetUsername() string {
	if o == nil || o.Username == nil {
		var ret string
		return ret
	}
	return *o.Username
}

// GetUsernameOk returns a tuple with the Username field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Project) GetUsernameOk() (*string, bool) {
	if o == nil || o.Username == nil {
		return nil, false
	}
	return o.Username, true
}

// HasUsername returns a boolean if a field has been set.
func (o *Project) HasUsername() bool {
	if o != nil && o.Username != nil {
		return true
	}

	return false
}

// SetUsername gets a reference to the given string and assigns it to the Username field.
func (o *Project) SetUsername(v string) {
	o.Username = &v
}

// GetVcsType returns the VcsType field value if set, zero value otherwise.
func (o *Project) GetVcsType() string {
	if o == nil || o.VcsType == nil {
		var ret string
		return ret
	}
	return *o.VcsType
}

// GetVcsTypeOk returns a tuple with the VcsType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Project) GetVcsTypeOk() (*string, bool) {
	if o == nil || o.VcsType == nil {
		return nil, false
	}
	return o.VcsType, true
}

// HasVcsType returns a boolean if a field has been set.
func (o *Project) HasVcsType() bool {
	if o != nil && o.VcsType != nil {
		return true
	}

	return false
}

// SetVcsType gets a reference to the given string and assigns it to the VcsType field.
func (o *Project) SetVcsType(v string) {
	o.VcsType = &v
}

// GetVcsUrl returns the VcsUrl field value if set, zero value otherwise.
func (o *Project) GetVcsUrl() string {
	if o == nil || o.VcsUrl == nil {
		var ret string
		return ret
	}
	return *o.VcsUrl
}

// GetVcsUrlOk returns a tuple with the VcsUrl field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *Project) GetVcsUrlOk() (*string, bool) {
	if o == nil || o.VcsUrl == nil {
		return nil, false
	}
	return o.VcsUrl, true
}

// HasVcsUrl returns a boolean if a field has been set.
func (o *Project) HasVcsUrl() bool {
	if o != nil && o.VcsUrl != nil {
		return true
	}

	return false
}

// SetVcsUrl gets a reference to the given string and assigns it to the VcsUrl field.
func (o *Project) SetVcsUrl(v string) {
	o.VcsUrl = &v
}

func (o Project) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Aws != nil {
		toSerialize["aws"] = o.Aws
	}
	if o.Branches != nil {
		toSerialize["branches"] = o.Branches
	}
	if o.CampfireNotifyPrefs != nil {
		toSerialize["campfire_notify_prefs"] = o.CampfireNotifyPrefs
	}
	if o.CampfireRoom != nil {
		toSerialize["campfire_room"] = o.CampfireRoom
	}
	if o.CampfireSubdomain != nil {
		toSerialize["campfire_subdomain"] = o.CampfireSubdomain
	}
	if o.CampfireToken != nil {
		toSerialize["campfire_token"] = o.CampfireToken
	}
	if o.Compile != nil {
		toSerialize["compile"] = o.Compile
	}
	if o.DefaultBranch != nil {
		toSerialize["default_branch"] = o.DefaultBranch
	}
	if o.Dependencies != nil {
		toSerialize["dependencies"] = o.Dependencies
	}
	if o.Extra != nil {
		toSerialize["extra"] = o.Extra
	}
	if o.FeatureFlags != nil {
		toSerialize["feature_flags"] = o.FeatureFlags
	}
	if o.FlowdockApiToken != nil {
		toSerialize["flowdock_api_token"] = o.FlowdockApiToken
	}
	if o.Followed != nil {
		toSerialize["followed"] = o.Followed
	}
	if o.HasUsableKey != nil {
		toSerialize["has_usable_key"] = o.HasUsableKey
	}
	if o.HerokuDeployUser != nil {
		toSerialize["heroku_deploy_user"] = o.HerokuDeployUser
	}
	if o.HipchatApiToken != nil {
		toSerialize["hipchat_api_token"] = o.HipchatApiToken
	}
	if o.HipchatNotify != nil {
		toSerialize["hipchat_notify"] = o.HipchatNotify
	}
	if o.HipchatRoom != nil {
		toSerialize["hipchat_room"] = o.HipchatRoom
	}
	if o.IrcChannel != nil {
		toSerialize["irc_channel"] = o.IrcChannel
	}
	if o.IrcKeyword != nil {
		toSerialize["irc_keyword"] = o.IrcKeyword
	}
	if o.IrcNotifyPrefs != nil {
		toSerialize["irc_notify_prefs"] = o.IrcNotifyPrefs
	}
	if o.IrcPassword != nil {
		toSerialize["irc_password"] = o.IrcPassword
	}
	if o.IrcServer != nil {
		toSerialize["irc_server"] = o.IrcServer
	}
	if o.IrcUsername != nil {
		toSerialize["irc_username"] = o.IrcUsername
	}
	if o.Language != nil {
		toSerialize["language"] = o.Language
	}
	if o.Oss != nil {
		toSerialize["oss"] = o.Oss
	}
	if o.Parallel != nil {
		toSerialize["parallel"] = o.Parallel
	}
	if o.Reponame != nil {
		toSerialize["reponame"] = o.Reponame
	}
	if o.Scopes != nil {
		toSerialize["scopes"] = o.Scopes
	}
	if o.Setup != nil {
		toSerialize["setup"] = o.Setup
	}
	if o.SlackApiToken != nil {
		toSerialize["slack_api_token"] = o.SlackApiToken
	}
	if o.SlackChannel != nil {
		toSerialize["slack_channel"] = o.SlackChannel
	}
	if o.SlackChannelOverride != nil {
		toSerialize["slack_channel_override"] = o.SlackChannelOverride
	}
	if o.SlackNotifyPrefs != nil {
		toSerialize["slack_notify_prefs"] = o.SlackNotifyPrefs
	}
	if o.SlackSubdomain != nil {
		toSerialize["slack_subdomain"] = o.SlackSubdomain
	}
	if o.SlackWebhookUrl != nil {
		toSerialize["slack_webhook_url"] = o.SlackWebhookUrl
	}
	if o.SshKeys != nil {
		toSerialize["ssh_keys"] = o.SshKeys
	}
	if o.Test != nil {
		toSerialize["test"] = o.Test
	}
	if o.Username != nil {
		toSerialize["username"] = o.Username
	}
	if o.VcsType != nil {
		toSerialize["vcs_type"] = o.VcsType
	}
	if o.VcsUrl != nil {
		toSerialize["vcs_url"] = o.VcsUrl
	}
	return json.Marshal(toSerialize)
}

type NullableProject struct {
	value *Project
	isSet bool
}

func (v NullableProject) Get() *Project {
	return v.value
}

func (v *NullableProject) Set(val *Project) {
	v.value = val
	v.isSet = true
}

func (v NullableProject) IsSet() bool {
	return v.isSet
}

func (v *NullableProject) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableProject(val *Project) *NullableProject {
	return &NullableProject{value: val, isSet: true}
}

func (v NullableProject) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableProject) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
