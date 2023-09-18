package infrastructure

import (
	"github.com/go-playground/locales/en"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	en_translations "github.com/go-playground/validator/v10/translations/en"
)

func InitValidation() (*validator.Validate, ut.Translator) {
	v := validator.New()
	en := en.New()
	uni := ut.New(en, en)
	trans, _ := uni.GetTranslator("id")
	_ = en_translations.RegisterDefaultTranslations(v, trans)

	return v, trans
}
