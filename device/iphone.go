package device

import (
	"regexp"
)

type iPhone struct {
	*core
}

func (i iPhone) ID() string {
	return "iphone"
}

func (i iPhone) Name() string {
	return "iPhone"
}

func (i iPhone) match() bool {
	return regexp.MustCompile("iPhone").Match([]byte(i.ua))
}
