// Copyright 2019, Chef.  All rights reserved.
// https://github.com/q191201771/lal
//
// Use of this source code is governed by a MIT-style license
// that can be found in the License file.
//
// Author: Chef (191201771@qq.com)

package rtmp

// amf0.go
// @pure
// 提供amf0格式的编码与解码的操作

import (
	"io"
)

const (
	Amf0TypeMarkerNumber      = uint8(0x00)
	Amf0TypeMarkerBoolean     = uint8(0x01)
	Amf0TypeMarkerString      = uint8(0x02)
	Amf0TypeMarkerObject      = uint8(0x03)
	Amf0TypeMarkerNull        = uint8(0x05)
	Amf0TypeMarkerUndefined   = uint8(0x06)
	Amf0TypeMarkerEcmaArray   = uint8(0x08)
	Amf0TypeMarkerObjectEnd   = uint8(0x09) // end for both Object and Array
	Amf0TypeMarkerStrictArray = uint8(0x0a)
	Amf0TypeMarkerLongString  = uint8(0x0c)
	Amf0TypeMarkerUnsupported = uint8(0x0d)

	// 还没用到的类型
	//Amf0TypeMarkerMovieclip   = uint8(0x04)
	//Amf0TypeMarkerReference   = uint8(0x07)
	//Amf0TypeMarkerData        = uint8(0x0b)
	//Amf0TypeMarkerRecordset   = uint8(0x0e)
	//Amf0TypeMarkerXmlDocument = uint8(0x0f)
	//Amf0TypeMarkerTypedObject = uint8(0x10)
)

var (
	// Amf0TypeMarkerObjectEndBytes Amf0TypeMarkerArrayEndBytes:
	// object-end-type(0x00 0x00 0x09) 表示Object和EcmaArray类型的结束标识
	Amf0TypeMarkerObjectEndBytes = []byte{0, 0, Amf0TypeMarkerObjectEnd}
	Amf0TypeMarkerArrayEndBytes  = []byte{0, 0, Amf0TypeMarkerObjectEnd}
)

// ---------------------------------------------------------------------------------------------------------------------

type ObjectPair struct {
	Key   string      // Amf0TypeMarkerStrictArray类型的数据，Key为""
	Value interface{} // TODO(chef): [perf] 考虑换成泛型 202206
}

type ObjectPairArray []ObjectPair

func (o ObjectPairArray) Find(key string) interface{} { _ = "STUB: not implemented"; return nil }

func (o ObjectPairArray) FindString(key string) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

func (o ObjectPairArray) FindNumber(key string) (int, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

func (o ObjectPairArray) DebugString() string { _ = "STUB: not implemented"; return "" }

// ---------------------------------------------------------------------------------------------------------------------

type amf0 struct{}

var Amf0 amf0

func (amf0) WriteNumber(writer io.Writer, val float64) error { _ = "STUB: not implemented"; return nil }

func (amf0) WriteString(writer io.Writer, val string) error { _ = "STUB: not implemented"; return nil }

func (amf0) WriteNull(writer io.Writer) error { _ = "STUB: not implemented"; return nil }

func (amf0) WriteBoolean(writer io.Writer, b bool) error { _ = "STUB: not implemented"; return nil }

func (amf0) WriteObject(writer io.Writer, opa ObjectPairArray) error {
	_ = "STUB: not implemented"
	return nil
}

// ----------------------------------------------------------------------------
// read类型的方法集合
//
// 从输入参数<b>切片中读取函数名所指定的amf类型数据
// 注意，方法内部不会修改输入参数<b>切片的内容
//
// 返回值如无特殊说明，则
// 第1个参数为读取出的所指定类型的数据
// 第2个参数为读取时从<b>消耗的字节大小
// 第3个参数error，如果不等于nil，表示读取失败

func (amf0) ReadStringWithoutType(b []byte) (string, int, error) {
	_ = "STUB: not implemented"
	return "", 0, nil
}

func (amf0) ReadLongStringWithoutType(b []byte) (string, int, error) {
	_ = "STUB: not implemented"
	return "", 0, nil
}

func (amf0) ReadString(b []byte) (val string, l int, err error) {
	_ = "STUB: not implemented"
	return "", 0, nil
}

func (amf0) ReadNumber(b []byte) (float64, int, error) { _ = "STUB: not implemented"; return 0, 0, nil }

func (amf0) ReadBoolean(b []byte) (bool, int, error) {
	_ = "STUB: not implemented"
	return false, 0, nil
}

func (amf0) ReadNull(b []byte) (int, error) { _ = "STUB: not implemented"; return 0, nil }

func (amf0) ReadUndefinedOrUnsupported(b []byte) (int, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

// ReadObject
//
// @return ObjectPairArray: ...
// @return int: 读取时从 b 消耗的字节大小
// @return error: ...
func (amf0) ReadObject(b []byte) (ObjectPairArray, int, error) {
	_ = "STUB: not implemented"
	return *new(ObjectPairArray), 0, nil
}

// TODO chef: 实现WriteArray

// ReadArray Amf0TypeMarkerEcmaArray
func (amf0) ReadArray(b []byte) (ObjectPairArray, int, error) {
	_ = "STUB: not implemented"
	return *new(ObjectPairArray), 0, nil
}

func (amf0) ReadStrictArray(b []byte) (ObjectPairArray, int, error) {
	_ = "STUB: not implemented"
	return *new(ObjectPairArray), 0, nil
}

func (amf0) ReadObjectOrArray(b []byte) (ObjectPairArray, int, error) {
	_ = "STUB: not implemented"
	return *new(ObjectPairArray), 0, nil
}

func (amf0) read(b []byte, index int, k string, ops ObjectPairArray) (ObjectPairArray, int, error) {
	_ = "STUB: not implemented"
	return *new(ObjectPairArray), 0, nil
}
