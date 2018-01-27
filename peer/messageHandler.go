package main

import (
	"fmt"

	p2plib "./p2plib"

	blockchainlib "./blockchainlib"
)

func createMsgHandlerCases() map[string]func(p2plib.Peer, p2plib.Msg) {
	configuredMsgCases := make(map[string]func(p2plib.Peer, p2plib.Msg))
	configuredMsgCases["Block"] = func(peer p2plib.Peer, msg p2plib.Msg) {
		//TODO check if the block is signed by an autorized emitter
		//block = msg.Data converted to Block
		var block blockchainlib.Block
		if !blockchain.BlockExists(block) {
			blockchain.AddBlock(block)
			p2plib.PropagateData(peer, "block in string format")
		}
	}

	fmt.Println(configuredMsgCases)
	return configuredMsgCases
}
