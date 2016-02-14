package docker

import "github.com/webcanvas/pinch/modules"

type Module struct {
}

func (Module) GetVersions() []string {
	return nil
}

func init() {
	modules.RegisterModule("docker", &Module{})
}
