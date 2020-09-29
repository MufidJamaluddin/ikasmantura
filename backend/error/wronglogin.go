package error

import "fmt"

// @author Mufid Jamaluddin
type WrongLoginError struct {
	CustomError
	Username string
}

func (p *WrongLoginError) Error() string {
	return fmt.Sprintf("Username %s or password is false!", p.Username)
}
