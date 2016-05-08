package cmd

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"

	"github.com/gpf-org/gpf/server"
	"github.com/spf13/cobra"
)

type ListFlags struct {
	publicURL string
}

var listFlags = &ListFlags{}

var ListCmd = &cobra.Command{
	Use:   "list",
	Short: "list features",
	PreRunE: func(cmd *cobra.Command, args []string) error {
		if listFlags.publicURL == "" {
			return errors.New("missing required publicURL flag")
		}
		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {
		res, err := http.Get(reloadFlags.publicURL + "/list")
		if err != nil || res.StatusCode != 200 {
			fmt.Printf("%s\n", err)
			os.Exit(1)
		}

		body, err := ioutil.ReadAll(res.Body)
		if err != nil {
			fmt.Printf("unable to retrieve list of features")
			os.Exit(1)
		}

		features := make([]server.Feature, 0)
		if err := json.Unmarshal(body, &features); err != nil {
			fmt.Printf("unable to decode list of features")
			os.Exit(1)
		}

		fmt.Printf("Found %d features\n", len(features))
		for _, feature := range features {
			fmt.Printf("* feature: %s\n", feature.Name)

			fmt.Printf("\tprojects:\n")
			for _, branch := range feature.Branches {
				fmt.Printf("\t\tproject: %s - branch: %s\n", branch.ProjectName, branch.BranchName)
			}

			fmt.Printf("\tcommands:\n")
			for _, command := range feature.Commands {
				fmt.Printf("\t\tcommand: %s\n", command)
			}
		}
	},
}

func init() {
	ListCmd.PersistentFlags().StringVarP(&listFlags.publicURL, "publicURL", "", "http://localhost:5544", "Public URL")
}
