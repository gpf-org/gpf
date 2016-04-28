package cmd

import (
	"fmt"
	"log"
	"os"

	"github.com/spf13/cobra"

	"github.com/gpf-org/gpf/git"
)

var ReloadCmd = &cobra.Command{
	Use:   "reload",
	Short: "reconfigure all projects configuration",
	Run: func(cmd *cobra.Command, args []string) {
		gp, err := git.NewProvider(BaseURL, Token, Provider)
		if err != nil {
			fmt.Printf("%s\n", err)
			os.Exit(1)
		}

		var projs []*git.Project
		projs, err = gp.ListAllProjects()
		if err != nil {
			fmt.Printf("%s\n", err)
			os.Exit(1)
		}
		log.Printf("Projects available: %d", len(projs))

		for _, proj := range projs {
			log.Printf("Project %s: reloading webhook", *proj.Name)
			gp.CreateOrUpdateProjectHook(*proj.ID)

			log.Printf("Project %s: reloading information", *proj.Name)
			// TODO: save project information into the gpf database
		}
	},
}
