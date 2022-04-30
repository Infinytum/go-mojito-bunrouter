<p align="center">
    <img src="/.github/assets/gopher.png"
        height="300">
</p>

<p align="center">
    <a href="https://goreportcard.com/report/github.com/infinytum/go-mojito-bunrouter" alt="Go Report Card">
        <img src="https://goreportcard.com/badge/github.com/infinytum/go-mojito-bunrouter" /></a>
	<a href="https://github.com/infinytum/go-mojito-bunrouter" alt="Go Version">
        <img src="https://img.shields.io/github/go-mod/go-version/infinytum/go-mojito-bunrouter.svg" /></a>
	<a href="https://godoc.org/github.com/infinytum/go-mojito-bunrouter" alt="GoDoc reference">
        <img src="https://img.shields.io/badge/godoc-reference-blue.svg"/></a>
	<a href="https://github.com/Infinytum/go-mojito-bunrouter/blob/main/LICENSE" alt="Licence">
        <img src="https://img.shields.io/github/license/Ileriayo/markdown-badges?style=flat-square" /></a>
	<a href="https://makeapullrequest.com">
        <img src="https://img.shields.io/badge/PRs-welcome-brightgreen.svg?style=flat-square" alt="PRs Welcome"></a>
</p>

<p align="center">
    <a href="https://go.dev/" alt="Made with Go">
        <img src="https://ForTheBadge.com/images/badges/made-with-go.svg" /></a>
		
</p>

<p align="center">
Go-Mojito is a super-modular library to bootstrap your next Go web project. It can be used for strict API-only purposes as well as server-side rendering.
</p>

<p align="center"><sub>Icon made with <a href="https://gopherize.me">Gopherize</a> and <a href="https://www.flaticon.com/free-icon/mojito_920710">flaticon</a>.</sub></p>
<br>

## ⚡️ Quickstart
```go
package main

import (
	"github.com/infinytum/go-mojito"
    "github.com/infinytum/go-mojito-bunrouter"
)

func main() {
    // Register bunrouter as the default router.
    bunrouter.AsDefault()

    // Create your first route
    mojito.GET("/", func(req *mojito.Request, res *mojito.Response) error {
        res.String("Hello, World!")
        return nil
    })

    mojito.ListenAndServe(":8080")
}
```

## 🤖 Benchmarks

*TODO*

## ⚙️ Installation

- Make sure you have Go 1.18 or higher installed.
- Create a new folder and initialize your project with `go mod init github.com/your/project`.
- Install go-mojito with `go get -u github.com/infinytum/go-mojito`.
- Install go-mojito-bunrouter `go get -u github.com/infinytum/go-mojito-bunrouter`
