package internal

import (
	"context"
	"fmt"
	"log"
	"net"
	"time"

	"spacecore_registry/pb"

	"github.com/ipfs/kubo/client/rpc"
	"github.com/libp2p/go-libp2p"
	dht "github.com/libp2p/go-libp2p-kad-dht"
	kbucket "github.com/libp2p/go-libp2p-kbucket"
	"github.com/libp2p/go-libp2p/core/host"
	"github.com/redis/go-redis/v9"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

// *kbucket.RoutingTable

func SetupHost() (host.Host, *kbucket.RoutingTable, error) {
	ctx := context.Background()
	h, err := libp2p.New(
		libp2p.ListenAddrStrings("/ip4/0.0.0.0/tcp/0"),
	)
	if err != nil {
		return nil, nil, fmt.Errorf("failed to create libp2p host: %w", err)
	}

	// let's use kbucket DHT instead
	rt, err := kbucket.NewRoutingTable(
		20, kbucket.ConvertPeerID(h.ID()), 10*time.Second, h.Peerstore(), time.Second, nil,
	)

	// rt, err := kbucket.NewRoutingTable(
	// 	20, kbucket.ConvertPeerID(h.ID()), h.Peerstore(), 10*time.Second,
	// )
	// Set up the DHT
	// rt := kbucket.NewRoutingTable(20, kbucket.ConvertPeerID(h.ID()), h.Peerstore(), 10*time.Second)

	if err != nil {
		return nil, nil, fmt.Errorf("failed to create routing table: %w", err)
	}
	// dht, err := dht.New(ctx, h, dht.RoutingTable(rt))

	// Set up the DHT
	dht, err := dht.New(ctx, h)
	if err != nil {
		return nil, nil, fmt.Errorf("failed to create DHT: %w", err)
	}
	if err := dht.Bootstrap(ctx); err != nil {
		return nil, nil, fmt.Errorf("failed to bootstrap DHT: %w", err)
	}

	// return h, rt, nil
	return h, rt, nil
}

func Start(h host.Host, dht *kbucket.RoutingTable) {
	ipfsClient, err := rpc.NewLocalApi()
	if err != nil {
		log.Fatalf("Failed to create IPFS client h: %v", err)
	}
	grpcServer := grpc.NewServer()
	redisClient := redis.NewClient(&redis.Options{
		// Addr: "localhost:6379", make this dynamic from env
		Addr: fmt.Sprintf("%s:%s", "0.0.0.0", "6379"),
	})
	pb.RegisterPluginRegistryServer(grpcServer, &pluginRegistryServer{
		ipfsClient:  ipfsClient,
		redisClient: redisClient,
	})
	reflection.Register(grpcServer)

	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	log.Println("Starting Spacecore Registry server on port 50051")
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("Failed to serve gRPC server: %v", err)
	}
}
