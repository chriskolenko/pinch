package plugins

import (
	"testing"
)

func TestServicePlugin(t *testing.T) {
	plugin, err := LoadServicePlugin("test")
	if err != nil {
		t.Fatalf("SetupServicePlugin: %v", err)
	}

	if plugin == nil {
		t.Fatalf("Plugin is nil")
	}
}
