package p2plib

func SendPetition(peer Peer, petition string) {
	_, err := peer.Conn.Write([]byte(petition))
	check(err)
}
