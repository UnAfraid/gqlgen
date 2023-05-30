// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package followschema

import (
	"context"
	"errors"
	"fmt"
	"strconv"
	"sync"
	"sync/atomic"

	"github.com/99designs/gqlgen/graphql"
	"github.com/vektah/gqlparser/v2/ast"
)

// region    ************************** generated!.gotpl **************************

type ErrorsResolver interface {
	A(ctx context.Context, obj *Errors) (*Error, error)
	B(ctx context.Context, obj *Errors) (*Error, error)
	C(ctx context.Context, obj *Errors) (*Error, error)
	D(ctx context.Context, obj *Errors) (*Error, error)
	E(ctx context.Context, obj *Errors) (*Error, error)
}

// endregion ************************** generated!.gotpl **************************

// region    ***************************** args.gotpl *****************************

// endregion ***************************** args.gotpl *****************************

// region    ************************** directives.gotpl **************************

// endregion ************************** directives.gotpl **************************

// region    **************************** field.gotpl *****************************

func (ec *executionContext) _Error_id(ctx context.Context, field graphql.CollectedField, obj *Error) (ret graphql.Marshaler) {
	fc, err := ec.fieldContext_Error_id(ctx, field)
	if err != nil {
		return graphql.Null
	}
	ctx = graphql.WithFieldContext(ctx, fc)
	defer func() {
		if r := recover(); r != nil {
			ec.Error(ctx, ec.Recover(ctx, r))
			ret = graphql.Null
		}
	}()
	resTmp := ec._fieldMiddleware(ctx, obj, func(rctx context.Context) (interface{}, error) {
		ctx = rctx // use context from middleware stack in children
		return obj.ID, nil
	})

	if resTmp == nil {
		if !graphql.HasFieldError(ctx, fc) {
			ec.Errorf(ctx, "must not be null")
		}
		return graphql.Null
	}
	res := resTmp.(string)
	fc.Result = res
	return ec.marshalNID2string(ctx, field.Selections, res)
}

func (ec *executionContext) fieldContext_Error_id(ctx context.Context, field graphql.CollectedField) (fc *graphql.FieldContext, err error) {
	fc = &graphql.FieldContext{
		Object:     "Error",
		Field:      field,
		IsMethod:   false,
		IsResolver: false,
		Child: func(ctx context.Context, field graphql.CollectedField) (*graphql.FieldContext, error) {
			return nil, errors.New("field of type ID does not have child fields")
		},
	}
	return fc, nil
}

func (ec *executionContext) _Error_errorOnNonRequiredField(ctx context.Context, field graphql.CollectedField, obj *Error) (ret graphql.Marshaler) {
	fc, err := ec.fieldContext_Error_errorOnNonRequiredField(ctx, field)
	if err != nil {
		return graphql.Null
	}
	ctx = graphql.WithFieldContext(ctx, fc)
	defer func() {
		if r := recover(); r != nil {
			ec.Error(ctx, ec.Recover(ctx, r))
			ret = graphql.Null
		}
	}()
	resTmp := ec._fieldMiddleware(ctx, obj, func(rctx context.Context) (interface{}, error) {
		ctx = rctx // use context from middleware stack in children
		return obj.ErrorOnNonRequiredField()
	})

	if resTmp == nil {
		return graphql.Null
	}
	res := resTmp.(string)
	fc.Result = res
	return ec.marshalOString2string(ctx, field.Selections, res)
}

func (ec *executionContext) fieldContext_Error_errorOnNonRequiredField(ctx context.Context, field graphql.CollectedField) (fc *graphql.FieldContext, err error) {
	fc = &graphql.FieldContext{
		Object:     "Error",
		Field:      field,
		IsMethod:   true,
		IsResolver: false,
		Child: func(ctx context.Context, field graphql.CollectedField) (*graphql.FieldContext, error) {
			return nil, errors.New("field of type String does not have child fields")
		},
	}
	return fc, nil
}

