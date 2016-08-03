package btos

import (
	"reflect"
	"unsafe"
)

func BytesToString(b []byte) string {
	bytesHeader := (*reflect.SliceHeader)(unsafe.Pointer(&b))
	strHeader := reflect.StringHeader{
		bytesHeader.Data,
		bytesHeader.Len,
	}
	return *(*string)(unsafe.Pointer(&strHeader))
}

func StringToBytes(s string) []byte {
	strHeader := (*reflect.StringHeader)(unsafe.Pointer(&s))
	bytesHeader := reflect.SliceHeader{
		strHeader.Data,
		strHeader.Len,
		strHeader.Len,
	}
	return *(*[]byte)(unsafe.Pointer(&bytesHeader))
}
