package p2plib

import (
	"fmt"
	"math/rand"
	"time"
)

func InitializePeer(role, ip, port, restport, serverip, serverport string) ThisPeer {
	//initialize some vars
	rand.Seed(time.Now().Unix())

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
