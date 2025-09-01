package utils

import (
	"encoding/base64"

	"github.com/pkg/errors"
	"google.golang.org/protobuf/proto"

	modelpb "github.com/smartmemos/memos/internal/proto/model"
)

func GetPageToken(page, pageSize int64) (string, error) {
	return MarshalPageToken(&modelpb.PageToken{
		Limit:  int32(pageSize),
		Offset: int32(page-1) * int32(pageSize),
	})
}

func MarshalPageToken(pageToken *modelpb.PageToken) (string, error) {
	b, err := proto.Marshal(pageToken)
	if err != nil {
		return "", errors.Wrapf(err, "failed to marshal page token")
	}
	return base64.StdEncoding.EncodeToString(b), nil
}

func UnmarshalPageToken(s string, pageToken *modelpb.PageToken) error {
	b, err := base64.StdEncoding.DecodeString(s)
	if err != nil {
		return errors.Wrapf(err, "failed to decode page token")
	}
	if err := proto.Unmarshal(b, pageToken); err != nil {
		return errors.Wrapf(err, "failed to unmarshal page token")
	}
	return nil
}
