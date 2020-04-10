package parhelion

import (
	"io/ioutil"
	"regexp"

	"github.com/parheliondb/ParhelionDB/util"
)

type dbDirectory struct {
	path string
}

type DBDirectory interface {
	GetPath() string
	ListDataFiles() ([]string, error)
	ListIndexFiles() ([]string, error)
	ListTombstoneFiles() ([]string, error)
	SyncMetadata() error
}

func NewDBDirectory(path string) (DBDirectory, error) {
	err := util.CreateDirectoryIfNotExists(path)
	if err != nil {
		return nil, err
	}

	return &dbDirectory{
		path: path,
	}, nil
}

func (d *dbDirectory) GetPath() string {
	return d.path
}

func (d *dbDirectory) ListDataFiles() ([]string, error) {
	return d.findFilesByRegexpPattern(DataFilePattern)
}

func (d *dbDirectory) ListIndexFiles() ([]string, error) {
	return d.findFilesByRegexpPattern(IndexFilePattern)
}

func (d *dbDirectory) ListTombstoneFiles() ([]string, error) {
	return d.findFilesByRegexpPattern(TombstoneFilePattern)
}

func (d *dbDirectory) findFilesByRegexpPattern(re *regexp.Regexp) ([]string, error) {
	files, err := ioutil.ReadDir(d.path)
	if err != nil {
		return nil, err
	}

	filtered := make([]string, 0, len(files))
	for _, file := range files {
		if re.Match([]byte(file.Name())) {
			filtered = append(filtered, file.Name())
		}
	}

	return filtered, nil
}

func (d *dbDirectory) SyncMetadata() error {
	// TODO: should be considered how to achieve ReadOnlyFileChannel in golang
	return nil
}
