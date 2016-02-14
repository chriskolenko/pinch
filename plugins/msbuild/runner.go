package msbuild

import (
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"

	"github.com/webcanvas/pinch/shared/commanders"
	"github.com/webcanvas/pinch/shared/models"
)

// Runner holds information for the Run method
type Runner struct {
	// Configs has the information about the different msbuild configurations and platforms
	Configs []BuildConfig
	// Targets are the different targets passed to msbuild, build and rebuild have post build operations
	Targets []BuildTarget
	// Projects can be full paths to sln or csproj files
	Projects []BuildProject
	// DefaultProjects has a list of csproj names, each csprojs artifacts will be copied to the output directory
	DefaultProjects []string
	// NuGetVersion if octopack is installed this will pass in the nuget package version
	NuGetVersion string

	cmd commanders.Commander
}

func (r Runner) run(project BuildProject, config BuildConfig, target BuildTarget) error {
	tempDir := filepath.Join(config.TempDirectory, "Pinch-"+config.Configuration)
	binDir := filepath.Join(tempDir, "bin")

	args := []string{
		project.Path,
		fmt.Sprintf(`/t:%s`, target.Target),
		fmt.Sprintf(`/p:Configuration=%s`, config.Configuration),
		fmt.Sprintf(`/p:Platform=%s`, config.Platform),
		fmt.Sprintf(`/p:WebProjectOutputDir=%s`, binDir),
		fmt.Sprintf(`/p:OutDir=%s`, tempDir),
		fmt.Sprintf(`/p:OctoPackPublishPackageToFileShare=%s`, config.OutputDirectory),
		fmt.Sprintf(`/p:OctoPackPackageVersion=%s`, r.NuGetVersion),
		`/p:RunOctoPack=true`,
		`/m`,
		`/v:quiet`,
		`/nologo`,
	}

	output, err := r.cmd.ExecOutput(args...)
	if err != nil {
		fmt.Print(string(output))
		return err
	}

	if !target.PostBuild {
		return nil
	}

	// find all the test directories.
	dir := filepath.Dir(project.Path)
	testProjects, err := getTestProjects(dir)
	if err != nil {
		return err
	}

	// create the path.
	err = os.RemoveAll(config.TestsDirectory)
	if err != nil {
		return err
	}

	// # Artifact default project
	// $project_build_output_dir = $(Join-Path "$default_project_dir" "$output_dir_name")
	// Copy-Item "$project_build_output_dir" "$artifact_dir" -Recurse -Force

	// create the test directory.
	err = ensureDirectory(config.TestsDirectory)
	if err != nil {
		return err
	}

	for _, testProject := range testProjects {
		fmt.Println("Test", testProject)

		info, err := os.Stat(testProject)
		if err != nil {
			return err
		}

		filename := info.Name()
		extension := filepath.Ext(filename)
		name := filename[0 : len(filename)-len(extension)]

		testOutput := filepath.Join(config.TestsDirectory, name)

		testDir := filepath.Join(filepath.Dir(testProject), tempDir)
		err = copyDir(testDir, testOutput)
		if err != nil {
			return err
		}
	}

	// # Artifact default project
	// $project_build_output_dir = $(Join-Path "$default_project_dir" "$output_dir_name")
	// Copy-Item "$project_build_output_dir" "$artifact_dir" -Recurse -Force

	// # Artifact tests
	// 	$test_projects = Get-TestProjectsFromSolution
	// foreach ($test_project in $test_projects) {
	// 	$test_project_artifact_dir = $( Join-Path $(Join-Path "$test_artifacts_dir" "$($test_project.Type)") "$($test_project.Name)")
	// 	$test_build_output_dir = $(Join-Path "$($test_project.Directory)" "$output_dir_name")
	// 	Copy-Item "$test_build_output_dir" "$test_project_artifact_dir" -Recurse -Force
	// }

	// write the output.
	return nil
}

