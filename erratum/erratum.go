package erratum

import (
	"errors"
	"fmt"
)

func Use(opener ResourceOpener, input string) (err error) {
	r, err := opener()
	if err != nil {
		if _, ok := err.(TransientError); ok {
			return Use(opener, input)
		} else {
			return err
		}
	}

	defer func() {
		if e := recover(); e != nil {
			if err, ok := e.(FrobError); ok {
				r.Defrob(err.defrobTag)
			}
			err = errors.New(fmt.Sprintf("%v", e))
			_ = r.Close()
		}
	}()

	r.Frob(input)
	err = r.Close()
	if err != nil {
		return err
	}

	return err
}
