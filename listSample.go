package main

import (
	"container/list"
	"fmt"
)

// リストのサンプルコード
func main() {

	// 空のリストを作成
	l := list.New()

	// リストに要素を追加
	l.PushBack("a")
	l.PushBack("b")
	l.PushBack(3)

	// リストの要素を繰り返し出力
	for e := l.Front(); e != nil; e = e.Next() {

		// Value関数で要素を取得
		fmt.Println(e.Value)
	}

	// Len関数でリストの要素数を取得（ Lが大文字なので注意 ）
	fmt.Println("要素数: ", l.Len())

	/*
	 リストの要素の値を取得・変更する
	 => Go言語のリストは先頭か最後尾から順に要素をたどらないと特定の要素に到達しないため、
	    インデックスを使って要素へ直接アクセスすることはできない
	*/

	// リストから３番目（インデックス2）の要素を取り出す（getElementは自作関数）
	target := getElement(l, 2)
	fmt.Println("変更前： ", target.Value)

	// Value関数でアクセスし、リストの値を変更
	target.Value = "changed"
	fmt.Println("変更後： ", target.Value)
	fmt.Println()

	/* リスト同士の結合 */
	l2 := list.New()

	// リストに要素を追加
	for i := 0; i < 5; i++ {
		l2.PushBack(i)
	}

	// リストlとl2を結合
	l.PushBackList(l2)

	fmt.Println("【リストの結合結果】")
	for e := l.Front(); e != nil; e = e.Next() {
		fmt.Print(e.Value, " ")
	}
	fmt.Println()
	fmt.Println()

	/* リストを空にする */
	l.Init()
	fmt.Println("Init後の要素数: ", l.Len())
}

// リストの要素（Element）を取り出す関数
func getElement(l *list.List, index int) *list.Element {

	// リストの要素をインデックス付きで繰り返し取得
	for e, i := l.Front(), 0; e != nil; e = e.Next() {

		// パラメータに渡したインデックスの要素を返す
		if i == index {
			return e
		}

		i++
	}

	panic("インデックスが正しくありません")
}
