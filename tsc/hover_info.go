package tsc

import (
	"fmt"
	"strings"

	"github.com/nimblebun/tsc-language-server/config"
	"github.com/nimblebun/tsc-language-server/utils"
)

// GetHoverInfo will retrieve information about an event or a command starting
// from a given position
func GetHoverInfo(str string, at int, config config.Config) string {
	if IsEvent(str) {
		return fmt.Sprintf("Event **%s**", str[:5])
	}

	commandTokenIdx := strings.LastIndex(str[at:], "<")

	if commandTokenIdx == -1 {
		return ""
	}

	targetCommand := utils.Substring(str, commandTokenIdx, 4)
	command, found := config.GetTSCDefinition(targetCommand)

	if !found {
		return ""
	}

	info := fmt.Sprintf("Command **%s**\n\n%s", command.Format, command.Documentation)

	args := []string{}
	strWithoutCommand := utils.Substring(str, commandTokenIdx+4, 0)

	if command.Nargs > 0 {
		for i := 0; i < command.Nargs; i++ {
			arg := utils.Substring(strWithoutCommand, i*5, 4)

			if IsValidArgument(arg) {
				args = append(args, arg)
			}
		}

		if len(args) == command.Nargs {
			for i := 0; i < command.Nargs; i++ {
				arg := args[i]
				value := config.GetArgumentValue(command, i, arg)

				if arg != value {
					args[i] = fmt.Sprintf("%s: %s", arg, value)
				}
			}
		}

		info += fmt.Sprintf("\n\n* %s", strings.Join(args, "\n* "))
	}

	return info
}
