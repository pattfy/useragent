package platform

// Platform is a public wrapper of platform.core
type Platform struct {
	base
}

// Parse will parse the details from user-agent given
func Parse(ua string) Platform {
	ms := matchers(ua)
	for _, p := range ms {
		if p.match() {
			return Platform{p}
		}
	}
	return Platform{
		&core{UA: ua},
	}
}

// IsWindows will return true if it's a windows-base platform.
func IsWindows(ua string) bool {
	c := &core{UA: ua}
	return windows{c}.match()
}

func matchers(ua string) []base {
	c := &core{UA: ua}
	return []base{
		ios{c},
		mac{c},
		windowsPhone{c},
		windows{c},
		android{c},
		linux{c},
	}
}
