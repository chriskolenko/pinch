package plugins

import (
	"fmt"
	"sync"
)

var (
	servicePluginMu sync.Mutex
	servicePlugins  = make(map[string]ServicePlugin)
)

// ServicePlugin default interface for a pinch service plugin
type ServicePlugin interface{}

// RegisterServicePlugin allows external packages to register a pinch service plugin
func RegisterServicePlugin(name string, plugin ServicePlugin) {
	servicePluginMu.Lock()
	defer servicePluginMu.Unlock()

	if plugin == nil {
		panic("RegisterServicePlugin: plugin is nil")
	}

	if _, dup := servicePlugins[name]; dup {
		panic("RegisterServicePlugin: Duplicate plugin for " + name)
	}

	servicePlugins[name] = plugin
}

// SetupServicePlugin finds a service plugin and returns it
func SetupServicePlugin(name string) (ServicePlugin, error) {
	// get the plugin
	servicePluginMu.Lock()
	plugin, ok := servicePlugins[name]
	servicePluginMu.Unlock()

	if !ok {
		return nil, fmt.Errorf("SetupServicePlugin: unknown plugin %q (forgotten import?)", name)
	}

	return plugin, nil
}
