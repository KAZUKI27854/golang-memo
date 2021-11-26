// 必ずpackageから始まる
package main

/*
複数行のコメントアウト
*/
import "fmt"

// 文字数を正しくカウントする 1/2
import "unicode/utf8"

func main() {

	// string型の文字列を宣言
	var en string

	// string型の変数に値を代入
	en = "golang"

	var ja string = "Go言語"

	// 文字列の長さを出力 => len組込み関数によるバイト長の出力
	fmt.Println(en, " 文字列の長さ（バイト長）：", len(en))
	fmt.Println(ja, " 文字列の長さ（バイト長）：", len(ja))

	/* Go言語の文字列はUTF-8エンコードされた「byte」配列のように振る舞うことができるので、
	「len」組み込み関数を使っての長さ（バイト長）の取得や、インデックスでのアクセスも可能。

	=> 文字数ではなくバイト長を返す点に注意！アルファベットや数字、日本語では1文字あたりのバイト数が違う
	*/

	// 文字数を正しくカウントするには「unicode/utf8」パッケージをインポートして「RuneCountInString」を使う
	// 文字数を正しくカウントする 2/2
	fmt.Println(ja, " 文字列の長さ（文字数）：", utf8.RuneCountInString(ja))
}