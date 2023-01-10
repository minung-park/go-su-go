package main

import f "fmt"

type test interface { // 인터페이스 정의
}

func main() {
	var i test // 인터페이스 선언
	f.Println(i)
}
