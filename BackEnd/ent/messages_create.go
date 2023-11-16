// Code generated by ent, DO NOT EDIT.

package ent

import (
	"GoTrain/BackEnd/ent/messages"
	"GoTrain/BackEnd/ent/users"
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// MessagesCreate is the builder for creating a Messages entity.
type MessagesCreate struct {
	config
	mutation *MessagesMutation
	hooks    []Hook
}

// SetContent sets the "content" field.
func (mc *MessagesCreate) SetContent(s string) *MessagesCreate {
	mc.mutation.SetContent(s)
	return mc
}

// SetCreatedAt sets the "created_at" field.
func (mc *MessagesCreate) SetCreatedAt(t time.Time) *MessagesCreate {
	mc.mutation.SetCreatedAt(t)
	return mc
}

// SetNillableCreatedAt sets the "created_at" field if the given value is not nil.
func (mc *MessagesCreate) SetNillableCreatedAt(t *time.Time) *MessagesCreate {
	if t != nil {
		mc.SetCreatedAt(*t)
	}
	return mc
}

// SetUSERID sets the "USER_ID" field.
func (mc *MessagesCreate) SetUSERID(i int) *MessagesCreate {
	mc.mutation.SetUSERID(i)
	return mc
}

// SetUserID sets the "user" edge to the Users entity by ID.
func (mc *MessagesCreate) SetUserID(id int) *MessagesCreate {
	mc.mutation.SetUserID(id)
	return mc
}

// SetNillableUserID sets the "user" edge to the Users entity by ID if the given value is not nil.
func (mc *MessagesCreate) SetNillableUserID(id *int) *MessagesCreate {
	if id != nil {
		mc = mc.SetUserID(*id)
	}
	return mc
}

// SetUser sets the "user" edge to the Users entity.
func (mc *MessagesCreate) SetUser(u *Users) *MessagesCreate {
	return mc.SetUserID(u.ID)
}

// Mutation returns the MessagesMutation object of the builder.
func (mc *MessagesCreate) Mutation() *MessagesMutation {
	return mc.mutation
}

// Save creates the Messages in the database.
func (mc *MessagesCreate) Save(ctx context.Context) (*Messages, error) {
	mc.defaults()
	return withHooks(ctx, mc.sqlSave, mc.mutation, mc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (mc *MessagesCreate) SaveX(ctx context.Context) *Messages {
	v, err := mc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (mc *MessagesCreate) Exec(ctx context.Context) error {
	_, err := mc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (mc *MessagesCreate) ExecX(ctx context.Context) {
	if err := mc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (mc *MessagesCreate) defaults() {
	if _, ok := mc.mutation.CreatedAt(); !ok {
		v := messages.DefaultCreatedAt()
		mc.mutation.SetCreatedAt(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (mc *MessagesCreate) check() error {
	if _, ok := mc.mutation.Content(); !ok {
		return &ValidationError{Name: "content", err: errors.New(`ent: missing required field "Messages.content"`)}
	}
	if _, ok := mc.mutation.CreatedAt(); !ok {
		return &ValidationError{Name: "created_at", err: errors.New(`ent: missing required field "Messages.created_at"`)}
	}
	if _, ok := mc.mutation.USERID(); !ok {
		return &ValidationError{Name: "USER_ID", err: errors.New(`ent: missing required field "Messages.USER_ID"`)}
	}
	return nil
}

func (mc *MessagesCreate) sqlSave(ctx context.Context) (*Messages, error) {
	if err := mc.check(); err != nil {
		return nil, err
	}
	_node, _spec := mc.createSpec()
	if err := sqlgraph.CreateNode(ctx, mc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	mc.mutation.id = &_node.ID
	mc.mutation.done = true
	return _node, nil
}

func (mc *MessagesCreate) createSpec() (*Messages, *sqlgraph.CreateSpec) {
	var (
		_node = &Messages{config: mc.config}
		_spec = sqlgraph.NewCreateSpec(messages.Table, sqlgraph.NewFieldSpec(messages.FieldID, field.TypeInt))
	)
	if value, ok := mc.mutation.Content(); ok {
		_spec.SetField(messages.FieldContent, field.TypeString, value)
		_node.Content = value
	}
	if value, ok := mc.mutation.CreatedAt(); ok {
		_spec.SetField(messages.FieldCreatedAt, field.TypeTime, value)
		_node.CreatedAt = value
	}
	if value, ok := mc.mutation.USERID(); ok {
		_spec.SetField(messages.FieldUSERID, field.TypeInt, value)
		_node.USERID = value
	}
	if nodes := mc.mutation.UserIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: false,
			Table:   messages.UserTable,
			Columns: []string{messages.UserColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(users.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// MessagesCreateBulk is the builder for creating many Messages entities in bulk.
type MessagesCreateBulk struct {
	config
	err      error
	builders []*MessagesCreate
}

// Save creates the Messages entities in the database.
func (mcb *MessagesCreateBulk) Save(ctx context.Context) ([]*Messages, error) {
	if mcb.err != nil {
		return nil, mcb.err
	}
	specs := make([]*sqlgraph.CreateSpec, len(mcb.builders))
	nodes := make([]*Messages, len(mcb.builders))
	mutators := make([]Mutator, len(mcb.builders))
	for i := range mcb.builders {
		func(i int, root context.Context) {
			builder := mcb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*MessagesMutation)
				if !ok {
					return nil, fmt.Errorf("unexpected mutation type %T", m)
				}
				if err := builder.check(); err != nil {
					return nil, err
				}
				builder.mutation = mutation
				var err error
				nodes[i], specs[i] = builder.createSpec()
				if i < len(mutators)-1 {
					_, err = mutators[i+1].Mutate(root, mcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, mcb.driver, spec); err != nil {
						if sqlgraph.IsConstraintError(err) {
							err = &ConstraintError{msg: err.Error(), wrap: err}
						}
					}
				}
				if err != nil {
					return nil, err
				}
				mutation.id = &nodes[i].ID
				if specs[i].ID.Value != nil {
					id := specs[i].ID.Value.(int64)
					nodes[i].ID = int(id)
				}
				mutation.done = true
				return nodes[i], nil
			})
			for i := len(builder.hooks) - 1; i >= 0; i-- {
				mut = builder.hooks[i](mut)
			}
			mutators[i] = mut
		}(i, ctx)
	}
	if len(mutators) > 0 {
		if _, err := mutators[0].Mutate(ctx, mcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (mcb *MessagesCreateBulk) SaveX(ctx context.Context) []*Messages {
	v, err := mcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (mcb *MessagesCreateBulk) Exec(ctx context.Context) error {
	_, err := mcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (mcb *MessagesCreateBulk) ExecX(ctx context.Context) {
	if err := mcb.Exec(ctx); err != nil {
		panic(err)
	}
}