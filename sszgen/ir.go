package sszgen

import (
	"fmt"
	"go/ast"
	"reflect"
)

type field struct {
	Name string
	Expr expr
}

type expr interface {
}

type (
	identExpr struct {
		Type reflect.Kind
	}
)

type astStruct1 struct {
	name   string
	fields []*field
}

func (e *env) Parse(astStruct *astStruct) *astStruct1 {
	res := &astStruct1{
		name:   astStruct.name,
		fields: []*field{},
	}
	for _, obj := range astStruct.obj.Fields.List {
		field := &field{
			Name: obj.Names[0].Name,
			Expr: e.parseField(obj),
		}
		res.fields = append(res.fields, field)
	}
	return res
}

func (e *env) parseField(field *ast.Field) expr {
	switch obj := field.Type.(type) {
	case *ast.Ident:
		typ, ok := basicTypes[obj.Name]
		if ok {
			return &identExpr{Type: typ}
		}
		panic("")

	default:
		panic(fmt.Errorf("unknown type '%s'", reflect.TypeOf(field.Type).String()))
	}
}

var basicTypes = map[string]reflect.Kind{
	"uint32": reflect.Uint32,
	"uint64": reflect.Uint64,
}
