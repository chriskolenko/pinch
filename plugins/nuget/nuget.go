package nuget

import (
	"bufio"
	"bytes"
	"regexp"
	"strings"

	"github.com/webcanvas/pinch/plugins"
	"github.com/webcanvas/pinch/shared/commanders"
	"github.com/webcanvas/pinch/shared/models"
)

var versionex = regexp.MustCompile("[0-9.]+")

// NuGet this is exported just for fun. This will be used inside the pinch engine
type NuGet struct {
	commander *commanders.Commander

	Version string
}

// Setup initializes the NuGet plugin
func (n *NuGet) Setup() error {

	cmd, err := commanders.Open("nuget.exe", "./", "./.nuget")
	if err != nil {
		return err
	}

	output, err := cmd.ExecOutput()
	if err != nil {
		return err
	}

	version, err := getVersion(output)
	if err != nil {
		return err
	}

	n.Version = version
	n.commander = cmd

	return nil
}

// Gather gathers all the facts for a pinch
func (n *NuGet) Gather(opts map[string]string) (*models.Result, error) {
	return &models.Result{}, nil
}

func getVersion(data []byte) (string, error) {
	reader := bytes.NewReader(data)
	scanner := bufio.NewScanner(reader)

	var version string

	for scanner.Scan() {
		txt := scanner.Text()
		if strings.HasPrefix(txt, "NuGet Version: ") {
			version = versionex.FindString(txt)
		}

		if version != "" {
			break
		}
	}

	err := scanner.Err()
	return version, err
}

func init() {
	plugins.RegisterFactPlugin("nuget", &NuGet{})
}
