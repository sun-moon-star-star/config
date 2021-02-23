package serialize

type Serializable interface {
	Serialize(interface{}) ([]byte, error)
	Unserialize([]byte) (interface{}, error)
}

func DefaultSerializer() Serializable {
	return &JSON{}
}
