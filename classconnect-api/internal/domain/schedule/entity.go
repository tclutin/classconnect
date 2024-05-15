package schedule

import "errors"

var (
	ErrEvenGroup     = errors.New("parity error")
	ErrAlreadyExists = errors.New("schedule already exists")
	ErrNotExists     = errors.New("schedule not exists")
	ErrDaysCount     = errors.New("schedule has a lot of days or less")
)
