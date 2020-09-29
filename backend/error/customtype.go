package error

// @author Mufid Jamaluddin
type CustomError struct {
	error
}

func (p *CustomError) IsCanSendToClient() bool {
	return true
}
