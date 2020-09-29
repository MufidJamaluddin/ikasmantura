package error

import "fmt"

// @author Mufid Jamaluddin
type TypeCastError struct {
	CustomError
	Type string
}

func (p *TypeCastError) Error() string {
	return fmt.Sprintf("Error in casting type %s", p.Type)
}
