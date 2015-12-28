package plugins

import (
	"testing"
)

func TestPlugin(t *testing.T) {
	plugin, err := LoadPlugin("test")
	if err != nil {
		t.Fatalf("LoadPlugin: %v", err)
	}

	if plugin == nil {
		t.Fatalf("Plugin is nil")
	}
}
