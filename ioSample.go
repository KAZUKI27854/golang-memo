package main

import (
	"fmt"
	"os"
)

func main() {

	/* 入出力
	【fmtパッケージ】
	*/

	// Sprintf: 値を文字列にフォーマット（出力するわけではない）
	s := fmt.Sprintf("%sの降水確率は %d %%です", "東京", 30)

	fmt.Println("Sprintfのフォーマット結果： ", s)
	fmt.Println()

	/*
	 【出力】
	 Print: 標準出力に出力
	 Println: 標準出力に出力し、改行
	 Printf: 文字列にフォーマットし、標準出力に出力
	*/

	/*
	 【標準エラー出力（osパッケージと併用）】
	 FPrint: 標準エラー出力に出力
	 FPrintln: 標準エラー出力に出力し、改行
	 FPrintf: 文字列にフォーマットし、標準エラー出力に出力

	 => これらの関数の第1パラメータは出力先になるので、第1パラメータへ os.Stderr(標準エラー出力)を指定する。
	*/
	fmt.Fprint(os.Stderr, "標準エラー出力: ")
	fmt.Fprintf(os.Stderr, "%sの降水確率は %d %%です", "大阪", 50)
	fmt.Println()

	/*
	 【入力】
	 Scan: 標準入力に入力された値を読み取り
	 Scanln: 標準入力に入力された値を改行含めて読み取り
	 Scanf: 文字列にフォーマットし、標準出力に出力
	*/
	var name string
	fmt.Println("名前を入力してください")
	fmt.Scanln(&name) // 入力値を読み取り変数nameへバインド

	var age int
	fmt.Println("年齢を入力してください")
	fmt.Scanf("%d", &age) // 入力値を読み取り、書式を指定して変数ageへバインド

	fmt.Printf("名前： %s 年齢： %d", name, age)
	fmt.Println()

}
