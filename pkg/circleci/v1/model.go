// Copyright 2020 The ccictl Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package v1

type TriggerBuildSummary struct {
	BuildParameters interface{} `json:"build_parameters,omitempty"`
	Parallel        string      `json:"parallel,omitempty"`
	Revision        string      `json:"revision,omitempty"`
	Tag             string      `json:"tag,omitempty"`
}

type SSHKey struct {
	HostName   string `json:"hostname"`
	PrivateKey string `json:"private_key"`
}

type TriggerBuild struct {
	BuildParameters interface{} `json:"build_parameters,omitempty"`
	Parallel        string      `json:"parallel,omitempty"`
	Revision        string      `json:"revision,omitempty"`
}
