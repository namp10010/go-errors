package main

import (
	"fmt"
	"os"

	"github.com/pkg/errors"
	"github.com/ztrue/tracerr"
	"golang.org/x/xerrors"
)

func main() {
	err := a()
	fmt.Printf("%+v\n", err)

	if err != nil {
		fmt.Printf("%+v\n", Wrap(err, "call A"))
	}

	if err := read(); err != nil {
		//tracerr.PrintSourceColor(err)
		tracerr.Print(err)
	}
	//err1 := a1()
	//
	//if err1 != nil {
	//	err1 := Wrap(err1, "call A1")
	//	fmt.Printf("%+v\n", err1)
	//}

	//panicky()
}

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
	//p.Printf("%s\n", i.s)
	if p.Detail() {
		i.f.Format(p)
	}

	wrapped := i.Unwrap()
	if wrapped == nil {
		return nil
	}

	err := &XError{}
	if !errors.As(wrapped, &err) {
		p.Print(wrapped.Error())
		return nil // TODO need to continue to unwrap other errors
	}

	return err.FormatError(p)
}

func Wrap(err error, s string) error {
	return &XError{e: err, s: s, f: xerrors.Caller(1)}
}

func panicky() {
	pannick2()
}

func pannick2() {
	panic("panic!")
}

func a1() error {
	err := b1()

	return errors.Wrap(err, "call B")
}

func b1() error {
	err := fmt.Errorf("external service is down")

	return errors.Wrap(err, "call external service")
}

func a() error {
	err := b()

	return Wrap(err, "call B")

}

func b() error {
	err := fmt.Errorf("external service is down")

	return Wrap(err, "call external service")
}
func read() error {
	return readNonExistent()
}

func readNonExistent() error {
	_, err := os.ReadFile("/tmp/non_existent_file")
	// Add stack trace to existing error, no matter if it's nil.
	return tracerr.Wrap(err)
}
