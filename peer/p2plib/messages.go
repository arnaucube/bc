package p2plib

import (
	"encoding/json"
	"log"
	"time"
)

type Msg struct {
	Type      string    `json:"type"`
	Date      time.Time `json:"date"`
	Content   string    `json:"content"`
	PeersList PeersList `json:"peerslist"`
	Data      []byte    `json:"data"`
}

type Case struct {
	Case     string
	Function func(Peer, Msg)
}

var msgCases map[string]func(Peer, Msg)

func MessageHandler(peer Peer, msg Msg) {

	log.Println("[New msg]")
	log.Println(msg)

	/*for c, caseFunction := range msgCases {
		if msg.Type == c {
			caseFunction(peer, msg)
		}
	}*/
	msgCases[msg.Type](peer, msg)

}
func (msg *Msg) Construct(msgtype string, msgcontent string) {
	msg.Type = msgtype
	msg.Content = msgcontent
	msg.Date = time.Now()
}
func (msg Msg) ToBytes() []byte {
	msgS, err := json.Marshal(msg)
	check(err)
	l := string(msgS) + "\n"
	r := []byte(l)
	return r
}
func (msg Msg) CreateFromBytes(bytes []byte) Msg {
	err := json.Unmarshal(bytes, &msg)
	check(err)
	return msg
}
