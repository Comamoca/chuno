package chuno

import (
	"bytes"
	"embed"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strings"
	"text/template"

	"github.com/yuin/goldmark"
	"github.com/yuin/goldmark/extension"
	"github.com/yuin/goldmark/renderer/html"
)

var md = goldmark.New(
	goldmark.WithExtensions(extension.GFM),
	goldmark.WithRendererOptions(
		html.WithUnsafe(),
	),
)

//go:embed js/prism.js
//go:embed css/prism.css
var assets embed.FS

//go:embed html/index.html
var baseHTML []byte

var style string

func load(path string) ([]byte, error) {
	log.Println("load", path)
	f, err := os.Open(path)
	if err != nil {
		log.Fatal("File Open Error\n", err)
		return nil, err
	}

	defer f.Close()

	b, err := ioutil.ReadAll(f)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	return b, nil
}

func buildSytle(css string) string {
	return fmt.Sprintf("<style>%s</style>", css)
}

func render(mdtext []byte, isdark bool) ([]byte, error) {
	var buf bytes.Buffer
	css := buildSytle(cssLight)

	if isdark {
		css = buildSytle(cssDark)
		style = `background: #0d1117;`
	}

	err := md.Convert(mdtext, &buf)
	if err != nil {
		log.Fatal(err)

		return nil, err
	}
	tmpl, err := template.New("tmpl").Parse(string(baseHTML))
	if err != nil {
		log.Fatal(err)

		return nil, err
	}
	writer := new(strings.Builder)

	mp := map[string]interface{}{
		"Content":   buf.String(),
		"Css":       css,
		"DarkStyle": style,
	}

	tmpl.Execute(writer, mp)

	return []byte(writer.String()), nil
}
