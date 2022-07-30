// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"moments/ent/message"
	"moments/ent/privatechat"
	"moments/ent/user"
	"time"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// PrivateChatCreate is the builder for creating a PrivateChat entity.
type PrivateChatCreate struct {
	config
	mutation *PrivateChatMutation
	hooks    []Hook
}

// SetCreatedDate sets the "created_date" field.
func (pcc *PrivateChatCreate) SetCreatedDate(t time.Time) *PrivateChatCreate {
	pcc.mutation.SetCreatedDate(t)
	return pcc
}

// SetNillableCreatedDate sets the "created_date" field if the given value is not nil.
func (pcc *PrivateChatCreate) SetNillableCreatedDate(t *time.Time) *PrivateChatCreate {
	if t != nil {
		pcc.SetCreatedDate(*t)
	}
	return pcc
}

// SetUpdatedDate sets the "updated_date" field.
func (pcc *PrivateChatCreate) SetUpdatedDate(t time.Time) *PrivateChatCreate {
	pcc.mutation.SetUpdatedDate(t)
	return pcc
}

// SetNillableUpdatedDate sets the "updated_date" field if the given value is not nil.
func (pcc *PrivateChatCreate) SetNillableUpdatedDate(t *time.Time) *PrivateChatCreate {
	if t != nil {
		pcc.SetUpdatedDate(*t)
	}
	return pcc
}

// SetDeletedDate sets the "deleted_date" field.
func (pcc *PrivateChatCreate) SetDeletedDate(t time.Time) *PrivateChatCreate {
	pcc.mutation.SetDeletedDate(t)
	return pcc
}

// SetNillableDeletedDate sets the "deleted_date" field if the given value is not nil.
func (pcc *PrivateChatCreate) SetNillableDeletedDate(t *time.Time) *PrivateChatCreate {
	if t != nil {
		pcc.SetDeletedDate(*t)
	}
	return pcc
}

// SetReceiverID sets the "receiver_id" field.
func (pcc *PrivateChatCreate) SetReceiverID(i int) *PrivateChatCreate {
	pcc.mutation.SetReceiverID(i)
	return pcc
}

// SetSenderID sets the "sender_id" field.
func (pcc *PrivateChatCreate) SetSenderID(i int) *PrivateChatCreate {
	pcc.mutation.SetSenderID(i)
	return pcc
}

// SetSender sets the "sender" edge to the User entity.
func (pcc *PrivateChatCreate) SetSender(u *User) *PrivateChatCreate {
	return pcc.SetSenderID(u.ID)
}

// SetReceiver sets the "receiver" edge to the User entity.
func (pcc *PrivateChatCreate) SetReceiver(u *User) *PrivateChatCreate {
	return pcc.SetReceiverID(u.ID)
}

// AddChatIDs adds the "chats" edge to the Message entity by IDs.
func (pcc *PrivateChatCreate) AddChatIDs(ids ...int) *PrivateChatCreate {
	pcc.mutation.AddChatIDs(ids...)
	return pcc
}

// AddChats adds the "chats" edges to the Message entity.
func (pcc *PrivateChatCreate) AddChats(m ...*Message) *PrivateChatCreate {
	ids := make([]int, len(m))
	for i := range m {
		ids[i] = m[i].ID
	}
	return pcc.AddChatIDs(ids...)
}

// Mutation returns the PrivateChatMutation object of the builder.
func (pcc *PrivateChatCreate) Mutation() *PrivateChatMutation {
	return pcc.mutation
}

