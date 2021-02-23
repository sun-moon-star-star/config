package storage

type Node struct {
	ID       int64 `json:"id"` // unique ID
	FatherID int64 `json:"father_id"`

	Name     string `json:"name"`
	FullPath string `json:"full_path"`

	CreatedTime uint64 `json:"created_time"` // nano
	UpdatedTime uint64 `json:"updated_time"` // nano

	Type int64  `json:"type"`    // last four bits is ConfigItem, first 60 bits no use
	Meta []byte `json:"extends"` // user defined data

	Value interface{} `json:"value"`
}

type StorageAPI interface {
	Create(*Node) error
	Read(*Node) error
	Update(*Node) error
	Delete(*Node) error
}
