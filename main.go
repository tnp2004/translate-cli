package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/tnp2004/translate-cli/config"
	"github.com/tnp2004/translate-cli/modules"
)

func envPath() string {
	return *flag.String("env", "./.env", "use for select specific env file")
}

func main() {
	config := config.LoadConfig(envPath())

	if len(os.Args) == 1 {
		fmt.Println("WARNING: Please input word for translate")
		return
	} else if len(os.Args) > 4 {
		fmt.Println("WARNING: input should be less than 4")
		return
	}

	word := os.Args[1]
	targetLang := "th"
	sourceLang := "en"

	if len(os.Args) >= 3 {
		sourceLang = os.Args[2]
	}

	if len(os.Args) >= 4 {
		targetLang = os.Args[3]
	}

	module := modules.InitModule(config)
	module.Translate(word, targetLang, sourceLang)
}
