# テンプレートをレンダリングする

Web アプリケーションの実装にはテンプレートが欠かせません。
テンプレートとは以下のような HTML 文字列の一部に変数や制御構造などの
書き換え処理を埋め込んだものです。

* hello.tmpl

```
<html>
<head>
  <title>Hello!</title>
</head>
<body>
<h1>Hello {{ .Name }}</h1>
</body>
</html>
```

上の例は `{{ .Name }}` の部分を Name という変数で置き換えるものです。
今回はこのテンプレートを使って動的に Web ページを表示してみましょう。

## プログラムを書く

実際のサンプルプログラムを書いてみます。
これは最初の hello world プログラムに改良を加えたものです。

* main.go

```go
package main

import (
	"html/template"
	"net/http"
	"strings"
	"log"
)

var (
	helloTemplate *template.Template
)

func init() {
	var err error
	helloTemplate, err = template.New("hello.tmpl").ParseFiles("./hello.tmpl")
	if err != nil {
		panic(err)
	}
}

func HelloHandler(w http.ResponseWriter, req *http.Request) {
	pathParams := strings.Split(req.URL.Path, "/")
	params := struct {
		Name string
	}{
		pathParams[2],
	}
	err := helloTemplate.Execute(w, params)
	if err != nil {
		panic(err)
	}
}

func main() {
	http.HandleFunc("/hello/", HelloHandler)
	log.Fatal(http.ListenAndServe(":8000", nil))
}

```

まずプログラムの起動時にテンプレートを読み込みます。以下のコードが該当箇所です。
init() 関数を書いておくと、Go パッケージが読み込まれたときに初期化処理として
実行されます。

```go
var (
	helloTemplate *template.Template
)

func init() {
	var err error
	helloTemplate, err = template.New("hello.tmpl").ParseFiles("./hello.tmpl")
	if err != nil {
		panic(err)
	}
}
```

テンプレートの読み込みは基本的に起動時に行いましょう。リクエストを処理する
ごとに読み込むのは CPU/IO リソースの無駄だからです。

読み込んだテンプレートをリクエストで受け取るときに処理します。

```go
func HelloHandler(w http.ResponseWriter, req *http.Request) {
	pathParams := strings.Split(req.URL.Path, "/")
	params := struct {
		Name string
	}{
		pathParams[2],
	}
	err := helloTemplate.Execute(w, params)
	if err != nil {
		panic(err)
	}
}
```

ここではリクエストパスからパラメターを取得します。`/hello/freind` でアクセスした場合、
2 番目の要素である `friend` を取得します。これを Name という名前でパラメータとして
設定し、template.Execute() に渡します。

プログラムを実行して `http://localhost:8000/hello/friend` にアクセスしてみましょう。
h1 要素の中身が `Hello friend!` となって表示されているはずです。

## Next

それでは次にデータを投稿する処理を行ってみましょう。

次へ: [データを投稿する](./04_post_data.md)


