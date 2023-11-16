// Code generated by ent, DO NOT EDIT.

package ent

import (
	"GoTrain/BackEnd/ent/messages"
	"GoTrain/BackEnd/ent/predicate"
	"GoTrain/BackEnd/ent/users"
	"context"
	"errors"
	"fmt"
	"sync"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
)

const (
	// Operation types.
	OpCreate    = ent.OpCreate
	OpDelete    = ent.OpDelete
	OpDeleteOne = ent.OpDeleteOne
	OpUpdate    = ent.OpUpdate
	OpUpdateOne = ent.OpUpdateOne

	// Node types.
	TypeMessages = "Messages"
	TypeUsers    = "Users"
)

// MessagesMutation represents an operation that mutates the Messages nodes in the graph.
type MessagesMutation struct {
	config
	op            Op
	typ           string
	id            *int
	content       *string
	created_at    *time.Time
	_USER_ID      *int
	add_USER_ID   *int
	clearedFields map[string]struct{}
	user          *int
	cleareduser   bool
	done          bool
	oldValue      func(context.Context) (*Messages, error)
	predicates    []predicate.Messages
}

var _ ent.Mutation = (*MessagesMutation)(nil)

// messagesOption allows management of the mutation configuration using functional options.
type messagesOption func(*MessagesMutation)

// newMessagesMutation creates new mutation for the Messages entity.
func newMessagesMutation(c config, op Op, opts ...messagesOption) *MessagesMutation {
	m := &MessagesMutation{
		config:        c,
		op:            op,
		typ:           TypeMessages,
		clearedFields: make(map[string]struct{}),
	}
	for _, opt := range opts {
		opt(m)
	}
	return m
}

// withMessagesID sets the ID field of the mutation.
func withMessagesID(id int) messagesOption {
	return func(m *MessagesMutation) {
		var (
			err   error
			once  sync.Once
			value *Messages
		)
		m.oldValue = func(ctx context.Context) (*Messages, error) {
			once.Do(func() {
				if m.done {
					err = errors.New("querying old values post mutation is not allowed")
				} else {
					value, err = m.Client().Messages.Get(ctx, id)
				}
			})
			return value, err
		}
		m.id = &id
	}
}

// withMessages sets the old Messages of the mutation.
func withMessages(node *Messages) messagesOption {
	return func(m *MessagesMutation) {
		m.oldValue = func(context.Context) (*Messages, error) {
			return node, nil
		}
		m.id = &node.ID
	}
}

// Client returns a new `ent.Client` from the mutation. If the mutation was
// executed in a transaction (ent.Tx), a transactional client is returned.
func (m MessagesMutation) Client() *Client {
	client := &Client{config: m.config}
	client.init()
	return client
}

// Tx returns an `ent.Tx` for mutations that were executed in transactions;
// it returns an error otherwise.
func (m MessagesMutation) Tx() (*Tx, error) {
	if _, ok := m.driver.(*txDriver); !ok {
		return nil, errors.New("ent: mutation is not running in a transaction")
	}
	tx := &Tx{config: m.config}
	tx.init()
	return tx, nil
}

// ID returns the ID value in the mutation. Note that the ID is only available
// if it was provided to the builder or after it was returned from the database.
func (m *MessagesMutation) ID() (id int, exists bool) {
	if m.id == nil {
		return
	}
	return *m.id, true
}

// IDs queries the database and returns the entity ids that match the mutation's predicate.
// That means, if the mutation is applied within a transaction with an isolation level such
// as sql.LevelSerializable, the returned ids match the ids of the rows that will be updated
// or updated by the mutation.
func (m *MessagesMutation) IDs(ctx context.Context) ([]int, error) {
	switch {
	case m.op.Is(OpUpdateOne | OpDeleteOne):
		id, exists := m.ID()
		if exists {
			return []int{id}, nil
		}
		fallthrough
	case m.op.Is(OpUpdate | OpDelete):
		return m.Client().Messages.Query().Where(m.predicates...).IDs(ctx)
	default:
		return nil, fmt.Errorf("IDs is not allowed on %s operations", m.op)
	}
}

// SetContent sets the "content" field.
func (m *MessagesMutation) SetContent(s string) {
	m.content = &s
}

// Content returns the value of the "content" field in the mutation.
func (m *MessagesMutation) Content() (r string, exists bool) {
	v := m.content
	if v == nil {
		return
	}
	return *v, true
}

