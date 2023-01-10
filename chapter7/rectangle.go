package main

import (
	f "fmt"
	"time"
)

type rectangle struct { // 구조체 선언
	width, height int
	lastUpdated   time.Time
}

func main() {
	f.Println("========== method ==========")

	r := rectangle{width: 20, height: 5}
	f.Println("area : ", r.area())
	f.Println("perim: ", r.perim())
	f.Println("last updated at: ", r.lastUpdated)

	rp := &r
	rp.width = 5
	rp.height = 10

	f.Println("area : ", rp.area())
	f.Println("perim: ", rp.perim())
	f.Println("last updated at: ", rp.lastUpdated)
	// rp 는 r 을 참조하여 사용했기 때문에, r 을 호출하면 rp 의 값으로 변경됨
	// f.Println(r.width, r.height)
}

// method
func (r *rectangle) area() int {
	r.lastUpdated = time.Now()
	f.Println(r.lastUpdated, " ---> pointer receiver")
	return r.width * r.height
}

func (r rectangle) perim() int {
	r.lastUpdated = time.Now()
	f.Println(r.lastUpdated, " ---> value receiver")
	return 2*r.width + 2*r.height
}
