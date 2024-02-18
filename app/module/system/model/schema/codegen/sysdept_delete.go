// Code generated by ent, DO NOT EDIT.

package codegen

import (
	"context"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/wangxg422/XishangOS-backend/app/module/system/model/schema/codegen/predicate"
	"github.com/wangxg422/XishangOS-backend/app/module/system/model/schema/codegen/sysdept"
)

// SysDeptDelete is the builder for deleting a SysDept entity.
type SysDeptDelete struct {
	config
	hooks    []Hook
	mutation *SysDeptMutation
}

// Where appends a list predicates to the SysDeptDelete builder.
func (sdd *SysDeptDelete) Where(ps ...predicate.SysDept) *SysDeptDelete {
	sdd.mutation.Where(ps...)
	return sdd
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (sdd *SysDeptDelete) Exec(ctx context.Context) (int, error) {
	return withHooks(ctx, sdd.sqlExec, sdd.mutation, sdd.hooks)
}

// ExecX is like Exec, but panics if an error occurs.
func (sdd *SysDeptDelete) ExecX(ctx context.Context) int {
	n, err := sdd.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (sdd *SysDeptDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := sqlgraph.NewDeleteSpec(sysdept.Table, sqlgraph.NewFieldSpec(sysdept.FieldID, field.TypeInt64))
	if ps := sdd.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	affected, err := sqlgraph.DeleteNodes(ctx, sdd.driver, _spec)
	if err != nil && sqlgraph.IsConstraintError(err) {
		err = &ConstraintError{msg: err.Error(), wrap: err}
	}
	sdd.mutation.done = true
	return affected, err
}

// SysDeptDeleteOne is the builder for deleting a single SysDept entity.
type SysDeptDeleteOne struct {
	sdd *SysDeptDelete
}

// Where appends a list predicates to the SysDeptDelete builder.
func (sddo *SysDeptDeleteOne) Where(ps ...predicate.SysDept) *SysDeptDeleteOne {
	sddo.sdd.mutation.Where(ps...)
	return sddo
}

// Exec executes the deletion query.
func (sddo *SysDeptDeleteOne) Exec(ctx context.Context) error {
	n, err := sddo.sdd.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{sysdept.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (sddo *SysDeptDeleteOne) ExecX(ctx context.Context) {
	if err := sddo.Exec(ctx); err != nil {
		panic(err)
	}
}