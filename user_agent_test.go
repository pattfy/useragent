package useragent_test

import (
	"testing"

	"github.com/pattfy/useragent"
	"github.com/stretchr/testify/assert"
)

var cases = []struct {
	desc               string
	ua                 string
	isMobile           bool
	isTablet           bool
	isBot              bool
	isSearchEngine     bool
	deviceID           string
	deviceName         string
	platformID         string
	platformName       string
	platformVersion    string
	browserID          string
	browserName        string
	browserFullVersion string
	browserVersion     string
}{
	// Bots
	{
		desc:               "GoogleBot",
		ua:                 "Mozilla/5.0 (compatible; Googlebot/2.1; +http://www.google.com/bot.html)",
		isBot:              true,
		isSearchEngine:     true,
		deviceID:           "",
		deviceName:         "",
		platformID:         "",
		platformName:       "",
		platformVersion:    "",
		browserID:          "",
		browserName:        "",
		browserFullVersion: "",
		browserVersion:     "",
	},
	{
		desc:               "GoogleBotSmartphone (iPhone)",
		ua:                 "Mozilla/5.0 (iPhone; CPU iPhone OS 6_0 like Mac OS X) AppleWebKit/536.26 (KHTML, like Gecko) Version/6.0 Mobile/10A5376e Safari/8536.25 (compatible; Googlebot/2.1; +http://www.google.com/bot.html)",
		isMobile:           true,
		isTablet:           false,
		isBot:              true,
		isSearchEngine:     true,
		deviceID:           "iphone",
		deviceName:         "iPhone",
		platformID:         "ios",
		platformName:       "iOS",
		platformVersion:    "6.0",
		browserID:          "safari",
		browserName:        "Safari",
		browserFullVersion: "6.0",
		browserVersion:     "6",
	},
	{
		desc:               "GoogleBotSmartphone (Android)",
		ua:                 "Mozilla/5.0 (Linux; Android 6.0.1; Nexus 5X Build/MMB29P) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/41.0.2272.96 Mobile Safari/537.36 (compatible; Googlebot/2.1; +http://www.google.com/bot.html)",
		isMobile:           true,
		isTablet:           false,
		isBot:              true,
		isSearchEngine:     true,
		deviceID:           "",
		deviceName:         "",
		platformID:         "android",
		platformName:       "Android",
		platformVersion:    "6.0.1",
		browserID:          "chrome",
		browserName:        "Chrome",
		browserFullVersion: "41.0.2272.96",
		browserVersion:     "41",
	},
	{
		desc:               "GoogleBotEmulateMozilla",
		ua:                 "Mozilla/5.0 AppleWebKit/537.36 (KHTML, like Gecko; compatible; Googlebot/2.1; +http://www.google.com/bot.html) Safari/537.36",
		isMobile:           false,
		isTablet:           false,
		isBot:              true,
		isSearchEngine:     true,
		deviceID:           "",
		deviceName:         "",
		platformID:         "",
		platformName:       "",
		platformVersion:    "",
		browserID:          "safari",
		browserName:        "Safari",
		browserFullVersion: "537.36",
		browserVersion:     "537",
	},
	{
		desc:               "BingBot",
		ua:                 "Mozilla/5.0 (compatible; bingbot/2.0; +http://www.bing.com/bingbot.htm)",
		isMobile:           false,
		isTablet:           false,
		isBot:              true,
		isSearchEngine:     true,
		deviceID:           "",
		deviceName:         "",
		platformID:         "",
		platformName:       "",
		platformVersion:    "",
		browserID:          "",
		browserName:        "",
		browserFullVersion: "",
		browserVersion:     "",
	},
	{
		desc:           "BaiduBot",
		ua:             "Mozilla/5.0 (compatible; Baiduspider/2.0; +http://www.baidu.com/search/spider.html)",
		isMobile:       false,
		isTablet:       false,
		isBot:          true,
		isSearchEngine: true,
	},
	{
		desc:           "Twitterbot",
		ua:             "Twitterbot",
		isMobile:       false,
		isTablet:       false,
		isBot:          true,
		isSearchEngine: false,
	},
	{
		desc:           "YahooBot",
		ua:             "Mozilla/5.0 (compatible; Yahoo! Slurp; http://help.yahoo.com/help/us/ysearch/slurp)",
		isMobile:       false,
		isTablet:       false,
		isBot:          true,
		isSearchEngine: true,
	},
	{
		desc:           "FacebookPlatform",
		ua:             "facebookplatform/1.0 (+http://developers.facebook.com)",
		isMobile:       false,
		isTablet:       false,
		isBot:          true,
		isSearchEngine: false,
	},
	{
		desc:           "FacebookExternalHit",
		ua:             "facebookexternalhit/1.1 (+http://www.facebook.com/externalhit_uatext.php)",
		isMobile:       false,
		isTablet:       false,
		isBot:          true,
		isSearchEngine: false,
	},
	{
		desc:           "FaceBot",
		ua:             "FaceBot",
		isMobile:       false,
		isTablet:       false,
		isBot:          true,
		isSearchEngine: false,
	},
	{
		desc:           "NutchCVS",
		ua:             "NutchCVS/0.8-dev (Nutch; http://lucene.apache.org/nutch/bot.html; nutch-agent@lucene.apache.org)",
		isMobile:       false,
		isTablet:       false,
		isBot:          true,
		isSearchEngine: false,
	},
	{
		desc:           "MJ12bot",
		ua:             "Mozilla/5.0 (compatible; MJ12bot/v1.2.4; http://www.majestic12.co.uk/bot.php?+)",
		isMobile:       false,
		isTablet:       false,
		isBot:          true,
		isSearchEngine: false,
	},
	{
		desc:           "MJ12bot",
		ua:             "MJ12bot/v1.0.8 (http://majestic12.co.uk/bot.php?+)",
		isMobile:       false,
		isTablet:       false,
		isBot:          true,
		isSearchEngine: false,
	},
	{
		desc:           "AhrefsBot",
		ua:             "Mozilla/5.0 (compatible; AhrefsBot/4.0; +http://ahrefs.com/robot/)",
		isMobile:       false,
		isTablet:       false,
		isBot:          true,
		isSearchEngine: false,
	},
	{
		desc:           "AdsBotGoogle",
		ua:             "AdsBot-Google (+http://www.google.com/adsbot.html)",
		isMobile:       false,
		isTablet:       false,
		isBot:          true,
		isSearchEngine: false,
	},
	{
		desc:               "AdsBotGoogleMobile",
		ua:                 "Mozilla/5.0 (iPhone; CPU iPhone OS 9_1 like Mac OS X) AppleWebKit/601.1.46 (KHTML, like Gecko) Version/9.0 Mobile/13B143 Safari/601.1 (compatible; AdsBot-Google-Mobile; +http://www.google.com/mobile/adsbot.html)",
		isMobile:           true,
		isTablet:           false,
		isBot:              true,
		isSearchEngine:     false,
		deviceID:           "iphone",
		deviceName:         "iPhone",
		platformID:         "ios",
		platformName:       "iOS",
		platformVersion:    "9.1",
		browserID:          "safari",
		browserName:        "Safari",
		browserFullVersion: "9.0",
		browserVersion:     "9",
	},
	{
		desc:  "Slackbot-LinkExpanding",
		ua:    "Slackbot-LinkExpanding 1.0 (+https://api.slack.com/robots)",
		isBot: true,
	},
	{
		desc:  "Slack-ImgProxy",
		ua:    "Slack-ImgProxy (+https://api.slack.com/robots)",
		isBot: true,
	},
	{
		desc:  "Pingdom",
		ua:    "Pingdom.com_bot_version_1.4_(http://www.pingdom.com)/",
		isBot: true,
	},
	// IE
	{
		desc:               "IE11Win7",
		ua:                 "Mozilla/5.0 (Windows NT 6.1; Trident/7.0; rv:11.0) like Gecko",
		isBot:              false,
		isSearchEngine:     false,
		platformID:         "windows",
		platformName:       "Windows",
		platformVersion:    "6.1",
		browserID:          "ie",
		browserName:        "Internet Explorer",
		browserFullVersion: "11.0",
		browserVersion:     "11",
	},
	{
		desc:               "IE10",
		ua:                 "Mozilla/5.0 (compatible; MSIE 10.0; Windows NT 6.2; Trident/6.0)",
		isBot:              false,
		isSearchEngine:     false,
		platformID:         "windows",
		platformName:       "Windows",
		platformVersion:    "6.2",
		browserID:          "ie",
		browserName:        "Internet Explorer",
		browserFullVersion: "10.0",
		browserVersion:     "10",
	},
	{
		desc:               "Tablet",
		ua:                 "Mozilla/4.0 (compatible; MSIE 7.0; Windows NT 6.2; ARM; Trident/6.0; Touch; .NET4.0E; .NET4.0C; Tablet PC 2.0)",
		isTablet:           true,
		isBot:              false,
		isSearchEngine:     false,
		deviceID:           "surface",
		deviceName:         "Microsoft Surface",
		platformID:         "windows",
		platformName:       "Windows",
		platformVersion:    "6.2",
		browserID:          "ie",
		browserName:        "Internet Explorer",
		browserFullVersion: "10.0",
		browserVersion:     "10",
	},
	{
		desc:               "Touch",
		ua:                 "Mozilla/5.0 (compatible; MSIE 10.0; Windows NT 6.2; ARM; Trident/6.0; Touch)",
		isTablet:           true,
		isBot:              false,
		isSearchEngine:     false,
		deviceID:           "surface",
		deviceName:         "Microsoft Surface",
		platformID:         "windows",
		platformName:       "Windows",
		platformVersion:    "6.2",
		browserID:          "ie",
		browserName:        "Internet Explorer",
		browserFullVersion: "10.0",
		browserVersion:     "10",
	},
	{
		desc:               "Phone",
		ua:                 "Mozilla/4.0 (compatible; MSIE 7.0; Windows Phone OS 7.0; Trident/3.1; IEMobile/7.0; SAMSUNG; SGH-i917)",
		isMobile:           true,
		isTablet:           false,
		isBot:              false,
		isSearchEngine:     false,
		deviceID:           "",
		deviceName:         "",
		platformID:         "windows_phone",
		platformName:       "Windows Phone",
		platformVersion:    "",
		browserID:          "ie",
		browserName:        "Internet Explorer",
		browserFullVersion: "7.0",
		browserVersion:     "7",
	},
	{
		desc:               "IE6",
		ua:                 "Mozilla/4.0 (compatible; MSIE 6.0; Windows NT 5.0; .NET CLR 1.1.4322)",
		isMobile:           false,
		isTablet:           false,
		isBot:              false,
		isSearchEngine:     false,
		deviceID:           "",
		deviceName:         "",
		platformID:         "windows",
		platformName:       "Windows",
		platformVersion:    "5.0",
		browserID:          "ie",
		browserName:        "Internet Explorer",
		browserFullVersion: "6.0",
		browserVersion:     "6",
	},
	{
		desc:               "IE8Compatibility",
		ua:                 "Mozilla/4.0 (compatible; MSIE 7.0; Windows NT 6.1; WOW64; Trident/4.0; SLCC2; .NET CLR 2.0.50727; .NET CLR 3.5.30729; .NET CLR 3.0.30729; Media Center PC 6.0; .NET4.0C; .NET4.0E; InfoPath.3; MS-RTC LM 8)",
		isMobile:           false,
		isTablet:           false,
		isBot:              false,
		isSearchEngine:     false,
		deviceID:           "",
		deviceName:         "",
		platformID:         "windows",
		platformName:       "Windows",
		platformVersion:    "6.1",
		browserID:          "ie",
		browserName:        "Internet Explorer",
		browserFullVersion: "8.0",
		browserVersion:     "8",
	},
	{
		desc:               "IE10Compatibility",
		ua:                 "Mozilla/4.0 (compatible; MSIE 7.0; Windows NT 6.1; WOW64; Trident/6.0; SLCC2; .NET CLR 2.0.50727; .NET CLR 3.5.30729; .NET CLR 3.0.30729; Media Center PC 6.0; .NET4.0C; .NET4.0E; InfoPath.3; MS-RTC LM 8)",
		isMobile:           false,
		isTablet:           false,
		isBot:              false,
		isSearchEngine:     false,
		deviceID:           "",
		deviceName:         "",
		platformID:         "windows",
		platformName:       "Windows",
		platformVersion:    "6.1",
		browserID:          "ie",
		browserName:        "Internet Explorer",
		browserFullVersion: "10.0",
		browserVersion:     "10",
	},
	{
		desc:               "IE11Win81",
		ua:                 "Mozilla/5.0 (Windows NT 6.3; Trident/7.0; rv:11.0) like Gecko",
		isMobile:           false,
		isTablet:           false,
		isBot:              false,
		isSearchEngine:     false,
		deviceID:           "",
		deviceName:         "",
		platformID:         "windows",
		platformName:       "Windows",
		platformVersion:    "6.3",
		browserID:          "ie",
		browserName:        "Internet Explorer",
		browserFullVersion: "11.0",
		browserVersion:     "11",
	},
	{
		desc:               "IE11Win7",
		ua:                 "Mozilla/5.0 (Windows NT 6.1; Trident/7.0; rv:11.0) like Gecko",
		isMobile:           false,
		isTablet:           false,
		isBot:              false,
		isSearchEngine:     false,
		deviceID:           "",
		deviceName:         "",
		platformID:         "windows",
		platformName:       "Windows",
		platformVersion:    "6.1",
		browserID:          "ie",
		browserName:        "Internet Explorer",
		browserFullVersion: "11.0",
		browserVersion:     "11",
	},
	{
		desc:               "IE11b32Win7b64",
		ua:                 "Mozilla/5.0 (Windows NT 6.1; WOW64; Trident/7.0; rv:11.0) like Gecko",
		isMobile:           false,
		isTablet:           false,
		isBot:              false,
		isSearchEngine:     false,
		deviceID:           "",
		deviceName:         "",
		platformID:         "windows",
		platformName:       "Windows",
		platformVersion:    "6.1",
		browserID:          "ie",
		browserName:        "Internet Explorer",
		browserFullVersion: "11.0",
		browserVersion:     "11",
	},
	{
		desc:               "IE11b32Win7b64MDDRJS",
		ua:                 "Mozilla/5.0 (Windows NT 6.1; WOW64; Trident/7.0; MDDRJS; rv:11.0) like Gecko",
		isMobile:           false,
		isTablet:           false,
		isBot:              false,
		isSearchEngine:     false,
		deviceID:           "",
		deviceName:         "",
		platformID:         "windows",
		platformName:       "Windows",
		platformVersion:    "6.1",
		browserID:          "ie",
		browserName:        "Internet Explorer",
		browserFullVersion: "11.0",
		browserVersion:     "11",
	},
	{
		desc:               "IE11Compatibility",
		ua:                 "Mozilla/4.0 (compatible; MSIE 7.0; Windows NT 6.3; Trident/7.0)",
		isMobile:           false,
		isTablet:           false,
		isBot:              false,
		isSearchEngine:     false,
		deviceID:           "",
		deviceName:         "",
		platformID:         "windows",
		platformName:       "Windows",
		platformVersion:    "6.3",
		browserID:          "ie",
		browserName:        "Internet Explorer",
		browserFullVersion: "11.0",
		browserVersion:     "11",
	},
	// Microsoft Edge
	{
		desc:               "EdgeDesktop",
		ua:                 "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/42.0.2311.135 Safari/537.36 Edge/12.10240",
		isMobile:           false,
		isTablet:           false,
		isBot:              false,
		isSearchEngine:     false,
		deviceID:           "",
		deviceName:         "",
		platformID:         "windows",
		platformName:       "Windows",
		platformVersion:    "10.0",
		browserID:          "edge",
		browserName:        "Edge",
		browserFullVersion: "12.10240",
		browserVersion:     "12",
	},
	{
		desc:               "EdgeMobile",
		ua:                 "Mozilla/5.0 (Windows Phone 10.0; Android 4.2.1; DEVICE INFO) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/42.0.2311.135 Mobile Safari/537.36 Edge/12.10240",
		isMobile:           true,
		isTablet:           false,
		isBot:              false,
		isSearchEngine:     false,
		deviceID:           "",
		deviceName:         "",
		platformID:         "windows_phone",
		platformName:       "Windows Phone",
		platformVersion:    "10.0",
		browserID:          "edge",
		browserName:        "Edge",
		browserFullVersion: "12.10240",
		browserVersion:     "12",
	},
	// Opera
	{
		desc:               "OperaMac",
		ua:                 "Opera/9.27 (Macintosh; Intel Mac OS X; U; en)",
		isMobile:           false,
		isTablet:           false,
		isBot:              false,
		isSearchEngine:     false,
		deviceID:           "",
		deviceName:         "",
		platformID:         "mac",
		platformName:       "Mac",
		platformVersion:    "",
		browserID:          "opera",
		browserName:        "Opera",
		browserFullVersion: "",
		browserVersion:     "",
	},
	{
		desc:               "OperaWin",
		ua:                 "Opera/9.27 (Windows NT 5.1; U; en)",
		isMobile:           false,
		isTablet:           false,
		isBot:              false,
		isSearchEngine:     false,
		deviceID:           "",
		deviceName:         "",
		platformID:         "windows",
		platformName:       "Windows",
		platformVersion:    "5.1",
		browserID:          "opera",
		browserName:        "Opera",
		browserFullVersion: "",
		browserVersion:     "",
	},
	{
		desc:               "OperaWinNoLocale",
		ua:                 "Opera/9.80 (Windows NT 5.1) Presto/2.12.388 Version/12.10",
		isMobile:           false,
		isTablet:           false,
		isBot:              false,
		isSearchEngine:     false,
		deviceID:           "",
		deviceName:         "",
		platformID:         "windows",
		platformName:       "Windows",
		platformVersion:    "5.1",
		browserID:          "opera",
		browserName:        "Opera",
		browserFullVersion: "12.10",
		browserVersion:     "12",
	},
	{
		desc:               "OperaWin2Comment",
		ua:                 "Opera/9.80 (Windows NT 6.0; WOW64) Presto/2.12.388 Version/12.15",
		isMobile:           false,
		isTablet:           false,
		isBot:              false,
		isSearchEngine:     false,
		deviceID:           "",
		deviceName:         "",
		platformID:         "windows",
		platformName:       "Windows",
		platformVersion:    "6.0",
		browserID:          "opera",
		browserName:        "Opera",
		browserFullVersion: "12.15",
		browserVersion:     "12",
	},
	{
		desc:               "OperaMinimal",
		ua:                 "Opera/9.80",
		isMobile:           false,
		isTablet:           false,
		isBot:              false,
		isSearchEngine:     false,
		deviceID:           "",
		deviceName:         "",
		platformID:         "",
		platformName:       "",
		platformVersion:    "",
		browserID:          "opera",
		browserName:        "Opera",
		browserFullVersion: "",
	},
	{
		desc:               "OperaFull",
		ua:                 "Opera/9.80 (Windows NT 6.0; U; en) Presto/2.2.15 Version/10.10",
		isMobile:           false,
		isTablet:           false,
		isBot:              false,
		isSearchEngine:     false,
		deviceID:           "",
		deviceName:         "",
		platformID:         "windows",
		platformName:       "Windows",
		platformVersion:    "6.0",
		browserID:          "opera",
		browserName:        "Opera",
		browserFullVersion: "10.10",
		browserVersion:     "10",
	},
	{
		desc:               "OperaLinux",
		ua:                 "Opera/9.80 (X11; Linux x86_64) Presto/2.12.388 Version/12.10",
		isMobile:           false,
		isTablet:           false,
		isBot:              false,
		isSearchEngine:     false,
		deviceID:           "",
		deviceName:         "",
		platformID:         "linux",
		platformName:       "Generic Linux",
		platformVersion:    "",
		browserID:          "opera",
		browserName:        "Opera",
		browserFullVersion: "12.10",
		browserVersion:     "12",
	},
	{
		desc:               "OperaLinux - Ubuntu V41",
		ua:                 "Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/54.0.2840.99 Safari/537.36 OPR/41.0.2353.69",
		isMobile:           false,
		isTablet:           false,
		isBot:              false,
		isSearchEngine:     false,
		deviceID:           "",
		deviceName:         "",
		platformID:         "linux",
		platformName:       "Generic Linux",
		platformVersion:    "",
		browserID:          "opera",
		browserName:        "Opera",
		browserFullVersion: "41.0.2353.69",
		browserVersion:     "41",
	},
	{
		desc:               "OperaAndroid",
		ua:                 "Opera/9.80 (Android 4.2.1; Linux; Opera Mobi/ADR-1212030829) Presto/2.11.355 Version/12.10",
		isMobile:           true,
		isTablet:           false,
		isBot:              false,
		isSearchEngine:     false,
		deviceID:           "",
		deviceName:         "",
		platformID:         "android",
		platformName:       "Android",
		platformVersion:    "4.2.1",
		browserID:          "opera",
		browserName:        "Opera",
		browserFullVersion: "12.10",
		browserVersion:     "12",
	},
	{
		desc:               "OperaNested",
		ua:                 "Opera/9.80 (Windows NT 5.1; MRA 6.0 (build 5831)) Presto/2.12.388 Version/12.10",
		isMobile:           false,
		isTablet:           false,
		isBot:              false,
		isSearchEngine:     false,
		deviceID:           "",
		deviceName:         "",
		platformID:         "windows",
		platformName:       "Windows",
		platformVersion:    "5.1",
		browserID:          "opera",
		browserName:        "Opera",
		browserFullVersion: "12.10",
		browserVersion:     "12",
	},
	{
		desc:               "OperaMRA",
		ua:                 "Opera/9.80 (Windows NT 6.1; U; MRA 5.8 (build 4139); en) Presto/2.9.168 Version/11.50",
		isMobile:           false,
		isTablet:           false,
		isBot:              false,
		isSearchEngine:     false,
		deviceID:           "",
		deviceName:         "",
		platformID:         "windows",
		platformName:       "Windows",
		platformVersion:    "6.1",
		browserID:          "opera",
		browserName:        "Opera",
		browserFullVersion: "11.50",
		browserVersion:     "11",
	},
	// Firefox
	{
		desc:               "FirefoxMac",
		ua:                 "Mozilla/5.0 (Macintosh; Intel Mac OS X 10.6; rv:2.0b8) Gecko/20100101 Firefox/4.0b8",
		isBot:              false,
		isSearchEngine:     false,
		deviceID:           "",
		deviceName:         "",
		platformID:         "mac",
		platformName:       "Mac",
		platformVersion:    "10.6",
		browserID:          "firefox",
		browserName:        "Firefox",
		browserFullVersion: "4.0",
		browserVersion:     "4",
	},
	{
		desc:               "FirefoxLinux",
		ua:                 "Mozilla/5.0 (X11; Linux x86_64; rv:17.0) Gecko/20100101 Firefox/17.0",
		isBot:              false,
		isSearchEngine:     false,
		deviceID:           "",
		deviceName:         "",
		platformID:         "linux",
		platformName:       "Generic Linux",
		platformVersion:    "",
		browserID:          "firefox",
		browserName:        "Firefox",
		browserFullVersion: "17.0",
		browserVersion:     "17",
	},
	{
		desc:               "AndroidFirefox",
		ua:                 "Mozilla/5.0 (Android; Mobile; rv:17.0) Gecko/17.0 Firefox/17.0",
		isMobile:           true,
		isTablet:           false,
		isBot:              false,
		isSearchEngine:     false,
		deviceID:           "",
		deviceName:         "",
		platformID:         "android",
		platformName:       "Android",
		platformVersion:    "",
		browserID:          "firefox",
		browserName:        "Firefox",
		browserFullVersion: "17.0",
		browserVersion:     "17",
	},
	{
		desc:               "AndroidFirefoxTablet",
		ua:                 "Mozilla/5.0 (Android; Tablet; rv:26.0) Gecko/26.0 Firefox/26.0",
		isMobile:           false,
		isTablet:           true,
		isBot:              false,
		isSearchEngine:     false,
		deviceID:           "",
		deviceName:         "",
		platformID:         "android",
		platformName:       "Android",
		platformVersion:    "",
		browserID:          "firefox",
		browserName:        "Firefox",
		browserFullVersion: "26.0",
		browserVersion:     "26",
	},
	{
		desc:               "FirefoxOS",
		ua:                 "Mozilla/5.0 (Mobile; rv:26.0) Gecko/26.0 Firefox/26.0",
		isMobile:           true,
		isTablet:           false,
		isBot:              false,
		isSearchEngine:     false,
		deviceID:           "",
		deviceName:         "",
		platformID:         "",
		platformName:       "",
		platformVersion:    "",
		browserID:          "firefox",
		browserName:        "Firefox",
		browserFullVersion: "26.0",
		browserVersion:     "26",
	},
	// Chrome
	{
		desc:               "ChromeLinux",
		ua:                 "Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.11 (KHTML, like Gecko) Chrome/23.0.1271.97 Safari/537.11",
		isMobile:           false,
		isTablet:           false,
		isBot:              false,
		isSearchEngine:     false,
		deviceID:           "",
		deviceName:         "",
		platformID:         "linux",
		platformName:       "Generic Linux",
		platformVersion:    "",
		browserID:          "chrome",
		browserName:        "Chrome",
		browserFullVersion: "23.0.1271.97",
		browserVersion:     "23",
	},
	{
		desc:               "ChromeWin7",
		ua:                 "Mozilla/5.0 (Windows NT 6.1; WOW64) AppleWebKit/535.19 (KHTML, like Gecko) Chrome/18.0.1025.168 Safari/535.19",
		isMobile:           false,
		isTablet:           false,
		isBot:              false,
		isSearchEngine:     false,
		deviceID:           "",
		deviceName:         "",
		platformID:         "windows",
		platformName:       "Windows",
		platformVersion:    "6.1",
		browserID:          "chrome",
		browserName:        "Chrome",
		browserFullVersion: "18.0.1025.168",
		browserVersion:     "18",
	},
	{
		desc:               "ChromeMinimal",
		ua:                 "Mozilla/5.0 AppleWebKit/534.10 Chrome/8.0.552.215 Safari/534.10",
		isMobile:           false,
		isTablet:           false,
		isBot:              false,
		isSearchEngine:     false,
		deviceID:           "",
		deviceName:         "",
		platformID:         "",
		platformName:       "",
		platformVersion:    "",
		browserID:          "chrome",
		browserName:        "Chrome",
		browserFullVersion: "8.0.552.215",
		browserVersion:     "8",
	},
	{
		desc:               "ChromeMac",
		ua:                 "Mozilla/5.0 (Macintosh; U; Intel Mac OS X 10_6_5; en-US) AppleWebKit/534.10 (KHTML, like Gecko) Chrome/8.0.552.231 Safari/534.10",
		isMobile:           false,
		isTablet:           false,
		isBot:              false,
		isSearchEngine:     false,
		deviceID:           "",
		deviceName:         "",
		platformID:         "mac",
		platformName:       "Mac",
		platformVersion:    "10.6.5",
		browserID:          "chrome",
		browserName:        "Chrome",
		browserFullVersion: "8.0.552.231",
		browserVersion:     "8",
	},
	{
		desc:               "ChromeAndroid",
		ua:                 "Mozilla/5.0 (Linux; Android 4.2.1; Galaxy Nexus Build/JOP40D) AppleWebKit/535.19 (KHTML, like Gecko) Chrome/18.0.1025.166 Mobile Safari/535.19",
		isMobile:           true,
		isTablet:           false,
		isBot:              false,
		isSearchEngine:     false,
		deviceID:           "",
		deviceName:         "",
		platformID:         "android",
		platformName:       "Android",
		platformVersion:    "4.2.1",
		browserID:          "chrome",
		browserName:        "Chrome",
		browserFullVersion: "18.0.1025.166",
		browserVersion:     "18",
	},
	{
		desc:               "Chrome for iOS",
		ua:                 "Mozilla/5.0 (iPhone; CPU iPhone OS 11_3_1 like Mac OS X) AppleWebKit/604.1.34 (KHTML, like Gecko) CriOS/67.0.3396.87 Mobile/15E302 Safari/604.1",
		isMobile:           true,
		isTablet:           false,
		isBot:              false,
		isSearchEngine:     false,
		deviceID:           "iphone",
		deviceName:         "iPhone",
		platformID:         "ios",
		platformName:       "iOS",
		platformVersion:    "11.3.1",
		browserID:          "chrome",
		browserName:        "Chrome",
		browserFullVersion: "67.0.3396.87",
		browserVersion:     "67",
	},
	{
		desc:               "Epiphany",
		ua:                 "Mozilla/5.0 (X11; Ubuntu; Linux x86_64) AppleWebKit/605.1.15 (KHTML, like Gecko) Version/11.0 Safari/605.1.15 Ubuntu/16.04 (3.18.11-0ubuntu1) Epiphany/3.18.11",
		isMobile:           false,
		isTablet:           false,
		isBot:              false,
		isSearchEngine:     false,
		deviceID:           "",
		deviceName:         "",
		platformID:         "linux",
		platformName:       "Generic Linux",
		platformVersion:    "",
		browserID:          "epiphany",
		browserName:        "Epiphany",
		browserFullVersion: "3.18.11",
		browserVersion:     "3",
	},
	// others
	{
		desc:  "empty",
		ua:    "",
		isBot: true,
	},
	{
		desc: "Mozilla",
		ua:   "Mozilla/5.0",
	},
	{
		desc:  "Curl",
		ua:    "curl/7.28.1",
		isBot: true,
	},
	{
		desc:  "Python",
		ua:    "Python-urllib/2.7",
		isBot: true,
	},
}

func TestParse(t *testing.T) {
	for _, tc := range cases {
		t.Run(tc.desc, func(t *testing.T) {
			ua := useragent.New(tc.ua)
			assert.NotNil(t, ua)
			assert.Equal(t, tc.isMobile, ua.IsMobile())
			assert.Equal(t, tc.isTablet, ua.IsTablet())
			assert.Equal(t, tc.isBot, ua.Bot.IsBot())
			assert.Equal(t, tc.isSearchEngine, ua.Bot.IsSearchEngine())
			assert.Equal(t, tc.deviceID, ua.Device.ID())
			assert.Equal(t, tc.deviceName, ua.Device.Name())
			assert.Equal(t, tc.platformID, ua.Platform.ID())
			assert.Equal(t, tc.platformName, ua.Platform.Name())
			assert.Equal(t, tc.platformVersion, ua.Platform.Version())
			assert.Equal(t, tc.browserID, ua.Browser.ID())
			assert.Equal(t, tc.browserName, ua.Name())
			assert.Equal(t, tc.browserFullVersion, ua.FullVersion())
			assert.Equal(t, tc.browserVersion, ua.Version())
		})
	}
}

func BenchmarkUserAgent(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range cases {
			useragent.New(tc.ua)
		}
	}
}
