// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/NpoolPlatform/authing-gateway/pkg/db/ent/authhistory"
	"github.com/google/uuid"
)

// AuthHistoryCreate is the builder for creating a AuthHistory entity.
type AuthHistoryCreate struct {
	config
	mutation *AuthHistoryMutation
	hooks    []Hook
	conflict []sql.ConflictOption
}

// SetAppID sets the "app_id" field.
func (ahc *AuthHistoryCreate) SetAppID(u uuid.UUID) *AuthHistoryCreate {
	ahc.mutation.SetAppID(u)
	return ahc
}

// SetUserID sets the "user_id" field.
func (ahc *AuthHistoryCreate) SetUserID(u uuid.UUID) *AuthHistoryCreate {
	ahc.mutation.SetUserID(u)
	return ahc
}

// SetNillableUserID sets the "user_id" field if the given value is not nil.
func (ahc *AuthHistoryCreate) SetNillableUserID(u *uuid.UUID) *AuthHistoryCreate {
	if u != nil {
		ahc.SetUserID(*u)
	}
	return ahc
}

// SetResource sets the "resource" field.
func (ahc *AuthHistoryCreate) SetResource(s string) *AuthHistoryCreate {
	ahc.mutation.SetResource(s)
	return ahc
}

// SetMethod sets the "method" field.
func (ahc *AuthHistoryCreate) SetMethod(s string) *AuthHistoryCreate {
	ahc.mutation.SetMethod(s)
	return ahc
}

// SetAllowed sets the "allowed" field.
func (ahc *AuthHistoryCreate) SetAllowed(b bool) *AuthHistoryCreate {
	ahc.mutation.SetAllowed(b)
	return ahc
}

// SetCreateAt sets the "create_at" field.
func (ahc *AuthHistoryCreate) SetCreateAt(u uint32) *AuthHistoryCreate {
	ahc.mutation.SetCreateAt(u)
	return ahc
}

// SetNillableCreateAt sets the "create_at" field if the given value is not nil.
func (ahc *AuthHistoryCreate) SetNillableCreateAt(u *uint32) *AuthHistoryCreate {
	if u != nil {
		ahc.SetCreateAt(*u)
	}
	return ahc
}

// SetID sets the "id" field.
func (ahc *AuthHistoryCreate) SetID(u uuid.UUID) *AuthHistoryCreate {
	ahc.mutation.SetID(u)
	return ahc
}

// SetNillableID sets the "id" field if the given value is not nil.
func (ahc *AuthHistoryCreate) SetNillableID(u *uuid.UUID) *AuthHistoryCreate {
	if u != nil {
		ahc.SetID(*u)
	}
	return ahc
}

// Mutation returns the AuthHistoryMutation object of the builder.
func (ahc *AuthHistoryCreate) Mutation() *AuthHistoryMutation {
	return ahc.mutation
}

