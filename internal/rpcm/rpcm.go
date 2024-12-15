// Package rpcm is for models (structs) that is used for
// net/rpc communication
package rpcm

type Args struct {
	Keys []string
}

type Reply struct {
	Results []string
}

// ShortenerService interface is intended to used like this in server side:
//
// var _ rpcm.ShortenerService = (*Shortener)(nil)
type ShortenerService interface {
	NewUrl(args *Args, reply *Reply) error
	GetUrl(args *Args, reply *Reply) error
}

const (
	NewUrl = "Shortener.NewUrl"
	GetUrl = "Shortener.GetUrl"
)
