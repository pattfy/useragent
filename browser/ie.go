package browser

import (
	"regexp"
)

type ie struct {
	*core
}

func (b ie) ID() string {
	return "ie"
}

func (b ie) Name() string {
	return "Internet Explorer"
}

func (b ie) FullVersion() string {
	if tridentVer := tridentVersion(b.UA); tridentVer != "" {
		return tridentVer
	}
	return msIEFullVersion(b.UA)
}

func (b ie) match() bool {
	return isMsIE(b.UA) || isModernIE(b.UA)
}

func isMsIE(ua string) bool {
	return regexp.MustCompile("MSIE").Match([]byte(ua)) && !regexp.MustCompile("Opera").Match([]byte(ua))
}

func msIEFullVersion(ua string) string {
	reg := regexp.MustCompile("MSIE ([\\d.]+)|Trident/.*?; rv:([\\d.]+)")
	m := reg.FindSubmatch([]byte(ua))
	if len(m) <= 1 {
		return ""
	}
	if string(m[1]) != "" {
		return string(m[1])
	}
	return string(m[2])
}

var tridentVersionMap = map[string]string{
	"4.0": "8.0",
	"5.0": "9.0",
	"6.0": "10.0",
	"7.0": "11.0",
	"8.0": "12.0",
}

func tridentVersion(ua string) string {
	reg := regexp.MustCompile("Trident/([0-9.]+)")
	m := reg.FindSubmatch([]byte(ua))
	if len(m) > 1 {
		return tridentVersionMap[string(m[1])]
	}
	return ""
}

func isModernIE(ua string) bool {
	return regexp.MustCompile("Trident/.*?; rv:(.*?)").Match([]byte(ua))
}
