package validation

import (
	"fmt"
	"go-gin/pkg/secret"

	"github.com/go-playground/locales/en"
	"github.com/go-playground/locales/id"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	en_translations "github.com/go-playground/validator/v10/translations/en"
	id_translations "github.com/go-playground/validator/v10/translations/id"
)

var (
	Translator ut.Translator
)

func Setup() error {
	locale := "en"
	fmt.Println("locale", locale, secret.APP_LOCALE)
	uni := ut.New(
		en.New(),
		id.New(),
		en.New(),
		id.New(),
	)

	var found bool
	Translator, found = uni.GetTranslator(secret.APP_LOCALE)
	if !found {
		return fmt.Errorf("translator for locale %s not found", secret.APP_LOCALE)
	}

	validate := validator.New()

	switch secret.APP_LOCALE {
	case "en":
		return en_translations.RegisterDefaultTranslations(validate, Translator)
	case "id":
		return id_translations.RegisterDefaultTranslations(validate, Translator)
	default:
		return en_translations.RegisterDefaultTranslations(validate, Translator)
	}
}
