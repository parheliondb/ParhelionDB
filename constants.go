package parhelion

import "regexp"

var (
	DataFileName          = ".data"
	CompactedDataFileName = DataFileName + "c"
	IndexFileName         = ".index"
	TombstoneFileName     = ".tombstone"
	MetadataFileName      = "META"

	DataFilePattern      = regexp.MustCompile("([0-9]+)" + DataFileName + "c?")
	IndexFilePattern     = regexp.MustCompile("([0-9]+)" + IndexFileName)
	TombstoneFilePattern = regexp.MustCompile("([0-9]+)" + TombstoneFileName)
	StorageFilePattern   = regexp.MustCompile("([0-9]+)\\.[a-z]+")
)
