// Copyright 2020 The ccictl Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"

	circleci_v1 "github.com/zchee/ccictl/pkg/circleci/v1"
)

// NewMeCmd returns the new command.
func NewMeCmd(svc *circleci_v1.Service, args ...string) *cobra.Command {
	cmd := &cobra.Command{
		Use:   "me",
		Short: "information about the signed in user",
	}
	cmd.RunE = func(cmd *cobra.Command, args []string) error {
		user, err := svc.API.Me().Do()
		if err != nil {
			return err
		}
		fmt.Fprintf(os.Stdout, "%#v\n", user)

		return nil
	}

	return cmd
}
