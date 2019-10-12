package platform

import (
	"regexp"
)

type windows struct {
	*core
}

func (p windows) ID() string {
	return "windows"
}

func (p windows) Name() string {
	return "Windows"
}

func (p windows) Version() string {
	reg := regexp.MustCompile("(?i)Windows NT\\s*([0-9_\\.]+)?")
	m := reg.FindSubmatch([]byte(p.UA))
	if len(m) > 1 {
		return string(m[1])
	}
	return ""
}

func (p windows) match() bool {
	return regexp.MustCompile("(?i)Windows").Match([]byte(p.UA))
}
