package main

import "fmt"
import "os"

/*
 【戻り値と自作関数によるエラーハンドリング】※パニックとリカバリによるエラーハンドリングはpanic.go参照
 ・Go言語には、「errorインターフェース型」という型が組み込まれているため、この型をエラーハンドリングに利用する。
  => 例えば、errorインターフェース型が持つErrorメソッドを利用することでエラーの詳細情報が取得できる。
*/


/*
 自作関数を使ったエラー処理（errorインターフェース型を利用）
*/

// MyError型の構造体
type MyError struct {

	// エラー詳細のフィールド
	message string
}

// Errorメソッドの実装（レシーバにMyError型の値を渡しているためMyError型と紐づくメソッドになる）
func (err MyError) ErrorMessage() string {

	// messageフィールドに格納されたエラーメッセージを返す
	return err.message
}



func main() {

	// ファイルのオープン
	file, error := os.Open("test.txt")

	// オープンに失敗したか判定
	if error != nil {

		// エラーの詳細情報を出力
		fmt.Println(error.Error())

		// 自作関数を使ったエラー内容の出力
		myError := MyError{"ファイルのオープンに失敗しました"}
		fmt.Println(myError.ErrorMessage())

		// 終了
		os.Exit(1)
	}

	// ファイルのクローズ
	file.Close()

	fmt.Println("OK")
}