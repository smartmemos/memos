package model

import (
	"regexp"
)

const (
	TableUser        = "user"
	TableUserSetting = "user_setting"
)

// Role is the type of a role.
type Role string

const (
	// RoleHost is the HOST role.
	RoleHost Role = "HOST"
	// RoleAdmin is the ADMIN role.
	RoleAdmin Role = "ADMIN"
	// RoleUser is the USER role.
	RoleUser Role = "USER"
)

type RowStatus string

const (
	// Normal is the status for a normal row.
	Normal RowStatus = "NORMAL"
	// Archived is the status for an archived row.
	Archived RowStatus = "ARCHIVED"
)

var UsernameReg = regexp.MustCompile("^[a-zA-Z0-9]([a-zA-Z0-9-]{1,30}[a-zA-Z0-9])$")
