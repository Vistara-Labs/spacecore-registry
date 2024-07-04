package main

import (
	"log"
	"spacecore_registry/internal"
	// "github.com/ipfs/kubo/client/rpc"
)

func main() {
	// ctx := context.Background()
	h, dht, err := internal.SetupHost()

	if err != nil {
		log.Fatalf("Failed to set up host: %v", err)
	}
	// setup routing discovery
	// rdisc := routing.NewRoutingDiscovery(dht)
	// rdisc.Advertise(ctx, "spacecore-registry")

	// rt, err := kbucket.NewRoutingTable(
	// 	20, kbucket.ConvertPeerID(h.ID()), 10*time.Second, h.Peerstore(), time.Second, nil,
	// )

	// Start the gRPC server
	internal.Start(h, dht)

	// Log the peer ID
	log.Printf("Peer ID: %v\n", h.ID())

	// Keep the process alive
	select {}

}
