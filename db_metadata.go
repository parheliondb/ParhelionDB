package parhelion

import (
	"encoding/gob"
	"hash/crc32"
	"os"
	"sync"
)

type DBMetadata struct {
	Checksum       uint32
	Version        uint8
	Open           bool
	SequenceNumber uint64
	MaxFileSize    uint64
}

type dbMetadataLoader struct {
	mu  *sync.Mutex
	dir DBDirectory
}

type DBMetadataLoader interface {
	Load() (*DBMetadata, error)
	Store(DBMetadata) error
}

func NewDBMetadataLoader(dir DBDirectory) DBMetadataLoader {
	return &dbMetadataLoader{
		mu:  &sync.Mutex{},
		dir: dir,
	}
}

func (l *dbMetadataLoader) Load() (*DBMetadata, error) {
	filename := l.dir.GetPath() + "/" + MetadataFileName

	l.mu.Lock()
	defer l.mu.Unlock()

	_, err := os.Stat(filename)
	if err != nil {
		if !os.IsNotExist(err) {
			return nil, err
		}
		return &DBMetadata{}, err
	}

	f, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	var m DBMetadata
	if err := gob.NewDecoder(f).Decode(&m); err != nil {
		return nil, err
	}

	return &m, nil
}

func (l *dbMetadataLoader) Store(m DBMetadata) error {
	filename := l.dir.GetPath() + "/" + MetadataFileName

	l.mu.Lock()
	defer l.mu.Unlock()

	f, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer f.Close()

	return gob.NewEncoder(f).Encode(m)
}

func computeChecksum(bs []byte) uint32 {
	return crc32.ChecksumIEEE(bs)
}
