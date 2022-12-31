package chuno

import (
	"bytes"
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

const base = `<!doctype html>
<html>
<meta name="viewport" content="width=device-width, initial-scale=1">
<head>
<script src="http://localhost:35729/livereload.js"></script>
</head>
{{.Css}}
<style>
.markdown-body {
		box-sizing: border-box;
		min-width: 200px;
		max-width: 980px;
		margin: 0 auto;
		padding: 45px;
	}

	@media (max-width: 767px) {
		.markdown-body {
			padding: 15px;
		}
	}
</style>
<body style="{{ .DarkStyle }}">
<article class="markdown-body">
  {{ .Content }}
</article>
</body>
</html>
`

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
	tmpl, err := template.New("tmpl").Parse(base)
	if err != nil {
		log.Fatal(err)

		return nil, err
	}
	writer := new(strings.Builder)

	mp :=  map[string]interface{}{
		"Content": buf.String(),
		"Css": css,
		"DarkStyle": style,
    }

	tmpl.Execute(writer, mp)

	return []byte(writer.String()), nil
}
