package p2p

// Peer is an interface that represents the remote node
type Peer interface {
}

// Transport is anything that handles communication
// between nodes in network
// Can be in form of TCP, UDP, websockets, ...
type Transport interface {
	ListendAndAccept() error
}
