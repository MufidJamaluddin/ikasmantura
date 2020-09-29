package error

import "fmt"

type WrongLoginError struct {
	error
	Username string
}

func (p *WrongLoginError) Error() string {
	return fmt.Sprintf("Username %s or password is false!", p.Username)
}
