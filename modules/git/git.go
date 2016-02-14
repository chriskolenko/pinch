package git

import (
	"fmt"
	"os/exec"
	"regexp"

	"github.com/webcanvas/pinch/modules"
)

type Module struct {
	Version string
}

func (m Module) GetVersions() []string {
	if m.Version == "" {
		return nil
	}

	return []string{m.Version}
}

func init() {
	// find the git client.
	if _, err := exec.LookPath("git"); err != nil {
		fmt.Errorf("No git: %s", err)
		return
	}

	args := []string{"version"}
	output, err := exec.Command("git", args...).CombinedOutput()
	if err != nil {
		fmt.Errorf("Error trying to use git: %s (%s)", err, output)
		return
	}

	r := regexp.MustCompile(`[\d.]+`)
	bytes := r.Find(output)

	mod := &Module{
		Version: string(bytes[:]),
	}
	fmt.Println("At git module")
	modules.RegisterModule("git", mod)
}
