package p2p

import (
	"fmt"
	"net"
	"sync"
)

// TCPPeer represents the remote node over a TCP established connection
type TCPPeer struct {
	// con is the underlying connection
	conn net.Conn

	// if we dial and retrive a connection => outbound true
	// if we accept  and retrive  => outbound false
	outbound bool
}

func NewTCPPeer(conn net.Conn, outbound bool) *TCPPeer {
	return &TCPPeer{
		conn:     conn,
		outbound: outbound,
	}
}

type TCPTransport struct {
	listenAddres string
	listener     net.Listener
	shakeHands   HandShakeFunc
	decoder      Decoder
	peers        map[net.Addr]Peer
	mu           sync.RWMutex
}

func NewTCPTransport(listenAddress string) *TCPTransport {
	return &TCPTransport{
		shakeHands:   NOPHandShakeFunc,
		listenAddres: listenAddress,
	}
}

func (tr *TCPTransport) ListenAndAccept() error {
	var err error
	tr.listener, err = net.Listen("tcp", tr.listenAddres)
	if err != nil {
		return err
	}
	go tr.startAcceptLoop()

	return nil
}

func (tr *TCPTransport) startAcceptLoop() {
	for {
		conn, err := tr.listener.Accept()
		if err != nil {
			fmt.Printf("TCP accept errors: %v\n", err)
		}

		fmt.Printf("New TCP connection: %v\n", conn)

		go tr.handleConnection(conn)
	}
}

type Temp struct {
}

func (tr *TCPTransport) handleConnection(conn net.Conn) {
	peer := NewTCPPeer(conn, true)

	if err := tr.shakeHands(peer); err != nil {
		fmt.Printf("Handshake error: %v\n", err)
		return
	}

	// Read Loop
	msg := &Temp{}
	for {
		if err := tr.decoder.Decode(conn, msg); err != nil {
			fmt.Printf("TCP Error: decoding message: %v\n", err)
			continue
		}
	}

}
