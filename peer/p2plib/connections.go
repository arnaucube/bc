package p2plib

import (
	"bufio"
	"fmt"
	"log"
	"net"

	"github.com/fatih/color"
)

func (tp *ThisPeer) AcceptPeers(peer Peer) {
	fmt.Println("accepting peers at: " + peer.Port)
	l, err := net.Listen("tcp", peer.IP+":"+peer.Port)
	if err != nil {
		log.Println("Error accepting peers. Listening port: " + peer.Port)
		tp.Running = false
	}
	for tp.Running {
		conn, err := l.Accept()
		if err != nil {
			log.Println("Error accepting peers. Error accepting connection")
			tp.Running = false
		}
		var newPeer Peer
		newPeer.IP = GetIPFromConn(conn)
		newPeer.Port = GetPortFromConn(conn)
		newPeer.Conn = conn
		globalTP.PeersConnections.Incoming = AppendPeerIfNoExist(globalTP.PeersConnections.Incoming, newPeer)
		go HandleConn(conn, newPeer)
	}
}
func ConnectToPeer(peer Peer) {
	color.Green("connecting to new peer")
	log.Println("Connecting to new peer: " + peer.IP + ":" + peer.Port)
	conn, err := net.Dial("tcp", peer.IP+":"+peer.Port)
	if err != nil {
		log.Println("Error connecting to: " + peer.IP + ":" + peer.Port)
		return
	}
	peer.Conn = conn
	globalTP.PeersConnections.Outcoming = AppendPeerIfNoExist(globalTP.PeersConnections.Outcoming, peer)
	go HandleConn(conn, peer)
}
func HandleConn(conn net.Conn, connPeer Peer) {
	connRunning := true
	log.Println("handling conn: " + conn.RemoteAddr().String())
	//reply to the conn, send the peerList
	var msg Msg
	msg.Construct("PeersList", "here my outcomingPeersList")
	msg.PeersList = globalTP.PeersConnections.Outcoming
	msgB := msg.ToBytes()
	_, err := conn.Write(msgB)
	if err != nil {
		log.Println(err)
	}

	for connRunning {
		/*
			buffer := make([]byte, 1024)
			bytesRead, err := conn.Read(buffer)
		*/
		newmsg, err := bufio.NewReader(conn).ReadString('\n')
		if err != nil {
			log.Println(err)
			connRunning = false
		} else {
			var msg Msg
			//msg = msg.createFromBytes([]byte(string(buffer[0:bytesRead])))
			msg = msg.CreateFromBytes([]byte(newmsg))
			MessageHandler(connPeer, msg)
		}
	}
	//TODO add that if the peer closed is the p2p server, show a warning message at the peer
	log.Println("Peer: " + conn.RemoteAddr().String() + " connection closed")
	conn.Close()
	//TODO delete the peer from the outcomingPeersList --> DONE
	DeletePeerFromPeersList(connPeer, &globalTP.PeersConnections.Outcoming)
	/*color.Yellow("peer deleted, current peerList:")
	PrintPeersList()*/
}
