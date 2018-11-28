package sdk

/*
#cgo CFLAGS: -I../include/
#cgo LDFLAGS: -L../ -lcoinapi-nojni
#include "uclib/value.h"
#include "coinapi.h"
*/
import "C"
import (
	"fmt"
	"unsafe"
	"strings"
)

type CoinApi struct {
	h unsafe.Pointer
}

func (self *CoinApi) Init() bool  {
	self.h  =  unsafe.Pointer(C.init())
	if self.h != nil {
		return true
	}
	return false
}

func (self CoinApi) Cleanup() {
	if self.h == nil {
		return;
	}
	C.cleanup(C.HUNIT(self.h))
}

func (self CoinApi) CreateAllCoinMnemonicCode() ([]string, bool) {
	if self.h == nil {
		return nil, false;
	}
	mne := C.createAllCoinMnemonicCode(C.HUNIT(self.h))

	if mne != nil {
		defer C.free(unsafe.Pointer(mne))
	}

	mnego := C.GoString(mne)
	fmt.Println(mnego)

	words := strings.Split(mnego, " ")
	if len(words) != 12 {
		return nil, false
	}
	return words, true
}

func (self CoinApi) CheckMnemonicCode(words []string) bool  {
	if self.h == nil {
		return false
	}

	if len(words) != 12 {
		return false
	}

	mne := strings.Join(words, " ")
	mnec := C.CString(mne)
	ret := C.checkMnemonicCode(C.HUNIT(self.h), mnec)
	if ret == 0 {
		return false
	}

	return true
}

func (self CoinApi) CreateWallet(words string, password string, netParams *NetParams) (*BtWallet, bool) {
	if self.h == nil {
		return nil, false
	}

	netparam := netParams.ToC()
	mne := C.CString(words)
	pwd := C.CString(password)

	var bw C.struct_btwallet
	ret := C.createWallet(C.HUNIT(self.h), mne, pwd, netparam, &bw);
	if ret == 0 {
		return nil, false
	}

	wallet := &BtWallet{}
	wallet.FromC(&bw)

	return wallet, true
}

func (self CoinApi) CreateSignTransaction(signParams *BtcTransactionParams, netParams *NetParams) map[string]string {
	ret := make(map[string]string)

	if self.h == nil {
		return ret
	}

	netparam := netParams.ToC()
	signparam := signParams.ToC()

	result := C.createSignTransaction(C.HUNIT(self.h), C.transaction_params(signparam), netparam)

	size := result.array.size
	p := unsafe.Pointer(result.array.data)
	fmt.Println(uintptr(p))

	for i := 0; i < int(size); i++ {
		v := (*C.struct__value_t)(p)
		vp := C.value_pointer(*v);
		kv := (*C.struct__key_value_t)(vp)

		key := ""
		if kv.key != nil {
			key = C.GoString(kv.key.str)
		}

		value := C.value_str(kv.value)
		val := C.GoString(value.str)

		ret[key] = val

		p = unsafe.Pointer(uintptr(p) + unsafe.Sizeof(C.struct__value_t{}))
	}

	return ret
}