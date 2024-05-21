package gotrans

import (
	logger "github.com/kordar/gologger"
	"github.com/kordar/govalidator"
)

var translations *Translations

func GetTranslations() *Translations {
	return translations
}

func InitValidateAndTranslations(tr ...ITranslation) {
	validate := govalidator.GetValidate()
	if validate == nil {
		logger.Fatal("请先加载validate句柄")
		return
	}
	translations = NewTranslations(validate).RegisterTranslators(tr...)
}

func RegTranslations(tag string, section string, key string) {
	translations.RegisterTranslationWithGI18n(tag, section, key)
}

func RegValidation(valid govalidator.IValidation) {
	govalidator.AddValidation(valid)
	if section, key := valid.I18n(); section != "" && key != "" {
		translations.RegisterTranslationWithGI18n(valid.Tag(), section, key)
	}
}
