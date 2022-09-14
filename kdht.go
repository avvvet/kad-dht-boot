package main

import (
	"context"

	"github.com/libp2p/go-libp2p-core/host"
	dht "github.com/libp2p/go-libp2p-kad-dht"
)

func NewKDHT(ctx context.Context, host host.Host) (*dht.IpfsDHT, error) {
	var options []dht.Option

	options = append(options, dht.Mode(dht.ModeServer))

	kdht, err := dht.New(ctx, host, options...)
	if err != nil {
		return nil, err
	}

	if err = kdht.Bootstrap(ctx); err != nil {
		return nil, err
	}

	return kdht, nil
}
