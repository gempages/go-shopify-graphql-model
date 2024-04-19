package model

func (e FileError) Error() string {
	if e.Details != nil {
		return *e.Details
	}
	return e.Message
}

func (e MediaError) Error() string {
	if e.Details != nil {
		return *e.Details
	}
	return e.Message
}
