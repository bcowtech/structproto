package structproto

import "reflect"

type StructProtoContext Struct

func (ctx *StructProtoContext) Target() reflect.Value {
	return ctx.target
}

func (ctx *StructProtoContext) Field(name string) FieldInfo {
	if field, ok := ctx.fields[name]; ok {
		return field
	}
	return nil
}

func (ctx *StructProtoContext) AllFields() []string {
	var fields []string = make([]string, len(ctx.fields))
	for _, v := range ctx.fields {
		fields[v.index] = v.name
	}
	return fields
}

func (ctx *StructProtoContext) AllRequiredFields() []string {
	return ctx.requiredFields
}

func (ctx *StructProtoContext) IsRequiredField(name string) bool {
	field := ctx.Field(name)
	if field != nil {
		return field.HasFlag(RequiredFlag)
	}
	return false
}

func (ctx *StructProtoContext) CheckIfMissingRequiredFields(fieldVisitFunc func() <-chan string) error {
	if ctx.requiredFields.IsEmpty() {
		return nil
	}

	var requiredFields = ctx.requiredFields.Clone()

	for field := range fieldVisitFunc() {
		index := requiredFields.IndexOf(field)
		if index != -1 {
			requiredFields.RemoveIndex(index)
		}
	}

	if !requiredFields.IsEmpty() {
		field, _ := requiredFields.Get(0)
		return &MissingRequiredFieldError{field, nil}
	}
	return nil
}

func buildStructProtoContext(prototype *Struct) (*StructProtoContext, error) {
	context := StructProtoContext(*prototype)
	return &context, nil
}
