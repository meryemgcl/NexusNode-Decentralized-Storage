package network

import (
	"context"
	"fmt"
	"log"

	"github.com/libp2p/go-libp2p"
	"github.com/libp2p/go-libp2p/core/host"
	"github.com/libp2p/go-libp2p/core/peer"
	"github.com/libp2p/go-libp2p/p2p/discovery/mdns"
)

// NexusNode represents a P2P node in the network.
type NexusNode struct {
	Host host.Host
}

// NewNode creates and starts a new P2P node.
func NewNode(listenPort int) (*NexusNode, error) {
	addr := fmt.Sprintf("/ip4/0.0.0.0/tcp/%d", listenPort)

	// Create a new libp2p Host that listens on the given port
	h, err := libp2p.New(
		libp2p.ListenAddrStrings(addr),
	)
	if err != nil {
		return nil, err
	}

	return &NexusNode{Host: h}, nil
}

// SetupDiscovery sets up mDNS discovery to find other peers on the local network.
func (n *NexusNode) SetupDiscovery(rendezvous string) error {
	mdnsService := mdns.NewMdnsService(n.Host, rendezvous, &discoveryNotifee{h: n.Host})
	return mdnsService.Start()
}

// discoveryNotifee gets notified when we find a new peer via mDNS
type discoveryNotifee struct {
	h host.Host
}

// HandlePeerFound connects to peers discovered via mDNS
func (n *discoveryNotifee) HandlePeerFound(pi peer.AddrInfo) {
	log.Printf("Discovered new peer: %s\n", pi.ID)
	err := n.h.Connect(context.Background(), pi)
	if err != nil {
		log.Printf("Error connecting to peer %s: %v\n", pi.ID, err)
	} else {
		log.Printf("Connected to peer %s\n", pi.ID)
	}
}
