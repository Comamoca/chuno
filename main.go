package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"net/http"
)

//shareing variable
var content []byte

var path string
var isDark = false
var port = 3535

func handler(w http.ResponseWriter, r *http.Request) {
	content, err := load(path)
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}

	rendred, err := render(content, isDark)
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}

	w.Write(rendred)
}

func LaunchPreviewServer(path string, port int) error {
	strPort := fmt.Sprintf(":%d", port)
	// Start setver
	go watch(path)
	http.HandleFunc("/", handler)
	fmt.Println("🌸 Server is starting at http://localhost" + strPort)
	err := http.ListenAndServe(strPort, nil)
	if err != nil {
		log.Fatal(err)
		return err
	}
	return nil
}

func strContains(arr []string, str string) bool{
	for _, v := range arr{
		if v == str{
			return true
		}
	}
	return false
}

func main() {
	flag.Parse()

	args := flag.Args()

	if len(args) == 0 {
		// red
		help()
		log.Fatal("No file specified to preview!!")
		os.Exit(1)
	}

	if strContains(args, "--dark") {
		isDark = true

	}

	path = args[0]

	err := LaunchPreviewServer(path,
	3535)
	if err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
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
