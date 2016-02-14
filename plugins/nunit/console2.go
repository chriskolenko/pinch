package nunit

import "path/filepath"

// Console2 can find and run nunit console v2.
type Console2 struct {
	isloaded bool
	pattern  string
	path     string
}

func (c *Console2) load() error {
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
func (c *Console2) IsFound() (bool, error) {
	err := c.load()
	if err != nil {
		return true, nil
	}

	return false, err
}

// Run runs the nunit console.
func (c *Console2) Run() error {
	return nil
}

// NewConsole2 returns NUnit Console 2
func NewConsole2(wd string) *Console2 {
	return &Console2{
		pattern: filepath.Join(wd, "packages", "*", "tools", "nunit-console.exe"),
	}
}
