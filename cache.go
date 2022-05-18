package odeskidb

import (
	"bytes"
)

func deserializeToMap(data []byte) (map[string][]byte, error) {
	/* NOTES:
	1. This function doesn't currently return an error,
	but it should return an improper format error.

	2. This function is not the most efficient choice for
	this problem as the data grows.
		2b. This whole system isn't prepared for data that
		doesn't fit in memory
		2c. This system is not sorted.
	*/
	keyValues := bytes.Split(
		data[1:len(data)-1],
		[]byte{newLineByte},
	)
	output := make(map[string][]byte, len(keyValues))

	for _, kv := range keyValues {
		splitKV := bytes.Split(kv, []byte{colonByte})
		output[string(splitKV[0])] = splitKV[1]
	}

	return output, nil
}

func serializeToBytes(data map[string][]byte) ([]byte, error) {
	/* NOTES:
	1. ALL THINGS CONSIDERED this method works as intended, but due to
	how I'm testing it comes can "randomly" fail tests.
	*/
	output := []byte{newLineByte}

	for key, value := range data {
		var entry []byte = append(
			[]byte(key),
			append(
				[]byte{colonByte},
				append(value, newLineByte)...,
			)...,
		)
		output = append(output, entry...)
	}
	return output, nil
}
