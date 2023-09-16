package modules

import (
	"encoding/json"
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/tnp2004/translate-cli/config"
)

type IModule interface {
	Translate(word, target, source string)
}

type module struct {
	config config.IConfig
}

func InitModule(config config.IConfig) IModule {
	return &module{config}
}

type TranslateResponse struct {
	Data struct {
		Translations []Translate `json:"translations"`
	} `json:"data"`
}

type Translate struct {
	TranslatedText string `json:"translatedText"`
}

func (m *module) setHeader(agent *fiber.Agent) {
	agent.Set("content-type", "application/x-www-form-urlencoded")
	agent.Set("Accept-Encoding", "application/gzip")
	agent.Set("X-RapidAPI-Key", m.config.App().ApiKey())
	agent.Set("X-RapidAPI-Host", m.config.App().ApiHost())
}

func (m *module) Translate(word, target, source string) {
	payload := fmt.Sprintf("q=%v&target=%v&source=%v", word, target, source)
	agent := fiber.Post(m.config.App().Url()).Body([]byte(payload))
	m.setHeader(agent)
	_, body, errs := agent.Bytes()
	if errs != nil {
		fmt.Println("Error: ", errs)
	}

	res := new(TranslateResponse)
	if err := json.Unmarshal(body, res); err != nil {
		fmt.Printf("json unmarshal failed: %v", err)
	}

	translateRes := res.Data.Translations
	if len(translateRes) == 1 {
		fmt.Println(translateRes[0].TranslatedText)
		return
	}

	for i, t := range translateRes {
		if i == len(translateRes)-1 {
			fmt.Printf("%v\n", t.TranslatedText)
			return
		}

		fmt.Printf("%v,", t.TranslatedText)
	}
}
