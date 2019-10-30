package task

type errorString struct {
	s string
}

func (e *errorString) Error() string {
	return e.s
}

func Error(text string) error {
	return &errorString{text}
}