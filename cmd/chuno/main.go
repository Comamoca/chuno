package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/Comamoca/chuno"
)

func main() {
	flag.Parse()

	args := flag.Args()

	log.Println(args)

	if len(args) == 0 {
		// red
		help()
		log.Fatal("No file specified to preview!!")
		os.Exit(1)
	}

	path := args[0]

	var isDark = false

	if contain(args, "--dark") {
		isDark = true
	}

	err := chuno.LaunchPreviewServer(path,
	3535, isDark)
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
}

func contain(arr []string, str string) bool {
	for _, v := range arr{
		if v == str{
			return true
		}
	}
	return false
}

func help() {
	const help = `Chuno -

Usage:
chuno [PATH] [Option]

PATH Preview file path

- Option
--dark Preview DarkMode` + "\n\n"

	fmt.Print(help)
}
