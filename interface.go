package main

import "fmt"

/* 
   インターフェース： 中身（ロジック）を持たない関数。例えば、「2つの整数の計算した結果を返す」という処理がある場合に、
   どのような計算を行うかがロジックにあたる箇所で、「2つの整数の計算した結果を返す」という表面的な振る舞いがインターフェースにあたる

   【書き方】
   interface {
	   関数名（パラメータリスト）戻り値の型
   }
*/

/*
 基本的には、以下のようにtypeを使ってインターフェース型として宣言する
 関数を1つしか持たないインターフェースの場合、インターフェース名を「関数名 + er」とするのが慣わし（ex: Reader)
*/
type Calculator interface {

	// 関数の定義（ロジックは持たない）
	Calculate(a int, b int) int
}

// 足し算型
type Add struct {
	// フィールドは持たない
}

// Add型にCalculatorインターフェースのCalculate関数を実装
func (x Add) Calculate(a int, b int) int {

	// 足し算
	return a + b
}

// 引き算型
type Sub struct {
	// フィールドは持たない
}

// Add型にCalculatorインターフェースのCalculate関数を実装
func (x Sub) Calculate(a int, b int) int {

	// 引き算
	return a - b
}

func main() {

	// Calculatorインターフェースを実装した型の変数を宣言
	var add Add
	var sub Sub

	// Calculatorインターフェース型の変数を宣言
	var cal Calculator

	// Add型の値を代入
	cal = add

	// インターフェース経由でメソッドを呼び出す
	fmt.Println("和：", cal.Calculate(1, 2))

	// Sub型の値を代入
	cal = sub

	// インターフェース経由でメソッドを呼び出す
	fmt.Println("差：", cal.Calculate(1, 2))


	/*
	 型アサーション： インターフェース型の値を他の型へ変換する時、または変換できるか確認する時に使う
	   => 型の変換先となり得るのは、インターフェースに格納されている値が持つ型と同一か、もしくはインターフェースに
	    　格納されている値の型が実装しているインターフェースだけ
	 */

	 // 空インターフェースに、string型の値を格納
	 var i interface{} = "test"

	 // 型アサーションを使いstring型へ
	 var s string = i.(string)

	 // 出力中
	 fmt.Println(s)


	 // 型switch文による判定
	 showType(nil)
	 showType(12345)
	 showType("abcdef")
	 showType(3.14)
	 
}

/* 
	型switch文： 「式switch文」が「値」によって分岐するのに対し、「型switch文」は値が持つ「型」によって分岐する
				case節には型だけでなく、nilを指定することもできるが、fallthrough文は使用できないので注意
*/

// 型を判定する関数
func showType(x interface{}) {

	// 型switch
	switch x.(type) {

	// nilも使用可能
	case nil:
		fmt.Println("nil")

	// 整数
	case int, int32, int64:
		fmt.Println("整数")

	// 文字列
	case string:
		fmt.Println("文字列")

	// その他
	default:
		fmt.Println("不明")
	}
}