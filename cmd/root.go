package cmd

import (
	"errors"

	"github.com/spf13/cobra"
)

var (
	Provider string
	Token    string
	BaseURL  string
)

var RootCmd = &cobra.Command{
	Use: "gpf",
	PreRunE: func(cmd *cobra.Command, args []string) error {
		if Token == "" {
			return errors.New("missing required token flag")
		}
		return nil
	},
}

func init() {
	RootCmd.PersistentFlags().StringVarP(&Provider, "provider", "p", "", "Git provider (e.g. gitlab)")
	RootCmd.PersistentFlags().StringVarP(&Token, "token", "t", "", "Private token")
	RootCmd.PersistentFlags().StringVarP(&BaseURL, "baseURL", "b", "https://gitlab.com/api/v3", "Base URL")
	RootCmd.AddCommand(ServerCmd)
	RootCmd.AddCommand(ReloadCmd)
	RootCmd.AddCommand(ListCmd)
	RootCmd.AddCommand(CodeReviewCmd)
}
