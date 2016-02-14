package msbuild

// Options holds information for the msbuild process
type Options struct {
	WorkingDirectory string `mapstructure:"WD"`
	OutputDirectory  string `mapstructure:"Output_Dir"`
	TestsDirectory   string `mapstructure:"Tests_Dir"`
	TempDirectory    string `mapstructure:"Temp_Dir"`
	Configuration    string
	Platform         string
	Is32Bit          bool
	ToolsVersion     string
	Path             string
	Targets          string
}
