package main

import "fmt"

/* 
  関数まとめ
  【書き方】
  func 関数名(パラメータリスト) 戻り値の型 {

	処理

	return 戻り値
  }
  ※ パラメータリストは「変数名」、「型」の順に , で区切って列挙する。戻り値がない場合は戻り値の型は省略可能。
*/

// main関数よりも前に実行したい処理がある場合はinit関数を使用する。初期化処理を記述することが多い。
func init() {
	fmt.Println("初期化処理")
}

// main関数はパラメータと戻り値どちらも持たない(プログラム実行時に自動的に呼び出される特別な関数)
func main() {

	// plus関数を呼び出す
	answer := plus(1, 2)

	fmt.Println(answer)

	// 複数の戻り値を返す関数を呼び出す
	add, sub, mul, div := calc(1, 2)
	fmt.Println(add, sub, mul, div)

	// 関数のパラメータリストが、別の関数の戻り値の型と一致する場合は直接渡すこともできる
	printFunc(calc(1, 2))

	// 可変長パラメータを持つ関数の呼び出し
	holiday(1, "元旦", "成人の日")
	holiday(2, "建国記念の日")
	holiday(3, "春分の日")

	// namedCalc関数を呼び出す
	namedAdd, namedSub := namedCalc(30, 20)
	fmt.Println(namedAdd, namedSub)


	/*
	 【関数リテラルまとめ】
	 ・関数名を持たない
	 ・他の関数内に記述する関数で、単体で宣言しない
	 ・関数リテラルは、関数リテラルの外側で宣言された変数にアクセスができる。これをクロージャという。
	 ・関数リテラルは記述と呼び出しを同時に行うか、一旦変数に代入してから、その変数を経由して呼び出しを行う。
	*/

	val := 123

	// 関数リテラルの記述と呼び出しを同時に行う
	func(i int) {

		// 関数リテラルの外側の変数valにアクセス可能
		fmt.Println(i * val)

	// 呼び出し箇所
	}(10)

	// 関数リテラルを変数に代入
	f := func(s string) {
		fmt.Println(s)
	}

	// 関数リテラルを変数経由で呼び出し
	f("関数リテラルが呼び出されました")


	/*
	 【関数型の変数まとめ】
	 ・関数を変数に格納するためには、関数型を使う
	 ・書式は以下のパターン
	     - func(パラメータリスト)
	     - func(パラメータリスト) 戻り値の型
		 - func(パラメータリスト) (戻り値のリスト)
	*/

	// 関数型の変数宣言
	var funcVar func(int, int) int

	// 関数型の変数に、関数リテラルの値を代入
	funcVar = func(a int, b int) int {
		return a + b
	}

	// 関数型の変数経由で関数を呼び出し1
	fmt.Println("関数型の変数経由での呼び出し結果1: ", funcVar(200, 300))

	// 関数型の変数に、関数（の値）を代入
	funcVar = plus

	// 関数型の変数経由で関数を呼び出し2
	fmt.Println("関数型の変数経由での呼び出し結果2: ", funcVar(2000, 3000))
}


// 関数1: 足し算を行う関数の宣言
func plus(a int, b int) int {
	return a + b
}

// 関数2: 同一関数内で複数の戻り値を同時に返す
func calc(x int, y int) (int, int, int, float32) {

	// 列挙した戻り値の型に対応する形で記述
	return x + y, x - y, x * y, float32(x) / float32(y)
}

// 関数3: 関数2の戻り値と同じ型のパラメータリストを持つ関数(戻り値なし)
func printFunc(num1 int, num2 int, num3 int, num4 float32) {
	fmt.Println(num1, num2, num3, num4)
}

// 関数4: 可変長パラメータを持つ関数
/*
 可変長パラメータ: 「days...string」のように記述すると、スライスとして値を受け取る可変長パラメータとなり、関数を呼び出す側でパラメータにいくつ値を渡すか決められるようになる。
                複数のパラメータを持つ関数では最後のパラメータしか可変長パラメータにすることはできない。
*/
func holiday(month int, days...string) {

	// printf関数（print関数については以下の記事参照）
	// https://qiita.com/rock619/items/14eb2b32f189514b5c3c
	fmt.Printf("%d月の祝日は%d日あります。\n", month, len(days))
	for _, day := range days {
		fmt.Println(day)
	}
	fmt.Println()
}

// 関数5: 戻り値に名前をつける
func namedCalc(a int, b int) (namedAdd int, namedSub int) {

	// 名前をつけた戻り値に値を入れる
	namedAdd = a + b
	namedSub = a - b

	// この場合もreturnは必要
	return
}