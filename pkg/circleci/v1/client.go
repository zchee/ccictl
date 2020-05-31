// Copyright 2020 The ccictl Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Code generated file by oapi-generator.

package v1

import (
	"bytes"
	"compress/gzip"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
	"path"
	"strconv"
	"strings"
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
)

const (
	// APIVersion CircleCI REST API version.
	APIVersion = "v1.1"

	// UserAgent is the default of user-agent.
	UserAgent = "ccictl/" + APIVersion
)

const (
	basePath = "https://circleci.com/api/v1.1"
)

// Service represents a V1 Services.
type Service struct {
	Client    *http.Client
	ctx       context.Context
	BasePath  string // API endpoint base URL
	UserAgent string // optional additional User-Agent fragment

	API *API
}

// NewService creates a new V1 Service.
func NewService(ctx context.Context, client *http.Client) *Service {
	if client == nil {
		client = &http.Client{}
	}
	uri, _ := url.Parse(basePath)
	svc := &Service{
		Client:   client,
		ctx:      ctx,
		BasePath: uri.String(),
	}
	svc.API = NewAPI(svc)

	return svc
}

func (s *Service) Token() string {
	return os.Getenv("CIRCLECI_TOKEN")
}

func (s *Service) userAgent() string {
	if s.UserAgent == "" {
		return UserAgent
	}
	return UserAgent + " " + s.UserAgent
}

// SchemaDescriptor returns the Schema file descriptor which is generated code to this file.
func SchemaDescriptor() (interface{}, error) {
	zr, err := gzip.NewReader(bytes.NewReader(fileDescriptor))
	if err != nil {
		return nil, err
	}

	var buf bytes.Buffer
	_, err = buf.ReadFrom(zr)
	if err != nil {
		return nil, err
	}

	var v interface{}
	if err := json.NewDecoder(bytes.NewReader(buf.Bytes())).Decode(&v); err != nil {
		return nil, err
	}

	return v, nil
}

