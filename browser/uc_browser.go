package browser

import (
	"regexp"
)

type uCBrowser struct {
	*core
}

func (u uCBrowser) ID() string {
	return "uc_browser"
}

func (u uCBrowser) Name() string {
	return "UC Browser"
}

func (u uCBrowser) FullVersion() string {
	m := regexp.MustCompile(`UCBrowser/([\d.]+)`).FindSubmatch([]byte(u.UA))
	if len(m) > 1 {
		return string(m[1])
	}
	return ""
}

func (u uCBrowser) match() bool {
	return regexp.MustCompile("UCBrowser").Match([]byte(u.UA))
}
