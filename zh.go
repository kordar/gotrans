package gotrans

import (
	"github.com/go-playground/locales"
	"github.com/go-playground/locales/zh"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	zhtranslations "github.com/go-playground/validator/v10/translations/zh"
)

type ZhTranslation struct {
	translator locales.Translator
}

func NewZhTranslation() *ZhTranslation {
	return &ZhTranslation{translator: zh.New()}
}

func (z ZhTranslation) GetTranslator() locales.Translator {
	return z.translator
}

func (z ZhTranslation) RegisterTranslatorAndValidate(trans ut.Translator, validate *validator.Validate) error {
	return zhtranslations.RegisterDefaultTranslations(validate, trans)
}
