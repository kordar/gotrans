package gotrans

import (
	"github.com/go-playground/locales"
	"github.com/go-playground/locales/en"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	entranslations "github.com/go-playground/validator/v10/translations/en"
)

type EnTranslation struct {
	translator locales.Translator
}

func NewEnTranslation() *EnTranslation {
	return &EnTranslation{translator: en.New()}
}

func (e EnTranslation) GetTranslator() locales.Translator {
	return e.translator
}

func (e EnTranslation) RegisterTranslatorAndValidate(trans ut.Translator, validate *validator.Validate) error {
	return entranslations.RegisterDefaultTranslations(validate, trans)
}
