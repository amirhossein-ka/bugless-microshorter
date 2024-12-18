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

	r.l = listener

	for {
		conn, err := listener.Accept()
		if err != nil {
			if errors.Is(err, net.ErrClosed) {
				log.Println("Listener closed, stopping accept loop")
				break
			}
			log.Println("Accept error:", err)
			continue
		}

		go r.server.ServeConn(conn)
	}

	return nil
}

func (r *Shortener) Stop() error {
	var err error

	if r.shortenerService != nil {
		if serr := r.shortenerService.Stop(); serr != nil {
			err = fmt.Errorf("failed to stop shortener service: %w", serr)
		}
	}

	if r.l != nil {
		if cerr := r.l.Close(); cerr != nil {
			if err != nil {
				err = fmt.Errorf("%s; failed to close listener: %w", err, cerr)
			} else {
				err = fmt.Errorf("failed to close listener: %w", cerr)
			}
		}
	}

	return err
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
