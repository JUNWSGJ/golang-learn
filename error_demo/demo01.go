package main

import (
	"database/sql"
	"fmt"
	"github.com/pkg/errors"
)

func foo() error {
	return errors.Wrap(sql.ErrNoRows, "foo failed")
}

func bar() error {
	return errors.WithMessage(foo(), "bar failed")
}

func main() {
	err := bar()
	if errors.Cause(err) == sql.ErrNoRows {
		fmt.Println("errors.Cause(err) == sql.ErrNoRows")
		fmt.Printf("%+v\n", err)
	}
	if errors.Is(err, sql.ErrNoRows) {
		fmt.Println("errors.Is(err, sql.ErrNoRows)")
		fmt.Printf("%+v\n", err)
	}

	if err != nil {
		// unknown error
		fmt.Errorf("unknown error, %+v", err)
	}
}
