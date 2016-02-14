package msbuild

import "path/filepath"

// BuildProject holds information about a sln or csproj file
type BuildProject struct {
	Path string
}

// NewBuildProjects returns all BuildProjects based on a wildcard
func NewBuildProjects(workingDir, pattern string) ([]BuildProject, error) {
	// get all the solutions or projects
	projectPath := filepath.Join(workingDir, pattern)
	projectFiles, err := filepath.Glob(projectPath)
	if err != nil {
		return nil, err
	}

	projects := make([]BuildProject, len(projectFiles))
	for i, f := range projectFiles {
		projects[i] = BuildProject{
			Path: f,
		}
	}

	return projects, nil
}
