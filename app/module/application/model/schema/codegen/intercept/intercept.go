// Code generated by ent, DO NOT EDIT.

package intercept

import (
	"context"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"github.com/wangxg422/XishangOS-backend/app/module/application/model/schema/codegen"
	"github.com/wangxg422/XishangOS-backend/app/module/application/model/schema/codegen/appinstance"
	"github.com/wangxg422/XishangOS-backend/app/module/application/model/schema/codegen/apppackage"
	"github.com/wangxg422/XishangOS-backend/app/module/application/model/schema/codegen/predicate"
)

// The Query interface represents an operation that queries a graph.
// By using this interface, users can write generic code that manipulates
// query builders of different types.
type Query interface {
	// Type returns the string representation of the query type.
	Type() string
	// Limit the number of records to be returned by this query.
	Limit(int)
	// Offset to start from.
	Offset(int)
	// Unique configures the query builder to filter duplicate records.
	Unique(bool)
	// Order specifies how the records should be ordered.
	Order(...func(*sql.Selector))
	// WhereP appends storage-level predicates to the query builder. Using this method, users
	// can use type-assertion to append predicates that do not depend on any generated package.
	WhereP(...func(*sql.Selector))
}

// The Func type is an adapter that allows ordinary functions to be used as interceptors.
// Unlike traversal functions, interceptors are skipped during graph traversals. Note that the
// implementation of Func is different from the one defined in entgo.io/ent.InterceptFunc.
type Func func(context.Context, Query) error

// Intercept calls f(ctx, q) and then applied the next Querier.
func (f Func) Intercept(next codegen.Querier) codegen.Querier {
	return codegen.QuerierFunc(func(ctx context.Context, q codegen.Query) (codegen.Value, error) {
		query, err := NewQuery(q)
		if err != nil {
			return nil, err
		}
		if err := f(ctx, query); err != nil {
			return nil, err
		}
		return next.Query(ctx, q)
	})
}

// The TraverseFunc type is an adapter to allow the use of ordinary function as Traverser.
// If f is a function with the appropriate signature, TraverseFunc(f) is a Traverser that calls f.
type TraverseFunc func(context.Context, Query) error

// Intercept is a dummy implementation of Intercept that returns the next Querier in the pipeline.
func (f TraverseFunc) Intercept(next codegen.Querier) codegen.Querier {
	return next
}

// Traverse calls f(ctx, q).
func (f TraverseFunc) Traverse(ctx context.Context, q codegen.Query) error {
	query, err := NewQuery(q)
	if err != nil {
		return err
	}
	return f(ctx, query)
}

// The AppInstanceFunc type is an adapter to allow the use of ordinary function as a Querier.
type AppInstanceFunc func(context.Context, *codegen.AppInstanceQuery) (codegen.Value, error)

// Query calls f(ctx, q).
func (f AppInstanceFunc) Query(ctx context.Context, q codegen.Query) (codegen.Value, error) {
	if q, ok := q.(*codegen.AppInstanceQuery); ok {
		return f(ctx, q)
	}
	return nil, fmt.Errorf("unexpected query type %T. expect *codegen.AppInstanceQuery", q)
}

// The TraverseAppInstance type is an adapter to allow the use of ordinary function as Traverser.
type TraverseAppInstance func(context.Context, *codegen.AppInstanceQuery) error

// Intercept is a dummy implementation of Intercept that returns the next Querier in the pipeline.
func (f TraverseAppInstance) Intercept(next codegen.Querier) codegen.Querier {
	return next
}

// Traverse calls f(ctx, q).
func (f TraverseAppInstance) Traverse(ctx context.Context, q codegen.Query) error {
	if q, ok := q.(*codegen.AppInstanceQuery); ok {
		return f(ctx, q)
	}
	return fmt.Errorf("unexpected query type %T. expect *codegen.AppInstanceQuery", q)
}

// The AppPackageFunc type is an adapter to allow the use of ordinary function as a Querier.
type AppPackageFunc func(context.Context, *codegen.AppPackageQuery) (codegen.Value, error)

// Query calls f(ctx, q).
func (f AppPackageFunc) Query(ctx context.Context, q codegen.Query) (codegen.Value, error) {
	if q, ok := q.(*codegen.AppPackageQuery); ok {
		return f(ctx, q)
	}
	return nil, fmt.Errorf("unexpected query type %T. expect *codegen.AppPackageQuery", q)
}

// The TraverseAppPackage type is an adapter to allow the use of ordinary function as Traverser.
type TraverseAppPackage func(context.Context, *codegen.AppPackageQuery) error

// Intercept is a dummy implementation of Intercept that returns the next Querier in the pipeline.
func (f TraverseAppPackage) Intercept(next codegen.Querier) codegen.Querier {
	return next
}

// Traverse calls f(ctx, q).
func (f TraverseAppPackage) Traverse(ctx context.Context, q codegen.Query) error {
	if q, ok := q.(*codegen.AppPackageQuery); ok {
		return f(ctx, q)
	}
	return fmt.Errorf("unexpected query type %T. expect *codegen.AppPackageQuery", q)
}

// NewQuery returns the generic Query interface for the given typed query.
func NewQuery(q codegen.Query) (Query, error) {
	switch q := q.(type) {
	case *codegen.AppInstanceQuery:
		return &query[*codegen.AppInstanceQuery, predicate.AppInstance, appinstance.OrderOption]{typ: codegen.TypeAppInstance, tq: q}, nil
	case *codegen.AppPackageQuery:
		return &query[*codegen.AppPackageQuery, predicate.AppPackage, apppackage.OrderOption]{typ: codegen.TypeAppPackage, tq: q}, nil
	default:
		return nil, fmt.Errorf("unknown query type %T", q)
	}
}

type query[T any, P ~func(*sql.Selector), R ~func(*sql.Selector)] struct {
	typ string
	tq  interface {
		Limit(int) T
		Offset(int) T
		Unique(bool) T
		Order(...R) T
		Where(...P) T
	}
}

func (q query[T, P, R]) Type() string {
	return q.typ
}

func (q query[T, P, R]) Limit(limit int) {
	q.tq.Limit(limit)
}

func (q query[T, P, R]) Offset(offset int) {
	q.tq.Offset(offset)
}

func (q query[T, P, R]) Unique(unique bool) {
	q.tq.Unique(unique)
}

func (q query[T, P, R]) Order(orders ...func(*sql.Selector)) {
	rs := make([]R, len(orders))
	for i := range orders {
		rs[i] = orders[i]
	}
	q.tq.Order(rs...)
}

func (q query[T, P, R]) WhereP(ps ...func(*sql.Selector)) {
	p := make([]P, len(ps))
	for i := range ps {
		p[i] = ps[i]
	}
	q.tq.Where(p...)
}