// Save creates the AuthHistory in the database.
func (ahc *AuthHistoryCreate) Save(ctx context.Context) (*AuthHistory, error) {
	var (
		err  error
		node *AuthHistory
	)
	ahc.defaults()
	if len(ahc.hooks) == 0 {
		if err = ahc.check(); err != nil {
			return nil, err
		}
		node, err = ahc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*AuthHistoryMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = ahc.check(); err != nil {
				return nil, err
			}
			ahc.mutation = mutation
			if node, err = ahc.sqlSave(ctx); err != nil {
				return nil, err
			}
			mutation.id = &node.ID
			mutation.done = true
			return node, err
		})
		for i := len(ahc.hooks) - 1; i >= 0; i-- {
			if ahc.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = ahc.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, ahc.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (ahc *AuthHistoryCreate) SaveX(ctx context.Context) *AuthHistory {
	v, err := ahc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (ahc *AuthHistoryCreate) Exec(ctx context.Context) error {
	_, err := ahc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ahc *AuthHistoryCreate) ExecX(ctx context.Context) {
	if err := ahc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (ahc *AuthHistoryCreate) defaults() {
	if _, ok := ahc.mutation.CreateAt(); !ok {
		v := authhistory.DefaultCreateAt()
		ahc.mutation.SetCreateAt(v)
	}
	if _, ok := ahc.mutation.ID(); !ok {
		v := authhistory.DefaultID()
		ahc.mutation.SetID(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (ahc *AuthHistoryCreate) check() error {
	if _, ok := ahc.mutation.AppID(); !ok {
		return &ValidationError{Name: "app_id", err: errors.New(`ent: missing required field "AuthHistory.app_id"`)}
	}
	if _, ok := ahc.mutation.Resource(); !ok {
		return &ValidationError{Name: "resource", err: errors.New(`ent: missing required field "AuthHistory.resource"`)}
	}
	if _, ok := ahc.mutation.Method(); !ok {
		return &ValidationError{Name: "method", err: errors.New(`ent: missing required field "AuthHistory.method"`)}
	}
	if _, ok := ahc.mutation.Allowed(); !ok {
		return &ValidationError{Name: "allowed", err: errors.New(`ent: missing required field "AuthHistory.allowed"`)}
	}
	if _, ok := ahc.mutation.CreateAt(); !ok {
		return &ValidationError{Name: "create_at", err: errors.New(`ent: missing required field "AuthHistory.create_at"`)}
	}
	return nil
}

func (ahc *AuthHistoryCreate) sqlSave(ctx context.Context) (*AuthHistory, error) {
	_node, _spec := ahc.createSpec()
	if err := sqlgraph.CreateNode(ctx, ahc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return nil, err
	}
	if _spec.ID.Value != nil {
		if id, ok := _spec.ID.Value.(*uuid.UUID); ok {
			_node.ID = *id
		} else if err := _node.ID.Scan(_spec.ID.Value); err != nil {
			return nil, err
		}
	}
	return _node, nil
}

func (ahc *AuthHistoryCreate) createSpec() (*AuthHistory, *sqlgraph.CreateSpec) {
	var (
		_node = &AuthHistory{config: ahc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: authhistory.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: authhistory.FieldID,
			},
		}
	)
	_spec.OnConflict = ahc.conflict
	if id, ok := ahc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = &id
	}
	if value, ok := ahc.mutation.AppID(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Value:  value,
			Column: authhistory.FieldAppID,
		})
		_node.AppID = value
	}
	if value, ok := ahc.mutation.UserID(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Value:  value,
			Column: authhistory.FieldUserID,
		})
		_node.UserID = value
	}
	if value, ok := ahc.mutation.Resource(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: authhistory.FieldResource,
		})
		_node.Resource = value
	}
	if value, ok := ahc.mutation.Method(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: authhistory.FieldMethod,
		})
		_node.Method = value
	}
	if value, ok := ahc.mutation.Allowed(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeBool,
			Value:  value,
			Column: authhistory.FieldAllowed,
		})
		_node.Allowed = value
	}
	if value, ok := ahc.mutation.CreateAt(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: authhistory.FieldCreateAt,
		})
		_node.CreateAt = value
	}
	return _node, _spec
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.AuthHistory.Create().
//		SetAppID(v).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.AuthHistoryUpsert) {
//			SetAppID(v+v).
//		}).
//		Exec(ctx)
//
func (ahc *AuthHistoryCreate) OnConflict(opts ...sql.ConflictOption) *AuthHistoryUpsertOne {
	ahc.conflict = opts
	return &AuthHistoryUpsertOne{
		create: ahc,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.AuthHistory.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
//
func (ahc *AuthHistoryCreate) OnConflictColumns(columns ...string) *AuthHistoryUpsertOne {
	ahc.conflict = append(ahc.conflict, sql.ConflictColumns(columns...))
	return &AuthHistoryUpsertOne{
		create: ahc,
	}
}

type (
	// AuthHistoryUpsertOne is the builder for "upsert"-ing
	//  one AuthHistory node.
	AuthHistoryUpsertOne struct {
		create *AuthHistoryCreate
	}

	// AuthHistoryUpsert is the "OnConflict" setter.
	AuthHistoryUpsert struct {
		*sql.UpdateSet
	}
)

// SetAppID sets the "app_id" field.
func (u *AuthHistoryUpsert) SetAppID(v uuid.UUID) *AuthHistoryUpsert {
	u.Set(authhistory.FieldAppID, v)
	return u
}

// UpdateAppID sets the "app_id" field to the value that was provided on create.
func (u *AuthHistoryUpsert) UpdateAppID() *AuthHistoryUpsert {
	u.SetExcluded(authhistory.FieldAppID)
	return u
}

// SetUserID sets the "user_id" field.
func (u *AuthHistoryUpsert) SetUserID(v uuid.UUID) *AuthHistoryUpsert {
	u.Set(authhistory.FieldUserID, v)
	return u
}

// UpdateUserID sets the "user_id" field to the value that was provided on create.
func (u *AuthHistoryUpsert) UpdateUserID() *AuthHistoryUpsert {
	u.SetExcluded(authhistory.FieldUserID)
	return u
}

// ClearUserID clears the value of the "user_id" field.
func (u *AuthHistoryUpsert) ClearUserID() *AuthHistoryUpsert {
	u.SetNull(authhistory.FieldUserID)
	return u
}

// SetResource sets the "resource" field.
func (u *AuthHistoryUpsert) SetResource(v string) *AuthHistoryUpsert {
	u.Set(authhistory.FieldResource, v)
	return u
}

// UpdateResource sets the "resource" field to the value that was provided on create.
func (u *AuthHistoryUpsert) UpdateResource() *AuthHistoryUpsert {
	u.SetExcluded(authhistory.FieldResource)
	return u
}

// SetMethod sets the "method" field.
func (u *AuthHistoryUpsert) SetMethod(v string) *AuthHistoryUpsert {
	u.Set(authhistory.FieldMethod, v)
	return u
}

// UpdateMethod sets the "method" field to the value that was provided on create.
func (u *AuthHistoryUpsert) UpdateMethod() *AuthHistoryUpsert {
	u.SetExcluded(authhistory.FieldMethod)
	return u
}

// SetAllowed sets the "allowed" field.
func (u *AuthHistoryUpsert) SetAllowed(v bool) *AuthHistoryUpsert {
	u.Set(authhistory.FieldAllowed, v)
	return u
}

// UpdateAllowed sets the "allowed" field to the value that was provided on create.
func (u *AuthHistoryUpsert) UpdateAllowed() *AuthHistoryUpsert {
	u.SetExcluded(authhistory.FieldAllowed)
	return u
}

// SetCreateAt sets the "create_at" field.
func (u *AuthHistoryUpsert) SetCreateAt(v uint32) *AuthHistoryUpsert {
	u.Set(authhistory.FieldCreateAt, v)
	return u
}

// UpdateCreateAt sets the "create_at" field to the value that was provided on create.
func (u *AuthHistoryUpsert) UpdateCreateAt() *AuthHistoryUpsert {
	u.SetExcluded(authhistory.FieldCreateAt)
	return u
}

// AddCreateAt adds v to the "create_at" field.
func (u *AuthHistoryUpsert) AddCreateAt(v uint32) *AuthHistoryUpsert {
	u.Add(authhistory.FieldCreateAt, v)
	return u
}

// UpdateNewValues updates the mutable fields using the new values that were set on create except the ID field.
// Using this option is equivalent to using:
//
//	client.AuthHistory.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//			sql.ResolveWith(func(u *sql.UpdateSet) {
//				u.SetIgnore(authhistory.FieldID)
//			}),
//		).
//		Exec(ctx)
//
func (u *AuthHistoryUpsertOne) UpdateNewValues() *AuthHistoryUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(s *sql.UpdateSet) {
		if _, exists := u.create.mutation.ID(); exists {
			s.SetIgnore(authhistory.FieldID)
		}
	}))
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//  client.AuthHistory.Create().
//      OnConflict(sql.ResolveWithIgnore()).
//      Exec(ctx)
//
func (u *AuthHistoryUpsertOne) Ignore() *AuthHistoryUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *AuthHistoryUpsertOne) DoNothing() *AuthHistoryUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the AuthHistoryCreate.OnConflict
// documentation for more info.
func (u *AuthHistoryUpsertOne) Update(set func(*AuthHistoryUpsert)) *AuthHistoryUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&AuthHistoryUpsert{UpdateSet: update})
	}))
	return u
}

