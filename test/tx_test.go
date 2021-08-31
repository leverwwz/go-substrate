package test

import (
	"fmt"
	"github.com/leverwwz/go-substrate/client"
	"github.com/leverwwz/go-substrate/expand"
	"github.com/leverwwz/go-substrate/tx"
	"github.com/leverwwz/go-substrate-crypto/crypto"
	"testing"
)

func Test_Tx2(t *testing.T) {
	// 1. init client
	c, err := client.New("")
	if err != nil {
		t.Fatal(err)
	}
	//2. for addr prefix
	//expand.SetSerDeOptions(false)
	from := ""
	to := ""
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
