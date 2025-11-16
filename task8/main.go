package main

import "fmt"

//type Error error

type MultiError struct {
	errors []error
}

func (e *MultiError) Error() string {
	if e.errors == nil || len(e.errors) == 0 {
		return ""
	}

	count := len(e.errors)
	errStr := fmt.Sprintf("%d errors occured:\n", count)

	for _, err := range e.errors {
		errStr += fmt.Sprintf("\t* %v", err.Error())
	}

	return errStr + "\n"
}

func Append(err error, errs ...error) *MultiError {
	switch err := err.(type) {
	case *MultiError:
		if err == nil {
			err = new(MultiError)
		}

		for _, e := range errs {
			switch e := e.(type) {
			case *MultiError:
				if e != nil {
					err.errors = append(err.errors, e.errors...)
				}
			default:
				if e != nil {
					err.errors = append(err.errors, e)
				}
			}
		}

		return err
	default:
		newErrs := make([]error, 0, len(errs)+1)
		if err != nil {
			newErrs = append(newErrs, err)
		}
		newErrs = append(newErrs, errs...)

		return Append(&MultiError{}, newErrs...)
	}
}
