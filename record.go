package parhelion

import (
	"bytes"
	"unsafe"
)

const (
	// crc              - 4 bytes.
	// version          - 1 byte.
	// key size         - 1 bytes.
	// value size       - 4 bytes.
	// sequence number  - 8 bytes.

	RecordChecksumSize       = 4
	RecordVersionSize        = 1
	RecordKeySizeSize        = 1
	RecordValueSizeSize      = 4
	RecordSequenceNumberSize = 8

	RecordHeaderSize = RecordChecksumSize +
		RecordVersionSize +
		RecordKeySizeSize +
		RecordValueSizeSize +
		RecordSequenceNumberSize
)

type recordHeader struct {
	checksum       uint32
	version        uint8
	keySize        uint8
	valueSize      uint32
	sequenceNumber int64
}

type RecordHeader interface {
	Serialize() (*bytes.Buffer, error)
	KeySize() uint8
	ValueSize() uint32
	setChecksum(checksum uint32)
}

func NewRecordHeader(checkSum uint32, version, keySize uint8, valueSize uint32, sequenceNumber int64) RecordHeader {
	return &recordHeader{
		checksum:       checkSum,
		version:        version,
		keySize:        keySize,
		valueSize:      valueSize,
		sequenceNumber: sequenceNumber,
	}
}

func NewRecordHeaderFromBuffer(buf bytes.Buffer) (RecordHeader, error) {
	rh := new(recordHeader)
	reader := func(size uint8) ([]byte, error) {
		bs := make([]byte, size)
		_, err := buf.Read(bs)
		if err != nil {
			return nil, err
		}
		return bs, nil
	}

	checksumBytes, err := reader(RecordChecksumSize)
	if err != nil {
		return nil, err
	}
	rh.checksum = *(*uint32)(unsafe.Pointer(&checksumBytes))

	versionBytes, err := reader(RecordVersionSize)
	if err != nil {
		return nil, err
	}
	rh.version = *(*uint8)(unsafe.Pointer(&versionBytes))

	keySizeBytes, err := reader(RecordKeySizeSize)
	if err != nil {
		return nil, err
	}
	rh.keySize = *(*uint8)(unsafe.Pointer(&keySizeBytes))

	valueSizeBytes, err := reader(RecordValueSizeSize)
	if err != nil {
		return nil, err
	}
	rh.valueSize = *(*uint32)(unsafe.Pointer(&valueSizeBytes))

	sequenceNumberBytes, err := reader(RecordSequenceNumberSize)
	if err != nil {
		return nil, err
	}
	rh.sequenceNumber = *(*int64)(unsafe.Pointer(&sequenceNumberBytes))

	return rh, nil
}

func (rh *recordHeader) Serialize() (*bytes.Buffer, error) {
	buf := bytes.NewBuffer(make([]byte, RecordHeaderSize))

	_, err := buf.Write(*(*[]byte)(unsafe.Pointer(&rh.checksum)))
	if err != nil {
		return nil, err
	}
	_, err = buf.Write(*(*[]byte)(unsafe.Pointer(&rh.version)))
	if err != nil {
		return nil, err
	}
	_, err = buf.Write(*(*[]byte)(unsafe.Pointer(&rh.keySize)))
	if err != nil {
		return nil, err
	}
	_, err = buf.Write(*(*[]byte)(unsafe.Pointer(&rh.valueSize)))
	if err != nil {
		return nil, err
	}
	_, err = buf.Write(*(*[]byte)(unsafe.Pointer(&rh.sequenceNumber)))
	if err != nil {
		return nil, err
	}

	return buf, nil
}

func (rh *recordHeader) KeySize() uint8 {
	return rh.keySize
}

func (rh *recordHeader) ValueSize() uint32 {
	return rh.valueSize
}

func (rh *recordHeader) setChecksum(checksum uint32) {
	rh.checksum = checksum
}

type record struct {
	key   []byte
	value []byte
	// metadata InMemoryIndexMetadata
	header RecordHeader
}

type Record interface {
	Serialize() (*bytes.Buffer, error)
}

func NewRecord(key, value []byte) Record {
	return &record{
		key:    key,
		value:  value,
		header: NewRecordHeader(0, CurrentDataFileVersion, uint8(len(key)), uint32(len(value)), -1),
	}
}

func NewRecordFromBuffer(buf bytes.Buffer, keySize uint8, valueSize uint32) (Record, error) {
	key := make([]byte, keySize)
	value := make([]byte, valueSize)
	_, err := buf.Read(key)
	if err != nil {
		return nil, err
	}
	_, err = buf.Read(value)
	if err != nil {
		return nil, err
	}
	return NewRecord(key, value), nil
}

func (r *record) Serialize() (*bytes.Buffer, error) {
	buf := bytes.NewBuffer(make([]byte, RecordHeaderSize+len(r.key)+len(r.value)))

	cs, err := r.computeChecksum()
	if err != nil {
		return nil, err
	}
	r.header.setChecksum(cs)
	headerBuf, err := r.header.Serialize()
	if err != nil {
		return nil, err
	}
	_, err = buf.Write(headerBuf.Bytes())
	if err != nil {
		return nil, err
	}

	_, err = buf.Write(r.key)
	if err != nil {
		return nil, err
	}

	_, err = buf.Write(r.value)
	if err != nil {
		return nil, err
	}

	return buf, nil
}

func (r *record) computeChecksum() (uint32, error) {
	bs := make([]byte, 0, RecordHeaderSize-RecordChecksumSize+len(r.key)+len(r.value))
	headerBuf, err := r.header.Serialize()
	if err != nil {
		return 0, err
	}

	bs = append(bs, headerBuf.Bytes()[RecordChecksumSize:]...)
	bs = append(bs, r.key...)
	bs = append(bs, r.value...)

	return computeChecksum(bs), nil
}
