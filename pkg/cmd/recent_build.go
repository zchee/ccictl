// Copyright 2020 The ccictl Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package cmd

import (
	"fmt"

	"github.com/davecgh/go-spew/spew"
	"github.com/spf13/cobra"

	circleci_v1 "github.com/zchee/ccictl/pkg/circleci/v1"
)

// NewRecentBuildCmd returns the new command.
func NewRecentBuildCmd(svc *circleci_v1.Service, args ...string) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "recent-build",
		Short: "Build summary for recent builds",
	}
	cmd.RunE = func(cmd *cobra.Command, args []string) error {
		recents, err := svc.API.RecentBuilds().Shallow(true).Do()
		if err != nil {
			return err
		}
		fmt.Printf("recents: %s", spew.Sdump(recents))

		return nil
	}

	return cmd
}
