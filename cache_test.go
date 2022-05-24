package odeskidb

import (
	"bytes"
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
		"hhhhhhhhhhh": []byte("16"),
		"hello":       []byte("world"),
		"jack":        []byte("be nimble, jack be quick, jack jump over the candle stick"),
		"venom":       []byte("i'm venom"),
	}
	// expectedOutputOne := []byte("\nz:the end\nhello:world\nhhhhhhhhhhh:16\njack:be nimble, jack be quick, jack jump over the candle stick\nvenom:i'm venom\n")

	outputOne, err := serializeToBytes(testCaseOne)
	if err != nil {
		t.Fatalf("An error occurred while trying to serialize test case one. This is the error\n%s", err)
	}

	checkOne, err2 := bytesToSet(outputOne)
	if err != nil {
		t.Fatalf("An error occurred while trying to serialize test case one, that was not caught by func serializeToBytes(). This is the error\n%s", err2)
	}

	if !reflect.DeepEqual(testCaseOne, checkOne) {
		t.Error("Test case one was not properly serialized")
	}
}

func bytesToSet(data []byte) (map[string][]byte, error) {
	keyValues := bytes.Split(
		data[1:len(data)-1], // excudes the trailing and leading '\n' runes
		[]byte{newLineByte},
	)
	output := make(map[string][]byte, len(keyValues))

	for _, kv := range keyValues {
		splitKV := bytes.Split(kv, []byte{colonByte})
		key := string(splitKV[0])

		if len(splitKV) != 2 {
			return nil, ImproperlyFormatedDataError{
				"func TestSerializeToBytes(): func bytesToSet()",
				`A subslice of the data, that was considered to a Key Value pair, 
				was found to have more that one colon. This is not allowed.`,
			}
		}

		if _, ok := output[key]; ok {
			return nil, ImproperlyFormatedDataError{
				"func TestSerializeToBytes(): func bytesToSet()",
				`There where multiple uses of same key in this data set. 
				That is not allowed.`,
			}
		}
		output[key] = splitKV[1]
	}

	return output, nil
}
