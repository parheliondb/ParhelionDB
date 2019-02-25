package parhelion

type parhelionDB struct{}

type ParhelionDB interface {
	Get(key []byte) ([]byte, error)
	Put(key []byte, value []byte) error
	Delete(key []byte) error
	Close() error
	Size() int64
	// Stats() ParhelionDBStats
	ResetStats() error
	// NewIterator() ParhelionDBIterator
	PauseCompaction() error
	ResumeCompaction() error
}

func NewParhelionDB(dirName string, options ParhelionDBOptions) (ParhelionDB, error) {
	return &parhelionDB{}, nil
}

func (p *parhelionDB) Get(key []byte) ([]byte, error) {
	return nil, nil
}

func (p *parhelionDB) Put(key []byte, value []byte) error {
	return nil
}

func (p *parhelionDB) Delete(key []byte) error {
	return nil
}

func (p *parhelionDB) Close() error {
	return nil
}

func (p *parhelionDB) Size() int64 {
	return 0
}

// func (p *parhelionDB) Stats() ParhelionDBStats {
//
// }

func (p *parhelionDB) ResetStats() error {
	return nil
}

// func (p *parhelionDB) NewIterator() ParhelionDBIterator {
//
// }

func (p *parhelionDB) PauseCompaction() error {
	return nil
}

func (p *parhelionDB) ResumeCompaction() error {
	return nil
}
