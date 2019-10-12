package browser

import (
	"testing"
)

func TestMatch(t *testing.T) {

	cases := []struct {
		desc        string
		ua          string
		fullVersion string
		version     string
		matched     bool
	}{
		{
			"chrome for tablet",
			"Mozilla/5.0 (Linux; Android 4.0.4; Galaxy Nexus Build/IMM76B) AppleWebKit/535.19 (KHTML, like Gecko) Chrome/18.0.1025.133 Safari/535.19",
			"18.0.1025.133",
			"18",
			true,
		},
		{
			"chrome for mobile",
			"Mozilla/5.0 (Linux; Android 4.0.4; Galaxy Nexus Build/IMM76B) AppleWebKit/535.19 (KHTML, like Gecko) Chrome/18.0.1025.133 Mobile Safari/535.19",
			"18.0.1025.133",
			"18",
			true,
		},
		{
			"chrome for mac",
			"Mozilla/5.0 (Macintosh; Intel Mac OS X 10_13_6) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/69.0.3497.100 Safari/537.36",
			"69.0.3497.100",
			"69",
			true,
		},
		{
			"chrome for windows",
			"Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/69.0.3497.100",
			"69.0.3497.100",
			"69",
			true,
		},
		{
			"chrome for iOS",
			"Mozilla/5.0 (iPhone; CPU iPhone OS 10_3 like Mac OS X) AppleWebKit/602.1.50 (KHTML, like Gecko) CriOS/56.0.2924.75 Mobile/14E5239e Safari/602.1",
			"56.0.2924.75",
			"56",
			true,
		},
		{
			"safari",
			"Mozilla/5.0 (Macintosh; Intel Mac OS X 10_13_6) AppleWebKit/605.1.15 (KHTML, like Gecko) Version/13.0.1 Safari/605.1.15",
			"",
			"",
			false,
		},
		{
			"opera",
			"Mozilla/5.0 (Windows NT 6.1; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/28.0.1500.52 Safari/537.36 OPR/15.0.1147.130",
			"",
			"",
			false,
		},
	}

	for _, tc := range cases {
		c := chrome{&core{UA: tc.ua}}

		if c.match() != tc.matched {
			t.Fatalf("[%s] match: `%t` expected, `%t` got.", tc.desc, tc.matched, c.match())
		}
		if c.match() {
			if c.ID() != "chrome" {
				t.Fatalf("[%s] id: `chrome` expected, `%s` got.", tc.desc, c.ID())
			}
			if c.Name() != "Chrome" {
				t.Fatalf("[%s] name: `Chrome` expected, `%s` got.", tc.desc, c.Name())
			}
			if c.FullVersion() != tc.fullVersion {
				t.Fatalf("[%s] full_version: `%s` expected, `%s` got.", tc.desc, tc.fullVersion, c.FullVersion())
			}
		}
	}
}
