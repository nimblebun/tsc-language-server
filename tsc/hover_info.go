package tsc

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/nimblebun/tsc-language-server/config"
	"github.com/nimblebun/tsc-language-server/utils"
)

// isEvent checks whether the text hovered over is an event (#0000, #1234, etc.)
func isEvent(str string) bool {
	// Events are 5 characters long (# and 4 digits)
	if len(str) < 5 {
		return false
	}

	// Events must start with #
	if str[0] != '#' {
		return false
	}

	// The 4 characters after the # must be valid integers
	_, err := strconv.Atoi(str[1:5])
	return err == nil
}

// isValidArgument checks whether the provided argument is valid (4 digits or
// 'V' and 3 digits)
func isValidArgument(arg string) bool {
	// Arguments are always 4 characters long
	if len(arg) != 4 {
		return false
	}

	// Arguments are either valid integers or start with V
	_, err := strconv.Atoi(arg)
	if err != nil && arg[0] != 'V' {
		return false
	}

	// The last 3 characters must always be integers
	_, err = strconv.Atoi(arg[1:])
	return err == nil
}

// GetHoverInfo will retrieve information about an event or a command starting
// from a given position
func GetHoverInfo(str string, at int, config config.Config) string {
	if isEvent(str) {
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

			if isValidArgument(arg) {
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
