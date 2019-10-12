package platform

import (
	"regexp"
	"strings"
)

type mac struct {
	*core
}

func (m mac) ID() string {
	return "mac"
}

func (m mac) Name() string {
	return "Mac"
}

func (m mac) Version() string {
	reg := regexp.MustCompile("(?i)Mac OS X\\s*([0-9_\\.]+)?")
	sub := reg.FindSubmatch([]byte(m.UA))
	if len(sub) > 1 {
		return strings.Replace(string(sub[1]), "_", ".", -1)
	}
	return ""
}

func (m mac) match() bool {
	reg := regexp.MustCompile("(?i)Mac")
	return reg.Match([]byte(m.UA))
}
