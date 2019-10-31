package nomad

type NomadInvocation interface {
	Install() error
	Upgrade() error
	Uninstall() error
}

type nomadInvocation struct {
}

func New() NomadInvocation {
	return nomadInvocation{}
}
