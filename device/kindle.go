package device

import (
	"regexp"
)

type kindle struct {
	*core
}

func (i kindle) ID() string {
	return "kindle"
}

func (i kindle) Name() string {
	return "Kindle"
}

func (i kindle) match() bool {
	return regexp.MustCompile("(?i)Kindle").Match([]byte(i.ua))
}
