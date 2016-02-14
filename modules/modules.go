package modules

var modules = make(map[string]Module)

type Module interface {
	GetVersions() []string
}

func RegisterModule(name string, module Module) {
	// TODO maybe valid if it's a good module.
	modules[name] = module
}

func GetModules() map[string]Module {
	return modules
}
