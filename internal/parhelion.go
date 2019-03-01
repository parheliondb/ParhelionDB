package internal

import (
	"sync"

	parhelion "github.com/parheliondb/ParhelionDB"
)

type parhelionDBInternal struct {
	DBDirectory parhelion.DBDirectory
	Options     parhelion.ParhelionDBOptions
	WriteLock   sync.Mutex
}

type ParhelionDBInternal interface {
	Get(key []byte, attemptNumber int) ([]byte, error)
	Put(key []byte, value []byte) bool
	Delete(key []byte) error
	Close() error
	Size() int64
	SetIOErrorFlag() error
	PauseCompaction() error
	ResumeCompaction() error
}

func NewParhelionDBInternal(dirName string, options parhelion.ParhelionDBOptions) (ParhelionDBInternal, error) {
	dbDirectory, err := parhelion.NewDBDirectory(dirName)
	if err != nil {
		return nil, err
	}

	return &parhelionDBInternal{
		DBDirectory: dbDirectory,
		Options:     options,
		WriteLock:   *new(sync.Mutex),
	}, nil
}

func (p *parhelionDBInternal) Get(key []byte, attemptNumber int) ([]byte, error) {

	return nil, nil
}

func (p *parhelionDBInternal) Put(key []byte, value []byte) bool {
	return false
}

func (p *parhelionDBInternal) Delete(key []byte) error {
	return nil
}

func (p *parhelionDBInternal) Close() error {
	return nil
}

func (p *parhelionDBInternal) Size() int64 {
	return 0
}

func (p *parhelionDBInternal) SetIOErrorFlag() error {
	return nil
}

func (p *parhelionDBInternal) PauseCompaction() error {
	return nil
}

func (p *parhelionDBInternal) ResumeCompaction() error {
	return nil
}
