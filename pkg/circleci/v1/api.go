// Copyright 2020 The ccictl Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Code generated file by oapi-generator.

package v1

import (
	"bytes"
	"compress/gzip"
	"context"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"path"
	"strconv"
	"strings"

	json "github.com/goccy/go-json"
	htransport "google.golang.org/api/transport/http"
)

// Always reference these packages, just in case the auto-generated code below doesn't.
var (
	_ = bytes.NewBuffer
	_ = context.Canceled
	_ = json.NewDecoder
	_ = errors.New
	_ = fmt.Sprintf
	_ = io.Copy
	_ = ioutil.ReadAll
	_ = http.NewRequest
	_ = url.Parse
	_ = strconv.Itoa
	_ = path.Join
	_ = strings.Replace
	_ = gzip.NewReader
	_ = htransport.NewClient
)

// API represents a CircleCI v1.1 API.
type API struct {
	s *Service
}

// New returns the new API.
func New(s *Service) *API {
	rs := &API{s: s}
	return rs
}

// MeCall represents a provides information about the signed in user.
type MeCall struct {
	s      *Service
	params url.Values
}

// Me returns the information about the signed in user.
func (r *API) Me() *MeCall {
	c := &MeCall{
		s:      r.s,
		params: url.Values{},
	}
	return c
}

// Context sets context.
func (c *MeCall) Context(ctx context.Context) *MeCall {
	c.s.ctx = ctx
	return c
}

// Do executes the Me.
func (c *MeCall) Do() (*User, error) {
	uri := path.Join(c.s.BasePath, "/me")
	if len(c.params) > 0 {
		uri += "?" + c.params.Encode()
	}

	req, err := http.NewRequestWithContext(c.s.ctx, http.MethodGet, uri, nil)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Accept-Encoding", "application/json")

	resp, err := c.s.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		return nil, errors.New(resp.Status)
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var result User
	if err := json.Unmarshal(body, &result); err != nil {
		return nil, err
	}

	return &result, nil
}

// BuildCall represents a build summary for each of the last 30 builds for a single git repo.
type BuildCall struct {
	s      *Service
	params url.Values

	// path fields
	username string
	project  string

	// query fields
	limit  int
	offset int
	filter string
}

// Build returns the build summary for each of the last 30 builds for a single git repo.
func (r *API) Build(username, project string) *BuildCall {
	c := &BuildCall{
		s:        r.s,
		params:   url.Values{},
		username: username,
		project:  project,
	}
	return c
}

// Limit sets limit.
func (c *BuildCall) Limit(limit int) *BuildCall {
	c.params.Add("limit", fmt.Sprintf("%v", limit))
	c.limit = limit
	return c
}

// Offset sets offset.
func (c *BuildCall) Offset(offset int) *BuildCall {
	c.params.Add("offset", fmt.Sprintf("%v", offset))
	c.offset = offset
	return c
}

// Filter sets filter.
func (c *BuildCall) Filter(filter string) *BuildCall {
	c.params.Add("filter", fmt.Sprintf("%v", filter))
	c.filter = filter
	return c
}

// Context sets context.
func (c *BuildCall) Context(ctx context.Context) *BuildCall {
	c.s.ctx = ctx
	return c
}

// Do executes the Build.
func (c *BuildCall) Do() (*Build, error) {
	uri := path.Join(c.s.BasePath, "/project/"+fmt.Sprintf("%s/%s", c.username, c.project))
	if len(c.params) > 0 {
		uri += "?" + c.params.Encode()
	}

	req, err := http.NewRequestWithContext(c.s.ctx, http.MethodGet, uri, nil)
	if err != nil {
		return nil, err
	}

	req.Header.Set("User-Agent", c.s.userAgent())
	req.Header.Set("Accept-Encoding", "application/json")

	resp, err := c.s.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		return nil, errors.New(resp.Status)
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var result Build
	if err := json.Unmarshal(body, &result); err != nil {
		return nil, err
	}

	return &result, nil
}

// TriggerBuildAndSummaryCall represents a triggers a new build, and returns the summary of the build.
type TriggerBuildAndSummaryCall struct {
	s      *Service
	params url.Values

	// path fields
	username string
	project  string
	trigger  *TriggerBuildSummary
}

