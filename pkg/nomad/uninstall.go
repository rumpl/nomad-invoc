package nomad

import "fmt"

func (n nomadInvocation) Uninstall() error {
	fmt.Println("install")
	return nil
}
