package main

import "fmt"

/* 
 【スライスまとめ】
  ・配列と似ているが、スライス自体は要素データを持たない。配列全体のポインタ（ptr）、長さ（len）、容量（cap）のデータ構造からなる。
  ・スライスは配列と紐づけられており、スライスの要素を参照すると、実際にはそのスライスに紐づけられた配列の持つ要素が間接的に参照される
  ・長さはあらかじめ決まっておらず、動的に変わるので、配列と違い要素の追加が可能

  ・配列とどう見分けるのか：
      -- スライス: []要素型 の形で記述
	  -- 配列：    [要素数]要素型 or [...]要素型　の形で記述

  ・スライス式： 配列からスライスを作成する、もしくはスライスから新たにスライスを作成するために使用する
      -- 作成元の配列またはスライス[下限値： 上限値]               => 完全スライス式
	  -- 作成元の配列またはスライス[下限値： 上限値: キャパシティ]   => 簡易スライス式

  ・スライスは参照型で、データの実体を持たずデータへの参照情報を持たない。Go言語の中では参照型は「スライス」「マップ」「チャネル」の3種類ある。
    => 配列の場合、配列の値を関数間で受け渡すと受け渡し先で配列の要素が全てコピーされ、メモリ消費が大きくなる。スライスでは、参照情報のみがコピーされるので、
	   長さやデータ量に左右されないというメリットがある。また、スライスの要素を呼び出し先の関数内で変更すると、実際には参照先である変数の要素が変更されるので、
	   呼び出し元にも変更が反映される。
	   つまりスライスから渡されているパラメータは値そのものではなく、値のメモリアドレスである。
	
  ・スライスを作るためだけに配列を作るのは手間なので、配列とそれを参照するスライスを同時に作成するmake関数を使う。
      -- make(スライス型, 長さ, キャパシティ)
	  -- make(スライス型, 長さ)

  ・スライスを初期化する方法として、スライスリテラルを使うことで暗黙的に配列が作成される。
      -- スライス型 {要素の初期値1, 要素の初期値2, ...}

   参照URL:  (※2015年の記事)https://qiita.com/mizukmb/items/b543f88bc37c9a75673f
           （スライスの基本的な使い方）https://golang.hateblo.jp/entry/golang-slice-array
		    (capなど、スライスで間違いやすい箇所) https://qiita.com/Kashiwara/items/e621a4ad8ec00974f025
*/
func main() {

	// 配列を宣言
	x := [5]string{"a", "b", "c", "d", "e"}

	// スライス型の変数を宣言
	var s1 []string

	// 配列全体をスライス（スライス： 配列やスライスの要素を指定してスライスに代入すること）
	s1 = x[:]
	fmt.Println(s1)   // a, b, c, d, e

	// インデックス1~3までをスライス
	s2 := x[1:4]
	fmt.Println(s2)   // b, c, d

	// インデックス3~をスライス
	s3 := x[3:]
	fmt.Println(s3)   // d, e

	// インデックス~3をスライス
	s4 := x[:4]
	fmt.Println(s4)   // a, b, c, d

	values := [...]int{0, 1, 2, 3, 4}

	// スライスして関数に受け渡す
	double(values[:])

	// double関数での変更が反映されている
	fmt.Println(values)
	fmt.Println()




	/* 
	 cap関数
	 => スライスの容量（メモリ領域）を確認する関数
	*/
	numbers := [...]int{0, 1, 2, 3, 4}

	// 配列全体をスライス
	n1 := numbers[:]

	fmt.Println("全体をスライス: ", n1)   // 0, 1, 2, 3, 4
	fmt.Println("len: ", len(n1))      // 5
	fmt.Println("cap: ", cap(n1))      // 5

	// インデックス1~3までをスライス
	n2 := numbers[1:4]

	fmt.Println("インデックス1~3までをスライス: ", n2)  // 1, 2, 3
	fmt.Println("len: ", len(n2))                  // 3
	fmt.Println("cap: ", cap(n2))                  // 4

	// インデックス3までをスライス
	/*
	 【注意！！】
	 : の左側が０、または省略した場合は元の配列分の容量を確保する
	*/
	n3 := numbers[:4]

	fmt.Println("インデックス3までをスライス: ", n3)    // 0, 1, 2, 3
	fmt.Println("len: ", len(n3))                  // 4
	fmt.Println("cap: ", cap(n3))                  // 5

	// 完全スライス式（上記の記述は簡易スライス式で、完全スライス式ではキャパシティを指定できる）
	// [下限値： 上限値: キャパシティ]
	n4 := numbers[1:3:4]

	fmt.Println("インデックス1~2までをスライス: ", n4)  // 1, 2
	fmt.Println("len: ", len(n4))                  // 2
	fmt.Println("cap: ", cap(n4))                  // 3
	fmt.Println()




	/* 
	 文字列のスライス
	 => 文字列をスライスしたときに返される値は文字列だが、日本語のように1文字単位で
	 　　複数バイトを使用する文字の場合は正しく文字列の切り出しができない可能性もある点に注意が必要。
	*/

	// 文字列のスライス
	var str1 string = "abcde"[1:4]
	fmt.Println(str1)

	// 平仮名はUTF-8で3バイトなので「いう」が取り出せる
	var str2 string = "あいうえお"[3:9]
	fmt.Println(str2)

	// UTF-8の文字の境界として不適切な値を指定すると文字化けする
	var str3 string = "あいうえお"[1:4]
	fmt.Println(str3)
	fmt.Println()



	/* 
	 append関数
	 => スライスへ要素を追加する関数で、戻り値としてスライスを返す。
	    追加先のスライスに充分なキャパシティが無い場合、元の配列の容量と同じ容量の配列が追加される（2倍になる）ため、
		返されるスライスが参照している配列が、元のスライスが参照していた配列と異なる可能性があるという点に注意が必要。
	
	 【書き方】
	  append(追加先のスライス, 追加する要素1, 追加する要素2)

	  append(追加先のスライス, 追加するスライス...)  // ... は必須
	*/

	// スライスを宣言
	sl := []int{1, 2, 3, 4}

	// スライスに要素を追加
	sl2 := append(sl, 5, 6)

	// スライスにスライスを追加
	sl3 := append(sl2, sl...)  // ... は必須

	fmt.Println(sl3)           // [1 2 3 4 5 6 1 2 3 4]

	// appendの特殊な使い方
	var b1 []byte
	b2 := append(b1, "abc"...) // byteスライスに文字列を追加できる
	fmt.Println(b2)            // [97 98 99]





	/*
	 コピー関数
	 => スライスの要素を別の要素にコピーする関数で、戻り値としてコピーされた要素の数を返す。

	 【書き方】
	 copy(コピー先スライス, コピー元スライス)
	*/

	dst := []int{1, 2, 3, 4}
	src := []int{5, 6, 7}

	count := copy(dst[2:], src)

	fmt.Println(dst)  // [1 2 5 6]
	fmt.Println("コピーされた要素数: ", count)  // 2





	/*
	 スライスを可変長パラメータに直接渡す場合
	 => 関数の可変長パラメータに渡された値は、呼び出された関数内ではスライスとして受け取られる。
	    つまり、新しいスライスが暗黙的に作成され、可変長パラメータとして与えられた各値がそのスライスの要素となる。
		そのため、可変長パラメータにスライスをそのまま渡しても、新規にスライスが作成されてしまい、スライスを格納した
		スライスが渡されることになるので、加工せずに渡すためにはスライスの後ろに ... を記述する必要がある。
	*/

	s := []string{"a", "b", "c"}

	// スライスをそのまま渡すには ... をつける
	test(s...)    // 渡される値は test("a", "b", "c")と書いた時と同じ





	/*
	 make関数
	 => スライスを作るためだけに配列を作るのは手間なので、配列とそれを参照するスライスを同時に作成するmake関数を使う。

	 【書き方】
      make(スライス型, 長さ, キャパシティ)
	  make(スライス型, 長さ)
	*/

	// 長さ10, キャパシティ20のスライスを作成
	test1 := make([]int, 10, 20)

	fmt.Println("make関数でスライス作成1: ", test1) 
	fmt.Println("len: ", len(test1))
	fmt.Println("cap: ", cap(test1))
	fmt.Println()

	// 長さ, キャパシティ共に5のスライスを作成
	test2 := make([]float32, 5)

	fmt.Println("make関数でスライス作成2: ", test2) 
	fmt.Println("len: ", len(test2))
	fmt.Println("cap: ", cap(test2))
	fmt.Println()





	/*
	 スライスリテラル
	 => スライスを初期化する方法として、スライスリテラルを使うことで暗黙的に配列が作成される。

	 【書き方】
      スライス型 {要素の初期値1, 要素の初期値2, ...}
	*/

	// スライスリテラルでスライスを作成
	ltr := []int{1, 2, 3, 4}

	fmt.Println(ltr)  // [1 2 3 4]
}



// スライスの要素の値を倍にする関数
func double(values []int) {

	/*
	 rangeでスライスの要素にアクセスしようとすると、値のコピーが返ってくるため元の値は変更されない
	 => パラメータがintのポインタのスライス（[]*int）であればrangeでも値の変更、反映ができる
	 参照URL: https://inabajunmr.hatenablog.com/entry/2019/11/07/093217

	 for _, value := range values {
	 	value *= 2
	 }
	*/

	// 以下の形であればポインタではない値のスライスでも元の値を変更できる（ポインタのスライスでも同様に変更できる）
	for i := 0; i < len(values); i++ {
		values[i] *= 2
	}
}


// 可変長パラメータを持つ関数
func test(s...string) {

	// スライスの長さと値を出力
	fmt.Println(len(s), s)
	fmt.Println()
}