// fileDescriptor gzipped JSON marshaled Schema object.
var fileDescriptor = []byte{
	// 4051 bytes of a gzipped Schema file descriptor
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x09, 0x6e, 0x88, 0x02, 0xff, 0xdc, 0x5c, 0xfd, 0x8e, 0x1b, 0x37,
	0x92, 0x7f, 0x15, 0xa2, 0xf7, 0x00, 0xef, 0x02, 0x92, 0x7a, 0x1c, 0x07, 0xb8, 0xc3, 0xfc, 0xe7,
	0xb5, 0x77, 0x6f, 0x83, 0xf8, 0x62, 0x23, 0xe3, 0x00, 0xb7, 0xf0, 0x18, 0x0d, 0xaa, 0xbb, 0x5a,
	0xcd, 0x88, 0x4d, 0x76, 0xf8, 0x21, 0x8d, 0x32, 0x18, 0xe0, 0x5e, 0xe3, 0x5e, 0xef, 0x9e, 0xe4,
	0x50, 0x24, 0xfb, 0x4b, 0x62, 0x4b, 0xad, 0x89, 0xe3, 0xdd, 0xbb, 0xfc, 0x63, 0x4d, 0xb3, 0x48,
	0x16, 0x8b, 0xf5, 0xf1, 0x63, 0xb1, 0x98, 0xc7, 0x64, 0x4d, 0x35, 0x7c, 0xa0, 0xa6, 0x4a, 0x6e,
	0x93, 0x94, 0x36, 0x2c, 0xdd, 0xbd, 0x4c, 0x16, 0x49, 0x2e, 0xeb, 0x46, 0x0a, 0x10, 0x46, 0x27,
	0xb7, 0x8f, 0x4f, 0x8b, 0xa4, 0x80, 0x92, 0x09, 0x66, 0x98, 0x14, 0xf8, 0x21, 0x79, 0xad, 0x0c,
	0x2b, 0x69, 0x6e, 0xf0, 0x77, 0xa3, 0x64, 0x03, 0xca, 0x30, 0x70, 0x2d, 0x42, 0x16, 0x90, 0x31,
	0x51, 0xc0, 0x03, 0xfe, 0x65, 0x0e, 0x0d, 0x24, 0xb7, 0x09, 0x13, 0x06, 0x36, 0xa0, 0x92, 0xa7,
	0x45, 0xd2, 0xb8, 0x99, 0xba, 0x16, 0x6d, 0x14, 0x13, 0x1b, 0xd7, 0xa0, 0xc0, 0x98, 0x43, 0x36,
	0xd9, 0x6e, 0x15, 0x8f, 0x7c, 0x7f, 0x5a, 0xb4, 0x5f, 0xe4, 0xfa, 0x67, 0xc8, 0x0d, 0x52, 0xb6,
	0xcc, 0x39, 0x7e, 0x98, 0x81, 0xda, 0xfd, 0xf8, 0x17, 0x05, 0x65, 0x72, 0x9b, 0xfc, 0x21, 0x2d,
	0xfa, 0xa5, 0xa4, 0xdd, 0x3a, 0xfa, 0x71, 0xa8, 0x52, 0xf4, 0xe0, 0x86, 0xd9, 0xeb, 0xd3, 0xe5,
	0x6d, 0xe1, 0xd0, 0x50, 0xa6, 0x7a, 0x4e, 0x3e, 0xb5, 0xac, 0x2c, 0x12, 0x61, 0x39, 0x4f, 0x3e,
	0x47, 0x59, 0xfa, 0xb3, 0x65, 0xbc, 0x38, 0x1d, 0x6d, 0x2d, 0x8b, 0x03, 0xfe, 0x5b, 0x80, 0xce,
	0x15, 0x6b, 0x90, 0xa7, 0xe4, 0x16, 0x65, 0x5f, 0x33, 0x43, 0x6a, 0xd0, 0x9a, 0x6e, 0x80, 0x38,
	0xa2, 0xc5, 0xa9, 0x40, 0xd6, 0x8a, 0x8a, 0x3c, 0x2e, 0xab, 0x35, 0x4e, 0x97, 0x19, 0x56, 0x43,
	0x56, 0x33, 0xce, 0x99, 0x8e, 0xef, 0x85, 0x27, 0x0b, 0x82, 0x2d, 0xa5, 0xaa, 0xa9, 0x49, 0x6e,
	0x13, 0xab, 0x58, 0x6c, 0x3a, 0xcf, 0x95, 0x01, 0x95, 0x41, 0x4d, 0xd9, 0xb8, 0x8b, 0xff, 0x72,
	0xb6, 0x93, 0xa0, 0x35, 0x44, 0x79, 0x2d, 0xa4, 0x30, 0xd9, 0xba, 0x95, 0xcf, 0x58, 0x12, 0x0a,
	0xa8, 0x96, 0x82, 0xec, 0xab, 0x03, 0xd9, 0x03, 0x29, 0x58, 0x21, 0x5e, 0x18, 0xe2, 0x68, 0x17,
	0x84, 0x95, 0xc7, 0xdf, 0x3a, 0x06, 0x4e, 0xf7, 0x64, 0x91, 0x70, 0x56, 0x42, 0x7e, 0xc8, 0x39,
	0x4c, 0xea, 0xc2, 0xbb, 0x8e, 0xc2, 0x6b, 0xe3, 0x8e, 0x49, 0x3b, 0xad, 0x39, 0x1f, 0x02, 0x81,
	0xdf, 0xd9, 0xa7, 0x45, 0xf2, 0x8b, 0x05, 0x0b, 0x45, 0x46, 0xcd, 0xe9, 0x32, 0x70, 0x23, 0x3c,
	0x8b, 0x64, 0x4f, 0x35, 0xf1, 0x94, 0xc9, 0xa2, 0x97, 0x5f, 0x41, 0x0d, 0x2c, 0x91, 0x2a, 0x26,
	0x43, 0x05, 0x8d, 0x9c, 0x94, 0x9e, 0x02, 0xa3, 0x0e, 0x99, 0x2c, 0x4f, 0x27, 0xf5, 0x9b, 0x2b,
	0x6c, 0x4d, 0x64, 0x49, 0x4c, 0xd5, 0x32, 0x60, 0x2a, 0xa6, 0x09, 0xd3, 0x84, 0x12, 0xd7, 0x95,
	0xc8, 0x72, 0x20, 0xb6, 0x56, 0x37, 0x06, 0x72, 0xd3, 0x86, 0x2a, 0xe3, 0x54, 0xe9, 0xec, 0xba,
	0x1c, 0xd9, 0x15, 0x8b, 0xd2, 0x46, 0x36, 0x97, 0x47, 0x45, 0x71, 0xeb, 0xea, 0x9a, 0x61, 0xad,
	0x37, 0xb9, 0xa8, 0x03, 0xd1, 0xa0, 0x26, 0xe5, 0xb8, 0xcb, 0xf5, 0x5c, 0x43, 0xd8, 0x57, 0x11,
	0x9b, 0xd5, 0x95, 0x54, 0x86, 0x78, 0x22, 0x02, 0x0f, 0x0d, 0xa7, 0x4c, 0xe0, 0x4f, 0x14, 0x7c,
	0xab, 0xc5, 0x7e, 0x55, 0xe6, 0x74, 0xcc, 0x49, 0x8f, 0xf1, 0x16, 0x4c, 0x30, 0xb5, 0xf1, 0x6c,
	0xad, 0x76, 0x76, 0x7a, 0x3f, 0x76, 0x2b, 0x94, 0xf3, 0xcc, 0x5b, 0x5e, 0x56, 0xb8, 0x11, 0xa6,
	0xd5, 0xf8, 0x8d, 0x23, 0x7b, 0x1b, 0xa8, 0xbc, 0xc5, 0x36, 0x54, 0xc1, 0xa4, 0x20, 0x62, 0xb6,
	0xf5, 0xb3, 0x5c, 0x4f, 0xdb, 0x37, 0x06, 0x84, 0x41, 0x83, 0xeb, 0x34, 0x30, 0xb0, 0x4c, 0xdb,
	0x3c, 0x07, 0xad, 0x4b, 0xcb, 0x7b, 0x2f, 0x30, 0xcf, 0xe2, 0x50, 0x81, 0xc3, 0x8a, 0x5b, 0xde,
	0xd6, 0x52, 0x72, 0xa0, 0x62, 0xa4, 0xc0, 0xba, 0xca, 0x40, 0xd0, 0x35, 0x87, 0xe2, 0x02, 0x25,
	0x2a, 0x55, 0x21, 0xad, 0xb9, 0x40, 0x66, 0xd1, 0x2b, 0x67, 0x23, 0x73, 0x9f, 0xa7, 0x9c, 0xa8,
	0x80, 0x93, 0xab, 0xfb, 0x09, 0x1b, 0xa7, 0x35, 0xe1, 0x03, 0x55, 0xb4, 0x06, 0x03, 0x4a, 0x9f,
	0x6a, 0xc3, 0xeb, 0xa2, 0x70, 0x83, 0x50, 0x4e, 0x40, 0xec, 0x98, 0x92, 0xa2, 0x06, 0x61, 0xc8,
	0x8e, 0x2a, 0x86, 0xcb, 0xd6, 0xc4, 0x48, 0xc2, 0x04, 0x0e, 0x47, 0x98, 0x30, 0x72, 0xe0, 0x0c,
	0x06, 0xe4, 0x2b, 0xf2, 0x9a, 0xd4, 0xb4, 0x41, 0x67, 0x81, 0x5b, 0xe9, 0x3a, 0xed, 0x28, 0xb7,
	0xa0, 0x57, 0xf7, 0x22, 0x99, 0x62, 0xeb, 0xce, 0xd6, 0x35, 0x55, 0x87, 0xd3, 0xc8, 0x46, 0x8b,
	0xe2, 0x4a, 0xe9, 0x74, 0xfe, 0x2a, 0x1e, 0xab, 0xa4, 0x35, 0xb9, 0xac, 0xa7, 0xbd, 0xf7, 0xfb,
	0xd0, 0x8e, 0xaa, 0x65, 0xd1, 0x67, 0x5c, 0x35, 0xb9, 0x36, 0xd4, 0x9c, 0x71, 0xf6, 0x77, 0xbe,
	0x39, 0xf8, 0x09, 0xd4, 0x43, 0xed, 0x44, 0x3f, 0x45, 0x5e, 0xd1, 0x97, 0x67, 0xf6, 0x72, 0x06,
	0x2e, 0xe9, 0x74, 0xfc, 0x18, 0x94, 0x0c, 0xed, 0x35, 0x22, 0x75, 0x6b, 0x2a, 0xa9, 0x32, 0x5c,
	0xea, 0xfc, 0xb5, 0x87, 0x4e, 0x57, 0x44, 0xf5, 0xd0, 0x83, 0xcb, 0x0d, 0x13, 0x51, 0x9b, 0x0f,
	0x04, 0x93, 0x3e, 0xa1, 0xc5, 0x3d, 0x13, 0x78, 0xe1, 0x82, 0x60, 0x5b, 0xb2, 0xeb, 0x91, 0xcb,
	0x75, 0x82, 0xf9, 0x8d, 0x88, 0x67, 0x5a, 0x3c, 0x33, 0x50, 0xd1, 0x74, 0x20, 0x8b, 0xa9, 0xd5,
	0xd8, 0x8b, 0x5f, 0xd4, 0xae, 0x91, 0x12, 0x45, 0x94, 0xec, 0x2f, 0x62, 0xb7, 0xa3, 0x2a, 0x82,
	0xed, 0x27, 0xa3, 0x27, 0x3a, 0x8a, 0x99, 0xbc, 0xfa, 0xc1, 0x67, 0x70, 0x19, 0xb8, 0x88, 0xf0,
	0xf7, 0x3d, 0x44, 0x3c, 0x4e, 0xc9, 0xc4, 0x06, 0x54, 0xa3, 0x98, 0x30, 0x53, 0xe7, 0x8b, 0x12,
	0x94, 0x1a, 0x46, 0x81, 0xce, 0xbb, 0x3b, 0xa7, 0xb1, 0xe6, 0x2c, 0xcf, 0xb6, 0x10, 0x57, 0xcc,
	0x38, 0x54, 0xd9, 0x57, 0x20, 0x9c, 0x3b, 0xdd, 0xc2, 0xc1, 0x41, 0x3b, 0xa6, 0xf5, 0x35, 0xd0,
	0xce, 0x7f, 0x38, 0xc1, 0xff, 0x54, 0x90, 0x35, 0x90, 0xfb, 0xa4, 0x80, 0x86, 0xcb, 0xc3, 0x72,
	0x0b, 0x87, 0xfb, 0x84, 0x48, 0x45, 0xee, 0x93, 0x0d, 0x33, 0x95, 0x5d, 0x2f, 0xad, 0x06, 0xe5,
	0x3f, 0x3b, 0xdf, 0x0c, 0xce, 0x6b, 0x7e, 0x1a, 0xd0, 0x27, 0x8b, 0x63, 0xd2, 0xe4, 0xf3, 0x2c,
	0xd8, 0xf1, 0x3d, 0x1c, 0x66, 0x6c, 0x0d, 0x6e, 0x40, 0x64, 0x5f, 0xde, 0x0d, 0x31, 0x76, 0xcb,
	0x54, 0x87, 0x76, 0x75, 0x5e, 0x41, 0x61, 0xb9, 0xfb, 0x2d, 0xa4, 0xc9, 0x94, 0x15, 0xfd, 0x2f,
	0xe1, 0x21, 0x45, 0xff, 0xab, 0x83, 0x7e, 0x9f, 0x23, 0x72, 0x7b, 0xdf, 0x07, 0x83, 0x76, 0x9a,
	0x9c, 0x8a, 0x1c, 0xfc, 0xe0, 0x4c, 0x94, 0x8a, 0x6a, 0xa3, 0x6c, 0x6e, 0xac, 0x82, 0xac, 0x0c,
	0x96, 0xda, 0x46, 0xf7, 0x45, 0x82, 0x5f, 0x02, 0x1b, 0x99, 0x01, 0x6d, 0x34, 0x72, 0xe7, 0x61,
	0x48, 0x74, 0x3a, 0x0c, 0xbe, 0x9c, 0x43, 0x04, 0x88, 0x7d, 0xac, 0x80, 0x08, 0x5b, 0xaf, 0x41,
	0x61, 0xe0, 0xcc, 0xa5, 0x30, 0x94, 0x09, 0x50, 0x2e, 0x7a, 0x5a, 0x0d, 0xf8, 0x8f, 0xb2, 0xa2,
	0x8f, 0xb8, 0x2b, 0xf2, 0x16, 0x4a, 0x6a, 0xb9, 0x41, 0x08, 0x8e, 0x88, 0x82, 0x50, 0x51, 0xb8,
	0xe6, 0x46, 0x49, 0x17, 0xa1, 0x8b, 0xbe, 0xdd, 0x6a, 0x28, 0x46, 0xc1, 0x77, 0xc0, 0xd1, 0x08,
	0x0b, 0x5d, 0x8b, 0x0f, 0x2f, 0x84, 0xda, 0x99, 0xa7, 0xc7, 0x99, 0x31, 0x33, 0xa6, 0x65, 0x1f,
	0xfc, 0x62, 0x23, 0x01, 0x6c, 0x7f, 0xe6, 0xac, 0xbe, 0xd7, 0xfd, 0xb1, 0x77, 0x08, 0xfb, 0x06,
	0x03, 0xe7, 0xb4, 0x6e, 0x4a, 0xa6, 0x20, 0x13, 0xd2, 0xb0, 0xf2, 0x90, 0xa1, 0xd1, 0x9f, 0x63,
	0x11, 0xe5, 0xf9, 0x5e, 0xfd, 0x10, 0x30, 0x69, 0xd7, 0x5b, 0x49, 0x59, 0x5f, 0xdf, 0x4b, 0xdb,
	0x75, 0x21, 0x6b, 0xca, 0xc4, 0xf5, 0x5d, 0x8d, 0xdc, 0xc2, 0x15, 0xdd, 0x64, 0xdd, 0x30, 0x3e,
	0x71, 0x9c, 0xf6, 0xfa, 0x93, 0x9d, 0xc9, 0x0e, 0x14, 0xd0, 0x80, 0x28, 0x40, 0xe4, 0x23, 0xf0,
	0x3c, 0x20, 0x80, 0x07, 0xa3, 0x68, 0xb4, 0xa5, 0x04, 0xea, 0x6d, 0x8a, 0xd3, 0x4d, 0x24, 0x3b,
	0xe2, 0x34, 0x67, 0x59, 0x4a, 0xb5, 0x5d, 0x36, 0x4a, 0xc7, 0x1d, 0x6d, 0xc9, 0x01, 0x2e, 0x41,
	0xec, 0x9f, 0xad, 0x60, 0x26, 0xde, 0x5d, 0x6a, 0x3d, 0xd5, 0xf0, 0x10, 0x6f, 0xd0, 0x60, 0x96,
	0xc1, 0x19, 0xf6, 0x2a, 0x7b, 0x4a, 0x66, 0x94, 0xd5, 0xe6, 0xb0, 0x5c, 0x83, 0xa1, 0x31, 0x82,
	0x98, 0x1a, 0x97, 0x5c, 0xee, 0x0b, 0x99, 0x6f, 0x33, 0xda, 0xb0, 0x2b, 0xb7, 0xb0, 0x94, 0x9c,
	0xcb, 0xfd, 0x54, 0x30, 0xaa, 0xa8, 0xce, 0xac, 0x46, 0xf0, 0x7e, 0x14, 0x90, 0x86, 0x34, 0xa0,
	0xe4, 0xd6, 0x66, 0xde, 0xe7, 0x67, 0x67, 0xcf, 0x16, 0xc7, 0x93, 0x57, 0xac, 0xc9, 0x2b, 0x6a,
	0x9e, 0xc1, 0x76, 0xdb, 0xd3, 0x9b, 0xd7, 0x33, 0xbb, 0xf5, 0x56, 0x39, 0x7d, 0xa4, 0x6c, 0x7b,
	0x5c, 0x65, 0x89, 0x4c, 0xe5, 0x59, 0x5e, 0x51, 0x21, 0xbc, 0xab, 0x9e, 0xdd, 0x67, 0x0b, 0x87,
	0xbd, 0x54, 0xc5, 0x55, 0x7d, 0x9e, 0xe5, 0x60, 0xb0, 0x63, 0x43, 0xb5, 0xbe, 0x7a, 0x36, 0x0d,
	0x6a, 0x77, 0xc5, 0x06, 0x63, 0x97, 0x61, 0xba, 0x63, 0x56, 0x27, 0x4e, 0xc5, 0xc6, 0xd2, 0x4d,
	0xdc, 0xad, 0x4c, 0x9a, 0x5d, 0x33, 0x88, 0x8d, 0xa7, 0x31, 0xe2, 0x6c, 0xee, 0x4a, 0xe7, 0xb2,
	0x81, 0x19, 0x68, 0xe3, 0x0e, 0xe9, 0x62, 0x78, 0x43, 0x83, 0xb1, 0x4d, 0x7c, 0x68, 0x4e, 0x9f,
	0x65, 0x96, 0xbe, 0xdf, 0xb5, 0x3a, 0x34, 0xea, 0x95, 0xc9, 0x1d, 0x28, 0xc5, 0x0a, 0xb8, 0xb2,
	0xfb, 0xb3, 0x14, 0xca, 0x77, 0xbd, 0x3e, 0xf0, 0xf8, 0x7e, 0x7b, 0x58, 0x57, 0x52, 0x6e, 0xe7,
	0x9e, 0xa3, 0xb4, 0xae, 0xd0, 0x54, 0xc6, 0x3b, 0x36, 0x81, 0x69, 0xfb, 0x4d, 0x42, 0x80, 0xf5,
	0xbc, 0x7c, 0x5c, 0x8b, 0x8e, 0x9f, 0x9f, 0xac, 0x3b, 0x83, 0x3d, 0x66, 0xe8, 0x5d, 0x8b, 0x52,
	0x22, 0x8b, 0xfa, 0x71, 0x90, 0x02, 0x38, 0x45, 0x84, 0xba, 0x81, 0x9c, 0x95, 0x2c, 0x27, 0x6d,
	0xa6, 0x00, 0x71, 0xa0, 0xc7, 0x7f, 0xf7, 0x62, 0x0a, 0x00, 0x56, 0x40, 0x8b, 0x2e, 0x55, 0xeb,
	0x82, 0x77, 0x87, 0x01, 0xc9, 0x1b, 0x2a, 0x84, 0x34, 0x78, 0x28, 0xc0, 0xbf, 0xc9, 0x9e, 0x99,
	0x8a, 0x18, 0xba, 0x21, 0x4d, 0x9b, 0x15, 0x9a, 0x80, 0x89, 0xde, 0x72, 0x06, 0x28, 0x79, 0xaf,
	0x98, 0x81, 0xa5, 0x06, 0x63, 0x98, 0xd8, 0x20, 0xec, 0xdd, 0x31, 0xd8, 0x2f, 0xd7, 0x3e, 0x2b,
	0x81, 0xe6, 0x4a, 0x8b, 0x61, 0xab, 0x51, 0x6c, 0xb3, 0x01, 0xd5, 0x13, 0x50, 0x8e, 0x50, 0x3a,
	0x84, 0x51, 0xc4, 0xcf, 0x02, 0xa2, 0x88, 0xd9, 0x1d, 0xd4, 0x63, 0x5b, 0x77, 0xd7, 0x45, 0xe0,
	0x96, 0x25, 0x9f, 0xc5, 0x43, 0x90, 0x7a, 0x0d, 0x84, 0xef, 0x0f, 0x10, 0x83, 0x23, 0x43, 0x0b,
	0xeb, 0xcf, 0x1d, 0x39, 0x02, 0xed, 0x00, 0xf9, 0x97, 0xec, 0xc1, 0x13, 0x9f, 0x39, 0x01, 0x8c,
	0xec, 0xe7, 0x6c, 0xe0, 0xfa, 0x48, 0x37, 0x71, 0xa5, 0xc0, 0xfd, 0xea, 0xd4, 0xe0, 0xf8, 0x18,
	0x10, 0xdf, 0xe2, 0x4e, 0x7d, 0x2e, 0xed, 0xf3, 0x47, 0xb7, 0x94, 0x13, 0x38, 0x66, 0xda, 0xcf,
	0x9d, 0x9e, 0x8f, 0xdb, 0x73, 0x4e, 0xb5, 0x9e, 0x34, 0xc1, 0x72, 0x0a, 0x62, 0x86, 0x8b, 0xa9,
	0x78, 0xb6, 0x77, 0xfa, 0xa2, 0x42, 0x5b, 0x6e, 0x66, 0x64, 0xd8, 0x94, 0x15, 0xdd, 0xc5, 0x40,
	0x97, 0x33, 0xc6, 0x13, 0x96, 0x73, 0x41, 0xd2, 0xaa, 0x7c, 0x6e, 0xa6, 0x61, 0x6c, 0xb4, 0x31,
	0x8a, 0x9f, 0x02, 0x6a, 0x3a, 0xce, 0x5d, 0xd6, 0xa3, 0xc4, 0xcd, 0x20, 0xda, 0x51, 0xce, 0x7d,
	0x32, 0x68, 0x2c, 0xd6, 0xcb, 0x69, 0xa1, 0x63, 0x07, 0x42, 0x05, 0xe5, 0x07, 0xc3, 0x72, 0x9d,
	0xb1, 0x22, 0x9e, 0x42, 0xdb, 0x51, 0x43, 0xd5, 0x5c, 0xd7, 0xbc, 0xa6, 0x9a, 0xe5, 0x9e, 0xb3,
	0x63, 0x74, 0x35, 0xa4, 0x62, 0x66, 0x6d, 0xf3, 0xed, 0x18, 0x7e, 0x47, 0xee, 0x7c, 0x3a, 0xba,
	0xcc, 0x27, 0xf2, 0xd8, 0xaf, 0x53, 0x40, 0xb5, 0x3f, 0xf1, 0xc6, 0xe3, 0x7f, 0xae, 0x80, 0x9a,
	0x2b, 0x73, 0xb1, 0x05, 0x3d, 0xe8, 0x8c, 0x43, 0x69, 0x32, 0x26, 0x32, 0xa3, 0x18, 0x9d, 0xc0,
	0x16, 0x05, 0xec, 0xb2, 0x33, 0x5b, 0x05, 0x42, 0x49, 0xce, 0xa1, 0xc8, 0x10, 0xd1, 0x5f, 0x19,
	0xb3, 0xfc, 0x71, 0x61, 0xb4, 0x35, 0x51, 0x41, 0x05, 0x3a, 0x89, 0x62, 0xca, 0x22, 0x70, 0xe6,
	0xf2, 0x44, 0x2a, 0xec, 0xf3, 0xc5, 0xa9, 0x02, 0xdc, 0x47, 0x48, 0x33, 0x3a, 0x12, 0xc4, 0x7c,
	0x10, 0x13, 0x6e, 0xd1, 0x59, 0xa3, 0xe4, 0x46, 0xd1, 0x3a, 0x2e, 0x9f, 0xe9, 0xf4, 0xe4, 0xa4,
	0x0d, 0x4b, 0xb5, 0xa1, 0x82, 0xfd, 0x4a, 0xd1, 0x64, 0x4f, 0xd4, 0xac, 0xb7, 0xaa, 0x16, 0x12,
	0x32, 0x3d, 0x91, 0x60, 0x68, 0x38, 0x15, 0xe7, 0x57, 0xd0, 0x0c, 0x82, 0xf4, 0xe9, 0xf8, 0x56,
	0x57, 0xa0, 0xa6, 0x0c, 0x47, 0x03, 0x87, 0x1c, 0x75, 0xee, 0x8a, 0xa4, 0xad, 0x66, 0x1b, 0x81,
	0xda, 0x96, 0x4b, 0x3b, 0xca, 0x1e, 0x0e, 0x78, 0x76, 0x8a, 0x98, 0x81, 0x28, 0xe6, 0x2a, 0xf2,
	0xa9, 0xbf, 0xc1, 0x5d, 0x94, 0x08, 0x86, 0x92, 0x9c, 0xa9, 0x9c, 0x43, 0xce, 0x56, 0xb9, 0xac,
	0x7d, 0xb4, 0x93, 0xf1, 0xa0, 0xf1, 0xc6, 0x11, 0xbe, 0xf9, 0x8e, 0xbc, 0xfe, 0xf0, 0x9d, 0xbf,
	0xb5, 0xfd, 0xf1, 0x2f, 0x77, 0x1f, 0x4b, 0xcb, 0x17, 0xa4, 0xb4, 0x9c, 0x1f, 0x96, 0xe1, 0x40,
	0x5e, 0x38, 0x02, 0x53, 0x51, 0x43, 0x28, 0x1e, 0x2c, 0x35, 0x39, 0x48, 0x8b, 0xe1, 0xa6, 0x90,
	0x84, 0xf2, 0x5a, 0x6a, 0x43, 0xa8, 0x38, 0x98, 0x8a, 0x89, 0x0d, 0x61, 0xa2, 0x1b, 0x76, 0x75,
	0x2f, 0xfe, 0x2e, 0x2d, 0xc9, 0xa9, 0x20, 0xd4, 0x05, 0x3f, 0xec, 0x4d, 0x90, 0x1d, 0x5c, 0x1f,
	0x46, 0x1e, 0x87, 0x4e, 0x3c, 0x0a, 0x70, 0x6d, 0x34, 0x77, 0xfe, 0x7a, 0x75, 0x2f, 0x90, 0x3d,
	0x29, 0xf8, 0x81, 0xf8, 0x51, 0xf7, 0x40, 0x0a, 0x29, 0xfe, 0xe7, 0xbf, 0xfe, 0xdb, 0x90, 0x46,
	0xc9, 0x1d, 0x2b, 0xa0, 0x1d, 0xd2, 0x48, 0xe4, 0x7c, 0xcd, 0x38, 0x47, 0xba, 0xd2, 0x0a, 0x3f,
	0xc4, 0x82, 0xec, 0x2b, 0x96, 0x57, 0xa4, 0xb6, 0xda, 0x85, 0xbd, 0x42, 0x0a, 0x20, 0xa5, 0x92,
	0xb5, 0x43, 0x40, 0xdd, 0xc2, 0xf7, 0xb0, 0x26, 0x3f, 0x7d, 0xe7, 0xe3, 0x1e, 0x67, 0x39, 0x08,
	0x0d, 0x7d, 0x2a, 0x3a, 0x79, 0x6d, 0x8c, 0x62, 0x6b, 0x8b, 0x03, 0x2e, 0x7f, 0x90, 0xe2, 0x8d,
	0xac, 0x6b, 0x50, 0x39, 0xa3, 0x7c, 0x79, 0x57, 0x51, 0x05, 0xaf, 0x39, 0xdb, 0x02, 0xf9, 0x76,
	0x75, 0x43, 0xbe, 0x13, 0x06, 0xd1, 0xa6, 0xbf, 0x26, 0x4b, 0x42, 0x35, 0x49, 0x52, 0x19, 0xd3,
	0xdc, 0xa6, 0xa9, 0xf3, 0x52, 0x6c, 0x07, 0xb9, 0xac, 0x6b, 0x5c, 0x9c, 0x54, 0x9b, 0x34, 0xcc,
	0xa5, 0xd3, 0xf5, 0x61, 0x29, 0xf2, 0xa5, 0xa6, 0xe9, 0xb7, 0xab, 0x9b, 0xd4, 0xa7, 0x82, 0x0d,
	0x06, 0xc6, 0xa4, 0x65, 0x11, 0x51, 0x14, 0x28, 0x8f, 0x06, 0x13, 0x57, 0x3c, 0xf3, 0xb0, 0xa4,
	0x0d, 0xd3, 0x1b, 0xab, 0xec, 0x32, 0xa7, 0x06, 0x36, 0xd2, 0x5f, 0x53, 0x7e, 0x42, 0x97, 0x05,
	0x1c, 0x43, 0x4d, 0x66, 0xa4, 0xe4, 0x0e, 0x6a, 0x3c, 0x2c, 0xb9, 0xdc, 0xb8, 0xdd, 0xef, 0x59,
	0xd2, 0xb7, 0x69, 0x4a, 0x1b, 0xb6, 0xc2, 0x51, 0x56, 0x38, 0x4c, 0xba, 0xfb, 0x26, 0xcd, 0x69,
	0x5e, 0x41, 0x8a, 0xc4, 0xa9, 0xa3, 0xc9, 0xbc, 0xe7, 0xd0, 0x37, 0x2b, 0xef, 0x85, 0x10, 0x4e,
	0xa3, 0x3b, 0x06, 0x61, 0x50, 0xad, 0x32, 0x9b, 0xbd, 0xfc, 0xe6, 0xd5, 0xcb, 0x7f, 0xfb, 0xd7,
	0x9b, 0x55, 0xe3, 0x95, 0xfc, 0x61, 0x29, 0x15, 0x73, 0x36, 0xff, 0x69, 0xa0, 0xc1, 0x7a, 0x4f,
	0x37, 0xde, 0xd5, 0x8c, 0xe7, 0x57, 0x74, 0x1f, 0x1f, 0x38, 0x5d, 0x73, 0x99, 0x6f, 0x41, 0x30,
	0x99, 0x22, 0xce, 0x5d, 0xb6, 0xba, 0x9c, 0xd6, 0x54, 0x1b, 0x50, 0x69, 0x18, 0x70, 0x75, 0xa8,
	0xf9, 0x48, 0x32, 0xdf, 0xac, 0x6e, 0x92, 0x27, 0xb7, 0xe0, 0x41, 0x26, 0xde, 0x28, 0x0b, 0xfe,
	0x93, 0x53, 0x19, 0xf5, 0x83, 0xdf, 0xd6, 0x23, 0xfb, 0x78, 0x58, 0x5a, 0x21, 0xcb, 0x92, 0xe1,
	0xc6, 0xde, 0x35, 0x90, 0xfb, 0x7e, 0xe8, 0x90, 0x1a, 0x10, 0xb4, 0x61, 0xc9, 0x6d, 0x92, 0x78,
	0xb7, 0xd3, 0x5f, 0x91, 0xf6, 0xf9, 0xaf, 0x53, 0xab, 0x0a, 0xf0, 0x1a, 0x75, 0x88, 0xe8, 0x4a,
	0x5a, 0x5e, 0x38, 0xdc, 0xa5, 0xf8, 0x12, 0x44, 0x2e, 0x8b, 0x36, 0xeb, 0xca, 0x5c, 0x2a, 0x95,
	0x9a, 0x2a, 0x69, 0x7d, 0x62, 0x3b, 0x2a, 0xc2, 0x99, 0x5f, 0x2c, 0x1b, 0x2c, 0xe1, 0xfc, 0x75,
	0xe6, 0x98, 0x85, 0xff, 0xec, 0xfe, 0x9b, 0x9c, 0xa6, 0xeb, 0x3c, 0x39, 0xd3, 0xc0, 0x35, 0x81,
	0xbb, 0x28, 0xe9, 0xee, 0x92, 0xae, 0x9d, 0xcc, 0xfd, 0x33, 0x63, 0x45, 0x25, 0xe3, 0xc6, 0xc3,
	0xa5, 0xf1, 0x0c, 0x3f, 0x02, 0xd2, 0xe4, 0x46, 0x07, 0x83, 0xf6, 0x87, 0x05, 0x42, 0x15, 0x10,
	0x05, 0xc6, 0x2a, 0xe1, 0xe4, 0x79, 0x07, 0x06, 0x3d, 0xc1, 0xbd, 0x4b, 0x5e, 0x72, 0x30, 0x50,
	0xdc, 0x27, 0x0b, 0x72, 0x9f, 0xf4, 0xd7, 0xff, 0xfe, 0x6f, 0x0f, 0xe2, 0xfd, 0xef, 0x00, 0xd7,
	0xf1, 0x0f, 0xa9, 0xda, 0xcc, 0xb8, 0x73, 0x28, 0x42, 0x12, 0xcf, 0xce, 0x6a, 0x74, 0x01, 0xd2,
	0x8d, 0xdd, 0xc3, 0xf9, 0xd2, 0xf2, 0xe1, 0xd1, 0xa0, 0x3d, 0x01, 0x7c, 0x0e, 0xa2, 0xf8, 0xc5,
	0x82, 0x3a, 0xf4, 0xb2, 0x08, 0x6b, 0x8c, 0xae, 0x7e, 0x74, 0xbd, 0x74, 0xad, 0x90, 0x87, 0xdd,
	0x67, 0xc8, 0x9a, 0xb3, 0x70, 0x0d, 0x1a, 0x16, 0x9d, 0xdc, 0xbe, 0xba, 0x59, 0x9c, 0xbd, 0x7b,
	0x08, 0x52, 0x37, 0x32, 0x08, 0x7d, 0x45, 0xfe, 0x83, 0x3e, 0xb0, 0xda, 0xd6, 0xe4, 0xe5, 0xcd,
	0xcd, 0x62, 0x24, 0xbc, 0x57, 0x37, 0x03, 0xfd, 0x6e, 0x05, 0x50, 0x7b, 0xea, 0xe4, 0x16, 0xc9,
	0x5b, 0xae, 0x3d, 0x1b, 0x31, 0x9d, 0x93, 0x65, 0xa9, 0x61, 0xcc, 0x61, 0x94, 0x41, 0x0c, 0x4b,
	0x9e, 0x21, 0xdd, 0xb2, 0xe8, 0x4a, 0x85, 0x5c, 0x28, 0xf0, 0xce, 0x9e, 0x69, 0xe2, 0x47, 0x1b,
	0x73, 0x19, 0x63, 0x32, 0xb0, 0x15, 0x26, 0x8f, 0xf1, 0xd5, 0xf4, 0x37, 0x0b, 0x13, 0x5b, 0x34,
	0xb5, 0x43, 0x6d, 0xcf, 0x19, 0xbb, 0x33, 0xcc, 0x5c, 0x5c, 0x39, 0x4b, 0xd7, 0xf5, 0xe2, 0x34,
	0xa1, 0x56, 0xd2, 0x39, 0xb3, 0xd4, 0xcf, 0xb5, 0x81, 0xc8, 0xc2, 0x3e, 0x78, 0xd7, 0xa9, 0xc7,
	0x61, 0x7b, 0x2d, 0xad, 0x71, 0x81, 0x14, 0xa1, 0x0d, 0x14, 0x18, 0xf4, 0x71, 0x6a, 0x2f, 0x53,
	0x05, 0xba, 0x91, 0x18, 0xdc, 0x70, 0xb4, 0x6f, 0x6e, 0x6e, 0x22, 0xe5, 0x4c, 0xa3, 0x5e, 0xe1,
	0x24, 0x5d, 0xd3, 0x4b, 0xd5, 0x2b, 0x4f, 0xc8, 0x75, 0x1a, 0x24, 0x99, 0x3e, 0xb6, 0x8b, 0x7d,
	0x4a, 0x1f, 0xc3, 0xb7, 0xa7, 0xc9, 0x65, 0xfc, 0xd9, 0x97, 0x91, 0xf9, 0x7a, 0x12, 0x52, 0x4a,
	0x45, 0x80, 0xe6, 0x55, 0x9b, 0x10, 0xe1, 0x54, 0x1b, 0xf2, 0xea, 0xa6, 0x55, 0x21, 0x6c, 0xa6,
	0x44, 0x33, 0xb1, 0xe1, 0x40, 0x36, 0xcc, 0x10, 0x05, 0x8d, 0xf4, 0x6b, 0x1b, 0x46, 0x81, 0x4f,
	0x03, 0x6e, 0xfb, 0xef, 0xa9, 0xd7, 0xea, 0xa7, 0x45, 0xbc, 0x35, 0x28, 0xd7, 0x54, 0x73, 0xf0,
	0x10, 0x18, 0xc9, 0x2e, 0x8b, 0xf1, 0x9a, 0x45, 0x5d, 0x16, 0x72, 0x28, 0x1e, 0xf1, 0x52, 0x9e,
	0xb1, 0xce, 0x4e, 0xd7, 0xa6, 0xd6, 0xd2, 0xaa, 0x3c, 0x2e, 0xa6, 0x91, 0x3e, 0x4d, 0x97, 0x4b,
	0xa1, 0x6d, 0xed, 0xe1, 0x0a, 0x6d, 0x1a, 0xce, 0x72, 0xa7, 0x51, 0xe9, 0xcf, 0x5a, 0x0a, 0xf4,
	0x9a, 0x47, 0x26, 0xee, 0xa1, 0x21, 0x02, 0x53, 0x01, 0xfb, 0xb6, 0x34, 0xb3, 0x35, 0x78, 0xda,
	0x2d, 0x7d, 0x58, 0x82, 0x18, 0xdd, 0x27, 0x67, 0x2a, 0xa1, 0xd2, 0xb5, 0x0d, 0x82, 0xfe, 0xaf,
	0x5e, 0x28, 0xb1, 0x4b, 0xcc, 0x71, 0xd0, 0x9f, 0x16, 0xdb, 0xa0, 0x7e, 0xea, 0x28, 0x69, 0x1d,
	0x4f, 0xfd, 0xb5, 0x04, 0x2e, 0x73, 0x71, 0xa1, 0xdc, 0xa7, 0x4b, 0x06, 0x22, 0x38, 0xf4, 0xe9,
	0x9f, 0x28, 0xdd, 0x47, 0x1a, 0x3f, 0x16, 0x9c, 0xaa, 0xd2, 0xcb, 0x58, 0x29, 0xec, 0x39, 0x99,
	0xce, 0x54, 0x9e, 0xb6, 0x5c, 0x6b, 0x8e, 0xa5, 0xa6, 0xfe, 0x32, 0xcf, 0xc1, 0x4e, 0xcf, 0x0e,
	0x46, 0xd6, 0x53, 0xc6, 0xde, 0x70, 0xa0, 0x4a, 0x3b, 0x56, 0x1c, 0x6d, 0x30, 0xce, 0x30, 0xcc,
	0x6c, 0x87, 0xe3, 0xf2, 0x3f, 0x6d, 0xcd, 0xf3, 0x99, 0x6d, 0x3f, 0xb9, 0xc2, 0x3b, 0x73, 0xde,
	0xfa, 0xf2, 0x76, 0x72, 0x51, 0x68, 0x79, 0x05, 0xf9, 0x56, 0x5a, 0xb3, 0x0c, 0x87, 0xf4, 0xa8,
	0xaf, 0x7b, 0xc7, 0xb4, 0xd1, 0xa4, 0x25, 0x25, 0x5b, 0x38, 0xe8, 0xb9, 0x72, 0x1a, 0x75, 0xba,
	0xbc, 0xeb, 0xae, 0x9c, 0xe3, 0xf7, 0x75, 0x18, 0x47, 0xda, 0xe0, 0x12, 0x3e, 0xad, 0x3f, 0x18,
	0x72, 0xbb, 0xba, 0x17, 0xef, 0xf1, 0xb8, 0xe8, 0xef, 0x35, 0x7d, 0x82, 0x93, 0xba, 0xf0, 0xe2,
	0xcf, 0xae, 0x72, 0x0b, 0x22, 0xea, 0x1a, 0x22, 0xa9, 0xd4, 0x43, 0x03, 0xa8, 0xfe, 0x5b, 0x38,
	0x20, 0x52, 0xf0, 0x49, 0x26, 0x97, 0x42, 0x45, 0x1c, 0xff, 0xa2, 0x2f, 0x83, 0x79, 0x81, 0xb8,
	0xf1, 0xc5, 0x51, 0x25, 0xcc, 0x8b, 0x01, 0xb0, 0x18, 0x7b, 0x1c, 0xa7, 0x3f, 0x43, 0x99, 0x3e,
	0xb7, 0xb0, 0x66, 0x22, 0x38, 0x74, 0x22, 0xed, 0x1a, 0x53, 0x1c, 0x60, 0x8e, 0x31, 0x0e, 0xf5,
	0x2a, 0x7d, 0x1c, 0xa0, 0xc9, 0xa7, 0x73, 0xb6, 0xf9, 0xd6, 0x7d, 0x27, 0xf4, 0x78, 0x23, 0xbe,
	0xac, 0x49, 0x0e, 0xf2, 0xc1, 0xad, 0xc0, 0xde, 0x7f, 0x3f, 0xab, 0xde, 0xc8, 0x2d, 0x3c, 0x6a,
	0x22, 0xff, 0x0e, 0x66, 0x36, 0xdb, 0xd3, 0x62, 0xfd, 0xc2, 0x5a, 0x3f, 0x09, 0x09, 0x7a, 0x6c,
	0x3f, 0xc3, 0x45, 0x40, 0x57, 0x5d, 0x77, 0xc6, 0x39, 0xa0, 0x3b, 0x8d, 0x17, 0xf3, 0xa2, 0x7b,
	0xbd, 0x0d, 0x83, 0x5d, 0x23, 0x12, 0x08, 0x75, 0x77, 0xff, 0x30, 0x67, 0x10, 0x5b, 0xce, 0xf5,
	0x0b, 0x98, 0x65, 0x2d, 0x9e, 0x34, 0x7d, 0x74, 0xdf, 0x2f, 0x1b, 0xc8, 0xb4, 0xb8, 0x5d, 0x8e,
	0xa0, 0x20, 0x2f, 0x6e, 0xf1, 0xdf, 0x17, 0x33, 0x0d, 0xe7, 0x19, 0x83, 0x7e, 0x65, 0xd3, 0xf2,
	0xbc, 0x55, 0xac, 0x28, 0x40, 0xf8, 0x2a, 0x6f, 0x74, 0xa7, 0x51, 0x5e, 0x1d, 0x93, 0xcf, 0xdb,
	0xa6, 0xaf, 0x63, 0x7d, 0xc3, 0xec, 0xc7, 0x0c, 0xeb, 0xd3, 0xba, 0x6a, 0x63, 0xf3, 0x3f, 0x07,
	0x84, 0xf6, 0x46, 0x42, 0xa8, 0x20, 0x5a, 0x57, 0x2e, 0xa4, 0xb9, 0x0b, 0x40, 0x23, 0xdb, 0xfc,
	0x29, 0x3c, 0x18, 0x64, 0x82, 0x13, 0x7d, 0xd0, 0x06, 0x6a, 0xed, 0xd3, 0xbc, 0xe1, 0xf4, 0x48,
	0xee, 0xee, 0xfe, 0x86, 0x9d, 0x96, 0x6b, 0x8a, 0xbd, 0xa8, 0x35, 0x15, 0x08, 0x13, 0x66, 0x8c,
	0x45, 0xd3, 0x56, 0x8d, 0x62, 0x8c, 0xb9, 0x90, 0x58, 0x01, 0x2d, 0xfc, 0xd5, 0x83, 0x0f, 0x8a,
	0x6f, 0x7c, 0xea, 0x6f, 0xf9, 0xd1, 0x07, 0xc7, 0x4b, 0x87, 0xe3, 0x73, 0x50, 0xfe, 0xb8, 0xef,
	0x94, 0xd6, 0x57, 0x52, 0x9b, 0xc9, 0x4b, 0x88, 0x46, 0xb1, 0x1d, 0x35, 0x30, 0x51, 0xa9, 0x3b,
	0x03, 0x5e, 0x77, 0x09, 0x8b, 0x63, 0xc3, 0x10, 0x92, 0xb4, 0x84, 0x04, 0x1e, 0x1a, 0x77, 0x8b,
	0x90, 0x2c, 0x12, 0x78, 0xa0, 0x75, 0xc3, 0xc3, 0xd5, 0xe0, 0xb1, 0xd4, 0x86, 0x36, 0x9a, 0x20,
	0xea, 0x75, 0xcc, 0xb9, 0x6d, 0x64, 0x9a, 0x74, 0x0b, 0x7e, 0x7a, 0x9a, 0x65, 0xe3, 0x73, 0xac,
	0xfa, 0xa2, 0x86, 0x1b, 0x05, 0x90, 0x3e, 0xfa, 0x4c, 0xe5, 0xd3, 0x17, 0xd7, 0xf3, 0x89, 0xf6,
	0x90, 0x18, 0xfd, 0xea, 0x27, 0xc9, 0xf7, 0x4d, 0x78, 0xf9, 0xe2, 0x1f, 0xb4, 0xf4, 0x1c, 0x91,
	0x50, 0x45, 0xad, 0xc1, 0x10, 0xab, 0x99, 0xd8, 0xa0, 0x85, 0xe1, 0xae, 0x2a, 0x86, 0x0e, 0x8e,
	0x72, 0x44, 0x9d, 0xab, 0x7b, 0x71, 0x2f, 0x7e, 0x90, 0x06, 0x6e, 0xc9, 0x5f, 0xa5, 0x22, 0xb5,
	0x54, 0x10, 0x52, 0x27, 0xc7, 0xc3, 0x21, 0x2b, 0xb4, 0x08, 0x8d, 0x9f, 0xfc, 0x80, 0x5d, 0x2b,
	0xfb, 0x15, 0x8a, 0x70, 0x8a, 0xff, 0xfc, 0xc7, 0x36, 0x81, 0x3e, 0xcc, 0x61, 0xa7, 0x85, 0xcc,
	0x75, 0x3a, 0xa2, 0x0f, 0x35, 0x15, 0xe9, 0x9f, 0xfe, 0xff, 0x9e, 0x86, 0x7f, 0xb7, 0xa3, 0xae,
	0x77, 0x51, 0x6e, 0x80, 0x77, 0xd2, 0x6b, 0x55, 0x72, 0xfb, 0x38, 0xb6, 0xb1, 0x73, 0xcf, 0x6a,
	0xe6, 0x58, 0xd1, 0x63, 0x97, 0x83, 0x9f, 0xce, 0x59, 0xfd, 0xd5, 0x72, 0x4e, 0xc2, 0x0b, 0xbb,
	0x71, 0x4e, 0x2a, 0x54, 0x79, 0x7c, 0x74, 0xaf, 0xff, 0x82, 0x4b, 0x61, 0x22, 0xe7, 0xb6, 0x00,
	0x7f, 0xc7, 0x16, 0x56, 0x54, 0x32, 0x70, 0xf9, 0xac, 0xf6, 0xda, 0x6b, 0x3d, 0x4c, 0x18, 0xb9,
	0x4b, 0xb6, 0xf0, 0x5a, 0x93, 0x6b, 0xff, 0x6e, 0xab, 0xa1, 0x07, 0x2e, 0x69, 0xe1, 0x66, 0xc3,
	0xbf, 0x3f, 0xb9, 0xb2, 0xb1, 0x60, 0x5a, 0x24, 0x14, 0x74, 0xe9, 0xcf, 0x7f, 0xf4, 0x1a, 0x97,
	0x4b, 0x51, 0xb2, 0x8d, 0x55, 0xde, 0xf0, 0xfe, 0xe0, 0x4b, 0xcc, 0xfe, 0xb4, 0x20, 0x4c, 0x84,
	0x5c, 0x7d, 0x4e, 0x35, 0x84, 0x14, 0xec, 0xda, 0x3f, 0x10, 0xf3, 0xd0, 0xc0, 0x63, 0x02, 0x8c,
	0x41, 0xce, 0x95, 0x05, 0xc0, 0x12, 0x26, 0x7f, 0x31, 0xf7, 0x04, 0x71, 0x41, 0x3e, 0x33, 0x93,
	0x16, 0xed, 0x3b, 0x95, 0xaf, 0x86, 0x28, 0xfa, 0xdb, 0x97, 0x19, 0x78, 0x62, 0xa0, 0x27, 0x29,
	0x1d, 0x3e, 0x3c, 0x9f, 0x04, 0xf7, 0x4e, 0xc0, 0x1d, 0x29, 0x69, 0x94, 0x2c, 0x6c, 0x8e, 0x0e,
	0xe4, 0x40, 0x28, 0xd9, 0xb0, 0x1d, 0x88, 0x61, 0x9e, 0xec, 0xb2, 0x94, 0xaf, 0x19, 0xf3, 0xb2,
	0xc8, 0xfb, 0xc7, 0xf3, 0xff, 0x07, 0x04, 0xee, 0x2b, 0xc6, 0xbe, 0x5a, 0x8c, 0x1b, 0xb0, 0x39,
	0x79, 0xe4, 0x71, 0x1c, 0xe9, 0xde, 0x98, 0xe7, 0x65, 0x43, 0x2f, 0xef, 0xf2, 0x17, 0xc9, 0xff,
	0x5d, 0xed, 0xfa, 0x52, 0xf7, 0x48, 0xfc, 0x9f, 0x49, 0xc0, 0x3f, 0xfa, 0x47, 0xbf, 0xb3, 0x04,
	0xdc, 0x01, 0x89, 0xdf, 0x2c, 0xe4, 0x6e, 0xa4, 0xdf, 0x4d, 0xd0, 0x5d, 0xad, 0xde, 0xf9, 0x4b,
	0x1e, 0x24, 0x23, 0x35, 0x18, 0x5a, 0x50, 0x43, 0x83, 0x4f, 0x75, 0xa3, 0xb4, 0x38, 0xe6, 0xd3,
	0x3b, 0xa0, 0x4a, 0x90, 0x4a, 0xee, 0xd1, 0x7d, 0x3b, 0xf4, 0xd3, 0x90, 0x83, 0xb4, 0x6a, 0x70,
	0x43, 0x98, 0x4b, 0xce, 0xd1, 0xd9, 0x8f, 0x06, 0x3b, 0x87, 0x5c, 0x90, 0x70, 0xd9, 0x12, 0x06,
	0xc8, 0x72, 0x59, 0x9a, 0x67, 0x78, 0xbd, 0x2c, 0x46, 0x5f, 0xd2, 0xf8, 0x0f, 0xf7, 0x41, 0x17,
	0x5c, 0xb9, 0x2c, 0x5d, 0x34, 0x1f, 0x3c, 0xe0, 0x72, 0x95, 0x37, 0x2f, 0x14, 0x10, 0xff, 0xc2,
	0x03, 0x71, 0xa2, 0xec, 0x0b, 0x6e, 0x16, 0x3e, 0xef, 0xe9, 0xc3, 0xfc, 0xf0, 0xc2, 0x2e, 0x94,
	0x56, 0x79, 0x97, 0xed, 0xc1, 0xf4, 0x55, 0xfe, 0xff, 0x5a, 0x46, 0xe6, 0x6c, 0x41, 0x57, 0x29,
	0xdd, 0x2a, 0xb3, 0x82, 0x1c, 0x4f, 0x82, 0xeb, 0xee, 0x11, 0xf3, 0x6f, 0xb8, 0xc9, 0xf3, 0x63,
	0x05, 0xa5, 0x5c, 0x10, 0xa9, 0x0a, 0x50, 0x61, 0xf5, 0xed, 0x4e, 0x7c, 0xe1, 0x0b, 0xbd, 0x2f,
	0x78, 0x63, 0x37, 0x62, 0xfe, 0xaa, 0x8b, 0x3b, 0x14, 0x23, 0xea, 0x6a, 0xea, 0xeb, 0xf9, 0xba,
	0x54, 0x44, 0xd4, 0xd7, 0xbd, 0x2e, 0x0a, 0xed, 0x6d, 0xf7, 0x6f, 0x8e, 0xda, 0x25, 0xca, 0x43,
	0xd2, 0xbb, 0xd7, 0x28, 0x43, 0xb7, 0x08, 0x2a, 0x1b, 0x86, 0x2d, 0xd4, 0x01, 0xad, 0xda, 0x9f,
	0x4d, 0x1c, 0x6e, 0x8b, 0xa8, 0xd1, 0xb7, 0x37, 0xaf, 0x4e, 0x27, 0xfb, 0x7b, 0x64, 0x1e, 0x04,
	0x9f, 0x62, 0x47, 0x39, 0x0b, 0x1e, 0xf4, 0x4b, 0x1d, 0x61, 0xfd, 0x45, 0x3d, 0x42, 0x94, 0xc9,
	0xd3, 0xe1, 0x88, 0xe1, 0x3e, 0x65, 0x7a, 0x72, 0xdd, 0x7e, 0x59, 0xfa, 0xed, 0x3b, 0xe2, 0xae,
	0x52, 0x46, 0xff, 0xa6, 0x81, 0xfc, 0x93, 0xc6, 0xb0, 0x6f, 0xd3, 0x97, 0x33, 0xb3, 0xee, 0x66,
	0x7c, 0x96, 0xcc, 0xd1, 0x79, 0x49, 0x38, 0x17, 0x8c, 0xcb, 0xd7, 0x90, 0x5b, 0xc5, 0xcc, 0xc1,
	0xe9, 0xbc, 0xdf, 0xdb, 0xe4, 0xf6, 0xd3, 0xe7, 0xa7, 0x41, 0xd3, 0xdb, 0xf1, 0xff, 0x92, 0xa9,
	0x25, 0x7a, 0x8c, 0x96, 0x4c, 0x78, 0x9f, 0xbe, 0xf4, 0xaf, 0x61, 0xfa, 0x3a, 0xd4, 0x86, 0x79,
	0x2e, 0x16, 0x5d, 0x1d, 0x58, 0x28, 0xd4, 0xfa, 0xdf, 0x01, 0x00, 0x80, 0x53, 0xd9, 0xe4, 0x17,
	0x4a, 0x00, 0x00,
}
