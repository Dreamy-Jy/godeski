package odeskidb

import (
	"bytes"
	"strings"
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

	3. We need a way for functions to know where their being
	called so we can put it in error messages and logs.
	*/
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
				"func deserializeToMap()",
				`A subslice of the data, that was considered to a Key Value pair, 
				was found to have more that one colon. This is not allowed.`,
			}
		}

		if _, ok := output[key]; ok {
			return nil, ImproperlyFormatedDataError{
				"func deserializeToMap()",
				`There where multiple uses of same key in this data set. 
				That is not allowed.`,
			}
		}
		output[key] = splitKV[1]
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
		if strings.Contains(key, "\n:") {
			return nil, ImproperlyFormatedDataError{
				"func serializeToBytes()",
				"A key contains one or both of the forbidden characters, '\n' or/and ':'.",
			}
		}

		if bytes.Contains(value, []byte{colonByte, newLineByte}) {
			return nil, ImproperlyFormatedDataError{
				"func serializeToBytes()",
				"A value contains one or both of the forbidden bytes, '\n'(10) or/and ':'(58).",
			}
		}

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
