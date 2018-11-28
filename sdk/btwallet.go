package sdk
/*
#cgo CFLAGS: -I../include/
#include "uclib/array.h"
#include "uclib/value.h"
#include "bt_wallet.h"
*/
import "C"
import "unsafe"

type BtWallet struct {
	BtSeed *BtSeed
	Pubkey string
	Address string
	Symbol string
}

func (self BtWallet) ToC() (*C.struct_btwallet) {
	sizeof_btwallet := unsafe.Sizeof(C.struct_btwallet{})
	bt := (*C.struct_btwallet)(C.malloc(C.size_t(sizeof_btwallet)))
	C.btwallet_init(bt)

	if self.BtSeed != nil {
		bt.btSeed = self.BtSeed.ToC()
	}

	if self.Pubkey != "" {
		pubkeyc := C.CString(self.Pubkey)
		size := len(self.Pubkey)
		bt.pubkey = C.str_create(pubkeyc, C.uint32_t(size), 0)
	}

	if self.Address != "" {
		addressc := C.CString(self.Address)
		size := len(self.Address)
		bt.address = C.str_create(addressc, C.uint32_t(size), 0)
	}

	if self.Symbol != "" {
		symbolc := C.CString(self.Symbol)
		size := len(self.Symbol)
		bt.symbol = C.str_create(symbolc, C.uint32_t(size), 0)
	}

	return bt
}

func (self *BtWallet) FromC(bt *C.struct_btwallet) bool {
	if bt.btSeed != nil {
		self.BtSeed = &BtSeed{}
		self.BtSeed.FromC(bt.btSeed)
	}

	if bt.pubkey != nil {
		self.Pubkey = C.GoString(bt.pubkey.str)
	}

	if bt.address != nil {
		self.Address = C.GoString(bt.address.str)
	}

	if bt.symbol != nil {
		self.Symbol = C.GoString(bt.symbol.str)
	}
	return true
}