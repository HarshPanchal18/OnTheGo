package main

import "errors"

func createError(str string) error {
	return errors.New(str)
}