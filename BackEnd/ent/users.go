// Code generated by ent, DO NOT EDIT.

package ent

import (
	"GoTrain/BackEnd/ent/messages"
	"GoTrain/BackEnd/ent/users"
	"fmt"
	"strings"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
)

// Users is the model entity for the Users schema.
type Users struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// Name holds the value of the "name" field.
	Name string `json:"name,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the UsersQuery when eager-loading is set.
	Edges         UsersEdges `json:"edges"`
	messages_user *int
	selectValues  sql.SelectValues
}

// UsersEdges holds the relations/edges for other nodes in the graph.
type UsersEdges struct {
	// Messages holds the value of the messages edge.
	Messages *Messages `json:"messages,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [1]bool
}

// MessagesOrErr returns the Messages value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e UsersEdges) MessagesOrErr() (*Messages, error) {
	if e.loadedTypes[0] {
		if e.Messages == nil {
			// Edge was loaded but was not found.
			return nil, &NotFoundError{label: messages.Label}
		}
		return e.Messages, nil
	}
	return nil, &NotLoadedError{edge: "messages"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Users) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case users.FieldID:
			values[i] = new(sql.NullInt64)
		case users.FieldName:
			values[i] = new(sql.NullString)
		case users.ForeignKeys[0]: // messages_user
			values[i] = new(sql.NullInt64)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Users fields.
func (u *Users) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case users.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			u.ID = int(value.Int64)
		case users.FieldName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field name", values[i])
			} else if value.Valid {
				u.Name = value.String
			}
		case users.ForeignKeys[0]:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for edge-field messages_user", value)
			} else if value.Valid {
				u.messages_user = new(int)
				*u.messages_user = int(value.Int64)
			}
		default:
			u.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the Users.
// This includes values selected through modifiers, order, etc.
func (u *Users) Value(name string) (ent.Value, error) {
	return u.selectValues.Get(name)
}

// QueryMessages queries the "messages" edge of the Users entity.
func (u *Users) QueryMessages() *MessagesQuery {
	return NewUsersClient(u.config).QueryMessages(u)
}

// Update returns a builder for updating this Users.
// Note that you need to call Users.Unwrap() before calling this method if this Users
// was returned from a transaction, and the transaction was committed or rolled back.
func (u *Users) Update() *UsersUpdateOne {
	return NewUsersClient(u.config).UpdateOne(u)
}

// Unwrap unwraps the Users entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (u *Users) Unwrap() *Users {
	_tx, ok := u.config.driver.(*txDriver)
	if !ok {
		panic("ent: Users is not a transactional entity")
	}
	u.config.driver = _tx.drv
	return u
}

// String implements the fmt.Stringer.
func (u *Users) String() string {
	var builder strings.Builder
	builder.WriteString("Users(")
	builder.WriteString(fmt.Sprintf("id=%v, ", u.ID))
	builder.WriteString("name=")
	builder.WriteString(u.Name)
	builder.WriteByte(')')
	return builder.String()
}

// UsersSlice is a parsable slice of Users.
type UsersSlice []*Users
