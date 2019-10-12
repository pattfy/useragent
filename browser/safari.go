package browser

import (
	"regexp"
)

type safari struct {
	*core
}

func (s safari) ID() string {
	return "safari"
}

func (s safari) Name() string {
	return "Safari"
}

func (s safari) FullVersion() string {
	patterns := []string{
		"Version/([\\d.]+)",
		"Safari/([\\d.]+)",
		"AppleWebKit/([\\d.]+)",
	}
	for _, p := range patterns {
		m := regexp.MustCompile(p).FindSubmatch([]byte(s.UA))
		if len(m) > 1 {
			return string(m[1])
		}
	}
	return ""
}

func (s safari) match() bool {
	return regexp.MustCompile("Safari").Match([]byte(s.UA)) &&
		!regexp.MustCompile("Chrome|CriOS|PhantomJS|FxiOS").Match([]byte(s.UA))
}
