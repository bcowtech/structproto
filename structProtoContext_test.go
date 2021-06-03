package structproto

import (
	"reflect"
	"testing"

	"github.com/bcowtech/structproto/valuebinder"
)

func TestStructProtoContext(t *testing.T) {
	c := mockCharacter{}
	prototype, err := Prototypify(&c, &StructProtoOption{
		TagName:             "demo",
		ValueBinderProvider: valuebinder.BuildStringArgsBinder,
	})
	if err != nil {
		t.Error(err)
	}

	context, err := buildStructProtoContext(prototype)
	if err != nil {
		t.Error(err)
	}

	expectedFields := []string{"NAME", "AGE", "ALIAS", "DATE_OF_BIRTH", "REMARK", "NUMBERS"}
	if !reflect.DeepEqual(expectedFields, context.AllFields()) {
		t.Errorf("assert 'structprotoContext.AllFields()':: expected '%#v', got '%#v'", expectedFields, context.AllFields())
	}
	expectedRequiredFields := []string{"AGE", "NAME"}
	if !reflect.DeepEqual(expectedRequiredFields, context.AllRequiredFields()) {
		t.Errorf("assert 'structprotoContext.AllRequiredFields()':: expected '%#v', got '%#v'", expectedRequiredFields, context.AllRequiredFields())
	}

	{
		field := context.Field("NAME")
		if field == nil {
			t.Errorf("assert 'structprotoContext.Field(\"NAME\")':: expected not nil, got '%#v'", field)
		}
		if field.Name() != "NAME" {
			t.Errorf("assert 'structprotoField.Name()':: expected '%#v', got '%#v'", "NAME", field.Name())
		}
		if field.Index() != 0 {
			t.Errorf("assert 'structprotoField.Index()':: expected '%#v', got '%#v'", "NAME", field.Name())
		}
		expectedFlags := []string{"required"}
		if !reflect.DeepEqual(expectedFlags, field.Flags()) {
			t.Errorf("assert 'structprotoField.Flags()':: expected '%#v', got '%#v'", expectedFlags, field.Flags())
		}
	}

	if !context.IsRequireField("NAME") {
		t.Errorf("assert 'structprotoContext.IsRequireField(\"NAME\")':: expected '%#v', got '%#v'", expectedRequiredFields, context.IsRequireField("NAME"))
	}
	if context.IsRequireField("unknown") {
		t.Errorf("assert 'structprotoContext.IsRequireField(\"unknown\")':: expected '%#v', got '%#v'", expectedRequiredFields, context.IsRequireField("unknown"))
	}

	// TODO: test context.ChechIfMissingRequireFields
}