func (ec *executionContext) _Error_errorOnRequiredField(ctx context.Context, field graphql.CollectedField, obj *Error) (ret graphql.Marshaler) {
	fc, err := ec.fieldContext_Error_errorOnRequiredField(ctx, field)
	if err != nil {
		return graphql.Null
	}
	ctx = graphql.WithFieldContext(ctx, fc)
	defer func() {
		if r := recover(); r != nil {
			ec.Error(ctx, ec.Recover(ctx, r))
			ret = graphql.Null
		}
	}()
	resTmp := ec._fieldMiddleware(ctx, obj, func(rctx context.Context) (interface{}, error) {
		ctx = rctx // use context from middleware stack in children
		return obj.ErrorOnRequiredField()
	})

	if resTmp == nil {
		if !graphql.HasFieldError(ctx, fc) {
			ec.Errorf(ctx, "must not be null")
		}
		return graphql.Null
	}
	res := resTmp.(string)
	fc.Result = res
	return ec.marshalNString2string(ctx, field.Selections, res)
}

func (ec *executionContext) fieldContext_Error_errorOnRequiredField(ctx context.Context, field graphql.CollectedField) (fc *graphql.FieldContext, err error) {
	fc = &graphql.FieldContext{
		Object:     "Error",
		Field:      field,
		IsMethod:   true,
		IsResolver: false,
		Child: func(ctx context.Context, field graphql.CollectedField) (*graphql.FieldContext, error) {
			return nil, errors.New("field of type String does not have child fields")
		},
	}
	return fc, nil
}

func (ec *executionContext) _Error_nilOnRequiredField(ctx context.Context, field graphql.CollectedField, obj *Error) (ret graphql.Marshaler) {
	fc, err := ec.fieldContext_Error_nilOnRequiredField(ctx, field)
	if err != nil {
		return graphql.Null
	}
	ctx = graphql.WithFieldContext(ctx, fc)
	defer func() {
		if r := recover(); r != nil {
			ec.Error(ctx, ec.Recover(ctx, r))
			ret = graphql.Null
		}
	}()
	resTmp := ec._fieldMiddleware(ctx, obj, func(rctx context.Context) (interface{}, error) {
		ctx = rctx // use context from middleware stack in children
		return obj.NilOnRequiredField(), nil
	})

	if resTmp == nil {
		if !graphql.HasFieldError(ctx, fc) {
			ec.Errorf(ctx, "must not be null")
		}
		return graphql.Null
	}
	res := resTmp.(*string)
	fc.Result = res
	return ec.marshalNString2ᚖstring(ctx, field.Selections, res)
}

func (ec *executionContext) fieldContext_Error_nilOnRequiredField(ctx context.Context, field graphql.CollectedField) (fc *graphql.FieldContext, err error) {
	fc = &graphql.FieldContext{
		Object:     "Error",
		Field:      field,
		IsMethod:   true,
		IsResolver: false,
		Child: func(ctx context.Context, field graphql.CollectedField) (*graphql.FieldContext, error) {
			return nil, errors.New("field of type String does not have child fields")
		},
	}
	return fc, nil
}

func (ec *executionContext) _Errors_a(ctx context.Context, field graphql.CollectedField, obj *Errors) (ret graphql.Marshaler) {
	fc, err := ec.fieldContext_Errors_a(ctx, field)
	if err != nil {
		return graphql.Null
	}
	ctx = graphql.WithFieldContext(ctx, fc)
	defer func() {
		if r := recover(); r != nil {
			ec.Error(ctx, ec.Recover(ctx, r))
			ret = graphql.Null
		}
	}()
	resTmp := ec._fieldMiddleware(ctx, obj, func(rctx context.Context) (interface{}, error) {
		ctx = rctx // use context from middleware stack in children
		return ec.resolvers.Errors().A(rctx, obj)
	})

	if resTmp == nil {
		if !graphql.HasFieldError(ctx, fc) {
			ec.Errorf(ctx, "must not be null")
		}
		return graphql.Null
	}
	res := resTmp.(*Error)
	fc.Result = res
	return ec.marshalNError2ᚖgithubᚗcomᚋ99designsᚋgqlgenᚋcodegenᚋtestserverᚋfollowschemaᚐError(ctx, field.Selections, res)
}

func (ec *executionContext) fieldContext_Errors_a(ctx context.Context, field graphql.CollectedField) (fc *graphql.FieldContext, err error) {
	fc = &graphql.FieldContext{
		Object:     "Errors",
		Field:      field,
		IsMethod:   true,
		IsResolver: true,
		Child: func(ctx context.Context, field graphql.CollectedField) (*graphql.FieldContext, error) {
			switch field.Name {
			case "id":
				return ec.fieldContext_Error_id(ctx, field)
			case "errorOnNonRequiredField":
				return ec.fieldContext_Error_errorOnNonRequiredField(ctx, field)
			case "errorOnRequiredField":
				return ec.fieldContext_Error_errorOnRequiredField(ctx, field)
			case "nilOnRequiredField":
				return ec.fieldContext_Error_nilOnRequiredField(ctx, field)
			}
			return nil, fmt.Errorf("no field named %q was found under type Error", field.Name)
		},
	}
	return fc, nil
}

