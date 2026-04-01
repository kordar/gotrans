package gotrans

import (
	"log/slog"
	"os"

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
		slog.Error("please load the \"validate\" object first!")
		os.Exit(1)
		return
	}
	translations = NewTrans(validate).RegisterTranslators(tr...)
}
