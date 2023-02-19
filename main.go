package main

import (
	"errors"
	"fmt"
	"golang.org/x/xerrors"
)

type XError struct {
	e error
	s string
	f xerrors.Frame
}

func (i *XError) Error() string {
	return i.s
}

func (i *XError) Unwrap() error {
	return i.e
}

func (i *XError) Format(f fmt.State, c rune) { // implements fmt.Formatter
	xerrors.FormatError(i, f, c)
}

func (i *XError) FormatError(p xerrors.Printer) error { // implements xerrors.Formatter
	p.Printf("%s\n", i.s)
	if p.Detail() {
		i.f.Format(p)
	}

	wrapped := i.Unwrap()
	if wrapped == nil {
		return nil
	}

	wrappedInternal := &XError{}
	if !errors.As(wrapped, &wrappedInternal) {
		p.Print(wrapped.Error())
		return nil // TODO need to continue to unwrap other errors
	}

	return wrappedInternal.FormatError(p)
}

func Wrap(err error, s string) error {
	return &XError{e: err, s: s, f: xerrors.Caller(1)}
}

func main() {
	err := callA()

	if err != nil {
		fmt.Printf("%+v\n", Wrap(err, "call A"))
	}
}

func callA() error {
	err := callB()

	return Wrap(err, "call B")

}

func callB() error {
	err := errors.New("external service is down")

	return Wrap(err, "call external service")
}