func (ec *executionContext) _Errors_b(ctx context.Context, field graphql.CollectedField, obj *Errors) (ret graphql.Marshaler) {
	fc, err := ec.fieldContext_Errors_b(ctx, field)
	if err != nil {
		return graphql.Null
	}
	ctx = graphql.WithFieldContext(ctx, fc)
	defer func() {
		if r := recover(); r != nil {
			ec.Error(ctx, ec.Recover(ctx, r))
			ret = graphql.Null
		}
	}()
	resTmp := ec._fieldMiddleware(ctx, obj, func(rctx context.Context) (interface{}, error) {
		ctx = rctx // use context from middleware stack in children
		return ec.resolvers.Errors().B(rctx, obj)
	})

	if resTmp == nil {
		if !graphql.HasFieldError(ctx, fc) {
			ec.Errorf(ctx, "must not be null")
		}
		return graphql.Null
	}
	res := resTmp.(*Error)
	fc.Result = res
	return ec.marshalNError2ᚖgithubᚗcomᚋ99designsᚋgqlgenᚋcodegenᚋtestserverᚋfollowschemaᚐError(ctx, field.Selections, res)
}

func (ec *executionContext) fieldContext_Errors_b(ctx context.Context, field graphql.CollectedField) (fc *graphql.FieldContext, err error) {
	fc = &graphql.FieldContext{
		Object:     "Errors",
		Field:      field,
		IsMethod:   true,
		IsResolver: true,
		Child: func(ctx context.Context, field graphql.CollectedField) (*graphql.FieldContext, error) {
			switch field.Name {
			case "id":
				return ec.fieldContext_Error_id(ctx, field)
			case "errorOnNonRequiredField":
				return ec.fieldContext_Error_errorOnNonRequiredField(ctx, field)
			case "errorOnRequiredField":
				return ec.fieldContext_Error_errorOnRequiredField(ctx, field)
			case "nilOnRequiredField":
				return ec.fieldContext_Error_nilOnRequiredField(ctx, field)
			}
			return nil, fmt.Errorf("no field named %q was found under type Error", field.Name)
		},
	}
	return fc, nil
}

func (ec *executionContext) _Errors_c(ctx context.Context, field graphql.CollectedField, obj *Errors) (ret graphql.Marshaler) {
	fc, err := ec.fieldContext_Errors_c(ctx, field)
	if err != nil {
		return graphql.Null
	}
	ctx = graphql.WithFieldContext(ctx, fc)
	defer func() {
		if r := recover(); r != nil {
			ec.Error(ctx, ec.Recover(ctx, r))
			ret = graphql.Null
		}
	}()
	resTmp := ec._fieldMiddleware(ctx, obj, func(rctx context.Context) (interface{}, error) {
		ctx = rctx // use context from middleware stack in children
		return ec.resolvers.Errors().C(rctx, obj)
	})

	if resTmp == nil {
		if !graphql.HasFieldError(ctx, fc) {
			ec.Errorf(ctx, "must not be null")
		}
		return graphql.Null
	}
	res := resTmp.(*Error)
	fc.Result = res
	return ec.marshalNError2ᚖgithubᚗcomᚋ99designsᚋgqlgenᚋcodegenᚋtestserverᚋfollowschemaᚐError(ctx, field.Selections, res)
}

func (ec *executionContext) fieldContext_Errors_c(ctx context.Context, field graphql.CollectedField) (fc *graphql.FieldContext, err error) {
	fc = &graphql.FieldContext{
		Object:     "Errors",
		Field:      field,
		IsMethod:   true,
		IsResolver: true,
		Child: func(ctx context.Context, field graphql.CollectedField) (*graphql.FieldContext, error) {
			switch field.Name {
			case "id":
				return ec.fieldContext_Error_id(ctx, field)
			case "errorOnNonRequiredField":
				return ec.fieldContext_Error_errorOnNonRequiredField(ctx, field)
			case "errorOnRequiredField":
				return ec.fieldContext_Error_errorOnRequiredField(ctx, field)
			case "nilOnRequiredField":
				return ec.fieldContext_Error_nilOnRequiredField(ctx, field)
			}
			return nil, fmt.Errorf("no field named %q was found under type Error", field.Name)
		},
	}
	return fc, nil
}

