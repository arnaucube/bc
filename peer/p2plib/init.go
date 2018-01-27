package p2plib

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/fatih/color"
)

func InitializePeer(role, ip, port, restport, serverip, serverport string,
	configuredMsgCases map[string]func(Peer, Msg)) ThisPeer {
	//initialize some vars
	rand.Seed(time.Now().Unix())

	InitializeDefaultMsgCases(configuredMsgCases)

	var tp ThisPeer
	tp.Running = true
	tp.RunningPeer.Role = role
	tp.RunningPeer.Port = port
	tp.RunningPeer.RESTPort = restport
	tp.RunningPeer.ID = HashPeer(tp.RunningPeer)

	tp.ID = tp.RunningPeer.ID
	globalTP.PeersConnections.Outcoming.PeerID = tp.RunningPeer.ID
	fmt.Println(tp.RunningPeer)
	//outcomingPeersList.Peers = append(outcomingPeersList.Peers, peer.RunningPeer)
	globalTP.PeersConnections.Outcoming = AppendPeerIfNoExist(globalTP.PeersConnections.Outcoming, tp.RunningPeer)
	fmt.Println(globalTP.PeersConnections.Outcoming)

	if tp.RunningPeer.Role == "server" {
		go tp.AcceptPeers(tp.RunningPeer)
	}
	if tp.RunningPeer.Role == "client" {
		var serverPeer Peer
		serverPeer.IP = serverip
		serverPeer.Port = serverport
		serverPeer.Role = "server"
		serverPeer.ID = HashPeer(serverPeer)
		go tp.AcceptPeers(tp.RunningPeer)
		ConnectToPeer(serverPeer)
	}
	globalTP = tp

	return tp

}

func InitializeDefaultMsgCases(configuredMsgCases map[string]func(Peer, Msg)) {
	//msgCases := make(map[string]func(Peer, Msg)) --> no, it's used the global
	msgCases = make(map[string]func(Peer, Msg))

	//get the user configured msgCases
	for k, v := range configuredMsgCases {
		msgCases[k] = v
	}

	msgCases["Hi"] = func(peer Peer, msg Msg) {
		color.Yellow(msg.Type)
		color.Green(msg.Content)
	}

	msgCases["PeersList"] = func(peer Peer, msg Msg) {
		color.Blue("newPeerslist")
		fmt.Println(msg.PeersList)
		color.Red("PeersList")

		UpdateNetworkPeersList(peer.Conn, msg.PeersList)
		PropagatePeersList(peer)
		PrintPeersList()
	}

	msgCases["Data"] = func(peer Peer, msg Msg) {
		color.Yellow(msg.Type)
		color.Green(msg.Content)
		PropagateData(peer, "data")
	}
	fmt.Println(msgCases)
}
