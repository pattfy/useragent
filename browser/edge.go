package browser

import (
	"regexp"
)

type edge struct {
	*core
}

func (e edge) ID() string {
	return "edge"
}

func (e edge) Name() string {
	return "Edge"
}

func (e edge) FullVersion() string {
	m := regexp.MustCompile("Edge/([\\d.]+)").FindSubmatch([]byte(e.UA))
	if len(m) > 1 {
		return string(m[1])
	}
	return ""
}

func (e edge) match() bool {
	return regexp.MustCompile("(?i)Edge/[\\d.]+|Trident/8").Match([]byte(e.UA))
}
