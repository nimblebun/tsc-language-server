package tsc

import "strconv"

// IsEvent checks whether the text hovered over is an event (#0000, #1234, etc.)
func IsEvent(str string) bool {
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

// IsValidArgument checks whether the provided argument is valid (4 digits or
// 'V' and 3 digits)
func IsValidArgument(arg string) bool {
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
