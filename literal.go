package main

import "fmt"

// リテラル：　値をプログラムコード内に記述するための表記法で、定数を記述するために使用する
func main() {
	
	// 論理値リテラル
	t := true
	f := false
	fmt.Println("論理値リテラル")
	fmt.Println("t: ", t)
	fmt.Println("f: ", f)

	// 整数リテラル
	a := 1234
	b := -1234
	c := 0123        // 0で始まる場合、以降は8進数とみなされる
	d := 0x1234ABCD  // 0xで始まる場合、以降は16進数とみなされる
	fmt.Println("整数リテラル")
	fmt.Println("a: ", a)
	fmt.Println("b: ", b)
	fmt.Println("c: ", c)
	fmt.Println("d: ", d)

	// 浮動小数点リテラル
	e := 0.0
	fmt.Println("浮動小数点リテラル")
	fmt.Println("e: ", e)

	// 以下を再宣言すると　no new variables on left side of :=　というコンパイルエラーとなるためコメントアウト
	// f := .12345e+5   // e以降は符号と指数
	// h := .12345e5    // 符号は省略可能
	
	// 文字列リテラル
	str1 := "ダブルクォートではエスケープ文字が解釈される"
	str2 := `バッククォートではエスケープ文字が解釈されないが、
	改行を含めることができる`
	fmt.Println("str1: ", str1)
	fmt.Println("str2: ", str2)
}