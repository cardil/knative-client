package main // nolint:testpackage

import (
	"bytes"
	"openshift/release/update_plugins/cli"
	"strings"
	"testing"

	"gotest.tools/v3/assert"
)

func TestMainFunc(t *testing.T) {
	buf := bytes.NewBuffer([]byte{})
	cmd = &cli.Command{
		Out: buf,
		Exit: func(code int) {
			assert.Equal(t, 0, code)
		},
	}

	main()

	out := buf.String()
	assert.Check(t, strings.Contains(out, "Manage plugins of OpenShift Serverless CLI"))
}
