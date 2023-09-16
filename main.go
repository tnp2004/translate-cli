package main

import (
	"os"

	"github.com/tnp2004/translate-cli/config"
	"github.com/tnp2004/translate-cli/modules"
)

func envPath() string {
	if len(os.Args) == 1 {
		return ".env"
	}

	return os.Args[1]
}

func main() {
	config := config.LoadConfig(envPath())

	module := modules.InitModule(config)
	module.Translate("hello world", "th", "en")
}
