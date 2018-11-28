#ifndef _key_path_h_
#define _key_path_h_

#include "define.h"

#ifdef __cplusplus
extern "C" {
#endif

	struct keypath
	{
		int path1;
		int path2;
		int path3;
		int path4;
		int path5;
		
		my_bool hd1;
		my_bool hd2;
		my_bool hd3;
		my_bool hd4;
		my_bool hd5;

		str_t *symbol;
	};
	typedef struct keypath keypath;

	COIN_API_EXPORT void keypath_init(struct keypath* kp);
	my_bool keypath_to_cpp(struct keypath* kp, void* cpp);
	COIN_API_EXPORT void keypath_free(struct keypath* kp);


#ifdef __cplusplus
}
#endif

#endif