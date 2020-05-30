// Copyright 2020 The ccictl Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package v1

import (
	"fmt"
	"time"

	json "github.com/goccy/go-json"
)

// Artifact represents a Artifact.
type Artifact struct {
	NodeIndex  int32  `json:"node_index,omitempty"`
	Path       string `json:"path,omitempty"`
	PrettyPath string `json:"pretty_path,omitempty"`
	URL        string `json:"url,omitempty"`
}

// Build represents a Build.
type Build struct {
	// Body commit message body
	Body            string         `json:"body,omitempty"`
	Branch          string         `json:"branch,omitempty"`
	BuildTimeMillis int32          `json:"build_time_millis,omitempty"`
	BuildURL        string         `json:"build_url,omitempty"`
	CommitterEmail  string         `json:"committer_email,omitempty"`
	CommitterName   string         `json:"committer_name,omitempty"`
	DontBuild       string         `json:"dont_build,omitempty"`
	Lifecycle       Lifecycle      `json:"lifecycle,omitempty"`
	Previous        *PreviousBuild `json:"previous,omitempty"`
	// QueuedAt time build was queued
	QueuedAt time.Time `json:"queued_at,omitempty"`
	Reponame string    `json:"reponame,omitempty"`
	RetryOf  int32     `json:"retry_of,omitempty"`
	// StartTime time build started
	StartTime time.Time `json:"start_time,omitempty"`
	// StopTime time build finished
	StopTime time.Time `json:"stop_time,omitempty"`
	Subject  string    `json:"subject,omitempty"`
	Username string    `json:"username,omitempty"`
	VCSURL   string    `json:"vcs_url,omitempty"`
	// Why short string explaining the reason we built
	Why string `json:"why,omitempty"`
}

// BuildDetail previous build
type BuildDetail struct {
	AllCommitDetails        []*CommitDetail `json:"all_commit_details,omitempty"`
	JobName                 string          `json:"job_name,omitempty"`
	PreviousSuccessfulBuild *PreviousBuild  `json:"previous_successful_build,omitempty"`
	UsageQueuedAt           time.Time       `json:"usage_queued_at,omitempty"`
	User                    *User           `json:"user,omitempty"`
}

// BuildSummary represents a BuildSummary.
type BuildSummary struct {
	AddedAt     time.Time `json:"added_at,omitempty"`
	BuildNum    int32     `json:"build_num,omitempty"`
	Outcome     *Outcome  `json:"outcome,omitempty"`
	PushedAt    time.Time `json:"pushed_at,omitempty"`
	Status      *Status   `json:"status,omitempty"`
	VCSRevision string    `json:"vcs_revision,omitempty"`
}

// CommitDetail represents a CommitDetail.
type CommitDetail struct {
	AuthorDate     time.Time `json:"author_date,omitempty"`
	AuthorEmail    string    `json:"author_email,omitempty"`
	AuthorLogin    string    `json:"author_login,omitempty"`
	AuthorName     string    `json:"author_name,omitempty"`
	Body           string    `json:"body,omitempty"`
	Commit         string    `json:"commit,omitempty"`
	CommitURL      string    `json:"commit_url,omitempty"`
	CommitterDate  time.Time `json:"committer_date,omitempty"`
	CommitterEmail string    `json:"committer_email,omitempty"`
	CommitterLogin string    `json:"committer_login,omitempty"`
	CommitterName  string    `json:"committer_name,omitempty"`
	Subject        string    `json:"subject,omitempty"`
}

// Envvar represents a Envvar.
type Envvar struct {
	Name  string `json:"name,omitempty"`
	Value string `json:"value,omitempty"`
}

// CheckoutKey represents a CheckoutKey.
type CheckoutKey struct {
	Fingerprint string `json:"fingerprint,omitempty"`
	Preferred   bool   `json:"preferred,omitempty"`
	PublicKey   string `json:"public_key,omitempty"`
	// when the key was issued
	Time time.Time `json:"time,omitempty"`
	// can be "deploy-key" or "github-user-key"
	Type string `json:"type,omitempty"`
}

// Lifecycle the model 'Lifecycle'
type Lifecycle uint8

