package model

import (
	"time"

	"github.com/smartmemos/memos/internal/pkg/db"
)

type Attachment struct {
	db.Model

	// UID is the user defined unique identifier for the attachment.
	UID string

	// Standard fields
	CreatorID int32
	// The related memo ID.
	MemoID int64

	// Domain specific fields
	Filename    string
	Blob        []byte
	Type        string
	Size        int64
	StorageType AttachmentStorageType
	Reference   string
	Payload     AttachmentPayload
}

func (Attachment) TableName() string {
	return TableAttachment
}

type AttachmentPayload struct {
	S3Object *AttachmentPayloadS3Object
}

type AttachmentStorageType string

const (
	// Attachment is stored locally. AKA, local file system.
	AttachmentStorageTypeLocal AttachmentStorageType = "local"
	// Attachment is stored in S3.
	AttachmentStorageTypeS3 AttachmentStorageType = "s3"
	// Attachment is stored in an external storage. The reference is a URL.
	AttachmentStorageTypeExternal AttachmentStorageType = "external"
)

type AttachmentPayloadS3Object struct {
	S3Config *StorageS3Config
	Key      string
	// last_presigned_time is the last time the object was presigned.
	// This is used to determine if the presigned URL is still valid.
	LastPresignedTime time.Time
}

type StorageS3Config struct {
	AccessKeyId     string
	AccessKeySecret string
	Endpoint        string
	Region          string
	Bucket          string
	UsePathStyle    bool
}

type AttachmentFilter struct {
	db.Query

	ID db.F[int64]
}
