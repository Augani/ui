// 28 february 2014

package ui

import (
	"unsafe"
)

// #cgo LDFLAGS: -lobjc -framework Foundation
// #include <stdlib.h>
// #include "objc_darwin.h"
import "C"

func objc_getClass(class string) C.id {
	cclass := C.CString(class)
	defer C.free(unsafe.Pointer(cclass))

	return C.objc_getClass(cclass)
}

func sel_getUid(sel string) C.SEL {
	csel := C.CString(sel)
	defer C.free(unsafe.Pointer(csel))

	return C.sel_getUid(csel)
}

// Common Objective-C types and selectors.
var (
	_NSObject = objc_getClass("NSObject")
	_NSString = objc_getClass("NSString")

	_alloc = sel_getUid("alloc")
	_new = sel_getUid("new")
	_release = sel_getUid("release")
	_stringWithUTF8String = sel_getUid("stringWithUTF8String:")
	_UTF8String = sel_getUid("UTF8String")
)

// some helper functions

func objc_alloc(class C.id) C.id {
	return C.objc_msgSend_noargs(class, _alloc)
}

func toNSString(str string) C.id {
	cstr := C.CString(str)
	defer C.free(unsafe.Pointer(cstr))

	return C.objc_msgSend_str(_NSString,
		_stringWithUTF8String,
		cstr)
}

func fromNSString(str C.id) string {
	cstr := C.objc_msgSend_noargs(str, _UTF8String)
	return C.GoString((*C.char)(unsafe.Pointer(cstr)))
}
