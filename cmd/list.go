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

var ListCmd = &cobra.Command{
	Use:               "list",
	Short:             "List issues",
	PersistentPreRunE: clientPersistentPreRunE,
	RunE: func(cmd *cobra.Command, args []string) error {
		res, err := http.Get(clientFlags.serviceURL + "/issues")
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
	addClientFlags(ListCmd)
}
