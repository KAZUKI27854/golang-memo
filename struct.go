package main

import "fmt"

// 構造体： フィールドと呼ばれる要素の集合体で、各フィールドはそれぞれ名前と型を持つ
func main() {

	// 構造体の変数を宣言（フィールド名の先頭の文字を小文字にすると他の関数などからアクセスできなくなる）
	var x struct {
		i1, i2 int  // int型フィールド
		f float32   // float32型フィールド
		s string    // string型フィールド
	}

	// 構造体の各フィールドに値をセット
	x.i1 = 1
	x.i2 = 2
	x.f = 3.14
	x.s = "go"

	fmt.Println(x)

	type embedded struct {
		i int
	}

	emb := new(embedded)
	emb.i = 10

	// 通常は、以下のように構造体にtypeでわかりやすい名前をつけて使用する（以下の構造体はMyData型となる）
	type MyData struct {
		s string
		_ string  // ブランクフィールド： メモリ領域としての確保は必要だが（パディング）、実際にそのフィールドにアクセスすることはないということを明示的に表すために使用する
		int       // 匿名フィールド： フィールド名を省略することで、型の名前がそのフィールドの名前になる
		embedded  // メソッドを埋め込み、継承と似たようなことを実現する
	}

	/* typeを用いて構造体型の変数を宣言した場合の初期化方法まとめ */
	// 構造体の初期化1: 構造体の各フィールドを個別に初期化
	myData := new(MyData)  // var myData MyDataも可能

	myData.s = "マイデータ"
	myData.int = 20
	myData.embedded.i = 10
	fmt.Println(myData)

	// 構造体の初期化2: 構造体リテラルで初期化（フィールド名と値のペアを記述）
	myData2 := MyData{s: "マイデータ2", int: 30}
	fmt.Println(myData2)

	// 構造体の初期化3: 構造体リテラルで初期化（フィールドの宣言順に値を記述）
	myData3 := MyData{"マイデータ3", "", 40, embedded{30}}
	fmt.Println(myData3)

	/* 
	 【構造体リテラル記述時の注意事項： 書き方によってコンパイルエラーになるため注意！】

	 OK (最後のデータと } が同じ行にあるのでOK)
	 Person{
		 age: 30,
		 name: "Tom"}

	NG (最後のデータと } が違う行にあるのでNG)
	 Person{
		 age: 30,
		 name: "Tom"
		}

	OK (最後のデータと } が違う行にある場合でも、最後のデータの後ろに , があればOK)
	 Person{
		 age: 30,
		 name: "Tom",
		}
	*/
}