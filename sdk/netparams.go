package sdk
/*
#cgo CFLAGS: -I../include/
#include "uclib/str.h"
#include "net_params.h"
*/
import "C"
import "unsafe"

type NetParams struct {
	Symbol string
	CoinType C.enum_coin_type
	NetworkType C.enum_network_type
	KeyPath KeyPath
	Version uint32
	HDprivate uint32
	HDpublic uint32
	P2KH uint32
	P2SH uint32
	KeyPrefixes uint8
	ApiVersion uint16
	N uint32
	R uint32
	P uint32
}

func (self *NetParams) Init()  {
	self.KeyPath.Init()
}

func (self NetParams) ToC() (*C.struct_netparams) {
	sizeof_netparams := unsafe.Sizeof(C.struct_netparams{})
	np := (*C.struct_netparams)(C.malloc(C.size_t(sizeof_netparams)))
	C.netparams_init(np)

	if self.Symbol != "" {
		symbolc := C.CString(self.Symbol)
		size := len(self.Symbol)
		np.symbol = C.str_create(symbolc, C.uint32_t(size), C.uint32_t(size))
	}
	np.coinType = C.coin_type(self.CoinType)
	np.nettype = C.network_type(self.NetworkType)
	np.keyPath = *self.KeyPath.ToC()
	np.version = C.uint32_t(self.Version)
	np.HDprivate = C.uint32_t(self.HDprivate)
	np.HDpublic = C.uint32_t(self.HDpublic)
	np.P2KH = C.uint32_t(self.P2KH)
	np.P2SH = C.uint32_t(self.P2SH)
	np.keyprefixes = C.uint8_t(self.KeyPrefixes)
	np.ApiVersion = C.uint16_t(self.ApiVersion)
	np.N = C.uint32_t(self.N)
	np.P = C.uint32_t(self.P)
	np.R = C.uint32_t(self.R)

	return np
}