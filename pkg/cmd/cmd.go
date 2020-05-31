// Copyright 2020 The ccictl Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package cmd

import (
	"context"

	"github.com/spf13/cobra"

	circleci_v1 "github.com/zchee/ccictl/pkg/circleci/v1"
)

// NewCommand returns the new command.
func NewCommand(ctx context.Context, args []string) *cobra.Command {
	c := &cobra.Command{
		Use:   "ccictl",
		Short: "CircleCI controller.",
	}

	svc := circleci_v1.NewService(ctx, nil)
	c.AddCommand(NewMeCmd(svc, args...))

	return c
}
