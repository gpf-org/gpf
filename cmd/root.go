package cmd

import (
	"errors"

	"github.com/spf13/cobra"
)

var RootCmd = &cobra.Command{Use: "gpf"}

func init() {
	RootCmd.AddCommand(ServerCmd)
	RootCmd.AddCommand(ReloadCmd)
	RootCmd.AddCommand(ListCmd)
	RootCmd.AddCommand(CodeReviewCmd)
}
