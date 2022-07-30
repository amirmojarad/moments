// Code generated by ent, DO NOT EDIT.

package post

import (
	"time"
)

const (
	// Label holds the string label denoting the post type in the database.
	Label = "post"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldCreatedDate holds the string denoting the created_date field in the database.
	FieldCreatedDate = "created_date"
	// FieldUpdatedDate holds the string denoting the updated_date field in the database.
	FieldUpdatedDate = "updated_date"
	// FieldDeletedDate holds the string denoting the deleted_date field in the database.
	FieldDeletedDate = "deleted_date"
	// FieldPlainText holds the string denoting the plain_text field in the database.
	FieldPlainText = "plain_text"
	// FieldLikes holds the string denoting the likes field in the database.
	FieldLikes = "likes"
	// FieldLinkURL holds the string denoting the link_url field in the database.
	FieldLinkURL = "link_url"
	// EdgeOwner holds the string denoting the owner edge name in mutations.
	EdgeOwner = "owner"
	// Table holds the table name of the post in the database.
	Table = "posts"
	// OwnerTable is the table that holds the owner relation/edge.
	OwnerTable = "posts"
	// OwnerInverseTable is the table name for the User entity.
	// It exists in this package in order to avoid circular dependency with the "user" package.
	OwnerInverseTable = "users"
	// OwnerColumn is the table column denoting the owner relation/edge.
	OwnerColumn = "user_posts"
)

// Columns holds all SQL columns for post fields.
var Columns = []string{
	FieldID,
	FieldCreatedDate,
	FieldUpdatedDate,
	FieldDeletedDate,
	FieldPlainText,
	FieldLikes,
	FieldLinkURL,
}

// ForeignKeys holds the SQL foreign-keys that are owned by the "posts"
// table and are not defined as standalone fields in the schema.
var ForeignKeys = []string{
	"user_posts",
}

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	for i := range ForeignKeys {
		if column == ForeignKeys[i] {
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
	// DefaultLikes holds the default value on creation for the "likes" field.
	DefaultLikes uint64
)
