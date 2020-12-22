package requests

import (
	"log"
	"testing"
)

func TestReqbinEchoPost(t *testing.T) {
	err := ReqbinEchoPost()
	if err != nil {
		log.Fatal(err)
	}
}
