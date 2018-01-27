package p2plib

import (
	"fmt"
	"net"
	"time"

	"github.com/fatih/color"
)

type Peer struct {
	ID       string   `json:"id"` //in the future, this will be the peer hash
	IP       string   `json:"ip"`
	Port     string   `json:"port"`
	RESTPort string   `json:"restport"`
	Role     string   `json:"role"` //client or server
	Conn     net.Conn `json:"conn"`
}

type PeersList struct {
	PeerID string
	Peers  []Peer    `json:"peerslist"`
	Date   time.Time `json:"date"`
}

type PeersConnections struct {
	Incoming  PeersList
	Outcoming PeersList
	Network   PeersList //the peers that have been received in the lists from other peers
}

type ThisPeer struct {
	Running          bool
	ID               string
	RunningPeer      Peer
	PeersConnections PeersConnections
}

var globalTP ThisPeer

func PeerIsInPeersList(p Peer, pl []Peer) int {
	r := -1
	for i, peer := range pl {
		if peer.IP+":"+peer.Port == p.IP+":"+p.Port {
			r = i
		}
	}
	return r
}

func DeletePeerFromPeersList(p Peer, pl *PeersList) {
	i := PeerIsInPeersList(p, pl.Peers)
	if i != -1 {
		//delete peer from pl.Peers
		pl.Peers = append(pl.Peers[:i], pl.Peers[i+1:]...)
	}
}
func AppendPeerIfNoExist(pl PeersList, p Peer) PeersList {
	i := PeerIsInPeersList(p, pl.Peers)
	if i == -1 {
		pl.Peers = append(pl.Peers, p)
	}
	return pl
}
func UpdateNetworkPeersList(conn net.Conn, newPeersList PeersList) {
	for _, peer := range newPeersList.Peers {
		if GetIPPortFromConn(conn) == peer.IP+":"+peer.Port {
			peer.ID = newPeersList.PeerID
			color.Yellow(peer.ID)
		}
		i := PeerIsInPeersList(peer, globalTP.PeersConnections.Network.Peers)
		if i == -1 {
			globalTP.PeersConnections.Network.Peers = append(globalTP.PeersConnections.Network.Peers, peer)
		} else {
			fmt.Println(globalTP.PeersConnections.Network.Peers[i])
			globalTP.PeersConnections.Network.Peers[i].ID = peer.ID
		}
	}
}
func SearchPeerAndUpdate(p Peer) {
	for _, peer := range globalTP.PeersConnections.Outcoming.Peers {
		color.Red(p.IP + ":" + p.Port)
		color.Yellow(peer.IP + ":" + peer.Port)
		if p.IP+":"+p.Port == peer.IP+":"+peer.Port {
			peer.ID = p.ID
		}
	}
}

//send the outcomingPeersList to all the peers except the peer p that has send the outcomingPeersList
func PropagatePeersList(p Peer) {
	for _, peer := range globalTP.PeersConnections.Network.Peers {
		if peer.Conn != nil {
			if peer.ID != p.ID && p.ID != "" {
				color.Yellow(peer.ID + " - " + p.ID)
				var msg Msg
				msg.Construct("PeersList", "here my outcomingPeersList")
				msg.PeersList = globalTP.PeersConnections.Outcoming
				msgB := msg.ToBytes()
				_, err := peer.Conn.Write(msgB)
				check(err)
			} else {
				/*
					for the moment, this is not being called, due that in the IncomingPeersList,
					there is no peer.ID, so in the comparation wih the peer that has send the
					peersList, is comparing ID with "", so nevere enters this 'else' section

					maybe it's not needed. TODO check if it's needed the PeerList_Response
					For the moment is working without it
				*/
				//to the peer that has sent the peerList, we send our PeersList

				var msg Msg
				msg.Construct("PeersList_Response", "here my outcomingPeersList")
				msg.PeersList = globalTP.PeersConnections.Outcoming
				msgB := msg.ToBytes()
				_, err := peer.Conn.Write(msgB)
				check(err)
			}
		} else {
			//connect to peer
			if peer.ID != p.ID && peer.ID != globalTP.RunningPeer.ID {
				if PeerIsInPeersList(peer, globalTP.PeersConnections.Outcoming.Peers) == -1 {
					color.Red("no connection, connecting to peer: " + peer.Port)
					ConnectToPeer(peer)
				}
			}

		}
	}
}
func PrintPeersList() {
	fmt.Println("")
	color.Blue("runningPeer.ID: " + globalTP.RunningPeer.ID)
	color.Green("OUTCOMING PEERSLIST:")
	for _, peer := range globalTP.PeersConnections.Outcoming.Peers {
		fmt.Println(peer)
	}
	color.Green("INCOMING PEERSLIST:")
	for _, peer := range globalTP.PeersConnections.Incoming.Peers {
		fmt.Println(peer)
	}

	color.Green("NETWORK PEERSLIST:")
	for _, peer := range globalTP.PeersConnections.Network.Peers {
		fmt.Println(peer)
	}
	fmt.Println("")
}
func PropagateData(p Peer, s string) {
	//prepare the msg to send to all connected peers
	var msg Msg
	msg.Construct("Data", "new Data")
	msg.Data = []byte(s)
	msgB := msg.ToBytes()
	for _, peer := range globalTP.PeersConnections.Outcoming.Peers {
		if peer.Conn != nil {
			if peer.ID != p.ID && p.ID != "" {
				_, err := peer.Conn.Write(msgB)
				check(err)
			}
		}
	}
}

//send the block to all the peers of the outcomingPeersList
/*func PropagateBlock(b Block) {
	//prepare the msg to send to all connected peers
	var msg Msg
	msg.construct("Block", "new block")
	msg.Block = b
	msgB := msg.toBytes()
	for _, peer := range outcomingPeersList.Peers {
		if peer.Conn != nil {
			_, err := peer.Conn.Write(msgB)
			check(err)
		}
	}
}*/
