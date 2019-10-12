package device_test

import (
	"testing"

	"github.com/pattfy/useragent/device"
)

func TestParse(t *testing.T) {
	cases := []struct {
		desc     string
		ua       string
		id       string
		name     string
		isMobile bool
		isTablet bool
	}{
		// ipad
		{
			"ipad safari",
			"Mozilla/5.0 (iPad; CPU OS 12_0 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) Version/12.0 Mobile/15E148 Safari/604.1",
			"ipad",
			"iPad",
			false,
			true,
		},
		{
			"ipad chrome",
			"Mozilla/5.0 (iPad; CPU OS 12_0 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) CriOS/69.0.3497.91 Mobile/15E148 Safari/605.1",
			"ipad",
			"iPad",
			false,
			true,
		},
		{
			"ipad firefox",
			"Mozilla/5.0 (iPad; CPU OS 12_0 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) FxiOS/13.2b11866 Mobile/16A366 Safari/605.1.15",
			"ipad",
			"iPad",
			false,
			true,
		},
		{
			"ipad opera",
			"Mozilla/5.0 (iPad; CPU OS 12_0 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) Mobile/16A366",
			"ipad",
			"iPad",
			false,
			true,
		},
		{
			"ipad pro chrome",
			"Mozilla/5.0 (iPad; CPU OS 12_4_1 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) Version/12.1.2 Mobile/15E148 Safari/604.1",
			"ipad",
			"iPad",
			false,
			true,
		},

		// iphone
		{
			"iphone safari",
			"Mozilla/5.0 (iPhone; CPU iPhone OS 12_4_1 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) Version/12.1.2 Mobile/15E148 Safari/604.1",
			"iphone",
			"iPhone",
			true,
			false,
		},
		{
			"iphone chrome",
			"Mozilla/5.0 (iPhone; CPU iPhone OS 12_4 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) CriOS/76.0.3809.123 Mobile/15E148 Safari/605.1",
			"iphone",
			"iPhone",
			true,
			false,
		},

		// surface
		{
			"surface",
			"Mozilla/5.0 (compatible; MSIE 10.0; Windows NT 6.2; Trident/6.0; Touch)",
			"surface",
			"Microsoft Surface",
			false,
			true,
		},

		// kindle
		{
			"kindle",
			"Mozilla/5.0 (Linux; U; Android 2.3.4; en-us; Kindle Fire Build/GINGERBREAD) AppleWebKit/533.1 (KHTML, like Gecko) Version/4.0 Mobile Safari/533.1",
			"kindle",
			"Kindle",
			true,
			false,
		},
		// others
		{
			"huawei mediapad",
			"Mozilla/5.0 (Linux; U; Android 4.0.3; en-us; MediaPad 7 Lite Build/HuaweiMediaPad) AppleWebKit/534.30 (KHTML, like Gecko) Version/4.0 Safari/534.30",
			"",
			"",
			false,
			true,
		},
		// no match
		{
			"no match",
			"hogehoge",
			"",
			"",
			false,
			false,
		},
	}

	for _, tc := range cases {
		d := device.Parse(tc.ua)
		if d.ID() != tc.id {
			t.Fatalf("[%s] id: `%s` expected, `%s` got.", tc.desc, tc.id, d.ID())
		}
		if d.Name() != tc.name {
			t.Fatalf("[%s] name: `%s` expected, `%s` got.", tc.desc, tc.name, d.Name())
		}
		if d.IsMobile() != tc.isMobile {
			t.Fatalf("[%s] isMobile: `%t` expected, `%t` got.", tc.desc, tc.isMobile, d.IsMobile())
		}
		if d.IsTablet() != tc.isTablet {
			t.Fatalf("[%s] isTablet: `%t` expected, `%t` got.", tc.desc, tc.isTablet, d.IsTablet())
		}
	}
}
