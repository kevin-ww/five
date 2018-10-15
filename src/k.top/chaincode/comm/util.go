package comm

import "encoding/json"

//common utilities
func Marshal(a interface{}) []byte {
	bytes, _ := json.Marshal(a)
	return bytes
}

func Unmarshal(data []byte, target interface{}) error {
	return json.Unmarshal(data, target)
}
