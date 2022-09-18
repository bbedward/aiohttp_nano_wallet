// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/appditto/pippin_nano_wallet/libs/database/ent/account"
	"github.com/appditto/pippin_nano_wallet/libs/database/ent/predicate"
	"github.com/appditto/pippin_nano_wallet/libs/database/ent/wallet"
	"github.com/google/uuid"
)

// WalletUpdate is the builder for updating Wallet entities.
type WalletUpdate struct {
	config
	hooks    []Hook
	mutation *WalletMutation
}

// Where appends a list predicates to the WalletUpdate builder.
func (wu *WalletUpdate) Where(ps ...predicate.Wallet) *WalletUpdate {
	wu.mutation.Where(ps...)
	return wu
}

// SetSeed sets the "seed" field.
func (wu *WalletUpdate) SetSeed(s string) *WalletUpdate {
	wu.mutation.SetSeed(s)
	return wu
}

// SetRepresentative sets the "representative" field.
func (wu *WalletUpdate) SetRepresentative(s string) *WalletUpdate {
	wu.mutation.SetRepresentative(s)
	return wu
}

// SetNillableRepresentative sets the "representative" field if the given value is not nil.
func (wu *WalletUpdate) SetNillableRepresentative(s *string) *WalletUpdate {
	if s != nil {
		wu.SetRepresentative(*s)
	}
	return wu
}

// ClearRepresentative clears the value of the "representative" field.
func (wu *WalletUpdate) ClearRepresentative() *WalletUpdate {
	wu.mutation.ClearRepresentative()
	return wu
}

// SetEncrypted sets the "encrypted" field.
func (wu *WalletUpdate) SetEncrypted(b bool) *WalletUpdate {
	wu.mutation.SetEncrypted(b)
	return wu
}

// SetNillableEncrypted sets the "encrypted" field if the given value is not nil.
func (wu *WalletUpdate) SetNillableEncrypted(b *bool) *WalletUpdate {
	if b != nil {
		wu.SetEncrypted(*b)
	}
	return wu
}

// SetWork sets the "work" field.
func (wu *WalletUpdate) SetWork(b bool) *WalletUpdate {
	wu.mutation.SetWork(b)
	return wu
}

// SetNillableWork sets the "work" field if the given value is not nil.
func (wu *WalletUpdate) SetNillableWork(b *bool) *WalletUpdate {
	if b != nil {
		wu.SetWork(*b)
	}
	return wu
}

// AddAccountIDs adds the "accounts" edge to the Account entity by IDs.
func (wu *WalletUpdate) AddAccountIDs(ids ...uuid.UUID) *WalletUpdate {
	wu.mutation.AddAccountIDs(ids...)
	return wu
}

// AddAccounts adds the "accounts" edges to the Account entity.
func (wu *WalletUpdate) AddAccounts(a ...*Account) *WalletUpdate {
	ids := make([]uuid.UUID, len(a))
	for i := range a {
		ids[i] = a[i].ID
	}
	return wu.AddAccountIDs(ids...)
}

// Mutation returns the WalletMutation object of the builder.
func (wu *WalletUpdate) Mutation() *WalletMutation {
	return wu.mutation
}

// ClearAccounts clears all "accounts" edges to the Account entity.
func (wu *WalletUpdate) ClearAccounts() *WalletUpdate {
	wu.mutation.ClearAccounts()
	return wu
}

// RemoveAccountIDs removes the "accounts" edge to Account entities by IDs.
func (wu *WalletUpdate) RemoveAccountIDs(ids ...uuid.UUID) *WalletUpdate {
	wu.mutation.RemoveAccountIDs(ids...)
	return wu
}

