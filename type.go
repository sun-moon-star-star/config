package config

type ConfigItem uint8

var ConfigItemBool = ConfigItem(0)
var ConfigItemInt32 = ConfigItem(1)
var ConfigItemUint32Item = ConfigItem(2)
var ConfigItemInt64 = ConfigItem(3)
var ConfigItemUint64 = ConfigItem(4)
var ConfigItemFloat32 = ConfigItem(5)
var ConfigItemFloat64 = ConfigItem(6)
var ConfigItemByte = ConfigItem(7)
var ConfigItemBytes = ConfigItem(8)
var ConfigItemArray = ConfigItem(9)

var ConfigItemCount = uint8(12)

type ConfigItemBoolType bool
type ConfigItemInt32Type int32
type ConfigItemUint32ItemType uint32
type ConfigItemInt64Type int64
type ConfigItemUint64Type uint64
type ConfigItemFloat32Type float32
type ConfigItemFloat64Type float64
type ConfigItemByteType byte
type ConfigItemBytesType []byte
type ConfigItemArrayType []interface{}
