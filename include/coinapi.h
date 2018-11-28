// coinapi.cpp: 定义 DLL 应用程序的导出函数。
//

#ifndef _coin_api_h_
#define _coin_api_h_

#include "define.h"
#include "bt_wallet.h"
#include "net_params.h"
#include "uclib/map.h"


typedef void* HUNIT;
typedef void* transaction_params;

#ifdef __cplusplus
extern "C" {
#endif

	/*
	库初始化
	*/
	COIN_API_EXPORT HUNIT init();

	/*
	获取助记词列表(要释放内存)
	*/
	COIN_API_EXPORT char* createAllCoinMnemonicCode(HUNIT h);

	/*
	检查助记词列表
	*/
	COIN_API_EXPORT int checkMnemonicCode(HUNIT h, char* pMne);

	/*
	一次创建一个地址
	*/
	COIN_API_EXPORT int createWallet(HUNIT h, char* words, char* password, netparams* netParams, btwallet* bw);

	/*
	根据 加密种子获取 解密后的私钥(要释放内存)
	*/
	COIN_API_EXPORT char* getPriKeyFromBtSeed(HUNIT h, btseed *btSeed, char* password, netparams* netParams);

	/*
	根据 加密种子获取 解密后的助记词(要释放内存)
	*/
	COIN_API_EXPORT char* getMnemonicCodeFromBtSeed(HUNIT h, btseed *btSeed, char* password, netparams* netParams);

	/*
	创建签名交易
	*/
	COIN_API_EXPORT map_t* createSignTransaction(HUNIT h, transaction_params signParams, netparams* netParams);

	/*
	库销毁
	*/
	COIN_API_EXPORT void cleanup(HUNIT h);


#ifdef __cplusplus
}
#endif

#endif


