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

// StartCommand starts the server in stdio mode
func StartCommand(c *cli.Context) error {
	logger := log.New(os.Stdout, "", log.LstdFlags|log.Lshortfile)
	ctx := context.Background()

	server := langserver.New(ctx, handlers.NewSession)
	server.SetLogger(logger)

	err := server.StartAndWait(os.Stdin, os.Stdout)
	if err != nil {
		return fmt.Errorf("failed to start the language server in stdio mode")
	}

	return nil
}
