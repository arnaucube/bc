package p2plib

import (
	"crypto/sha256"
	"encoding/base64"
	"math/rand"
	"net"
	"strings"
)

func GetIPPortFromConn(conn net.Conn) string {
	ip := GetIPFromConn(conn)
	port := GetPortFromConn(conn)
	return ip + ":" + port
}
func GetIPFromConn(conn net.Conn) string {
	s := conn.RemoteAddr().String()
	s = strings.Split(s, ":")[0]
	s = strings.Trim(s, ":")
	return s
}
func GetPortFromConn(conn net.Conn) string {
	s := conn.RemoteAddr().String()
	s = strings.Split(s, ":")[1]
	s = strings.Trim(s, ":")
	return s
}
func RandInt(min int, max int) int {
	r := rand.Intn(max-min) + min
	return r
}
func HashPeer(p Peer) string {
	peerString := p.IP + ":" + p.Port

	h := sha256.New()
	h.Write([]byte(peerString))
	return base64.URLEncoding.EncodeToString(h.Sum(nil))
}
