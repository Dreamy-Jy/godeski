package odeskidb

import (
	"strings"
)

type remove struct {
	key []byte
}

func (q remove) Key() string {
	return string(q.key)
}

func (q *remove) SetKey(newKey string) error {
	if !strings.ContainsAny(newKey, ":\n") {
		q.key = []byte(newKey)
		return nil
	}

	return ImproperlyFormatedQueryFieldError{
		"struct remove: func SetKey()",
		"the key you are trying to use for this REMOVE query is not properly formatted",
	}
}

/*
NOTE: Should I return the value of the deleted entry?

NOTE: When should we over write data?
*/
func (q remove) Execute() (query, string, error) {
	var result string

	if !isDatabaseInitalized() {
		return q, result, UninitializedDatabaseError{
			"struct remove: func Execute()",
			"You must initalize your database once before use",
		}
	}

	value, ok := databaseCache[q.Key()]

	delete(databaseCache, q.Key())

	newData, serializationError := serializeToBytes(databaseCache)
	if serializationError != nil {
		if ok {
			databaseCache[q.Key()] = value
		}
		return q, result, serializationError
	}

	databaseWriteError := overwriteDatabase(newData)
	if databaseWriteError != nil {
		// how can we solve this in a way that does not lead to nil again?
		//TODO: how to revert from a failed overwrite?
		return q, result, databaseWriteError
	}

	return q, result, nil
}
