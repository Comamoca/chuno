<div align="center">

![Last commit](https://img.shields.io/github/last-commit/Comamoca/chuno?style=flat-square)
![Repository Stars](https://img.shields.io/github/stars/Comamoca/chuno?style=flat-square)
![Issues](https://img.shields.io/github/issues/Comamoca/chuno?style=flat-square)
![Open Issues](https://img.shields.io/github/issues-raw/Comamoca/chuno?style=flat-square)
![Bug Issues](https://img.shields.io/github/issues/Comamoca/chuno/bug?style=flat-square)

</div>


<div align="center" style="margin-bottom: 4rem">

<div style="font-size: 5rem;">🦊</div>
<div style="font-size: 2rem;">Chuno</div>
<div style="font-size: 1.3rem;">Go製の簡易的なMarkDown プレビュー サーバー💨</div>
</div>


</div>

<table>
  <thead>
    <tr>
      <th style="text-align:center">🍡日本語</th>
      <th style="text-align:center"><a href="README.md">🍔English</a></th>
    </tr>
  </thead>
</table>

<div align="center">

</div>

## 🚀 使い方

- CLIツールとして使う場合  
`chuno README.md`

- ライブラリとして使う場合  
```go
err := LaunchPreviewServer("README.md", 3535)

if err != nil {
	log.Fatal(err)
		os.Exit(1)
}
```

## ⬇️  Install

- CLIツールとして使う場合  
`go install github.com/Comamoca/chuno@latest`


- ライブラリとして使う場合  
`go get github.com/Comamoca/chuno`

## ⛏️   開発

```sh
// debug run
go run . README.md
```

## 📝 Todo

- [ ] シンタックスハイライトの追加

## 📜 ライセンス

MIT

### 🧩 Modules

- [lrserver](https://github.com/jaschaephraim/lrserver)
- [goldmark](https://github.com/yuin/goldmark)
- [fsnotify](https://github.com/fsnotify/fsnotify)

## 👏 Affected projects

[好きなエディタでMarkdownをプレビューするためのLive Reloadサーバを立てるコマンドをGoで作った](https://zenn.dev/fj68/articles/e00cc62c84225f)