func (ec *executionContext) _Errors_d(ctx context.Context, field graphql.CollectedField, obj *Errors) (ret graphql.Marshaler) {
	fc, err := ec.fieldContext_Errors_d(ctx, field)
	if err != nil {
		return graphql.Null
	}
	ctx = graphql.WithFieldContext(ctx, fc)
	defer func() {
		if r := recover(); r != nil {
			ec.Error(ctx, ec.Recover(ctx, r))
			ret = graphql.Null
		}
	}()
	resTmp := ec._fieldMiddleware(ctx, obj, func(rctx context.Context) (interface{}, error) {
		ctx = rctx // use context from middleware stack in children
		return ec.resolvers.Errors().D(rctx, obj)
	})

	if resTmp == nil {
		if !graphql.HasFieldError(ctx, fc) {
			ec.Errorf(ctx, "must not be null")
		}
		return graphql.Null
	}
	res := resTmp.(*Error)
	fc.Result = res
	return ec.marshalNError2ᚖgithubᚗcomᚋ99designsᚋgqlgenᚋcodegenᚋtestserverᚋfollowschemaᚐError(ctx, field.Selections, res)
}

func (ec *executionContext) fieldContext_Errors_d(ctx context.Context, field graphql.CollectedField) (fc *graphql.FieldContext, err error) {
	fc = &graphql.FieldContext{
		Object:     "Errors",
		Field:      field,
		IsMethod:   true,
		IsResolver: true,
		Child: func(ctx context.Context, field graphql.CollectedField) (*graphql.FieldContext, error) {
			switch field.Name {
			case "id":
				return ec.fieldContext_Error_id(ctx, field)
			case "errorOnNonRequiredField":
				return ec.fieldContext_Error_errorOnNonRequiredField(ctx, field)
			case "errorOnRequiredField":
				return ec.fieldContext_Error_errorOnRequiredField(ctx, field)
			case "nilOnRequiredField":
				return ec.fieldContext_Error_nilOnRequiredField(ctx, field)
			}
			return nil, fmt.Errorf("no field named %q was found under type Error", field.Name)
		},
	}
	return fc, nil
}

func (ec *executionContext) _Errors_e(ctx context.Context, field graphql.CollectedField, obj *Errors) (ret graphql.Marshaler) {
	fc, err := ec.fieldContext_Errors_e(ctx, field)
	if err != nil {
		return graphql.Null
	}
	ctx = graphql.WithFieldContext(ctx, fc)
	defer func() {
		if r := recover(); r != nil {
			ec.Error(ctx, ec.Recover(ctx, r))
			ret = graphql.Null
		}
	}()
	resTmp := ec._fieldMiddleware(ctx, obj, func(rctx context.Context) (interface{}, error) {
		ctx = rctx // use context from middleware stack in children
		return ec.resolvers.Errors().E(rctx, obj)
	})

	if resTmp == nil {
		if !graphql.HasFieldError(ctx, fc) {
			ec.Errorf(ctx, "must not be null")
		}
		return graphql.Null
	}
	res := resTmp.(*Error)
	fc.Result = res
	return ec.marshalNError2ᚖgithubᚗcomᚋ99designsᚋgqlgenᚋcodegenᚋtestserverᚋfollowschemaᚐError(ctx, field.Selections, res)
}

func (ec *executionContext) fieldContext_Errors_e(ctx context.Context, field graphql.CollectedField) (fc *graphql.FieldContext, err error) {
	fc = &graphql.FieldContext{
		Object:     "Errors",
		Field:      field,
		IsMethod:   true,
		IsResolver: true,
		Child: func(ctx context.Context, field graphql.CollectedField) (*graphql.FieldContext, error) {
			switch field.Name {
			case "id":
				return ec.fieldContext_Error_id(ctx, field)
			case "errorOnNonRequiredField":
				return ec.fieldContext_Error_errorOnNonRequiredField(ctx, field)
			case "errorOnRequiredField":
				return ec.fieldContext_Error_errorOnRequiredField(ctx, field)
			case "nilOnRequiredField":
				return ec.fieldContext_Error_nilOnRequiredField(ctx, field)
			}
			return nil, fmt.Errorf("no field named %q was found under type Error", field.Name)
		},
	}
	return fc, nil
}

