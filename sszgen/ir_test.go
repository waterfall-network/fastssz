package sszgen

import (
	"fmt"
	"io/ioutil"
	"path/filepath"
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIR(t *testing.T) {

	tmpDir, err := ioutil.TempDir("/tmp", "ssz-")
	assert.NoError(t, err)

	code := `
	type A struct {
		A1 uint32
	}
	type B struct {
		B1 uint64
	}
	type C struct {
		C1 []byte
	}
	`

	filePath := filepath.Join(tmpDir, "./test.go")

	fullCode := `package main
	` + code
	err = ioutil.WriteFile(filePath, []byte(fullCode), 0755)
	assert.NoError(t, err)

	input, err := parseInput(filePath)
	assert.NoError(t, err)

	env := &env{}

	result := []*astStruct1{}
	for _, file := range input {
		res := decodeASTStruct(file)
		for _, obj := range res.objs {
			result = append(result, env.Parse(obj))
		}
	}

	expect := []*astStruct1{
		{
			name: "A",
			fields: []*field{
				{
					Name: "A1",
					Expr: &identExpr{
						Type: reflect.Uint32,
					},
				},
			},
		},
		{
			name: "B",
			fields: []*field{
				{
					Name: "B1",
					Expr: &identExpr{
						Type: reflect.Uint64,
					},
				},
			},
		},
	}
	fmt.Println(expect[0].fields[0])
	fmt.Println(result[0].fields[0])
	fmt.Println(reflect.DeepEqual(expect, result))
}
