package main

import (
	"errors"
	"time"
)

func intToMonth(num int) (month string, err error) {
	if num > 12 || num < 1 {
		return "", errors.New("error: nummber must be between 1 and 12")
	}

	month = time.Month(num).String()

	return month, nil
}
