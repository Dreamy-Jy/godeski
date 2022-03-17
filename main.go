package odeskidb

import (
	"os"
)

type query interface {
	isQuery() bool
}

type Get struct {
	key string
}
type Set struct {
	key   string
	value string
}
type Delete struct {
	key string
}
type Group struct {
	queries []interface{}
}

func (q *Get) isQuery() bool {
	return true
}
func (q *Set) isQuery() bool {
	return true
}
func (q *Delete) isQuery() bool {
	return true
}
func (q *Group) isQuery() bool {
	return true
}

// CreateDatabaseIfExist creates a database in the directory
// specifed by the databaseDirectory param, if a database does
// not exist there
func CreateDatabaseIfExist(databaseDirectory string) {
	err := os.WriteFile(
		databaseDirectory+"odeski.db",
		[]byte(""),
		0644,
	)

	if err != nil {
		panic(err)
	}
}

// DeleteDatabaseIfExist deletes the database at the directory
// specifed by the databaseDirectory param, if a database does
// not exist there
func DeleteDatabaseIfExist(databaseDirectory string) {
	err := os.Remove(databaseDirectory + "odeski.db")

	if err != nil {
		panic(err)
	}
}
