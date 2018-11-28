#ifndef _btc_vout_h_
#define _btc_vout_h_

#include "define.h"
#include "bigint.h"


#ifdef __cplusplus
extern "C" {
#endif

	struct btc_vout
	{
		str_t* hash;
		bigint_t* value;
		int n;
		my_bool coinBase;
	};
	typedef struct btc_vout btc_vout;

	COIN_API_EXPORT void btc_vout_init(struct btc_vout* bv);
	my_bool btc_vout_to_cpp(struct btc_vout* bv, void* cpp);
	my_bool cpp_to_btc_vout(void* cpp, struct btc_vout* bv);
	COIN_API_EXPORT void btc_vout_free(struct btc_vout* bv);


#ifdef __cplusplus
}
#endif

#endif
