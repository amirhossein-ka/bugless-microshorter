// Package rsrv is short for rpc server (:
package rsrv

import (
	"errors"
	"fmt"
	"log"
	"net"
	"net/rpc"
	"ush/internal/pkg/config"
	"ush/internal/rpcm"
	"ush/internal/shortener/controller"
	"ush/internal/shortener/service"
)

type Shortener struct {
	shortenerService service.ShortenerService
	server           *rpc.Server
	cfg              *config.ShortenerConfig
	l                net.Listener
}

func (r *Shortener) NewUrl(args *rpcm.Args, reply *rpcm.Reply) error {
	if len(args.Keys) == 0 {
		return errors.New("what you wanna put in")
	}
	if len(args.Keys) > 1 {
		return errors.New("you can only add one url per request pls wait for more updates")
	}

	key, err := r.shortenerService.AddUrl(args.Keys[0])
	if err != nil {
		return err
	}

	reply.Results = append(reply.Results, key)

	return nil
}

func (r *Shortener) GetUrl(args *rpcm.Args, reply *rpcm.Reply) error {
	if len(args.Keys) == 0 {
		return errors.New("what you looking for dude")
	}
	if len(args.Keys) > 1 {
		return errors.New("you can only add one url per request pls wait for more updates")
	}

	url, err := r.shortenerService.GetUrl(args.Keys[0])
	if err != nil {
		return err
	}

	reply.Results = append(reply.Results, url)

	return nil
}

func (r *Shortener) Start(s string) error {

	listener, err := net.Listen("tcp", fmt.Sprintf("0.0.0.0:%d", r.cfg.ListenPort))
	if err != nil {
		return err
	}

	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Println(err)
			continue
		}

		go r.server.ServeConn(conn)
	}
}

func (r *Shortener) Stop() error {
	if err := r.shortenerService.Stop(); err != nil {
		return nil
	}

	if err := r.l.Close(); err != nil {
		return err
	}
	return nil
}

func New(config *config.ShortenerConfig, srv service.ShortenerService) (controller.RPC, error) {
	// check that what i wrote implements the thing in rpcm package
	var _ rpcm.ShortenerService = (*Shortener)(nil)

	sh := new(Shortener)

	s := rpc.NewServer()
	if err := s.Register(sh); err != nil {
		return nil, err
	}

	sh.server = s
	sh.cfg = config
	sh.shortenerService = srv

	return sh, nil
}
