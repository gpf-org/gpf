package cmd

import (
	"fmt"
	"log"
	"os"

	"github.com/spf13/cobra"

	"github.com/gpf-org/gpf/git"
)

type BranchesFlags struct {
	status string
}

var branchesFlags = &BranchesFlags{}

var BrahnchesCmd = &cobra.Command{
	Use:   "branches",
	Short: "list feature branches and their statuses",
}

var listBranchesCmd = &cobra.Command{
	Use:   "list",
	Short: "List branches",
	Run: func(cmd *cobra.Command, args []string) {
		gp, err := git.NewProvider(BaseURL, Token, Provider)
		if err != nil {
			fmt.Printf("%s\n", err)
			os.Exit(1)
		}

		// TODO: get pid by argument
		pid := 1
		var result []*git.Branch
		result, err = gp.ListAllBranches(pid)
		if err != nil {
			fmt.Printf("%s\n", err)
			os.Exit(1)
		}

		// util.PrintJSONList(result)
		log.Printf(">>>%v", result)
	},
}

func init() {
	// listBranchesCmd flags
	listBranchesCmd.Flags().StringVarP(&branchesFlags.status, "status", "s", "", "Filter by status")

	// add commands into the parent command
	BrahnchesCmd.AddCommand(listBranchesCmd)
}
