package gotrans

import (
	"github.com/go-playground/locales"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	"github.com/kordar/gologger"
	"regexp"
	"strings"
)

type ITranslation interface {
	GetTranslator() locales.Translator
	RegisterValidate(trans ut.Translator, validate *validator.Validate) error
}

type Translations struct {
	uni      *ut.UniversalTranslator
	validate *validator.Validate
	trans    map[string]ut.Translator
}

// NewTranslations 创建翻译组件，并注册为默认组件
func NewTranslations(validate *validator.Validate) *Translations {
	// 使用英语作为默认组件
	en := NewEnTranslation()
	translator := en.GetTranslator()
	uni := ut.New(translator)
	trans, _ := uni.GetTranslator(translator.Locale())
	_ = en.RegisterValidate(trans, validate)
	return &Translations{uni: uni, validate: validate, trans: map[string]ut.Translator{trans.Locale(): trans}}
}

func (t *Translations) RegisterTranslators(translators ...ITranslation) *Translations {
	for i := range translators {
		translation := translators[i]
		translator := translation.GetTranslator()
		err := t.uni.AddTranslator(translator, true)
		if err != nil {
			logger.Errorf("failed to add \"translator\", err=%v", err)
			continue
		}
		locale := translator.Locale()
		if trans, found := t.GetTrans(locale); found {
			t.trans[locale] = trans
			_ = translation.RegisterValidate(trans, t.validate)
		}
	}
	return t
}

// GetTrans 通过en、ar、zh等获取本地组件
func (t *Translations) GetTrans(locale string) (trans ut.Translator, found bool) {
	return t.uni.GetTranslator(locale)
}

func (t *Translations) GetValidate() *validator.Validate {
	return t.validate
}

func (t *Translations) RegisterTranslationWithGI18n(tag string, section string, key string) *Translations {

	for locale, trans := range t.trans {

		regFn := func(ut ut.Translator) error {
			text := GetSectionValue(locale, section, key)
			return ut.Add(tag, text, true) // see universal-translator for details
		}

		transFn := func(ut ut.Translator, fe validator.FieldError) string {

			if fe.Param() != "" && strings.Contains(fe.Param(), "msg(") {
				// 定义一个正则表达式来匹配圆括号中的内容
				re := regexp.MustCompile(`msg\((.*?)\)`)
				// 使用正则表达式查找匹配项
				matches := re.FindAllStringSubmatch(fe.Param(), -1)
				if len(matches) > 0 {
					return matches[0][1]
				}
			}

			text := GetDictValue(locale, fe.Field())
			if text == "" {
				text = fe.Field()
			}
			t2, _ := ut.T(tag, text)
			return t2
		}

		err := t.validate.RegisterTranslation(tag, trans, regFn, transFn)
		if err != nil {
			logger.Errorf("unable to register translator for locale value '%s', err=%v", locale, err)
		}
	}

	return t
}
