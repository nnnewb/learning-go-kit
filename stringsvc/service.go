package stringsvc

import (
	"strings"
)

type StringService interface {
	Uppercase(string) (string, error)
	Count(string) int
}

type StringServiceImpl struct{}

func (s *StringServiceImpl) Uppercase(arg string) (string, error) {
	if arg == "" {
		return "", nil
	}
	return strings.ToUpper(arg), nil
}

func (s *StringServiceImpl) Count(arg string) int {
	return len(arg)
}
