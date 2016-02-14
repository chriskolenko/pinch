package msbuild

import "github.com/webcanvas/pinch/modules"

type Module struct {
	versions Versions
}

func (m Module) GetVersions() []string {
	vers := make([]string, len(m.versions))

	for i, v := range m.versions {
		vers[i] = v.Version
	}

	return vers
}

func init() {
	// we can search for it here.
	versions, err := FindVersions()
	if err != nil {
		panic(err)
	}

	module := &Module{
		versions: versions,
	}

	modules.RegisterModule("msbuild", module)
}
