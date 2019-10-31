package main

import "github.com/rumpl/nomad_invoc/pkg/nomad"

func uninstallAction(n nomad.NomadInvocation, name string) error {
	return n.Uninstall()
}
