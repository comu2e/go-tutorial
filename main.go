package main

import "fmt"

func main() {
	//print
	fmt.Println("Hello")
	fmt.Println("Good")
	//変数の設定
	name := 23;
	fmt.Println(name)
	//条件分岐
	if name == 2 {
		fmt.Println("True");
	//elseは改行したらエラー
	} else {fmt.Println("False")}

	//配列
	a := [2][2]string{{"satp","suzuki"},{"takahashi","tanaka"}}

	fmt.Println(a[1][0])

	age := 22

	if age >= 20{
		fmt.Println("adult")
	}else if age == 0{
		fmt.Println("baby")
	}else{
		fmt.Println("child")
	}

}
