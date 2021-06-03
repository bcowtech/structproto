package structproto

import (
	"github.com/bcowtech/structproto/tagresolver"
	"github.com/bcowtech/structproto/valuebinder"
)

func Prototypify(target interface{}, option *StructProtoOption) (*Struct, error) {
	if target == nil {
		panic("specified argument 'target' cannot be nil")
	}
	if option == nil {
		panic("specified argument 'option' cannot be nil")
	}

	if option.ValueBinderProvider == nil {
		option.ValueBinderProvider = valuebinder.BuildNilBinder
	}

	r := &StructProtoResolver{
		TagName:             option.TagName,
		ValueBinderProvider: option.ValueBinderProvider,
		TagResolver:         option.TagResolver,
	}
	if r.TagResolver == nil {
		r.TagResolver = tagresolver.StdTagResolver
	}
	prototype, err := r.Resolve(target)
	if err != nil {
		return nil, err
	}
	return prototype, nil
}
