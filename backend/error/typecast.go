package error

import "fmt"

type TypeCastError struct {
	error
	Type string
}

func (p *TypeCastError) Error() string {
	return fmt.Sprintf("Error in casting type %s", p.Type)
}
