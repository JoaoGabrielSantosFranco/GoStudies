package main

import "fmt"

type Rectangle struct {
	width, heigth float64
}

func (r Rectangle) area() float64 {
	return r.width * r.heigth
}

func main() {

	r := Rectangle{width: 5, heigth: 5}
	fmt.Print("area = ", r.area())

}
