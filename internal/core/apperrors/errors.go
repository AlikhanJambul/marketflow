package apperrors

import "errors"

var (
	ErrInvalidSymbol   = errors.New("Symbol is not valid")
	ErrInvalidExchange = errors.New("Exchange is not valid")
)

func CheckCode(err error) int {
	if err == ErrInvalidSymbol || err == ErrInvalidExchange {
		return 400
	}
	return 500
}
