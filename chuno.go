package chuno

import (
	"fmt"
	"log"
	"net/http"
	"os"
	// "path/filepath"
)

//shareing variable
var content []byte

var Path string
var isDark = false
var port = 3535

func handler(w http.ResponseWriter, r *http.Request) {
	content, err := load(Path)
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

func strContains(arr []string, str string) bool{
	for _, v := range arr{
		if v == str{
			return true
		}
	}
	return false
}

func LaunchPreviewServer(path string, port int, darkmode bool) error {

	Path = path
	isDark = darkmode

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