// SetAppID sets the "app_id" field.
func (u *AuthHistoryUpsertOne) SetAppID(v uuid.UUID) *AuthHistoryUpsertOne {
	return u.Update(func(s *AuthHistoryUpsert) {
		s.SetAppID(v)
	})
}

// UpdateAppID sets the "app_id" field to the value that was provided on create.
func (u *AuthHistoryUpsertOne) UpdateAppID() *AuthHistoryUpsertOne {
	return u.Update(func(s *AuthHistoryUpsert) {
		s.UpdateAppID()
	})
}

// SetUserID sets the "user_id" field.
func (u *AuthHistoryUpsertOne) SetUserID(v uuid.UUID) *AuthHistoryUpsertOne {
	return u.Update(func(s *AuthHistoryUpsert) {
		s.SetUserID(v)
	})
}

// UpdateUserID sets the "user_id" field to the value that was provided on create.
func (u *AuthHistoryUpsertOne) UpdateUserID() *AuthHistoryUpsertOne {
	return u.Update(func(s *AuthHistoryUpsert) {
		s.UpdateUserID()
	})
}

// ClearUserID clears the value of the "user_id" field.
func (u *AuthHistoryUpsertOne) ClearUserID() *AuthHistoryUpsertOne {
	return u.Update(func(s *AuthHistoryUpsert) {
		s.ClearUserID()
	})
}

