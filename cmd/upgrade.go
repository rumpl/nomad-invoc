package main

import "github.com/rumpl/nomad_invoc/pkg/nomad"

func upgradeAction(n nomad.NomadInvocation, name string) error {
	return n.Upgrade()
}
