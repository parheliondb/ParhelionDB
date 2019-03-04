package internal

import (
	"sync"

	"github.com/juju/fslock"
	parhelion "github.com/parheliondb/ParhelionDB"
)

type parhelionDB struct {
	DBDirectory parhelion.DBDirectory
	Options     parhelion.ParhelionDBOptions
	FileLock    *fslock.Lock
	WriteLock   *sync.Mutex
}

type ParhelionDB interface {
	Get(key []byte, attemptNumber int) ([]byte, error)
	Put(key []byte, value []byte) bool
	Delete(key []byte) error
	Close() error
	Size() int64
	SetIOErrorFlag() error
	PauseCompaction() error
	ResumeCompaction() error
}

func NewParhelionDB(dirName string, options parhelion.ParhelionDBOptions) (ParhelionDB, error) {
	dbDirectory, err := parhelion.NewDBDirectory(dirName)
	if err != nil {
		return nil, err
	}

	fileLock := fslock.New(dbDirectory.GetPath())
	err = fileLock.TryLock()
	if err != nil {
		return nil, err
	}

	return &parhelionDB{
		DBDirectory: dbDirectory,
		Options:     options,
		FileLock:    fileLock,
		WriteLock:   new(sync.Mutex),
	}, nil
}

func (p *parhelionDB) Get(key []byte, attemptNumber int) ([]byte, error) {

	return nil, nil
}

func (p *parhelionDB) Put(key []byte, value []byte) bool {
	return false
}

func (p *parhelionDB) Delete(key []byte) error {
	return nil
}

func (p *parhelionDB) Close() error {
	if p.FileLock != nil {
		p.FileLock.Unlock()
	}

	return nil
}

func (p *parhelionDB) Size() int64 {
	return 0
}

func (p *parhelionDB) SetIOErrorFlag() error {
	return nil
}

func (p *parhelionDB) PauseCompaction() error {
	return nil
}

func (p *parhelionDB) ResumeCompaction() error {
	return nil
}
