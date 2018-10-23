package errors

import "github.com/pkg/errors"

func New(msg string) error {
  return errors.New(msg)
}

func Wrap(err error, msg string) error {
  return errors.Wrap(err, msg)
}

func Wrapped(msg string) func(err error) error {
  return func(err error) error {
    return errors.Wrap(err, msg)
  }
}
