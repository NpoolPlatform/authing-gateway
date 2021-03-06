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
	"github.com/NpoolPlatform/authing-gateway/pkg/db/ent/appuserauth"
	"github.com/google/uuid"
)

// AppUserAuthCreate is the builder for creating a AppUserAuth entity.
type AppUserAuthCreate struct {
	config
	mutation *AppUserAuthMutation
	hooks    []Hook
	conflict []sql.ConflictOption
}

// SetAppID sets the "app_id" field.
func (auac *AppUserAuthCreate) SetAppID(u uuid.UUID) *AppUserAuthCreate {
	auac.mutation.SetAppID(u)
	return auac
}

// SetUserID sets the "user_id" field.
func (auac *AppUserAuthCreate) SetUserID(u uuid.UUID) *AppUserAuthCreate {
	auac.mutation.SetUserID(u)
	return auac
}

// SetResource sets the "resource" field.
func (auac *AppUserAuthCreate) SetResource(s string) *AppUserAuthCreate {
	auac.mutation.SetResource(s)
	return auac
}

// SetMethod sets the "method" field.
func (auac *AppUserAuthCreate) SetMethod(s string) *AppUserAuthCreate {
	auac.mutation.SetMethod(s)
	return auac
}

// SetCreateAt sets the "create_at" field.
func (auac *AppUserAuthCreate) SetCreateAt(u uint32) *AppUserAuthCreate {
	auac.mutation.SetCreateAt(u)
	return auac
}

// SetNillableCreateAt sets the "create_at" field if the given value is not nil.
func (auac *AppUserAuthCreate) SetNillableCreateAt(u *uint32) *AppUserAuthCreate {
	if u != nil {
		auac.SetCreateAt(*u)
	}
	return auac
}

// SetUpdateAt sets the "update_at" field.
func (auac *AppUserAuthCreate) SetUpdateAt(u uint32) *AppUserAuthCreate {
	auac.mutation.SetUpdateAt(u)
	return auac
}

// SetNillableUpdateAt sets the "update_at" field if the given value is not nil.
func (auac *AppUserAuthCreate) SetNillableUpdateAt(u *uint32) *AppUserAuthCreate {
	if u != nil {
		auac.SetUpdateAt(*u)
	}
	return auac
}

// SetDeleteAt sets the "delete_at" field.
func (auac *AppUserAuthCreate) SetDeleteAt(u uint32) *AppUserAuthCreate {
	auac.mutation.SetDeleteAt(u)
	return auac
}

// SetNillableDeleteAt sets the "delete_at" field if the given value is not nil.
func (auac *AppUserAuthCreate) SetNillableDeleteAt(u *uint32) *AppUserAuthCreate {
	if u != nil {
		auac.SetDeleteAt(*u)
	}
	return auac
}

// SetID sets the "id" field.
func (auac *AppUserAuthCreate) SetID(u uuid.UUID) *AppUserAuthCreate {
	auac.mutation.SetID(u)
	return auac
}

// SetNillableID sets the "id" field if the given value is not nil.
func (auac *AppUserAuthCreate) SetNillableID(u *uuid.UUID) *AppUserAuthCreate {
	if u != nil {
		auac.SetID(*u)
	}
	return auac
}

// Mutation returns the AppUserAuthMutation object of the builder.
func (auac *AppUserAuthCreate) Mutation() *AppUserAuthMutation {
	return auac.mutation
}

