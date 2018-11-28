#ifndef _net_params_h_
#define _net_params_h_

#include <stdint.h>
#include "define.h"
#include "key_path.h"

#ifdef __cplusplus
extern "C" {
#endif

	struct netparams
	{
		str_t* symbol;
		coin_type coinType;
		network_type nettype;
		struct keypath keyPath;
		uint32_t version;

		uint32_t HDprivate;
		uint32_t HDpublic;
		uint32_t P2KH;
		uint32_t P2SH;
		uint8_t keyprefixes;

		uint16_t ApiVersion;

		uint32_t N;
		uint32_t R;
		uint32_t P;
	};
	typedef struct netparams netparams;

	COIN_API_EXPORT void netparams_init(struct netparams* np);
	my_bool netparams_to_cpp(struct netparams* np, void* cpp);
	COIN_API_EXPORT void netparams_free(struct netparams* np);


#ifdef __cplusplus
}
#endif

#endif
