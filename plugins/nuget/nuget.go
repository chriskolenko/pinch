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

type nuget struct {
	commander *commanders.Commander

	Version string
}

// Setup initializes the NuGet plugin
func (n *nuget) Setup(models.Raw) (result models.Result, err error) {
	cmd, err := commanders.Open("nuget.exe", "./", "./.nuget")
	if err != nil {
		return
	}

	output, err := cmd.ExecOutput()
	if err != nil {
		return
	}

	version, err := getVersion(output)
	if err != nil {
		return
	}

	n.Version = version
	n.commander = cmd

	return
}

// Exec runs the pinch
func (n *nuget) Exec(opts models.Raw) (models.Result, error) {
	return models.Result{}, nil
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
	plugins.RegisterPinchPlugin("nuget", &nuget{})
}
