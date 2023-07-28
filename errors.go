package gonep

type UnsuccessfulLoginError struct {
	Message string
}

func (e UnsuccessfulLoginError) Error() string {
	return e.Message
}
