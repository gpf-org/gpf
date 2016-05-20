package cmd

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/spf13/cobra"

	"github.com/gpf-org/gpf/core"
)

var ListCmd = &cobra.Command{
	Use:               "list",
	Short:             "List issues",
	PersistentPreRunE: clientPersistentPreRunE,
	Run: func(cmd *cobra.Command, args []string) {
		res, err := http.Get(clientFlags.serviceURL + "/issues")
		if err != nil || res.StatusCode != 200 {
			log.Fatal(err)
		}

		body, err := ioutil.ReadAll(res.Body)
		if err != nil {
			log.Fatal("Unable to retrieve list of issues")
		}

		issues := make([]core.Issue, 0)
		if err := json.Unmarshal(body, &issues); err != nil {
			log.Fatal("Unable to decode list of issues")
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
	addClientFlags(ListCmd)
}
