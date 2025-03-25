package browser

import (
	"regexp"
	"strings"
)

type epiphany struct {
	*core
}

func (e epiphany) ID() string {
	return "epiphany"
}

func (e epiphany) Name() string {
	return "Epiphany"
}

func (e epiphany) FullVersion() string {
	m := regexp.MustCompile(`Epiphany/([\d.]+)`).FindSubmatch([]byte(e.UA))
	if len(m) > 1 {
		return string(m[1])
	}
	return ""
}

func (e epiphany) match() bool {
	return strings.Contains(e.UA, "Epiphany")
}
