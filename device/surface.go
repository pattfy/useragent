package device

import (
	"regexp"

	"github.com/pattfy/useragent/platform"
)

type surface struct {
	*core
}

func (s surface) ID() string {
	return "surface"
}

func (s surface) Name() string {
	return "Microsoft Surface"
}

func (s surface) match() bool {
	return platform.IsWindows(s.ua) && regexp.MustCompile("Touch").Match([]byte(s.ua))
}
