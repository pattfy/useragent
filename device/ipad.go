package device

import (
	"regexp"
)

type iPad struct {
	*core
}

func (i iPad) ID() string {
	return "ipad"
}

func (i iPad) Name() string {
	return "iPad"
}

func (i iPad) match() bool {
	return regexp.MustCompile("iPad").Match([]byte(i.ua))
}
