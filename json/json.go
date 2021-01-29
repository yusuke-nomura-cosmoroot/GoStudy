package main

import (
	"fmt"
	"io/ioutil"
	"log"
)

func main() {
	// ファイルにメッセージを書き込む
	hello := []byte("hello world\n")
	err := ioutil.WriteFile("./file.txt", hello, 0666)
	if err != nil { // エラー処理
		log.Fatal(err)
	}

	// ファイルの中身を全て読み出す
	message, err := ioutil.ReadFile("./file.txt")
	if err != nil { // エラー処理
		log.Fatal(err)
	}
	fmt.Print(string(message)) // 文字列にして表示
}
