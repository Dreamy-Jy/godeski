package odeskidb

import (
	"strings"
)

type get struct {
	key []byte
}

func (q get) Key() string {
	return string(q.key)
}

func (q *get) SetKey(newKey string) error {
	if !strings.ContainsAny(newKey, ":\n") {
		q.key = []byte(newKey)
		return nil
	}

	return ImproperlyFormatedQueryFieldError{
		"struct Get: func SetKey()",
		"the key you are trying to use for this GET query is not properly formatted",
	}
}

/*
TODO: make better error message/format

WARNING: this method won't work when we buffer changes
*/
func (q get) Execute() (query, string, error) {
	var result string

	if !isDatabaseInitalized() {
		return q, result, UninitializedDatabaseError{
			"struct get: func Execute()",
			"You must initalize your database once before use",
		}
	}

	if value, ok := databaseCache[q.Key()]; ok {
		result = string(value)
	}

	return q, result, nil
}
