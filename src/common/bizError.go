package common

type BizError struct {
	Mesage string
}

func (e *BizError) Error() string {
	return e.Mesage
}
