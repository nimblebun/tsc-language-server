package uri

import (
	"net/url"
	"path/filepath"
)

// FromPath converts a given path to a valid URI
func FromPath(path string) string {
	p := filepath.ToSlash(path)
	p = normalizePath(p)

	u := &url.URL{
		Scheme: "file",
		Path:   p,
	}

	return u.String()
}
