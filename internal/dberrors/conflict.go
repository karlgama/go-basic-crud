package dberrors

type ConflictError struct {
}

func (ConflictError) Error() string {
	return "attempt to create a record with an existing key"
}
