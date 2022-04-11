// Copyright 2022 The yyjson-go Authors
// SPDX-License-Identifier: BSD-3-Clause

package yyjson

import (
	"fmt"
	"unsafe"

	"modernc.org/libc"

	"github.com/zchee/yyjson-go/internal/rt"
	cyyjson "github.com/zchee/yyjson-go/internal/yyjson"
)

const (
	nullUintptr   = uintptr(0)
	SizeofUintptr = unsafe.Sizeof(nullUintptr)

	// PtrSize is the size of a pointer in bytes. unsafe.Sizeof(uintptr(0)) but as an ideal constant.
	// It is also the size of the machine's native word size (that is, 4 on 32-bit systems, 8 on 64-bit).
	PtrSize = 4 << (^nullUintptr >> 63)
)

// ReadFlag represents an options for JSON reader.
type ReadFlag = cyyjson.READ_FLAG

// List of ReadCode.
const (
	// ReadNoflag default option (RFC 8259 compliant):
	//
	// - Read positive integer as uint64_t.
	// - Read negative integer as int64_t.
	// - Read floating-point number as double with correct rounding.
	// - Read integer which cannot fit in uint64_t or int64_t as double.
	// - Report error if real number is infinity.
	// - Report error if string contains invalid UTF-8 character or BOM.
	// - Report error on trailing commas, comments, inf and nan literals.
	ReadNoflag = cyyjson.READ_NOFLAG

	// ReadInsitu reads the input data in-situ.
	//
	// This option allows the reader to modify and use input data to store string
	// values, which can increase reading speed slightly.
	// The caller should hold the input data before free the document.
	// The input data must be padded by at least `YYJSON_PADDING_SIZE` byte.
	// For example: "[1,2]" should be "[1,2]\0\0\0\0", length should be 5.
	ReadInsitu = cyyjson.READ_INSITU

	// ReadStopWhenDone stops when done instead of issues an error if there's additional content
	// after a JSON document. This option may used to parse small pieces of JSON
	// in larger data, such as NDJSON.
	ReadStopWhenDone = cyyjson.READ_STOP_WHEN_DONE

	// ReadAllowTrailingCommas allows single trailing comma at the end of an object or array,
	// such as [1,2,3,] {"a":1,"b":2,}.
	ReadAllowTrailingCommas = cyyjson.READ_ALLOW_TRAILING_COMMAS

	// ReadAllowComments allows C-style single line and multiple line comments.
	ReadAllowComments = cyyjson.READ_ALLOW_COMMENTS

	// ReadAllowInfAndNan allows inf/nan number and literal, case-insensitive,
	// such as 1e999, NaN, inf, -Infinity.
	ReadAllowInfAndNan = cyyjson.READ_ALLOW_INF_AND_NAN
)

// ReadCode represents a result code for JSON reader.
type ReadCode = cyyjson.READ_CODE

// List of ReadCode.
const (
	// ReadErrorInvalidParameter invalid parameter, such as NULL string or invalid file path.
	ReadErrorInvalidParameter = cyyjson.READ_ERROR_INVALID_PARAMETER

	// ReadErrorMemoryAllocation memory allocation failure occurs.
	ReadErrorMemoryAllocation = cyyjson.READ_ERROR_MEMORY_ALLOCATION

	// ReadErrorEmptyContent input JSON string is empty.
	ReadErrorEmptyContent = cyyjson.READ_ERROR_EMPTY_CONTENT

	// ReadErrorUnexpectedContent unexpected content after document, such as "[1]#".
	ReadErrorUnexpectedContent = cyyjson.READ_ERROR_UNEXPECTED_CONTENT

	// ReadErrorUnexpectedEnd unexpected ending, such as "[123".
	ReadErrorUnexpectedEnd = cyyjson.READ_ERROR_UNEXPECTED_END

	// ReadErrorUnexpectedCharacter unexpected character inside the document, such as "[#]".
	ReadErrorUnexpectedCharacter = cyyjson.READ_ERROR_UNEXPECTED_CHARACTER

	// ReadErrorJSONStructure invalid JSON structure, such as "[1,]".
	ReadErrorJSONStructure = cyyjson.READ_ERROR_JSON_STRUCTURE

	// ReadErrorInvalidComment invalid comment, such as unclosed multi-line comment.
	ReadErrorInvalidComment = cyyjson.READ_ERROR_INVALID_COMMENT

	// ReadErrorInvalidNumber invalid number, such as "123.e12", "000".
	ReadErrorInvalidNumber = cyyjson.READ_ERROR_INVALID_NUMBER

	// ReadErrorInvalidString invalid string, such as invalid escaped character inside a string.
	ReadErrorInvalidString = cyyjson.READ_ERROR_INVALID_STRING

	// ReadErrorLiteral invalid JSON literal, such as "truu".
	ReadErrorLiteral = cyyjson.READ_ERROR_LITERAL

	// ReadErrorFileOpen failed to open a file.
	ReadErrorFileOpen = cyyjson.READ_ERROR_FILE_OPEN

	// ReadErrorFileRead failed to read a file.
	ReadErrorFileRead = cyyjson.READ_ERROR_FILE_READ
)

