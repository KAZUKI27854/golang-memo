package main

import "fmt"

/*
 【配列まとめ】
 ・ [長さ]要素型 の形で記述し、後から長さを変更することはできない
 ・　配列[インデックス] の形で配列の要素へアクセスできる
*/
func main() {

	// 配列型変数の宣言
	var array1 [1]byte                 // 長さが1, 要素型がbyteである配列
	var array2 [5]*int                 // ポインタの配列
	var array3 [8][3]int64             // 2次元配列
	var array4 [2]struct{ x, y int }   // 構造体の配列

	// len組込み関数で配列の長さを取得する
	fmt.Println(len(array1))
	fmt.Println(len(array2))
	fmt.Println(len(array3), len(array3[0])) // 2次元目の長さも出力
	fmt.Println(len(array4))

	var direction [4]string

	// インデックスを使用して要素に値を代入
	direction[0] = "上"
	direction[1] = "下"
	direction[2] = "右"
	direction[3] = "左"

	// 配列の要素を繰り返し出力1
	for i := 0; i < len(direction); i++ {
		fmt.Print(direction[i], " ")
	}

	fmt.Println()

	// 配列の要素を繰り返し出力2（rangeから返される1つ目の値はインデックスだが、使用しないのでブランク識別子に代入）
	for _, value := range direction {
		fmt.Print(value, " ")
	}

	fmt.Println()

	// 配列リテラルによる配列の初期化
	// 長さ5の配列で、各要素はゼロ値で初期化される
	arrayOne := [5]float32{}

	// 長さ6の配列で、要素の不足分はゼロ値で初期化される
	arrayTwo := [6]int{1, 2, 3, 4}

	// 長さに ... と記述すると要素数が長さとして使用される
	arrayThree := [...]string{"One", "Two", "Three"}

	// 出力
	fmt.Println(arrayOne)
	fmt.Println(arrayTwo)
	fmt.Println(arrayThree)
}