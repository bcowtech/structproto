package structproto

import (
	"reflect"
)

type Struct struct {
	value reflect.Value

	fields         map[string]*Field
	requiredFields FieldFlagSet

	valueBinderProvider ValueBinderProvider
}

func (s *Struct) Bind(binder StructBinder) error {
	if binder == nil {
		panic("specified argument 'binder' cannot be nil")
	}

	var err error

	context, err := buildStructProtoContext(s)
	if err != nil {
		return err
	}

	if err = binder.Init(context); err != nil {
		return err
	}

	// bind all fields
	for _, field := range s.fields {
		err := binder.Bind(field, s.value.Field(field.index))
		if err != nil {
			return err
		}
	}

	if err = binder.Deinit(context); err != nil {
		return err
	}

	return nil
}

func (s *Struct) BindValues(values KeyValuePairIterator) error {
	if s == nil {
		return nil
	}

	var requiredFields = s.requiredFields.Clone()

	// mapping values
	for p := range values.Iterate() {
		field, val := p.Key, p.Value
		if val != nil {
			binder := s.makeFieldBinder(s.value, field)
			if binder != nil {
				err := binder.Bind(val)
				if err != nil {
					return &FieldBindingError{field, val, err}
				}

				index := requiredFields.IndexOf(field)
				if index != -1 {
					// eliminate the field from slice if found
					requiredFields.RemoveIndex(index)
				}
			}
		}
	}

	// check if the requiredFields still have fields don't be set
	if !requiredFields.IsEmpty() {
		field, _ := requiredFields.Get(0)
		return &MissingRequiredFieldError{field, nil}
	}

	return nil
}

func (s *Struct) makeFieldBinder(rv reflect.Value, name string) ValueBinder {
	if s == nil {
		return nil
	}
	if f, ok := s.fields[name]; ok {
		binder := s.valueBinderProvider(rv.Field(f.index))
		return binder
	}
	return nil
}

func buildStruct(value reflect.Value) *Struct {
	prototype := Struct{
		value:  value,
		fields: make(map[string]*Field),
	}
	return &prototype
}
