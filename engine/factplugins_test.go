package engine

import (
	"testing"
)

func TestFactPlugin(t *testing.T) {
	factplugin, err := SetupFactPlugin("test")
	if err != nil {
		t.Fatalf("SetupFactPlugin: %v", err)
	}

	if factplugin == nil {
		t.Fatalf("Fact Plugin is nil")
	}
}
