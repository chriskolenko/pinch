package capabilities

import "github.com/webcanvas/pinch/modules"

type Capabilities []Capability

type Capability struct {
	Name    string
	Version string
}

func Load() (Capabilities, error) {
	// for each of the module find something that has the capablitiy interface.
	// test to see if the system supports it.
	var caps []Capability

	// get all the modules.
	mods := modules.GetModules()
	for name, mod := range mods {
		// A capability is something we have installed on a machine.
		// There could be multiple versions of the tool too.
		vers := mod.GetVersions()
		for _, ver := range vers {
			caps = append(caps, Capability{
				Name:    name,
				Version: ver,
			})
		}
	}

	return caps, nil
}
