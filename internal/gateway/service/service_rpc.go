package service

import "ush/internal/rpcm"

func (s *SrvImpl) newUrlRpc(urls []string) ([]string, error) {
	args := rpcm.Args{Keys: urls}
	reply := rpcm.Reply{}
	err := s.rpcClient.Call(rpcm.NewUrl, &args, &reply)
	if err != nil {
		return nil, err
	}

	return args.Keys, nil
}

func (s *SrvImpl) getUrlRpc(keys []string) ([]string, error) {
	args := rpcm.Args{Keys: keys}
	reply := rpcm.Reply{}
	err := s.rpcClient.Call(rpcm.GetUrl, &args, &reply)
	if err != nil {
		return nil, err
	}

	return args.Keys, nil
}
