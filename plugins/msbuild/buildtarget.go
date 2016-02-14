package msbuild

import "strings"

// BuildTarget holds information for a msbuild target
type BuildTarget struct {
	Target    string
	PostBuild bool
}

func contains(src, val string) bool {
	src, val = strings.ToLower(src), strings.ToLower(val)
	return strings.Contains(src, val)
}

// NewBuildTargets takes a comma seperated string and converts it to a list of build targets
func NewBuildTargets(version, targets string) []BuildTarget {
	// trim spaces
	target := strings.Replace(targets, " ", "", -1)

	// set up the default targets
	if target == "" {
		target = "Clean,Build"
	}

	// if we are using msbuild tools 14 we can send them through comma seperated.
	if version == "14.0" {
		return []BuildTarget{
			BuildTarget{
				Target:    target,
				PostBuild: contains(target, "build"),
			},
		}
	}

	buildtargets := []BuildTarget{}
	for _, t := range strings.Split(target, ",") {
		buildtargets = append(buildtargets, BuildTarget{
			Target:    t,
			PostBuild: contains(t, "build"),
		})
	}

	return buildtargets
}
