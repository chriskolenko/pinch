package nunit

import "errors"

var errNotFound = errors.New("NUnit Console not found")

// Consoles holds the list of nunit consoles
type Consoles []Console

// Find returns a console.
func (c Consoles) Find() (Console, error) {

	for _, console := range c {
		b, err := console.IsFound()
		if err != nil {
			return nil, err
		}

		if b {
			return console, nil
		}
	}

	// not found.
	return nil, errNotFound
}
