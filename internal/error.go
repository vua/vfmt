package internal

import "errors"

const (
	SyntaxERR int=1<<iota
	Unknown
)

func Err(code int) error{
	var err error
	switch code {
	case SyntaxERR:
		err=errors.New("v_fmt syntax error")
	case Unknown:
		err=errors.New("v_fmt unknown error")
	}
	return err
}
