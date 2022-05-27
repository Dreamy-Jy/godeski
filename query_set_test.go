package odeskidb

import (
	"bytes"
	"testing"
)

func TestSet_Execute(t *testing.T) {
	/* NOTE:
	1. We probably will want to have a universal start-up
	function for test that need to use the initalized
	database like this one

	2. Currently we need to check both the physical files
	and the in memory database cache.
	*/

	if startError := InitalizeOdeski("./"); startError != nil {
		t.Fatalf("Odeski could not be initalized. The error is on the following line.\n%s", startError)
	}

	s1, setError := NewSet("jordane", "thomas")
	if setError != nil {
		t.Fatalf("The SET query was not created. The error is on the following line.\n%s", setError)
	}

	_, _, setExecutionError := s1.Execute()
	if setExecutionError != nil {
		t.Fatalf("The SET query was not created. The error is on the following line.\n%s", setExecutionError)
	}

	if value, ok := databaseCache[s1.Key()]; !ok {
		t.Fatal("The SET query was unable to create the desired database entry.")
	} else if ok && !bytes.Equal(value, s1.value) {
		t.Fatal("The SET query was unable to change the value of the desired database key.")
	}

	// TODO: Find a way to test the contents of the file.

}