// OldContent returns the old "content" field's value of the Messages entity.
// If the Messages object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *MessagesMutation) OldContent(ctx context.Context) (v string, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldContent is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldContent requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldContent: %w", err)
	}
	return oldValue.Content, nil
}

// ResetContent resets all changes to the "content" field.
func (m *MessagesMutation) ResetContent() {
	m.content = nil
}

// SetCreatedAt sets the "created_at" field.
func (m *MessagesMutation) SetCreatedAt(t time.Time) {
	m.created_at = &t
}

// CreatedAt returns the value of the "created_at" field in the mutation.
func (m *MessagesMutation) CreatedAt() (r time.Time, exists bool) {
	v := m.created_at
	if v == nil {
		return
	}
	return *v, true
}

// OldCreatedAt returns the old "created_at" field's value of the Messages entity.
// If the Messages object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *MessagesMutation) OldCreatedAt(ctx context.Context) (v time.Time, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldCreatedAt is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldCreatedAt requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldCreatedAt: %w", err)
	}
	return oldValue.CreatedAt, nil
}

// ResetCreatedAt resets all changes to the "created_at" field.
func (m *MessagesMutation) ResetCreatedAt() {
	m.created_at = nil
}

// SetUSERID sets the "USER_ID" field.
func (m *MessagesMutation) SetUSERID(i int) {
	m._USER_ID = &i
	m.add_USER_ID = nil
}

// USERID returns the value of the "USER_ID" field in the mutation.
func (m *MessagesMutation) USERID() (r int, exists bool) {
	v := m._USER_ID
	if v == nil {
		return
	}
	return *v, true
}

// OldUSERID returns the old "USER_ID" field's value of the Messages entity.
// If the Messages object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *MessagesMutation) OldUSERID(ctx context.Context) (v int, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldUSERID is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldUSERID requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldUSERID: %w", err)
	}
	return oldValue.USERID, nil
}

// AddUSERID adds i to the "USER_ID" field.
func (m *MessagesMutation) AddUSERID(i int) {
	if m.add_USER_ID != nil {
		*m.add_USER_ID += i
	} else {
		m.add_USER_ID = &i
	}
}

// AddedUSERID returns the value that was added to the "USER_ID" field in this mutation.
func (m *MessagesMutation) AddedUSERID() (r int, exists bool) {
	v := m.add_USER_ID
	if v == nil {
		return
	}
	return *v, true
}

// ResetUSERID resets all changes to the "USER_ID" field.
func (m *MessagesMutation) ResetUSERID() {
	m._USER_ID = nil
	m.add_USER_ID = nil
}

// SetUserID sets the "user" edge to the Users entity by id.
func (m *MessagesMutation) SetUserID(id int) {
	m.user = &id
}

// ClearUser clears the "user" edge to the Users entity.
func (m *MessagesMutation) ClearUser() {
	m.cleareduser = true
}

// UserCleared reports if the "user" edge to the Users entity was cleared.
func (m *MessagesMutation) UserCleared() bool {
	return m.cleareduser
}

// UserID returns the "user" edge ID in the mutation.
func (m *MessagesMutation) UserID() (id int, exists bool) {
	if m.user != nil {
		return *m.user, true
	}
	return
}

// UserIDs returns the "user" edge IDs in the mutation.
// Note that IDs always returns len(IDs) <= 1 for unique edges, and you should use
// UserID instead. It exists only for internal usage by the builders.
func (m *MessagesMutation) UserIDs() (ids []int) {
	if id := m.user; id != nil {
		ids = append(ids, *id)
	}
	return
}

// ResetUser resets all changes to the "user" edge.
func (m *MessagesMutation) ResetUser() {
	m.user = nil
	m.cleareduser = false
}

// Where appends a list predicates to the MessagesMutation builder.
func (m *MessagesMutation) Where(ps ...predicate.Messages) {
	m.predicates = append(m.predicates, ps...)
}

// WhereP appends storage-level predicates to the MessagesMutation builder. Using this method,
// users can use type-assertion to append predicates that do not depend on any generated package.
func (m *MessagesMutation) WhereP(ps ...func(*sql.Selector)) {
	p := make([]predicate.Messages, len(ps))
	for i := range ps {
		p[i] = ps[i]
	}
	m.Where(p...)
}

// Op returns the operation name.
func (m *MessagesMutation) Op() Op {
	return m.op
}

