package errors

import "github.com/pkg/errors"

func Wrapped(msg string) func(err error) error {
  return func(err error) error {
    return errors.Wrap(err, msg)
  }
}

func New(msg string) error {
  return errors.New(msg)
}
