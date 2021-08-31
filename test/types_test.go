package test

import (
	"encoding/json"
	"fmt"
	"github.com/leverwwz/go-substrate-crypto/ss58"
	"github.com/leverwwz/go-substrate/client"
	"testing"
)

func Test_SubGameTypes(t *testing.T) {
	// c, err := client.New("wss://mainnet.subgame.org")
	c, err := client.New("wss://subgamenode.guanfantech.com")
	if err != nil {
		panic(err)
	}

	c.SetPrefix(ss58.SubGamePrefix)
	resp, err := c.GetBlockByNumber(99099)
	if err != nil {
		panic(err)
	}

	if len(resp.Extrinsic) == 0 {
		d, _ := json.Marshal(resp)
		fmt.Println(string(d))
		t.Fatal("Empty Extrinsic")
	}
}
