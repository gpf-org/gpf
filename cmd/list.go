package cmd

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"

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
		res, err := http.Get(reloadFlags.publicURL + "/issues")
	RunE: func(cmd *cobra.Command, args []string) error {
		if err != nil || res.StatusCode != 200 {
			return err
		}

		body, err := ioutil.ReadAll(res.Body)
		if err != nil {
			return errors.New("Unable to retrieve list of issues")
		}

		issues := make([]core.Issue, 0)
		if err := json.Unmarshal(body, &issues); err != nil {
			return errors.New("Unable to decode list of issues")
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

		return nil
	},
}

func init() {
	ListCmd.PersistentFlags().StringVarP(&listFlags.publicURL, "publicURL", "", "http://localhost:5544", "Public URL")
}
