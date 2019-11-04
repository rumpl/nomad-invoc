package nomad

// NomadInvocation represents the actions that
// it can perform
type NomadInvocation interface {
	Install(name string) error
	Upgrade() error
	Uninstall() error
}

type nomadInvocation struct {
}

// New creates a new Nomad invocation
func New() NomadInvocation {
	return nomadInvocation{}
}
