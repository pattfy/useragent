package platform

import (
	"regexp"
)

type linux struct {
	*core
}

func (l linux) ID() string {
	return "linux"
}

func (l linux) Name() string {
	return "Generic Linux"
}

func (l linux) Version() string {
	return ""
}

func (l linux) match() bool {
	reg := regexp.MustCompile("(?i)Linux")
	return reg.Match([]byte(l.UA))
}
