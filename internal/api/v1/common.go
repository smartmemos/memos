package v1

import (
	"encoding/base64"
	"fmt"
	"strconv"
	"strings"

	"github.com/pkg/errors"
	"google.golang.org/protobuf/proto"

	commonpb "github.com/smartmemos/memos/internal/proto/model/common"
)

const (
	WorkspaceSettingNamePrefix = "settings/"
	UserNamePrefix             = "users/"
	MemoNamePrefix             = "memos/"
	ResourceNamePrefix         = "resources/"
	InboxNamePrefix            = "inboxes/"
	IdentityProviderNamePrefix = "identityProviders/"
	ActivityNamePrefix         = "activities/"
)

const (
	// DefaultPageSize is the default page size for requests.
	DefaultPageSize = 10
)

// GetNameParentTokens returns the tokens from a resource name.
func GetNameParentTokens(name string, tokenPrefixes ...string) ([]string, error) {
	parts := strings.Split(name, "/")
	if len(parts) != 2*len(tokenPrefixes) {
		return nil, errors.Errorf("invalid request %q", name)
	}

	var tokens []string
	for i, tokenPrefix := range tokenPrefixes {
		if fmt.Sprintf("%s/", parts[2*i]) != tokenPrefix {
			return nil, errors.Errorf("invalid prefix %q in request %q", tokenPrefix, name)
		}
		if parts[2*i+1] == "" {
			return nil, errors.Errorf("invalid request %q with empty prefix %q", name, tokenPrefix)
		}
		tokens = append(tokens, parts[2*i+1])
	}
	return tokens, nil
}

func ExtractWorkspaceSettingKeyFromName(name string) (string, error) {
	tokens, err := GetNameParentTokens(name, WorkspaceSettingNamePrefix)
	if err != nil {
		return "", err
	}
	return tokens[0], nil
}

// ExtractUserIDFromName returns the uid from a resource name.
func ExtractUserIDFromName(name string) (int32, error) {
	tokens, err := GetNameParentTokens(name, UserNamePrefix)
	if err != nil {
		return 0, err
	}
	id, err := strconv.Atoi(tokens[0])
	if err != nil {
		return 0, errors.Errorf("invalid user ID %q", tokens[0])
	}
	return int32(id), nil
}

// ExtractMemoUIDFromName returns the memo UID from a resource name.
// e.g., "memos/uuid" -> "uuid".
func ExtractMemoUIDFromName(name string) (string, error) {
	tokens, err := GetNameParentTokens(name, MemoNamePrefix)
	if err != nil {
		return "", err
	}
	id := tokens[0]
	return id, nil
}

// ExtractResourceUIDFromName returns the resource UID from a resource name.
func ExtractResourceUIDFromName(name string) (string, error) {
	tokens, err := GetNameParentTokens(name, ResourceNamePrefix)
	if err != nil {
		return "", err
	}
	id := tokens[0]
	return id, nil
}

// ExtractInboxIDFromName returns the inbox ID from a resource name.
func ExtractInboxIDFromName(name string) (int32, error) {
	tokens, err := GetNameParentTokens(name, InboxNamePrefix)
	if err != nil {
		return 0, err
	}
	id, err := strconv.Atoi(tokens[0])
	if err != nil {
		return 0, errors.Errorf("invalid inbox ID %q", tokens[0])
	}
	return int32(id), nil
}

func ExtractIdentityProviderIDFromName(name string) (int32, error) {
	tokens, err := GetNameParentTokens(name, IdentityProviderNamePrefix)
	if err != nil {
		return 0, err
	}
	id, err := strconv.Atoi(tokens[0])
	if err != nil {
		return 0, errors.Errorf("invalid identity provider ID %q", tokens[0])
	}
	return int32(id), nil
}

func ExtractActivityIDFromName(name string) (int32, error) {
	tokens, err := GetNameParentTokens(name, ActivityNamePrefix)
	if err != nil {
		return 0, err
	}
	id, err := strconv.Atoi(tokens[0])
	if err != nil {
		return 0, errors.Errorf("invalid activity ID %q", tokens[0])
	}
	return int32(id), nil
}

func getPageToken(pageSize int, page int) (string, error) {
	return marshalPageToken(&commonpb.PageToken{
		Limit:  int32(pageSize),
		Offset: int32((page - 1) * pageSize),
	})
}

func marshalPageToken(pageToken *commonpb.PageToken) (string, error) {
	b, err := proto.Marshal(pageToken)
	if err != nil {
		return "", errors.Wrapf(err, "failed to marshal page token")
	}
	return base64.StdEncoding.EncodeToString(b), nil
}

func unmarshalPageToken(s string, pageToken *commonpb.PageToken) error {
	b, err := base64.StdEncoding.DecodeString(s)
	if err != nil {
		return errors.Wrapf(err, "failed to decode page token")
	}
	if err := proto.Unmarshal(b, pageToken); err != nil {
		return errors.Wrapf(err, "failed to unmarshal page token")
	}
	return nil
}
