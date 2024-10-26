package utils

import "errors"

func SizeChecker(width, height int) error {
	if width <= 0 || height <= 0 {
		return errors.New("width and height must be greater than 0")
	}
	return nil
}
