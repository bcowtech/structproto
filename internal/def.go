package internal

const (
	RequiredFlag = "required"
)

type (
	TagResolver func(fieldname, token string) (*Tag, error)

	ValueBinder interface {
		Bind(v interface{}) error
	}

	Tag struct {
		Name  string
		Flags []string
		Desc  string
	}
)
