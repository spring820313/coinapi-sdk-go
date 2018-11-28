package main
/*
#cgo CFLAGS: -I./include/
#include "define.h"
*/
import "C"

import "coinapi-sdk/sdk"

import (
	"fmt"
	"strings"
	"math/big"
)

func main()  {
	bigint := big.Int{}
	bigint.SetString("12345678901234567890", 10)
	bigstr := bigint.String()
	fmt.Println(bigstr)

	coinapi := sdk.CoinApi {}
	ret := coinapi.Init()

	words, ret := coinapi.CreateAllCoinMnemonicCode()
	fmt.Println(words)

	ret = coinapi.CheckMnemonicCode(words)
	fmt.Println(ret)

	keypath := sdk.KeyPath{
		Path1: 44,
		Path2: 133,
		Path3: 0,
		Path4: 0,
		Path5: 0,
		Hd1: true,
		Hd2: true,
		Hd3: true,
		Hd4: false,
		Hd5: false,
		Symbol: "BTC",
	}

	keypathc := keypath.ToC()
	fmt.Println(keypathc)

	netparams := sdk.NetParams{
		Symbol: "BTC",
		CoinType:C.CT_BTC,
		NetworkType:C.NT_TEST,
		KeyPath: keypath,
		Version: 2,
		HDprivate: 0x0488ADE4,
		HDpublic: 0x0488B21E,
		P2KH: 0x00,
		P2SH: 0x05,
		KeyPrefixes: 128,
		ApiVersion: 2,
		N: 32768,
		R: 8,
		P: 1,
	}
	fmt.Println(netparams)

	netparamsc := netparams.ToC()
	fmt.Println(netparamsc)

	mne := strings.Join(words, " ")
	wallet, ret := coinapi.CreateWallet(mne, "12345", &netparams)
	fmt.Println(wallet)

	sendValue := big.Int{}
	sendValue.SetString("100000000", 10)
	feePerKb := big.Int{}
	feePerKb.SetString("100000", 10)

	value := big.Int{}
	value.SetString("200000000", 10)

	vout := sdk.BtcVout{
		Hash:"842596976630aa2664fa2a58b9c5b618ca1c34c23ccdfc7b6de281e393b3f7c4",
		Value:value,
		N:1,
		CoinBase:false,
	}

	voutList := make([]sdk.BtcVout, 0)
	voutList = append(voutList, vout)

	btcTransactionParams := sdk.BtcTransactionParams{
		Seed:wallet.BtSeed,
		FromAddress:"n2w5TLN2hgYuumewuWrCTEsizhUAvxhSge",
		SendAddress:"2N5SZ9bMxhEfDZ3bWXHnHW4RngczVE9nb16",
		SendValue:sendValue,
		FeePerKb:feePerKb,
		Password:"12345",
		PriKey:"cVraRiXNUghfAgsms7ZtwXkXCd5KoauvBs64hzsP4J7UMVyJXgiW",
		RecipientsPayFees: false,
		BtcVoutFormList:voutList,
	}

	result := coinapi.CreateSignTransaction(&btcTransactionParams, &netparams)
	fmt.Println(result)

	defer sdk.FreeKeyPathC(keypathc)

	coinapi.Cleanup()
}
