package sdk
/*
#cgo CFLAGS: -I../include/
#include "uclib/str.h"
#include "key_path.h"
*/
import "C"
import (
	"fmt"
	"unsafe"
)

type KeyPath struct {
	Path1 int32
	Path2 int32
	Path3 int32
	Path4 int32
	Path5 int32
	Hd1 bool
	Hd2 bool
	Hd3 bool
	Hd4 bool
	Hd5 bool
	Symbol string
}

func (self *KeyPath) Init()  {
	self.Path1 = -1
	self.Path2 = -1
	self.Path3 = -1
	self.Path4 = -1
	self.Path5 = -1
	self.Hd1 = false
	self.Hd2 = false
	self.Hd3 = false
	self.Hd4 = false
	self.Hd5 = false
	self.Symbol = ""
}

func (self KeyPath) ToC() (*C.struct_keypath)  {
	sizeof_keypath := unsafe.Sizeof(C.struct_keypath{})
	kp := (*C.struct_keypath)(C.malloc(C.size_t(sizeof_keypath)))
	C.keypath_init(kp)
	kp.path1 = C.int(self.Path1)
	kp.path2 = C.int(self.Path2)
	kp.path3 = C.int(self.Path3)
	kp.path4 = C.int(self.Path4)
	kp.path5 = C.int(self.Path5)
	if self.Hd1 {
		kp.hd1 = C.my_bool(1)
	} else {
		kp.hd1 = C.my_bool(0)
	}
	if self.Hd2 {
		kp.hd2 = C.my_bool(1)
	} else {
		kp.hd2 = C.my_bool(0)
	}
	if self.Hd3 {
		kp.hd3 = C.my_bool(1)
	} else {
		kp.hd3 = C.my_bool(0)
	}
	if self.Hd4 {
		kp.hd4 = C.my_bool(1)
	} else {
		kp.hd4 = C.my_bool(0)
	}
	if self.Hd5 {
		kp.hd5 = C.my_bool(1)
	} else {
		kp.hd5 = C.my_bool(0)
	}
	if self.Symbol != "" {
		symbolc := C.CString(self.Symbol)
		size := len(self.Symbol)
		kp.symbol = C.str_create(symbolc, C.uint32_t(size), C.uint32_t(size))
	}
	fmt.Println(kp)
	return kp
}

func FreeKeyPathC(kp *C.struct_keypath)  {
	if kp == nil {
		return
	}
	C.str_unref(kp.symbol)
	C.free(unsafe.Pointer(kp))
}
