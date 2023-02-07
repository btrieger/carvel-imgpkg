// Copyright 2020 VMware, Inc.
// SPDX-License-Identifier: Apache-2.0

package cmd

import (
	"github.com/spf13/cobra"
)

type LockOutputFlags struct {
	LockFilePath string
}

// SetOnCopy Sets the lock-output flag for Copy command
func (l *LockOutputFlags) SetOnCopy(cmd *cobra.Command) {
	cmd.Flags().StringVar(&l.LockFilePath, "lock-output", "",
		"Location to output the generated lockfile. Option only available when using --bundle or --lock flags")
}

// SetOnPush Sets the lock-output flag for Push command
func (l *LockOutputFlags) SetOnPush(cmd *cobra.Command) {
	cmd.Flags().StringVar(&l.LockFilePath, "lock-output", "",
		"Location to output the generated lockfile. Option only available when using --bundle flag")
}
