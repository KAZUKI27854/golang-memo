package main

import "fmt"

/*
 メソッド: 関数の一種だが、メソッドの型となるレシーバを持ち、特定の型と紐づけられる点が異なる。
 【書き方】
 func (レシーバの変数名 レシーバの型) メソッド名(パラメータリスト)　戻り値の型 {

	処理

	return 戻り値
 }
*/

// int型から新たにmyType型を宣言
type myType int

// myType型をレシーバに持つ関数を宣言（厳密には、レシーバの値を出力するmyType型のメソッド）
func (value myType) println() {

	fmt.Println(value)
}


/* 
　レシーバ変数の値をメソッド内で変更するには、レシーバをポインタにする必要がある。
　また、はじめからポインタ型のレシーバ変数を使うことはできない。
*/
// レシーバが値（非ポインタ）のメソッド
func (value myType) setByValue(newValue myType) {

	// ポインタでは無いため、値のコピーを変更しているだけで、実際の値は変更されていない
	value = newValue
}

// レシーバがポインタのメソッド
func (value *myType) setByPointer(newValue myType) {

	// 代入した値へと変更される
	*value = newValue
}


// 値を加算するメソッド(メソッド値で使用)
func (value *myType) add(increment myType) myType {
	*value += increment
	return *value
}


func main() {

	// myType型の変数を宣言
	var z myType = 1234

	// myType型のメソッド呼び出し: メソッドを持つ型の値の後ろに . を記述し、続けてメソッド名を記述する。
	z.println() // 1234

	// このメソッドのレシーバは値なので、値は変更できない
	z.setByValue(1)
	z.println() // 1234

	// このメソッドのレシーバはポインタなので、値は変更できる
	z.setByPointer(2)
	z.println() // 2


	// メソッド値： 「z.println」のような、値とメソッドの紐付けを扱う値
	// 変数zに対してaddメソッドを呼び出し
	fmt.Println(z.add(3))

	// メソッド値を取得。このメソッドの値の型は「func(myType) myType」
	mv := z.add

	// メソッド値に対してaddメソッドを呼び出し
	fmt.Println(mv(3))
}