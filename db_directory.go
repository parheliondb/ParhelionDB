package parhelion

import (
	"github.com/parheliondb/ParhelionDB/util"
)

type dbDirectory struct {
	Path string
}

type DBDirectory interface {
	Close() error
	GetPath() string
	// ListDataFiles() []string
	// ListIndexFiles() []int
	// ListTombstoneFiles() []string
	SyncMetaData() error
}

func NewDBDirectory(path string) (DBDirectory, error) {
	err := util.CreateDirectoryIfNotExists(path)
	if err != nil {
		return nil, err
	}

	return &dbDirectory{
		Path: path,
	}, nil
}

func (d *dbDirectory) Close() error {
	return nil
}

func (d *dbDirectory) GetPath() string {
	return d.Path
}

// func (d *dbDirectory) ListDataFiles() []string {
//
// }
//
// func (d *dbDirectory) ListIndexFiles() []int {
//
// }
//
// func (d *dbDirectory) ListTombstoneFiles() []string {
//
// }

func (d *dbDirectory) SyncMetaData() error {
	return nil
}
