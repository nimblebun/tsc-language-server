package instance

import (
	"github.com/creachadair/jrpc2"
	"github.com/creachadair/jrpc2/channel"
)

// ServerInstance is a jrpc2.Server with support for server.Service
type ServerInstance struct {
	Server     *jrpc2.Server
	CancelFunc func(jrpc2.Assigner, jrpc2.ServerStatus)
}

// Start will call jrpc2.Server's Start method
func (instance *ServerInstance) Start(ch channel.Channel) {
	instance.Server = instance.Server.Start(ch)
}

// Wait will cancel the current instance with the server's wait status
func (instance *ServerInstance) Wait() {
	status := instance.Server.WaitStatus()
	instance.CancelFunc(nil, status)
}

// StartAndWait will start the server and then put it in "wait" status
func (instance *ServerInstance) StartAndWait(ch channel.Channel) {
	instance.Start(ch)
	instance.Wait()
}

// Stop will call jrpc2.Server's Stop method
func (instance *ServerInstance) Stop() {
	instance.Server.Stop()
}
