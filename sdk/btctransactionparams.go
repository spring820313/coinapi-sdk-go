package sdk
/*
#cgo CFLAGS: -I../include/
#include "uclib/str.h"
#include "bigint.h"
#include "btc_vout.h"
#include "btc_transaction_params.h"

void btc_vout_destroy(void* data) {
	if (data) {
		btc_vout_free((struct btc_vout*)data);
		free(data);
	}
}

*/
import "C"
import (
	"math/big"
	"unsafe"
)

type BtcTransactionParams struct {
	Seed *BtSeed
	FromAddress string
	SendAddress string
	SendValue big.Int
	FeePerKb big.Int
	Password string
	PriKey string
	RecipientsPayFees bool
	BtcVoutFormList []BtcVout
	TxType int
	Data []byte
}

func (self BtcTransactionParams) ToC() (*C.struct_btc_transaction_params) {
	sizeof_btc_transaction_params := unsafe.Sizeof(C.struct_btc_transaction_params{})
	btp := (*C.struct_btc_transaction_params)(C.malloc(C.size_t(sizeof_btc_transaction_params)))
	C.btc_transaction_params_init(btp)

	if self.Seed != nil {
		btp.seed = self.Seed.ToC()
	}

	if self.FromAddress != "" {
		fromAddressc := C.CString(self.FromAddress)
		size := len(self.FromAddress)
		btp.fromAddress = C.str_create(fromAddressc, C.uint32_t(size), 0)
	}

	if self.SendAddress != "" {
		sendAddressc := C.CString(self.SendAddress)
		size := len(self.SendAddress)
		btp.sendAddress = C.str_create(sendAddressc, C.uint32_t(size), 0)
	}

	sendValue := self.SendValue.String()
	btp.sendValue = C.bigint_from_string(C.CString(sendValue), 10)

	feePerKb := self.FeePerKb.String()
	btp.feePerKb = C.bigint_from_string(C.CString(feePerKb), 10)

	if self.Password != "" {
		passwordc := C.CString(self.Password)
		size := len(self.Password)
		btp.password = C.str_create(passwordc, C.uint32_t(size), 0)
	}

	if self.PriKey != "" {
		prikeyc := C.CString(self.PriKey)
		size := len(self.PriKey)
		btp.priKey = C.str_create(prikeyc, C.uint32_t(size), 0)
	}

	if self.RecipientsPayFees {
		btp.recipientsPayFees = C.my_bool(1)
	} else {
		btp.recipientsPayFees = C.my_bool(0)
	}

	btp.btcvoutFormList = C.array_create()
	for _, vout := range self.BtcVoutFormList{
		voutc := vout.ToC()

		var vt C.struct__value_t
		C.value_set_pointer(&vt, unsafe.Pointer(voutc), C.destroy_t(unsafe.Pointer(C.btc_vout_destroy)))
		C.array_append(btp.btcvoutFormList, vt)
	}

	btp.txType = C.int(self.TxType)

	iv_data := C.array_create()
	for _, v := range self.Data {
		var vt C.struct__value_t
		C.value_set_uint8(&vt, C.uint8_t(v))
		C.array_append(iv_data, vt);
	}
	btp.data = iv_data

	return btp
}