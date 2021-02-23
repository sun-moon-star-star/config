package serialize

type Serializable interface {
	Serialize(map[string]interface{}) ([]byte, error)
	Unserialize([]byte) (map[string]interface{}, error)
}

func DefaultSerializer() Serializable {
	return &JSON{}
}
