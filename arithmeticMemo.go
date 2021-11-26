package main

import "fmt"

// 算術演算子まとめ（型が違うもの同士は計算できないため変換必須）
func main() {

	// 数値の足し算
	fmt.Println("1 + 2 =", 1+2)

	// 文字列の足し算
	fmt.Println("\"abc\" + \"XYZ\" =", "abc"+"XYZ")

	// 引き算
	fmt.Println("3 - 2 =", 3-2)

	// 掛け算
	fmt.Println("2 * 3 =", 2*3)

	// 割り算(答えは2.5だがゼロに近づけ丸められるので2)
	fmt.Println("5 / 2 =", 5/2)

	// 割り算(答えは-2.5だがゼロに近づけ丸められるので-2)
	fmt.Println("-5 / 2 =", -5/2)

	// 余り
	fmt.Println("5 % 2 =", 5%2)

	// ビット演算(and)
	fmt.Println("3 & 6 =", 3&6)

	// ビット演算(or)
	fmt.Println("3 | 6 =", 3|6)

	// ビットクリア(and not)
	fmt.Println("3 &^ 6 =", 3&^6)

	// 左シフト演算
	fmt.Println("2 << 1 =", 2<<1)

	// 右シフト演算
	fmt.Println("2 >> 1 =", 2>>1)
}
