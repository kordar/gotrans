package gotrans

import (
	"github.com/go-playground/locales"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
)

type GoTranslation interface {
	GetTranslator() locales.Translator
	RegisterTranslatorAndValidate(trans ut.Translator, validate *validator.Validate) error
}
