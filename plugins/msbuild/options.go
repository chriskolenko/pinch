package msbuild

// Options holds information for the msbuild process
type Options struct {
	WorkingDirectory string `mapstructure:"WD"`
	OutputDirectory  string `mapstructure:"Output_Dir"`
	Configuration    string
	Platform         string
	Is32Bit          bool
	ToolsVersion     string
	Path             string
	Targets          string
}
