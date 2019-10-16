# UserAgent 
[![CircleCI](https://circleci.com/gh/pattfy/useragent.svg?style=svg)](https://circleci.com/gh/pattfy/useragent)
UserAgent is a Golang library that parses HTTP User-Agent.

## Installation

```
go get -u github.com/pattfy/useragent
```

## Struct and Interface

```golang
// UserAgent
type UserAgent struct {
    UA       string  // User-Agent string
    Bot      bot.Bot
    Device   device.Device
    Platform platform.Platform
    Browser  browser.Browser
}
```

### bot.Bot
- `Name` (field)
Name of the bot
- `IsBot()`
IsBot check whether the User-Agent is identified as a bot.
- `IsSearchEngine()`
IsSearchEngine check whether the User-Agent is identified as a search engine.

### browser.Browser
- `ID()`
return the ID of the browser(lower-case).
- `Name()`
return the name of the browser.
- `FullVersion()`
return the full version of the browser.
- `Version()`
return the major version of the browser.

### device.Device
Now such devices are availaable.
- iPad
- iPhone
- Kindle
- Surface

#### Interface
- `ID()`
return the ID of the device(lower-case).
- `Name()`
return the name of the device.
- `IsMobile()`
check whether the User-Agent is a mobile device.
- `IsTablet()`
check whether the User-Agent is a tablet device.

### platform.Platform
Now such platforms are available.
- Android
- iOS
- Linux
- Mac
- Windows
- Windows Phone

#### Interface
- `ID()`
return the ID of the platform(lower-case).
- `Name()`
return the name of the platform.
- `Version()`
return the version of the platform.

## Usage

```golang
package main

import (
	"fmt"

	"github.com/pattfy/useragent"
)

func main() {
	userAgent := useragent.New("Mozilla/5.0 (iPhone; CPU iPhone OS 11_0 like Mac OS X) AppleWebKit/604.1.38 (KHTML, like Gecko) Version/11.0 Mobile/15A372 Safari/604.1")

	fmt.Printf("IsBot: %t\n", userAgent.IsBot())       // alias of bot.IsBot()
	fmt.Printf("IsMobile: %t\n", userAgent.IsMobile()) // alias of device.IsMobile()
	fmt.Printf("IsTablet: %t\n", userAgent.IsTablet()) // alias of device.IsTablet()

	fmt.Printf("FullVersion(): %q\n", userAgent.FullVersion()) // alias of browser.FullVersion()
	fmt.Printf("Version(): %q\n", userAgent.Version())         // alias of browser.Version()
	fmt.Printf("Name(): %q\n", userAgent.Name())               // alias of browser.Name()

	// bot
	bot := userAgent.Bot
	fmt.Printf("bot.IsBot(): %t\n", bot.IsBot())
	fmt.Printf("bot.IsSearchEngine(): %t\n", bot.IsBot())

	// device
	device := userAgent.Device
	fmt.Printf("device.ID(): %q\n", device.ID())     // id of device(lowser-case)
	fmt.Printf("device.Name(): %q\n", device.Name()) // name of device

	// browser
	browser := userAgent.Browser
	fmt.Printf("browser.ID(): %q\n", browser.ID())                   // id of browser(lower-case)
	fmt.Printf("browser.Name(): %q\n", browser.Name())               // name of browser
	fmt.Printf("browser.FullVersion(): %q\n", browser.FullVersion()) // full version of browser
	fmt.Printf("browser.Version(): %q\n", browser.Version())         // major version of browser

	// platform
	platform := userAgent.Platform
	fmt.Printf("platform.ID(): %q\n", platform.ID())           // id of platform(lower-case)
	fmt.Printf("platform.Name(): %q\n", platform.Name())       // name of platform
	fmt.Printf("platform.Version(): %q\n", platform.Version()) // version of platform
}
```

## License
MIT License.
