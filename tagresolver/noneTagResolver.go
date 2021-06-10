package tagresolver

import "github.com/bcowtech/structproto/internal"

var _ internal.TagResolver = NoneTagResolver

func NoneTagResolver(fieldname, token string) (*internal.Tag, error) {
	var tag = &internal.Tag{
		Name: fieldname,
	}
	return tag, nil
}
