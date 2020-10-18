package commands

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/nimblebun/tsc-language-server/langserver"
	"github.com/nimblebun/tsc-language-server/langserver/handlers"
	"github.com/urfave/cli/v2"
)

// TCPCommand starts the server in stdio mode
func TCPCommand(c *cli.Context) error {
	port := c.Int("port")

	logger := log.New(os.Stdout, "", log.LstdFlags|log.Lshortfile)
	ctx := context.Background()

	server := langserver.New(ctx, handlers.NewSession)
	server.SetLogger(logger)

	err := server.StartTCP(fmt.Sprintf("localhost:%d", port))
	if err != nil {
		return fmt.Errorf("failed to start the language server in TCP mode")
	}

	return nil
}
