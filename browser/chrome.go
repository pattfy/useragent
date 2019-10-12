package browser

import (
	"regexp"
)

type chrome struct {
	*core
}

func (c chrome) ID() string {
	return "chrome"
}

func (c chrome) Name() string {
	return "Chrome"
}

func (c chrome) FullVersion() string {
	patterns := []string{
		"Chrome/([\\d.]+)",
		"CriOS/([\\d.]+)",
	}
	for _, p := range patterns {
		m := regexp.MustCompile(p).FindSubmatch([]byte(c.UA))
		if len(m) > 1 {
			return string(m[1])
		}
	}
	return ""
}

// see: https://developer.chrome.com/multidevice/user-agent
func (c chrome) match() bool {
	isOpera := opera{c.core}.match()
	isEdge := edge{c.core}.match()
	return regexp.MustCompile("Chrome|CriOS").Match([]byte(c.UA)) &&
		!isOpera &&
		!isEdge
}
