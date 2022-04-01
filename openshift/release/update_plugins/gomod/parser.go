package gomod

import (
	"io/fs"

	"golang.org/x/mod/modfile"
)

type Parser struct {
	fs.FS
}
type PluginInfo struct {
	Version   string
	Source    string
	DisplayAs string
}

func (p Parser) ResolveInfo(plugin string) (*PluginInfo, error) {
	filename := "go.mod"
	bytes, err := fs.ReadFile(p.FS, filename)
	if err != nil {
		return nil, err
	}
	modf, err := modfile.Parse(filename, bytes, nil)
	if err != nil {
		return nil, err
	}
	pi := &PluginInfo{Source: plugin}
	for _, req := range modf.Require {
		if req.Mod.Path == plugin {
			pi.Version = req.Mod.Version
			pi.DisplayAs = pi.Version
		}
	}
	for _, rep := range modf.Replace {
		if rep.Old.Path == plugin {
			pi.Version = rep.New.Version
			pi.Source = rep.New.Path
		}
	}
	return pi, nil
}
