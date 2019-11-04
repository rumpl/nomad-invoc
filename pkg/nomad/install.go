package nomad

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"

	"github.com/docker/app/render"
	"github.com/docker/app/types"
	"github.com/hashicorp/nomad/api"
	"github.com/pkg/errors"
)

func (n nomadInvocation) Install(name string) error {
	fmt.Println("Installing app", name)
	// Placeholder to keep the import of nomad because I am scared
	// of go dependency management
	job := api.Job{}
	fmt.Println(job)

	app, err := getApp()
	if err != nil {
		return err
	}

	// nil params and image map
	rendered, err := render.Render(app, nil, nil)
	if err != nil {
		return err
	}

	for _, service := range rendered.Services {
		fmt.Println("Service", service.Name)
	}

	return nil
}

func getApp() (*types.App, error) {
	cwd, err := os.Getwd()
	if err != nil {
		return nil, errors.Wrap(err, "cannot resolve current working directory")
	}

	name, err := findApp(cwd)

	if err != nil {
		return nil, err
	}

	return types.NewAppFromDefaultFiles(name)
}

// findApp looks for an app in CWD or subdirs
func findApp(cwd string) (string, error) {
	if strings.HasSuffix(cwd, ".dockerapp") {
		return cwd, nil
	}
	content, err := ioutil.ReadDir(cwd)
	if err != nil {
		return "", errors.Wrap(err, "failed to read current working directory")
	}
	hit := ""
	for _, c := range content {
		if strings.HasSuffix(c.Name(), ".dockerapp") {
			if hit != "" {
				return "", fmt.Errorf("multiple applications found in current directory, specify the application name on the command line")
			}
			hit = c.Name()
		}
	}
	if hit == "" {
		return "", fmt.Errorf("no application found in current directory")
	}
	return filepath.Join(cwd, hit), nil
}
