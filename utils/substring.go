package utils

import "math"

// Substring will return a substring of given length from a given index. If
// the from index is bigger than the length of the string, it will return an
// empty string. If the given length would result in an OOB error, the function
// will return the substring up to the length of the string. If the from index
// or the length is a negative value, it will return an empty string. If the
// length is 0, it will return everything starting from the provided index to
// the end of the string.
func Substring(str string, from int, length int) string {
	if from > len(str) || from < 0 || length < 0 {
		return ""
	}

	if length == 0 {
		return str[from:]
	}

	to := int(math.Min(float64(from+length), float64(len(str))))
	return str[from:to]
}
