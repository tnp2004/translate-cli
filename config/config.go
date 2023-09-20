package config

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

func LoadConfig() IConfig {
	return &config{
		app: &app{
			url:     "https://microsoft-translator-text.p.rapidapi.com",
			apiKey:  "REPLACE YOUR X-RapidAPI-Key HERE",
			apiHost: "microsoft-translator-text.p.rapidapi.com",
		},
	}
}
