package main

import (
	"fmt"
	"math/rand"
)


// 条件分岐まとめ
func main() {
	
	// if文による条件分岐1
	fmt.Println("if文による条件分岐1")

	for i := 0; i < 5; i++ {
		if i % 2 == 0 {
			fmt.Println(i, "は偶数")
		} else {
			fmt.Println(i, "は奇数")
		}
	}

	// if文による条件分岐2
	fmt.Println("if文による条件分岐2")

	// 0 ~ 100までのランダムな整数を返す
	x := rand.Intn(100)

	// 条件式の前に ; で区切って文を記述することが可能
	if y := rand.Intn(100); x < y {
		fmt.Println(x, "は", y, "より小さい")
	} else if x > y {
		fmt.Println(x, "は", y, "より大きい")
	} else {
		fmt.Println(x, "は", y, "と等しい")
	}

	// switch文による条件分岐
	fmt.Println("switch文による条件分岐")

	for i := 0; i < 5; i++ {
		
		// どれかのケースに該当した時点で暗黙的にbreak処理が行われる（明示的に記述することも可能）
		switch i {
		case 0:
		       fmt.Println("0")
		case 1, 2:
		       fmt.Println("1または2")
		case 3:
			   fallthrough  // 暗黙的なbreak処理を行わずに、すぐ下のcase節の処理を実行する
		default:
		       fmt.Println("その他")
		}

		// switch trueとすることで、case節の式に条件式が使えるようになり、利用範囲が広がる
		switch true {
		case i == 0:
			    fmt.Println("0")
		case i < 3:
			    fmt.Println("1または2") 
		default:
			    fmt.Println("その他")
		}
	}
}