package usecase

import "fmt"

func errNoRepository(local string) error {
	return fmt.Errorf("cann't start a service without a repository on %s", local)
}