// TriggerBuildAndSummary triggers a new build, and returns the summary of the build.
func (r *API) TriggerBuildAndSummary(username, project string, trigger *TriggerBuildSummary) *TriggerBuildAndSummaryCall {
	c := &TriggerBuildAndSummaryCall{
		s:        r.s,
		params:   url.Values{},
		username: username,
		project:  project,
		trigger:  trigger,
	}
	return c
}

// Context sets context.
func (c *TriggerBuildAndSummaryCall) Context(ctx context.Context) *TriggerBuildAndSummaryCall {
	c.s.ctx = ctx
	return c
}

// Do executes the TriggerBuildAndSummary.
func (c *TriggerBuildAndSummaryCall) Do() (*BuildSummary, error) {
	uri := path.Join(c.s.BasePath, "/project/"+fmt.Sprintf("%s/%s", c.username, c.project))
	if len(c.params) > 0 {
		uri += "?" + c.params.Encode()
	}

	var r io.Reader
	if c.trigger != nil {
		reqBody, err := json.Marshal(&c.trigger)
		if err != nil {
			return nil, err
		}
		r = bytes.NewReader(reqBody)
	}

	req, err := http.NewRequestWithContext(c.s.ctx, http.MethodPost, uri, r)
	if err != nil {
		return nil, err
	}

	req.Header.Set("User-Agent", c.s.userAgent())
	req.Header.Set("Accept-Encoding", "application/json")

	resp, err := c.s.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		return nil, errors.New(resp.Status)
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var result BuildSummary
	if err := json.Unmarshal(body, &result); err != nil {
		return nil, err
	}

	return &result, nil
}

// DeleteBuildCacheCall represents a clears the cache for a project.
type DeleteBuildCacheCall struct {
	s      *Service
	params url.Values

	// path fields
	username string
	project  string
}

// DeleteBuildCache clears the cache for a project.
func (r *API) DeleteBuildCache(username, project string) *DeleteBuildCacheCall {
	c := &DeleteBuildCacheCall{
		s:        r.s,
		params:   url.Values{},
		username: username,
		project:  project,
	}
	return c
}

// Context sets context.
func (c *DeleteBuildCacheCall) Context(ctx context.Context) *DeleteBuildCacheCall {
	c.s.ctx = ctx
	return c
}

// Do executes the DeleteBuildCache.
func (c *DeleteBuildCacheCall) Do() (*Response, error) {
	uri := path.Join(c.s.BasePath, "/project/"+fmt.Sprintf("%s/%s", c.username, c.project), "build-cache")
	if len(c.params) > 0 {
		uri += "?" + c.params.Encode()
	}

	req, err := http.NewRequestWithContext(c.s.ctx, http.MethodDelete, uri, nil)
	if err != nil {
		return nil, err
	}

	req.Header.Set("User-Agent", c.s.userAgent())
	req.Header.Set("Accept-Encoding", "application/json")

	resp, err := c.s.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		return nil, errors.New(resp.Status)
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var result Response
	if err := json.Unmarshal(body, &result); err != nil {
		return nil, err
	}

	return &result, nil
}

// ListCheckoutKeyCall represents a lists checkout keys.
type ListCheckoutKeyCall struct {
	s      *Service
	params url.Values

	// path fields
	username string
	project  string
}

// ListCheckoutKey lists checkout keys.
func (r *API) ListCheckoutKey(username, project string) *ListCheckoutKeyCall {
	c := &ListCheckoutKeyCall{
		s:        r.s,
		params:   url.Values{},
		username: username,
		project:  project,
	}
	return c
}

// Context sets context.
func (c *ListCheckoutKeyCall) Context(ctx context.Context) *ListCheckoutKeyCall {
	c.s.ctx = ctx
	return c
}

// Do executes the ListCheckoutKey.
func (c *ListCheckoutKeyCall) Do() ([]*CheckoutKey, error) {
	uri := path.Join(c.s.BasePath, "/project/"+fmt.Sprintf("%s/%s", c.username, c.project), "checkout-key")
	if len(c.params) > 0 {
		uri += "?" + c.params.Encode()
	}

	req, err := http.NewRequestWithContext(c.s.ctx, http.MethodGet, uri, nil)
	if err != nil {
		return nil, err
	}

	req.Header.Set("User-Agent", c.s.userAgent())
	req.Header.Set("Accept-Encoding", "application/json")

	resp, err := c.s.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		return nil, errors.New(resp.Status)
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var result []*CheckoutKey
	if err := json.Unmarshal(body, &result); err != nil {
		return nil, err
	}

	return result, nil
}

