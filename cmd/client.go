package cmd

import (
	"errors"

	"github.com/spf13/cobra"
)

var clientFlags = struct {
	serviceURL string
}{}

func addClientFlags(command *cobra.Command) {
	command.PersistentFlags().StringVarP(&clientFlags.serviceURL, "serviceURL", "", "", "Service URL")
}

func clientPersistentPreRunE(cmd *cobra.Command, args []string) error {
	if clientFlags.serviceURL == "" {
		return errors.New("missing required serviceURL flag")
	}
	return nil
}
