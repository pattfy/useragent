package platform

import (
	"regexp"
	"strings"
)

type ios struct {
	*core
}

func (i ios) ID() string {
	return "ios"
}

func (i ios) Name() string {
	return "iOS"
}

func (i ios) Version() string {
	reg := regexp.MustCompile("(?i)OS ([0-9_\\.]+)?")
	sub := reg.FindSubmatch([]byte(i.UA))
	if len(sub) > 1 {
		return strings.Replace(string(sub[1]), "_", ".", -1)
	}
	return ""
}

func (i ios) match() bool {
	reg := regexp.MustCompile("(?i)(iPhone|iPad|iPod)")
	return reg.Match([]byte(i.UA))
}
