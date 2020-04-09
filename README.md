# 実行方法

## パターン 1

ビルド & 実行ができる。
ただ、少し時間がかかる

```sh
$ go run main.go
```

## パターン 2

ビルドしてから、作られたファイルを実行
一手間かかるけど、実行速度は早い

```sh
$ go build main.go
$ ./main
```

# Go 言語の簡単なチュートリアル

↓ ディレクトリ構造がしたの様だとする

```
./main.go
./sub/sub1.go
./sub/sub2.go
```

```main.go
// main.go
package main

import (
	"./sub" // ここではフォルダ(package)を指定
)

func main() { // mainファイルが実行される。mainは一個だけ。
    sub.ClassPrint()
}


```

```sub2.go
// sub2.go
// ファイル名はおそらく何でもいい。

// ここもフォルダ名を指定する。
package sub

import "fmt"

// JavaやPythonのClassの概念に近いね。
type mydata struct {
	num int
	str string
}

// 大文字じゃ無いとダメ！！！
func ClassPrint() {
	var x mydata
	x.num = 10
	x.str = "something"
	fmt.Printf("x.num=%d, x.str=%s\n", x.num, x.str)
}
```

# 簡単なサーバ

今回のプロジェクトでは、簡単なサーバが実装してある。
下記のコマンドで実行できる。

```sh
$ go build main.go
$ ./main
```

http://localhost:9090/?url_long=111&url_long=222

↑ これでサーバに接続

# 詳しく知りたいなら！

- [Go 言語の初心者が見ると幸せになれる場所](https://qiita.com/tenntenn/items/0e33a4959250d1a55045)
- [Go で簡単な Web サーバを立てる](https://astaxie.gitbooks.io/build-web-application-with-golang/content/ja/03.2.html)
- [Go の構造体を package に外出ししてハマッたこと](https://qiita.com/zurazurataicho/items/4a95e0daf0d960cfc2f7)
