package validator

import (
	"fmt"
	"reflect"
	"strings"

	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
)

var (
	validate *validator.Validate
)

// Init 初始化验证器
func Init() {
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		validate = v

		// 注册自定义验证
		validate.RegisterTagNameFunc(func(fld reflect.StructField) string {
			name := strings.SplitN(fld.Tag.Get("json"), ",", 2)[0]
			if name == "-" {
				return ""
			}
			return name
		})
	}
}

// GetError 获取验证错误
func GetError(err error) string {
	if errs, ok := err.(validator.ValidationErrors); ok {
		for _, e := range errs {
			return fmt.Sprintf("%s %s", e.Field(), getMessage(e))
		}
	}
	return err.Error()
}

// getMessage 获取验证消息
func getMessage(e validator.FieldError) string {
	switch e.Tag() {
	case "required":
		return "不能为空"
	case "email":
		return "格式不正确"
	case "min":
		return fmt.Sprintf("长度不能小于%s", e.Param())
	case "max":
		return fmt.Sprintf("长度不能大于%s", e.Param())
	case "len":
		return fmt.Sprintf("长度必须为%s", e.Param())
	default:
		return "格式不正确"
	}
}