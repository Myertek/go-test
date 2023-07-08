package db

import (
	"errors"
	"fmt"
)

func Hello(name string) (string, error) {
	//if no name given return error
	if name == "" {
		return "", errors.New("Name cannot be empty!")
	}
	msg := fmt.Sprintf("Hi %v, Welcom from DB!!", name)
	return msg, nil
}