// SetResource sets the "resource" field.
func (u *AuthHistoryUpsertOne) SetResource(v string) *AuthHistoryUpsertOne {
	return u.Update(func(s *AuthHistoryUpsert) {
		s.SetResource(v)
	})
}

// UpdateResource sets the "resource" field to the value that was provided on create.
func (u *AuthHistoryUpsertOne) UpdateResource() *AuthHistoryUpsertOne {
	return u.Update(func(s *AuthHistoryUpsert) {
		s.UpdateResource()
	})
}

// SetMethod sets the "method" field.
func (u *AuthHistoryUpsertOne) SetMethod(v string) *AuthHistoryUpsertOne {
	return u.Update(func(s *AuthHistoryUpsert) {
		s.SetMethod(v)
	})
}

// UpdateMethod sets the "method" field to the value that was provided on create.
func (u *AuthHistoryUpsertOne) UpdateMethod() *AuthHistoryUpsertOne {
	return u.Update(func(s *AuthHistoryUpsert) {
		s.UpdateMethod()
	})
}

// SetAllowed sets the "allowed" field.
func (u *AuthHistoryUpsertOne) SetAllowed(v bool) *AuthHistoryUpsertOne {
	return u.Update(func(s *AuthHistoryUpsert) {
		s.SetAllowed(v)
	})
}

// UpdateAllowed sets the "allowed" field to the value that was provided on create.
func (u *AuthHistoryUpsertOne) UpdateAllowed() *AuthHistoryUpsertOne {
	return u.Update(func(s *AuthHistoryUpsert) {
		s.UpdateAllowed()
	})
}

// SetCreateAt sets the "create_at" field.
func (u *AuthHistoryUpsertOne) SetCreateAt(v uint32) *AuthHistoryUpsertOne {
	return u.Update(func(s *AuthHistoryUpsert) {
		s.SetCreateAt(v)
	})
}

// AddCreateAt adds v to the "create_at" field.
func (u *AuthHistoryUpsertOne) AddCreateAt(v uint32) *AuthHistoryUpsertOne {
	return u.Update(func(s *AuthHistoryUpsert) {
		s.AddCreateAt(v)
	})
}

// UpdateCreateAt sets the "create_at" field to the value that was provided on create.
func (u *AuthHistoryUpsertOne) UpdateCreateAt() *AuthHistoryUpsertOne {
	return u.Update(func(s *AuthHistoryUpsert) {
		s.UpdateCreateAt()
	})
}

