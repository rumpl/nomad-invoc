package main

import (
	"fmt"

	"github.com/docker/app/types"
	"github.com/hashicorp/nomad/api"
	"github.com/rumpl/nomad-invoc/pkg/nomad"
)

func installAction(n nomad.NomadInvocation, name string) error {
	job := api.Job{}
	fmt.Println(job)
	app, err := getApp()
	if err != nil {
		return err
	}
	fmt.Println(app)
	return n.Install()
}

func getApp() (*types.App, error) {
	return types.NewAppFromDefaultFiles("./")
}
