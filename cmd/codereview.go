package cmd

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/spf13/cobra"

	"github.com/gpf-org/gpf/core"
)

var CodeReviewCmd = &cobra.Command{
	Use:               "code-review <action>",
	Short:             "Code Review",
	PersistentPreRunE: clientPersistentPreRunE,
}

var requestCmd = &cobra.Command{
	Use:   "request <issue>",
	Short: "Request code review",
	PreRunE: func(cmd *cobra.Command, args []string) error {
		if len(args) != 1 {
			return errors.New("missing required issue name")
		}
		return nil
	},
	Run: func(cmd *cobra.Command, args []string) {
		url := fmt.Sprintf("%s/issues/%s/command/%d", clientFlags.serviceURL, args[0], core.CommandCodeReviewRequest)
		res, err := http.Post(url, "application/json", bytes.NewBuffer([]byte("")))
		if err != nil {
			log.Fatal(err)
		}

		defer res.Body.Close()

		if res.StatusCode != 200 {
			body, _ := ioutil.ReadAll(res.Body)
			log.Fatalf("Invalid response [%d]: %v", res.StatusCode, string(body))
		}

		body, err := ioutil.ReadAll(res.Body)
		if err != nil {
			log.Fatalf("Unable to retrieve result: %v", err)
		}

		affectedIssueBranches := make([]core.AffectedIssueBranch, 0)
		if err := json.Unmarshal(body, &affectedIssueBranches); err != nil {
			log.Fatalf("Unable to decode list of affected issue branches: %v", err)
		}

		for _, affected := range affectedIssueBranches {
			if affected.Error != nil {
				fmt.Print("[Failed]\t")
			} else {
				fmt.Print("[OK]\t")
			}

			fmt.Printf("Project: %s (%s)\n", affected.IssueBranch.ProjectName, affected.IssueBranch.BranchName)
		}
	},
}

func init() {
	addClientFlags(CodeReviewCmd)
	CodeReviewCmd.AddCommand(requestCmd)
}
