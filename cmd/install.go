package main

import (
	"github.com/rumpl/nomad-invoc/pkg/nomad"
)

func installAction(n nomad.NomadInvocation, name string) error {
	return n.Install(name)
}
