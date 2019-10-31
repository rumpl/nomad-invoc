package main

import "github.com/rumpl/nomad-invoc/pkg/nomad"

func uninstallAction(n nomad.NomadInvocation, name string) error {
	return n.Uninstall()
}
