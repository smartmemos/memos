package model

const (
	TableMemo     = "memo"
	TableResource = "resource"
	TableReaction = "reaction"
)

// Visibility is the type of a visibility.
type Visibility string

const (
	// Public is the PUBLIC visibility.
	Public Visibility = "PUBLIC"
	// Protected is the PROTECTED visibility.
	Protected Visibility = "PROTECTED"
	// Private is the PRIVATE visibility.
	Private Visibility = "PRIVATE"
)

func (v Visibility) String() string {
	switch v {
	case Public:
		return "PUBLIC"
	case Protected:
		return "PROTECTED"
	case Private:
		return "PRIVATE"
	}
	return "PRIVATE"
}

type RelationType string

const (
	RelationReference RelationType = "REFERENCE"
	RelationComment   RelationType = "COMMENT"
)

type ResStorageType string

const (
	ResStorageTypeLocal         ResStorageType = "LOCAL"
	ResourceStorageTypeS3       ResStorageType = "S3"
	ResourceStorageTypeExternal ResStorageType = "EXTERNAL"
)
