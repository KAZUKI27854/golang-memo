package main

import "fmt"

// 繰り返し処理まとめ
func main() {
	
	// for
	/* 
	 for 初期化文; 条件式; 後処理文 {
		 このブロックの処理が繰り返される
	 }
	*/

	// for文での繰り返し処理1
	fmt.Println("for文での繰り返し処理1")
	for i := 0; i < 5; i++ {
		fmt.Println(i)
	}

	// for文での繰り返し処理2
	fmt.Println("for文での繰り返し処理2")
	n := 0

	for n < 5 {
		fmt.Println(n)
		n++
	}

	// for文での繰り返し処理3
	fmt.Println("for文での繰り返し処理3")
	num := 0

	for {

		if num == 2 {
			num++

			// 条件式がtrueの時の処理をスキップ
			continue
		}
		
		fmt.Println(num)
		num++

		if num == 5 {

			// 繰り返し処理から抜ける
			break
		}
	}

	// rangeを使った繰り返し処理：　配列などの値から要素をひとつずつ取り出してループする
	/*
	 for 変数 := range 要素を取り出す値 {
		 このブロックの処理が繰り返される
	 }
	 */

	 // range式に使用する式の値の型によって戻り値の型や、数が異なるので注意!（P.72参照）
	 // 例えば文字列型を使った場合、1番目の戻り値が「文字列内のバイトインデックス（int型）」で、2番目の戻り値が「Unicodeコードポイント（rune型）」となるので、インデックスが連続した値にならない可能性もある

	 fmt.Println("rangeを使った繰り返し処理")
	 array := [...]string{"apple", "banana", "りんご", "バナナ"}

	 // 配列型を使った場合、1番目の戻り値が「配列内のインデックス（int型）」で、2番目の戻り値が「配列の要素型」となる
	 for index, value := range array {
		 fmt.Println(index, value)
	 }
}