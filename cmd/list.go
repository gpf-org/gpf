package cmd

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"

	"github.com/spf13/cobra"

	"github.com/gpf-org/gpf/core"
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
			fmt.Printf("unable to retrieve list of issues")
			os.Exit(1)
		}

		issues := make([]core.Issue, 0)
		if err := json.Unmarshal(body, &issues); err != nil {
			fmt.Printf("unable to decode list of issues")
			os.Exit(1)
		}

		for _, issue := range issues {
			fmt.Printf("* issue: %s\n", issue.Name)

			fmt.Printf("\tprojects:\n")
			for _, issueBranch := range issue.IssueBranches {
				fmt.Printf("\t\tproject: %s - branch: %s\n", issueBranch.ProjectName, issueBranch.BranchName)
			}

			fmt.Printf("\tcommands:\n")
			for _, command := range issue.Commands {
				fmt.Printf("\t\tcommand: %s\n", core.CommandText(command))
			}
		}
	},
}

func init() {
	ListCmd.PersistentFlags().StringVarP(&listFlags.publicURL, "publicURL", "", "http://localhost:5544", "Public URL")
}
