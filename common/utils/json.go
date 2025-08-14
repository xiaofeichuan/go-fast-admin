package utils

import "encoding/json"

//type JsonUtil struct{}
//
//var Json = new(JsonUtil)

// Serialize 序列化
func Serialize(value any) string {
	data, err := json.Marshal(value)
	if err != nil {
		return ""
	}
	return string(data)
}

// Deserialize 反序列化
func Deserialize(data string, v any) {
	err := json.Unmarshal([]byte(data), &v)
	if err != nil {
		v = nil
	}
}
