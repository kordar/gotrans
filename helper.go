package gotrans

import (
	logger "github.com/kordar/gologger"
	"github.com/kordar/govalidator"
)

var (
	translations *Trans
)

func Get() *Trans {
	return translations
}

func Exists() bool {
	return Get() != nil
}

// Initialize 初始化翻译组件，参数注册函数、翻译函数，翻译组件
func Initialize(tr ...GoTranslation) {
	validate := govalidator.GetValidate()
	if validate == nil {
		logger.Fatal("please load the \"validate\" object first!")
		return
	}
	translations = NewTrans(validate).RegisterTranslators(tr...)
}
