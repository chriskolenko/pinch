package nunit

import "path/filepath"

// Console3 can find and run nunit console v3.
type Console3 struct {
	isloaded bool
	pattern  string
	path     string
}

func (c *Console3) load() error {
	if c.isloaded {
		return nil
	}

	// load it
	found, err := filepath.Glob(c.pattern)
	if err != nil {
		return err
	}

	if len(found) > 0 {
		c.path = found[0]
		c.isloaded = true
		return nil
	}

	// not found.
	c.isloaded = false
	return errNotFound
}

// IsFound finds all
func (c *Console3) IsFound() (bool, error) {
	err := c.load()
	if err != nil {
		return true, nil
	}

	return false, err
}

// Run runs the nunit console.
func (c *Console3) Run() error {
	return nil
}

// NewConsole3 returns NUnit Console 3
func NewConsole3(wd string) *Console3 {
	return &Console3{
		pattern: filepath.Join(wd, "packages", "*", "tools", "nunit3-console.exe"),
	}
}
