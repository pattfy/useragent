package bot

import (
	"fmt"
	"io/ioutil"
	"strings"

	"github.com/pattfy/useragent/assets"
	yaml "gopkg.in/yaml.v2"
)

// List contains all the Bots defined.
type List struct {
	Bots          map[string]string `yaml:"bots"`
	SearchEngines map[string]string `yaml:"search_engines"`
}

// Bot contains the bot info defined.
type Bot struct {
	Name           string
	isBot          bool
	isSearchEngine bool
}

// IsBot check whether the User-Agent is identified as a bot.
func (b Bot) IsBot() bool {
	return b.isBot
}

// IsSearchEngine check whether the User-Agent is identified as a search engine.
func (b Bot) IsSearchEngine() bool {
	return b.isSearchEngine
}

// Parse will parse the bot info from the given user-agent.
func Parse(ua string) Bot {
	if ua == "" || strings.TrimSpace(ua) == "" {
		return Bot{
			Name:           "",
			isBot:          true,
			isSearchEngine: false,
		}
	}
	result := Bot{}
	lcUA := strings.ToLower(ua)
	for keyword, name := range botList.Bots {
		if strings.Contains(lcUA, keyword) {
			return Bot{
				Name:           name,
				isBot:          true,
				isSearchEngine: matchSearchEngine(lcUA),
			}
		}
	}
	return result
}

func matchSearchEngine(lcUA string) bool {
	for keyword := range botList.SearchEngines {
		if strings.Contains(lcUA, keyword) {
			return true
		}
	}
	return false
}

var botList List

func init() {
	f, err := assets.Assets.Open("/bots.yml")
	if err != nil {
		panic(err)
	}
	defer f.Close()
	buf, err := ioutil.ReadAll(f)
	if err != nil {
		panic(err)
	}

	err = yaml.Unmarshal(buf, &botList)
	if err != nil {
		fmt.Printf("buf: %+v\n", string(buf))
		panic(err)
	}
}
