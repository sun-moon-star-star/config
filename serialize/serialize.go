package serialize

type Serializable interface {
	serialize(map[string]interface{}) []byte
	unserialize([]byte) map[string]interface{}
}
