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