// List of Lifecycle
const (
	LifecycleQueued     = Lifecycle(1) + iota // "queued"
	LifecycleScheduled                        // "scheduled"
	LifecycleNotRun                           // "not_run"
	LifecycleNotRunning                       // "not_running"
	LifecycleRunning                          // "running"
	LifecycleFinished                         // "finished"
	lifecycleEnd
)

func (v Lifecycle) String() string {
	switch v {
	case LifecycleQueued:
		return "queued"
	case LifecycleScheduled:
		return "scheduled"
	case LifecycleNotRun:
		return "not_run"
	case LifecycleNotRunning:
		return "not_running"
	case LifecycleRunning:
		return "running"
	case LifecycleFinished:
		return "finished"
	default:
		return ""
	}
}

// UnmarshalJSON implements json.Unmarshaler.
func (v *Lifecycle) UnmarshalJSON(src []byte) error {
	var i Lifecycle
	err := json.Unmarshal(src, &i)
	if err != nil {
		return err
	}
	if lifecycleEnd >= i {
		return fmt.Errorf("%+v is not a valid Lifecycle", i)
	}
	*v = i
	return nil
}

// Outcome the model 'Outcome'
type Outcome uint8

// List of Outcome
const (
	OutcomeCanceled           = Outcome(1) + iota // "canceled"
	OutcomeInfrastructureFail                     // "infrastructure_fail"
	OutcomeTimedout                               // "timedout"
	OutcomeFailed                                 // "failed"
	OutcomeNoTests                                // "no_tests"
	OutcomeSuccess                                // "success"
	outcomeEnd
)

func (v Outcome) String() string {
	switch v {
	case OutcomeCanceled:
		return "canceled"
	case OutcomeInfrastructureFail:
		return "infrastructure_fail"
	case OutcomeTimedout:
		return "timedout"
	case OutcomeFailed:
		return "failed"
	case OutcomeNoTests:
		return "no_tests"
	case OutcomeSuccess:
		return "success"
	default:
		return ""
	}
}

// UnmarshalJSON implements json.Unmarshaler.
func (v *Outcome) UnmarshalJSON(src []byte) error {
	var i Outcome
	err := json.Unmarshal(src, &i)
	if err != nil {
		return err
	}
	if outcomeEnd >= i {
		return fmt.Errorf("%+v is not a valid Outcome", i)
	}
	*v = i
	return nil
}

// PreviousBuild previous build
type PreviousBuild struct {
	BuildNum        int32  `json:"build_num,omitempty"`
	BuildTimeMillis int32  `json:"build_time_millis,omitempty"`
	Status          Status `json:"status,omitempty"`
}

// Project represents a Project.
type Project struct {
	AWS                  map[string]string      `json:"aws,omitempty"`
	Branches             map[string]interface{} `json:"branches,omitempty"`
	CampfireNotifyPrefs  string                 `json:"campfire_notify_prefs,omitempty"`
	CampfireRoom         string                 `json:"campfire_room,omitempty"`
	CampfireSubdomain    string                 `json:"campfire_subdomain,omitempty"`
	CampfireToken        string                 `json:"campfire_token,omitempty"`
	Compile              string                 `json:"compile,omitempty"`
	DefaultBranch        string                 `json:"default_branch,omitempty"`
	Dependencies         string                 `json:"dependencies,omitempty"`
	Extra                string                 `json:"extra,omitempty"`
	FeatureFlags         ProjectFeatureFlags    `json:"feature_flags,omitempty"`
	FlowdockAPIToken     string                 `json:"flowdock_api_token,omitempty"`
	Followed             bool                   `json:"followed,omitempty"`
	HasUsableKey         bool                   `json:"has_usable_key,omitempty"`
	HerokuDeployUser     string                 `json:"heroku_deploy_user,omitempty"`
	HipchatAPIToken      string                 `json:"hipchat_api_token,omitempty"`
	HipchatNotify        string                 `json:"hipchat_notify,omitempty"`
	HipchatRoom          string                 `json:"hipchat_room,omitempty"`
	IRCChannel           string                 `json:"irc_channel,omitempty"`
	IRCKeyword           string                 `json:"irc_keyword,omitempty"`
	IRCNotifyPrefs       string                 `json:"irc_notify_prefs,omitempty"`
	IRCPassword          string                 `json:"irc_password,omitempty"`
	IRCServer            string                 `json:"irc_server,omitempty"`
	IRCUsername          string                 `json:"irc_username,omitempty"`
	Language             string                 `json:"language,omitempty"`
	OSS                  bool                   `json:"oss,omitempty"`
	Parallel             int32                  `json:"parallel,omitempty"`
	Reponame             string                 `json:"reponame,omitempty"`
	Scopes               []Scope                `json:"scopes,omitempty"`
	Setup                string                 `json:"setup,omitempty"`
	SlackAPIToken        string                 `json:"slack_api_token,omitempty"`
	SlackChannel         string                 `json:"slack_channel,omitempty"`
	SlackChannelOverride string                 `json:"slack_channel_override,omitempty"`
	SlackNotifyPrefs     string                 `json:"slack_notify_prefs,omitempty"`
	SlackSubdomain       string                 `json:"slack_subdomain,omitempty"`
	SlackWebhookURL      string                 `json:"slack_webhook_url,omitempty"`
	SSHKeys              []SSHKey               `json:"ssh_keys,omitempty"`
	Test                 string                 `json:"test,omitempty"`
	Username             string                 `json:"username,omitempty"`
	VCSType              string                 `json:"vcs_type,omitempty"`
	VCSURL               string                 `json:"vcs_url,omitempty"`
}