// endregion **************************** field.gotpl *****************************

// region    **************************** input.gotpl *****************************

// endregion **************************** input.gotpl *****************************

// region    ************************** interface.gotpl ***************************

// endregion ************************** interface.gotpl ***************************

// region    **************************** object.gotpl ****************************

var errorImplementors = []string{"Error"}

func (ec *executionContext) _Error(ctx context.Context, sel ast.SelectionSet, obj *Error) graphql.Marshaler {
	fields := graphql.CollectFields(ec.OperationContext, sel, errorImplementors)

	out := graphql.NewFieldSet(fields)
	deferred := make(map[string]*graphql.FieldSet)
	for i, field := range fields {
		switch field.Name {
		case "__typename":
			out.Values[i] = graphql.MarshalString("Error")
		case "id":
			out.Values[i] = ec._Error_id(ctx, field, obj)
			if out.Values[i] == graphql.Null {
				out.Invalids++
			}
		case "errorOnNonRequiredField":
			out.Values[i] = ec._Error_errorOnNonRequiredField(ctx, field, obj)
		case "errorOnRequiredField":
			out.Values[i] = ec._Error_errorOnRequiredField(ctx, field, obj)
			if out.Values[i] == graphql.Null {
				out.Invalids++
			}
		case "nilOnRequiredField":
			out.Values[i] = ec._Error_nilOnRequiredField(ctx, field, obj)
			if out.Values[i] == graphql.Null {
				out.Invalids++
			}
		default:
			panic("unknown field " + strconv.Quote(field.Name))
		}
	}
	out.Dispatch(ctx)
	if out.Invalids > 0 {
		return graphql.Null
	}

	ec.deferred.Add(int32(len(deferred)))

	for label, dfs := range deferred {
		ec.processDeferredGroup(graphql.DeferredGroup{
			Label:    label,
			Path:     graphql.GetPath(ctx),
			FieldSet: dfs,
			Context:  ctx,
		})
	}

	return out
}

var errorsImplementors = []string{"Errors"}