// Exec executes the query.
func (u *AuthHistoryUpsertOne) Exec(ctx context.Context) error {
	if len(u.create.conflict) == 0 {
		return errors.New("ent: missing options for AuthHistoryCreate.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *AuthHistoryUpsertOne) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}

// Exec executes the UPSERT query and returns the inserted/updated ID.
func (u *AuthHistoryUpsertOne) ID(ctx context.Context) (id uuid.UUID, err error) {
	if u.create.driver.Dialect() == dialect.MySQL {
		// In case of "ON CONFLICT", there is no way to get back non-numeric ID
		// fields from the database since MySQL does not support the RETURNING clause.
		return id, errors.New("ent: AuthHistoryUpsertOne.ID is not supported by MySQL driver. Use AuthHistoryUpsertOne.Exec instead")
	}
	node, err := u.create.Save(ctx)
	if err != nil {
		return id, err
	}
	return node.ID, nil
}

// IDX is like ID, but panics if an error occurs.
func (u *AuthHistoryUpsertOne) IDX(ctx context.Context) uuid.UUID {
	id, err := u.ID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// AuthHistoryCreateBulk is the builder for creating many AuthHistory entities in bulk.
type AuthHistoryCreateBulk struct {
	config
	builders []*AuthHistoryCreate
	conflict []sql.ConflictOption
}

// Save creates the AuthHistory entities in the database.
func (ahcb *AuthHistoryCreateBulk) Save(ctx context.Context) ([]*AuthHistory, error) {
	specs := make([]*sqlgraph.CreateSpec, len(ahcb.builders))
	nodes := make([]*AuthHistory, len(ahcb.builders))
	mutators := make([]Mutator, len(ahcb.builders))
	for i := range ahcb.builders {
		func(i int, root context.Context) {
			builder := ahcb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*AuthHistoryMutation)
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
					_, err = mutators[i+1].Mutate(root, ahcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					spec.OnConflict = ahcb.conflict
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, ahcb.driver, spec); err != nil {
						if sqlgraph.IsConstraintError(err) {
							err = &ConstraintError{err.Error(), err}
						}
					}
				}
				if err != nil {
					return nil, err
				}
				mutation.id = &nodes[i].ID
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
		if _, err := mutators[0].Mutate(ctx, ahcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (ahcb *AuthHistoryCreateBulk) SaveX(ctx context.Context) []*AuthHistory {
	v, err := ahcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (ahcb *AuthHistoryCreateBulk) Exec(ctx context.Context) error {
	_, err := ahcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ahcb *AuthHistoryCreateBulk) ExecX(ctx context.Context) {
	if err := ahcb.Exec(ctx); err != nil {
		panic(err)
	}
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.AuthHistory.CreateBulk(builders...).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.AuthHistoryUpsert) {
//			SetAppID(v+v).
//		}).
//		Exec(ctx)
//
func (ahcb *AuthHistoryCreateBulk) OnConflict(opts ...sql.ConflictOption) *AuthHistoryUpsertBulk {
	ahcb.conflict = opts
	return &AuthHistoryUpsertBulk{
		create: ahcb,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.AuthHistory.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
//
func (ahcb *AuthHistoryCreateBulk) OnConflictColumns(columns ...string) *AuthHistoryUpsertBulk {
	ahcb.conflict = append(ahcb.conflict, sql.ConflictColumns(columns...))
	return &AuthHistoryUpsertBulk{
		create: ahcb,
	}
}

// AuthHistoryUpsertBulk is the builder for "upsert"-ing
// a bulk of AuthHistory nodes.
type AuthHistoryUpsertBulk struct {
	create *AuthHistoryCreateBulk
}

// UpdateNewValues updates the mutable fields using the new values that
// were set on create. Using this option is equivalent to using:
//
//	client.AuthHistory.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//			sql.ResolveWith(func(u *sql.UpdateSet) {
//				u.SetIgnore(authhistory.FieldID)
//			}),
//		).
//		Exec(ctx)
//
func (u *AuthHistoryUpsertBulk) UpdateNewValues() *AuthHistoryUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(s *sql.UpdateSet) {
		for _, b := range u.create.builders {
			if _, exists := b.mutation.ID(); exists {
				s.SetIgnore(authhistory.FieldID)
				return
			}
		}
	}))
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//	client.AuthHistory.Create().
//		OnConflict(sql.ResolveWithIgnore()).
//		Exec(ctx)
//
func (u *AuthHistoryUpsertBulk) Ignore() *AuthHistoryUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *AuthHistoryUpsertBulk) DoNothing() *AuthHistoryUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the AuthHistoryCreateBulk.OnConflict
// documentation for more info.
func (u *AuthHistoryUpsertBulk) Update(set func(*AuthHistoryUpsert)) *AuthHistoryUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&AuthHistoryUpsert{UpdateSet: update})
	}))
	return u
}

// SetAppID sets the "app_id" field.
func (u *AuthHistoryUpsertBulk) SetAppID(v uuid.UUID) *AuthHistoryUpsertBulk {
	return u.Update(func(s *AuthHistoryUpsert) {
		s.SetAppID(v)
	})
}

// UpdateAppID sets the "app_id" field to the value that was provided on create.
func (u *AuthHistoryUpsertBulk) UpdateAppID() *AuthHistoryUpsertBulk {
	return u.Update(func(s *AuthHistoryUpsert) {
		s.UpdateAppID()
	})
}

// SetUserID sets the "user_id" field.
func (u *AuthHistoryUpsertBulk) SetUserID(v uuid.UUID) *AuthHistoryUpsertBulk {
	return u.Update(func(s *AuthHistoryUpsert) {
		s.SetUserID(v)
	})
}

// UpdateUserID sets the "user_id" field to the value that was provided on create.
func (u *AuthHistoryUpsertBulk) UpdateUserID() *AuthHistoryUpsertBulk {
	return u.Update(func(s *AuthHistoryUpsert) {
		s.UpdateUserID()
	})
}

// ClearUserID clears the value of the "user_id" field.
func (u *AuthHistoryUpsertBulk) ClearUserID() *AuthHistoryUpsertBulk {
	return u.Update(func(s *AuthHistoryUpsert) {
		s.ClearUserID()
	})
}

// SetResource sets the "resource" field.
func (u *AuthHistoryUpsertBulk) SetResource(v string) *AuthHistoryUpsertBulk {
	return u.Update(func(s *AuthHistoryUpsert) {
		s.SetResource(v)
	})
}

// UpdateResource sets the "resource" field to the value that was provided on create.
func (u *AuthHistoryUpsertBulk) UpdateResource() *AuthHistoryUpsertBulk {
	return u.Update(func(s *AuthHistoryUpsert) {
		s.UpdateResource()
	})
}

// SetMethod sets the "method" field.
func (u *AuthHistoryUpsertBulk) SetMethod(v string) *AuthHistoryUpsertBulk {
	return u.Update(func(s *AuthHistoryUpsert) {
		s.SetMethod(v)
	})
}

// UpdateMethod sets the "method" field to the value that was provided on create.
func (u *AuthHistoryUpsertBulk) UpdateMethod() *AuthHistoryUpsertBulk {
	return u.Update(func(s *AuthHistoryUpsert) {
		s.UpdateMethod()
	})
}

// SetAllowed sets the "allowed" field.
func (u *AuthHistoryUpsertBulk) SetAllowed(v bool) *AuthHistoryUpsertBulk {
	return u.Update(func(s *AuthHistoryUpsert) {
		s.SetAllowed(v)
	})
}

// UpdateAllowed sets the "allowed" field to the value that was provided on create.
func (u *AuthHistoryUpsertBulk) UpdateAllowed() *AuthHistoryUpsertBulk {
	return u.Update(func(s *AuthHistoryUpsert) {
		s.UpdateAllowed()
	})
}

// SetCreateAt sets the "create_at" field.
func (u *AuthHistoryUpsertBulk) SetCreateAt(v uint32) *AuthHistoryUpsertBulk {
	return u.Update(func(s *AuthHistoryUpsert) {
		s.SetCreateAt(v)
	})
}

// AddCreateAt adds v to the "create_at" field.
func (u *AuthHistoryUpsertBulk) AddCreateAt(v uint32) *AuthHistoryUpsertBulk {
	return u.Update(func(s *AuthHistoryUpsert) {
		s.AddCreateAt(v)
	})
}

// UpdateCreateAt sets the "create_at" field to the value that was provided on create.
func (u *AuthHistoryUpsertBulk) UpdateCreateAt() *AuthHistoryUpsertBulk {
	return u.Update(func(s *AuthHistoryUpsert) {
		s.UpdateCreateAt()
	})
}

// Exec executes the query.
func (u *AuthHistoryUpsertBulk) Exec(ctx context.Context) error {
	for i, b := range u.create.builders {
		if len(b.conflict) != 0 {
			return fmt.Errorf("ent: OnConflict was set for builder %d. Set it on the AuthHistoryCreateBulk instead", i)
		}
	}
	if len(u.create.conflict) == 0 {
		return errors.New("ent: missing options for AuthHistoryCreateBulk.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *AuthHistoryUpsertBulk) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}
