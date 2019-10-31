package nomad

import "fmt"

func (n nomadInvocation) Upgrade() error {
	fmt.Println("install")
	return nil
}
