package commanders

import (
	"os"
	"os/exec"
	"path"
)

// Commander the way to execute commands
type Commander struct {
	Binary string
}

// ExecOutput executes a command and returns the stdout as string
func (c *Commander) ExecOutput(args ...string) ([]byte, error) {
	out, err := exec.Command(c.Binary, args...).Output()
	if err != nil {
		return nil, err
	}
	return out, nil
}

// Open returns a commander
func Open(binary string, directories ...string) (*Commander, error) {
	for _, dir := range directories {
		p := path.Join(dir, binary)
		if _, err := os.Stat(p); err == nil {
			return &Commander{Binary: p}, nil
		}
	}

	p, err := exec.LookPath(binary)
	if err != nil {
		return nil, nil
	}

	return &Commander{Binary: p}, nil
}
