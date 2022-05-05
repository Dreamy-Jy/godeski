package odeskidb

import (
	"bytes"
	"errors"
	"strings"
)

type delete struct {
	key []byte
}

func (q delete) Key() string {
	return string(q.key)
}

func (q *delete) SetKey(newKey string) error {
	if !strings.ContainsAny(newKey, ":\n") {
		q.key = []byte(newKey)
		return nil
	}

	return ImproperlyFormatedQueryFieldError{
		"struct delete: func SetKey()",
		"the key you are trying to use for this DELETE query is not properly formatted",
	}
}

func (q delete) Execute() (query, string, error) {
	fileFormatError := errors.New("odeskidb: struct delete: func Execute(): The database's file(s) is not properily formated")

	var result string

	if !isDatabaseInitalized() {
		return q, result, UninitializedDatabaseError{
			"struct delete: func Execute()",
			"You must initalize your database once before use",
		}
	}

	const colonByte byte = 58
	const newLineByte byte = 10

	if databaseCache[len(databaseCache)-1] != newLineByte {
		return q, result, fileFormatError
	}

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

		newDatabaseFileData := bytes.Replace(
			databaseCache,
			databaseCache[keyStartIndex+1:valueEndIndex+1],
			[]byte{},
			1,
		)

		if writingFileError := overwriteDatabase(
			newDatabaseFileData,
		); writingFileError != nil {
			_ = overwriteDatabase(databaseCache) // this is a revert
			return q, result, writingFileError
		}

		databaseCache = newDatabaseFileData
	}

	return q, result, nil
}
