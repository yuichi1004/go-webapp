# リクエストをルーティングする

Web アプリケーションを作る際には指定された URL に応じて表示されるページを
切り替える必要があります。このリクエストのルーティング処理を見てみましょう。

# http.HandlerFunc

実は先の Hello World プログラムでもリクエストのルーティングを行っています。
http.HandlerFunc です。http.HandlerFunc は第一引数に指定されたパスにリクエスト
が来ると、第二引数に指定したハンドラに処理を渡します。

```go
http.HandleFunc("/hello", HelloHandler)
```

つまりこのように指定してやると /hello にリクエストが来た場合には HelloHandler
で処理が行われるわけです。これを複数書くことにより、ページの出し分けができます。

さらに、HandleFunc は / をパスの区切りとして認識します。以下のようにコードを書い
てみます。

```go
package main

import (
	"io"
	"net/http"
	"log"
)

func NotFoundHandler(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "request path not found\n")
}

func GreetingHandler(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "hi, forks!\n")
}

func ByeHandler(w http.ResponseWriter, req *http.Request) {
	io.WriteString(w, "bye bye!\n")
}

func main() {
	http.HandleFunc("/", NotFoundHandler)
	http.HandleFunc("/greeting/", GreetingHandler)
	http.HandleFunc("/greeting/bye", ByeHandler)
	log.Fatal(http.ListenAndServe(":8000", nil))
}
```

このプログラムに対して、以下のようなパスにアクセスすると対応するレスポンスは以下の
ようになります。

* `/`               => "request path not found"
* `/greeting/`      => "hi, forks!"
* `/greeting/bye`   => "bye bye!"
* `/greeting/hello` => "hi, forks!"

このように、リクエストパスに合致するものがない場合、パスの途中の改装
まで合致するものがあればそちらで処理が行われます。これによって特定のパス
以下の処理をまとめてハンドリングするといったことが可能です。

## 参考: http.ServerMux (マルチプレクサ) について

ちなみにこの仕組は、標準ライブラリの net/http における http.ServerMux の仕組み
によって実現されています。net/http 自体が DefaultMux (デフォルトマルチプレクサ)
を持っていて、デフォルトではこちらが使われます。

その一方、サードパーティが独自に ServerMux を実装して net/http 上で使うことがで
きます。この場合、より柔軟なリクエストハンドリングが可能です。

最も有名な ServerMux に httprouter (https://github.com/julienschmidt/httprouter)
があります。httprouter では、リクエストパスをパラメターとして扱ったり、
自動的におかしなパスを修正したりするなどの、柔軟なルーティングが可能になっています。

httprouter を用いた実装についてはより実践的な後の章で取り扱いたいと思います。

# Next

さて次の節ではテンプレートのレンダリングについて学びましょう。





