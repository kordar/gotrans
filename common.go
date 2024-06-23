package gotrans

import (
	"github.com/kordar/gocfg"
	logger "github.com/kordar/gologger"
	"github.com/kordar/govalidator"
)

var (
	translations *Translations
	i18nPkg      = "language"
)

func SetI18nPkg(pkg string) {
	i18nPkg = pkg
}

func GetI18nPkg() string {
	return i18nPkg
}

func GetTranslations() *Translations {
	return translations
}

func InitValidateAndTranslations(tr ...ITranslation) {
	validate := govalidator.GetValidate()
	if validate == nil {
		logger.Fatal("please load the \"validate\" object first!")
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

func GetSectionValue(locale string, section string, key string) string {
	return gocfg.GetSectionValue(locale, section+"."+key, GetI18nPkg())
}

func GetDictValue(locale string, key string) string {
	return gocfg.GetSectionValue(locale, "dictionary."+key, GetI18nPkg())
}
