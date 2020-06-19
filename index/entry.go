package index

import (
	"bytes"
	"unsafe"
)

const (
	// crc              - 4 bytes.
	// version          - 1 byte.
	// key size         - 1 bytes.
	// record size      - 4 bytes.
	// record offset    - 4 bytes.
	// sequence number  - 8 bytes.

	IndexChecksumSize       = 4
	IndexVersionSize        = 1
	IndexKeySizeSize        = 1
	IndexRecordSizeSize     = 4
	IndexRecordOffsetSize   = 4
	IndexSequenceNumberSize = 8

	IndexHeaderSize = IndexChecksumSize +
		IndexVersionSize +
		IndexKeySizeSize +
		IndexRecordSizeSize +
		IndexRecordOffsetSize +
		IndexSequenceNumberSize
)

type entryHeader struct {
	checkSum       uint32
	version        uint8
	keySize        uint8
	recordSize     uint32
	recordOffset   uint32
	sequenceNumber int64
}

type EntryHeader interface {
	Serialize() (*bytes.Buffer, error)
}

// TODO: checkSum required?
func NewEntryHeader(
	checkSum uint32,
	version, keySize uint8,
	recordSize, recordOffset uint32,
	sequenceNumber int64,
) EntryHeader {
	return &entryHeader{
		checkSum:       checkSum,
		version:        version,
		keySize:        keySize,
		recordSize:     recordSize,
		recordOffset:   recordOffset,
		sequenceNumber: sequenceNumber,
	}
}

func (eh *entryHeader) Serialize() (*bytes.Buffer, error) {
	buf := bytes.NewBuffer(make([]byte, IndexHeaderSize))

	_, err := buf.Write(*(*[]byte)(unsafe.Pointer(&eh.checkSum)))
	if err != nil {
		return nil, err
	}
	_, err = buf.Write(*(*[]byte)(unsafe.Pointer(&eh.version)))
	if err != nil {
		return nil, err
	}
	_, err = buf.Write(*(*[]byte)(unsafe.Pointer(&eh.keySize)))
	if err != nil {
		return nil, err
	}
	_, err = buf.Write(*(*[]byte)(unsafe.Pointer(&eh.recordSize)))
	if err != nil {
		return nil, err
	}
	_, err = buf.Write(*(*[]byte)(unsafe.Pointer(&eh.recordOffset)))
	if err != nil {
		return nil, err
	}
	_, err = buf.Write(*(*[]byte)(unsafe.Pointer(&eh.sequenceNumber)))
	if err != nil {
		return nil, err
	}

	return buf, nil
}

type entry struct {
	key    []byte
	header EntryHeader
}

type Entry interface {
	Serialize() (*bytes.Buffer, error)
}

func NewEntry(
	key []byte,
	checkSum uint32,
	version, keySize uint8,
	recordSize, recordOffset uint32,
	sequenceNumber int64,
) Entry {
	return &entry{
		key: key,
		header: NewEntryHeader(
			checkSum,
			version,
			keySize,
			recordSize,
			recordOffset,
			sequenceNumber,
		),
	}
}

func (e *entry) Serialize() (*bytes.Buffer, error) {
	buf := bytes.NewBuffer(make([]byte, IndexHeaderSize+len(e.key)))

	// TODO: compute checksum?
	headerBuf, err := e.header.Serialize()
	if err != nil {
		return nil, err
	}
	_, err = buf.Write(headerBuf.Bytes())
	if err != nil {
		return nil, err
	}

	_, err = buf.Write(e.key)
	if err != nil {
		return nil, err
	}

	return nil, nil
}
