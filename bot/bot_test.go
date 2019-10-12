package bot

import (
	"strings"
	"testing"
)

func TestIsBot(t *testing.T) {
	cases := []struct {
		desc           string
		ua             string
		isBot          bool
		isSearchEngine bool
	}{
		{
			"360spider",
			"Mozilla/5.0 (Windows; U; Windows NT 5.1; zh-CN; rv:1.8.0.11) Firefox/1.5.0.11; 360Spider",
			true,
			false,
		},
		{
			"googleBot",
			"Googlebot/2.1 (+http://www.google.com/bot.html)",
			true,
			true,
		},
		{
			"curl",
			"curl/7.37.0",
			true,
			false,
		},
		{
			"baidu",
			"Mozilla/5.0 (compatible; Baiduspider/2.0; +http://www.baidu.com/search/spider.html)",
			true,
			true,
		},
		// bing https://www.bing.com/webmaster/help/which-crawlers-does-bing-use-8c184ec0
		{
			"bingbot",
			"Mozilla/5.0 (iPhone; CPU iPhone OS 7_0 like Mac OS X) AppleWebKit/537.51.1 (KHTML, like Gecko) Version/7.0 Mobile/11A465 Safari/9537.53 (compatible; bingbot/2.0; +http://www.bing.com/bingbot.htm)",
			true,
			true,
		},
		{
			"AdIdxBot",
			"Mozilla/5.0 (compatible; adidxbot/2.0; +http://www.bing.com/bingbot.htm)",
			true,
			true,
		},
		{
			"bing preview",
			"Mozilla/5.0 (Windows NT 6.1; WOW64) AppleWebKit/534+ (KHTML, like Gecko) BingPreview/1.0b",
			true,
			false,
		},
		{
			"empty",
			"",
			true,
			false,
		},
		{
			"empty space",
			" ",
			true,
			false,
		},
	}

	for _, tc := range cases {
		b := Parse(tc.ua)
		if b.IsBot() != tc.isBot {
			t.Fatalf("[%s] IsBot: `%t` expected, `%t` got", tc.desc, tc.isBot, b.IsBot())
		}
		if b.IsSearchEngine() != tc.isSearchEngine {
			t.Fatalf("[%s] IsSearchEngine: `%t` expected, `%t` got", tc.desc, tc.isSearchEngine, b.IsSearchEngine())
		}
	}
}

func TestMatchSearchEngine(t *testing.T) {
	googleBot := "Mozilla/5.0 (compatible; Googlebot/2.1; +http://www.google.com/bot.html)"
	if m := matchSearchEngine(strings.ToLower(googleBot)); !m {
		t.Fatalf("matchSearchEngine: `true` expected, `%t` got", m)
	}
	chrome := "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_13_6) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/77.0.3865.90 Safari/537.36"
	if m := matchSearchEngine(strings.ToLower(chrome)); m {
		t.Fatalf("matchSearchEngine: `false` expected, `%t` got", m)
	}
}

func TestBotList(t *testing.T) {
	if len(botList.Bots) == 0 {
		t.Fatal("bot.botList.Bots is expected to be initialized by init()")
	}
	if len(botList.SearchEngines) == 0 {
		t.Fatal("bot.botList.SearchEngines is expected to be initialized by init()")
	}
}
