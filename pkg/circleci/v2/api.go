// Copyright 2020 The ccictl Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Code generated file by oapi-generator.

package v2

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

// API represents a CircleCI v2 API.
type API struct {
	s *Service
}

// New returns the new API.
func New(s *Service) *API {
	rs := &API{s: s}
	return rs
}

// PipelineByIDCall provides the get a pipeline.
type PipelineByIDCall struct {
	s      *Service
	header http.Header
	params url.Values

	// path fields
	pipelineID string
}

// PipelineByID returns the PipelineByIDCall for get a pipeline.
func (r *API) PipelineByID(pipelineID string) *PipelineByIDCall {
	c := &PipelineByIDCall{
		s:          r.s,
		header:     make(http.Header),
		params:     url.Values{},
		pipelineID: pipelineID,
	}
	return c
}

// Do executes the PipelineByID.
func (c *PipelineByIDCall) Do(ctx context.Context) (interface{}, error) {
	uri := path.Join(c.s.BasePath, "/pipeline/"+fmt.Sprintf("%v", c.pipelineID)+"")
	if len(c.params) > 0 {
		uri += "?" + c.params.Encode()
	}

	req, err := http.NewRequest(http.MethodGet, uri, nil)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Accept-Encoding", "application/json")

	resp, err := c.s.client.Do(req.WithContext(ctx))
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

	var result interface{}
	if err := json.Unmarshal(body, &result); err != nil {
		return nil, err
	}

	return result, nil
}

// PipelineConfigByIDCall provides the get a pipeline's configuration.
type PipelineConfigByIDCall struct {
	s      *Service
	header http.Header
	params url.Values

	// path fields
	pipelineID string
}

// PipelineConfigByID returns the PipelineConfigByIDCall for get a pipeline's configuration.
func (r *API) PipelineConfigByID(pipelineID string) *PipelineConfigByIDCall {
	c := &PipelineConfigByIDCall{
		s:          r.s,
		header:     make(http.Header),
		params:     url.Values{},
		pipelineID: pipelineID,
	}
	return c
}

// Do executes the PipelineConfigByID.
func (c *PipelineConfigByIDCall) Do(ctx context.Context) (interface{}, error) {
	uri := path.Join(c.s.BasePath, "/pipeline/"+fmt.Sprintf("%v", c.pipelineID)+"/config")
	if len(c.params) > 0 {
		uri += "?" + c.params.Encode()
	}

	req, err := http.NewRequest(http.MethodGet, uri, nil)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Accept-Encoding", "application/json")

	resp, err := c.s.client.Do(req.WithContext(ctx))
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

	var result interface{}
	if err := json.Unmarshal(body, &result); err != nil {
		return nil, err
	}

	return result, nil
}

// ListWorkflowsByPipelineIDCall provides the get a pipeline's workflows.
type ListWorkflowsByPipelineIDCall struct {
	s      *Service
	header http.Header
	params url.Values

	// path fields
	pipelineID string
	// query fields
	pageToken string
}

// ListWorkflowsByPipelineID returns the ListWorkflowsByPipelineIDCall for get a pipeline's workflows.
func (r *API) ListWorkflowsByPipelineID(pipelineID string) *ListWorkflowsByPipelineIDCall {
	c := &ListWorkflowsByPipelineIDCall{
		s:          r.s,
		header:     make(http.Header),
		params:     url.Values{},
		pipelineID: pipelineID,
	}
	return c
}

func (c *ListWorkflowsByPipelineIDCall) PageToken(pageToken string) *ListWorkflowsByPipelineIDCall {
	c.params.Set("pageToken", fmt.Sprintf("%v", pageToken))
	return c
}

// Do executes the ListWorkflowsByPipelineID.
func (c *ListWorkflowsByPipelineIDCall) Do(ctx context.Context) (interface{}, error) {
	uri := path.Join(c.s.BasePath, "/pipeline/"+fmt.Sprintf("%v", c.pipelineID)+"/workflow")
	if len(c.params) > 0 {
		uri += "?" + c.params.Encode()
	}

	req, err := http.NewRequest(http.MethodGet, uri, nil)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Accept-Encoding", "application/json")

	resp, err := c.s.client.Do(req.WithContext(ctx))
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

	var result interface{}
	if err := json.Unmarshal(body, &result); err != nil {
		return nil, err
	}

	return result, nil
}

// ListPipelinesForProjectCall provides the get all pipelines.
type ListPipelinesForProjectCall struct {
	s      *Service
	header http.Header
	params url.Values

	// path fields
	projectSlug string
	// query fields
	branch    string
	pageToken string
}

// ListPipelinesForProject returns the ListPipelinesForProjectCall for get all pipelines.
func (r *API) ListPipelinesForProject(projectSlug string) *ListPipelinesForProjectCall {
	c := &ListPipelinesForProjectCall{
		s:           r.s,
		header:      make(http.Header),
		params:      url.Values{},
		projectSlug: projectSlug,
	}
	return c
}

