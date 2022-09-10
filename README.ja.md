<div align="center">
  <img alt="DajareGo Logo" src="logo/DajareGo.svg" width="300">
</div>

# DajareGo

DajareGo は、ダジャレを検出するGo ライブラリです。

[![MIT LICENSE](https://img.shields.io/github/license/jugesuke/dajareGo)](./LICENSE)
![GitHub go.mod Go version](https://img.shields.io/github/go-mod/go-version/jugesuke/dajareGo)
[![GoDev][godev-image]][godev-url]

[[English version]](./README.md)

[go-version]: https://img.shields.io/github/go-mod/go-version/jugesuke/dajareGo
[godev-image]: https://pkg.go.dev/badge/github.com/jugesuke/dajareGo
[godev-url]: https://pkg.go.dev/github.com/jugesuke/dajareGo

## Getting Started

### Usage

まず、あなたのプロジェクトにDajareGoをインポートしましょう。
```bash
go get github.com/jugesuke/dajareGo
```
```golang
import "github.com/jugesuke/dajareGo"
```

次に形態素解析器をロードします。

```golang
if err := dajareGo.Init(); err != nil {
  panic(err)
}
```

そして、`IsDajare`関数を使ってダジャレの判定処理を行います。
```golang
result := dajareGo.IsDajare("アルミ缶の上にあるミカン")
```

ダジャレ判定の結果は、 `result.IsDajare`で`bool`値として取得できます。

```golang
result := dajareGo.IsDajare("アルミ缶の上にあるミカン")
if result.IsDajare {
  fmt.Println("This is Dajare")
} else {
  fmt.Println("This is not Dajare")
}
```
より詳細な仕様は[Document](https://pkg.go.dev/github.com/jugesuke/dajareGo/) をお読みください。

## このパッケージでのダジャレの定義
このパッケージでは、ダジャレを、
> **読みが同じ**または**似ている**が**それぞれ意味の違う言葉**が含まれるフレーズ

と定義して、検出を行っています。

## ロゴについて
The Gopher character is based on the Go mascot designed by [Renée French](https://reneefrench.blogspot.com/).
