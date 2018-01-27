package blockchainlib

import (
	p2plib "../p2plib"
)

func (bc *Blockchain) InitializeBlockchain(role, ip, port, restport, serverip, serverport string) p2plib.ThisPeer {

	//read the stored blockchain
	err := bc.ReadFromDisk()
	check(err)
	bc.Print()

	//get blockchain msgHandlerCases
	configuredMsgCases := bc.CreateMsgHandlerCases()
	//initialize p2plib, adding the configuredMsgCases to the p2plib msgCases to handle
	tp := p2plib.InitializePeer(role, ip, port, restport, serverip,
		serverport, configuredMsgCases)
	//return thisPeer (tp)
	return tp
}