// Run is the way we run things around here.
func (r *Runner) Run() (models.Result, error) {
	result := models.Result{}

	for _, project := range r.Projects {
		// maybe todo get information about the msbuild solution!.
		for _, buildconfig := range r.Configs {
			// run msbuild!.
			for _, target := range r.Targets {
				// this doesn't take in account the result yet.
				err := r.run(project, buildconfig, target)
				if err != nil {
					return result, err
				}
			}
		}
	}

	return result, nil
}

// CreateRunner creates the object which will run the msbuild operations
func CreateRunner(versions Versions, opts Options) (*Runner, error) {
	// get the tools version
	version, err := versions.GetVersion(opts.ToolsVersion)
	if err != nil {
		return nil, err
	}

	path, err := version.GetPath(opts.Is32Bit)
	if err != nil {
		return nil, err
	}

	// get the msbuild.
	cmd, err := commanders.Open("msbuild.exe", path)
	if err != nil {
		return nil, err
	}

	// create the targets
	targets := NewBuildTargets(version.Version, opts.Targets)

	// create the configs
	configs := NewBuildConfigs(opts.WorkingDirectory, opts.OutputDirectory, opts.TestsDirectory, opts.TempDirectory, opts.Configuration, opts.Platform)

	// finds all projects
	projects, err := NewBuildProjects(opts.WorkingDirectory, opts.Path)
	if err != nil {
		return nil, err
	}

	defaultProject := []string{"something"}

	return &Runner{
		Configs:         configs,
		Targets:         targets,
		Projects:        projects,
		DefaultProjects: defaultProject,

		cmd: *cmd,
	}, nil
}

func ensureDirectory(dir string) error {
	if _, err := os.Stat(dir); os.IsNotExist(err) {
		// path/to/whatever does not exist
		err := os.MkdirAll(dir, 0777)
		if err != nil {
			return err
		}
	}

	return nil
}

func getTestProjects(dir string) ([]string, error) {
	testProjects := []string{}

	walker := func(fp string, fi os.FileInfo, err error) error {
		if err != nil {
			return nil
		}
		if !!fi.IsDir() {
			return nil // not a file
		}
		matched, err := filepath.Match("*Tests.csproj", fi.Name())
		if err != nil {
			return err // this is bad.
		}
		if matched {
			testProjects = append(testProjects, fp)
		}
		return nil
	}

	err := filepath.Walk(dir, walker)
	if err != nil {
		return nil, err
	}

	return testProjects, nil
}

// Recursively copies a directory tree, attempting to preserve permissions.
// Source directory must exist, destination directory must *not* exist.
func copyDir(source string, dest string) (err error) {

	// get properties of source dir
	fi, err := os.Stat(source)
	if err != nil {
		return err
	}

	if !fi.IsDir() {
		return fmt.Errorf("Source is not a directory")
	}

	// ensure dest dir does not already exist
	_, err = os.Open(dest)
	if !os.IsNotExist(err) {
		return fmt.Errorf("Destination already exists")
	}

	// create dest dir
	err = os.MkdirAll(dest, fi.Mode())
	if err != nil {
		return err
	}

	entries, err := ioutil.ReadDir(source)

	for _, entry := range entries {
		sfp := source + "/" + entry.Name()
		dfp := dest + "/" + entry.Name()
		if entry.IsDir() {
			err = copyDir(sfp, dfp)
			if err != nil {
				log.Println(err)
			}
		} else {
			// perform copy
			err = copyFile(sfp, dfp)
			if err != nil {
				log.Println(err)
			}
		}

	}
	return
}

// Copies file source to destination dest.
func copyFile(source string, dest string) (err error) {
	sf, err := os.Open(source)
	if err != nil {
		return err
	}
	defer sf.Close()
	df, err := os.Create(dest)
	if err != nil {
		return err
	}
	defer df.Close()
	_, err = io.Copy(df, sf)
	if err == nil {
		si, err := os.Stat(source)
		if err != nil {
			err = os.Chmod(dest, si.Mode())
		}
	}

	return
}
