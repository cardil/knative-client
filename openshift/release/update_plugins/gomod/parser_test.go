package gomod_test

import (
	"openshift/release/update_plugins/gomod"
	"testing"
	"testing/fstest"
	"time"

	"gotest.tools/v3/assert"
)

func TestParser(t *testing.T) {
	t.Run("ResolveInfo", testResolveInfo)
}

func testResolveInfo(t *testing.T) {
	src := `module example

go 1.17

require (
	knative.dev/kn-plugin-event v1.0.1
)

replace (
	knative.dev/kn-plugin-event => github.com/openshift-knative/kn-plugin-event v0.27.1-0.20220223114256-af13ecf492aa
)
`
	parser := gomod.Parser{FS: fstest.MapFS{
		"go.mod": {
			Data:    []byte(src),
			Mode:    0,
			ModTime: time.Now(),
		},
	}}
	info, err := parser.ResolveInfo("knative.dev/kn-plugin-event")
	assert.NilError(t, err)
	assert.Equal(t, info.DisplayAs, "v1.0.1")
	assert.Equal(t, info.Version, "v0.27.1-0.20220223114256-af13ecf492aa")
	assert.Equal(t, info.Source, "github.com/openshift-knative/kn-plugin-event")
}