// ReadError represents an error of read JSON.
type ReadError cyyjson.Yyjson_read_err

// Error returns a string representation of the ReadError.
func (e ReadError) Error() string {
	return fmt.Sprintf("code: %d, message: %s", ReadCode(e.Code), libc.GoString(e.Msg))
}

// UnmarshalOptions configures the unmarshaler.
type UnmarshalOptions struct {
	tls *libc.TLS

	Flag     cyyjson.Yyjson_read_flag
	Allocate uintptr
	ReadErr  uintptr
}

// Unmarshal parses the JSON-encoded data and stores the result
// in the value pointed to by v. If v is nil or not a pointer,
// Unmarshal returns an ReadError.
func Unmarshal(data []byte, v interface{}) error {
	opt := &UnmarshalOptions{
		tls: libc.NewTLS(),
	}
	defer opt.tls.Close()

	return opt.Unmarshal(data, v)
}

// Unmarshal parses the JSON-encoded data and stores the result
// in the value pointed to by v. If v is nil or not a pointer,
// Unmarshal returns an ReadError.
func (o UnmarshalOptions) Unmarshal(data []byte, v interface{}) error {
	d := string(data)
	s := cString(d)
	ret := cyyjson.Yyjson_read_opts(o.tls, s, uint64(len(d)), o.Flag, o.Allocate, o.ReadErr)
	if o.ReadErr != nullUintptr {
		return *(**ReadError)(unsafe.Pointer(&o.ReadErr))
	}

	doc := *(**cyyjson.Yyjson_doc)(unsafe.Pointer(&ret))
	cyyjson.Dump(doc)

	vv := rt.UnpackEface(v)
	cyyjson.Dump(vv)
	vp := vv.Value
	cyyjson.Dump(vp)

	return nil
}

// ReadFile reads a JSON file.
//
// This function is thread-safe if you make sure that:
// 1. The file is not modified by other threads.
// 2. The `alc` is thread-safe or NULL.
//
// path is the JSON file's path. If you pass an invalid path, you will get NULL result.
//
// flg is the JSON read options. You can combine multiple options using bitwise `|` operator.
//
// alc is the memory allocator used by JSON reader. Pass NULL to use the libc's default allocator (thread-safe).
//
// err is a pointer to receive error information. Pass NULL if you don't need error information.
//
// Return a new JSON document, or NULL if error occurs. You should use yyjson_doc_free() to release it when it's no longer needed.
func ReadFile(path string) error {
	opt := UnmarshalOptions{tls: libc.NewTLS()}
	defer opt.tls.Close()

	return opt.ReadFile(path)
}

// ReadFile reads a JSON file.
//
// This function is thread-safe if you make sure that:
// 1. The file is not modified by other threads.
// 2. The `alc` is thread-safe or NULL.
//
// path is the JSON file's path. If you pass an invalid path, you will get NULL result.
//
// flg is the JSON read options. You can combine multiple options using bitwise `|` operator.
//
// alc is the memory allocator used by JSON reader. Pass NULL to use the libc's default allocator (thread-safe).
//
// err is a pointer to receive error information. Pass NULL if you don't need error information.
//
// Return a new JSON document, or NULL if error occurs. You should use yyjson_doc_free() to release it when it's no longer needed.
func (o *UnmarshalOptions) ReadFile(path string) error {
	file := cString(path)

	doc := cyyjson.Yyjson_read_file(o.tls, file, ReadNoflag, nullUintptr, o.ReadErr)
	if o.ReadErr != nullUintptr {
		return *(**ReadError)(unsafe.Pointer(&o.ReadErr))
	}

	root := cyyjson.Yyjson_doc_get_root(o.tls, doc)

	titleName := cString("title")
	title := cyyjson.Yyjson_obj_get(o.tls, root, titleName)

	str := cyyjson.Yyjson_get_str(o.tls, title)
	cyyjson.Dump(libc.GoString(str))

	pentries := cString("entries")
	entries := cyyjson.Yyjson_obj_get(o.tls, root, pentries)

	ArrForeach(o.tls, entries, func(idx uint64, val uintptr) {
		ptitle := cString("title")
		title := cyyjson.Yyjson_obj_get(o.tls, val, ptitle)
		if title != 0 {
			cyyjson.Dump(libc.GoString(cyyjson.Yyjson_get_str(o.tls, title)))
		}
	})

	return nil
}

// ForeachFunc type of ArrForeach arg.
type ForeachFunc func(idx uint64, val uintptr)

// ArrForeach is the iterating over an array.
func ArrForeach(tls *libc.TLS, arr uintptr, fn ForeachFunc) {
	var val uintptr
	for idx := uint64(0); idx < cyyjson.Yyjson_arr_size(tls, arr); idx++ {
		if idx == 0 {
			val = cyyjson.Yyjson_arr_get_first(tls, arr)
		}

		fn(idx, val)

		val = cyyjson.Unsafe_yyjson_get_next(tls, val)
	}
}
