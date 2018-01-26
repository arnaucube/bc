package blockchainlib

import (
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
	"log"
)

func HashBlock(b Block) string {
	blockJson, err := json.Marshal(b)
	if err != nil {
		log.Println(err)
	}
	blockString := string(blockJson)

	h := sha256.New()
	h.Write([]byte(blockString))
	return base64.URLEncoding.EncodeToString(h.Sum(nil))
}
