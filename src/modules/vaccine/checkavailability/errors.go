package checkavailability

import "errors"

var (
	ErrNumRequired   = errors.New("NUM_REQUIRED")
	ErrPinInvalid    = errors.New("PIN_INVALID")
	ErrPinLenInvalid = errors.New("PIN_LEN_INVALID")
)
