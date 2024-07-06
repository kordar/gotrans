# gotrans

## 安装
```go
go get github.com/kordar/gotrans v0.1.0
```

## 自定义多语言

自定义组件接口定义

```go
type GoTranslation interface {
	GetTranslator() locales.Translator
	RegisterTranslatorAndValidate(trans ut.Translator, validate *validator.Validate) error
}
```

实现GoTranslation
```go
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
```

## 组件集成使用

- 初始化组件

```go
func Init(tr ...gotrans.GoTranslation) {
	gotrans.Initialize(tr...)
}
```

- 集成[`govalidator`](https://github.com/kordar/govalidator)

```go
func AddValidate(validations ...govalidator.IValidation) {
	for _, validation := range validations {
		govalidator.AddValidation(validation)
		if !gotrans.Exists() {
			continue
		}
		trans := gotrans.Get()
		trans.BindTranslatorToValidate(
			validation.Tag(),
			func(locale string) (string, bool) {
				defaultTpl, override := validation.DefaultTpl()
				section, key := validation.Tpl()
				if section == "" || key == "" {
					return defaultTpl, override
				}
				sk := fmt.Sprintf("%s.%s", section, key)
				value := gocfg.GetSectionValue(locale, sk, "language")
				if value == "" {
					return defaultTpl, override
				} else {
					return value, override
				}
			},
			func(locale string, fe validator.FieldError) []string {
				n := validation.I18n(fe, locale)
				if n == nil || len(n) == 0 {
					text := gocfg.GetSectionValue(locale, "dictionary."+fe.StructNamespace(), "language")
					if text == "" {
						text = fe.Field()
					}
					return []string{text}
				}
				//logger.Infof("=========field======%+v", fe.Field())
				//logger.Infof("=========param======%+v", fe.Param())
				//logger.Infof("=========tag======%+v", fe.Tag())
				//logger.Infof("=========error======%+v", fe.Error())
				//logger.Infof("=========StructField======%+v", fe.StructField())
				//logger.Infof("=========Namespace======%+v", fe.Namespace())
				//logger.Infof("=========StructNamespace======%+v", fe.StructNamespace())
				//logger.Infof("=========ActualTag======%+v", fe.ActualTag())
				return n
			})
	}
}
```


## 常用国家语言标识

语言标识|国家地区  
---|:---  
zh_CN  |  简体中文(中国)  
zh_TW  |  繁体中文(台湾地区)  
zh_HK  |  繁体中文(香港)  
en_HK  |  英语(香港)  
en_US  |  英语(美国)  
en_GB  |  英语(英国)  
en_WW  |  英语(全球)  
en_CA  |  英语(加拿大)  
en_AU  |  英语(澳大利亚)  
en_IE  |  英语(爱尔兰)  
en_FI  |  英语(芬兰)  
fi_FI  |  芬兰语(芬兰)  
en_DK  |  英语(丹麦)  
da_DK  |  丹麦语(丹麦)  
en_IL  |  英语(以色列)  
he_IL  |  希伯来语(以色列)  
en_ZA  |  英语(南非)  
en_IN  |  英语(印度)  
en_NO  |  英语(挪威)  
en_SG  |  英语(新加坡)  
en_NZ  |  英语(新西兰)  
en_ID  |  英语(印度尼西亚)  
en_PH  |  英语(菲律宾)  
en_TH  |  英语(泰国)  
en_MY  |  英语(马来西亚)  
en_XA  |  英语(阿拉伯)  
ko_KR  |  韩文(韩国)  
ja_JP  |  日语(日本)  
nl_NL  |  荷兰语(荷兰)  
nl_BE  |  荷兰语(比利时)  
pt_PT  |  葡萄牙语(葡萄牙)  
pt_BR  |  葡萄牙语(巴西)  
fr_FR  |  法语(法国)  
fr_LU  |  法语(卢森堡)  
fr_CH  |  法语(瑞士)  
fr_BE  |  法语(比利时)  
fr_CA  |  法语(加拿大)  
es_LA  |  西班牙语(拉丁美洲)  
es_ES  |  西班牙语(西班牙)  
es_AR  |  西班牙语(阿根廷)  
es_US  |  西班牙语(美国)  
es_MX  |  西班牙语(墨西哥)  
es_CO  |  西班牙语(哥伦比亚)  
es_PR  |  西班牙语(波多黎各)  
de_DE  |  德语(德国)  
de_AT  |  德语(奥地利)  
de_CH  |  德语(瑞士)  
ru_RU  |  俄语(俄罗斯)  
it_IT  |  意大利语(意大利)  
el_GR  |  希腊语(希腊)  
no_NO  |  挪威语(挪威)  
hu_HU  |  匈牙利语(匈牙利)  
tr_TR  |  土耳其语(土耳其)  
cs_CZ  |  捷克语(捷克共和国)  
sl_SL  |  斯洛文尼亚语   
pl_PL  |  波兰语(波兰)  
sv_SE  |  瑞典语(瑞典)  
es_CL  |  西班牙语 (智利)  