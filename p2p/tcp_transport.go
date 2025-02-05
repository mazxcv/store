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
	peers        map[net.Addr]Peer
	mu           sync.RWMutex
}

func NewTCPTransport(listenAddress string) *TCPTransport {
	return &TCPTransport{
		listenAddres: listenAddress,
		peers:        make(map[net.Addr]Peer),
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
		go tr.handleConnection(conn)
	}
}

func (tr *TCPTransport) handleConnection(conn net.Conn) {
	peer := NewTCPPeer(conn, false)
	fmt.Printf("New incomming connection: %v\n", peer)
}
