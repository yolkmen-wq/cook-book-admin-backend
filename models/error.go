package models

import "fmt"

type DbError struct {
	Msg string
}

func (e *DbError) Error() string {
	return fmt.Sprintf("Msg: %s", e.Msg)
}

func NewDbError(msg string) *DbError {
	return &DbError{
		Msg: msg,
	}
}
