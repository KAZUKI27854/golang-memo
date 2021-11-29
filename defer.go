package main

import "fmt"

// 遅延実行： deferを使用して関数を呼び出すと、呼び出し元の関数から抜けたときに関数が実行される
func main() {

	fmt.Println("開始")     // 実行順：1

	// delay関数を遅延実行
	defer delay()          // 実行順：3

	fmt.Println("終了")     // 実行順：2

	/* 
	 【遅延実行のユースケース： ファイルのクローズなどのリソース解放に適している】

	 ファイルのオープン
	 file, err := os.OpenFile("test.txt", os.O_WRONLY|os.O_CREATE, 0666) 

	 オープンに失敗した時は終了
	 if err != nil {
		 os.Exit(1)
	 }

	 ファイルのクローズを遅延指定
	 defer file.close()

	 ファイルへテキスト出力
	 file.WriteString("あいうえお")
	*/
}

func delay() {
	fmt.Println("delayが呼び出されました")
}