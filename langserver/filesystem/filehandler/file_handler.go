package filehandler

import (
	"net/url"
	"path/filepath"
	"strings"

	"pkg.nimblebun.works/go-lsp"
	"pkg.nimblebun.works/tsc-language-server/langserver/filesystem/uri"
)

// FileHandler is a type that makes it easy to access files from within the
// language server.
type FileHandler struct {
	URI   string
	IsDir bool
}

// ParsePath will parse the file handler's URI and return it if it's valid.
func (fh *FileHandler) ParsePath() (string, error) {
	u, err := url.ParseRequestURI(string(fh.URI))
	if err != nil {
		return "", err
	}

	return url.PathUnescape(u.Path)
}

// Valid is a slim wrapper around FileHandler#ParsePath.
func (fh *FileHandler) Valid() bool {
	_, err := fh.ParsePath()
	return err != nil
}

// Dir will return the path of filehandler's directory.
func (fh *FileHandler) Dir() (string, error) {
	if fh.IsDir {
		return fh.FullPath()
	}

	path, err := fh.FullPath()
	if err != nil {
		return "", err
	}

	return filepath.Dir(path), nil
}

// Filename will return the path of the filehandler's file.
func (fh *FileHandler) Filename() (string, error) {
	path, err := fh.FullPath()
	if err != nil {
		return "", err
	}

	return filepath.Base(path), nil
}

// DocumentURI will convert the filehandler's URI into a valid LSP-compatible
// Document URI
func (fh *FileHandler) DocumentURI() lsp.DocumentURI {
	return lsp.DocumentURI(fh.URI)
}

// VersionedFileHandler is an extension of FileHandler with added support for
// a version parameter.
type VersionedFileHandler struct {
	FileHandler
	Version int
}

// FromDocumentURI will return a FileHandler from a given LSP document URI.
func FromDocumentURI(uri lsp.DocumentURI) *FileHandler {
	return &FileHandler{URI: string(uri)}
}

// FromDirURI will return a FileHandler with IsDir set to true from a given LSP
// document URI. It will also normalize the path before passing it onto the
// struct field.
func FromDirURI(uri lsp.DocumentURI) *FileHandler {
	dirURI := strings.TrimSuffix(string(uri), "/")

	return &FileHandler{
		URI:   dirURI,
		IsDir: true,
	}
}

// FromPath will return a FileHandler from a given path.
func FromPath(path string) *FileHandler {
	return &FileHandler{URI: uri.FromPath(path)}
}

// FromDirPath will return a FileHandler with IsDir set to true from a given
// path.
func FromDirPath(path string) *FileHandler {
	path = filepath.Clean(path)

	return &FileHandler{
		URI:   uri.FromPath(path),
		IsDir: true,
	}
}

// VersionedFileHandlerFromDocument will return a file handler from a given LSP
// document.
func VersionedFileHandlerFromDocument(document lsp.VersionedTextDocumentIdentifier) *VersionedFileHandler {
	return &VersionedFileHandler{
		FileHandler: FileHandler{URI: string(document.URI)},
		Version:     document.Version,
	}
}