// Save creates the AppUserAuth in the database.
func (auac *AppUserAuthCreate) Save(ctx context.Context) (*AppUserAuth, error) {
	var (
		err  error
		node *AppUserAuth
	)
	auac.defaults()
	if len(auac.hooks) == 0 {
		if err = auac.check(); err != nil {
			return nil, err
		}
		node, err = auac.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*AppUserAuthMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = auac.check(); err != nil {
				return nil, err
			}
			auac.mutation = mutation
			if node, err = auac.sqlSave(ctx); err != nil {
				return nil, err
			}
			mutation.id = &node.ID
			mutation.done = true
			return node, err
		})
		for i := len(auac.hooks) - 1; i >= 0; i-- {
			if auac.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = auac.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, auac.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (auac *AppUserAuthCreate) SaveX(ctx context.Context) *AppUserAuth {
	v, err := auac.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (auac *AppUserAuthCreate) Exec(ctx context.Context) error {
	_, err := auac.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (auac *AppUserAuthCreate) ExecX(ctx context.Context) {
	if err := auac.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (auac *AppUserAuthCreate) defaults() {
	if _, ok := auac.mutation.CreateAt(); !ok {
		v := appuserauth.DefaultCreateAt()
		auac.mutation.SetCreateAt(v)
	}
	if _, ok := auac.mutation.UpdateAt(); !ok {
		v := appuserauth.DefaultUpdateAt()
		auac.mutation.SetUpdateAt(v)
	}
	if _, ok := auac.mutation.DeleteAt(); !ok {
		v := appuserauth.DefaultDeleteAt()
		auac.mutation.SetDeleteAt(v)
	}
	if _, ok := auac.mutation.ID(); !ok {
		v := appuserauth.DefaultID()
		auac.mutation.SetID(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (auac *AppUserAuthCreate) check() error {
	if _, ok := auac.mutation.AppID(); !ok {
		return &ValidationError{Name: "app_id", err: errors.New(`ent: missing required field "AppUserAuth.app_id"`)}
	}
	if _, ok := auac.mutation.UserID(); !ok {
		return &ValidationError{Name: "user_id", err: errors.New(`ent: missing required field "AppUserAuth.user_id"`)}
	}
	if _, ok := auac.mutation.Resource(); !ok {
		return &ValidationError{Name: "resource", err: errors.New(`ent: missing required field "AppUserAuth.resource"`)}
	}
	if _, ok := auac.mutation.Method(); !ok {
		return &ValidationError{Name: "method", err: errors.New(`ent: missing required field "AppUserAuth.method"`)}
	}
	if _, ok := auac.mutation.CreateAt(); !ok {
		return &ValidationError{Name: "create_at", err: errors.New(`ent: missing required field "AppUserAuth.create_at"`)}
	}
	if _, ok := auac.mutation.UpdateAt(); !ok {
		return &ValidationError{Name: "update_at", err: errors.New(`ent: missing required field "AppUserAuth.update_at"`)}
	}
	if _, ok := auac.mutation.DeleteAt(); !ok {
		return &ValidationError{Name: "delete_at", err: errors.New(`ent: missing required field "AppUserAuth.delete_at"`)}
	}
	return nil
}

func (auac *AppUserAuthCreate) sqlSave(ctx context.Context) (*AppUserAuth, error) {
	_node, _spec := auac.createSpec()
	if err := sqlgraph.CreateNode(ctx, auac.driver, _spec); err != nil {
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

func (auac *AppUserAuthCreate) createSpec() (*AppUserAuth, *sqlgraph.CreateSpec) {
	var (
		_node = &AppUserAuth{config: auac.config}
		_spec = &sqlgraph.CreateSpec{
			Table: appuserauth.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: appuserauth.FieldID,
			},
		}
	)
	_spec.OnConflict = auac.conflict
	if id, ok := auac.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = &id
	}
	if value, ok := auac.mutation.AppID(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Value:  value,
			Column: appuserauth.FieldAppID,
		})
		_node.AppID = value
	}
	if value, ok := auac.mutation.UserID(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeUUID,
			Value:  value,
			Column: appuserauth.FieldUserID,
		})
		_node.UserID = value
	}
	if value, ok := auac.mutation.Resource(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: appuserauth.FieldResource,
		})
		_node.Resource = value
	}
	if value, ok := auac.mutation.Method(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: appuserauth.FieldMethod,
		})
		_node.Method = value
	}
	if value, ok := auac.mutation.CreateAt(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: appuserauth.FieldCreateAt,
		})
		_node.CreateAt = value
	}
	if value, ok := auac.mutation.UpdateAt(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: appuserauth.FieldUpdateAt,
		})
		_node.UpdateAt = value
	}
	if value, ok := auac.mutation.DeleteAt(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeUint32,
			Value:  value,
			Column: appuserauth.FieldDeleteAt,
		})
		_node.DeleteAt = value
	}
	return _node, _spec
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.AppUserAuth.Create().
//		SetAppID(v).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.AppUserAuthUpsert) {
//			SetAppID(v+v).
//		}).
//		Exec(ctx)
//
func (auac *AppUserAuthCreate) OnConflict(opts ...sql.ConflictOption) *AppUserAuthUpsertOne {
	auac.conflict = opts
	return &AppUserAuthUpsertOne{
		create: auac,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.AppUserAuth.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
//
func (auac *AppUserAuthCreate) OnConflictColumns(columns ...string) *AppUserAuthUpsertOne {
	auac.conflict = append(auac.conflict, sql.ConflictColumns(columns...))
	return &AppUserAuthUpsertOne{
		create: auac,
	}
}

type (
	// AppUserAuthUpsertOne is the builder for "upsert"-ing
	//  one AppUserAuth node.
	AppUserAuthUpsertOne struct {
		create *AppUserAuthCreate
	}

	// AppUserAuthUpsert is the "OnConflict" setter.
	AppUserAuthUpsert struct {
		*sql.UpdateSet
	}
)

// SetAppID sets the "app_id" field.
func (u *AppUserAuthUpsert) SetAppID(v uuid.UUID) *AppUserAuthUpsert {
	u.Set(appuserauth.FieldAppID, v)
	return u
}

// UpdateAppID sets the "app_id" field to the value that was provided on create.
func (u *AppUserAuthUpsert) UpdateAppID() *AppUserAuthUpsert {
	u.SetExcluded(appuserauth.FieldAppID)
	return u
}

// SetUserID sets the "user_id" field.
func (u *AppUserAuthUpsert) SetUserID(v uuid.UUID) *AppUserAuthUpsert {
	u.Set(appuserauth.FieldUserID, v)
	return u
}

// UpdateUserID sets the "user_id" field to the value that was provided on create.
func (u *AppUserAuthUpsert) UpdateUserID() *AppUserAuthUpsert {
	u.SetExcluded(appuserauth.FieldUserID)
	return u
}

// SetResource sets the "resource" field.
func (u *AppUserAuthUpsert) SetResource(v string) *AppUserAuthUpsert {
	u.Set(appuserauth.FieldResource, v)
	return u
}

// UpdateResource sets the "resource" field to the value that was provided on create.
func (u *AppUserAuthUpsert) UpdateResource() *AppUserAuthUpsert {
	u.SetExcluded(appuserauth.FieldResource)
	return u
}

// SetMethod sets the "method" field.
func (u *AppUserAuthUpsert) SetMethod(v string) *AppUserAuthUpsert {
	u.Set(appuserauth.FieldMethod, v)
	return u
}

// UpdateMethod sets the "method" field to the value that was provided on create.
func (u *AppUserAuthUpsert) UpdateMethod() *AppUserAuthUpsert {
	u.SetExcluded(appuserauth.FieldMethod)
	return u
}

// SetCreateAt sets the "create_at" field.
func (u *AppUserAuthUpsert) SetCreateAt(v uint32) *AppUserAuthUpsert {
	u.Set(appuserauth.FieldCreateAt, v)
	return u
}

// UpdateCreateAt sets the "create_at" field to the value that was provided on create.
func (u *AppUserAuthUpsert) UpdateCreateAt() *AppUserAuthUpsert {
	u.SetExcluded(appuserauth.FieldCreateAt)
	return u
}

// AddCreateAt adds v to the "create_at" field.
func (u *AppUserAuthUpsert) AddCreateAt(v uint32) *AppUserAuthUpsert {
	u.Add(appuserauth.FieldCreateAt, v)
	return u
}

// SetUpdateAt sets the "update_at" field.
func (u *AppUserAuthUpsert) SetUpdateAt(v uint32) *AppUserAuthUpsert {
	u.Set(appuserauth.FieldUpdateAt, v)
	return u
}

// UpdateUpdateAt sets the "update_at" field to the value that was provided on create.
func (u *AppUserAuthUpsert) UpdateUpdateAt() *AppUserAuthUpsert {
	u.SetExcluded(appuserauth.FieldUpdateAt)
	return u
}

// AddUpdateAt adds v to the "update_at" field.
func (u *AppUserAuthUpsert) AddUpdateAt(v uint32) *AppUserAuthUpsert {
	u.Add(appuserauth.FieldUpdateAt, v)
	return u
}

// SetDeleteAt sets the "delete_at" field.
func (u *AppUserAuthUpsert) SetDeleteAt(v uint32) *AppUserAuthUpsert {
	u.Set(appuserauth.FieldDeleteAt, v)
	return u
}

// UpdateDeleteAt sets the "delete_at" field to the value that was provided on create.
func (u *AppUserAuthUpsert) UpdateDeleteAt() *AppUserAuthUpsert {
	u.SetExcluded(appuserauth.FieldDeleteAt)
	return u
}

// AddDeleteAt adds v to the "delete_at" field.
func (u *AppUserAuthUpsert) AddDeleteAt(v uint32) *AppUserAuthUpsert {
	u.Add(appuserauth.FieldDeleteAt, v)
	return u
}

// UpdateNewValues updates the mutable fields using the new values that were set on create except the ID field.
// Using this option is equivalent to using:
//
//	client.AppUserAuth.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//			sql.ResolveWith(func(u *sql.UpdateSet) {
//				u.SetIgnore(appuserauth.FieldID)
//			}),
//		).
//		Exec(ctx)
//
func (u *AppUserAuthUpsertOne) UpdateNewValues() *AppUserAuthUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(s *sql.UpdateSet) {
		if _, exists := u.create.mutation.ID(); exists {
			s.SetIgnore(appuserauth.FieldID)
		}
	}))
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//  client.AppUserAuth.Create().
//      OnConflict(sql.ResolveWithIgnore()).
//      Exec(ctx)
//
func (u *AppUserAuthUpsertOne) Ignore() *AppUserAuthUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *AppUserAuthUpsertOne) DoNothing() *AppUserAuthUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the AppUserAuthCreate.OnConflict
// documentation for more info.
func (u *AppUserAuthUpsertOne) Update(set func(*AppUserAuthUpsert)) *AppUserAuthUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&AppUserAuthUpsert{UpdateSet: update})
	}))
	return u
}

