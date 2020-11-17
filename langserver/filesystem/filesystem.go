package filesystem

import (
	"io/ioutil"
	"log"
	"os"

	"github.com/spf13/afero"
	"pkg.nimblebun.works/tsc-language-server/langserver/filesystem/filehandler"
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
func (fs *FileSystem) Remove(fh filehandler.FileHandler) error {
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
