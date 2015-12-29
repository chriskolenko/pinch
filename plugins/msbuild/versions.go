package msbuild

import "errors"

var errNotFound = errors.New("MSBuild not found")

// Versions is an array of build versions
type Versions []Version

// GetVersion finds a version by tools version. If empty return the first one.
func (versions Versions) GetVersion(version string) (Version, error) {
	// get it.
	var ver = versions[0]

	if version != "" {
		for _, v := range versions {
			if v.Version == version {
				return v, nil
			}
		}

		return ver, errNotFound
	}

	return ver, nil
}
