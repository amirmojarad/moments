// Code generated by ent, DO NOT EDIT.

package room

import (
	"fmt"
	"time"
)

const (
	// Label holds the string label denoting the room type in the database.
	Label = "room"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldCreatedDate holds the string denoting the created_date field in the database.
	FieldCreatedDate = "created_date"
	// FieldUpdatedDate holds the string denoting the updated_date field in the database.
	FieldUpdatedDate = "updated_date"
	// FieldDeletedDate holds the string denoting the deleted_date field in the database.
	FieldDeletedDate = "deleted_date"
	// FieldTitle holds the string denoting the title field in the database.
	FieldTitle = "title"
	// FieldType holds the string denoting the type field in the database.
	FieldType = "type"
	// EdgeUsers holds the string denoting the users edge name in mutations.
	EdgeUsers = "users"
	// EdgeMessages holds the string denoting the messages edge name in mutations.
	EdgeMessages = "messages"
	// Table holds the table name of the room in the database.
	Table = "rooms"
	// UsersTable is the table that holds the users relation/edge. The primary key declared below.
	UsersTable = "user_rooms"
	// UsersInverseTable is the table name for the User entity.
	// It exists in this package in order to avoid circular dependency with the "user" package.
	UsersInverseTable = "users"
	// MessagesTable is the table that holds the messages relation/edge.
	MessagesTable = "messages"
	// MessagesInverseTable is the table name for the Message entity.
	// It exists in this package in order to avoid circular dependency with the "message" package.
	MessagesInverseTable = "messages"
	// MessagesColumn is the table column denoting the messages relation/edge.
	MessagesColumn = "room_messages"
)

// Columns holds all SQL columns for room fields.
var Columns = []string{
	FieldID,
	FieldCreatedDate,
	FieldUpdatedDate,
	FieldDeletedDate,
	FieldTitle,
	FieldType,
}

var (
	// UsersPrimaryKey and UsersColumn2 are the table columns denoting the
	// primary key for the users relation (M2M).
	UsersPrimaryKey = []string{"user_id", "room_id"}
)

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	return false
}

var (
	// DefaultCreatedDate holds the default value on creation for the "created_date" field.
	DefaultCreatedDate time.Time
	// DefaultUpdatedDate holds the default value on creation for the "updated_date" field.
	DefaultUpdatedDate time.Time
	// TitleValidator is a validator for the "title" field. It is called by the builders before save.
	TitleValidator func(string) error
)

// Type defines the type for the "type" enum field.
type Type string

// TypeNone is the default value of the Type enum.
const DefaultType = TypeNone

// Type values.
const (
	TypeGroup   Type = "group"
	TypePrivate Type = "private"
	TypeNone    Type = "none"
)

func (_type Type) String() string {
	return string(_type)
}

// TypeValidator is a validator for the "type" field enum values. It is called by the builders before save.
func TypeValidator(_type Type) error {
	switch _type {
	case TypeGroup, TypePrivate, TypeNone:
		return nil
	default:
		return fmt.Errorf("room: invalid enum value for type field: %q", _type)
	}
}
