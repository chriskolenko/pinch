package plugins

import (
	"testing"
)

func TestPinchPlugin(t *testing.T) {
	plugin, err := SetupPinchPlugin("test")
	if err != nil {
		t.Fatalf("SetupPinchPlugin: %v", err)
	}

	if plugin == nil {
		t.Fatalf("Plugin is nil")
	}
}
