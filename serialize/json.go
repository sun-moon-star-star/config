package serialize

import "encoding/json"

type JSON struct{}

func (*JSON) Serialize(data interface{}) ([]byte, error) {
	return json.Marshal(data)
}

func (*JSON) Unserialize(jsonBytes []byte) (interface{}, error) {
	var data map[string]interface{}
	err := json.Unmarshal(jsonBytes, &data)
	return data, err
}
