<div align="center">

![Last commit](https://img.shields.io/github/last-commit/Comamoca/chuno?style=flat-square)
![Repository Stars](https://img.shields.io/github/stars/Comamoca/chuno?style=flat-square)
![Issues](https://img.shields.io/github/issues/Comamoca/chuno?style=flat-square)
![Open Issues](https://img.shields.io/github/issues-raw/Comamoca/chuno?style=flat-square)
![Bug Issues](https://img.shields.io/github/issues/Comamoca/chuno/bug?style=flat-square)

</div>


<div align="center">

<img src="https://emoji2svg.deno.dev/api/🍣" alt="eyecatch" height="100">

# Chuno
Instant preview server written in Go💨

<br>
<br>

</div>

</div>

<table>
  <thead>
    <tr>
      <th style="text-align:center">🍔English</th>
      <th style="text-align:center"><a href="README.ja.md">🍡日本語</a></th>
    </tr>
  </thead>
</table>

<div align="center">

</div>

## 🚀 How to use

- to CLI Tool  
`chuno README.md`

- to Library
```go
err := LaunchPreviewServer("README.md", 3535)

if err != nil {
	log.Fatal(err)
		os.Exit(1)
}
```

## ⬇️  Install

- to CLI Tool  
`go install github.com/Comamoca/chuno@latest`


- to Library  
`go get github.com/Comamoca/chuno`


## ⛏️   Development

```sh
// debug run
go run . README.md
```
## 📝 Todo

- [ ] Syntax hightlight

## 📜 License

MIT

### 🧩 Modules

- [lrserver](https://github.com/jaschaephraim/lrserver)
- [goldmark](https://github.com/yuin/goldmark)
- [fsnotify](https://github.com/fsnotify/fsnotify)

## 👏 Affected projects

[好きなエディタでMarkdownをプレビューするためのLive Reloadサーバを立てるコマンドをGoで作った](https://zenn.dev/fj68/articles/e00cc62c84225f)
