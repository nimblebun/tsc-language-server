// +build !windows

package filehandler

import "path/filepath"

// FullPath will attempt to parse and create a full path from the given
// file handler.
func (fh *FileHandler) FullPath() (string, error) {
	path, err := fh.ParsePath()
	if err != nil {
		return "", err
	}

	return filepath.FromSlash(path), nil
}
