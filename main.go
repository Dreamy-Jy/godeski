package odeskidb

import (
	"errors"
	"os"
)

//TODO: make better error message/format

// will path code work for windows and macos?
var databasePath string

// adding a global cache adds new problems
var databaseCache map[string][]byte // WARNING: Dangerous

const colonByte byte = 58
const newLineByte byte = 10

func NewGet(key string) (*get, error) {
	newQuery := get{}

	if err := newQuery.SetKey(key); err != nil {
		return nil, err
	}

	return &newQuery, nil
}

func NewSet(key string, value string) (*set, error) {
	newQuery := set{}

	if err := newQuery.SetKey(key); err != nil {
		return nil, err
	}

	if err := newQuery.SetValue(value); err != nil {
		return nil, err
	}

	return &newQuery, nil
}

func NewRemove(key string) (*remove, error) {
	newQuery := remove{}

	if err := newQuery.SetKey(key); err != nil {
		return nil, err
	}

	return &newQuery, nil
}

func InitalizeOdeski(path string) error {
	//check if directory is valid

	// will path code work for windows and macos?

	//at some point if we're editing the db path in multipule functions
	//we should make this it's own function
	var initalizeCacheError error
	initalizeDatabase := func(path string) error {
		if directoryInfo, err1 := os.Stat(path); err1 != nil {
			return err1
		} else if !directoryInfo.IsDir() {
			return errors.New("path param does not point to a directory")
		} else if path[len(path)-1] != '/' {
			return errors.New("path param must end with '/' character")
		}

		//check if database already initalized in directory and if not create a db
		if _, err2 := os.Stat(path + "odeski.db"); os.IsNotExist(err2) {
			if err3 := os.WriteFile(path+"odeski.db", []byte(""), 0644); err3 != nil {
				return err3
			}
		} else if err2 != nil {
			return err2
		}

		return nil
	}
	initalizeCache := func() (map[string][]byte, error) {
		fileFormatError := errors.New("odeskidb: func InitalizeOdeski(): The database's file(s) is/are not properily formated")

		var databaseFile, openingFileError = os.Open(databasePath + "odeski.db")
		if openingFileError != nil {
			return nil, openingFileError
		}

		var databaseFileInfo, gettingFileInforError = databaseFile.Stat()
		if gettingFileInforError != nil {
			return nil, gettingFileInforError
		}

		databaseData := make([]byte, databaseFileInfo.Size())
		var _, readingFileError = databaseFile.Read(databaseData)
		if readingFileError != nil {
			return nil, readingFileError
		}
		databaseFile.Close()

		if databaseData[0] != newLineByte ||
			databaseData[len(databaseData)-1] != newLineByte {
			return nil, fileFormatError
		}

		databaseMap, deserializationError := deserializeToMap(databaseData)
		if deserializationError != nil {
			return nil, deserializationError
		}

		return databaseMap, nil
	}

	databasePath = path

	initalizeDatabaseError := initalizeDatabase(path)
	if initalizeDatabaseError != nil {
		databasePath = ""
		databaseCache = nil
		return initalizeDatabaseError
	}

	databaseCache, initalizeCacheError = initalizeCache()
	if initalizeCacheError != nil {
		databasePath = ""
		databaseCache = nil
		return initalizeCacheError
	}

	return nil
}
