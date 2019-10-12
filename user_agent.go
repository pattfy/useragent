package useragent

import (
	"github.com/pattfy/useragent/bot"
	"github.com/pattfy/useragent/browser"
	"github.com/pattfy/useragent/device"
	"github.com/pattfy/useragent/platform"
)

// UserAgent contains the raw User-Agent, and the `bot`, `device`, `platform` and `browser` details parsed.
type UserAgent struct {
	UA       string
	Bot      bot.Bot
	Device   device.Device
	Platform platform.Platform
	Browser  browser.Browser
}

// New will create a new instance of UserAgent and parse the details.
func New(ua string) *UserAgent {
	d := device.Parse(ua)
	b := bot.Parse(ua)
	p := platform.Parse(ua)
	br := browser.Parse(ua)

	return &UserAgent{
		UA:       ua,
		Device:   d,
		Bot:      b,
		Platform: p,
		Browser:  br,
	}
}

// IsBot check whether the User-Agent is identified as a bot.
// alias of Bot.IsBot
func (ua UserAgent) IsBot() bool {
	return ua.Bot.IsBot()
}

// IsMobile check whether the User-Agent is identified as from smartphone device.
// alias of Device.IsMobile
func (ua UserAgent) IsMobile() bool {
	return ua.Device.IsMobile()
}

// IsTablet check whether the User-Agent is identified as from tablet device.
// alias of Device.IsTablet
func (ua UserAgent) IsTablet() bool {
	return ua.Device.IsTablet()
}

// FullVersion return the full version of the browser identified.
// alias of Browser.FullVersion
func (ua UserAgent) FullVersion() string {
	return ua.Browser.FullVersion()
}

// Name return the name of the browser identified.
// alias of Browser.Name
func (ua UserAgent) Name() string {
	return ua.Browser.Name()
}

// Version return the major version of the browser identified.
// alias of Browser.Version
func (ua UserAgent) Version() string {
	return ua.Browser.Version()
}
