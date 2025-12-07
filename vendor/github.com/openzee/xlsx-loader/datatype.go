package loader

import log "github.com/sirupsen/logrus"

type DataType int

const (
	UnsupportedType DataType = iota
	NoneType                 // = "NONE"
	Float32Type              // = "REAL"
	Float64Type              // = "DOUBLE"
	Int32Type                // = "INT32"
	Uint32Type               // = "UINT32"
	Int16Type                // = "INT16"
	Uint16Type               // = "UINT16"
	Int64Type                // = "INT64"
	Uint64Type               // = "UINT64"
	BoolType                 // = "BOOL"
	StringType               // = "STRING"
)

func ParseDataType(s string) DataType {

	if s == "NONE" {
		return NoneType
	}

	if s == "REAL" {
		return Float32Type
	}

	if s == "DOUBLE" {
		return Float64Type
	}

	if s == "INT32" {
		return Int32Type
	}

	if s == "INT64" {
		return Int64Type
	}

	if s == "UINT64" {
		return Uint64Type
	}

	if s == "INT16" {
		return Int16Type
	}

	if s == "UINT16" {
		return Uint16Type
	}

	if s == "BOOL" {
		return BoolType
	}

	if s == "STRING" {
		return StringType
	}

	log.Fatalf("不支持的数据类型:%v", s )

	return UnsupportedType
}
