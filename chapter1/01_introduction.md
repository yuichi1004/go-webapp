# はじめに

ここでは、Go 言語を使った Web アプリケーションの開発について、基本から
実際のアプリケーションを開発を行うまで解説していきます。

## 想定している読者

主に以下の読者を想定しています。

* Go 言語に興味がある、触ったことがある
* Go 言語による Web アプリケーションの開発に興味がある

逆に以下の読者は対象としていません。他の資料と合わせて参考にしていただければと思います。

* Go 言語の基本的な仕様・開発方法について知らない
* すでに Go による Web アプリケーションの開発の経験があってより先週的なトピックを必要としている

## なぜ Go 言語なのか

まずそもそも Web アプリケーションを作成するにあたってなぜ Go 言語なのでしょうか。
Go 言語は以下の理由で Web アプリケーションの開発に優れています。

* 並行処理の記述を自然に行うことができる
* 処理効率が良いためスケールする
* Web アプリケーション開発にあたって標準ライブラリが充実している

Go 言語は Groutine と呼ばれる自然に並行処理を行うことができる仕組みの上に、
ランタイムや標準ライブラリが実装されています。そのため、Go の流儀に則って
プログラムを書いていれば、自然とリクエストが平行に処理されます。

また、Go 言語はコンパイル型の言語であるため、CPU ヘビーな処理が効率よく書けます。
たくさんのトラフィックをさばく Web アプリケーションであっても、少ない CPU リソース
で効率良く動かすことができます。

また、標準ライブラリが非常に充実しています。例えば net/http という標準ライブラリ
を使えば簡単に Web アプリを書くことができます。リクエストのルーティング、
テンプレートエンジンと言った Web アプリケーションの構成に必要なさまざまな機能も
標準で備えています。

## Next

それでは次章から簡単なサンプルアプリを作ってみたいと思います。

次へ: [Hello World](../chapter2/01_hello.md)

