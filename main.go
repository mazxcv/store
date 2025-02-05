package main

import (
	"log"

	"github.com/mazxcv/store/p2p"
)

func main() {

	tr := p2p.NewTCPTransport("127.0.0.1:4000")
	if err := tr.ListenAndAccept(); err != nil {
		log.Fatal(err)
	}

	select {}
}
