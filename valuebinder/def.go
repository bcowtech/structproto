package valuebinder

import (
	"net/url"
	"reflect"
	"time"

	"github.com/bcowtech/structproto/internal"
	"github.com/bcowtech/types"
)

var (
	typeOfDuration   = reflect.TypeOf(time.Nanosecond)
	typeOfUrl        = reflect.TypeOf(url.URL{})
	typeOfTime       = reflect.TypeOf(time.Time{})
	typeOfRawContent = reflect.TypeOf(types.RawContent(nil))
)

var _ internal.ValueBindProvider = BuildIgnoreBinder

func BuildIgnoreBinder(rv reflect.Value) internal.ValueBinder { return nil }
