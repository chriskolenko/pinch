package plugins

import (
	"fmt"
	"sync"

	"github.com/webcanvas/pinch/shared/models"
)

var (
	pluginMu sync.Mutex
	plugins  = make(map[string]Plugin)
)

// Plugin default interface for a pinch plugin
type Plugin interface {
	Setup(models.PluginType, models.Raw) (interface{}, error)
}

// RegisterPlugin allows external packages to register a pinch plugin
func RegisterPlugin(name string, plugin Plugin) {
	pluginMu.Lock()
	defer pluginMu.Unlock()

	if plugin == nil {
		panic("RegisterFactPlugin: plugin is nil")
	}

	if _, dup := plugins[name]; dup {
		panic("RegisterPlugin: Duplicate plugin for " + name)
	}

	plugins[name] = plugin
}

// GetPlugin finds a plugin and returns it
func GetPlugin(name string) (Plugin, error) {
	// get the plugin
	pluginMu.Lock()
	plugin, ok := plugins[name]
	pluginMu.Unlock()

	if !ok {
		return nil, fmt.Errorf("LoadPlugin: unknown plugin %q (forgotten import?)", name)
	}

	return plugin, nil
}
