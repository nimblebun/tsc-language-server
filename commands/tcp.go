package commands

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/urfave/cli/v2"
	"pkg.nimblebun.works/tsc-language-server/langserver"
	"pkg.nimblebun.works/tsc-language-server/langserver/handlers"
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
