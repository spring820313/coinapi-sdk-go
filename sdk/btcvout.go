package sdk
/*
#cgo CFLAGS: -I../include/
#include "uclib/str.h"
#include "bigint.h"
#include "btc_vout.h"
*/
import "C"
import (
	"math/big"
	"unsafe"
)

type BtcVout struct {
	Value big.Int
	Hash string
	N int32
	CoinBase bool
}

func (self BtcVout) ToC() (*C.struct_btc_vout)  {
	sizeof_btc_vout := unsafe.Sizeof(C.struct_btc_vout{})
	bv := (*C.struct_btc_vout)(C.malloc(C.size_t(sizeof_btc_vout)))
	C.btc_vout_init(bv)

	valueStr := self.Value.String()
	bv.value = C.bigint_from_string(C.CString(valueStr), 10)

	if self.Hash != "" {
		hashc := C.CString(self.Hash)
		size := len(self.Hash)
		bv.hash = C.str_create(hashc, C.uint32_t(size), 0)
	}

	bv.n = C.int(self.N)
	if self.CoinBase {
		bv.coinBase = C.my_bool(1)
	} else {
		bv.coinBase = C.my_bool(0)
	}

	return bv
}