// SetOp allows setting the mutation operation.
func (m *MessagesMutation) SetOp(op Op) {
	m.op = op
}

// Type returns the node type of this mutation (Messages).
func (m *MessagesMutation) Type() string {
	return m.typ
}

// Fields returns all fields that were changed during this mutation. Note that in
// order to get all numeric fields that were incremented/decremented, call
// AddedFields().
func (m *MessagesMutation) Fields() []string {
	fields := make([]string, 0, 3)
	if m.content != nil {
		fields = append(fields, messages.FieldContent)
	}
	if m.created_at != nil {
		fields = append(fields, messages.FieldCreatedAt)
	}
	if m._USER_ID != nil {
		fields = append(fields, messages.FieldUSERID)
	}
	return fields
}

// Field returns the value of a field with the given name. The second boolean
// return value indicates that this field was not set, or was not defined in the
// schema.
func (m *MessagesMutation) Field(name string) (ent.Value, bool) {
	switch name {
	case messages.FieldContent:
		return m.Content()
	case messages.FieldCreatedAt:
		return m.CreatedAt()
	case messages.FieldUSERID:
		return m.USERID()
	}
	return nil, false
}

// OldField returns the old value of the field from the database. An error is
// returned if the mutation operation is not UpdateOne, or the query to the
// database failed.
func (m *MessagesMutation) OldField(ctx context.Context, name string) (ent.Value, error) {
	switch name {
	case messages.FieldContent:
		return m.OldContent(ctx)
	case messages.FieldCreatedAt:
		return m.OldCreatedAt(ctx)
	case messages.FieldUSERID:
		return m.OldUSERID(ctx)
	}
	return nil, fmt.Errorf("unknown Messages field %s", name)
}

// SetField sets the value of a field with the given name. It returns an error if
// the field is not defined in the schema, or if the type mismatched the field
// type.
func (m *MessagesMutation) SetField(name string, value ent.Value) error {
	switch name {
	case messages.FieldContent:
		v, ok := value.(string)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetContent(v)
		return nil
	case messages.FieldCreatedAt:
		v, ok := value.(time.Time)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetCreatedAt(v)
		return nil
	case messages.FieldUSERID:
		v, ok := value.(int)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetUSERID(v)
		return nil
	}
	return fmt.Errorf("unknown Messages field %s", name)
}

// AddedFields returns all numeric fields that were incremented/decremented during
// this mutation.
func (m *MessagesMutation) AddedFields() []string {
	var fields []string
	if m.add_USER_ID != nil {
		fields = append(fields, messages.FieldUSERID)
	}
	return fields
}

// AddedField returns the numeric value that was incremented/decremented on a field
// with the given name. The second boolean return value indicates that this field
// was not set, or was not defined in the schema.
func (m *MessagesMutation) AddedField(name string) (ent.Value, bool) {
	switch name {
	case messages.FieldUSERID:
		return m.AddedUSERID()
	}
	return nil, false
}

// AddField adds the value to the field with the given name. It returns an error if
// the field is not defined in the schema, or if the type mismatched the field
// type.
func (m *MessagesMutation) AddField(name string, value ent.Value) error {
	switch name {
	case messages.FieldUSERID:
		v, ok := value.(int)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.AddUSERID(v)
		return nil
	}
	return fmt.Errorf("unknown Messages numeric field %s", name)
}

// ClearedFields returns all nullable fields that were cleared during this
// mutation.
func (m *MessagesMutation) ClearedFields() []string {
	return nil
}

// FieldCleared returns a boolean indicating if a field with the given name was
// cleared in this mutation.
func (m *MessagesMutation) FieldCleared(name string) bool {
	_, ok := m.clearedFields[name]
	return ok
}

// ClearField clears the value of the field with the given name. It returns an
// error if the field is not defined in the schema.
func (m *MessagesMutation) ClearField(name string) error {
	return fmt.Errorf("unknown Messages nullable field %s", name)
}

// ResetField resets all changes in the mutation for the field with the given name.
// It returns an error if the field is not defined in the schema.
func (m *MessagesMutation) ResetField(name string) error {
	switch name {
	case messages.FieldContent:
		m.ResetContent()
		return nil
	case messages.FieldCreatedAt:
		m.ResetCreatedAt()
		return nil
	case messages.FieldUSERID:
		m.ResetUSERID()
		return nil
	}
	return fmt.Errorf("unknown Messages field %s", name)
}

