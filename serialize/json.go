package serialize

import "encoding/json"

type JSON struct{}

func (*JSON) serialize(data map[string]interface{}) ([]byte, error) {
	return json.Marshal(data)
}

func (*JSON) unserialize(jsonBytes []byte) (map[string]interface{}, error) {
	var data map[string]interface{}
	err := json.Unmarshal(jsonBytes, &data)
	return data, err
}
