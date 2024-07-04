package p2p

import (
	"context"
	"fmt"

	libp2p "github.com/libp2p/go-libp2p"
	dht "github.com/libp2p/go-libp2p-kad-dht"
	"github.com/libp2p/go-libp2p/core/host"
	"github.com/libp2p/go-libp2p/core/peer"

	// "github.com/libp2p/go-libp2p/core/routing"
	"github.com/libp2p/go-libp2p/p2p/discovery/routing"
)

func SetupHost(ctx context.Context) (host.Host, *dht.IpfsDHT, error) {
	h, err := libp2p.New(
		libp2p.ListenAddrStrings("/ip4/0.0.0.0/tcp/0"),
	)
	if err != nil {
		return nil, nil, fmt.Errorf("failed to create libp2p host: %w", err)
	}

	dht, err := dht.New(ctx, h)
	if err != nil {
		return nil, nil, fmt.Errorf("failed to create DHT: %w", err)
	}

	if err := dht.Bootstrap(ctx); err != nil {
		return nil, nil, fmt.Errorf("failed to bootstrap DHT: %w", err)
	}

	return h, dht, nil
}

func Advertise(ctx context.Context, h host.Host, dht *dht.IpfsDHT) error {
	rdisc := routing.NewRoutingDiscovery(dht)
	_, err := rdisc.Advertise(ctx, "spacecore-registry")
	return err
}

func DiscoverPeers(ctx context.Context, dht *dht.IpfsDHT) ([]peer.AddrInfo, error) {
	rdisc := routing.NewRoutingDiscovery(dht)
	peers, err := rdisc.FindPeers(ctx, "spacecore-registry")
	if err != nil {
		return nil, err
	}

	var results []peer.AddrInfo
	for p := range peers {
		if p.ID == "" {
			continue
		}
		results = append(results, p)
	}
	return results, nil
}
