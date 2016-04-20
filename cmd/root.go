package cmd

import (
	"errors"

	"github.com/spf13/cobra"
)

var (
	Token string
)

var RootCmd = &cobra.Command{
	PreRunE: func(cmd *cobra.Command, args []string) error {
		if Token == "" {
			return errors.New("missing required token flag")
		}
		return nil
	},
}

func init() {
	RootCmd.PersistentFlags().StringVarP(&Token, "token", "", "", "Private token")
}
