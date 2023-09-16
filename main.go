package main

import (
	"fmt"
	"os"

	"github.com/tnp2004/translate-cli/config"
)

func envPath() string {
	if len(os.Args) == 1 {
		return ".env"
	}

	return os.Args[1]
}

func main() {
	config := config.LoadConfig(envPath())
	fmt.Println(config.App().ApiKey())
}
