package controller

// Controller is just a simple interface to start/stop http server
// I do this so in future, if its needed to use another library for handling http requests,
// it would be easier to migrate
type Controller interface {
	Start(string) error
	Stop()
}
