package main

import "github.com/gpf-org/gpf/cmd"

func main() {
	cmd.RootCmd.AddCommand(cmd.ServerCmd)
	cmd.RootCmd.AddCommand(cmd.BrahnchesCmd)
	cmd.RootCmd.AddCommand(cmd.ReloadCmd)
	cmd.RootCmd.AddCommand(cmd.ListCmd)
	cmd.RootCmd.AddCommand(cmd.VersionCmd)
	cmd.RootCmd.Execute()
}
