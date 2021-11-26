package main

import "fmt"

// 型まとめ
func main() {

	// 型宣言の書式：　type {宣言する型名} {基になる型}

	// int型を基に、myInteger型を宣言
	type myInteger int

	// myInteger型の変数を宣言
	var i myInteger = 123

	// ベースがint型なのでintと同様に演算が可能
	i += 1

	fmt.Println(i)

	// 新しい構造体を作成
	type myStruct struct {
		a int
		b int
	}



	// 型変換（Go言語では型のチェックが厳しいため、明示的に変換を使用しないとint32型とint64型のような型の異なる整数同士で比較や演算ができない）
	// https://go.dev/ref/spec#Numeric_types
	var n int = 1234

	// int型からuint32型への変換
	var u uint32 = uint32(n)

	// uint型からfloat32型への変換
	var f float32 = float32(u)

	// int型からstring型への変換
	// 整数値をstring型に変換したときは、その整数をUnicodeコードポイントとする文字を持つ文字列を取得する
	var s string = string(n)

	// string型からスライス型への変換
	var b []byte = []byte("abc")

	fmt.Println(n)
	fmt.Println(u)
	fmt.Println(f)
	fmt.Println(s)
	fmt.Println(b)

	/*
	型名が演算子で始まるときは注意が必要で、型名を () で囲まないと意図した通りに変換されない

	NG: *Point(p)
	OK: (*Point)(p)
	*/
}