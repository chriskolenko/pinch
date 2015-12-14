package plugins

import (
	"fmt"
	"sync"

	"github.com/webcanvas/pinch/shared/models"
)

var (
	pinchPluginMu sync.Mutex
	pinchPlugins  = make(map[string]PinchPlugin)
)

// PinchPlugin default interface for a pinch pinch plugin
type PinchPlugin interface {
	Setup() error
	Exec(map[string]string) (models.Result, error)
}

// RegisterPinchPlugin allows external packages to register a pinch pinch plugin
func RegisterPinchPlugin(name string, plugin PinchPlugin) {
	pinchPluginMu.Lock()
	defer pinchPluginMu.Unlock()

	if plugin == nil {
		panic("RegisterPinchPlugin: plugin is nil")
	}

	if _, dup := pinchPlugins[name]; dup {
		panic("RegisterPinchPlugin: Duplicate plugin for " + name)
	}

	pinchPlugins[name] = plugin
}

// SetupPinchPlugin finds a pinch plugin and returns it
func SetupPinchPlugin(name string) (PinchPlugin, error) {
	// get the plugin
	pinchPluginMu.Lock()
	plugin, ok := pinchPlugins[name]
	pinchPluginMu.Unlock()

	if !ok {
		return nil, fmt.Errorf("SetupPinchPlugin: unknown plugin %q (forgotten import?)", name)
	}

	err := plugin.Setup()
	return plugin, err
}
