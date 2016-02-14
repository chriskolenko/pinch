package msbuild

import (
	"path/filepath"
	"strings"
)

// BuildConfig holds information for a msbuild configuration
type BuildConfig struct {
	OutputDirectory string
	TempDirectory   string
	TestsDirectory  string
	Configuration   string
	Platform        string
}

// NewBuildConfigs creates build configurations form a comma seperated list of configurations
func NewBuildConfigs(workingDir, outputDir, testsDir, tempDir, configurations, platform string) []BuildConfig {
	// trim spaces
	configurations = strings.Replace(configurations, " ", "", -1)

	// Set the defaults
	if configurations == "" {
		configurations = "Release"
	}

	if platform == "" {
		platform = "Any Cpu"
	}

	if outputDir == "" {
		outputDir = `.build`
	}

	if testsDir == "" {
		testsDir = filepath.Join(outputDir, "Tests")
	}

	if tempDir == "" {
		tempDir = `obj`
	}

	// Convert output directory to full path.
	if !filepath.IsAbs(outputDir) {
		outputDir = filepath.Join(workingDir, outputDir)
	}

	// Convert output directory to full path.
	if !filepath.IsAbs(testsDir) {
		testsDir = filepath.Join(workingDir, testsDir)
	}

	buildconfigs := []BuildConfig{}
	for _, cfg := range strings.Split(configurations, ",") {
		// do it.
		buildconfigs = append(buildconfigs, BuildConfig{
			OutputDirectory: outputDir,
			TestsDirectory:  testsDir,
			TempDirectory:   tempDir,
			Configuration:   cfg,
			Platform:        platform,
		})
	}

	return buildconfigs
}
