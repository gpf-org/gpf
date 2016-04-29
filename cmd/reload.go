package cmd

import (
	"errors"
	"fmt"
	"net/http"
	"os"

	"github.com/spf13/cobra"
)

type ReloadFlags struct {
	publicURL string
}

var reloadFlags = &ReloadFlags{}

var ReloadCmd = &cobra.Command{
	Use:   "reload",
	Short: "reconfigure all projects configuration",
	PreRunE: func(cmd *cobra.Command, args []string) error {
		if reloadFlags.publicURL == "" {
			return errors.New("missing required publicURL flag")
		}
		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {
		resp, err := http.Get(reloadFlags.publicURL + "/reload")
		if err != nil || resp.StatusCode != 200 {
			fmt.Printf("%s\n", err)
			os.Exit(1)
		}

		fmt.Printf("Finished reloading")
	},
}

func init() {
	ReloadCmd.PersistentFlags().StringVarP(&reloadFlags.publicURL, "publicURL", "", "http://localhost:5544", "Public URL")
}
