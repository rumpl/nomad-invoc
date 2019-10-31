package main

import (
	"fmt"

	"github.com/hashicorp/nomad/api"
)

func main() {
	name := "name"
	job := api.Job{
		Name: &name,
	}

	fmt.Println(job)
}
