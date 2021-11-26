package main

import "fmt"

// ポインタ： 変数が存在するメモリ上の場所をアドレスと言い、そのアドレスが格納可能な変数をポインタと呼ぶ
func main() {
	
	// int型のポインタ変数を宣言: int型のポインタ変数とは、int型変数のアドレスを格納できるポインタという意味
	var pointerVariable *int

	// int型の変数を宣言
	var number int = 12345

	// アドレス演算子(&)： 変数からアドレスを取得するときに用いる
	// アドレス演算子を使い、numberのアドレスをpointerVariableに代入
	pointerVariable = &number

	fmt.Println("numberのアドレス： ", &number)
	fmt.Println("pointerVariableの値（変数numberのアドレス）： ", pointerVariable)

	fmt.Println("numberの値： ", number)

	// 間接参照演算子(*)： アドレスが指し示す変数にアクセスするときに用いる
	// 間接参照演算子を使い、アドレスからnumberを参照
	fmt.Println("ポインタ経由のnumberの値： ", *pointerVariable)

	// ポインタ経由でnumberの値を変更
	*pointerVariable = 999

	fmt.Println("ポインタ経由で変更したnumberの値： ", number)

	/*
	 値渡し： 変数を関数に渡すと実際は値のコピーが渡されるので、呼び出した関数内で値を変更しても関数の呼び出し元の値には一切影響がない。
	 ポインタ渡し（参照渡し）： 値のポインタを関数に渡した時は、実際に保存されているメモリ上のデータを参照するため、呼び出した関数内で呼び出し元の値を変更することができる。
	*/

	a, b := 100, 100

	// aは値渡しで関数へ渡す
	// bはポインタ渡しで関数へ渡す
	double(a, &b)

	// 値渡しとポインタ渡しで、実際の値が変更されているか比較
	fmt.Println("値渡し： ", a)
	fmt.Println("ポインタ渡し： ", b)

	// new組込み関数を使って、明示的に新しいメモリを割り当てることも可能
	// var i *int = new(int)
}

func double(x int, y *int) {

	// 変数の値を変更
	x *= 2

	// 間接参照演算子を使用し、ポインタyが指し示す変数の値を変更
	*y *= 2
}