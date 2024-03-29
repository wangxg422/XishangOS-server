// Code generated by ent, DO NOT EDIT.

package codegen

import (
	"context"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/wangxg422/XishangOS-backend/app/module/system/model/schema/codegen/predicate"
	"github.com/wangxg422/XishangOS-backend/app/module/system/model/schema/codegen/sysloginlog"
)

// SysLoginLogDelete is the builder for deleting a SysLoginLog entity.
type SysLoginLogDelete struct {
	config
	hooks    []Hook
	mutation *SysLoginLogMutation
}

// Where appends a list predicates to the SysLoginLogDelete builder.
func (slld *SysLoginLogDelete) Where(ps ...predicate.SysLoginLog) *SysLoginLogDelete {
	slld.mutation.Where(ps...)
	return slld
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (slld *SysLoginLogDelete) Exec(ctx context.Context) (int, error) {
	return withHooks(ctx, slld.sqlExec, slld.mutation, slld.hooks)
}

// ExecX is like Exec, but panics if an error occurs.
func (slld *SysLoginLogDelete) ExecX(ctx context.Context) int {
	n, err := slld.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (slld *SysLoginLogDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := sqlgraph.NewDeleteSpec(sysloginlog.Table, sqlgraph.NewFieldSpec(sysloginlog.FieldID, field.TypeInt64))
	if ps := slld.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	affected, err := sqlgraph.DeleteNodes(ctx, slld.driver, _spec)
	if err != nil && sqlgraph.IsConstraintError(err) {
		err = &ConstraintError{msg: err.Error(), wrap: err}
	}
	slld.mutation.done = true
	return affected, err
}

// SysLoginLogDeleteOne is the builder for deleting a single SysLoginLog entity.
type SysLoginLogDeleteOne struct {
	slld *SysLoginLogDelete
}

// Where appends a list predicates to the SysLoginLogDelete builder.
func (slldo *SysLoginLogDeleteOne) Where(ps ...predicate.SysLoginLog) *SysLoginLogDeleteOne {
	slldo.slld.mutation.Where(ps...)
	return slldo
}

// Exec executes the deletion query.
func (slldo *SysLoginLogDeleteOne) Exec(ctx context.Context) error {
	n, err := slldo.slld.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{sysloginlog.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (slldo *SysLoginLogDeleteOne) ExecX(ctx context.Context) {
	if err := slldo.Exec(ctx); err != nil {
		panic(err)
	}
}
