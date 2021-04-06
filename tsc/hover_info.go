package tsc

import (
	"fmt"
	"strings"

	"pkg.nimblebun.works/tsc-language-server/config"
	"pkg.nimblebun.works/tsc-language-server/utils"
)

// GetHoverInfo will retrieve information about an event or a command starting
// from a given position
func GetHoverInfo(str string, to int, conf *config.Config) string {
	if IsEvent(str, conf.Setup.LooseChecking.Events) {
		return fmt.Sprintf("Event **%s**", str[:5])
	}

	commandTokenIdx := strings.LastIndex(str[:to], "<")

	if commandTokenIdx == -1 {
		return ""
	}

	targetCommand := utils.Substring(str, commandTokenIdx, 4)
	command, found := conf.GetTSCDefinition(targetCommand)

	if !found {
		return ""
	}

	info := fmt.Sprintf("Command **%s**\n\n%s", command.Format, command.Documentation)

	args := []string{}
	strWithoutCommand := utils.Substring(str, commandTokenIdx+4, 0)

	if command.Nargs() > 0 {
		for i := 0; i < command.Nargs(); i++ {
			arg := utils.Substring(strWithoutCommand, i*5, 4)

			if IsValidArgument(arg, conf.Setup.LooseChecking.Arguments) {
				args = append(args, arg)
			}
		}

		if len(args) == command.Nargs() {
			for i := 0; i < command.Nargs(); i++ {
				arg := args[i]
				value := conf.GetArgumentValue(command, i, arg)

				if arg != value {
					args[i] = fmt.Sprintf("%s: %s", arg, value)
				}
			}
		}

		info += fmt.Sprintf("\n\n* %s", strings.Join(args, "\n* "))
	}

	return info
}
