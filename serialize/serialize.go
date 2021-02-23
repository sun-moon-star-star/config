package serialize

type Serializable interface {
	serialize(map[string]interface{}) ([]byte, error)
	unserialize([]byte) (map[string]interface{}, error)
}

func DefaultSerializer() Serializable {
	return &JSON{}
}
