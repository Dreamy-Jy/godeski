/*
TODO: Place all errors here, for communication perposes
*/
package odeskidb

import (
	"fmt"
)

//TODO: Consider new error naming scheme
type ImproperlyFormatedQueryFieldError struct {
	location string
	message  string
	// field    string
}

func (err ImproperlyFormatedQueryFieldError) Error() string {
	return fmt.Sprintf("ImproperlyFormatedQueryFieldError | odeskidb: %s:\n\t%s", err.location, err.message)
}

type UninitializedDatabaseError struct {
	location string
	message  string
}

func (err UninitializedDatabaseError) Error() string {
	return fmt.Sprintf("UninitializedDatabaseError | odeskidb: %s:\n\t%s", err.location, err.message)
}