// AddedEdges returns all edge names that were set/added in this mutation.
func (m *MessagesMutation) AddedEdges() []string {
	edges := make([]string, 0, 1)
	if m.user != nil {
		edges = append(edges, messages.EdgeUser)
	}
	return edges
}

// AddedIDs returns all IDs (to other nodes) that were added for the given edge
// name in this mutation.
func (m *MessagesMutation) AddedIDs(name string) []ent.Value {
	switch name {
	case messages.EdgeUser:
		if id := m.user; id != nil {
			return []ent.Value{*id}
		}
	}
	return nil
}

// RemovedEdges returns all edge names that were removed in this mutation.
func (m *MessagesMutation) RemovedEdges() []string {
	edges := make([]string, 0, 1)
	return edges
}

// RemovedIDs returns all IDs (to other nodes) that were removed for the edge with
// the given name in this mutation.
func (m *MessagesMutation) RemovedIDs(name string) []ent.Value {
	return nil
}

// ClearedEdges returns all edge names that were cleared in this mutation.
func (m *MessagesMutation) ClearedEdges() []string {
	edges := make([]string, 0, 1)
	if m.cleareduser {
		edges = append(edges, messages.EdgeUser)
	}
	return edges
}

// EdgeCleared returns a boolean which indicates if the edge with the given name
// was cleared in this mutation.
func (m *MessagesMutation) EdgeCleared(name string) bool {
	switch name {
	case messages.EdgeUser:
		return m.cleareduser
	}
	return false
}

// ClearEdge clears the value of the edge with the given name. It returns an error
// if that edge is not defined in the schema.
func (m *MessagesMutation) ClearEdge(name string) error {
	switch name {
	case messages.EdgeUser:
		m.ClearUser()
		return nil
	}
	return fmt.Errorf("unknown Messages unique edge %s", name)
}

// ResetEdge resets all changes to the edge with the given name in this mutation.
// It returns an error if the edge is not defined in the schema.
func (m *MessagesMutation) ResetEdge(name string) error {
	switch name {
	case messages.EdgeUser:
		m.ResetUser()
		return nil
	}
	return fmt.Errorf("unknown Messages edge %s", name)
}

// UsersMutation represents an operation that mutates the Users nodes in the graph.
type UsersMutation struct {
	config
	op              Op
	typ             string
	id              *int
	name            *string
	clearedFields   map[string]struct{}
	messages        *int
	clearedmessages bool
	done            bool
	oldValue        func(context.Context) (*Users, error)
	predicates      []predicate.Users
}

var _ ent.Mutation = (*UsersMutation)(nil)

// usersOption allows management of the mutation configuration using functional options.
type usersOption func(*UsersMutation)

// newUsersMutation creates new mutation for the Users entity.
func newUsersMutation(c config, op Op, opts ...usersOption) *UsersMutation {
	m := &UsersMutation{
		config:        c,
		op:            op,
		typ:           TypeUsers,
		clearedFields: make(map[string]struct{}),
	}
	for _, opt := range opts {
		opt(m)
	}
	return m
}

// withUsersID sets the ID field of the mutation.
func withUsersID(id int) usersOption {
	return func(m *UsersMutation) {
		var (
			err   error
			once  sync.Once
			value *Users
		)
		m.oldValue = func(ctx context.Context) (*Users, error) {
			once.Do(func() {
				if m.done {
					err = errors.New("querying old values post mutation is not allowed")
				} else {
					value, err = m.Client().Users.Get(ctx, id)
				}
			})
			return value, err
		}
		m.id = &id
	}
}

// withUsers sets the old Users of the mutation.
func withUsers(node *Users) usersOption {
	return func(m *UsersMutation) {
		m.oldValue = func(context.Context) (*Users, error) {
			return node, nil
		}
		m.id = &node.ID
	}
}

// Client returns a new `ent.Client` from the mutation. If the mutation was
// executed in a transaction (ent.Tx), a transactional client is returned.
func (m UsersMutation) Client() *Client {
	client := &Client{config: m.config}
	client.init()
	return client
}

// Tx returns an `ent.Tx` for mutations that were executed in transactions;
// it returns an error otherwise.
func (m UsersMutation) Tx() (*Tx, error) {
	if _, ok := m.driver.(*txDriver); !ok {
		return nil, errors.New("ent: mutation is not running in a transaction")
	}
	tx := &Tx{config: m.config}
	tx.init()
	return tx, nil
}