// CreateCheckoutKeyCall represents a creates a new checkout key.
type CreateCheckoutKeyCall struct {
	s      *Service
	params url.Values

	// path fields
	username string
	project  string
}

// CreateCheckoutKey creates a new checkout key.
//
// Only usable with a user API token.
func (r *API) CreateCheckoutKey(username, project string) *CreateCheckoutKeyCall {
	c := &CreateCheckoutKeyCall{
		s:        r.s,
		params:   url.Values{},
		username: username,
		project:  project,
	}
	return c
}

// Context sets context.
func (c *CreateCheckoutKeyCall) Context(ctx context.Context) *CreateCheckoutKeyCall {
	c.s.ctx = ctx
	return c
}

// Do executes the CreateCheckoutKey.
func (c *CreateCheckoutKeyCall) Do() (*CheckoutKey, error) {
	uri := path.Join(c.s.BasePath, "/project/"+fmt.Sprintf("%s/%s", c.username, c.project), "checkout-key")
	if len(c.params) > 0 {
		uri += "?" + c.params.Encode()
	}

	req, err := http.NewRequestWithContext(c.s.ctx, http.MethodPost, uri, nil)
	if err != nil {
		return nil, err
	}

	req.Header.Set("User-Agent", c.s.userAgent())
	req.Header.Set("Accept-Encoding", "application/json")

	resp, err := c.s.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		return nil, errors.New(resp.Status)
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var result CheckoutKey
	if err := json.Unmarshal(body, &result); err != nil {
		return nil, err
	}

	return &result, nil
}

// DeleteCheckoutKeyCall represents a delete a checkout key.
type DeleteCheckoutKeyCall struct {
	s      *Service
	params url.Values

	// path fields
	username    string
	project     string
	fingerprint string
}

// DeleteCheckoutKey delete a checkout key.
func (r *API) DeleteCheckoutKey(username, project, fingerprint string) *DeleteCheckoutKeyCall {
	c := &DeleteCheckoutKeyCall{
		s:           r.s,
		params:      url.Values{},
		username:    username,
		project:     project,
		fingerprint: fingerprint,
	}
	return c
}

// Context sets context.
func (c *DeleteCheckoutKeyCall) Context(ctx context.Context) *DeleteCheckoutKeyCall {
	c.s.ctx = ctx
	return c
}

// Do executes the DeleteCheckoutKey.
func (c *DeleteCheckoutKeyCall) Do() (*Response, error) {
	uri := path.Join(c.s.BasePath, "/project/"+fmt.Sprintf("%s/%s/%s", c.username, c.project, c.fingerprint))
	if len(c.params) > 0 {
		uri += "?" + c.params.Encode()
	}

	req, err := http.NewRequestWithContext(c.s.ctx, http.MethodDelete, uri, nil)
	if err != nil {
		return nil, err
	}

	req.Header.Set("User-Agent", c.s.userAgent())
	req.Header.Set("Accept-Encoding", "application/json")

	resp, err := c.s.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		return nil, errors.New(resp.Status)
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var result Response
	if err := json.Unmarshal(body, &result); err != nil {
		return nil, err
	}

	return &result, nil
}

// CheckoutKeyCall represents a get a checkout key.
type CheckoutKeyCall struct {
	s      *Service
	params url.Values

	// path fields
	username    string
	project     string
	fingerprint string
}

// CheckoutKey get a checkout key.
func (r *API) CheckoutKey(username, project, fingerprint string) *CheckoutKeyCall {
	c := &CheckoutKeyCall{
		s:           r.s,
		params:      url.Values{},
		username:    username,
		project:     project,
		fingerprint: fingerprint,
	}
	return c
}

// Context sets context.
func (c *CheckoutKeyCall) Context(ctx context.Context) *CheckoutKeyCall {
	c.s.ctx = ctx
	return c
}

// Do executes the CheckoutKey.
func (c *CheckoutKeyCall) Do() (*CheckoutKey, error) {
	uri := path.Join(c.s.BasePath, "/project/"+fmt.Sprintf("%s/%s/checkout-key/%s", c.username, c.project, c.fingerprint))
	if len(c.params) > 0 {
		uri += "?" + c.params.Encode()
	}

	req, err := http.NewRequestWithContext(c.s.ctx, http.MethodGet, uri, nil)
	if err != nil {
		return nil, err
	}

	req.Header.Set("User-Agent", c.s.userAgent())
	req.Header.Set("Accept-Encoding", "application/json")

	resp, err := c.s.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		return nil, errors.New(resp.Status)
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var result CheckoutKey
	if err := json.Unmarshal(body, &result); err != nil {
		return nil, err
	}

	return &result, nil
}

