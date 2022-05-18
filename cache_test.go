package odeskidb

import (
	"reflect"
	"testing"
)

func TestDeserializeToMap(t *testing.T) {
	testCaseOne := []byte("\nhello:world\nhhhhhhhhhhh:16\njack:be nimble, jack be quick, jack jump over the candle stick\nvenom:i'm venom\n")
	expectedOutputOne := map[string][]byte{
		"hello":       []byte("world"),
		"hhhhhhhhhhh": []byte("16"),
		"jack":        []byte("be nimble, jack be quick, jack jump over the candle stick"),
		"venom":       []byte("i'm venom"),
	}

	outputOne, err := deserializeToMap(testCaseOne)
	if err != nil {
		t.Fatalf("An error occurred while trying to deserialize test case one. This is the error\n%s", err)
	}

	if !reflect.DeepEqual(expectedOutputOne, outputOne) {
		t.Error("test case one was not properly deserialized")
	}
}

func TestSerializeToBytes(t *testing.T) {
	/* NOTES:
	1. ALL THINGS CONSIDERED the method tested works as intended(for now, it's not yet robust), but due to
	how I'm testing it comes can "randomly" fail tests. I need a way to test of equality that does not
	care about order.
	*/
	testCaseOne := map[string][]byte{
		"z":           []byte("the end"),
		"hello":       []byte("world"),
		"hhhhhhhhhhh": []byte("16"),
		"jack":        []byte("be nimble, jack be quick, jack jump over the candle stick"),
		"venom":       []byte("i'm venom"),
	}
	expectedOutputOne := []byte("\nz:the end\nhello:world\nhhhhhhhhhhh:16\njack:be nimble, jack be quick, jack jump over the candle stick\nvenom:i'm venom\n")

	outputOne, err := serializeToBytes(testCaseOne)
	if err != nil {
		t.Fatalf("An error occurred while trying to serialize test case one. This is the error\n%s", err)
	}

	if !reflect.DeepEqual(expectedOutputOne, outputOne) {
		t.Error("test case one was not properly serialized")
	}
}