// ProjectFeatureFlags represents a ProjectFeatureFlags.
type ProjectFeatureFlags struct {
	BuildForkPRs    bool `json:"build-fork-prs,omitempty"`
	Junit           bool `json:"junit,omitempty"`
	OSS             bool `json:"oss,omitempty"`
	OSX             bool `json:"osx,omitempty"`
	SetGithubStatus bool `json:"set-github-status,omitempty"`
	TrustyBeta      bool `json:"trusty-beta,omitempty"`
}

// Response represents a response of CircleCI API.
type Response struct {
	Status  string `json:"status,omitempty"`
	Message string `json:"message,omitempty"`
}

// Scope the model Scope
type Scope uint8

// List of Scope
const (
	ScopeWriteSettings = Scope(1) + iota // "write-settings"
	ScopeViewBuilds                      // "view-builds"
	ScopeReadSettings                    // "read-settings"
	ScopeTriggeRBuilds                   // "trigger-builds"
	ScopeAll                             // "all"
	ScopeStatus                          // "status"
	ScopeNone                            // "none"
	scopeEnd
)

func (v Scope) String() string {
	switch v {
	case ScopeWriteSettings:
		return "write-settings"
	case ScopeViewBuilds:
		return "view-builds"
	case ScopeReadSettings:
		return "read-settings"
	case ScopeTriggeRBuilds:
		return "trigger-builds"
	case ScopeAll:
		return "all"
	case ScopeStatus:
		return "status"
	case ScopeNone:
		return "none"
	default:
		return ""
	}
}

// UnmarshalJSON implements json.Unmarshaler.
func (v *Scope) UnmarshalJSON(src []byte) error {
	var i Scope
	err := json.Unmarshal(src, &i)
	if err != nil {
		return err
	}
	if scopeEnd >= i {
		return fmt.Errorf("%+v is not a valid Scope", i)
	}
	*v = i
	return nil
}

// SSHKey represents a ssh hostname and privatekey pair.
type SSHKey struct {
	Hostname   string `json:"hostname"`
	PrivateKey string `json:"private_key"`
}

// Status the model Status
type Status uint8

// List of Status
const (
	StatusRetried            = Status(1) + iota // "retried"
	StatusCanceled                              // "canceled"
	StatusInfrastructureFail                    // "infrastructure_fail"
	StatusTimeout                               // "timedout"
	StatusNotRun                                // "not_run"
	StatusRunning                               // "running"
	StatusFailed                                // "failed"
	StatusQueued                                // "queued"
	StatusScheduled                             // "scheduled"
	StatusNotRunning                            // "not_running"
	StatusNoTests                               // "no_tests"
	StatusFixed                                 // "fixed"
	StatusSuccess                               // "success"
	statusEnd
)

func (v Status) String() string {
	switch v {
	case StatusRetried:
		return "queued"
	case StatusCanceled:
		return "scheduled"
	case StatusInfrastructureFail:
		return "not_run"
	case StatusNotRun:
		return "not_running"
	case StatusRunning:
		return "running"
	case StatusFailed:
		return "finished"
	case StatusQueued:
		return "finished"
	case StatusScheduled:
		return "finished"
	case StatusNotRunning:
		return "finished"
	case StatusNoTests:
		return "finished"
	case StatusFixed:
		return "finished"
	case StatusSuccess:
		return "finished"
	default:
		return ""
	}
}