// ListProjectEnvVarsCall represents a lists the environment variables for project.
type ListProjectEnvVarsCall struct {
	s      *Service
	params url.Values

	// path fields
	username string
	project  string
}

// ListProjectEnvVars lists the environment variables for project.
func (r *API) ListProjectEnvVars(username, project string) *ListProjectEnvVarsCall {
	c := &ListProjectEnvVarsCall{
		s:        r.s,
		params:   url.Values{},
		username: username,
		project:  project,
	}
	return c
}

// Context sets context.
func (c *ListProjectEnvVarsCall) Context(ctx context.Context) *ListProjectEnvVarsCall {
	c.s.ctx = ctx
	return c
}

// Do executes the ListProjectEnvVars.
func (c *ListProjectEnvVarsCall) Do() ([]Envvar, error) {
	uri := path.Join(c.s.BasePath, "project", fmt.Sprintf("%s/%s", c.username, c.project), "envvar")
	if len(c.params) > 0 {
		uri += "?" + c.params.Encode()
	}

	req, err := http.NewRequestWithContext(c.s.ctx, http.MethodGet, uri, nil)
	if err != nil {
		return nil, err
	}

	req.Header.Set("User-Agent", c.s.userAgent())
	req.Header.Set("Accept-Encoding", "application/json")

	resp, err := c.s.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		return nil, errors.New(resp.Status)
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var result []Envvar
	if err := json.Unmarshal(body, &result); err != nil {
		return nil, err
	}

	return result, nil
}

// CreateProjectEnvVarCall represents a creates a new environment variable.
type CreateProjectEnvVarCall struct {
	s      *Service
	params url.Values

	// path fields
	username string
	project  string
}

// CreateProjectEnvVar creates a new environment variable.
func (r *API) CreateProjectEnvVar(username, project string) *CreateProjectEnvVarCall {
	c := &CreateProjectEnvVarCall{
		s:        r.s,
		params:   url.Values{},
		username: username,
		project:  project,
	}
	return c
}

// Context sets context.
func (c *CreateProjectEnvVarCall) Context(ctx context.Context) *CreateProjectEnvVarCall {
	c.s.ctx = ctx
	return c
}

// Do executes the CreateProjectEnvVar.
func (c *CreateProjectEnvVarCall) Do() (*Envvar, error) {
	uri := path.Join(c.s.BasePath, "project", fmt.Sprintf("%s/%s", c.username, c.project), "envvar")
	if len(c.params) > 0 {
		uri += "?" + c.params.Encode()
	}

	req, err := http.NewRequestWithContext(c.s.ctx, http.MethodPost, uri, nil)
	if err != nil {
		return nil, err
	}

	req.Header.Set("User-Agent", c.s.userAgent())
	req.Header.Set("Accept-Encoding", "application/json")

	resp, err := c.s.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		return nil, errors.New(resp.Status)
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var result Envvar
	if err := json.Unmarshal(body, &result); err != nil {
		return nil, err
	}

	return &result, nil
}

// DeleteEnvVarCall represents a deletes the environment variable named.
type DeleteEnvVarCall struct {
	s      *Service
	params url.Values

	// path fields
	username string
	project  string
	name     string
}

// DeleteEnvVar deletes the environment variable named.
func (r *API) DeleteEnvVar(username, project, name string) *DeleteEnvVarCall {
	c := &DeleteEnvVarCall{
		s:        r.s,
		params:   url.Values{},
		username: username,
		project:  project,
		name:     name,
	}
	return c
}

// Context sets context.
func (c *DeleteEnvVarCall) Context(ctx context.Context) *DeleteEnvVarCall {
	c.s.ctx = ctx
	return c
}

// Do executes the DeleteEnvVar.
func (c *DeleteEnvVarCall) Do() (*Response, error) {
	uri := path.Join(c.s.BasePath, "project", fmt.Sprintf("%s/%s/envvar/%s", c.username, c.project, c.name))
	if len(c.params) > 0 {
		uri += "?" + c.params.Encode()
	}

	req, err := http.NewRequestWithContext(c.s.ctx, http.MethodDelete, uri, nil)
	if err != nil {
		return nil, err
	}

	req.Header.Set("User-Agent", c.s.userAgent())
	req.Header.Set("Accept-Encoding", "application/json")

	resp, err := c.s.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		return nil, errors.New(resp.Status)
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var result Response
	if err := json.Unmarshal(body, &result); err != nil {
		return nil, err
	}

	return &result, nil
}

