// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"shor_url/ent/predicate"
	"shor_url/ent/tinyurl"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// TinyURLUpdate is the builder for updating TinyURL entities.
type TinyURLUpdate struct {
	config
	hooks    []Hook
	mutation *TinyURLMutation
}

// Where appends a list predicates to the TinyURLUpdate builder.
func (tuu *TinyURLUpdate) Where(ps ...predicate.TinyURL) *TinyURLUpdate {
	tuu.mutation.Where(ps...)
	return tuu
}

// SetURL sets the "url" field.
func (tuu *TinyURLUpdate) SetURL(s string) *TinyURLUpdate {
	tuu.mutation.SetURL(s)
	return tuu
}

// Mutation returns the TinyURLMutation object of the builder.
func (tuu *TinyURLUpdate) Mutation() *TinyURLMutation {
	return tuu.mutation
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (tuu *TinyURLUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(tuu.hooks) == 0 {
		affected, err = tuu.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*TinyURLMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			tuu.mutation = mutation
			affected, err = tuu.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(tuu.hooks) - 1; i >= 0; i-- {
			if tuu.hooks[i] == nil {
				return 0, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = tuu.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, tuu.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (tuu *TinyURLUpdate) SaveX(ctx context.Context) int {
	affected, err := tuu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (tuu *TinyURLUpdate) Exec(ctx context.Context) error {
	_, err := tuu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (tuu *TinyURLUpdate) ExecX(ctx context.Context) {
	if err := tuu.Exec(ctx); err != nil {
		panic(err)
	}
}

func (tuu *TinyURLUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   tinyurl.Table,
			Columns: tinyurl.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUint64,
				Column: tinyurl.FieldID,
			},
		},
	}
	if ps := tuu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := tuu.mutation.URL(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: tinyurl.FieldURL,
		})
	}
	if n, err = sqlgraph.UpdateNodes(ctx, tuu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{tinyurl.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return 0, err
	}
	return n, nil
}

// TinyURLUpdateOne is the builder for updating a single TinyURL entity.
type TinyURLUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *TinyURLMutation
}

// SetURL sets the "url" field.
func (tuuo *TinyURLUpdateOne) SetURL(s string) *TinyURLUpdateOne {
	tuuo.mutation.SetURL(s)
	return tuuo
}

// Mutation returns the TinyURLMutation object of the builder.
func (tuuo *TinyURLUpdateOne) Mutation() *TinyURLMutation {
	return tuuo.mutation
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (tuuo *TinyURLUpdateOne) Select(field string, fields ...string) *TinyURLUpdateOne {
	tuuo.fields = append([]string{field}, fields...)
	return tuuo
}

// Save executes the query and returns the updated TinyURL entity.
func (tuuo *TinyURLUpdateOne) Save(ctx context.Context) (*TinyURL, error) {
	var (
		err  error
		node *TinyURL
	)
	if len(tuuo.hooks) == 0 {
		node, err = tuuo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*TinyURLMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			tuuo.mutation = mutation
			node, err = tuuo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(tuuo.hooks) - 1; i >= 0; i-- {
			if tuuo.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = tuuo.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, tuuo.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (tuuo *TinyURLUpdateOne) SaveX(ctx context.Context) *TinyURL {
	node, err := tuuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (tuuo *TinyURLUpdateOne) Exec(ctx context.Context) error {
	_, err := tuuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (tuuo *TinyURLUpdateOne) ExecX(ctx context.Context) {
	if err := tuuo.Exec(ctx); err != nil {
		panic(err)
	}
}

func (tuuo *TinyURLUpdateOne) sqlSave(ctx context.Context) (_node *TinyURL, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   tinyurl.Table,
			Columns: tinyurl.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUint64,
				Column: tinyurl.FieldID,
			},
		},
	}
	id, ok := tuuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "TinyURL.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := tuuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, tinyurl.FieldID)
		for _, f := range fields {
			if !tinyurl.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != tinyurl.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := tuuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := tuuo.mutation.URL(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: tinyurl.FieldURL,
		})
	}
	_node = &TinyURL{config: tuuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, tuuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{tinyurl.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return nil, err
	}
	return _node, nil
}
