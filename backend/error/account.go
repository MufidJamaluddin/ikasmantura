package error

type NotVerifiedAccount struct {
	CustomError
}

func (p *NotVerifiedAccount) Error() string {
	return "Your account is not verified"
}
