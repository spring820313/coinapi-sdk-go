#ifndef _btc_transaction_params_h
#define _btc_transaction_params_h

#include "define.h"
#include "bigint.h"
#include "bt_seed.h"
#include "btc_vout.h"

#ifdef __cplusplus
extern "C" {
#endif

	struct btc_transaction_params
	{
		btseed* seed;
		str_t* fromAddress;
		str_t* sendAddress;
		bigint_t* sendValue;
		bigint_t* feePerKb;
		str_t* password;
		str_t* priKey;
		my_bool recipientsPayFees;
		array_t* btcvoutFormList;
		int txType;
		array_t* data;
	};
	typedef struct btc_transaction_params btc_transaction_params;

	COIN_API_EXPORT void btc_transaction_params_init(struct btc_transaction_params* btp);
	my_bool btc_transaction_params_to_cpp(struct btc_transaction_params* btp, void* cpp);
	my_bool cpp_to_btc_transaction_params(void* cpp, struct btc_transaction_params* btp);
	COIN_API_EXPORT  void btc_transaction_params_free(struct btc_transaction_params* btp);


#ifdef __cplusplus
}
#endif

#endif
