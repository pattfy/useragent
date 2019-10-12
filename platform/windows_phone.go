package platform

import (
	"regexp"
)

type windowsPhone struct {
	*core
}

func (w windowsPhone) ID() string {
	return "windows_phone"
}

func (w windowsPhone) Name() string {
	return "Windows Phone"
}

func (w windowsPhone) Version() string {
	reg := regexp.MustCompile("(?i)Windows Phone ([\\d.]+)")
	m := reg.FindSubmatch([]byte(w.UA))
	if len(m) > 1 {
		return string(m[1])
	}
	return ""
}

func (w windowsPhone) match() bool {
	return regexp.MustCompile("Windows Phone").Match([]byte(w.UA))
}
