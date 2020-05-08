package memory

type metadata struct {
	fileID         uint64
	valueOffset    uint64
	valueSize      uint32
	sequenceNumber int64
}

type Metadata interface {
}

func NewMetadata(fileID, valueOffset uint64, valueSize uint32, sequenceNumber int64) Metadata {
	return &metadata{
		fileID:         fileID,
		valueOffset:    valueOffset,
		valueSize:      valueSize,
		sequenceNumber: sequenceNumber,
	}
}
