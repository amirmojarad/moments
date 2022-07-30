// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"moments/ent/publicchat"
	"strings"

	"entgo.io/ent/dialect/sql"
)

// PublicChat is the model entity for the PublicChat schema.
type PublicChat struct {
	config
	// ID of the ent.
	ID int `json:"id,omitempty"`
}

// scanValues returns the types for scanning values from sql.Rows.
func (*PublicChat) scanValues(columns []string) ([]interface{}, error) {
	values := make([]interface{}, len(columns))
	for i := range columns {
		switch columns[i] {
		case publicchat.FieldID:
			values[i] = new(sql.NullInt64)
		default:
			return nil, fmt.Errorf("unexpected column %q for type PublicChat", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the PublicChat fields.
func (pc *PublicChat) assignValues(columns []string, values []interface{}) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case publicchat.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			pc.ID = int(value.Int64)
		}
	}
	return nil
}

// Update returns a builder for updating this PublicChat.
// Note that you need to call PublicChat.Unwrap() before calling this method if this PublicChat
// was returned from a transaction, and the transaction was committed or rolled back.
func (pc *PublicChat) Update() *PublicChatUpdateOne {
	return (&PublicChatClient{config: pc.config}).UpdateOne(pc)
}

// Unwrap unwraps the PublicChat entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (pc *PublicChat) Unwrap() *PublicChat {
	_tx, ok := pc.config.driver.(*txDriver)
	if !ok {
		panic("ent: PublicChat is not a transactional entity")
	}
	pc.config.driver = _tx.drv
	return pc
}

// String implements the fmt.Stringer.
func (pc *PublicChat) String() string {
	var builder strings.Builder
	builder.WriteString("PublicChat(")
	builder.WriteString(fmt.Sprintf("id=%v", pc.ID))
	builder.WriteByte(')')
	return builder.String()
}

// PublicChats is a parsable slice of PublicChat.
type PublicChats []*PublicChat

func (pc PublicChats) config(cfg config) {
	for _i := range pc {
		pc[_i].config = cfg
	}
}