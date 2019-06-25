/* This is free and unencumbered software released into the public domain. */

package sdk

import (
	"context"
	"strings"

	"github.com/grandcat/zeroconf"
)

// GameEndpoint describes a discovered game endpoint (name, hostname, port).
type GameEndpoint struct {
	Name    string
	Host    string
	Port    int
	Version string
}

// DiscoverGames attempts to discover ongoing games on the local network.
func DiscoverGames(ctx context.Context) (<-chan *GameEndpoint, error) {
	mdnsResolver, err := zeroconf.NewResolver(nil)
	if err != nil {
		return nil, err
	}

	mdnsServices := make(chan *zeroconf.ServiceEntry)

	err = mdnsResolver.Browse(ctx, "_conreality._tcp", "local.", mdnsServices)
	if err != nil {
		return nil, err
	}

	endpoints := make(chan *GameEndpoint)

	go func(services <-chan *zeroconf.ServiceEntry) {
		for service := range services {
			var version string
			for i := range service.Text {
				if strings.HasPrefix(service.Text[i], "version=") {
					version = service.Text[i][8:]
					break
				}
			}
			endpoints <- &GameEndpoint{
				Name:    service.Instance,
				Host:    service.AddrIPv4[0].String(),
				Port:    service.Port,
				Version: version,
			}
		}
	}(mdnsServices)

	return endpoints, nil
}
