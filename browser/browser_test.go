package browser_test

import (
	"testing"

	"github.com/pattfy/useragent/browser"
)

func TestParse(t *testing.T) {
	cases := []struct {
		desc                string
		ua                  string
		expectedID          string
		expectedName        string
		expectedFullVersion string
		expectedVersion     string
	}{
		// firefox
		{
			"firefox",
			"Mozilla/5.0 (Windows NT x.y; Win64; x64; rv:10.0) Gecko/20100101 Firefox/10.0",
			"firefox",
			"Firefox",
			"10.0",
			"10",
		},
		// chrome
		{
			"chrome",
			"Mozilla/5.0 (Macintosh; Intel Mac OS X 10_13_6) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/76.0.3809.132 Safari/537.36",
			"chrome",
			"Chrome",
			"76.0.3809.132",
			"76",
		},
		// edge
		{
			"edge",
			"Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/64.0.3282.140 Safari/537.36 Edge/17.17134",
			"edge",
			"Edge",
			"17.17134",
			"17",
		},
		// opera
		{
			"opera",
			"Opera/9.80 (Android 2.3.3; Linux; Opera Mobi/ADR-1111101157; U; ja) Presto/2.9.201 Version/11.50",
			"opera",
			"Opera",
			"11.50",
			"11",
		},
		// IE
		{
			"ie 11.0",
			"Mozilla/5.0 (compatible; MSIE 10.0; Windows NT 6.3; WOW64; Trident/7.0; Touch; .NET4.0E; .NET4.0C; .NET CLR 3.5.30729; .NET CLR 2.0.50727; .NET CLR 3.0.30729; Tablet PC 2.0)",
			"ie",
			"Internet Explorer",
			"11.0",
			"11",
		},
		{
			"ie 10.0",
			"Mozilla/4.0 (compatible; MSIE 7.0; Windows NT 6.1; SLCC2; .NET CLR 2.0.50727; .NET CLR 3.5.30729; .NET CLR 3.0.30729; Media Center PC 6.0; InfoPath.2; .NET4.0C; .NET4.0E; Tablet PC 2.0)",
			"ie",
			"Internet Explorer",
			"7.0",
			"7",
		},
		{
			"i1 8.0",
			"Mozilla/4.0 (compatible; MSIE 8.0; Windows NT 6.1; Trident/4.0; SLCC2; .NET CLR 2.0.50727; .NET CLR 3.5.30729; .NET CLR 3.0.30729; Media Center PC 6.0; .NET4.0C; .NET4.0E; Tablet PC 2.0)",
			"ie",
			"Internet Explorer",
			"8.0",
			"8",
		},

		// UC browser
		// see: https://developers.whatismybrowser.com/useragents/explore/software_name/uc-browser/
		{
			"UC",
			"Mozilla/5.0 (Linux; U; Android 6.0.1; zh-CN; F5121 Build/34.0.A.1.247) AppleWebKit/537.36 (KHTML, like Gecko) Version/4.0 Chrome/40.0.2214.89 UCBrowser/11.5.1.944 Mobile Safari/537.36",
			"uc_browser",
			"UC Browser",
			"11.5.1.944",
			"11",
		},
		{
			"UC",
			"Opera/9.80 (Android; Opera Mini/7.28879/27.1662; U; en-us) Presto/2.8.119 Version/11.10 UCBrowser/8.6.1.262/145/355/",
			"uc_browser",
			"UC Browser",
			"8.6.1.262",
			"8",
		},
		{
			"UC empty version",
			"Opera/9.80 (Android; Opera Mini/7.28879/27.1662; U; en-us) Presto/2.8.119 Version/11.10 UCBrowser/xx/145/355/",
			"uc_browser",
			"UC Browser",
			"",
			"",
		},
	}

	for _, tc := range cases {
		b := browser.Parse(tc.ua)
		if b.ID() != tc.expectedID {
			t.Fatalf("[%s] id: `%s` expected, `%s` got.", tc.desc, tc.expectedID, b.ID())
		}
		if b.Name() != tc.expectedName {
			t.Fatalf("[%s] name: `%s` expected, `%s` got.", tc.desc, tc.expectedName, b.Name())
		}
		if b.FullVersion() != tc.expectedFullVersion {
			t.Fatalf("[%s] full_version: `%s` expected, `%s` got.", tc.desc, tc.expectedFullVersion, b.FullVersion())
		}
		if b.Version() != tc.expectedVersion {
			t.Fatalf("[%s] version: `%s` expected, `%s` got.", tc.desc, tc.expectedVersion, b.Version())
		}
	}
}
