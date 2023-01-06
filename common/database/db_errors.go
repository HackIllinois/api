package database

import (
	"errors"
	"fmt"

	"go.mongodb.org/mongo-driver/mongo"
)

var (
	ErrNotFound        = errors.New("Error: NOT_FOUND")
	ErrDuplicateKey    = errors.New("Error: DUPLICATE_KEY")
	ErrConnection      = errors.New("Error: CONNECTION_FAILED")
	ErrDisconnected    = errors.New("Error: CLIENT_DISCONNECTED")
	ErrUnknown         = errors.New("Error: UNKNOWN")
	ErrNilPassedToCRUD = errors.New("Error: NIL_PASSED_TO_CRUD") // this is an error the user should never get and is merely here for debugging
)

/*
	Converts internal mgo errors to external presented errors
*/
func convertMgoError(err error) error {
	// TODO: Need to try and encompass all possible errors
	switch err {
	case nil:
		return nil
	case mongo.ErrClientDisconnected:
		return ErrDisconnected
	case mongo.ErrNilDocument:
		return ErrNilPassedToCRUD
	case mongo.ErrNoDocuments:
		return ErrNotFound
	default:
		{
			var e mongo.WriteException
			if errors.As(err, &e) {
				for _, we := range e.WriteErrors {
					if we.Code == 11000 { // Error code for duplicate key error
						return ErrDuplicateKey
					}
				}
			}
		}
		fmt.Println("Unhandled error: ", err)
		// TODO: How can we embed error information into here?
		// It'll help a lot if an unexpected error comes up
		return ErrUnknown
	}
}
