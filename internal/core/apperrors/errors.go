package apperrors

import "errors"

var (
	ErrInvalidSymbol   = errors.New("Symbol is not valid")
	ErrInvalidExchange = errors.New("Exchange is not valid")
	ErrInavalidBody    = errors.New("Symbol or Exchange is not valid")
)

func CheckCode(err error) int {
	if err == ErrInvalidSymbol || err == ErrInvalidExchange || err == ErrInavalidBody {
		return 400
	}
	return 500
}