// SetAppID sets the "app_id" field.
func (u *AppUserAuthUpsertOne) SetAppID(v uuid.UUID) *AppUserAuthUpsertOne {
	return u.Update(func(s *AppUserAuthUpsert) {
		s.SetAppID(v)
	})
}

// UpdateAppID sets the "app_id" field to the value that was provided on create.
func (u *AppUserAuthUpsertOne) UpdateAppID() *AppUserAuthUpsertOne {
	return u.Update(func(s *AppUserAuthUpsert) {
		s.UpdateAppID()
	})
}

// SetUserID sets the "user_id" field.
func (u *AppUserAuthUpsertOne) SetUserID(v uuid.UUID) *AppUserAuthUpsertOne {
	return u.Update(func(s *AppUserAuthUpsert) {
		s.SetUserID(v)
	})
}

// UpdateUserID sets the "user_id" field to the value that was provided on create.
func (u *AppUserAuthUpsertOne) UpdateUserID() *AppUserAuthUpsertOne {
	return u.Update(func(s *AppUserAuthUpsert) {
		s.UpdateUserID()
	})
}

// SetResource sets the "resource" field.
func (u *AppUserAuthUpsertOne) SetResource(v string) *AppUserAuthUpsertOne {
	return u.Update(func(s *AppUserAuthUpsert) {
		s.SetResource(v)
	})
}

