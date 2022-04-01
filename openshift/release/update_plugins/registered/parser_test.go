package registered_test

import (
	"openshift/release/update_plugins/registered"
	"testing"
	"testing/fstest"
	"time"

	"gotest.tools/v3/assert"
)

func TestParser(t *testing.T) {
	t.Run("ListPlugins", testListPlugins)
}

func testListPlugins(t *testing.T) {
	src := `package root

import (
	_ "knative.dev/kn-plugin-event/pkg/plugin"
	_ "knative.dev/kn-plugin-func/plugin"
	_ "knative.dev/kn-plugin-source-kafka/plugin"
)

func RegisterInlinePlugins() {}
`
	parser := registered.Parser{FS: fstest.MapFS{
		"pkg/kn/root/plugin_register.go": {
			Data:    []byte(src),
			Mode:    0,
			ModTime: time.Now(),
		},
	}}
	got := parser.ListPlugins()
	want := []string{
		"knative.dev/kn-plugin-event",
		"knative.dev/kn-plugin-func",
		"knative.dev/kn-plugin-source-kafka",
	}
	assert.DeepEqual(t, got, want)
}