// HiddenEnvVarCall represents a gets the hidden value of environment variable.
type HiddenEnvVarCall struct {
	s      *Service
	params url.Values

	// path fields
	username string
	project  string
	name     string
}

// HiddenEnvVar gets the hidden value of environment variable.
func (r *API) HiddenEnvVar(username, project, name string) *HiddenEnvVarCall {
	c := &HiddenEnvVarCall{
		s:        r.s,
		params:   url.Values{},
		username: username,
		project:  project,
		name:     name,
	}
	return c
}

// Context sets context.
func (c *HiddenEnvVarCall) Context(ctx context.Context) *HiddenEnvVarCall {
	c.s.ctx = ctx
	return c
}

// Do executes the HiddenEnvVar.
func (c *HiddenEnvVarCall) Do() (*Envvar, error) {
	uri := path.Join(c.s.BasePath, "project", fmt.Sprintf("%s/%s/%s", c.username, c.project, c.name))

	if len(c.params) > 0 {
		uri += "?" + c.params.Encode()
	}

	req, err := http.NewRequestWithContext(c.s.ctx, http.MethodGet, uri, nil)
	if err != nil {
		return nil, err
	}

	req.Header.Set("User-Agent", c.s.userAgent())
	req.Header.Set("Accept-Encoding", "application/json")

	resp, err := c.s.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		return nil, errors.New(resp.Status)
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var result Envvar
	if err := json.Unmarshal(body, &result); err != nil {
		return nil, err
	}

	return &result, nil
}

// CreateSSHKeyCall represents a create an ssh key used to access external systems that require SSH key-based authentication.
type CreateSSHKeyCall struct {
	s      *Service
	params url.Values

	// path fields
	username string
	project  string
	sshKey   SSHKey
}

// CreateSSHKey create an ssh key used to access external systems that require SSH key-based authentication.
func (r *API) CreateSSHKey(username, project string, sshKey SSHKey) *CreateSSHKeyCall {
	c := &CreateSSHKeyCall{
		s:        r.s,
		params:   url.Values{},
		username: username,
		project:  project,
		sshKey:   sshKey,
	}
	return c
}

// Context sets context.
func (c *CreateSSHKeyCall) Context(ctx context.Context) *CreateSSHKeyCall {
	c.s.ctx = ctx
	return c
}

// Do executes the CreateSSHKey.
func (c *CreateSSHKeyCall) Do() (*Response, error) {
	uri := path.Join(c.s.BasePath, "project", fmt.Sprintf("%s/%s", c.username, c.project), "ssh-key")
	if len(c.params) > 0 {
		uri += "?" + c.params.Encode()
	}

	bodySSHkey, err := json.Marshal(&c.sshKey)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequestWithContext(c.s.ctx, http.MethodPost, uri, bytes.NewReader(bodySSHkey))
	if err != nil {
		return nil, err
	}

	req.Header.Set("User-Agent", c.s.userAgent())
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Accept-Encoding", "application/json")

	resp, err := c.s.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		return nil, errors.New(resp.Status)
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var result Response
	if err := json.Unmarshal(body, &result); err != nil {
		return nil, err
	}

	return &result, nil
}

// TriggerNewBuildCall represents a triggers a new build, returns a summary of the build.
type TriggerNewBuildCall struct {
	s      *Service
	params url.Values

	// path fields
	username string
	project  string
	branch   string

	// body fields
	trigger *TriggerBuild
}

// TriggerNewBuild triggers a new build, returns a summary of the build.
func (r *API) TriggerNewBuild(username, project, branch string) *TriggerNewBuildCall {
	c := &TriggerNewBuildCall{
		s:        r.s,
		params:   url.Values{},
		username: username,
		project:  project,
		branch:   branch,
	}
	return c
}

// Context sets context.
func (c *TriggerNewBuildCall) Context(ctx context.Context) *TriggerNewBuildCall {
	c.s.ctx = ctx
	return c
}

