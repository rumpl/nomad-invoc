package main

import (
	"errors"
	"fmt"
	"os"

	"github.com/rumpl/nomad_invoc/pkg/nomad"
)

type cnabAction func(nomad.NomadInvocation, string) error

var (
	cnabActions = map[string]cnabAction{
		"install":   installAction,
		"upgrade":   upgradeAction, // upgrade is implemented as reinstall.
		"uninstall": uninstallAction,
	}
)

func getCnabAction() (cnabAction, string, error) {
	// CNAB_ACTION should always be set. but in future we want to have
	// claim-less actions. So we don't fail if no installation is set
	actionName, ok := os.LookupEnv("CNAB_ACTION")
	if !ok {
		return nil, "", errors.New("no CNAB action specified")
	}
	action, ok := cnabActions[actionName]
	if !ok {
		return nil, "", fmt.Errorf("action %q not supported", actionName)
	}
	return action, actionName, nil
}

func main() {
	action, actionName, err := getCnabAction()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error while parsing CNAB operation: %s", err)
		os.Exit(1)
	}
	n := nomad.New()
	instanceName := os.Getenv("CNAB_INSTALLATION_NAME")
	if err := action(n, instanceName); err != nil {
		fmt.Fprintf(os.Stderr, "Action %q failed: %s", actionName, err)
		os.Exit(1)
	}
}