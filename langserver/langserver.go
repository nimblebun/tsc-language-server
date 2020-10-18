package langserver

import (
	"context"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net"

	"github.com/creachadair/jrpc2"
	"github.com/creachadair/jrpc2/channel"
	"github.com/creachadair/jrpc2/server"
	"github.com/nimblebun/tsc-language-server/langserver/instance"
	"github.com/nimblebun/tsc-language-server/langserver/session"
)

// LanguageServer contains definitions of JSON-RPC options, context, and logger
// to be used when instantiating a new language server
type LanguageServer struct {
	ctx            context.Context
	logger         *log.Logger
	opts           *jrpc2.ServerOptions
	sessionFactory session.Factory
}

// New instantiates a language server to be used in a JSON-RPC context
func New(ctx context.Context, sessionFactory session.Factory) *LanguageServer {
	opts := &jrpc2.ServerOptions{
		AllowPush:   true,
		Concurrency: 1,
	}

	return &LanguageServer{
		ctx:            ctx,
		logger:         log.New(ioutil.Discard, "", 0),
		opts:           opts,
		sessionFactory: sessionFactory,
	}
}

// SetLogger will overwrite the current logger of the language server instance
func (s *LanguageServer) SetLogger(logger *log.Logger) {
	s.opts.Logger = logger
	s.logger = logger
}

func (s *LanguageServer) newSevice() server.Service {
	service := s.sessionFactory(s.ctx)
	service.SetLogger(s.logger)
	return service
}

func (s *LanguageServer) start(reader io.Reader, writer io.WriteCloser) (*instance.ServerInstance, error) {
	server, err := CreateInstance(s.newSevice(), s.opts)
	if err != nil {
		return nil, err
	}

	server.Start(channel.LSP(reader, writer))

	return server, nil
}

// StartAndWait will start a new language server
func (s *LanguageServer) StartAndWait(reader io.Reader, writer io.WriteCloser) error {
	server, err := s.start(reader, writer)

	if err != nil {
		return err
	}

	s.logger.Print("Starting server...")

	ctx, cancelFunc := context.WithCancel(s.ctx)
	go func() {
		server.Wait()
		cancelFunc()
	}()

	select {
	case <-ctx.Done():
		s.logger.Printf("Stopping server...")
		server.Stop()
	}

	s.logger.Printf("Server stopped.")
	return nil
}

// StartTCP will start a language server in TCP mode
func (s *LanguageServer) StartTCP(addr string) error {
	s.logger.Printf("Starting TCP server (%q)...", addr)

	listener, err := net.Listen("tcp", addr)
	if err != nil {
		return fmt.Errorf("Failed to start TCP server: %s", err)
	}

	s.logger.Printf("TCP server running (%q)...", listener.Addr())

	go func() {
		s.logger.Println("Starting loop server...")
		err = server.Loop(listener, s.newSevice, &server.LoopOptions{
			Framing:       channel.LSP,
			ServerOptions: s.opts,
		})

		if err != nil {
			s.logger.Printf("Failed to start loop server: %s", err)
		}
	}()

	select {
	case <-s.ctx.Done():
		s.logger.Printf("Stopping TCP server...")
		err = listener.Close()
		if err != nil {
			s.logger.Printf("Failed to stop TCP server: %s", err)
			return err
		}
	}

	s.logger.Printf("TCP server stopped.")
	return nil
}

// CreateInstance will instantiate a JSON-RPC server
func CreateInstance(service server.Service, opts *jrpc2.ServerOptions) (*instance.ServerInstance, error) {
	assigner, err := service.Assigner()
	if err != nil {
		return nil, err
	}

	return &instance.ServerInstance{
		Server:     jrpc2.NewServer(assigner, opts),
		CancelFunc: service.Finish,
	}, nil
}
