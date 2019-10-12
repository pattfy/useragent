package browser

import (
	"regexp"
)

type opera struct {
	*core
}

func (o opera) ID() string {
	return "opera"
}

func (o opera) Name() string {
	return "Opera"
}

func (o opera) FullVersion() string {
	patterns := []string{
		"OPR/([\\d.]+)",
		"Version/([\\d.]+)",
	}
	for _, p := range patterns {
		m := regexp.MustCompile(p).FindSubmatch([]byte(o.UA))
		if len(m) > 1 {
			return string(m[1])
		}
	}
	return ""
}

func (o opera) match() bool {
	return regexp.MustCompile("(Opera|OPR\\/)").Match([]byte(o.UA))
}
