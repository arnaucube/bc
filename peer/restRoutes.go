package main

import (
	"encoding/json"
	"fmt"
	"net/http"

	blockchainlib "./blockchainlib"
	p2plib "./p2plib"
	"github.com/fatih/color"
	"github.com/gorilla/mux"
)

type Routes []Route

var routes = Routes{
	Route{
		"Index",
		"GET",
		"/",
		Index,
	},
	Route{
		"GetPeers",
		"GET",
		"/peers",
		GetPeers,
	},
	Route{
		"PostUser",
		"POST",
		"/register",
		PostUser,
	},
	Route{
		"GenesisBlock",
		"GET",
		"/blocks/genesis",
		GenesisBlock,
	},
	Route{
		"NextBlock",
		"GET",
		"/blocks/next/{blockhash}",
		NextBlock,
	},
	Route{
		"LastBlock",
		"GET",
		"/blocks/last",
		LastBlock,
	},
}

type Address struct {
	Address string `json:"address"` //the pubK of the user, to perform logins
}

func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, tp.ID)
}
func GetPeers(w http.ResponseWriter, r *http.Request) {
	jResp, err := json.Marshal(tp.PeersConnections.Outcoming)
	check(err)
	fmt.Fprintln(w, string(jResp))
}
func PostUser(w http.ResponseWriter, r *http.Request) {

	decoder := json.NewDecoder(r.Body)
	var address string
	err := decoder.Decode(&address)
	if err != nil {
		panic(err)
	}
	defer r.Body.Close()
	fmt.Println(address)
	color.Blue(address)

	//TODO add the verification of the address, to decide if it's accepted to create a new Block
	block := blockchain.CreateBlock(address)
	blockchain.AddBlock(block)

	go PropagateBlock(block)

	jResp, err := json.Marshal(blockchain)
	check(err)
	fmt.Fprintln(w, string(jResp))
}

func PropagateBlock(b blockchainlib.Block) {
	//prepare the msg to send to all connected peers
	var msg p2plib.Msg
	msg.Construct("Block", "new block")
	bJson, err := json.Marshal(b)
	check(err)

	msg.Data = []byte(bJson)
	msgB := msg.ToBytes()
	for _, peer := range tp.PeersConnections.Outcoming.Peers {
		if peer.Conn != nil {
			_, err := peer.Conn.Write(msgB)
			check(err)
		}
	}
}

func GenesisBlock(w http.ResponseWriter, r *http.Request) {
	var genesis blockchainlib.Block
	if len(blockchain.Blocks) > 0 {
		genesis = blockchain.Blocks[0]
	}

	jResp, err := json.Marshal(genesis)
	check(err)
	fmt.Fprintln(w, string(jResp))
}

func NextBlock(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	blockhash := vars["blockhash"]

	currBlock, err := blockchain.GetBlockByHash(blockhash)
	check(err)
	nextBlock, err := blockchain.GetBlockByHash(currBlock.NextHash)
	check(err)

	jResp, err := json.Marshal(nextBlock)
	check(err)
	fmt.Fprintln(w, string(jResp))
}

func LastBlock(w http.ResponseWriter, r *http.Request) {
	var genesis blockchainlib.Block
	if len(blockchain.Blocks) > 0 {
		genesis = blockchain.Blocks[len(blockchain.Blocks)-1]
	}

	jResp, err := json.Marshal(genesis)
	check(err)
	fmt.Fprintln(w, string(jResp))
}
