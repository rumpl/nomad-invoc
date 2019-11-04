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

	app, err := getApp()
	if err != nil {
		return err
	}

	// nil params and image map
	rendered, err := render.Render(app, nil, nil)
	if err != nil {
		return err
	}

	tasks := []*api.Task{}
	for _, service := range rendered.Services {
		fmt.Println("Service", service.Name)
		tasks = append(tasks, &api.Task{
			Name:   service.Name,
			Driver: "docker",
			Config: map[string]interface{}{
				"image": service.Image,
				"args":  service.Command,
			},
			Resources: &api.Resources{
				Networks: []*api.NetworkResource{
					&api.NetworkResource{
						ReservedPorts: []api.Port{
							api.Port{
								Label: "http",
								To:    int(service.Ports[0].Target),
								Value: int(service.Ports[0].Published),
							},
						},
					},
				},
			},
			Services: []*api.Service{
				&api.Service{
					Name:      service.Name,
					PortLabel: "http",
				},
			},
		})
	}

	dd := "hello"
	tname := "tasks"
	job := api.Job{
		ID: &dd,
		Datacenters: []string{
			"dc1",
		},
		TaskGroups: []*api.TaskGroup{
			&api.TaskGroup{
				Name:  &tname,
				Tasks: tasks,
			},
		},
	}

	config := api.DefaultConfig()
	config.Address = "http://host.docker.internal:4646"
	client, err := api.NewClient(config)
	if err != nil {
		return err
	}

	opts := &api.RegisterOptions{}
	resp, _, err := client.Jobs().RegisterOpts(&job, opts, nil)
	if err != nil {
		return err
	}
	if resp.Warnings != "" {
		fmt.Println(resp.Warnings)
	}
	fmt.Println("Eval ID", resp.EvalID)

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
