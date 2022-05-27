package odeskidb

import (
	"bytes"
	"strings"
)

type set struct {
	key   []byte
	value []byte
}

func (q set) Key() string {
	return string(q.key)
}

func (q *set) SetKey(newKey string) error {
	if !strings.ContainsAny(newKey, ":\n") {
		q.key = []byte(newKey)
		return nil
	}

	return ImproperlyFormatedQueryFieldError{
		"struct Set: func SetKey()",
		"the key you are trying to use for this SET query is not properly formatted",
	}
}

func (q set) Value() string {
	return string(q.value)
}

func (q *set) SetValue(newValue string) error {
	if !strings.ContainsAny(newValue, ":\n") {
		q.value = []byte(newValue)
		return nil
	}

	return ImproperlyFormatedQueryFieldError{
		"struct Set: func SetValue()",
		"the value you are trying to use for this SET query is not properly formatted",
	}
}

func (q set) Execute() (query, string, error) {
	//you currently have no convenient method to test wheither set is properily set.
	var result string

	if !isDatabaseInitalized() {
		return q, result, UninitializedDatabaseError{
			"struct set: func Execute()",
			"You must initalize your database once before use",
		}
	}

	value, ok := databaseCache[q.Key()]
	if !ok || !bytes.Equal(value, q.value) {
		databaseCache[q.Key()] = q.value
	}

	newData, serializationError := serializeToBytes(databaseCache)
	if serializationError != nil {
		if !ok {
			delete(databaseCache, q.Key())
		} else if ok && !bytes.Equal(value, q.value) {
			databaseCache[q.Key()] = value
		}

		return q, result, serializationError
	}

	if databaseWriteError := overwriteDatabase(newData); databaseWriteError != nil {
		// how can we solve this in a way that does not lead to nil again?
		//TODO: how to revert from a failed overwrite?
		return q, result, databaseWriteError
	}

	return q, result, nil
}