// ID returns the ID value in the mutation. Note that the ID is only available
// if it was provided to the builder or after it was returned from the database.
func (m *UsersMutation) ID() (id int, exists bool) {
	if m.id == nil {
		return
	}
	return *m.id, true
}

// IDs queries the database and returns the entity ids that match the mutation's predicate.
// That means, if the mutation is applied within a transaction with an isolation level such
// as sql.LevelSerializable, the returned ids match the ids of the rows that will be updated
// or updated by the mutation.
func (m *UsersMutation) IDs(ctx context.Context) ([]int, error) {
	switch {
	case m.op.Is(OpUpdateOne | OpDeleteOne):
		id, exists := m.ID()
		if exists {
			return []int{id}, nil
		}
		fallthrough
	case m.op.Is(OpUpdate | OpDelete):
		return m.Client().Users.Query().Where(m.predicates...).IDs(ctx)
	default:
		return nil, fmt.Errorf("IDs is not allowed on %s operations", m.op)
	}
}

// SetName sets the "name" field.
func (m *UsersMutation) SetName(s string) {
	m.name = &s
}

// Name returns the value of the "name" field in the mutation.
func (m *UsersMutation) Name() (r string, exists bool) {
	v := m.name
	if v == nil {
		return
	}
	return *v, true
}

// OldName returns the old "name" field's value of the Users entity.
// If the Users object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *UsersMutation) OldName(ctx context.Context) (v string, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, errors.New("OldName is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, errors.New("OldName requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldName: %w", err)
	}
	return oldValue.Name, nil
}

// ResetName resets all changes to the "name" field.
func (m *UsersMutation) ResetName() {
	m.name = nil
}

// SetMessagesID sets the "messages" edge to the Messages entity by id.
func (m *UsersMutation) SetMessagesID(id int) {
	m.messages = &id
}

// ClearMessages clears the "messages" edge to the Messages entity.
func (m *UsersMutation) ClearMessages() {
	m.clearedmessages = true
}

// MessagesCleared reports if the "messages" edge to the Messages entity was cleared.
func (m *UsersMutation) MessagesCleared() bool {
	return m.clearedmessages
}

// MessagesID returns the "messages" edge ID in the mutation.
func (m *UsersMutation) MessagesID() (id int, exists bool) {
	if m.messages != nil {
		return *m.messages, true
	}
	return
}

// MessagesIDs returns the "messages" edge IDs in the mutation.
// Note that IDs always returns len(IDs) <= 1 for unique edges, and you should use
// MessagesID instead. It exists only for internal usage by the builders.
func (m *UsersMutation) MessagesIDs() (ids []int) {
	if id := m.messages; id != nil {
		ids = append(ids, *id)
	}
	return
}

// ResetMessages resets all changes to the "messages" edge.
func (m *UsersMutation) ResetMessages() {
	m.messages = nil
	m.clearedmessages = false
}

// Where appends a list predicates to the UsersMutation builder.
func (m *UsersMutation) Where(ps ...predicate.Users) {
	m.predicates = append(m.predicates, ps...)
}

// WhereP appends storage-level predicates to the UsersMutation builder. Using this method,
// users can use type-assertion to append predicates that do not depend on any generated package.
func (m *UsersMutation) WhereP(ps ...func(*sql.Selector)) {
	p := make([]predicate.Users, len(ps))
	for i := range ps {
		p[i] = ps[i]
	}
	m.Where(p...)
}

// Op returns the operation name.
func (m *UsersMutation) Op() Op {
	return m.op
}

// SetOp allows setting the mutation operation.
func (m *UsersMutation) SetOp(op Op) {
	m.op = op
}

// Type returns the node type of this mutation (Users).
func (m *UsersMutation) Type() string {
	return m.typ
}

// Fields returns all fields that were changed during this mutation. Note that in
// order to get all numeric fields that were incremented/decremented, call
// AddedFields().
func (m *UsersMutation) Fields() []string {
	fields := make([]string, 0, 1)
	if m.name != nil {
		fields = append(fields, users.FieldName)
	}
	return fields
}

// Field returns the value of a field with the given name. The second boolean
// return value indicates that this field was not set, or was not defined in the
// schema.
func (m *UsersMutation) Field(name string) (ent.Value, bool) {
	switch name {
	case users.FieldName:
		return m.Name()
	}
	return nil, false
}

