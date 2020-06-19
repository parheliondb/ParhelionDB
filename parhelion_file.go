package parhelion

import (
	"bytes"
	"fmt"
	"os"
	"strconv"

	"github.com/parheliondb/ParhelionDB/memory"
)

type parhelionDBFile struct {
	fileID      uint64
	backingFile string
	dir         DBDirectory
	// indexFile IndexFile
	// TODO
}

type ParhelionDBFile interface {
	ReadFromFile(offset, length int64) ([]byte, error)
}

func NewParhelionDBFile(dir DBDirectory, filename string) (ParhelionDBFile, error) {
	fid, err := GetFileTimestamp(filename)
	if err != nil {
		return nil, err
	}

	return &parhelionDBFile{
		fileID:      fid,
		backingFile: filename,
		dir:         dir,
	}, nil
}

func GetFileTimestamp(filename string) (uint64, error) {
	ss := DataFilePattern.FindStringSubmatch(filename)

	if len(ss) == 0 {
		return 0, fmt.Errorf("filename pattern is invalid")
	}

	i, err := strconv.Atoi(ss[1])
	if err != nil {
		return 0, err
	}

	return uint64(i), nil
}

func (pf *parhelionDBFile) ReadFromFile(offset, length int64) ([]byte, error) {
	f, err := os.Open(pf.backingFile)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	_, err = f.Seek(offset, 0)
	if err != nil {
		return nil, err
	}

	bs := make([]byte, length)
	_, err = f.Read(bs)
	if err != nil {
		return nil, err
	}

	return bs, nil
}

func (pf *parhelionDBFile) readRecord(offset int64) (Record, error) {
	bs, err := pf.ReadFromFile(offset, RecordHeaderSize)
	if err != nil {
		return nil, err
	}
	offset += RecordHeaderSize

	rh, err := NewRecordHeaderFromBuffer(*bytes.NewBuffer(bs))
	if err != nil {
		return nil, err
	}

	bs, err = pf.ReadFromFile(offset, int64(rh.KeySize())+int64(rh.ValueSize()))
	if err != nil {
		return nil, err
	}

	r, err := NewRecordFromBuffer(*bytes.NewBuffer(bs), rh.KeySize(), rh.ValueSize())
	if err != nil {
		return nil, err
	}

	r.SetHeader(rh)
	valueOffset := uint64(offset) + uint64(rh.KeySize())
	r.SetMetadata(memory.NewMetadata(pf.fileID, valueOffset, rh.ValueSize(), rh.SequenceNumber()))

	return r, nil
}

func (pf *parhelionDBFile) writeRecord(rec Record) (memory.Metadata, error) {
	// writeToChannel

	return nil, nil
}
