package platform_test

import (
	"testing"

	"github.com/pattfy/useragent/platform"
)

func TestParse(t *testing.T) {
	cases := []struct {
		desc            string
		ua              string
		expectedID      string
		expectedName    string
		expectedVersion string
		isWindows       bool
	}{
		{
			"ipad",
			"Mozilla/5.0 (iPad; CPU OS 11_0 like Mac OS X) AppleWebKit/604.1.34 (KHTML, like Gecko) Version/11.0 Mobile/15A5341f Safari/604.1",
			"ios",
			"iOS",
			"11.0",
			false,
		},
		{
			"iphone",
			"Mozilla/5.0 (iPhone; CPU iPhone OS 11_0 like Mac OS X) AppleWebKit/604.1.38 (KHTML, like Gecko) Version/11.0 Mobile/15A372 Safari/604.1",
			"ios",
			"iOS",
			"11.0",
			false,
		},
		{
			"mac",
			"Mozilla/5.0 (Macintosh; Intel Mac OS X 10_13_6) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/76.0.3809.132 Safari/537.36",
			"mac",
			"Mac",
			"10.13.6",
			false,
		},
		{
			"mac with invalid os version",
			"Mozilla/5.0 (Macintosh; Intel Mac OS X abc) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/76.0.3809.132 Safari/537.36",
			"mac",
			"Mac",
			"",
			false,
		},
		{
			"windows",
			"Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/69.0.3497.100",
			"windows",
			"Windows",
			"10.0",
			true,
		},
		{
			"blackberry",
			"BlackBerry8520/5.0.0.681 Profile/MIDP-2.1 Configuration/CLDC-1.1 VendorID/600",
			"",
			"",
			"",
			false,
		},
	}
	for _, tc := range cases {
		p := platform.Parse(tc.ua)
		if p.ID() != tc.expectedID {
			t.Fatalf("[%s] id: `%s` expected, `%s` got", tc.desc, tc.expectedID, p.ID())
		}
		if p.Name() != tc.expectedName {
			t.Fatalf("[%s] name: `%s` expected, `%s` got", tc.desc, tc.expectedName, p.Name())
		}
		if p.Version() != tc.expectedVersion {
			t.Fatalf("[%s] version: `%s` expected, `%s` got", tc.desc, tc.expectedVersion, p.Version())
		}
		if platform.IsWindows(tc.ua) != tc.isWindows {
			t.Fatalf("[%s] isWindows: `%t` expected, `%t` got", tc.desc, tc.isWindows, platform.IsWindows(tc.ua))
		}
	}
}
