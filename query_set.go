package odeskidb

import (
	"bytes"
	"errors"
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
	fileFormatError := errors.New("odeskidb: struct Set: func Execute(): The database's file(s) is not properily formated")

	var result string

	if !isDatabaseInitalized() {
		return q, result, UninitializedDatabaseError{
			"struct set: func Execute()",
			"You must initalize your database once before use",
		}
	}

	const colonByte byte = 58
	const newLineByte byte = 10

	if databaseCache[len(databaseCache)-1] != newLineByte {
		return q, result, fileFormatError
	}

	var newDatabaseFileData []byte

	newDatabaseEntry := append(
		append(q.key, colonByte),
		append(q.value, newLineByte)...,
	)

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

		newDatabaseFileData = bytes.Replace(
			databaseCache,
			databaseCache[keyStartIndex+1:valueEndIndex+1],
			newDatabaseEntry,
			1,
		)
	} else {
		newDatabaseFileData = append(
			databaseCache,
			newDatabaseEntry...,
		)
	}

	if writingFileError := overwriteDatabase(
		newDatabaseFileData,
	); writingFileError != nil {
		_ = overwriteDatabase(databaseCache) // this is a revert
		return q, result, writingFileError
	}

	databaseCache = newDatabaseFileData

	return q, result, nil
}
