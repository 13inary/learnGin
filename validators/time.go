package validators

import (
	"time"

	"github.com/go-playground/validator/v10"
)

func BookableDate(f1 validator.FieldLevel) bool {

	// get value of parameter
	if date, ok := f1.Field().Interface().(time.Time); ok {
		today := time.Now()
		if today.Unix() > date.Unix() {
			return false
		}
	}

	return true

}
