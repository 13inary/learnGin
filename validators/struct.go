package validators

import (
	"github.com/go-playground/validator/v10"
)

func NameNotNullAndAdmin(f1 validator.FieldLevel) bool {

	if v, f := f1.Field().Interface().(string); f {
		return v != "" && !("admin" == v)
	}

	return true
}
