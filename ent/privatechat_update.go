// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"moments/ent/predicate"
	"moments/ent/privatechat"
	"moments/ent/user"

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

// SetFirstUserID sets the "first_user_id" field.
func (pcu *PrivateChatUpdate) SetFirstUserID(i int) *PrivateChatUpdate {
	pcu.mutation.SetFirstUserID(i)
	return pcu
}

// SetNillableFirstUserID sets the "first_user_id" field if the given value is not nil.
func (pcu *PrivateChatUpdate) SetNillableFirstUserID(i *int) *PrivateChatUpdate {
	if i != nil {
		pcu.SetFirstUserID(*i)
	}
	return pcu
}

// ClearFirstUserID clears the value of the "first_user_id" field.
func (pcu *PrivateChatUpdate) ClearFirstUserID() *PrivateChatUpdate {
	pcu.mutation.ClearFirstUserID()
	return pcu
}

// SetSecondUserID sets the "second_user_id" field.
func (pcu *PrivateChatUpdate) SetSecondUserID(i int) *PrivateChatUpdate {
	pcu.mutation.SetSecondUserID(i)
	return pcu
}

// SetNillableSecondUserID sets the "second_user_id" field if the given value is not nil.
func (pcu *PrivateChatUpdate) SetNillableSecondUserID(i *int) *PrivateChatUpdate {
	if i != nil {
		pcu.SetSecondUserID(*i)
	}
	return pcu
}

// ClearSecondUserID clears the value of the "second_user_id" field.
func (pcu *PrivateChatUpdate) ClearSecondUserID() *PrivateChatUpdate {
	pcu.mutation.ClearSecondUserID()
	return pcu
}

// SetFirstUser sets the "first_user" edge to the User entity.
func (pcu *PrivateChatUpdate) SetFirstUser(u *User) *PrivateChatUpdate {
	return pcu.SetFirstUserID(u.ID)
}

// SetSecondUser sets the "second_user" edge to the User entity.
func (pcu *PrivateChatUpdate) SetSecondUser(u *User) *PrivateChatUpdate {
	return pcu.SetSecondUserID(u.ID)
}

// Mutation returns the PrivateChatMutation object of the builder.
func (pcu *PrivateChatUpdate) Mutation() *PrivateChatMutation {
	return pcu.mutation
}

// ClearFirstUser clears the "first_user" edge to the User entity.
func (pcu *PrivateChatUpdate) ClearFirstUser() *PrivateChatUpdate {
	pcu.mutation.ClearFirstUser()
	return pcu
}

// ClearSecondUser clears the "second_user" edge to the User entity.
func (pcu *PrivateChatUpdate) ClearSecondUser() *PrivateChatUpdate {
	pcu.mutation.ClearSecondUser()
	return pcu
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (pcu *PrivateChatUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(pcu.hooks) == 0 {
		affected, err = pcu.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*PrivateChatMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
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
	if pcu.mutation.FirstUserCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   privatechat.FirstUserTable,
			Columns: []string{privatechat.FirstUserColumn},
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
	if nodes := pcu.mutation.FirstUserIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   privatechat.FirstUserTable,
			Columns: []string{privatechat.FirstUserColumn},
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
	if pcu.mutation.SecondUserCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   privatechat.SecondUserTable,
			Columns: []string{privatechat.SecondUserColumn},
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
	if nodes := pcu.mutation.SecondUserIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   privatechat.SecondUserTable,
			Columns: []string{privatechat.SecondUserColumn},
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

// SetFirstUserID sets the "first_user_id" field.
func (pcuo *PrivateChatUpdateOne) SetFirstUserID(i int) *PrivateChatUpdateOne {
	pcuo.mutation.SetFirstUserID(i)
	return pcuo
}

// SetNillableFirstUserID sets the "first_user_id" field if the given value is not nil.
func (pcuo *PrivateChatUpdateOne) SetNillableFirstUserID(i *int) *PrivateChatUpdateOne {
	if i != nil {
		pcuo.SetFirstUserID(*i)
	}
	return pcuo
}

// ClearFirstUserID clears the value of the "first_user_id" field.
func (pcuo *PrivateChatUpdateOne) ClearFirstUserID() *PrivateChatUpdateOne {
	pcuo.mutation.ClearFirstUserID()
	return pcuo
}

// SetSecondUserID sets the "second_user_id" field.
func (pcuo *PrivateChatUpdateOne) SetSecondUserID(i int) *PrivateChatUpdateOne {
	pcuo.mutation.SetSecondUserID(i)
	return pcuo
}

// SetNillableSecondUserID sets the "second_user_id" field if the given value is not nil.
func (pcuo *PrivateChatUpdateOne) SetNillableSecondUserID(i *int) *PrivateChatUpdateOne {
	if i != nil {
		pcuo.SetSecondUserID(*i)
	}
	return pcuo
}

// ClearSecondUserID clears the value of the "second_user_id" field.
func (pcuo *PrivateChatUpdateOne) ClearSecondUserID() *PrivateChatUpdateOne {
	pcuo.mutation.ClearSecondUserID()
	return pcuo
}

// SetFirstUser sets the "first_user" edge to the User entity.
func (pcuo *PrivateChatUpdateOne) SetFirstUser(u *User) *PrivateChatUpdateOne {
	return pcuo.SetFirstUserID(u.ID)
}

// SetSecondUser sets the "second_user" edge to the User entity.
func (pcuo *PrivateChatUpdateOne) SetSecondUser(u *User) *PrivateChatUpdateOne {
	return pcuo.SetSecondUserID(u.ID)
}

// Mutation returns the PrivateChatMutation object of the builder.
func (pcuo *PrivateChatUpdateOne) Mutation() *PrivateChatMutation {
	return pcuo.mutation
}

// ClearFirstUser clears the "first_user" edge to the User entity.
func (pcuo *PrivateChatUpdateOne) ClearFirstUser() *PrivateChatUpdateOne {
	pcuo.mutation.ClearFirstUser()
	return pcuo
}

// ClearSecondUser clears the "second_user" edge to the User entity.
func (pcuo *PrivateChatUpdateOne) ClearSecondUser() *PrivateChatUpdateOne {
	pcuo.mutation.ClearSecondUser()
	return pcuo
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
		node, err = pcuo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*PrivateChatMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
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
	if pcuo.mutation.FirstUserCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   privatechat.FirstUserTable,
			Columns: []string{privatechat.FirstUserColumn},
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
	if nodes := pcuo.mutation.FirstUserIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   privatechat.FirstUserTable,
			Columns: []string{privatechat.FirstUserColumn},
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
	if pcuo.mutation.SecondUserCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   privatechat.SecondUserTable,
			Columns: []string{privatechat.SecondUserColumn},
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
	if nodes := pcuo.mutation.SecondUserIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   privatechat.SecondUserTable,
			Columns: []string{privatechat.SecondUserColumn},
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
