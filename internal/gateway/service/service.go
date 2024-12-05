// This pacakge is where all magic happens for gateway
// including sending rpc requests, handling requests duplicates,
// the queue and 50ms wait mechanism, and maybe more stuff !
package service

type (
	// I know using interfaces too much can cause smoll performance problems,
	// But I just do it because its more future proof, like it will be easier
	// to change logic behind it.
	// Also its easier for me to know what should I implement (:
	Service interface {
		AddUrl(string) error
		GetFullURL(string) (string, error)
	}

	// SrvImp implements Service
	// It should contain the queue/wait things, stuff for rpc calls, and ofc some helper methodes to do those stuff
	SrvImpl struct{}
)

func NewService(...any) Service {
	return &SrvImpl{}
}
