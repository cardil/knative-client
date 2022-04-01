package cli

import (
	"fmt"
	"openshift/release/update_plugins/gomod"
	"openshift/release/update_plugins/registered"
	"os"
	"path"

	"github.com/spf13/cobra"
)

func list() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "list",
		Short: "List plugins",
		RunE: func(cmd *cobra.Command, args []string) error {
			fs := os.DirFS(path.Dir(cwd()))
			reg := registered.Parser{FS: fs}
			pls := reg.ListPlugins()
			for _, pl := range pls {
				gmp := gomod.Parser{FS: fs}
				pi, err := gmp.ResolveInfo(pl)
				if err != nil {
					return err
				}
				_, err = fmt.Fprintf(cmd.OutOrStdout(),
					" * %s (%s):\n     source: %s\n     version: %s\n\n",
					pl, pi.DisplayAs, pi.Source, pi.Version)
				if err != nil {
					return err
				}
			}
			if len(pls) == 0 {
				_, err := fmt.Fprint(cmd.OutOrStdout(), "No plugins registered")
				if err != nil {
					return err
				}
			}
			return nil
		},
	}

	return cmd
}

func cwd() string {
	wd, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	return wd
}
