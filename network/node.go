package network

import (
	"context"
	"fmt"
	"log"

	"github.com/libp2p/go-libp2p"
	dht "github.com/libp2p/go-libp2p-kad-dht"
	"github.com/libp2p/go-libp2p/core/host"
	"github.com/libp2p/go-libp2p/core/peer"
	"github.com/libp2p/go-libp2p/p2p/discovery/mdns"
	"github.com/libp2p/go-libp2p/p2p/discovery/routing"
	"github.com/libp2p/go-libp2p/p2p/discovery/util"
)

// NexusNode represents a P2P node in the network.
type NexusNode struct {
	Host host.Host
	DHT  *dht.IpfsDHT
}

// NewNode creates and starts a new P2P node.
func NewNode(listenPort int) (*NexusNode, error) {
	addr := fmt.Sprintf("/ip4/0.0.0.0/tcp/%d", listenPort)

	h, err := libp2p.New(
		libp2p.ListenAddrStrings(addr),
	)
	if err != nil {
		return nil, fmt.Errorf("NewNode: create libp2p host: %w", err)
	}

	// Create a Kademlia DHT for global (internet-wide) peer discovery.
	// dht.ModeServer makes this node participate in routing queries, not just lookups.
	ctx := context.Background()
	kadDHT, err := dht.New(ctx, h, dht.Mode(dht.ModeServer))
	if err != nil {
		return nil, fmt.Errorf("NewNode: create DHT: %w", err)
	}

	// Bootstrap the DHT by connecting to well-known bootstrap peers.
	if err := kadDHT.Bootstrap(ctx); err != nil {
		return nil, fmt.Errorf("NewNode: DHT bootstrap: %w", err)
	}

	return &NexusNode{Host: h, DHT: kadDHT}, nil
}

// SetupDiscovery enables both mDNS (local network) and DHT-based (global) peer discovery.
// mDNS finds peers on the same LAN; DHT finds peers across the internet.
func (n *NexusNode) SetupDiscovery(rendezvous string) error {
	// 1. mDNS — local area network discovery
	mdnsService := mdns.NewMdnsService(n.Host, rendezvous, &discoveryNotifee{h: n.Host})
	if err := mdnsService.Start(); err != nil {
		return fmt.Errorf("SetupDiscovery: start mDNS: %w", err)
	}
	log.Println("mDNS discovery started")

	// 2. DHT — global peer discovery via rendezvous point
	routingDiscovery := routing.NewRoutingDiscovery(n.DHT)
	util.Advertise(context.Background(), routingDiscovery, rendezvous)
	log.Println("Kademlia DHT discovery advertised rendezvous:", rendezvous)

	// Start discovering peers via DHT in the background
	go func() {
		peerChan, err := routingDiscovery.FindPeers(context.Background(), rendezvous)
		if err != nil {
			log.Printf("DHT FindPeers error: %v", err)
			return
		}
		for p := range peerChan {
			if p.ID == n.Host.ID() {
				continue // skip self
			}
			log.Printf("DHT discovered peer: %s", p.ID)
			if err := n.Host.Connect(context.Background(), p); err != nil {
				log.Printf("DHT connect error (%s): %v", p.ID, err)
			} else {
				log.Printf("DHT connected to peer: %s", p.ID)
			}
		}
	}()

	return nil
}

// discoveryNotifee gets notified when we find a new peer via mDNS.
type discoveryNotifee struct {
	h host.Host
}

// HandlePeerFound connects to peers discovered via mDNS.
func (n *discoveryNotifee) HandlePeerFound(pi peer.AddrInfo) {
	log.Printf("mDNS discovered peer: %s", pi.ID)
	if err := n.h.Connect(context.Background(), pi); err != nil {
		log.Printf("mDNS connect error (%s): %v", pi.ID, err)
	} else {
		log.Printf("mDNS connected to peer: %s", pi.ID)
	}
}
