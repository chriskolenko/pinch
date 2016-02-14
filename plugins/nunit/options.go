package nunit

// Options provided from the pinch file, or environment
type Options struct {
	WorkingDirectory string `mapstructure:"WD"`
	TestsDirectory   string `mapstructure:"Tests_Dir"`
}
