package base

import "github.com/leverwwz/go-substrate-rpc-client/types"

type HrmpChannelId struct {
	Sender   types.U32
	Receiver types.U32
}
