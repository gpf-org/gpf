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

		for _, proj := range projs {
			log.Printf("Project %s: reloading webhook", *proj.Name)
			gpf.CreateOrUpdateProjectHook(BaseURL, Token, *proj.ID)

			log.Printf("Project %s: reloading information", *proj.Name)
			// TODO: save project information into the gpf database
		}
	},
}
