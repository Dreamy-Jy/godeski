package odeskidb

import "testing"

func TestRemove_Execute(t *testing.T) {
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

	r1, removeError := NewRemove("hhhhhhhhhhh")
	if removeError != nil {
		t.Fatalf("The REMOVE query was not created. The error is on the following line.\n%s", removeError)
	}

	_, _, removeExecutionError := r1.Execute()
	if removeExecutionError != nil {
		t.Fatalf("The REMOVE query was not created. The error is on the following line.\n%s", removeExecutionError)
	}

	if _, ok := databaseCache[r1.Key()]; ok {
		t.Fatal("The REMOVE query was unable to delete the data specified in the query.")
	}

	// TODO: Find a way to test the contents of the file.
}