// UpdateResource sets the "resource" field to the value that was provided on create.
func (u *AppUserAuthUpsertOne) UpdateResource() *AppUserAuthUpsertOne {
	return u.Update(func(s *AppUserAuthUpsert) {
		s.UpdateResource()
	})
}

// SetMethod sets the "method" field.
func (u *AppUserAuthUpsertOne) SetMethod(v string) *AppUserAuthUpsertOne {
	return u.Update(func(s *AppUserAuthUpsert) {
		s.SetMethod(v)
	})
}

// UpdateMethod sets the "method" field to the value that was provided on create.
func (u *AppUserAuthUpsertOne) UpdateMethod() *AppUserAuthUpsertOne {
	return u.Update(func(s *AppUserAuthUpsert) {
		s.UpdateMethod()
	})
}

// SetCreateAt sets the "create_at" field.
func (u *AppUserAuthUpsertOne) SetCreateAt(v uint32) *AppUserAuthUpsertOne {
	return u.Update(func(s *AppUserAuthUpsert) {
		s.SetCreateAt(v)
	})
}

// AddCreateAt adds v to the "create_at" field.
func (u *AppUserAuthUpsertOne) AddCreateAt(v uint32) *AppUserAuthUpsertOne {
	return u.Update(func(s *AppUserAuthUpsert) {
		s.AddCreateAt(v)
	})
}

