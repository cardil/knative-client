package cli

import "github.com/spf13/cobra"

func root() *cobra.Command {
	cmd := &cobra.Command{
		Long: "Manage plugins of OpenShift Serverless CLI",
	}
	cmd.AddCommand(list())
	cmd.AddCommand(update())
	return cmd
}
