package filehandler

import (
	"path/filepath"
	"strings"
)

// FullPath will attempt to parse and create a full path from the given
// file handler. This method is different from the UNIX method because it will
// also trim the leading "/" on Windows.
func (fh *FileHandler) FullPath() (string, error) {
	path, err := fh.ParsePath()
	if err != nil {
		return "", err
	}

	path = strings.TrimPrefix(path, "/")
	return filepath.FromSlash(path), nil
}
