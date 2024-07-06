package gotrans

import (
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	logger "github.com/kordar/gologger"
)

type Trans struct {
	uni         *ut.UniversalTranslator
	validate    *validator.Validate
	translators map[string]ut.Translator
}

// NewTrans 创建翻译组件，并注册为默认组件
func NewTrans(validate *validator.Validate) *Trans {
	// 使用英语作为默认组件
	en := NewEnTranslation()
	translator := en.GetTranslator()
	uni := ut.New(translator)
	translators, _ := uni.GetTranslator(translator.Locale())
	_ = en.RegisterTranslatorAndValidate(translators, validate)
	return &Trans{uni: uni, validate: validate, translators: map[string]ut.Translator{translators.Locale(): translators}}
}

func (t *Trans) RegisterTranslators(translators ...GoTranslation) *Trans {
	for _, translation := range translators {
		// 获取原始translator
		translator := translation.GetTranslator()
		// 将translator添加到国际化组件中
		err := t.uni.AddTranslator(translator, true)
		if err != nil {
			logger.Errorf("failed to add \"translator\", err=%v", err)
			continue
		}
		locale := translator.Locale()
		// 获取添加的国际化组件，向该组件注册validate组件，完成locale对绑定的国际化初始化
		if target, found := t.GetTranslator(locale); found {
			t.translators[locale] = target
			_ = translation.RegisterTranslatorAndValidate(target, t.validate)
		}
	}
	return t
}

// GetTranslator 通过en、ar、zh等获取本地组件
func (t *Trans) GetTranslator(locale string) (trans ut.Translator, found bool) {
	return t.uni.GetTranslator(locale)
}

func (t *Trans) GetValidate() *validator.Validate {
	return t.validate
}

func (t *Trans) BindTranslatorToValidate(tag string, registerFn func(locale string) (string, bool), translationFn func(locale string, fe validator.FieldError) []string) *Trans {

	for locale, translator := range t.translators {
		// 向validate注册国际化，tag指向不同国际化，及其处理函数
		err := t.validate.RegisterTranslation(tag, translator, func(ut ut.Translator) error {
			text, override := registerFn(locale)
			return ut.Add(tag, text, override) // see universal-translator for details
		}, func(ut ut.Translator, fe validator.FieldError) string {
			tt := translationFn(locale, fe)
			value, err := ut.T(tag, tt...)
			if err != nil {
				logger.Warnf("translation function execution failed, fe = %s, err = %v", fe.Field(), err)
			}
			return value
		})

		if err != nil {
			logger.Errorf("unable to register translator for locale value '%s', err=%v", locale, err)
		}
	}

	return t
}
