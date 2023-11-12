package util

import "errors"

const (
	USD = "USD"
	IDR = "IDR"

	AuthorizationHeaderKey  = "authorization"
	AuthorizationTypeBearer = "bearer"
	AuthorizationPayloadKey = "authorization_payload"
)

var (
	ErrExpiredToken = errors.New("token has expired")
	ErrInvalidToken = errors.New("token is invalid")
)

func IsSupportedCurrency(currency string) bool {
	switch currency {
	case USD, IDR:
		return true
	}

	return false
}
