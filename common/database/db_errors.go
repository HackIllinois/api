package database

import (
	"errors"
	"gopkg.in/mgo.v2"
)

var (
	ErrNotFound   = errors.New("Error: NOT_FOUND")
	ErrConnection = errors.New("Error: CONNECTION_FAILED")
	ErrUnknown    = errors.New("Error: UNKNOWN")
)

/*
	Converts internal mgo errors to external presented errors
*/
func convertMgoError(err error) error {
	if err == nil {
		return nil
	} else if err == mgo.ErrNotFound {
		return ErrNotFound
	}

	return ErrUnknown
}