// UpdateCreateAt sets the "create_at" field to the value that was provided on create.
func (u *AppUserAuthUpsertOne) UpdateCreateAt() *AppUserAuthUpsertOne {
	return u.Update(func(s *AppUserAuthUpsert) {
		s.UpdateCreateAt()
	})
}

// SetUpdateAt sets the "update_at" field.
func (u *AppUserAuthUpsertOne) SetUpdateAt(v uint32) *AppUserAuthUpsertOne {
	return u.Update(func(s *AppUserAuthUpsert) {
		s.SetUpdateAt(v)
	})
}

// AddUpdateAt adds v to the "update_at" field.
func (u *AppUserAuthUpsertOne) AddUpdateAt(v uint32) *AppUserAuthUpsertOne {
	return u.Update(func(s *AppUserAuthUpsert) {
		s.AddUpdateAt(v)
	})
}

// UpdateUpdateAt sets the "update_at" field to the value that was provided on create.
func (u *AppUserAuthUpsertOne) UpdateUpdateAt() *AppUserAuthUpsertOne {
	return u.Update(func(s *AppUserAuthUpsert) {
		s.UpdateUpdateAt()
	})
}

// SetDeleteAt sets the "delete_at" field.
func (u *AppUserAuthUpsertOne) SetDeleteAt(v uint32) *AppUserAuthUpsertOne {
	return u.Update(func(s *AppUserAuthUpsert) {
		s.SetDeleteAt(v)
	})
}

// AddDeleteAt adds v to the "delete_at" field.
func (u *AppUserAuthUpsertOne) AddDeleteAt(v uint32) *AppUserAuthUpsertOne {
	return u.Update(func(s *AppUserAuthUpsert) {
		s.AddDeleteAt(v)
	})
}

// UpdateDeleteAt sets the "delete_at" field to the value that was provided on create.
func (u *AppUserAuthUpsertOne) UpdateDeleteAt() *AppUserAuthUpsertOne {
	return u.Update(func(s *AppUserAuthUpsert) {
		s.UpdateDeleteAt()
	})
}

