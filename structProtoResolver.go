package structproto

import (
	"fmt"
	"reflect"

	"github.com/bcowtech/structproto/tagresolver"
)

type StructProtoResolver struct {
	tagName     string
	tagResolver TagResolver
}

func NewStructProtoResolver(option *StructProtoResolveOption) *StructProtoResolver {
	if option == nil {
		panic("specified argument 'option' cannot be nil")
	}

	r := &StructProtoResolver{
		tagName:     option.TagName,
		tagResolver: option.TagResolver,
	}

	// use StdTagResolver if missing
	if r.tagResolver == nil {
		r.tagResolver = tagresolver.StdTagResolver
	}
	return r
}

func (r *StructProtoResolver) Resolve(target interface{}) (*Struct, error) {
	var rv reflect.Value
	switch target.(type) {
	case reflect.Value:
		rv = target.(reflect.Value)
	default:
		rv = reflect.ValueOf(target)
	}

	if !rv.IsValid() {
		return nil, fmt.Errorf("specified argument 'target' is invalid")
	}

	for i := 0; true; i++ {
		if i >= 32 {
			return nil, fmt.Errorf("exceed maximum processing calls")
		}
		switch rv.Kind() {
		case reflect.Struct:
			info, err := r.internalResolve(rv)
			if err != nil {
				return nil, err
			}
			return info, nil
		case reflect.Ptr:
			if rv.IsNil() {
				rv = reflect.New(rv.Type().Elem())
			}
			rv = rv.Elem()
		default:
			return nil, fmt.Errorf("specified argument 'target' must be pointer to struct")
		}
	}
	return nil, nil
}

func (r *StructProtoResolver) internalResolve(rv reflect.Value) (*Struct, error) {
	var prototype = buildStruct(rv)
	t := rv.Type()
	count := t.NumField()
	for i := 0; i < count; i++ {
		fieldname := t.Field(i).Name
		token := t.Field(i).Tag.Get(r.tagName)
		tag, err := r.tagResolver(fieldname, token)
		if err != nil {
			return nil, err
		}
		if tag != nil {
			field := &Field{
				name:  tag.Name,
				index: i,
				desc:  tag.Desc,
			}
			field.appendFlags(tag.Flags...)

			prototype.fields[tag.Name] = field
			if field.HasFlag(RequiredFlag) {
				prototype.requiredFields.Append(tag.Name)
			}
		}
	}
	return prototype, nil
}
