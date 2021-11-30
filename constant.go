package main

import "fmt"

// 定数：　コンパイル時に値が決定し、プログラム内では値が変更できない値のこと。定数となり得るのは、数値・文字列・論理値だけ。
func main() {

	// 定数宣言は変数の時とほぼ同じ（varだった箇所がconstに変わるくらい）

	// 複数の定数を一度に宣言
	const a, b, c int = 1, 2, 3
	fmt.Println(a, b, c)

	// 型の省略も可能（ ＝の右側の値か式によって型が決まる）
	const d = a
	fmt.Println(d)

	/*
	 宣言時に型が決まらないケース：　整数リテラルは型を持たない為、使用される文によって明示的・暗黙的に型がきまる
	 ※ この場合、整数リテラルが表す値より小さいサイズの変数へ代入しようとするとオーバーフローを起こし、コンパイルエラーとなるので注意
	*/
	const e = 1234567890 // この時点では型を持たない

	// int64型へは代入可能
	var i int64 = e
	fmt.Println(i)

	// オーバーフロー（コンパイルエラー）
	// var b byte = c

	// 丸括弧を使ったグルーピング
	const (
		ja string = "日本語"
		en string = "English"
	)
	fmt.Println(ja, en)

	// グルーピング時、値を省略すると前の行と同じ値になる
	const (
		X int = 1 // X = 1
		Y         // Y = 1
		Z         // Z = 1
	)
	fmt.Println("X = ", X)
	fmt.Println("Y = ", Y)
	fmt.Println("Z = ", Z)

	// iotaを使うと0からインクリメントするように定義できる
	const (
		ZERO = iota
		ONE
		TWO
		THREE
	)
	fmt.Println("iotaその1: ", ZERO, ONE, TWO, THREE)

	// 式の途中でiotaを使うことも可能（以下のケースでは2倍されていく）
	const (
		FIVE = 5 << iota
		TEN
		TWENTY
		FORTY
	)
	fmt.Println("iotaその2: ", FIVE, TEN, TWENTY, FORTY)
}
