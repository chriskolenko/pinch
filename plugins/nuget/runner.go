package nuget

import (
	"bufio"
	"bytes"
	"strings"

	"github.com/webcanvas/pinch/shared/commanders"
	"github.com/webcanvas/pinch/shared/models"
)

// Runner runs the plugin.
type Runner struct {
	cmd commanders.Commander

	Version string
}

// Run runs it!
func (r *Runner) Run() (models.Result, error) {
	return models.Result{}, nil
}

// CreateRunner finds nuget and returns an object that can run the plugin
func CreateRunner(opts Options) (*Runner, error) {

	cmd, err := commanders.Open("nuget.exe", "./", "./.nuget")
	if err != nil {
		return nil, err
	}

	output, err := cmd.ExecOutput()
	if err != nil {
		return nil, err
	}

	version, err := getVersion(output)
	if err != nil {
		return nil, err
	}

	return &Runner{
		cmd:     *cmd,
		Version: version,
	}, nil
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
