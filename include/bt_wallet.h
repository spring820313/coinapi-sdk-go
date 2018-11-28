#pragma once
#ifndef _bt_wallet_h_
#define _bt_wallet_h_

#include "define.h"
#include "bt_seed.h"


#ifdef __cplusplus
extern "C" {
#endif

	struct btwallet
	{
		btseed* btSeed;
		str_t *pubkey;
		str_t *address;
		str_t *symbol;
	};
	typedef struct btwallet btwallet;

	COIN_API_EXPORT void btwallet_init(struct btwallet* bw);
	my_bool btwallet_to_cpp(struct btwallet* bw, void* cpp);
	my_bool cpp_to_btwallet(void* cpp, struct btwallet* bw);
	COIN_API_EXPORT void btwallet_free(struct btwallet* bw);


#ifdef __cplusplus
}
#endif

#endif