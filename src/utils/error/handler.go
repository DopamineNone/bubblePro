package error

type Error struct {
	Code      int
	Message   string
	ExtraInfo string
}

func (e *Error) Error() string {
	return e.Message
}

func (e *Error) GetCode() int {
	return e.Code
}
