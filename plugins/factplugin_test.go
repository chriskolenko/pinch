package plugins

import (
	"testing"
)

func TestFactPlugin(t *testing.T) {
	plugin, err := LoadFactPlugin("test")
	if err != nil {
		t.Fatalf("LoadFactPlugin: %v", err)
	}

	if plugin == nil {
		t.Fatalf("Plugin is nil")
	}
}
