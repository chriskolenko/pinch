package commanders

import (
	"os"
	"os/exec"
	"path/filepath"

	"github.com/Sirupsen/logrus"
)

// Commander the way to execute commands
type Commander struct {
	Binary string
}

// ExecOutput executes a command and returns the stdout as string
func (c *Commander) ExecOutput(args ...string) ([]byte, error) {
	logrus.WithFields(logrus.Fields{"binary": c.Binary, "args": args}).Debug("Running ExecOutput")

	out, err := exec.Command(c.Binary, args...).CombinedOutput()
	return out, err
}

// Open returns a commander
func Open(binary string, directories ...string) (*Commander, error) {
	for _, dir := range directories {
		p := filepath.Join(dir, binary)
		if _, err := os.Stat(p); err == nil {
			return &Commander{Binary: p}, nil
		}
	}

	p, err := exec.LookPath(binary)
	if err != nil {
		return nil, err
	}

	return &Commander{Binary: p}, nil
}
