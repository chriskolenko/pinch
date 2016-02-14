package nunit

// Console is the interface for NUnit Consoles.
type Console interface {
	IsFound() (bool, error)
	Run() error
}