// Trigger sets trigger.
//
// The build parameters can be set using an experimental API.
func (c *TriggerNewBuildCall) Trigger(trigger *TriggerBuild) *TriggerNewBuildCall {
	c.params.Add("body", fmt.Sprintf("%v", trigger))
	c.trigger = trigger
	return c
}

// Do executes the TriggerNewBuild.
func (c *TriggerNewBuildCall) Do() (*Build, error) {
	uri := path.Join(c.s.BasePath, "project", fmt.Sprintf("%s/%s/tree/%s", c.username, c.project, c.branch))
	if len(c.params) > 0 {
		uri += "?" + c.params.Encode()
	}

	var r io.Reader
	if c.trigger != nil {
		reqBody, err := json.Marshal(&c.trigger)
		if err != nil {
			return nil, err
		}
		r = bytes.NewReader(reqBody)
	}

	req, err := http.NewRequestWithContext(c.s.ctx, http.MethodPost, uri, r)
	if err != nil {
		return nil, err
	}

	req.Header.Set("User-Agent", c.s.userAgent())
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Accept-Encoding", "application/json")

	resp, err := c.s.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		return nil, errors.New(resp.Status)
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var result Build
	if err := json.Unmarshal(body, &result); err != nil {
		return nil, err
	}

	return &result, nil
}

// BuildDetailCall represents a full details for a single build. The response includes all of the fields from the build summary.
type BuildDetailCall struct {
	s      *Service
	params url.Values

	// path fields
	username string
	project  string
	buildNum int

	// query fields
	limit int
}

// BuildDetail full details for a single build. The response includes all of the fields from the build summary.
//
// This is also the payload for the notification webhooks, in which case this object is the value to a key named 'payload'.
func (r *API) BuildDetail(username, project string, buildNum int) *BuildDetailCall {
	c := &BuildDetailCall{
		s:        r.s,
		params:   url.Values{},
		username: username,
		project:  project,
		buildNum: buildNum,
	}
	return c
}

// Context sets context.
func (c *BuildDetailCall) Context(ctx context.Context) *BuildDetailCall {
	c.s.ctx = ctx
	return c
}

// Do executes the BuildDetail.
func (c *BuildDetailCall) Do() (*BuildDetail, error) {
	uri := path.Join(c.s.BasePath, "project", fmt.Sprintf("%s/%s/%d", c.username, c.project, c.buildNum))
	if len(c.params) > 0 {
		uri += "?" + c.params.Encode()
	}

	req, err := http.NewRequestWithContext(c.s.ctx, http.MethodGet, uri, nil)
	if err != nil {
		return nil, err
	}

	req.Header.Set("User-Agent", c.s.userAgent())
	req.Header.Set("Accept-Encoding", "application/json")

	resp, err := c.s.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		return nil, errors.New(resp.Status)
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var result BuildDetail
	if err := json.Unmarshal(body, &result); err != nil {
		return nil, err
	}

	return &result, nil
}

// ListArtifactsCall represents a list the artifacts produced by a given build.
type ListArtifactsCall struct {
	s      *Service
	params url.Values

	// path fields
	username string
	project  string
	buildNum int
}

// ListArtifacts list the artifacts produced by a given build.
func (r *API) ListArtifacts(username, project string, buildNum int) *ListArtifactsCall {
	c := &ListArtifactsCall{
		s:        r.s,
		params:   url.Values{},
		username: username,
		project:  project,
		buildNum: buildNum,
	}
	return c
}

// Context sets context.
func (c *ListArtifactsCall) Context(ctx context.Context) *ListArtifactsCall {
	c.s.ctx = ctx
	return c
}

// Do executes the ListArtifacts.
func (c *ListArtifactsCall) Do() ([]Artifact, error) {
	uri := path.Join(c.s.BasePath, "project", fmt.Sprintf("%s/%s/%d", c.username, c.project, c.buildNum), "artifacts")
	if len(c.params) > 0 {
		uri += "?" + c.params.Encode()
	}

	req, err := http.NewRequestWithContext(c.s.ctx, http.MethodGet, uri, nil)
	if err != nil {
		return nil, err
	}

	req.Header.Set("User-Agent", c.s.userAgent())
	req.Header.Set("Accept-Encoding", "application/json")

	resp, err := c.s.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		return nil, errors.New(resp.Status)
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var result []Artifact
	if err := json.Unmarshal(body, &result); err != nil {
		return nil, err
	}

	return result, nil
}

