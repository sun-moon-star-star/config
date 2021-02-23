package storage

import "config"

type Node struct {
	ID          int64  `json:"id"`           // unique ID
	CreatedTime uint64 `json:"created_time"` // nano
	UpdatedTime uint64 `json:"updated_time"` // nano
	Type        int64  `json:"type"`         // last four bits is ConfigItem, first 60 bits no use
	Meta        []byte `json:"extends"`      // user defined data
}

type StorageAPI interface {
	ExistsNode(key string) error
	CreateNode(key string, nodeType config.ConfigItem, meta []byte, value interface{}) (int64, error)
	QueryNodeID(key string) (int64, error)
	QueryNodeType(key string) (int64, error)
	QueryNodeMeta(key string) ([]byte, error)
	UpdateNodeMeta(key string, meta []byte) error

	Update(key string, value interface{}) error
	Append(key string, value interface{}) error // only for ConfigItemArray

	// use ID...
	ExistsNodeByID(id int64) error
	CreateNodeWithID(id int64, nodeType config.ConfigItem, meta []byte, value interface{}) (int64, error)
	QueryNodeDescKey(id int64) (string, error)
	QueryNodeTypeByID(id int64) (int64, error)
	QueryNodeMetaByID(id int64) ([]byte, error)
	UpdateNodeMetaByID(id int64, meta []byte) error

	UpdateByID(id int64, value interface{}) error
	AppendByID(id int64, value interface{}) error // only for ConfigItemArray
}
