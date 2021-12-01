package main

import (
	"fmt"
	"strings"
)

// 文字列のサンプルコード
func main() {

	/*
	 【文字列の比較】
	 文字列を比較する際は、比較演算子を使う。不等号により大小を比較した際は、Unicodeコードポイント順で比較される。
	*/
	s1 := "alpha"
	s2 := "beta"

	if s1 == s2 {
		fmt.Println("同じ")
	}

	if s1 != s2 {
		fmt.Println("異なる")
	}

	if s1 < s2 {
		fmt.Println("s1の方が小さい")
	}

	if s1 > s2 {
		fmt.Println("s2の方が小さい")
	}
	fmt.Println()

	/*
	 【文字列の結合】
	 文字列を結合するには + 演算子または += 演算子を使用する。
	 文字列スライスを1つの文字列に結合するにはstringsパッケージのJoin関数を使用する
	*/
	s2 = s2 + "gamma"
	s1 += s2
	fmt.Println("【結合結果】", s1)
	fmt.Println()

	// 文字列スライスの用意
	arr := []string{"alpha", "beta", "gamma"}

	// 文字列スライスの結合（カンマで結合）
	fmt.Println("【Join関数での結合結果】", strings.Join(arr, ", "))
	fmt.Println()

	/*
	 【文字列の部分取得】
	 スライス式を使って部分的に文字列から取り出す（インデックスではなくバイト単位での取り出しである点に注意）
	*/
	str := "あいうえおかきくけこさしすせそ"
	fmt.Println(str[:12])
	fmt.Println(str[3:10]) // 文字の途中でバイト取り出しが終わるため、文字化けする
	fmt.Println(str[3:])
	fmt.Println(str[:])
	fmt.Println()

	/*
	 【文字列のトリム】
	 文字列のトリムにはstringsパッケージのTrimSpace関数を使用する。

	 先頭か末尾のどちらかだけをトリムするときはTrimLeft関数かTrimRight関数を使う。
	 => この場合は、取り除く対象の文字列を第2パラメータで指定する必要がある。
	*/
	strWithSpace := " xyz "

	// 前後トリム
	fmt.Println("【前後トリム】")
	fmt.Printf("[%s]\n", strings.TrimSpace(strWithSpace))
	fmt.Println()

	// 左右どちらかだけトリム（取り除く文字列を第2パラメータで指定）
	fmt.Println("【左トリム】")
	fmt.Printf("[%s]\n", strings.TrimLeft(strWithSpace, " "))
	fmt.Println()

	fmt.Println("【右トリム】")
	fmt.Printf("[%s]\n", strings.TrimRight(strWithSpace, " "))
	fmt.Println()

	/* 【文字列を大文字・小文字にする】 */
	greeting := "Hello"

	// 大文字に変換
	fmt.Println("【大文字変換】")
	fmt.Println(strings.ToUpper(greeting))
	fmt.Println()

	// 小文字に変換
	fmt.Println("【小文字変換】")
	fmt.Println(strings.ToLower(greeting))
	fmt.Println()

	/* 【文字列を分割する】 */
	numberStr := "one, two, three, four, five"

	// カンマで分割
	result := strings.Split(numberStr, ",")

	fmt.Println("【文字列の分割】")
	fmt.Println(result)
	fmt.Println()

	/* 【指定した文字の位置を取得する】 */
	message := "The Go Programming Language"

	// 文字列内の「Go」の位置（0~）を取得する
	fmt.Println("【文字列内の「Go」の位置（インデックス）】")
	fmt.Println(strings.Index(message, "Go"))

	// 見つからない時は-1が返される
	fmt.Println(strings.Index(message, "abc"))

}
