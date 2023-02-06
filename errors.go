package glinet

import "errors"

type Code int

const (
	CodeBadToken Code = -1
)

var ErrUnauthorized = errors.New("unauthorized")
var ErrUnexpected = errors.New("unexpected")
