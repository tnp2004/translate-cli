package config

import (
	"log"

	"github.com/joho/godotenv"
)

type IConfig interface {
	App() IAppConfig
}

type IAppConfig interface {
	Url() string
	ApiKey() string
	ApiHost() string
}

type config struct {
	app *app
}

type app struct {
	url     string
	apiKey  string
	apiHost string
}

func (c *config) App() IAppConfig {
	return c.app
}

func (a *app) Url() string {
	return a.url
}

func (a *app) ApiKey() string {
	return a.apiKey
}

func (a *app) ApiHost() string {
	return a.apiHost
}

func LoadConfig(path string) IConfig {
	envMap, err := godotenv.Read(path)
	if err != nil {
		log.Fatalf("load dotenv error: %v", err)
	}

	return &config{
		app: &app{
			url:     envMap["API_URL"],
			apiKey:  envMap["RAPID_API_KEY"],
			apiHost: envMap["RAPID_API_HOST"],
		},
	}
}
