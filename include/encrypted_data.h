#ifndef _encrypted_data_h_
#define _encrypted_data_h_

#include "define.h"


#ifdef __cplusplus
extern "C" {
#endif

	struct encrypted_data
	{
		array_t* initialisationVector;
		array_t* encryptedBytes;
	};
	typedef struct encrypted_data encrypted_data;

	COIN_API_EXPORT void encrypted_data_init(struct encrypted_data* ed);
	my_bool encrypted_data_to_cpp(struct encrypted_data* ed, void* cpp);
	my_bool cpp_to_encrypted_data(void* cpp, struct encrypted_data* ed);
	COIN_API_EXPORT void encrypted_data_free(struct encrypted_data* ed);


#ifdef __cplusplus
}
#endif

#endif
