// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"moments/ent/user"
	"strings"
	"time"

	"entgo.io/ent/dialect/sql"
)

// User is the model entity for the User schema.
type User struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// CreatedDate holds the value of the "created_date" field.
	CreatedDate time.Time `json:"created_date,omitempty"`
	// UpdatedDate holds the value of the "updated_date" field.
	UpdatedDate time.Time `json:"updated_date,omitempty"`
	// DeletedDate holds the value of the "deleted_date" field.
	DeletedDate *time.Time `json:"deleted_date,omitempty"`
	// Username holds the value of the "username" field.
	Username string `json:"username,omitempty"`
	// Email holds the value of the "email" field.
	Email string `json:"email,omitempty"`
	// Password holds the value of the "password" field.
	Password string `json:"password,omitempty"`
	// FullName holds the value of the "full_name" field.
	FullName *string `json:"full_name,omitempty"`
	// IsAdmin holds the value of the "is_admin" field.
	IsAdmin bool `json:"is_admin,omitempty"`
	// IsStaff holds the value of the "is_staff" field.
	IsStaff bool `json:"is_staff,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the UserQuery when eager-loading is set.
	Edges UserEdges `json:"edges"`
}

// UserEdges holds the relations/edges for other nodes in the graph.
type UserEdges struct {
	// Posts holds the value of the posts edge.
	Posts []*Post `json:"posts,omitempty"`
	// Followers holds the value of the followers edge.
	Followers []*User `json:"followers,omitempty"`
	// Following holds the value of the following edge.
	Following []*User `json:"following,omitempty"`
	// Rooms holds the value of the rooms edge.
	Rooms []*Room `json:"rooms,omitempty"`
	// Messages holds the value of the messages edge.
	Messages []*Message `json:"messages,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [5]bool
}

// PostsOrErr returns the Posts value or an error if the edge
// was not loaded in eager-loading.
func (e UserEdges) PostsOrErr() ([]*Post, error) {
	if e.loadedTypes[0] {
		return e.Posts, nil
	}
	return nil, &NotLoadedError{edge: "posts"}
}

// FollowersOrErr returns the Followers value or an error if the edge
// was not loaded in eager-loading.
func (e UserEdges) FollowersOrErr() ([]*User, error) {
	if e.loadedTypes[1] {
		return e.Followers, nil
	}
	return nil, &NotLoadedError{edge: "followers"}
}

// FollowingOrErr returns the Following value or an error if the edge
// was not loaded in eager-loading.
func (e UserEdges) FollowingOrErr() ([]*User, error) {
	if e.loadedTypes[2] {
		return e.Following, nil
	}
	return nil, &NotLoadedError{edge: "following"}
}

// RoomsOrErr returns the Rooms value or an error if the edge
// was not loaded in eager-loading.
func (e UserEdges) RoomsOrErr() ([]*Room, error) {
	if e.loadedTypes[3] {
		return e.Rooms, nil
	}
	return nil, &NotLoadedError{edge: "rooms"}
}

// MessagesOrErr returns the Messages value or an error if the edge
// was not loaded in eager-loading.
func (e UserEdges) MessagesOrErr() ([]*Message, error) {
	if e.loadedTypes[4] {
		return e.Messages, nil
	}
	return nil, &NotLoadedError{edge: "messages"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*User) scanValues(columns []string) ([]interface{}, error) {
	values := make([]interface{}, len(columns))
	for i := range columns {
		switch columns[i] {
		case user.FieldIsAdmin, user.FieldIsStaff:
			values[i] = new(sql.NullBool)
		case user.FieldID:
			values[i] = new(sql.NullInt64)
		case user.FieldUsername, user.FieldEmail, user.FieldPassword, user.FieldFullName:
			values[i] = new(sql.NullString)
		case user.FieldCreatedDate, user.FieldUpdatedDate, user.FieldDeletedDate:
			values[i] = new(sql.NullTime)
		default:
			return nil, fmt.Errorf("unexpected column %q for type User", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the User fields.
func (u *User) assignValues(columns []string, values []interface{}) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case user.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			u.ID = int(value.Int64)
		case user.FieldCreatedDate:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field created_date", values[i])
			} else if value.Valid {
				u.CreatedDate = value.Time
			}
		case user.FieldUpdatedDate:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field updated_date", values[i])
			} else if value.Valid {
				u.UpdatedDate = value.Time
			}
		case user.FieldDeletedDate:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field deleted_date", values[i])
			} else if value.Valid {
				u.DeletedDate = new(time.Time)
				*u.DeletedDate = value.Time
			}
		case user.FieldUsername:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field username", values[i])
			} else if value.Valid {
				u.Username = value.String
			}
		case user.FieldEmail:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field email", values[i])
			} else if value.Valid {
				u.Email = value.String
			}
		case user.FieldPassword:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field password", values[i])
			} else if value.Valid {
				u.Password = value.String
			}
		case user.FieldFullName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field full_name", values[i])
			} else if value.Valid {
				u.FullName = new(string)
				*u.FullName = value.String
			}
		case user.FieldIsAdmin:
			if value, ok := values[i].(*sql.NullBool); !ok {
				return fmt.Errorf("unexpected type %T for field is_admin", values[i])
			} else if value.Valid {
				u.IsAdmin = value.Bool
			}
		case user.FieldIsStaff:
			if value, ok := values[i].(*sql.NullBool); !ok {
				return fmt.Errorf("unexpected type %T for field is_staff", values[i])
			} else if value.Valid {
				u.IsStaff = value.Bool
			}
		}
	}
	return nil
}

