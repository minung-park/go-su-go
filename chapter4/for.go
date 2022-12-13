package main

import (
	"fmt"
)

func loop_for_1() {
	// standard
	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}
}

func loop_for_2() {
	i := 1
	// only conditional expression
	for i < 100 {
		fmt.Println(i)
		i = i * 2
	}
}

func loop_for_3() {
	i := 1
	for {
		fmt.Println("hello")
		if i > 10 {
			fmt.Println("bye!!!", i)
			break
		}
		i *= 2
	}
}

func continue_1() {
	for i := 1; i <= 100; i++ {
		if i%3 == 0 {
			if i%5 == 0 {
				fmt.Println("FizzBuzz")
			} else {
				fmt.Println("Fizz")
			}
		} else if i%5 == 0 {
			fmt.Println("Buzz")
		} else {
			fmt.Println(i)
		}
	}
}

func continue_2() {
	for i := 1; i <= 100; i++ {
		if i%3 == 0 && i%5 == 0 {
			fmt.Println("FizzBuzz")
			continue
		}
		if i%3 == 0 {
			fmt.Println("Fizz")
			continue
		}
		if i%5 == 0 {
			fmt.Println("Buzz")
			continue
		}
		fmt.Println(i)
	}
}

func for_range_1() {
	// python enumerate 와 비슷한 느낌?
	// for i, v in enumerate(evenVals):
	// 	print(i, v)
	evenVals := []int{2, 4, 6, 8, 10, 12}
	for i, v := range evenVals {
		fmt.Println(i, v)
	}
	fmt.Println("=======================")
	for _, v := range evenVals {
		fmt.Println(v)
	}
}

func for_range_map() {
	uniqueNames := map[string]bool{"Fred": true, "Raul": true, "Wilma": true}
	for k := range uniqueNames {
		fmt.Println(k)
	}
	fmt.Println("=======================")
	m := map[string]int{"a": 1, "b": 2, "c": 3}
	for i := 0; i < 3; i++ {
		fmt.Println("Loop: ", i)
		for k, v := range m {
			fmt.Println(k, v)
		}
	}
}

func for_range_string() {
	samples := []string{"hello", "apple_π!"}

outer: // label
	for _, sample := range samples {
		for i, r := range sample {
			fmt.Println(i, r, string(r))
			if r == 'l' {
				continue outer // label 로 빠져나감
			}
		}
	}
}

func main() {
	// fmt.Print("test1: ")
	// loop_for_1()
	// fmt.Print("test2: ")
	// loop_for_2()
	// fmt.Print("test3: ")
	// for i:=0; i<20; i++ {
	// 	for_range_1()
	// }
	for_range_string()
}