// CancelBuildCall represents a cancels the build, returns a summary of the build.
type CancelBuildCall struct {
	s      *Service
	params url.Values

	// path fields
	username string
	project  string
	buildNum int
}

// CancelBuild cancels the build, returns a summary of the build.
func (r *API) CancelBuild(username, project string, buildNum int) *CancelBuildCall {
	c := &CancelBuildCall{
		s:        r.s,
		params:   url.Values{},
		username: username,
		project:  project,
		buildNum: buildNum,
	}
	return c
}

// Context sets context.
func (c *CancelBuildCall) Context(ctx context.Context) *CancelBuildCall {
	c.s.ctx = ctx
	return c
}

// Do executes the CancelBuild.
func (c *CancelBuildCall) Do() (*Build, error) {
	uri := path.Join(c.s.BasePath, "project", fmt.Sprintf("%s/%s/%d", c.username, c.project, c.buildNum), "cancel")
	if len(c.params) > 0 {
		uri += "?" + c.params.Encode()
	}

	req, err := http.NewRequestWithContext(c.s.ctx, http.MethodPost, uri, nil)
	if err != nil {
		return nil, err
	}

	req.Header.Set("User-Agent", c.s.userAgent())
	req.Header.Set("Accept-Encoding", "application/json")

	resp, err := c.s.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		return nil, errors.New(resp.Status)
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var result Build
	if err := json.Unmarshal(body, &result); err != nil {
		return nil, err
	}

	return &result, nil
}

// RetryBuildCall represents a retries the build, returns a summary of the new build.
type RetryBuildCall struct {
	s      *Service
	params url.Values

	// path fields
	username string
	project  string
	buildNum int
}

// RetryBuild retries the build, returns a summary of the new build.
func (r *API) RetryBuild(username, project string, buildNum int) *RetryBuildCall {
	c := &RetryBuildCall{
		s:        r.s,
		params:   url.Values{},
		username: username,
		project:  project,
		buildNum: buildNum,
	}
	return c
}

// Context sets context.
func (c *RetryBuildCall) Context(ctx context.Context) *RetryBuildCall {
	c.s.ctx = ctx
	return c
}

// Do executes the RetryBuild.
func (c *RetryBuildCall) Do() (*Build, error) {
	uri := path.Join(c.s.BasePath, "project", fmt.Sprintf("%s/%s/%d", c.username, c.project, c.buildNum), "retry")
	if len(c.params) > 0 {
		uri += "?" + c.params.Encode()
	}

	req, err := http.NewRequestWithContext(c.s.ctx, http.MethodPost, uri, nil)
	if err != nil {
		return nil, err
	}

	req.Header.Set("User-Agent", c.s.userAgent())
	req.Header.Set("Accept-Encoding", "application/json")

	resp, err := c.s.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		return nil, errors.New(resp.Status)
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var result Build
	if err := json.Unmarshal(body, &result); err != nil {
		return nil, err
	}

	return &result, nil
}

// TestMetadataCall represents a provides test metadata for a build.
type TestMetadataCall struct {
	s      *Service
	params url.Values

	// path fields
	username string
	project  string
	buildNum int
}

// TestMetadata provides test metadata for a build.
func (r *API) TestMetadata(username, project string, buildNum int) *TestMetadataCall {
	c := &TestMetadataCall{
		s:        r.s,
		params:   url.Values{},
		username: username,
		project:  project,
		buildNum: buildNum,
	}
	return c
}

// Context sets context.
func (c *TestMetadataCall) Context(ctx context.Context) *TestMetadataCall {
	c.s.ctx = ctx
	return c
}

// Do executes the TestMetadata.
func (c *TestMetadataCall) Do() (*Tests, error) {
	uri := path.Join(c.s.BasePath, "project", fmt.Sprintf("%s/%s/%d", c.username, c.project, c.buildNum), "tests")
	if len(c.params) > 0 {
		uri += "?" + c.params.Encode()
	}

	req, err := http.NewRequestWithContext(c.s.ctx, http.MethodGet, uri, nil)
	if err != nil {
		return nil, err
	}

	req.Header.Set("User-Agent", c.s.userAgent())
	req.Header.Set("Accept-Encoding", "application/json")

	resp, err := c.s.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		return nil, errors.New(resp.Status)
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var result Tests
	if err := json.Unmarshal(body, &result); err != nil {
		return nil, err
	}

	return &result, nil
}