func (ec *executionContext) _Errors(ctx context.Context, sel ast.SelectionSet, obj *Errors) graphql.Marshaler {
	fields := graphql.CollectFields(ec.OperationContext, sel, errorsImplementors)

	out := graphql.NewFieldSet(fields)
	deferred := make(map[string]*graphql.FieldSet)
	for i, field := range fields {
		switch field.Name {
		case "__typename":
			out.Values[i] = graphql.MarshalString("Errors")
		case "a":
			field := field

			innerFunc := func(ctx context.Context, fs *graphql.FieldSet) (res graphql.Marshaler) {
				defer func() {
					if r := recover(); r != nil {
						ec.Error(ctx, ec.Recover(ctx, r))
					}
				}()
				res = ec._Errors_a(ctx, field, obj)
				if res == graphql.Null {
					atomic.AddUint32(&fs.Invalids, 1)
				}
				return res
			}

			if field.Deferrable != nil {
				dfs, ok := deferred[field.Deferrable.Label]
				di := 0
				if ok {
					dfs.AddField(field)
					di = len(dfs.Values) - 1
				} else {
					dfs = graphql.NewFieldSet([]graphql.CollectedField{field})
					deferred[field.Deferrable.Label] = dfs
				}
				dfs.Concurrently(di, func(ctx context.Context) graphql.Marshaler {
					return innerFunc(ctx, dfs)
				})

				// don't run the out.Concurrently() call below
				out.Values[i] = graphql.Null
				continue
			}

			out.Concurrently(i, func(ctx context.Context) graphql.Marshaler { return innerFunc(ctx, out) })
		case "b":
			field := field

			innerFunc := func(ctx context.Context, fs *graphql.FieldSet) (res graphql.Marshaler) {
				defer func() {
					if r := recover(); r != nil {
						ec.Error(ctx, ec.Recover(ctx, r))
					}
				}()
				res = ec._Errors_b(ctx, field, obj)
				if res == graphql.Null {
					atomic.AddUint32(&fs.Invalids, 1)
				}
				return res
			}

			if field.Deferrable != nil {
				dfs, ok := deferred[field.Deferrable.Label]
				di := 0
				if ok {
					dfs.AddField(field)
					di = len(dfs.Values) - 1
				} else {
					dfs = graphql.NewFieldSet([]graphql.CollectedField{field})
					deferred[field.Deferrable.Label] = dfs
				}
				dfs.Concurrently(di, func(ctx context.Context) graphql.Marshaler {
					return innerFunc(ctx, dfs)
				})

				// don't run the out.Concurrently() call below
				out.Values[i] = graphql.Null
				continue
			}

			out.Concurrently(i, func(ctx context.Context) graphql.Marshaler { return innerFunc(ctx, out) })
		case "c":
			field := field

			innerFunc := func(ctx context.Context, fs *graphql.FieldSet) (res graphql.Marshaler) {
				defer func() {
					if r := recover(); r != nil {
						ec.Error(ctx, ec.Recover(ctx, r))
					}
				}()
				res = ec._Errors_c(ctx, field, obj)
				if res == graphql.Null {
					atomic.AddUint32(&fs.Invalids, 1)
				}
				return res
			}

			if field.Deferrable != nil {
				dfs, ok := deferred[field.Deferrable.Label]
				di := 0
				if ok {
					dfs.AddField(field)
					di = len(dfs.Values) - 1
				} else {
					dfs = graphql.NewFieldSet([]graphql.CollectedField{field})
					deferred[field.Deferrable.Label] = dfs
				}
				dfs.Concurrently(di, func(ctx context.Context) graphql.Marshaler {
					return innerFunc(ctx, dfs)
				})

				// don't run the out.Concurrently() call below
				out.Values[i] = graphql.Null
				continue
			}

			out.Concurrently(i, func(ctx context.Context) graphql.Marshaler { return innerFunc(ctx, out) })
		case "d":
			field := field

			innerFunc := func(ctx context.Context, fs *graphql.FieldSet) (res graphql.Marshaler) {
				defer func() {
					if r := recover(); r != nil {
						ec.Error(ctx, ec.Recover(ctx, r))
					}
				}()
				res = ec._Errors_d(ctx, field, obj)
				if res == graphql.Null {
					atomic.AddUint32(&fs.Invalids, 1)
				}
				return res
			}

			if field.Deferrable != nil {
				dfs, ok := deferred[field.Deferrable.Label]
				di := 0
				if ok {
					dfs.AddField(field)
					di = len(dfs.Values) - 1
				} else {
					dfs = graphql.NewFieldSet([]graphql.CollectedField{field})
					deferred[field.Deferrable.Label] = dfs
				}
				dfs.Concurrently(di, func(ctx context.Context) graphql.Marshaler {
					return innerFunc(ctx, dfs)
				})

				// don't run the out.Concurrently() call below
				out.Values[i] = graphql.Null
				continue
			}

			out.Concurrently(i, func(ctx context.Context) graphql.Marshaler { return innerFunc(ctx, out) })
		case "e":
			field := field

			innerFunc := func(ctx context.Context, fs *graphql.FieldSet) (res graphql.Marshaler) {
				defer func() {
					if r := recover(); r != nil {
						ec.Error(ctx, ec.Recover(ctx, r))
					}
				}()
				res = ec._Errors_e(ctx, field, obj)
				if res == graphql.Null {
					atomic.AddUint32(&fs.Invalids, 1)
				}
				return res
			}

			if field.Deferrable != nil {
				dfs, ok := deferred[field.Deferrable.Label]
				di := 0
				if ok {
					dfs.AddField(field)
					di = len(dfs.Values) - 1
				} else {
					dfs = graphql.NewFieldSet([]graphql.CollectedField{field})
					deferred[field.Deferrable.Label] = dfs
				}
				dfs.Concurrently(di, func(ctx context.Context) graphql.Marshaler {
					return innerFunc(ctx, dfs)
				})

				// don't run the out.Concurrently() call below
				out.Values[i] = graphql.Null
				continue
			}

			out.Concurrently(i, func(ctx context.Context) graphql.Marshaler { return innerFunc(ctx, out) })
		default:
			panic("unknown field " + strconv.Quote(field.Name))
		}
	}
	out.Dispatch(ctx)
	if out.Invalids > 0 {
		return graphql.Null
	}

	ec.deferred.Add(int32(len(deferred)))

	for label, dfs := range deferred {
		ec.processDeferredGroup(graphql.DeferredGroup{
			Label:    label,
			Path:     graphql.GetPath(ctx),
			FieldSet: dfs,
			Context:  ctx,
		})
	}

	return out
}

// endregion **************************** object.gotpl ****************************

