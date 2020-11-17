package filesystem

import (
	"github.com/sourcegraph/go-lsp"
	"pkg.nimblebun.works/tsc-language-server/langserver/filesystem/filehandler"
)

// File is a structure that contains fields and methods that the LSP uses for
// handling document files.
type File struct {
	Handler *filehandler.FileHandler
	Text    []byte
	Version int
}

// URI returns the URI from the file handler.
func (f *File) URI() string {
	return f.Handler.URI
}

// FullPath returns the file's full path from the file handler.
func (f *File) FullPath() (string, error) {
	return f.Handler.FullPath()
}

// Dir returns the path to the directory of the file from the file handler.
func (f *File) Dir() (string, error) {
	return f.Handler.Dir()
}

// Filename returns the name of the file from the file handler.
func (f *File) Filename() (string, error) {
	return f.Handler.Filename()
}

// FileFromDocumentItem will create a File from a given LSP textdocument item.
func FileFromDocumentItem(doc lsp.TextDocumentItem) *File {
	return &File{
		Handler: filehandler.FromDocumentURI(doc.URI),
		Text:    []byte(doc.Text),
		Version: doc.Version,
	}
}
