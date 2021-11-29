package main

import "fmt"

// ラベル： breakやcontinue、gotoの宛先となる識別子
func main() {

// 宛先となるラベル
LABEL:
    
    for i := 0; i < 5; i++ {
		switch {
		case i == 3:

			// break後の行き先を指定（continueの場合も同様）
			break LABEL

		default:
			fmt.Println(i)  // 出力結果： 0 1 2
		}
	}

	// gotoの場合
	for n := 0; n < 10; n++ {
		if n % 2 == 0 {

			// 行き先を指定
			goto EVEN
		}
		fmt.Println(n)  // 出力結果： 1 3 5 7 9
	// ラベル
	EVEN:
	}
}