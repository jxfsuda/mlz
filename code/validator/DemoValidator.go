package validator

import (
	"github.com/gin-gonic/gin/binding"
	"gopkg.in/go-playground/validator.v8"
	"reflect"
	"strings"
)

func init() {
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		v.RegisterValidation("DemoName", DemoName)
	}

}


func DemoName(
	v *validator.Validate, topStruct reflect.Value, currentStructOrField reflect.Value,
	field reflect.Value, fieldType reflect.Type, fieldKind reflect.Kind, param string,
) bool {
	if name, ok := field.Interface().(string); ok {
		if strings.Contains(name,"j"){  //自定义验证,如果name包含j ,name验证不通过

			return false
		}
	}
	return true
}