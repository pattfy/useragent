package browser

import (
	"regexp"
)

type firefox struct {
	*core
}

func (f firefox) ID() string {
	return "firefox"
}

func (f firefox) Name() string {
	return "Firefox"
}

func (f firefox) FullVersion() string {
	m := regexp.MustCompile("(Firefox|FxiOS)/([\\d.]+)").FindSubmatch([]byte(f.UA))
	if len(m) > 2 {
		return string(m[2])
	}
	return ""
}

func (f firefox) match() bool {
	return regexp.MustCompile("(?i)Firefox|FxiOS").Match([]byte(f.UA))
}
