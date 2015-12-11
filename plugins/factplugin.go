package plugins

import (
	"fmt"
	"sync"
)

var (
	factPluginMu sync.Mutex
	factPlugins  = make(map[string]FactPlugin)
)

// FactPlugin default interface for a pinch fact plugin
type FactPlugin interface{}

// RegisterFactPlugin allows external packages to register a pinch fact plugin
func RegisterFactPlugin(name string, plugin FactPlugin) {
	factPluginMu.Lock()
	defer factPluginMu.Unlock()

	if plugin == nil {
		panic("RegisterFactPlugin: plugin is nil")
	}

	if _, dup := factPlugins[name]; dup {
		panic("RegisterFactPlugin: Duplicate plugin for " + name)
	}

	factPlugins[name] = plugin
}

// SetupFactPlugin finds a fact plugin and returns it
func SetupFactPlugin(name string) (FactPlugin, error) {
	// get the plugin
	factPluginMu.Lock()
	plugin, ok := factPlugins[name]
	factPluginMu.Unlock()

	if !ok {
		return nil, fmt.Errorf("SetupFactPlugin: unknown plugin %q (forgotten import?)", name)
	}

	return plugin, nil
}
