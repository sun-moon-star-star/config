package config

type Item interface{}

type ItemBool bool

type ItemInt32 int32
type ItemUint32 uint32
type ItemInt64 int64
type ItemUint64 uint64

type ItemFloat32 float32
type ItemFloat64 float64

type ItemByte byte
type ItemBytes []byte

type ItemArray []interface{}
type ItemObject map[string]interface{}
