package main

import "fmt"

func volume(length, width, height int) int {
	return length * width * height
}

type Dimension struct {
	length int
	width  int
	height int
}

func (d Dimension) Area() int {
	return d.length * d.width
}

func (d Dimension) Volume() int {
	return d.length * d.width * d.height
}

func dimensions(length, width, height int) (area int, volume int) {
	area = length * width
	volume = length * width * height
	return
}

func main() {
	area, vol := dimensions(5, 5, 5)
	vlm := volume(5, 5, 5)
	fmt.Println("area: ", area, "vol: ", vol, "volume: ", vlm)
	d := Dimension{10, 5, 6}
	fmt.Println(d.Area())
	fmt.Println(d.Volume())
}
