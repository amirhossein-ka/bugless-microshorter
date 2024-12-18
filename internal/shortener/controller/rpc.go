package controller

type RPC interface {
	Start(string) error
	Stop() error
}
