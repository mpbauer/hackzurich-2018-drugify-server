package validators

import (
	"reflect"
	"sync"


	"github.com/gin-gonic/gin/binding"
	"gopkg.in/go-playground/validator.v9"
)

type DefaultV9Validator struct {
	once     sync.Once
	validate *validator.Validate
}

var _ binding.StructValidator = &DefaultV9Validator{}

func (v *DefaultV9Validator) ValidateStruct(obj interface{}) error {

	if kindOfData(obj) == reflect.Struct {

		v.lazyinit()

		if err := v.validate.Struct(obj); err != nil {
			return error(err)
		}
	}

	return nil
}

func (v *DefaultV9Validator) Engine() interface{} {
	v.lazyinit()
	return v.validate
}

func (v *DefaultV9Validator) lazyinit() {
	v.once.Do(func() {
		v.validate = validator.New()
		v.validate.SetTagName("binding")

		// add any custom validations etc. here
	})
}

func kindOfData(data interface{}) reflect.Kind {

	value := reflect.ValueOf(data)
	valueType := value.Kind()

	if valueType == reflect.Ptr {
		valueType = value.Elem().Kind()
	}
	return valueType
}