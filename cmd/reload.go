package cmd

import (
	"fmt"
	"log"
	"os"

	"github.com/spf13/cobra"

	"github.com/gpf-org/gpf/gpf"
)

var ReloadCmd = &cobra.Command{
	Use:   "reload",
	Short: "reconfigure all projects configuration",
	Run: func(cmd *cobra.Command, args []string) {
		projs, err := gpf.ListAllProjects(BaseURL, Token)
		if err != nil {
			fmt.Printf("%s\n", err)
			os.Exit(1)
		}

		log.Printf(">>>%v", projs)

		for _, proj := range projs {
			gpf.CreateOrUpdateProjectHook(BaseURL, Token, *proj.ID)
			log.Printf(">>> Reloading webhook for project %s", *proj.Name)
		}
	},
}