// Exec executes the query.
func (u *AppUserAuthUpsertOne) Exec(ctx context.Context) error {
	if len(u.create.conflict) == 0 {
		return errors.New("ent: missing options for AppUserAuthCreate.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *AppUserAuthUpsertOne) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}

// Exec executes the UPSERT query and returns the inserted/updated ID.
func (u *AppUserAuthUpsertOne) ID(ctx context.Context) (id uuid.UUID, err error) {
	if u.create.driver.Dialect() == dialect.MySQL {
		// In case of "ON CONFLICT", there is no way to get back non-numeric ID
		// fields from the database since MySQL does not support the RETURNING clause.
		return id, errors.New("ent: AppUserAuthUpsertOne.ID is not supported by MySQL driver. Use AppUserAuthUpsertOne.Exec instead")
	}
	node, err := u.create.Save(ctx)
	if err != nil {
		return id, err
	}
	return node.ID, nil
}

// IDX is like ID, but panics if an error occurs.
func (u *AppUserAuthUpsertOne) IDX(ctx context.Context) uuid.UUID {
	id, err := u.ID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// AppUserAuthCreateBulk is the builder for creating many AppUserAuth entities in bulk.
type AppUserAuthCreateBulk struct {
	config
	builders []*AppUserAuthCreate
	conflict []sql.ConflictOption
}

// Save creates the AppUserAuth entities in the database.
func (auacb *AppUserAuthCreateBulk) Save(ctx context.Context) ([]*AppUserAuth, error) {
	specs := make([]*sqlgraph.CreateSpec, len(auacb.builders))
	nodes := make([]*AppUserAuth, len(auacb.builders))
	mutators := make([]Mutator, len(auacb.builders))
	for i := range auacb.builders {
		func(i int, root context.Context) {
			builder := auacb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*AppUserAuthMutation)
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
					_, err = mutators[i+1].Mutate(root, auacb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					spec.OnConflict = auacb.conflict
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, auacb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, auacb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (auacb *AppUserAuthCreateBulk) SaveX(ctx context.Context) []*AppUserAuth {
	v, err := auacb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (auacb *AppUserAuthCreateBulk) Exec(ctx context.Context) error {
	_, err := auacb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (auacb *AppUserAuthCreateBulk) ExecX(ctx context.Context) {
	if err := auacb.Exec(ctx); err != nil {
		panic(err)
	}
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.AppUserAuth.CreateBulk(builders...).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.AppUserAuthUpsert) {
//			SetAppID(v+v).
//		}).
//		Exec(ctx)
//
func (auacb *AppUserAuthCreateBulk) OnConflict(opts ...sql.ConflictOption) *AppUserAuthUpsertBulk {
	auacb.conflict = opts
	return &AppUserAuthUpsertBulk{
		create: auacb,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.AppUserAuth.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
//
func (auacb *AppUserAuthCreateBulk) OnConflictColumns(columns ...string) *AppUserAuthUpsertBulk {
	auacb.conflict = append(auacb.conflict, sql.ConflictColumns(columns...))
	return &AppUserAuthUpsertBulk{
		create: auacb,
	}
}

// AppUserAuthUpsertBulk is the builder for "upsert"-ing
// a bulk of AppUserAuth nodes.
type AppUserAuthUpsertBulk struct {
	create *AppUserAuthCreateBulk
}

// UpdateNewValues updates the mutable fields using the new values that
// were set on create. Using this option is equivalent to using:
//
//	client.AppUserAuth.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//			sql.ResolveWith(func(u *sql.UpdateSet) {
//				u.SetIgnore(appuserauth.FieldID)
//			}),
//		).
//		Exec(ctx)
//
func (u *AppUserAuthUpsertBulk) UpdateNewValues() *AppUserAuthUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(s *sql.UpdateSet) {
		for _, b := range u.create.builders {
			if _, exists := b.mutation.ID(); exists {
				s.SetIgnore(appuserauth.FieldID)
				return
			}
		}
	}))
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//	client.AppUserAuth.Create().
//		OnConflict(sql.ResolveWithIgnore()).
//		Exec(ctx)
//
func (u *AppUserAuthUpsertBulk) Ignore() *AppUserAuthUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *AppUserAuthUpsertBulk) DoNothing() *AppUserAuthUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the AppUserAuthCreateBulk.OnConflict
// documentation for more info.
func (u *AppUserAuthUpsertBulk) Update(set func(*AppUserAuthUpsert)) *AppUserAuthUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&AppUserAuthUpsert{UpdateSet: update})
	}))
	return u
}

// SetAppID sets the "app_id" field.
func (u *AppUserAuthUpsertBulk) SetAppID(v uuid.UUID) *AppUserAuthUpsertBulk {
	return u.Update(func(s *AppUserAuthUpsert) {
		s.SetAppID(v)
	})
}

// UpdateAppID sets the "app_id" field to the value that was provided on create.
func (u *AppUserAuthUpsertBulk) UpdateAppID() *AppUserAuthUpsertBulk {
	return u.Update(func(s *AppUserAuthUpsert) {
		s.UpdateAppID()
	})
}

// SetUserID sets the "user_id" field.
func (u *AppUserAuthUpsertBulk) SetUserID(v uuid.UUID) *AppUserAuthUpsertBulk {
	return u.Update(func(s *AppUserAuthUpsert) {
		s.SetUserID(v)
	})
}

// UpdateUserID sets the "user_id" field to the value that was provided on create.
func (u *AppUserAuthUpsertBulk) UpdateUserID() *AppUserAuthUpsertBulk {
	return u.Update(func(s *AppUserAuthUpsert) {
		s.UpdateUserID()
	})
}

// SetResource sets the "resource" field.
func (u *AppUserAuthUpsertBulk) SetResource(v string) *AppUserAuthUpsertBulk {
	return u.Update(func(s *AppUserAuthUpsert) {
		s.SetResource(v)
	})
}

// UpdateResource sets the "resource" field to the value that was provided on create.
func (u *AppUserAuthUpsertBulk) UpdateResource() *AppUserAuthUpsertBulk {
	return u.Update(func(s *AppUserAuthUpsert) {
		s.UpdateResource()
	})
}

// SetMethod sets the "method" field.
func (u *AppUserAuthUpsertBulk) SetMethod(v string) *AppUserAuthUpsertBulk {
	return u.Update(func(s *AppUserAuthUpsert) {
		s.SetMethod(v)
	})
}

// UpdateMethod sets the "method" field to the value that was provided on create.
func (u *AppUserAuthUpsertBulk) UpdateMethod() *AppUserAuthUpsertBulk {
	return u.Update(func(s *AppUserAuthUpsert) {
		s.UpdateMethod()
	})
}

// SetCreateAt sets the "create_at" field.
func (u *AppUserAuthUpsertBulk) SetCreateAt(v uint32) *AppUserAuthUpsertBulk {
	return u.Update(func(s *AppUserAuthUpsert) {
		s.SetCreateAt(v)
	})
}

// AddCreateAt adds v to the "create_at" field.
func (u *AppUserAuthUpsertBulk) AddCreateAt(v uint32) *AppUserAuthUpsertBulk {
	return u.Update(func(s *AppUserAuthUpsert) {
		s.AddCreateAt(v)
	})
}

// UpdateCreateAt sets the "create_at" field to the value that was provided on create.
func (u *AppUserAuthUpsertBulk) UpdateCreateAt() *AppUserAuthUpsertBulk {
	return u.Update(func(s *AppUserAuthUpsert) {
		s.UpdateCreateAt()
	})
}

// SetUpdateAt sets the "update_at" field.
func (u *AppUserAuthUpsertBulk) SetUpdateAt(v uint32) *AppUserAuthUpsertBulk {
	return u.Update(func(s *AppUserAuthUpsert) {
		s.SetUpdateAt(v)
	})
}

// AddUpdateAt adds v to the "update_at" field.
func (u *AppUserAuthUpsertBulk) AddUpdateAt(v uint32) *AppUserAuthUpsertBulk {
	return u.Update(func(s *AppUserAuthUpsert) {
		s.AddUpdateAt(v)
	})
}

// UpdateUpdateAt sets the "update_at" field to the value that was provided on create.
func (u *AppUserAuthUpsertBulk) UpdateUpdateAt() *AppUserAuthUpsertBulk {
	return u.Update(func(s *AppUserAuthUpsert) {
		s.UpdateUpdateAt()
	})
}

// SetDeleteAt sets the "delete_at" field.
func (u *AppUserAuthUpsertBulk) SetDeleteAt(v uint32) *AppUserAuthUpsertBulk {
	return u.Update(func(s *AppUserAuthUpsert) {
		s.SetDeleteAt(v)
	})
}

// AddDeleteAt adds v to the "delete_at" field.
func (u *AppUserAuthUpsertBulk) AddDeleteAt(v uint32) *AppUserAuthUpsertBulk {
	return u.Update(func(s *AppUserAuthUpsert) {
		s.AddDeleteAt(v)
	})
}

// UpdateDeleteAt sets the "delete_at" field to the value that was provided on create.
func (u *AppUserAuthUpsertBulk) UpdateDeleteAt() *AppUserAuthUpsertBulk {
	return u.Update(func(s *AppUserAuthUpsert) {
		s.UpdateDeleteAt()
	})
}

// Exec executes the query.
func (u *AppUserAuthUpsertBulk) Exec(ctx context.Context) error {
	for i, b := range u.create.builders {
		if len(b.conflict) != 0 {
			return fmt.Errorf("ent: OnConflict was set for builder %d. Set it on the AppUserAuthCreateBulk instead", i)
		}
	}
	if len(u.create.conflict) == 0 {
		return errors.New("ent: missing options for AppUserAuthCreateBulk.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *AppUserAuthUpsertBulk) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}
