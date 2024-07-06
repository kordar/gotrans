package gotrans

import (
	"github.com/go-playground/locales"
	"github.com/go-playground/locales/ar"
	"github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	artranslations "github.com/go-playground/validator/v10/translations/ar"
)

type ArTranslation struct {
	translator locales.Translator
}

func NewArTranslation() *ArTranslation {
	return &ArTranslation{translator: ar.New()}
}

func (z ArTranslation) GetTranslator() locales.Translator {
	return z.translator
}

func (z ArTranslation) RegisterTranslatorAndValidate(trans ut.Translator, validate *validator.Validate) error {
	return artranslations.RegisterDefaultTranslations(validate, trans)
}
