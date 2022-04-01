package main

import "openshift/release/update_plugins/cli"

// Suppress global check for testing purposes.
var cmd = &cli.Command{} //nolint:gochecknoglobals

func main() {
	cmd.Execute()
}
