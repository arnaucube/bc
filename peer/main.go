package main

import (
	"fmt"
	"os"
	"time"

	blockchainlib "./blockchainlib"
	p2plib "./p2plib"
	"github.com/fatih/color"
)

var tp p2plib.ThisPeer
var blockchain blockchainlib.Blockchain

func main() {

	if len(os.Args) < 3 {
		color.Red("need to call:")
		color.Red("./peer client 3001 3002")
		os.Exit(3)
	}

	color.Blue("Starting Peer")
	//read configuration file
	readConfig("config.json")

	//initialize blockchainlib
	//InitializeBlockchain(role, ip, port, restport, serverip, serverport)
	config.RESTPort = os.Args[3]
	tp := blockchain.InitializeBlockchain(os.Args[1], "127.0.0.1",
		os.Args[2], os.Args[3], config.ServerIP, config.ServerPort)

	if tp.RunningPeer.Role == "client" {
		color.Red("http://" + config.IP + ":" + config.ServerRESTPort)
		fmt.Println(blockchain.GenesisBlock)
		blockchain.ReconstructBlockchainFromBlock("http://"+config.IP+":"+config.ServerRESTPort, blockchain.GenesisBlock)
	}
	color.Blue("peer and blokcchain initialized")
	go runRestServer() //TODO this will not be necessary, due the communications will go full over tcp connections

	for tp.Running {
		time.Sleep(1000 * time.Millisecond)
	}
}
