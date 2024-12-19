// Package service is where all magic happens for gateway
// including sending rpc requests, handling requests duplicates,
// the queue and 50ms wait mechanism, and maybe more stuff !
package service

import (
	"fmt"
	"net/rpc"
	"ush/pkg/cache"
	"ush/pkg/config"
)

type (

	// Service I know using interfaces too much can cause smoll performance problems,
	// But I just do it because its more future-proof, like it will be easier
	// to change logic behind it.
	// Also, it's easier for me to know what should I implement (:
	Service interface {
		AddUrl(string) (string, error)
		GetFullURL(string) (string, error)
	}

	// srvImpl implements Service
	// It should contain the queue/wait things, stuff for rpc calls, and ofc some helper methods to do those stuff
	srvImpl struct {
		config    *config.GatewayConfig
		cache     cache.LRUCache[string, string]
		rpcClient *rpc.Client
	}
)

func NewService(cfg *config.GatewayConfig) (Service, error) {
	client, err := rpc.DialHTTP("tcp", fmt.Sprintf("%s:%d", cfg.ShortenerHost, cfg.ShortenerPort))
	if err != nil {
		return nil, err
	}
	lru, err := cache.NewLRU[string, string](cfg.CacheSize, nil)
	if err != nil {
		return nil, err
	}

	return &srvImpl{
		rpcClient: client,
		cache:     lru,
		config:    cfg,
	}, nil
}