// region    ***************************** type.gotpl *****************************

func (ec *executionContext) marshalNError2githubᚗcomᚋ99designsᚋgqlgenᚋcodegenᚋtestserverᚋfollowschemaᚐError(ctx context.Context, sel ast.SelectionSet, v Error) graphql.Marshaler {
	return ec._Error(ctx, sel, &v)
}

func (ec *executionContext) marshalNError2ᚖgithubᚗcomᚋ99designsᚋgqlgenᚋcodegenᚋtestserverᚋfollowschemaᚐError(ctx context.Context, sel ast.SelectionSet, v *Error) graphql.Marshaler {
	if v == nil {
		if !graphql.HasFieldError(ctx, graphql.GetFieldContext(ctx)) {
			ec.Errorf(ctx, "the requested element is null which the schema does not allow")
		}
		return graphql.Null
	}
	return ec._Error(ctx, sel, v)
}

func (ec *executionContext) marshalOError2ᚕᚖgithubᚗcomᚋ99designsᚋgqlgenᚋcodegenᚋtestserverᚋfollowschemaᚐError(ctx context.Context, sel ast.SelectionSet, v []*Error) graphql.Marshaler {
	if v == nil {
		return graphql.Null
	}
	ret := make(graphql.Array, len(v))
	var wg sync.WaitGroup
	isLen1 := len(v) == 1
	if !isLen1 {
		wg.Add(len(v))
	}
	for i := range v {
		i := i
		fc := &graphql.FieldContext{
			Index:  &i,
			Result: &v[i],
		}
		ctx := graphql.WithFieldContext(ctx, fc)
		f := func(i int) {
			defer func() {
				if r := recover(); r != nil {
					ec.Error(ctx, ec.Recover(ctx, r))
					ret = nil
				}
			}()
			if !isLen1 {
				defer wg.Done()
			}
			ret[i] = ec.marshalOError2ᚖgithubᚗcomᚋ99designsᚋgqlgenᚋcodegenᚋtestserverᚋfollowschemaᚐError(ctx, sel, v[i])
		}
		if isLen1 {
			f(i)
		} else {
			go f(i)
		}

	}
	wg.Wait()

	return ret
}

func (ec *executionContext) marshalOError2ᚕᚖgithubᚗcomᚋ99designsᚋgqlgenᚋcodegenᚋtestserverᚋfollowschemaᚐErrorᚄ(ctx context.Context, sel ast.SelectionSet, v []*Error) graphql.Marshaler {
	if v == nil {
		return graphql.Null
	}
	ret := make(graphql.Array, len(v))
	var wg sync.WaitGroup
	isLen1 := len(v) == 1
	if !isLen1 {
		wg.Add(len(v))
	}
	for i := range v {
		i := i
		fc := &graphql.FieldContext{
			Index:  &i,
			Result: &v[i],
		}
		ctx := graphql.WithFieldContext(ctx, fc)
		f := func(i int) {
			defer func() {
				if r := recover(); r != nil {
					ec.Error(ctx, ec.Recover(ctx, r))
					ret = nil
				}
			}()
			if !isLen1 {
				defer wg.Done()
			}
			ret[i] = ec.marshalNError2ᚖgithubᚗcomᚋ99designsᚋgqlgenᚋcodegenᚋtestserverᚋfollowschemaᚐError(ctx, sel, v[i])
		}
		if isLen1 {
			f(i)
		} else {
			go f(i)
		}

	}
	wg.Wait()

	for _, e := range ret {
		if e == graphql.Null {
			return graphql.Null
		}
	}

	return ret
}

func (ec *executionContext) marshalOError2ᚖgithubᚗcomᚋ99designsᚋgqlgenᚋcodegenᚋtestserverᚋfollowschemaᚐError(ctx context.Context, sel ast.SelectionSet, v *Error) graphql.Marshaler {
	if v == nil {
		return graphql.Null
	}
	return ec._Error(ctx, sel, v)
}

func (ec *executionContext) marshalOErrors2ᚖgithubᚗcomᚋ99designsᚋgqlgenᚋcodegenᚋtestserverᚋfollowschemaᚐErrors(ctx context.Context, sel ast.SelectionSet, v *Errors) graphql.Marshaler {
	if v == nil {
		return graphql.Null
	}
	return ec._Errors(ctx, sel, v)
}

// endregion ***************************** type.gotpl *****************************
