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
	}

	word := os.Args[1]
	targetLang := flag.String("t", "th", "use for target language")
	sourceLang := flag.String("s", "en", "use for source language")
	flag.Parse()
	module := modules.InitModule(config)
	module.Translate(word, *targetLang, *sourceLang)
}
