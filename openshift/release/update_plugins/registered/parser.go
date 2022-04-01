package registered

import (
	"go/parser"
	"go/token"
	"io/fs"
	"path"
	"strconv"
	"strings"
)

type Parser struct {
	fs.FS
}

func (p Parser) ListPlugins() []string {
	filename := "pkg/kn/root/plugin_register.go"
	srcBytes, err := fs.ReadFile(p.FS, filename)
	if err != nil {
		return []string{}
	}
	fset := token.NewFileSet() // positions are relative to fset
	// Parse src but stop after processing the imports.
	src := string(srcBytes)
	f, err := parser.ParseFile(fset, filename, src, parser.ImportsOnly)
	if err != nil {
		return []string{}
	}
	plugins := make([]string, 0, len(f.Imports))
	for _, imp := range f.Imports {
		if imp.Name.String() != "_" {
			continue
		}
		lit := imp.Path
		if importPath, err := strconv.Unquote(lit.Value); err == nil {
			parts := strings.Split(importPath, "/")
			if !strings.HasPrefix(parts[1], "kn-plugin") {
				continue
			}
			pluginName := path.Join(parts[0:2]...)
			plugins = append(plugins, pluginName)
		}
	}

	return plugins
}
