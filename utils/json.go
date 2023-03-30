package utils

import "encoding/json"

func JsonMarshal(v any) []byte {
	bz, err := json.Marshal(v)
	if err != nil {
		panic(err)
	}
	return bz
}

func JsonUnmarshal(bz []byte, v interface{}) {
	if err := json.Unmarshal(bz, v); err != nil {
		panic(err)
	}
}