func (c *ListPipelinesForProjectCall) Branch(branch string) *ListPipelinesForProjectCall {
	c.params.Set("branch", fmt.Sprintf("%v", branch))
	return c
}

func (c *ListPipelinesForProjectCall) PageToken(pageToken string) *ListPipelinesForProjectCall {
	c.params.Set("pageToken", fmt.Sprintf("%v", pageToken))
	return c
}

// Do executes the ListPipelinesForProject.
func (c *ListPipelinesForProjectCall) Do(ctx context.Context) (interface{}, error) {
	uri := path.Join(c.s.BasePath, "/project/"+fmt.Sprintf("%v", c.projectSlug)+"/pipeline")
	if len(c.params) > 0 {
		uri += "?" + c.params.Encode()
	}

	req, err := http.NewRequest(http.MethodGet, uri, nil)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Accept-Encoding", "application/json")

	resp, err := c.s.client.Do(req.WithContext(ctx))
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

	var result interface{}
	if err := json.Unmarshal(body, &result); err != nil {
		return nil, err
	}

	return result, nil
}

// ListMyPipelinesCall provides the get your pipelines.
type ListMyPipelinesCall struct {
	s      *Service
	header http.Header
	params url.Values

	// path fields
	projectSlug string
	// query fields
	pageToken string
}

// ListMyPipelines returns the ListMyPipelinesCall for get your pipelines.
func (r *API) ListMyPipelines(projectSlug string) *ListMyPipelinesCall {
	c := &ListMyPipelinesCall{
		s:           r.s,
		header:      make(http.Header),
		params:      url.Values{},
		projectSlug: projectSlug,
	}
	return c
}

func (c *ListMyPipelinesCall) PageToken(pageToken string) *ListMyPipelinesCall {
	c.params.Set("pageToken", fmt.Sprintf("%v", pageToken))
	return c
}

// Do executes the ListMyPipelines.
func (c *ListMyPipelinesCall) Do(ctx context.Context) (interface{}, error) {
	uri := path.Join(c.s.BasePath, "/project/"+fmt.Sprintf("%v", c.projectSlug)+"/pipeline/mine")
	if len(c.params) > 0 {
		uri += "?" + c.params.Encode()
	}

	req, err := http.NewRequest(http.MethodGet, uri, nil)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Accept-Encoding", "application/json")

	resp, err := c.s.client.Do(req.WithContext(ctx))
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

	var result interface{}
	if err := json.Unmarshal(body, &result); err != nil {
		return nil, err
	}

	return result, nil
}

// PipelineByNumberCall provides the get a pipeline.
type PipelineByNumberCall struct {
	s      *Service
	header http.Header
	params url.Values

	// path fields
	projectSlug    string
	pipelineNumber int64
}

// PipelineByNumber returns the PipelineByNumberCall for get a pipeline.
func (r *API) PipelineByNumber(projectSlug string, pipelineNumber int64) *PipelineByNumberCall {
	c := &PipelineByNumberCall{
		s:              r.s,
		header:         make(http.Header),
		params:         url.Values{},
		projectSlug:    projectSlug,
		pipelineNumber: pipelineNumber,
	}
	return c
}

// Do executes the PipelineByNumber.
func (c *PipelineByNumberCall) Do(ctx context.Context) (interface{}, error) {
	uri := path.Join(c.s.BasePath, "/project/"+fmt.Sprintf("%v", c.projectSlug)+"/pipeline/"+fmt.Sprintf("%v", c.pipelineNumber)+"")
	if len(c.params) > 0 {
		uri += "?" + c.params.Encode()
	}

	req, err := http.NewRequest(http.MethodGet, uri, nil)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Accept-Encoding", "application/json")

	resp, err := c.s.client.Do(req.WithContext(ctx))
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

	var result interface{}
	if err := json.Unmarshal(body, &result); err != nil {
		return nil, err
	}

	return result, nil
}

// WorkflowByIDCall provides the get a workflow.
type WorkflowByIDCall struct {
	s      *Service
	header http.Header
	params url.Values

	// path fields
	id string
}

// WorkflowByID returns the WorkflowByIDCall for get a workflow.
func (r *API) WorkflowByID(id string) *WorkflowByIDCall {
	c := &WorkflowByIDCall{
		s:      r.s,
		header: make(http.Header),
		params: url.Values{},
		id:     id,
	}
	return c
}

// Do executes the WorkflowByID.
func (c *WorkflowByIDCall) Do(ctx context.Context) (interface{}, error) {
	uri := path.Join(c.s.BasePath, "/workflow/"+fmt.Sprintf("%v", c.id)+"")
	if len(c.params) > 0 {
		uri += "?" + c.params.Encode()
	}

	req, err := http.NewRequest(http.MethodGet, uri, nil)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Accept-Encoding", "application/json")

	resp, err := c.s.client.Do(req.WithContext(ctx))
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

	var result interface{}
	if err := json.Unmarshal(body, &result); err != nil {
		return nil, err
	}

	return result, nil
}

