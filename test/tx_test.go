package test

import (
	"fmt"
	"github.com/leverwwz/go-substrate-crypto/crypto"
	"github.com/leverwwz/go-substrate/client"
	"github.com/leverwwz/go-substrate/expand"
	"github.com/leverwwz/go-substrate/tx"
	"testing"
)

func Test_Tx2(t *testing.T) {
	url := "wss://mainnet.subgame.org"
	// 1. init client
	c, err := client.New(url)
	if err != nil {
		t.Fatal(err)
	}
	//2. for addr prefix
	//expand.SetSerDeOptions(false)
	from := "3mPjMmNqf1bp9J1NdVRnsBVAogHGQbfRvrAkheomoE52njvh"
	to := "3kWtUyYDz4Hu6GpaSXDNyXDMD43GCH3hy9Ty7eDFPqVLnUjr"
	amount := uint64(10000000000)
	//3. account info
	acc, err := c.GetAccountInfo(from)
	if err != nil {
		t.Fatal(err)
	}
	nonce := uint64(acc.Nonce)
	//4. create substrate tx
	transaction := tx.NewSubstrateTransaction(from, nonce)
	//5. expand meta
	ed, err := expand.NewMetadataExpand(c.Meta)
	if err != nil {
		t.Fatal(err)
	}
	//6. init transfer call
	call, err := ed.BalanceTransferCall(to, amount)
	if err != nil {
		t.Fatal(err)
	}
	/*
		//Balances.transfer_keep_alive  call
		btkac,err:=ed.BalanceTransferKeepAliveCall(to,amount)
	*/

	/*
		toAmount:=make(map[string]uint64)
		toAmount[to] = amount
		//...
		//true: user Balances.transfer_keep_alive  false: Balances.transfer
		ubtc,err:=ed.UtilityBatchTxCall(toAmount,false)
	*/

	//7. add transfer params
	transaction.SetGenesisHashAndBlockHash(c.GetGenesisHash(), c.GetGenesisHash()).
		SetSpecAndTxVersion(uint32(c.SpecVersion), uint32(c.TransactionVersion)).
		SetCall(call)
	//8. sign
	sig, err := transaction.SignTransaction("", crypto.Sr25519Type)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println("signed tx :", sig)
	sig = "0x41028400b6fb57574cc4c6ac0f064e0d2aa39c28b14fe93f5c6f45502dca8a5cd707523b01f0749299ffc560b7e2e33856a27019c461eccd0197b16d1a2656f06bb8b5592971ab0155602ea35ea27c9489d4daf439707bbd74bf354cb2bc5c5bc4f39fba840001020003000090348a6b6eba4478d8446cdba0e7f90293352941c6339008d9da15a021b7166b075e4ce74608"
	//9. submit tx to node
	var result interface{}
	err = c.C.Client.Call(&result, "author_submitExtrinsic", sig)
	if err != nil {
		t.Fatal(err)
	}
	//10. txid
	txid := result.(string)
	fmt.Println(txid)
}

func TestSubmitTx(t *testing.T) {
	url := "wss://mainnet.subgame.org"
	// 1. init client
	c, err := client.New(url)
	if err != nil {
		t.Fatal(err)
	}
	sig := "0x41028400b6fb57574cc4c6ac0f064e0d2aa39c28b14fe93f5c6f45502dca8a5cd707523b01f0749299ffc560b7e2e33856a27019c461eccd0197b16d1a2656f06bb8b5592971ab0155602ea35ea27c9489d4daf439707bbd74bf354cb2bc5c5bc4f39fba840001020003000090348a6b6eba4478d8446cdba0e7f90293352941c6339008d9da15a021b7166b075e4ce74608"
	//9. submit tx to node
	var result interface{}
	err = c.C.Client.Call(&result, "author_submitExtrinsic", sig)
	if err != nil {
		t.Fatal(err)
	}
	//10. txid
	txid := result.(string)
	fmt.Println(txid)

}