// QueryPosts queries the "posts" edge of the User entity.
func (u *User) QueryPosts() *PostQuery {
	return (&UserClient{config: u.config}).QueryPosts(u)
}

// QueryFollowers queries the "followers" edge of the User entity.
func (u *User) QueryFollowers() *UserQuery {
	return (&UserClient{config: u.config}).QueryFollowers(u)
}

// QueryFollowing queries the "following" edge of the User entity.
func (u *User) QueryFollowing() *UserQuery {
	return (&UserClient{config: u.config}).QueryFollowing(u)
}

// QueryRooms queries the "rooms" edge of the User entity.
func (u *User) QueryRooms() *RoomQuery {
	return (&UserClient{config: u.config}).QueryRooms(u)
}

// QueryMessages queries the "messages" edge of the User entity.
func (u *User) QueryMessages() *MessageQuery {
	return (&UserClient{config: u.config}).QueryMessages(u)
}

// Update returns a builder for updating this User.
// Note that you need to call User.Unwrap() before calling this method if this User
// was returned from a transaction, and the transaction was committed or rolled back.
func (u *User) Update() *UserUpdateOne {
	return (&UserClient{config: u.config}).UpdateOne(u)
}

// Unwrap unwraps the User entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (u *User) Unwrap() *User {
	_tx, ok := u.config.driver.(*txDriver)
	if !ok {
		panic("ent: User is not a transactional entity")
	}
	u.config.driver = _tx.drv
	return u
}

// String implements the fmt.Stringer.
func (u *User) String() string {
	var builder strings.Builder
	builder.WriteString("User(")
	builder.WriteString(fmt.Sprintf("id=%v, ", u.ID))
	builder.WriteString("created_date=")
	builder.WriteString(u.CreatedDate.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("updated_date=")
	builder.WriteString(u.UpdatedDate.Format(time.ANSIC))
	builder.WriteString(", ")
	if v := u.DeletedDate; v != nil {
		builder.WriteString("deleted_date=")
		builder.WriteString(v.Format(time.ANSIC))
	}
	builder.WriteString(", ")
	builder.WriteString("username=")
	builder.WriteString(u.Username)
	builder.WriteString(", ")
	builder.WriteString("email=")
	builder.WriteString(u.Email)
	builder.WriteString(", ")
	builder.WriteString("password=")
	builder.WriteString(u.Password)
	builder.WriteString(", ")
	if v := u.FullName; v != nil {
		builder.WriteString("full_name=")
		builder.WriteString(*v)
	}
	builder.WriteString(", ")
	builder.WriteString("is_admin=")
	builder.WriteString(fmt.Sprintf("%v", u.IsAdmin))
	builder.WriteString(", ")
	builder.WriteString("is_staff=")
	builder.WriteString(fmt.Sprintf("%v", u.IsStaff))
	builder.WriteByte(')')
	return builder.String()
}

// Users is a parsable slice of User.
type Users []*User

func (u Users) config(cfg config) {
	for _i := range u {
		u[_i].config = cfg
	}
}