// RemoveAccounts removes "accounts" edges to Account entities.
func (wu *WalletUpdate) RemoveAccounts(a ...*Account) *WalletUpdate {
	ids := make([]uuid.UUID, len(a))
	for i := range a {
		ids[i] = a[i].ID
	}
	return wu.RemoveAccountIDs(ids...)
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (wu *WalletUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(wu.hooks) == 0 {
		if err = wu.check(); err != nil {
			return 0, err
		}
		affected, err = wu.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*WalletMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = wu.check(); err != nil {
				return 0, err
			}
			wu.mutation = mutation
			affected, err = wu.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(wu.hooks) - 1; i >= 0; i-- {
			if wu.hooks[i] == nil {
				return 0, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = wu.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, wu.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (wu *WalletUpdate) SaveX(ctx context.Context) int {
	affected, err := wu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (wu *WalletUpdate) Exec(ctx context.Context) error {
	_, err := wu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (wu *WalletUpdate) ExecX(ctx context.Context) {
	if err := wu.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (wu *WalletUpdate) check() error {
	if v, ok := wu.mutation.Seed(); ok {
		if err := wallet.SeedValidator(v); err != nil {
			return &ValidationError{Name: "seed", err: fmt.Errorf(`ent: validator failed for field "Wallet.seed": %w`, err)}
		}
	}
	if v, ok := wu.mutation.Representative(); ok {
		if err := wallet.RepresentativeValidator(v); err != nil {
			return &ValidationError{Name: "representative", err: fmt.Errorf(`ent: validator failed for field "Wallet.representative": %w`, err)}
		}
	}
	return nil
}

func (wu *WalletUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   wallet.Table,
			Columns: wallet.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: wallet.FieldID,
			},
		},
	}
	if ps := wu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := wu.mutation.Seed(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: wallet.FieldSeed,
		})
	}
	if value, ok := wu.mutation.Representative(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: wallet.FieldRepresentative,
		})
	}
	if wu.mutation.RepresentativeCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: wallet.FieldRepresentative,
		})
	}
	if value, ok := wu.mutation.Encrypted(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeBool,
			Value:  value,
			Column: wallet.FieldEncrypted,
		})
	}
	if value, ok := wu.mutation.Work(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeBool,
			Value:  value,
			Column: wallet.FieldWork,
		})
	}
	if wu.mutation.AccountsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   wallet.AccountsTable,
			Columns: []string{wallet.AccountsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: account.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := wu.mutation.RemovedAccountsIDs(); len(nodes) > 0 && !wu.mutation.AccountsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   wallet.AccountsTable,
			Columns: []string{wallet.AccountsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: account.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := wu.mutation.AccountsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   wallet.AccountsTable,
			Columns: []string{wallet.AccountsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: account.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, wu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{wallet.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	return n, nil
}

// WalletUpdateOne is the builder for updating a single Wallet entity.
type WalletUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *WalletMutation
}

// SetSeed sets the "seed" field.
func (wuo *WalletUpdateOne) SetSeed(s string) *WalletUpdateOne {
	wuo.mutation.SetSeed(s)
	return wuo
}

// SetRepresentative sets the "representative" field.
func (wuo *WalletUpdateOne) SetRepresentative(s string) *WalletUpdateOne {
	wuo.mutation.SetRepresentative(s)
	return wuo
}

// SetNillableRepresentative sets the "representative" field if the given value is not nil.
func (wuo *WalletUpdateOne) SetNillableRepresentative(s *string) *WalletUpdateOne {
	if s != nil {
		wuo.SetRepresentative(*s)
	}
	return wuo
}

// ClearRepresentative clears the value of the "representative" field.
func (wuo *WalletUpdateOne) ClearRepresentative() *WalletUpdateOne {
	wuo.mutation.ClearRepresentative()
	return wuo
}

// SetEncrypted sets the "encrypted" field.
func (wuo *WalletUpdateOne) SetEncrypted(b bool) *WalletUpdateOne {
	wuo.mutation.SetEncrypted(b)
	return wuo
}

// SetNillableEncrypted sets the "encrypted" field if the given value is not nil.
func (wuo *WalletUpdateOne) SetNillableEncrypted(b *bool) *WalletUpdateOne {
	if b != nil {
		wuo.SetEncrypted(*b)
	}
	return wuo
}

// SetWork sets the "work" field.
func (wuo *WalletUpdateOne) SetWork(b bool) *WalletUpdateOne {
	wuo.mutation.SetWork(b)
	return wuo
}

// SetNillableWork sets the "work" field if the given value is not nil.
func (wuo *WalletUpdateOne) SetNillableWork(b *bool) *WalletUpdateOne {
	if b != nil {
		wuo.SetWork(*b)
	}
	return wuo
}

// AddAccountIDs adds the "accounts" edge to the Account entity by IDs.
func (wuo *WalletUpdateOne) AddAccountIDs(ids ...uuid.UUID) *WalletUpdateOne {
	wuo.mutation.AddAccountIDs(ids...)
	return wuo
}

// AddAccounts adds the "accounts" edges to the Account entity.
func (wuo *WalletUpdateOne) AddAccounts(a ...*Account) *WalletUpdateOne {
	ids := make([]uuid.UUID, len(a))
	for i := range a {
		ids[i] = a[i].ID
	}
	return wuo.AddAccountIDs(ids...)
}

// Mutation returns the WalletMutation object of the builder.
func (wuo *WalletUpdateOne) Mutation() *WalletMutation {
	return wuo.mutation
}

// ClearAccounts clears all "accounts" edges to the Account entity.
func (wuo *WalletUpdateOne) ClearAccounts() *WalletUpdateOne {
	wuo.mutation.ClearAccounts()
	return wuo
}

// RemoveAccountIDs removes the "accounts" edge to Account entities by IDs.
func (wuo *WalletUpdateOne) RemoveAccountIDs(ids ...uuid.UUID) *WalletUpdateOne {
	wuo.mutation.RemoveAccountIDs(ids...)
	return wuo
}

// RemoveAccounts removes "accounts" edges to Account entities.
func (wuo *WalletUpdateOne) RemoveAccounts(a ...*Account) *WalletUpdateOne {
	ids := make([]uuid.UUID, len(a))
	for i := range a {
		ids[i] = a[i].ID
	}
	return wuo.RemoveAccountIDs(ids...)
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (wuo *WalletUpdateOne) Select(field string, fields ...string) *WalletUpdateOne {
	wuo.fields = append([]string{field}, fields...)
	return wuo
}

// Save executes the query and returns the updated Wallet entity.
func (wuo *WalletUpdateOne) Save(ctx context.Context) (*Wallet, error) {
	var (
		err  error
		node *Wallet
	)
	if len(wuo.hooks) == 0 {
		if err = wuo.check(); err != nil {
			return nil, err
		}
		node, err = wuo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*WalletMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = wuo.check(); err != nil {
				return nil, err
			}
			wuo.mutation = mutation
			node, err = wuo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(wuo.hooks) - 1; i >= 0; i-- {
			if wuo.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = wuo.hooks[i](mut)
		}
		v, err := mut.Mutate(ctx, wuo.mutation)
		if err != nil {
			return nil, err
		}
		nv, ok := v.(*Wallet)
		if !ok {
			return nil, fmt.Errorf("unexpected node type %T returned from WalletMutation", v)
		}
		node = nv
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (wuo *WalletUpdateOne) SaveX(ctx context.Context) *Wallet {
	node, err := wuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (wuo *WalletUpdateOne) Exec(ctx context.Context) error {
	_, err := wuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (wuo *WalletUpdateOne) ExecX(ctx context.Context) {
	if err := wuo.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (wuo *WalletUpdateOne) check() error {
	if v, ok := wuo.mutation.Seed(); ok {
		if err := wallet.SeedValidator(v); err != nil {
			return &ValidationError{Name: "seed", err: fmt.Errorf(`ent: validator failed for field "Wallet.seed": %w`, err)}
		}
	}
	if v, ok := wuo.mutation.Representative(); ok {
		if err := wallet.RepresentativeValidator(v); err != nil {
			return &ValidationError{Name: "representative", err: fmt.Errorf(`ent: validator failed for field "Wallet.representative": %w`, err)}
		}
	}
	return nil
}

func (wuo *WalletUpdateOne) sqlSave(ctx context.Context) (_node *Wallet, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   wallet.Table,
			Columns: wallet.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: wallet.FieldID,
			},
		},
	}
	id, ok := wuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "Wallet.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := wuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, wallet.FieldID)
		for _, f := range fields {
			if !wallet.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != wallet.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := wuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := wuo.mutation.Seed(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: wallet.FieldSeed,
		})
	}
	if value, ok := wuo.mutation.Representative(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: wallet.FieldRepresentative,
		})
	}
	if wuo.mutation.RepresentativeCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Column: wallet.FieldRepresentative,
		})
	}
	if value, ok := wuo.mutation.Encrypted(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeBool,
			Value:  value,
			Column: wallet.FieldEncrypted,
		})
	}
	if value, ok := wuo.mutation.Work(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeBool,
			Value:  value,
			Column: wallet.FieldWork,
		})
	}
	if wuo.mutation.AccountsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   wallet.AccountsTable,
			Columns: []string{wallet.AccountsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: account.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := wuo.mutation.RemovedAccountsIDs(); len(nodes) > 0 && !wuo.mutation.AccountsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   wallet.AccountsTable,
			Columns: []string{wallet.AccountsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: account.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := wuo.mutation.AccountsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   wallet.AccountsTable,
			Columns: []string{wallet.AccountsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: account.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &Wallet{config: wuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, wuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{wallet.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	return _node, nil
}
