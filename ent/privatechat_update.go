// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"moments/ent/message"
	"moments/ent/predicate"
	"moments/ent/privatechat"
	"moments/ent/user"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// PrivateChatUpdate is the builder for updating PrivateChat entities.
type PrivateChatUpdate struct {
	config
	hooks    []Hook
	mutation *PrivateChatMutation
}

// Where appends a list predicates to the PrivateChatUpdate builder.
func (pcu *PrivateChatUpdate) Where(ps ...predicate.PrivateChat) *PrivateChatUpdate {
	pcu.mutation.Where(ps...)
	return pcu
}

// SetCreatedDate sets the "created_date" field.
func (pcu *PrivateChatUpdate) SetCreatedDate(t time.Time) *PrivateChatUpdate {
	pcu.mutation.SetCreatedDate(t)
	return pcu
}

// SetNillableCreatedDate sets the "created_date" field if the given value is not nil.
func (pcu *PrivateChatUpdate) SetNillableCreatedDate(t *time.Time) *PrivateChatUpdate {
	if t != nil {
		pcu.SetCreatedDate(*t)
	}
	return pcu
}

// SetUpdatedDate sets the "updated_date" field.
func (pcu *PrivateChatUpdate) SetUpdatedDate(t time.Time) *PrivateChatUpdate {
	pcu.mutation.SetUpdatedDate(t)
	return pcu
}

// SetNillableUpdatedDate sets the "updated_date" field if the given value is not nil.
func (pcu *PrivateChatUpdate) SetNillableUpdatedDate(t *time.Time) *PrivateChatUpdate {
	if t != nil {
		pcu.SetUpdatedDate(*t)
	}
	return pcu
}

// SetDeletedDate sets the "deleted_date" field.
func (pcu *PrivateChatUpdate) SetDeletedDate(t time.Time) *PrivateChatUpdate {
	pcu.mutation.SetDeletedDate(t)
	return pcu
}

// SetNillableDeletedDate sets the "deleted_date" field if the given value is not nil.
func (pcu *PrivateChatUpdate) SetNillableDeletedDate(t *time.Time) *PrivateChatUpdate {
	if t != nil {
		pcu.SetDeletedDate(*t)
	}
	return pcu
}

// ClearDeletedDate clears the value of the "deleted_date" field.
func (pcu *PrivateChatUpdate) ClearDeletedDate() *PrivateChatUpdate {
	pcu.mutation.ClearDeletedDate()
	return pcu
}

// SetReceiverID sets the "receiver_id" field.
func (pcu *PrivateChatUpdate) SetReceiverID(i int) *PrivateChatUpdate {
	pcu.mutation.SetReceiverID(i)
	return pcu
}

// SetSenderID sets the "sender_id" field.
func (pcu *PrivateChatUpdate) SetSenderID(i int) *PrivateChatUpdate {
	pcu.mutation.SetSenderID(i)
	return pcu
}

// SetSender sets the "sender" edge to the User entity.
func (pcu *PrivateChatUpdate) SetSender(u *User) *PrivateChatUpdate {
	return pcu.SetSenderID(u.ID)
}

// SetReceiver sets the "receiver" edge to the User entity.
func (pcu *PrivateChatUpdate) SetReceiver(u *User) *PrivateChatUpdate {
	return pcu.SetReceiverID(u.ID)
}

// AddChatIDs adds the "chats" edge to the Message entity by IDs.
func (pcu *PrivateChatUpdate) AddChatIDs(ids ...int) *PrivateChatUpdate {
	pcu.mutation.AddChatIDs(ids...)
	return pcu
}

// AddChats adds the "chats" edges to the Message entity.
func (pcu *PrivateChatUpdate) AddChats(m ...*Message) *PrivateChatUpdate {
	ids := make([]int, len(m))
	for i := range m {
		ids[i] = m[i].ID
	}
	return pcu.AddChatIDs(ids...)
}

// Mutation returns the PrivateChatMutation object of the builder.
func (pcu *PrivateChatUpdate) Mutation() *PrivateChatMutation {
	return pcu.mutation
}

