package odeskidb

import (
	"bytes"
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

//TODO: make better error message/format
func (q get) Execute() (query, string, error) {

	var result string

	if !isDatabaseInitalized() {
		return q, result, UninitializedDatabaseError{
			"struct get: func Execute()",
			"You must initalize your database once before use",
		}
	}

	const colonByte byte = 58
	const newLineByte byte = 10

	keyStartIndex := bytes.Index(
		databaseCache,
		append(
			append([]byte{newLineByte}, q.key...),
			colonByte,
		),
	)

	if keyStartIndex >= 0 {
		valueStartIndex := keyStartIndex + 1 + len(q.key)
		valueEndIndex := valueStartIndex

		for databaseCache[valueEndIndex] != newLineByte {
			valueEndIndex++
		}

		result = string(databaseCache[valueStartIndex+1 : valueEndIndex])
	}

	return q, result, nil
}
