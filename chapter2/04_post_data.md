# データを投稿する

さて本格的なステップに踏み込みましょう。ユーザーがデータを投稿し、そのデータを
表示することによって Web アプリケーションは成り立ちます。

ここでは簡単な Chat アプリを作ってみましょう。

# テンプレートを書く

まずは形をイメージしやすいようにテンプレートを作ってみましょう。今回は以下のような
テンプレートを使います。

* chat.templ

```
<html>
<head>
  <title>My First Chat</title>
</head>
<body>

<h1>My First Chat</h1>
<form action="/" method="POST">
  name:
  <input type="text" name="name"></input>
  message:
  <input type="text" name="message"></input>
  <input type="submit"></input>
</form>

<dl>
  {{ range $.Posts }}
    <dt>{{ .Name }}</dt>
    <dd>{{ .Message }}</dd>
  {{ end }}
</dl>
</body>
</html>
```

データの投稿にはおなじみの form が使われます。ここから送信されたデータを
サーバー側でハンドリングします。

テンプレートの後半には range という制御構文が使われています。これはテンプレートに
渡された Posts という配列を順番に実行せよというものです。
ループ内では `{{ .Name }}` のようにドットが使わています。このドットは go のテンプレート
でコンテキストと呼ばれるものです。range 構文の文脈でいうとこれは、現在ループで
処理する変数を表します。すなわち `{{ .Name }}` は現在ループで処理する要素の
Name というフィールドを表します。

この制御構造を使うことに寄って、フォームで受け取ったポストを複数順番にならべて
表示することができます。

# POST を受け取る

では実際に Form による POST データを受け取るプログラムを書いてみましょう。
プログラムは以下のようになります。

```go
package main

import (
	"html/template"
	"net/http"
	"log"
)

var (
	chatTemplate *template.Template
	posts []Post
)

type Post struct {
	Name string
	Message string
}

func init() {
	var err error
	chatTemplate, err = template.New("chat.tmpl").ParseFiles("./chat.tmpl")
	if err != nil {
		panic(err)
	}
}

func ChatHandler(w http.ResponseWriter, req *http.Request) {
	if req.Method == "POST" {
		newPost := Post {
			Name: req.PostFormValue("name"),
			Message: req.PostFormValue("message"),
		}
		posts = append(posts, newPost)
	}

	params := map[string]interface{} {
		"Posts": posts,
	}
	err := chatTemplate.Execute(w, params)
	if err != nil {
		panic(err)
	}
}

func main() {
	http.HandleFunc("/", ChatHandler)
	log.Fatal(http.ListenAndServe(":8000", nil))
}
```

POST を処理しているのは以下の部分です。

```go
func ChatHandler(w http.ResponseWriter, req *http.Request) {
	if req.Method == "POST" {
		newPost := Post {
			Name: req.PostFormValue("name"),
			Message: req.PostFormValue("message"),
		}
		posts = append(posts, newPost)
	}
    ...
}
```

http.Request は Method という変数を持っています。これが POST か否かをチェックします。
もし POST の場合は PostFormValue() を用いてフォームで投稿された値を取得しておきます。

最後に posts スライスに新しい Post 構造体の値を追加することに寄って、今までの
投稿に今回の投稿を加えることができます。

最後に、新しい投稿を加えた posts をテンプレートに渡してレンダリングします。

```go
	params := map[string]interface{} {
		"Posts": posts,
	}
	err := chatTemplate.Execute(w, params)
	if err != nil {
		panic(err)
	}
```

さてこれで http://localhost:8000 にアクセスしてみましょう。画面上に名前と
メッセージを入力してボタンを押すと次々にメッセージが投稿されるはずです。

# Next

さて、これで基本的な話はおしまいです！
次回以降、より実践的なアプリケーションを作ってみましょう。


