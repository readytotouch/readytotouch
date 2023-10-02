package oauth_providers

import "errors"

var (
	ErrUnexpectedHttpStatusCode = errors.New("unexpected http status code")
	ErrAbsentEmail              = errors.New("absent email")
	ErrUnverifiedEmail          = errors.New("unverified email")
	ErrAbsentID                 = errors.New("absent id")
	ErrAbsentUsername           = errors.New("absent username")
)
