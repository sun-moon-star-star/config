package storage

type Node struct {
	ID          int64  `json:"id"`
	CreatedTime uint64 `json:"created_time"` // nano
	UpdatedTime uint64 `json:"updated_time"` // nano
	Extends     []byte `json:"extends"`
}

type StorageAPI interface {
	ExistsNode(key string) bool
	CreateNode(key string) bool

	Create(key string, value interface{}) error // key not exist
	Update(key string, value interface{}) error // key exist
	Query(key string) (interface{}, error)
	Delete(key string) error
}
