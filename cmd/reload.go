package cmd

import (
	"fmt"
	"log"
	"net/http"

	"github.com/spf13/cobra"
)

var ReloadCmd = &cobra.Command{
	Use:               "reload",
	Short:             "Force server to reload configuration",
	PersistentPreRunE: clientPersistentPreRunE,
	Run: func(cmd *cobra.Command, args []string) {
		resp, err := http.Get(clientFlags.serviceURL + "/reload")
		if err != nil || resp.StatusCode != 200 {
			log.Fatal(err)
		}

		fmt.Printf("Server configuration reloaded with success")
	},
}

func init() {
	addClientFlags(ReloadCmd)
}
