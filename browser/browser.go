package browser

import (
	"strings"
)

// Browser is a public wrapper of browser.core
type Browser struct {
	base
}

// Version return the major version of the browser.
// parsed from FullVersion()
func (browser Browser) Version() string {
	fullVer := browser.FullVersion()
	if fullVer == "" {
		return ""
	}
	return strings.Split(fullVer, ".")[0]
}

// Parse will parse the details from user-agent given
func Parse(ua string) Browser {
	ms := matchers(ua)
	for _, b := range ms {
		if b.match() {
			return Browser{b}
		}
	}
	return Browser{
		&core{UA: ua},
	}
}

func matchers(ua string) []base {
	c := &core{UA: ua}
	return []base{
		epiphany{c},
		firefox{c},
		edge{c},
		uCBrowser{c},
		opera{c},
		ie{c},
		chrome{c},
		safari{c},
	}
}
