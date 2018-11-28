package sdk
/*
#cgo CFLAGS: -I../include/
#include "uclib/array.h"
#include "uclib/value.h"
#include "encrypted_data.h"
*/
import "C"
import "unsafe"

type EncryptedData struct {
	initialisationVector []byte
	encryptedBytes []byte
}

func (self EncryptedData) ToC() (*C.struct_encrypted_data) {
	sizeof_encrypted_data := unsafe.Sizeof(C.struct_encrypted_data{})
	ed := (*C.struct_encrypted_data)(C.malloc(C.size_t(sizeof_encrypted_data)))
	C.encrypted_data_init(ed)

	iv_array := C.array_create()
	for _, v := range self.initialisationVector {
		var vt C.struct__value_t
		C.value_set_uint8(&vt, C.uint8_t(v))
		C.array_append(iv_array, vt);
	}
	ed.initialisationVector = iv_array

	encrypted_array := C.array_create()
	for _, v := range self.encryptedBytes {
		var vt C.struct__value_t
		C.value_set_uint8(&vt, C.uint8_t(v))
		C.array_append(encrypted_array, vt);
	}
	ed.encryptedBytes = encrypted_array

	return ed
}

func (self *EncryptedData) FromC(ed *C.struct_encrypted_data) bool {
	if ed.initialisationVector != nil {
		iv_size := C.array_size(ed.initialisationVector)
		self.initialisationVector = make([]byte, 0)
		for i := 0; i < int(iv_size); i++ {
			vt := C.array_get(ed.initialisationVector, C.uint32_t(i))
			v := C.value_uint8(vt)
			self.initialisationVector = append(self.initialisationVector, byte(v))
		}
	}
	if ed.encryptedBytes != nil {
		encrypted_size := C.array_size(ed.encryptedBytes)
		self.encryptedBytes = make([]byte, 0)
		for i := 0; i < int(encrypted_size); i++ {
			vt := C.array_get(ed.encryptedBytes, C.uint32_t(i))
			v := C.value_uint8(vt)
			self.encryptedBytes = append(self.encryptedBytes, byte(v))
		}
	}
	return true
}