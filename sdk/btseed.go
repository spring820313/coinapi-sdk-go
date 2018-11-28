package sdk
/*
#cgo CFLAGS: -I../include/
#include "uclib/array.h"
#include "uclib/value.h"
#include "bt_seed.h"
*/
import "C"
import "unsafe"

type BtSeed struct {
	Seed []byte
	MnemonicCode []string
	EncryptedMnemonicCode EncryptedData
	EncryptedSeed EncryptedData
	CreationTimeSeconds int32
	PwdHash string
	RandomSalt []byte
}

func (self *BtSeed) Init()  {
	self.CreationTimeSeconds = 0
	self.PwdHash = ""
}

func (self BtSeed) ToC() (*C.struct_btseed) {
	sizeof_btseed := unsafe.Sizeof(C.struct_btseed{})
	bs := (*C.struct_btseed)(C.malloc(C.size_t(sizeof_btseed)))
	C.btseed_init(bs)

	seed_array := C.array_create()
	for _, v := range self.Seed {
		var vt C.struct__value_t
		C.value_set_uint8(&vt, C.uint8_t(v))
		C.array_append(seed_array, vt);
	}
	bs.seed = seed_array

	mne_array := C.array_create()
	for _, v := range self.MnemonicCode {
		var vt C.struct__value_t

		vc := C.CString(v)
		size := len(v)
		word := C.str_create(vc, C.uint32_t(size), 0)
		C.value_set_str(&vt, word)
		C.array_append(mne_array, vt);
	}
	bs.mnemonicCode = mne_array

	bs.encryptedMnemonicCode = *self.EncryptedMnemonicCode.ToC()
	bs.encryptedSeed = *self.EncryptedSeed.ToC()
	bs.creationTimeSeconds = C.long(self.CreationTimeSeconds)
	if self.PwdHash != "" {
		pwdhashc := C.CString(self.PwdHash)
		size := len(self.PwdHash)
		bs.pwdhash = C.str_create(pwdhashc, C.uint32_t(size), 0)
	}

	salt_array := C.array_create()
	for _, v := range self.RandomSalt {
		var vt C.struct__value_t
		C.value_set_uint8(&vt, C.uint8_t(v))
		C.array_append(salt_array, vt);
	}
	bs.randomSalt = salt_array

	return bs
}

func (self *BtSeed) FromC(bs *C.struct_btseed) bool {
	if bs.seed != nil {
		seed_size := C.array_size(bs.seed)
		self.Seed = make([]byte, 0)
		for i := 0; i < int(seed_size); i++ {
			vt := C.array_get(bs.seed, C.uint32_t(i))
			v := C.value_uint8(vt)
			self.Seed = append(self.Seed, byte(v))
		}
	}

	if bs.mnemonicCode != nil {
		mne_size := C.array_size(bs.mnemonicCode)
		self.MnemonicCode = make([]string, 0)
		for i := 0; i < int(mne_size); i++ {
			vt := C.array_get(bs.mnemonicCode, C.uint32_t(i))
			v := C.value_str(vt)

			word := C.GoString(v.str)
			self.MnemonicCode = append(self.MnemonicCode, word)
		}
	}

	self.EncryptedMnemonicCode.FromC(&bs.encryptedMnemonicCode)
	self.EncryptedSeed.FromC(&bs.encryptedSeed)

	self.CreationTimeSeconds = int32(bs.creationTimeSeconds)
	if bs.pwdhash != nil {
		self.PwdHash = C.GoString(bs.pwdhash.str)
	}

	if bs.randomSalt != nil {
		salt_size := C.array_size(bs.randomSalt)
		self.RandomSalt = make([]byte, 0)
		for i := 0; i < int(salt_size); i++ {
			vt := C.array_get(bs.randomSalt, C.uint32_t(i))
			v := C.value_uint8(vt)
			self.RandomSalt = append(self.RandomSalt, byte(v))
		}
	}

	return true
}
