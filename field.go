package structproto

import "reflect"

var _ FieldInfo = new(Field)

type Field struct {
	name  string
	desc  string
	index int
	flags FieldFlagSet
	tag   reflect.StructTag
}

func (info *Field) Name() string {
	return info.name
}

func (info *Field) Desc() string {
	return info.desc
}

func (info *Field) Index() int {
	return info.index
}

func (info *Field) Flags() []string {
	return info.flags
}

func (info *Field) HasFlag(v string) bool {
	return info.flags.Has(v)
}

func (info *Field) Tag() reflect.StructTag {
	return info.tag
}

func (info *Field) appendFlags(values ...string) {
	if len(values) > 0 {
		for _, v := range values {
			if len(v) == 0 {
				continue
			}
			info.flags.Append(v)
		}
	}
}
