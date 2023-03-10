<div align="center">

![Last commit](https://img.shields.io/github/last-commit/Comamoca/chuno?style=flat-square)
![Repository Stars](https://img.shields.io/github/stars/Comamoca/chuno?style=flat-square)
![Issues](https://img.shields.io/github/issues/Comamoca/chuno?style=flat-square)
![Open Issues](https://img.shields.io/github/issues-raw/Comamoca/chuno?style=flat-square)
![Bug Issues](https://img.shields.io/github/issues/Comamoca/chuno/bug?style=flat-square)

</div>


<div align="center">

<img src="https://emoji2svg.deno.dev/api/ð¦" alt="eyecatch" height="100">

# Chuno
Instant preview server written in Goð¨

<br>
<br>


</div>

<table>
  <thead>
    <tr>
      <th style="text-align:center">ð¡æ¥æ¬èª</th>
      <th style="text-align:center"><a href="README.md">ðEnglish</a></th>
    </tr>
  </thead>
</table>

<div align="center">

</div>

## ð ä½¿ãæ¹

- CLIãã¼ã«ã¨ãã¦ä½¿ãå ´å  
`chuno README.md`

- ã©ã¤ãã©ãªã¨ãã¦ä½¿ãå ´å  
```go
err := chuno.LaunchPreviewServer(path,
		3535, isDark)
if err != nil {
	log.Fatal(err)
		os.Exit(1)
}
```

## â¬ï¸  Install

- CLIãã¼ã«ã¨ãã¦ä½¿ãå ´å  
`go install github.com/Comamoca/chuno/cmd/chuno@latest`

- ã©ã¤ãã©ãªã¨ãã¦ä½¿ãå ´å  
`go get github.com/Comamoca/chuno`

## âï¸   éçº

```sh
// debug run
go run ./cmd/chuno/main.go ./README.md
```

## ð Todo

- [ ] ã·ã³ã¿ãã¯ã¹ãã¤ã©ã¤ãã®è¿½å 

## ð ã©ã¤ã»ã³ã¹

MIT

### ð§© Modules

- [lrserver](https://github.com/jaschaephraim/lrserver)
- [goldmark](https://github.com/yuin/goldmark)
- [fsnotify](https://github.com/fsnotify/fsnotify)

## ð Affected projects

[å¥½ããªã¨ãã£ã¿ã§Markdownããã¬ãã¥ã¼ããããã®Live Reloadãµã¼ããç«ã¦ãã³ãã³ããGoã§ä½ã£ã](https://zenn.dev/fj68/articles/e00cc62c84225f)