// ListProjectCall represents a list of all the projects you're following on CircleCI, with build information organized by branch.
type ListProjectCall struct {
	s      *Service
	params url.Values
}

// ListProject list of all the projects you're following on CircleCI, with build information organized by branch.
func (r *API) ListProject() *ListProjectCall {
	c := &ListProjectCall{
		s:      r.s,
		params: url.Values{},
	}
	return c
}

// Context sets context.
func (c *ListProjectCall) Context(ctx context.Context) *ListProjectCall {
	c.s.ctx = ctx
	return c
}

// Do executes the ListProject.
func (c *ListProjectCall) Do() ([]*Project, error) {
	uri := path.Join(c.s.BasePath, "projects")
	if len(c.params) > 0 {
		uri += "?" + c.params.Encode()
	}

	req, err := http.NewRequestWithContext(c.s.ctx, http.MethodGet, uri, nil)
	if err != nil {
		return nil, err
	}

	req.Header.Set("User-Agent", c.s.userAgent())
	req.Header.Set("Accept-Encoding", "application/json")

	resp, err := c.s.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		return nil, errors.New(resp.Status)
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var result []*Project
	if err := json.Unmarshal(body, &result); err != nil {
		return nil, err
	}

	return result, nil
}

// RecentBuildsCall represents a build summary for each of the last 30 recent builds, ordered by build_num.
type RecentBuildsCall struct {
	s      *Service
	params url.Values

	// query fields
	limit  int
	offset int
}

// RecentBuilds build summary for each of the last 30 recent builds, ordered by build_num.
func (r *API) RecentBuilds() *RecentBuildsCall {
	c := &RecentBuildsCall{
		s:      r.s,
		params: url.Values{},
		limit:  30, // limit is 30 by default
	}
	return c
}

// Limit sets page limit.
func (c *RecentBuildsCall) Limit(limit int) *RecentBuildsCall {
	c.params.Add("limit", fmt.Sprintf("%v", limit))
	c.limit = limit
	return c
}

// Offset sets offset.
func (c *RecentBuildsCall) Offset(offset int) *RecentBuildsCall {
	c.params.Add("offset", fmt.Sprintf("%v", offset))
	c.offset = offset
	return c
}

// Context sets context.
func (c *RecentBuildsCall) Context(ctx context.Context) *RecentBuildsCall {
	c.s.ctx = ctx
	return c
}

// Do executes the RecentBuilds.
func (c *RecentBuildsCall) Do() ([]*Build, error) {
	uri := path.Join(c.s.BasePath, "recent-builds")
	if len(c.params) > 0 {
		uri += "?" + c.params.Encode()
	}

	req, err := http.NewRequestWithContext(c.s.ctx, http.MethodGet, uri, nil)
	if err != nil {
		return nil, err
	}

	req.Header.Set("User-Agent", c.s.userAgent())
	req.Header.Set("Accept-Encoding", "application/json")

	resp, err := c.s.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		return nil, errors.New(resp.Status)
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var result []*Build
	if err := json.Unmarshal(body, &result); err != nil {
		return nil, err
	}

	return result, nil
}

// AddHerokuKeyCall represents a adds your Heroku API key to CircleCI, takes apikey as form param name.
type AddHerokuKeyCall struct {
	s      *Service
	params url.Values
}

// AddHerokuKey adds your Heroku API key to CircleCI, takes apikey as form param name.
func (r *API) AddHerokuKey() *AddHerokuKeyCall {
	c := &AddHerokuKeyCall{
		s:      r.s,
		params: url.Values{},
	}
	return c
}

// Context sets context.
func (c *AddHerokuKeyCall) Context(ctx context.Context) *AddHerokuKeyCall {
	c.s.ctx = ctx
	return c
}

// Do executes the AddHerokuKey.
func (c *AddHerokuKeyCall) Do() (*Response, error) {
	uri := path.Join(c.s.BasePath, "user", "heroku-key")
	if len(c.params) > 0 {
		uri += "?" + c.params.Encode()
	}

	req, err := http.NewRequestWithContext(c.s.ctx, http.MethodPost, uri, nil)
	if err != nil {
		return nil, err
	}

	req.Header.Set("User-Agent", c.s.userAgent())
	req.Header.Set("Accept-Encoding", "application/json")

	resp, err := c.s.client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		return nil, errors.New(resp.Status)
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var result Response
	if err := json.Unmarshal(body, &result); err != nil {
		return nil, err
	}

	return &result, nil
}