// OldField returns the old value of the field from the database. An error is
// returned if the mutation operation is not UpdateOne, or the query to the
// database failed.
func (m *UsersMutation) OldField(ctx context.Context, name string) (ent.Value, error) {
	switch name {
	case users.FieldName:
		return m.OldName(ctx)
	}
	return nil, fmt.Errorf("unknown Users field %s", name)
}

// SetField sets the value of a field with the given name. It returns an error if
// the field is not defined in the schema, or if the type mismatched the field
// type.
func (m *UsersMutation) SetField(name string, value ent.Value) error {
	switch name {
	case users.FieldName:
		v, ok := value.(string)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetName(v)
		return nil
	}
	return fmt.Errorf("unknown Users field %s", name)
}

// AddedFields returns all numeric fields that were incremented/decremented during
// this mutation.
func (m *UsersMutation) AddedFields() []string {
	return nil
}

// AddedField returns the numeric value that was incremented/decremented on a field
// with the given name. The second boolean return value indicates that this field
// was not set, or was not defined in the schema.
func (m *UsersMutation) AddedField(name string) (ent.Value, bool) {
	return nil, false
}

// AddField adds the value to the field with the given name. It returns an error if
// the field is not defined in the schema, or if the type mismatched the field
// type.
func (m *UsersMutation) AddField(name string, value ent.Value) error {
	switch name {
	}
	return fmt.Errorf("unknown Users numeric field %s", name)
}

// ClearedFields returns all nullable fields that were cleared during this
// mutation.
func (m *UsersMutation) ClearedFields() []string {
	return nil
}

// FieldCleared returns a boolean indicating if a field with the given name was
// cleared in this mutation.
func (m *UsersMutation) FieldCleared(name string) bool {
	_, ok := m.clearedFields[name]
	return ok
}

// ClearField clears the value of the field with the given name. It returns an
// error if the field is not defined in the schema.
func (m *UsersMutation) ClearField(name string) error {
	return fmt.Errorf("unknown Users nullable field %s", name)
}

// ResetField resets all changes in the mutation for the field with the given name.
// It returns an error if the field is not defined in the schema.
func (m *UsersMutation) ResetField(name string) error {
	switch name {
	case users.FieldName:
		m.ResetName()
		return nil
	}
	return fmt.Errorf("unknown Users field %s", name)
}

// AddedEdges returns all edge names that were set/added in this mutation.
func (m *UsersMutation) AddedEdges() []string {
	edges := make([]string, 0, 1)
	if m.messages != nil {
		edges = append(edges, users.EdgeMessages)
	}
	return edges
}

// AddedIDs returns all IDs (to other nodes) that were added for the given edge
// name in this mutation.
func (m *UsersMutation) AddedIDs(name string) []ent.Value {
	switch name {
	case users.EdgeMessages:
		if id := m.messages; id != nil {
			return []ent.Value{*id}
		}
	}
	return nil
}

// RemovedEdges returns all edge names that were removed in this mutation.
func (m *UsersMutation) RemovedEdges() []string {
	edges := make([]string, 0, 1)
	return edges
}

// RemovedIDs returns all IDs (to other nodes) that were removed for the edge with
// the given name in this mutation.
func (m *UsersMutation) RemovedIDs(name string) []ent.Value {
	return nil
}

// ClearedEdges returns all edge names that were cleared in this mutation.
func (m *UsersMutation) ClearedEdges() []string {
	edges := make([]string, 0, 1)
	if m.clearedmessages {
		edges = append(edges, users.EdgeMessages)
	}
	return edges
}

// EdgeCleared returns a boolean which indicates if the edge with the given name
// was cleared in this mutation.
func (m *UsersMutation) EdgeCleared(name string) bool {
	switch name {
	case users.EdgeMessages:
		return m.clearedmessages
	}
	return false
}

// ClearEdge clears the value of the edge with the given name. It returns an error
// if that edge is not defined in the schema.
func (m *UsersMutation) ClearEdge(name string) error {
	switch name {
	case users.EdgeMessages:
		m.ClearMessages()
		return nil
	}
	return fmt.Errorf("unknown Users unique edge %s", name)
}

// ResetEdge resets all changes to the edge with the given name in this mutation.
// It returns an error if the edge is not defined in the schema.
func (m *UsersMutation) ResetEdge(name string) error {
	switch name {
	case users.EdgeMessages:
		m.ResetMessages()
		return nil
	}
	return fmt.Errorf("unknown Users edge %s", name)
}
