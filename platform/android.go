package platform

import (
	"regexp"
)

type android struct {
	*core
}

func (a android) ID() string {
	return "android"
}

func (a android) Name() string {
	return "Android"
}

func (a android) Version() string {
	reg := regexp.MustCompile("(?i)Android ([\\d.]+)")
	m := reg.FindSubmatch([]byte(a.UA))
	if len(m) > 1 {
		return string(m[1])
	}
	return ""
}

func (a android) match() bool {
	reg := regexp.MustCompile("(?i)Android")
	return reg.Match([]byte(a.UA))
}
