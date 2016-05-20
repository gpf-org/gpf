package cmd

import (
	"errors"
	"log"

	"github.com/gpf-org/gpf/server"
	"github.com/spf13/cobra"
)

var options = &server.ServerOptions{}

var ServerCmd = &cobra.Command{
	Use:   "server",
	Short: "A high performance gpf server",
	PreRunE: func(cmd *cobra.Command, args []string) error {
		if options.Provider == "" {
			return errors.New("missing required provider flag")
		}
		if options.Token == "" {
			return errors.New("missing required token flag")
		}
		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {
		s, err := server.NewServer(options)
		if err != nil {
			log.Fatal("Failed to create server instance: ", err)
		}

		if err := s.Reload(); err != nil {
			log.Fatal("Failed to reload server data: ", err)
		}

		if err := s.ListenAndServe(); err != nil {
			log.Fatal("Failed to start listening: ", err)
		}
	},
}

func init() {
	ServerCmd.PersistentFlags().StringVarP(&options.Provider, "provider", "p", "", "Git provider (e.g. gitlab, github)")
	ServerCmd.PersistentFlags().StringVarP(&options.Token, "token", "t", "", "Private token")
	ServerCmd.PersistentFlags().StringVarP(&options.BaseURL, "baseURL", "", "https://gitlab.com/api/v3", "Base URL")
	ServerCmd.PersistentFlags().StringVarP(&options.PublicURL, "publicURL", "", "", "Public URL")
	ServerCmd.PersistentFlags().StringVarP(&options.Bind, "bind", "", "127.0.0.1", "Interface to which the server will bind")
	ServerCmd.PersistentFlags().IntVarP(&options.Port, "port", "", 5544, "Port on which the server will listen")
}
