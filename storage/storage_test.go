package storage

import (
	"config"
	"time"
)

type StorageTest struct {
	root Node
}

func (s *StorageTest) Init() {
	s.root = Node{
		ID:       0,
		FatherID: -1,

		Name:     "root",
		FullPath: "/",

		CreatedTime: time.Now().UnixNano(),
		UpdatedTime: time.Now().UnixNano(),

		Type: config.ConfigItemBytes,
		Meta: []byte(`{"name":"zhaolu","age": 12}`),

		Value: "hello world",
	}
}

func (*StorageTest) Create(*Node) error {

	return nil
}

func (*StorageTest) Read(*Node) error {
	return nil
}

func (*StorageTest) Update(*Node) error {
	return nil
}

func (*StorageTest) Delete(*Node) error {
	return nil
}
