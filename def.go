package structproto

import (
	"reflect"

	"github.com/bcowtech/structproto/internal"
)

const (
	RequiredFlag = internal.RequiredFlag
)

type (
	ValueBinder = internal.ValueBinder
	TagResolver = internal.TagResolver
	Tag         = internal.Tag

	ValueBinderProvider func(rv reflect.Value) ValueBinder

	FieldValueEntry struct {
		Field string
		Value interface{}
	}

	FieldValueCollectionIterator interface {
		Iterate() <-chan FieldValueEntry
	}

	FieldInfo interface {
		Name() string
		Desc() string
		Index() int
		Flags() []string
		HasFlag(v string) bool
	}

	StructBinder interface {
		Init(context *StructProtoContext) error
		Bind(field FieldInfo, rv reflect.Value) error
		Deinit(context *StructProtoContext) error
	}

	StructProtoOption struct {
		TagName             string
		ValueBinderProvider ValueBinderProvider
		TagResolver         TagResolver
	}
)