// Save creates the PrivateChat in the database.
func (pcc *PrivateChatCreate) Save(ctx context.Context) (*PrivateChat, error) {
	var (
		err  error
		node *PrivateChat
	)
	pcc.defaults()
	if len(pcc.hooks) == 0 {
		if err = pcc.check(); err != nil {
			return nil, err
		}
		node, err = pcc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*PrivateChatMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = pcc.check(); err != nil {
				return nil, err
			}
			pcc.mutation = mutation
			if node, err = pcc.sqlSave(ctx); err != nil {
				return nil, err
			}
			mutation.id = &node.ID
			mutation.done = true
			return node, err
		})
		for i := len(pcc.hooks) - 1; i >= 0; i-- {
			if pcc.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = pcc.hooks[i](mut)
		}
		v, err := mut.Mutate(ctx, pcc.mutation)
		if err != nil {
			return nil, err
		}
		nv, ok := v.(*PrivateChat)
		if !ok {
			return nil, fmt.Errorf("unexpected node type %T returned from PrivateChatMutation", v)
		}
		node = nv
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (pcc *PrivateChatCreate) SaveX(ctx context.Context) *PrivateChat {
	v, err := pcc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (pcc *PrivateChatCreate) Exec(ctx context.Context) error {
	_, err := pcc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (pcc *PrivateChatCreate) ExecX(ctx context.Context) {
	if err := pcc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (pcc *PrivateChatCreate) defaults() {
	if _, ok := pcc.mutation.CreatedDate(); !ok {
		v := privatechat.DefaultCreatedDate
		pcc.mutation.SetCreatedDate(v)
	}
	if _, ok := pcc.mutation.UpdatedDate(); !ok {
		v := privatechat.DefaultUpdatedDate
		pcc.mutation.SetUpdatedDate(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (pcc *PrivateChatCreate) check() error {
	if _, ok := pcc.mutation.CreatedDate(); !ok {
		return &ValidationError{Name: "created_date", err: errors.New(`ent: missing required field "PrivateChat.created_date"`)}
	}
	if _, ok := pcc.mutation.UpdatedDate(); !ok {
		return &ValidationError{Name: "updated_date", err: errors.New(`ent: missing required field "PrivateChat.updated_date"`)}
	}
	if _, ok := pcc.mutation.ReceiverID(); !ok {
		return &ValidationError{Name: "receiver_id", err: errors.New(`ent: missing required field "PrivateChat.receiver_id"`)}
	}
	if _, ok := pcc.mutation.SenderID(); !ok {
		return &ValidationError{Name: "sender_id", err: errors.New(`ent: missing required field "PrivateChat.sender_id"`)}
	}
	if _, ok := pcc.mutation.SenderID(); !ok {
		return &ValidationError{Name: "sender", err: errors.New(`ent: missing required edge "PrivateChat.sender"`)}
	}
	if _, ok := pcc.mutation.ReceiverID(); !ok {
		return &ValidationError{Name: "receiver", err: errors.New(`ent: missing required edge "PrivateChat.receiver"`)}
	}
	return nil
}

func (pcc *PrivateChatCreate) sqlSave(ctx context.Context) (*PrivateChat, error) {
	_node, _spec := pcc.createSpec()
	if err := sqlgraph.CreateNode(ctx, pcc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	return _node, nil
}

func (pcc *PrivateChatCreate) createSpec() (*PrivateChat, *sqlgraph.CreateSpec) {
	var (
		_node = &PrivateChat{config: pcc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: privatechat.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: privatechat.FieldID,
			},
		}
	)
	if value, ok := pcc.mutation.CreatedDate(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: privatechat.FieldCreatedDate,
		})
		_node.CreatedDate = value
	}
	if value, ok := pcc.mutation.UpdatedDate(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: privatechat.FieldUpdatedDate,
		})
		_node.UpdatedDate = value
	}
	if value, ok := pcc.mutation.DeletedDate(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: privatechat.FieldDeletedDate,
		})
		_node.DeletedDate = &value
	}
	if nodes := pcc.mutation.SenderIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   privatechat.SenderTable,
			Columns: []string{privatechat.SenderColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: user.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.SenderID = nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := pcc.mutation.ReceiverIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   privatechat.ReceiverTable,
			Columns: []string{privatechat.ReceiverColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: user.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.ReceiverID = nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := pcc.mutation.ChatsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   privatechat.ChatsTable,
			Columns: []string{privatechat.ChatsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt,
					Column: message.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// PrivateChatCreateBulk is the builder for creating many PrivateChat entities in bulk.
type PrivateChatCreateBulk struct {
	config
	builders []*PrivateChatCreate
}

// Save creates the PrivateChat entities in the database.
func (pccb *PrivateChatCreateBulk) Save(ctx context.Context) ([]*PrivateChat, error) {
	specs := make([]*sqlgraph.CreateSpec, len(pccb.builders))
	nodes := make([]*PrivateChat, len(pccb.builders))
	mutators := make([]Mutator, len(pccb.builders))
	for i := range pccb.builders {
		func(i int, root context.Context) {
			builder := pccb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*PrivateChatMutation)
				if !ok {
					return nil, fmt.Errorf("unexpected mutation type %T", m)
				}
				if err := builder.check(); err != nil {
					return nil, err
				}
				builder.mutation = mutation
				nodes[i], specs[i] = builder.createSpec()
				var err error
				if i < len(mutators)-1 {
					_, err = mutators[i+1].Mutate(root, pccb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, pccb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, pccb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (pccb *PrivateChatCreateBulk) SaveX(ctx context.Context) []*PrivateChat {
	v, err := pccb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (pccb *PrivateChatCreateBulk) Exec(ctx context.Context) error {
	_, err := pccb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (pccb *PrivateChatCreateBulk) ExecX(ctx context.Context) {
	if err := pccb.Exec(ctx); err != nil {
		panic(err)
	}
}
