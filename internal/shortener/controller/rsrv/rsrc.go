// Package rsrv is short for rpc server (:
package rsrv

import (
	"fmt"
	"log"
	"net"
	"net/rpc"
	"ush/internal/pkg/config"
	"ush/internal/rpcm"
	"ush/internal/shortener/controller"
	"ush/internal/shortener/repository"
)

type Shortener struct {
	repo repository.Repository
	srv  *rpc.Server
	cfg  *config.ShortenerConfig
	l    net.Listener
}

func (r *Shortener) NewUrl(args *rpcm.Args, reply *rpcm.Reply) error {
	//TODO implement me
	panic("implement me")
}

func (r *Shortener) GetUrl(args *rpcm.Args, reply *rpcm.Reply) error {
	//TODO implement me
	panic("implement me")
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

		go r.srv.ServeConn(conn)
	}
}

func (r *Shortener) Stop() error {
	//TODO implement me
	panic("implement me")
}

func New(config *config.ShortenerConfig, repo repository.Repository) (controller.RPC, error) {
	// check that
	var _ rpcm.ShortenerService = (*Shortener)(nil)

	sh := new(Shortener)

	s := rpc.NewServer()
	if err := s.Register(sh); err != nil {
		return nil, err
	}

	sh.srv = s
	sh.cfg = config
	sh.repo = repo

	return sh, nil
}