// ClearSender clears the "sender" edge to the User entity.
func (pcu *PrivateChatUpdate) ClearSender() *PrivateChatUpdate {
	pcu.mutation.ClearSender()
	return pcu
}

// ClearReceiver clears the "receiver" edge to the User entity.
func (pcu *PrivateChatUpdate) ClearReceiver() *PrivateChatUpdate {
	pcu.mutation.ClearReceiver()
	return pcu
}

// ClearChats clears all "chats" edges to the Message entity.
func (pcu *PrivateChatUpdate) ClearChats() *PrivateChatUpdate {
	pcu.mutation.ClearChats()
	return pcu
}

// RemoveChatIDs removes the "chats" edge to Message entities by IDs.
func (pcu *PrivateChatUpdate) RemoveChatIDs(ids ...int) *PrivateChatUpdate {
	pcu.mutation.RemoveChatIDs(ids...)
	return pcu
}

// RemoveChats removes "chats" edges to Message entities.
func (pcu *PrivateChatUpdate) RemoveChats(m ...*Message) *PrivateChatUpdate {
	ids := make([]int, len(m))
	for i := range m {
		ids[i] = m[i].ID
	}
	return pcu.RemoveChatIDs(ids...)
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (pcu *PrivateChatUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(pcu.hooks) == 0 {
		if err = pcu.check(); err != nil {
			return 0, err
		}
		affected, err = pcu.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*PrivateChatMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = pcu.check(); err != nil {
				return 0, err
			}
			pcu.mutation = mutation
			affected, err = pcu.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(pcu.hooks) - 1; i >= 0; i-- {
			if pcu.hooks[i] == nil {
				return 0, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = pcu.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, pcu.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (pcu *PrivateChatUpdate) SaveX(ctx context.Context) int {
	affected, err := pcu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (pcu *PrivateChatUpdate) Exec(ctx context.Context) error {
	_, err := pcu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (pcu *PrivateChatUpdate) ExecX(ctx context.Context) {
	if err := pcu.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (pcu *PrivateChatUpdate) check() error {
	if _, ok := pcu.mutation.SenderID(); pcu.mutation.SenderCleared() && !ok {
		return errors.New(`ent: clearing a required unique edge "PrivateChat.sender"`)
	}
	if _, ok := pcu.mutation.ReceiverID(); pcu.mutation.ReceiverCleared() && !ok {
		return errors.New(`ent: clearing a required unique edge "PrivateChat.receiver"`)
	}
	return nil
}

func (pcu *PrivateChatUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   privatechat.Table,
			Columns: privatechat.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: privatechat.FieldID,
			},
		},
	}
	if ps := pcu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := pcu.mutation.CreatedDate(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: privatechat.FieldCreatedDate,
		})
	}
	if value, ok := pcu.mutation.UpdatedDate(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: privatechat.FieldUpdatedDate,
		})
	}
	if value, ok := pcu.mutation.DeletedDate(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: privatechat.FieldDeletedDate,
		})
	}
	if pcu.mutation.DeletedDateCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Column: privatechat.FieldDeletedDate,
		})
	}
	if pcu.mutation.SenderCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := pcu.mutation.SenderIDs(); len(nodes) > 0 {
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
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if pcu.mutation.ReceiverCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := pcu.mutation.ReceiverIDs(); len(nodes) > 0 {
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
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if pcu.mutation.ChatsCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := pcu.mutation.RemovedChatsIDs(); len(nodes) > 0 && !pcu.mutation.ChatsCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := pcu.mutation.ChatsIDs(); len(nodes) > 0 {
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
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, pcu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{privatechat.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	return n, nil
}

// PrivateChatUpdateOne is the builder for updating a single PrivateChat entity.
type PrivateChatUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *PrivateChatMutation
}

// SetCreatedDate sets the "created_date" field.
func (pcuo *PrivateChatUpdateOne) SetCreatedDate(t time.Time) *PrivateChatUpdateOne {
	pcuo.mutation.SetCreatedDate(t)
	return pcuo
}

// SetNillableCreatedDate sets the "created_date" field if the given value is not nil.
func (pcuo *PrivateChatUpdateOne) SetNillableCreatedDate(t *time.Time) *PrivateChatUpdateOne {
	if t != nil {
		pcuo.SetCreatedDate(*t)
	}
	return pcuo
}

// SetUpdatedDate sets the "updated_date" field.
func (pcuo *PrivateChatUpdateOne) SetUpdatedDate(t time.Time) *PrivateChatUpdateOne {
	pcuo.mutation.SetUpdatedDate(t)
	return pcuo
}

// SetNillableUpdatedDate sets the "updated_date" field if the given value is not nil.
func (pcuo *PrivateChatUpdateOne) SetNillableUpdatedDate(t *time.Time) *PrivateChatUpdateOne {
	if t != nil {
		pcuo.SetUpdatedDate(*t)
	}
	return pcuo
}

// SetDeletedDate sets the "deleted_date" field.
func (pcuo *PrivateChatUpdateOne) SetDeletedDate(t time.Time) *PrivateChatUpdateOne {
	pcuo.mutation.SetDeletedDate(t)
	return pcuo
}

// SetNillableDeletedDate sets the "deleted_date" field if the given value is not nil.
func (pcuo *PrivateChatUpdateOne) SetNillableDeletedDate(t *time.Time) *PrivateChatUpdateOne {
	if t != nil {
		pcuo.SetDeletedDate(*t)
	}
	return pcuo
}

// ClearDeletedDate clears the value of the "deleted_date" field.
func (pcuo *PrivateChatUpdateOne) ClearDeletedDate() *PrivateChatUpdateOne {
	pcuo.mutation.ClearDeletedDate()
	return pcuo
}

// SetReceiverID sets the "receiver_id" field.
func (pcuo *PrivateChatUpdateOne) SetReceiverID(i int) *PrivateChatUpdateOne {
	pcuo.mutation.SetReceiverID(i)
	return pcuo
}

// SetSenderID sets the "sender_id" field.
func (pcuo *PrivateChatUpdateOne) SetSenderID(i int) *PrivateChatUpdateOne {
	pcuo.mutation.SetSenderID(i)
	return pcuo
}

// SetSender sets the "sender" edge to the User entity.
func (pcuo *PrivateChatUpdateOne) SetSender(u *User) *PrivateChatUpdateOne {
	return pcuo.SetSenderID(u.ID)
}

// SetReceiver sets the "receiver" edge to the User entity.
func (pcuo *PrivateChatUpdateOne) SetReceiver(u *User) *PrivateChatUpdateOne {
	return pcuo.SetReceiverID(u.ID)
}

// AddChatIDs adds the "chats" edge to the Message entity by IDs.
func (pcuo *PrivateChatUpdateOne) AddChatIDs(ids ...int) *PrivateChatUpdateOne {
	pcuo.mutation.AddChatIDs(ids...)
	return pcuo
}

// AddChats adds the "chats" edges to the Message entity.
func (pcuo *PrivateChatUpdateOne) AddChats(m ...*Message) *PrivateChatUpdateOne {
	ids := make([]int, len(m))
	for i := range m {
		ids[i] = m[i].ID
	}
	return pcuo.AddChatIDs(ids...)
}

// Mutation returns the PrivateChatMutation object of the builder.
func (pcuo *PrivateChatUpdateOne) Mutation() *PrivateChatMutation {
	return pcuo.mutation
}

// ClearSender clears the "sender" edge to the User entity.
func (pcuo *PrivateChatUpdateOne) ClearSender() *PrivateChatUpdateOne {
	pcuo.mutation.ClearSender()
	return pcuo
}

// ClearReceiver clears the "receiver" edge to the User entity.
func (pcuo *PrivateChatUpdateOne) ClearReceiver() *PrivateChatUpdateOne {
	pcuo.mutation.ClearReceiver()
	return pcuo
}

// ClearChats clears all "chats" edges to the Message entity.
func (pcuo *PrivateChatUpdateOne) ClearChats() *PrivateChatUpdateOne {
	pcuo.mutation.ClearChats()
	return pcuo
}

// RemoveChatIDs removes the "chats" edge to Message entities by IDs.
func (pcuo *PrivateChatUpdateOne) RemoveChatIDs(ids ...int) *PrivateChatUpdateOne {
	pcuo.mutation.RemoveChatIDs(ids...)
	return pcuo
}

// RemoveChats removes "chats" edges to Message entities.
func (pcuo *PrivateChatUpdateOne) RemoveChats(m ...*Message) *PrivateChatUpdateOne {
	ids := make([]int, len(m))
	for i := range m {
		ids[i] = m[i].ID
	}
	return pcuo.RemoveChatIDs(ids...)
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (pcuo *PrivateChatUpdateOne) Select(field string, fields ...string) *PrivateChatUpdateOne {
	pcuo.fields = append([]string{field}, fields...)
	return pcuo
}

// Save executes the query and returns the updated PrivateChat entity.
func (pcuo *PrivateChatUpdateOne) Save(ctx context.Context) (*PrivateChat, error) {
	var (
		err  error
		node *PrivateChat
	)
	if len(pcuo.hooks) == 0 {
		if err = pcuo.check(); err != nil {
			return nil, err
		}
		node, err = pcuo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*PrivateChatMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = pcuo.check(); err != nil {
				return nil, err
			}
			pcuo.mutation = mutation
			node, err = pcuo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(pcuo.hooks) - 1; i >= 0; i-- {
			if pcuo.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = pcuo.hooks[i](mut)
		}
		v, err := mut.Mutate(ctx, pcuo.mutation)
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

// SaveX is like Save, but panics if an error occurs.
func (pcuo *PrivateChatUpdateOne) SaveX(ctx context.Context) *PrivateChat {
	node, err := pcuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (pcuo *PrivateChatUpdateOne) Exec(ctx context.Context) error {
	_, err := pcuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (pcuo *PrivateChatUpdateOne) ExecX(ctx context.Context) {
	if err := pcuo.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (pcuo *PrivateChatUpdateOne) check() error {
	if _, ok := pcuo.mutation.SenderID(); pcuo.mutation.SenderCleared() && !ok {
		return errors.New(`ent: clearing a required unique edge "PrivateChat.sender"`)
	}
	if _, ok := pcuo.mutation.ReceiverID(); pcuo.mutation.ReceiverCleared() && !ok {
		return errors.New(`ent: clearing a required unique edge "PrivateChat.receiver"`)
	}
	return nil
}

func (pcuo *PrivateChatUpdateOne) sqlSave(ctx context.Context) (_node *PrivateChat, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   privatechat.Table,
			Columns: privatechat.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: privatechat.FieldID,
			},
		},
	}
	id, ok := pcuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "PrivateChat.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := pcuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, privatechat.FieldID)
		for _, f := range fields {
			if !privatechat.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != privatechat.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := pcuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := pcuo.mutation.CreatedDate(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: privatechat.FieldCreatedDate,
		})
	}
	if value, ok := pcuo.mutation.UpdatedDate(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: privatechat.FieldUpdatedDate,
		})
	}
	if value, ok := pcuo.mutation.DeletedDate(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: privatechat.FieldDeletedDate,
		})
	}
	if pcuo.mutation.DeletedDateCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Column: privatechat.FieldDeletedDate,
		})
	}
	if pcuo.mutation.SenderCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := pcuo.mutation.SenderIDs(); len(nodes) > 0 {
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
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if pcuo.mutation.ReceiverCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := pcuo.mutation.ReceiverIDs(); len(nodes) > 0 {
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
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if pcuo.mutation.ChatsCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := pcuo.mutation.RemovedChatsIDs(); len(nodes) > 0 && !pcuo.mutation.ChatsCleared() {
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
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := pcuo.mutation.ChatsIDs(); len(nodes) > 0 {
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
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &PrivateChat{config: pcuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, pcuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{privatechat.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	return _node, nil
}