// ApprovePendingApprovalJobByIDCall provides the approve a job.
type ApprovePendingApprovalJobByIDCall struct {
	s      *Service
	header http.Header
	params url.Values

	// path fields
	id                string
	approvalRequestID string
}

// ApprovePendingApprovalJobByID returns the ApprovePendingApprovalJobByIDCall for approve a job.
func (r *API) ApprovePendingApprovalJobByID(id string, approvalRequestID string) *ApprovePendingApprovalJobByIDCall {
	c := &ApprovePendingApprovalJobByIDCall{
		s:                 r.s,
		header:            make(http.Header),
		params:            url.Values{},
		id:                id,
		approvalRequestID: approvalRequestID,
	}
	return c
}

// Do executes the ApprovePendingApprovalJobByID.
func (c *ApprovePendingApprovalJobByIDCall) Do(ctx context.Context) (interface{}, error) {
	uri := path.Join(c.s.BasePath, "/workflow/"+fmt.Sprintf("%v", c.id)+"/approve/"+fmt.Sprintf("%v", c.approvalRequestID)+"")
	if len(c.params) > 0 {
		uri += "?" + c.params.Encode()
	}

	req, err := http.NewRequest(http.MethodPost, uri, nil)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Accept-Encoding", "application/json")

	resp, err := c.s.client.Do(req.WithContext(ctx))
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

	var result interface{}
	if err := json.Unmarshal(body, &result); err != nil {
		return nil, err
	}

	return result, nil
}

// CancelWorkflowCall provides the cancel a workflow.
type CancelWorkflowCall struct {
	s      *Service
	header http.Header
	params url.Values

	// path fields
	id string
}

// CancelWorkflow returns the CancelWorkflowCall for cancel a workflow.
func (r *API) CancelWorkflow(id string) *CancelWorkflowCall {
	c := &CancelWorkflowCall{
		s:      r.s,
		header: make(http.Header),
		params: url.Values{},
		id:     id,
	}
	return c
}

// Do executes the CancelWorkflow.
func (c *CancelWorkflowCall) Do(ctx context.Context) (interface{}, error) {
	uri := path.Join(c.s.BasePath, "/workflow/"+fmt.Sprintf("%v", c.id)+"/cancel")
	if len(c.params) > 0 {
		uri += "?" + c.params.Encode()
	}

	req, err := http.NewRequest(http.MethodPost, uri, nil)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Accept-Encoding", "application/json")

	resp, err := c.s.client.Do(req.WithContext(ctx))
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

	var result interface{}
	if err := json.Unmarshal(body, &result); err != nil {
		return nil, err
	}

	return result, nil
}

// ListWorkflowJobsCall provides the get a workflow's jobs.
type ListWorkflowJobsCall struct {
	s      *Service
	header http.Header
	params url.Values

	// path fields
	id string
}

// ListWorkflowJobs returns the ListWorkflowJobsCall for get a workflow's jobs.
func (r *API) ListWorkflowJobs(id string) *ListWorkflowJobsCall {
	c := &ListWorkflowJobsCall{
		s:      r.s,
		header: make(http.Header),
		params: url.Values{},
		id:     id,
	}
	return c
}

// Do executes the ListWorkflowJobs.
func (c *ListWorkflowJobsCall) Do(ctx context.Context) (interface{}, error) {
	uri := path.Join(c.s.BasePath, "/workflow/"+fmt.Sprintf("%v", c.id)+"/job")
	if len(c.params) > 0 {
		uri += "?" + c.params.Encode()
	}

	req, err := http.NewRequest(http.MethodGet, uri, nil)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Accept-Encoding", "application/json")

	resp, err := c.s.client.Do(req.WithContext(ctx))
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

	var result interface{}
	if err := json.Unmarshal(body, &result); err != nil {
		return nil, err
	}

	return result, nil
}

// ReRunWorkflowCall provides the rerun a workflow.
type ReRunWorkflowCall struct {
	s      *Service
	header http.Header
	params url.Values

	// path fields
	id string
}

// ReRunWorkflow returns the RerunWorkflowCall for rerun a workflow.
func (r *API) ReRunWorkflow(id string) *ReRunWorkflowCall {
	c := &ReRunWorkflowCall{
		s:      r.s,
		header: make(http.Header),
		params: url.Values{},
		id:     id,
	}
	return c
}

// Do executes the RerunWorkflow.
func (c *ReRunWorkflowCall) Do(ctx context.Context) (interface{}, error) {
	uri := path.Join(c.s.BasePath, "/workflow/"+fmt.Sprintf("%v", c.id)+"/rerun")
	if len(c.params) > 0 {
		uri += "?" + c.params.Encode()
	}

	req, err := http.NewRequest(http.MethodPost, uri, nil)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Accept-Encoding", "application/json")

	resp, err := c.s.client.Do(req.WithContext(ctx))
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

	var result interface{}
	if err := json.Unmarshal(body, &result); err != nil {
		return nil, err
	}

	return result, nil
}
