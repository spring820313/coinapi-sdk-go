#ifndef _bt_seed_h_
#define _bt_seed_h_


#include "define.h"
#include "encrypted_data.h"


#ifdef __cplusplus
extern "C" {
#endif

	struct btseed
	{
		array_t* seed;
		array_t* mnemonicCode;
		encrypted_data encryptedMnemonicCode;
		encrypted_data encryptedSeed;
		long creationTimeSeconds;
		str_t* pwdhash;
		array_t* randomSalt;
	};
	typedef struct btseed btseed;

	COIN_API_EXPORT void btseed_init(struct btseed* bs);
	COIN_API_EXPORT my_bool btseed_to_cpp(struct btseed* bs, void* cpp);
	my_bool cpp_to_btseed(void* cpp, struct btseed* bs);
	COIN_API_EXPORT  void btseed_free(struct btseed* bs);


#ifdef __cplusplus
}
#endif

#endif
