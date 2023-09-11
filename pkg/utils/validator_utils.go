package utils

import (
	"errors"
	"github.com/go-playground/locales/fa_IR"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	fa_translations "github.com/go-playground/validator/v10/translations/fa"
	"github.com/gofiber/fiber/v2"
)

var (
	uni      *ut.UniversalTranslator
	validate *validator.Validate
)

func Validator(ctx *fiber.Ctx, req interface{}) (error, any) {
	fa := fa_IR.New()
	uni = ut.New(fa, fa)

	trans, _ := uni.GetTranslator("fa")

	validate = validator.New()
	fa_translations.RegisterDefaultTranslations(validate, trans)

	ctx.BodyParser(&req)
	err := validate.Struct(req)

	if err != nil {
		var errs validator.ValidationErrors
		errors.As(err, &errs)

		return err, errs.Translate(trans)
	}

	return nil, nil
}
