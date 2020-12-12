package filesystem

import (
	"bytes"
	"io/ioutil"
	"log"
	"os"

	"github.com/spf13/afero"
	"pkg.nimblebun.works/go-lsp"
	"pkg.nimblebun.works/tsc-language-server/langserver/filesystem/filehandler"
	"pkg.nimblebun.works/tsc-language-server/utils"
)

// FileSystem is a file I/O pool that uses afero to create an in-memory and
// an OS file system to store the documents in.
type FileSystem struct {
	MemoryFS afero.Fs
	OsFS     afero.Fs

	Logger *log.Logger
}

// New instantiates a new file system.
func New() *FileSystem {
	return &FileSystem{
		MemoryFS: afero.NewMemMapFs(),
		OsFS:     afero.NewReadOnlyFs(afero.NewOsFs()),

		Logger: log.New(ioutil.Discard, "", 0),
	}
}

// SetLogger will override the logger of the filesystem instance.
func (fs *FileSystem) SetLogger(logger *log.Logger) {
	fs.Logger = logger
}

// Create will create a file in the file system and add its metadata to the
// metadata map.
func (fs *FileSystem) Create(fh filehandler.FileHandler, contents []byte) error {
	path, err := fh.FullPath()
	if err != nil {
		return err
	}

	f, err := fs.MemoryFS.Create(path)
	if err != nil {
		return err
	}

	_, err = f.Write(contents)
	if err != nil {
		return err
	}

	return nil
}

// ReadFile will read the contents of a file with a given name.
func (fs *FileSystem) ReadFile(name string) ([]byte, error) {
	buf, err := afero.ReadFile(fs.MemoryFS, name)
	if err != nil && os.IsNotExist(err) {
		return afero.ReadFile(fs.OsFS, name)
	}

	return buf, err
}

// Change will update a file with the changes provided by the language client.
func (fs *FileSystem) Change(handler *filehandler.VersionedFileHandler, changes []lsp.TextDocumentContentChangeEvent) error {
	if len(changes) == 0 {
		return nil
	}

	path, err := handler.FullPath()
	if err != nil {
		return err
	}

	file, err := fs.MemoryFS.OpenFile(path, os.O_RDWR, 0700)
	if err != nil {
		return err
	}

	defer file.Close()

	var buffer bytes.Buffer
	_, err = buffer.ReadFrom(file)
	if err != nil {
		return err
	}

	for _, change := range changes {
		err := fs.applyChange(&buffer, change)
		if err != nil {
			return err
		}
	}

	err = file.Truncate(0)
	if err != nil {
		return err
	}

	_, err = file.Seek(0, 0)
	if err != nil {
		return err
	}

	_, err = file.Write(buffer.Bytes())
	if err != nil {
		return err
	}

	return nil
}

// Open will open the raw file to perform I/O operations on it.
func (fs *FileSystem) Open(name string) (afero.File, error) {
	f, err := fs.MemoryFS.Open(name)
	if err != nil && os.IsNotExist(err) {
		return fs.OsFS.Open(name)
	}

	return f, err
}

// Remove will mark the file as closed and then it will remove the file
// from the pool.
func (fs *FileSystem) Remove(fh *filehandler.FileHandler) error {
	path, err := fh.FullPath()
	if err != nil {
		return err
	}

	err = fs.MemoryFS.Remove(path)
	if err != nil {
		return err
	}

	return nil
}

func (fs *FileSystem) applyChange(buffer *bytes.Buffer, change lsp.TextDocumentContentChangeEvent) error {
	if change.Range == nil {
		buffer.Reset()
		_, err := buffer.WriteString(change.Text)
		return err
	}

	b := buffer.Bytes()

	from := utils.DocumentOffset(b, change.Range.Start)
	to := utils.DocumentOffset(b, change.Range.End)

	delta := to - from
	if delta > 0 {
		buffer.Grow(delta)
	}

	beforeChange := make([]byte, from, from)
	copy(beforeChange, buffer.Bytes())
	afterBytes := buffer.Bytes()[to:]
	afterChange := make([]byte, len(afterBytes), len(afterBytes))
	copy(afterChange, afterBytes)

	buffer.Reset()

	_, err := buffer.Write(beforeChange)
	if err != nil {
		return err
	}

	_, err = buffer.WriteString(change.Text)
	if err != nil {
		return err
	}

	_, err = buffer.Write(afterChange)
	if err != nil {
		return err
	}

	return nil
}
