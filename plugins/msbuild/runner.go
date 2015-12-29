package msbuild

import (
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/webcanvas/pinch/shared/commanders"
	"github.com/webcanvas/pinch/shared/models"
)

// Runner holds information for the Run method
type Runner struct {
	OutputDirectory string
	Configuration   string
	Targets         []string
	Platform        string
	ProjectFiles    []string

	cmd commanders.Commander
}

func (r Runner) ensureDirectories() error {
	return ensureDirectory(r.OutputDirectory)
}

func (r Runner) runMSBuild(projectFile, target string) error {
	// what should the output be?
	outputDir := `obj\pinch-` + r.Configuration

	// make the args
	args := []string{
		projectFile,
		"/t:" + target,
		"/p:Configuration=" + r.Configuration,
		"/p:Platform=" + r.Platform,
		"/p:WebProjectOutputDir=" + outputDir,
		"/p:OutDir=" + outputDir,
		// "/p:OutputPath=" + outputDir,
		"/m",
		"/v:quiet",
		"/nologo",
	}

	output, err := r.cmd.ExecOutput(args...)
	if err != nil {
		return err
	}

	// find all the test directories.
	dir := filepath.Dir(projectFile)
	testProjects, err := getTestProjects(dir)
	if err != nil {
		return err
	}

	// create the path.
	allTestsOutput := filepath.Join(r.OutputDirectory, "Tests")
	err = os.RemoveAll(allTestsOutput)
	if err != nil {
		return err
	}

	// create the test directory.
	err = ensureDirectory(allTestsOutput)
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

		testOutput := filepath.Join(allTestsOutput, name)

		testDir := filepath.Join(filepath.Dir(testProject), outputDir)
		err = copyDir(testDir, testOutput)
		if err != nil {
			return err
		}
	}

	// # Artifact tests
	// 	$test_projects = Get-TestProjectsFromSolution
	// foreach ($test_project in $test_projects) {
	// 	$test_project_artifact_dir = $( Join-Path $(Join-Path "$test_artifacts_dir" "$($test_project.Type)") "$($test_project.Name)")
	// 	$test_build_output_dir = $(Join-Path "$($test_project.Directory)" "$output_dir_name")
	// 	Copy-Item "$test_build_output_dir" "$test_project_artifact_dir" -Recurse -Force
	// }

	// write the output.
	fmt.Print(string(output))
	return nil
}

// Run is the way we run things around here.
func (r *Runner) Run() (models.Result, error) {
	result := models.Result{}

	err := r.ensureDirectories()
	if err != nil {
		return result, err
	}

	// /p:RunOctoPack=$run_octopack /p:OctoPackPublishPackageToFileShare="$build_output_dir" /p:OctoPackPackageVersion=$version /verbosity:quiet }
	// /p:WebProjectOutputDir="$output_dir_name" /p:OutDir="$(Join-Path "$output_dir_name" "bin\")" /p:RunOctoPack=$run_octopack /p:OctoPackPublishPackageToFileShare="$build_output_dir" /p:OctoPackPackageVersion="$nuget_version" /verbosity:quiet }

	for _, project := range r.ProjectFiles {

		// get information about the msbuild solution!.

		// run msbuild!.
		for _, target := range r.Targets {
			// this doesn't take in account the result yet.
			err := r.runMSBuild(project, target)
			if err != nil {
				return result, err
			}
		}
	}

	return result, nil
}

// CreateRunner creates the object which will run the msbuild operations
func CreateRunner(versions Versions, opts Options) (*Runner, error) {
	// try to find the msbuild !.

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

	// get the targets.
	target := strings.Replace(opts.Targets, " ", "", -1)

	var targets []string
	if target == "" {
		targets = []string{"Clean", "Build"}
	} else if version.Version == "14.0" {
		targets = []string{target}
	} else {
		targets = strings.Split(target, ",")
	}

	// get the configuration.
	configuration := opts.Configuration
	if configuration == "" {
		configuration = "Release"
	}

	platform := opts.Platform
	if platform == "" {
		platform = "Any Cpu"
	}

	// what about the output directory
	outputDir := opts.OutputDirectory
	if outputDir == "" {
		outputDir = `.build\output`
	}

	// Convert output directory to full path.
	if !filepath.IsAbs(outputDir) {
		outputDir = filepath.Join(opts.WorkingDirectory, outputDir)
	}

	// get all the solutions or projects
	projectPath := filepath.Join(opts.WorkingDirectory, opts.Path)
	projectFiles, err := filepath.Glob(projectPath)
	if err != nil {
		return nil, err
	}

	return &Runner{
		OutputDirectory: outputDir,
		Configuration:   configuration,
		Targets:         targets,
		Platform:        platform,
		ProjectFiles:    projectFiles,

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
