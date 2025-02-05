package p2p

// HandShakeFunc is a function that is called when a new connection is established
type HandShakeFunc func(Peer) error

func NOPHandShakeFunc(Peer) error {
	return nil
}
