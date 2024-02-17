package db

import (
	"reflect"
)

type Field struct {
	name     string
	kind     reflect.Kind
	nullable bool
}

func (f *Field) Name() string {
	return f.name
}

func (f *Field) Kind() reflect.Kind {
	return f.kind
}

func (f *Field) Nullable() bool {
	return f.nullable
}
