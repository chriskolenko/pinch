package models

// PluginType the different plugin types supported by pinch
type PluginType int

const (
	// FactPluginType self described
	FactPluginType PluginType = iota
	// ServicePluginType self described
	ServicePluginType
	// SetupPluginType self described
	SetupPluginType
	// BuildPluginType self described
	BuildPluginType
	// TestPluginType self described
	TestPluginType
	// PostPluginType self described
	PostPluginType
)
