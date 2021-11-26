package main

import "fmt"

func main() {
	// int型の変数を宣言
	var i int = 12345

	// int型からint64型への変換
	var i64 int64 = int64(i)

	fmt.Println(i64)
}