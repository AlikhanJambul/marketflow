package apperrors

import "errors"

var (
	ErrInvalidSymbol   = errors.New("Symbol is not valid")
	ErrInvalidExchange = errors.New("Exchange is not valid")
	ErrInavalidBody    = errors.New("Symbol or Exchange is not valid")
	ErrRedis           = errors.New("Redis doesn't work")
	ErrRedisNil        = errors.New("Redis is nil")
	ErrDB              = errors.New("DB doesn't work")
	ErrDuration        = errors.New("Period is invalid")
)

func CheckCode(err error) int {
	if err == ErrInvalidSymbol || err == ErrInvalidExchange || err == ErrInavalidBody || err == ErrDuration {
		return 400
	}

	if err == ErrRedisNil {
		return 404
	}

	return 500
}
