package main

import (
	"fmt"
	"time"
)

// 日付・時刻のサンプルコードまとめ
func main() {

	// 現在時刻の取得
	t := time.Now()

	// 現在時刻を出力
	fmt.Println("【現在時刻】")
	fmt.Println(t)
	fmt.Println()

	// 現在時刻を yyyy/mm/dd hh:mm:ss の形で文字列にフォーマットして出力
	fmt.Println("【フォーマットした現在時刻】")
	fmt.Printf("%04d%02d%02d %02d:%02d:%02d\n",
		t.Year(),
		t.Month(),
		t.Day(),
		t.Hour(),
		t.Minute(),
		t.Second())
	fmt.Println()

	// Time型に時間を加算・減算する
	// 日時を加算（24:30:05後）
	add := t.Add(time.Hour*24 + time.Minute*30 + time.Second*5)
	fmt.Println("【時間加算後の時刻】")
	fmt.Println(add)
	fmt.Println()

	// 日時を減算（24:30:05前）
	sub := t.Add(-time.Hour*24 - time.Minute*30 - time.Second*5)
	fmt.Println("【時間減算後の時刻】")
	fmt.Println(sub)
	fmt.Println()

	// Time型に日付を加算・減算する
	// 日付を加算（1年と2日後）
	addDate := t.AddDate(1, 0, 2)
	fmt.Println("【日付加算後の時刻】")
	fmt.Println(addDate)
	fmt.Println()

	// 日付を減算（1年と2日前）
	subDate := t.AddDate(-1, 0, -2)
	fmt.Println("【日付減算後の時刻】")
	fmt.Println(subDate)
	fmt.Println()

	// 曜日の取得と出力
	fmt.Println("【今日の曜日】")
	switch t.Weekday() {
	case time.Sunday:
		fmt.Println("日曜日")

	case time.Monday:
		fmt.Println("月曜日")

	case time.Tuesday:
		fmt.Println("火曜日")

	case time.Wednesday:
		fmt.Println("水曜日")

	case time.Thursday:
		fmt.Println("木曜日")

	case time.Friday:
		fmt.Println("金曜日")

	case time.Saturday:
		fmt.Println("土曜日")
	}
}
