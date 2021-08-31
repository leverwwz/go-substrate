package test

import (
	"encoding/json"
	"fmt"
	"github.com/leverwwz/go-substrate/client"
	"github.com/leverwwz/go-substrate-crypto/ss58"
	"testing"
)

func Test_GetBlockByNumber(t *testing.T) {
	c, err := client.New("wss://rpc.polkadot.io")
	if err != nil {
		t.Fatal(err)
	}
	c.SetPrefix(ss58.KsmPrefix)
	//expand.SetSerDeOptions(false)
	/*
		Ksm: 7834050
	*/
	resp, err := c.GetBlockByNumber(5429211)
	if err != nil {
		t.Fatal(err)
	}

	d, _ := json.Marshal(resp)
	fmt.Println(string(d))
}

func Test_GetAccountInfo(t *testing.T) {
	url := "wss://kusama-rpc.polkadot.io" // wss://kusama-rpc.polkadot.io
	address := "DXuShaYiV3gqYspg7mzdDmweS9p79Z9u3wEY9FH3rHaj6yN"

	c, err := client.New(url)
	if err != nil {
		t.Fatal(err)
	}
	// 000000000000000001000000000000000047ab56020000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000
	ai, err := c.GetAccountInfo(address)
	if err != nil {
		t.Fatal(err)
	}
	d, _ := json.Marshal(ai)
	fmt.Println(string(d))
	fmt.Println(ai.Data.Free.String())
}

func Test_GetBlockExtrinsic(t *testing.T) {
	url := "wss://kusama-rpc.polkadot.io" // wss://kusama-rpc.polkadot.io
	c, err := client.New(url)
	if err != nil {
		t.Fatal(err)
	}
	h, err := c.C.RPC.Chain.GetBlockHash(7812476)
	if err != nil {
		t.Fatal(err)
	}
	block, err := c.C.RPC.Chain.GetBlock(h)
	if err != nil {
		t.Fatal(err)
	}
	for _, extrinsic := range block.Block.Extrinsics {
		fmt.Println(extrinsic)
	}
}

func Test_SubGameGetBlockByNumber(t *testing.T) {
	c, err := client.New("wss://mainnet.subgame.org")
	if err != nil {
		t.Fatal(err)
	}
	c.SetPrefix(ss58.SubGamePrefix)

	resp, err := c.GetBlockByNumber(521297)
	if err != nil {
		t.Fatal(err)
	}

	d, _ := json.Marshal(resp)
	fmt.Println(string(d))
}

func Test_SubGameGetAccountInfo(t *testing.T) {
	url := "wss://mainnet.subgame.org"
	address := "3hKdcCz7fcu344acfo6p1JqnehydfdSX5VueExUy9rCjy64a"

	c, err := client.New(url)
	if err != nil {
		t.Fatal(err)
	}

	ai, err := c.GetAccountInfo(address)
	if err != nil {
		t.Fatal(err)
	}
	d, _ := json.Marshal(ai)
	fmt.Println(string(d))
	fmt.Println(ai.Data.Free.String())
}

func Test_SubGameGetBlockExtrinsic(t *testing.T) {
	url := "wss://mainnet.subgame.org"
	c, err := client.New(url)
	if err != nil {
		t.Fatal(err)
	}
	h, err := c.C.RPC.Chain.GetBlockHash(520429)
	if err != nil {
		t.Fatal(err)
	}
	block, err := c.C.RPC.Chain.GetBlock(h)
	if err != nil {
		t.Fatal(err)
	}
	for _, extrinsic := range block.Block.Extrinsics {
		fmt.Println(extrinsic)
	}
}
