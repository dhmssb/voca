package util

import "github.com/go-playground/validator/v10"

// GetValidCurrencyValidator returns the ValidCurrency validator.Func.
func GetValidCurrencyValidator() validator.Func {
	return ValidCurrency
}

// ValidCurrency is your custom currency validation function.
func ValidCurrency(fl validator.FieldLevel) bool {
	if currency, ok := fl.Field().Interface().(string); ok {
		return IsSupportedCurrency(currency)
	}
	return false
}

func CustomCurrencyValidator(fl validator.FieldLevel) bool {
	if currency, ok := fl.Field().Interface().(string); ok {
		return IsSupportedCurrency(currency)
	}
	return false
}

// func GetValidCurrencyValidator() validator.Func {
// 	return func(fl validator.FieldLevel) bool {
// 		if currency, ok := fl.Field().Interface().(string); ok {
// 			return IsSupportedCurrency(currency)
// 		}
// 		return false
// 	}
// }