// UnmarshalJSON implements json.Unmarshaler.
func (v *Status) UnmarshalJSON(src []byte) error {
	var i Status
	err := json.Unmarshal(src, &i)
	if err != nil {
		return err
	}
	if statusEnd >= i {
		return fmt.Errorf("%+v is not a valid Status", i)
	}
	*v = i
	return nil
}

// Test represents a Test.
type Test struct {
	Classname string  `json:"classname,omitempty"`
	File      string  `json:"file,omitempty"`
	Message   string  `json:"message,omitempty"`
	Name      string  `json:"name,omitempty"`
	Result    *Status `json:"result,omitempty"`
	RunTime   float32 `json:"run_time,omitempty"`
	Source    string  `json:"source,omitempty"`
}

// Tests represents a Tests.
type Tests struct {
	Tests []*Test `json:"tests,omitempty"`
}

// TriggerBuild represents a trigger build.
type TriggerBuild struct {
	// BuildParameters is the additional environment variables to inject into the build environment. A map of names to values.
	BuildParameters map[string]interface{} `json:"build_parameters,omitempty"`
	// Parallel is the number of containers to use to run the build. Default is null and the project default is used.
	Parallel string `json:"parallel,omitempty"`
	// Revision is the specific revision to build. Default is null and the head of the branch is used. Cannot be used with tag parameter.
	Revision string `json:"revision,omitempty"`
}

// TriggerBuildSummary represents a trigger build summary.
type TriggerBuildSummary struct {
	// BuildParameters is the additional environment variables to inject into the build environment. A map of names to values.
	BuildParameters map[string]interface{} `json:"build_parameters,omitempty"`
	// Parallel is the number of containers to use to run the build. Default is null and the project default is used.
	Parallel string `json:"parallel,omitempty"`
	// Revision is the specific revision to build. Default is null and the head of the branch is used. Cannot be used with tag parameter.
	Revision string `json:"revision,omitempty"`
	// Tag is the tag to build. Default is null. Cannot be used with revision parameter.
	Tag string `json:"tag,omitempty"`
}

// User represents a User.
type User struct {
	Admin               bool                   `json:"admin,omitempty"`
	AllEmails           []string               `json:"all_emails,omitempty"`
	AnalyticsID         string                 `json:"analytics_id,omitempty"`
	AvatarURL           string                 `json:"avatar_url,omitempty"`
	BasicEmailPrefs     string                 `json:"basic_email_prefs,omitempty"`
	Bitbucket           int32                  `json:"bitbucket,omitempty"`
	BitbucketAuthorized bool                   `json:"bitbucket_authorized,omitempty"`
	Containers          int32                  `json:"containers,omitempty"`
	CreatedAt           time.Time              `json:"created_at,omitempty"`
	DaysLeftInTrial     int32                  `json:"days_left_in_trial,omitempty"`
	DevAdmin            bool                   `json:"dev_admin,omitempty"`
	EnrolledBetas       []string               `json:"enrolled_betas,omitempty"`
	GithubID            int32                  `json:"github_id,omitempty"`
	GithubOauthScopes   []string               `json:"github_oauth_scopes,omitempty"`
	GravatarID          int32                  `json:"gravatar_id,omitempty"`
	HerokuAPIKey        string                 `json:"heroku_api_key,omitempty"`
	InBetaProgram       bool                   `json:"in_beta_program,omitempty"`
	Login               string                 `json:"login,omitempty"`
	Name                string                 `json:"name,omitempty"`
	OrganizationPrefs   map[string]interface{} `json:"organization_prefs,omitempty"`
	Parallelism         int32                  `json:"parallelism,omitempty"`
	Plan                string                 `json:"plan,omitempty"`
	Projects            map[string]interface{} `json:"projects,omitempty"`
	PusherID            string                 `json:"pusher_id,omitempty"`
	SelectedEmail       string                 `json:"selected_email,omitempty"`
	SignInCount         int32                  `json:"sign_in_count,omitempty"`
	TrialEnd            time.Time              `json:"trial_end,omitempty"`
}
