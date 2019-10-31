package nomad

import "fmt"

func (n nomadInvocation) Install() error {
	fmt.Println("install")
	return nil
}
