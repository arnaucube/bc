package blockchainlib

import (
	"fmt"

	p2plib "../p2plib"
)

func (bc *Blockchain) CreateMsgHandlerCases() map[string]func(p2plib.Peer, p2plib.Msg) {
	configuredMsgCases := make(map[string]func(p2plib.Peer, p2plib.Msg))
	configuredMsgCases["Block"] = func(peer p2plib.Peer, msg p2plib.Msg) {
		//TODO check if the block is signed by an autorized emitter
		//block = msg.Data converted to Block
		var block Block
		if !bc.BlockExists(block) {
			bc.AddBlock(block)
			p2plib.PropagateData(peer, "block in string format")
		}
	}

	fmt.Println(configuredMsgCases)
	return configuredMsgCases
}
