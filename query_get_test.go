package odeskidb

import "testing"

func TestGet_Execute(t *testing.T) {
	if startError := InitalizeOdeski("./"); startError != nil {
		t.Fatalf("Odeski could not be initalized. The error is on the following line.\n%s", startError)
	}

	g1, getError := NewGet("hello")
	if getError != nil {
		t.Fatalf("The GET query was not created. The error is on the following line.\n%s", getError)
	}
	_, result, getExecutionError := g1.Execute()
	if getExecutionError != nil {
		t.Fatalf("The GET query was not created. The error is on the following line.\n%s", getError)
	}
	if result != "world" {
		t.Error("GET did not return the proper result.")
	}

}
