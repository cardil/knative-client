package cli

import "github.com/spf13/cobra"

func update() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "update",
		Short: "Update plugins",
	}

	return cmd
}
