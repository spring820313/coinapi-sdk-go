#ifndef _define_h_
#define _define_h_


#if defined(_MSC_VER) 
#ifndef COIN_API_EXPORT
#ifdef COINAPI_DLL
#define COIN_API_EXPORT __declspec(dllexport)
#else
#define COIN_API_EXPORT __declspec(dllimport)
#endif
#endif
#else
#ifndef COIN_API_EXPORT
#define COIN_API_EXPORT extern
#endif
#endif

//#define LOGARITHMIC_GROWTH

#include <stdint.h>
#include "uclib/array_helper.h"
#include "uclib/value_helper.h"


#ifdef __cplusplus
extern "C" {
#endif

	typedef int my_bool;
#define TRUE 1
#define FALSE 0

	enum coin_type {
		CT_BTC = 1,
		CT_ETH,
		CT_WICC,
		CT_BTS,
		CT_AFT,
		CT_LTC,
		CT_SBTC,
		CT_DOGE,
		CT_ETC,
		CT_WBTC,
		CT_ZEC,
		CT_DSH,
		CT_BCH,
		CT_QTUM,
		CT_LBTC,
		CT_NEO,
		CT_GAS,
		CT_XZC,
		CT_USDT,
		CT_BCO,
		CT_BHD,
		CT_EOS
	};
    typedef enum coin_type coin_type;

	enum network_type {
		NT_MAIN = 1,
		NT_TEST = 2,
		NT_REGTEST = 3
	};
	typedef enum network_type network_type;

	enum tx_type {
		TT_NONE = 0,
		TT_LBTC_REGISTER,
		TT_LBTC_VOTE,
		TT_LBTC_CANCELVOTE,
		TT_QTUM_TOKEN_TRANSFER,
		TT_WICC_REGISTERACCOUNT,
		TT_WICC_COMMON,
		TT_WICC_TRANSFER_SPC,
		TT_WICC_BET
	};
	typedef enum tx_type tx_type;

	typedef char* coin_symbol;

	extern coin_symbol CS_BTC;
	extern coin_symbol CS_ETH;
	extern coin_symbol CS_WICC;
	extern coin_symbol CS_BTS;
	extern coin_symbol CS_AFT;
	extern coin_symbol CS_LTC;
	extern coin_symbol CS_SBTC;
	extern coin_symbol CS_DOGE;
	extern coin_symbol CS_ETC;
	extern coin_symbol CS_WBTC;
	extern coin_symbol CS_ZEC;
	extern coin_symbol CS_DSH;
	extern coin_symbol CS_BCH;
	extern coin_symbol CS_QTUM;
	extern coin_symbol CS_LBTC;
	extern coin_symbol CS_NEO;
	extern coin_symbol CS_GAS;
	extern coin_symbol CS_XZC;
	extern coin_symbol CS_USDT;
	extern coin_symbol CS_BCO;
	extern coin_symbol CS_BHD;
	extern coin_symbol CS_EOS;

#ifdef __cplusplus
}
#endif

#endif
