package plugins

import (
	"testing"
)

func TestFactPlugin(t *testing.T) {
	plugin, err := SetupFactPlugin("test")
	if err != nil {
		t.Fatalf("SetupFactPlugin: %v", err)
	}

	if plugin == nil {
		t.Fatalf("Plugin is nil")
	}
}
