package main

import (
	"log"
	"net/http"

	"./sub"
)

func main() {
	// struct の 実行結果
	sub.ClassPrint()
	http.HandleFunc("/", sub.SayhelloName)   //アクセスのルーティングを設定します。
	err := http.ListenAndServe(":9090", nil) //監視するポートを設定します。
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
