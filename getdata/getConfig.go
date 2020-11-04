package getdata

import (
	"io/ioutil"
	"strings"
)

// Config config struct
type Config struct {
	URL   string `json:"url"`
	Types string `json:"type"`
}

// GetConfig get config
func GetConfig(cors bool, config string) (c []*Config) {
	var list []string
	if config != "" {
		list = GetConfigFile(config)
	} else {
		list = []string{"https://t.me/ivmess", "https://t.me/s/V2List"}
	}
	var (
		types string = "sub"
	)
	for _, item := range list {
		if item != "" {
			if strings.HasPrefix(item, "http") {
				if strings.Contains(item, "https://t.me/s/") {
					types = "tg"
				}
				if cors {
					item = strings.Join([]string{"https://cors.izumana.ml", item}, "/?url=")
				}
				c = append(c, &Config{
					URL:   item,
					Types: types,
				})
			}
		}
	}
	return
}

// GetConfigFile get config file
func GetConfigFile(config string) (list []string) {
	byteValue, err := ioutil.ReadFile(config)
	if err != nil {
		return
	}
	return strings.Split(string(byteValue), "\n")
}
