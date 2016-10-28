# Hello World

それでは定番の Hello World アプリケーションを作って見たいと思います。
あるリクエストパスにアクセスすると、Hello World を表示する簡単な Web アプリ
を作ってみましょう。

## プログラムを書く

実際のサンプルプログラムを書いてみます。
net/http パッケージの godoc にサンプルプログラムが書いてありますので、
今回はほぼそちらのコードを拝借して作ってみましょう。

* main.go

```go
package main

import (
	"io"
	"net/http"
	"log"
)

func HelloHandler(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "hello, world!\n")
}

func main() {
	http.HandleFunc("/hello", HelloHandler)
	log.Fatal(http.ListenAndServe(":8000", nil))
}

```

基本的にサンプルプログラムをそのままです。プログラムは Go のお約束である GOPATH
に配置します。GOPATH について詳しくは公式の資料に譲ります。
仮に GOPATH が ${HOME}/go だった場合は以下のような配置になります。

* /homge/yuichi/go/src/github.com/yuichi1004/src/go-webapp/chapter2/01_hello
** main.go

GOPATH が構成できて、コードが配置できたら、コードを実行してみましょう。

```
go run ./main.go
```

ブラウザーで `http://localhost:8000/hello` にアクセスしてみましょう。
Hello, wordl! のメッセージが表示されるはずです。

## Next

それでは次の章でリクエストのルーティングについて調べてみましょう。

次へ: [リクエストをルーティングする](./02_routing.md)
