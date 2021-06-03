package valuebinder

import (
	"net/url"
	"reflect"
	"time"

	"github.com/bcowtech/structproto/internal"
)

var (
	typeOfDuration = reflect.TypeOf(time.Nanosecond)
	typeOfUrl      = reflect.TypeOf(url.URL{})
	typeOfTime     = reflect.TypeOf(time.Time{})
)

func BuildIgnoreBinder(rv reflect.Value) internal.ValueBinder { return nil }
